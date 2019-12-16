/*
Copyright 2015 - Jaume Arús

Author Jaume Arús - jaumearus@gmail.com

Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/

package lib_gc_messages_collector

import (
	"bytes"
	"expvar"
	"fmt"
	"sync"
	"time"

	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_fifo"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_runner"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_timeout_wrapper"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_config"

	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	EVENT_PUBLISHER "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event_publisher"
)

func init() {

	messagesPackageMaxSize = int32(lib_gc_config.PtrConfigProvider.GetPropertyINT64("LIB_MESSAGES_COLLECTOR", "PACKAGE_MAX_SIZE"))
	messagesPackageMaxWaitingTime = lib_gc_config.PtrConfigProvider.GetPropertyINT64("LIB_MESSAGES_COLLECTOR", "PACKAGE_MAX_WAITING_TIME")
	messagesPackageCheckTime = lib_gc_config.PtrConfigProvider.GetPropertyINT64("LIB_MESSAGES_COLLECTOR", "PACKAGE_CHECK_INTERVAL")
	messagesPackageMaxSendingRetries = int8(lib_gc_config.PtrConfigProvider.GetPropertyINT("LIB_MESSAGES_COLLECTOR", "MAX_SENDING_RETRIES"))
	messagesPackageSendingChanSize = int16(lib_gc_config.PtrConfigProvider.GetPropertyINT("LIB_MESSAGES_COLLECTOR", "SENDING_RETRIES_CHAN_SIZE"))
	EchoSendingErrors = lib_gc_config.PtrConfigProvider.GetPropertyBOOL("LIB_MESSAGES_COLLECTOR", "ECHO_SENDING_ERRORS")

	// Initalizing the messages channels
	MessagesCollectorChan = make(chan *Message, messagesPackageMaxSize+1)
	packagesToBeSentChan = make(chan *Package, messagesPackageSendingChanSize)
	packagesToBeRetiedChan = make(chan *Package, messagesPackageSendingChanSize*10)

	// Initializing maps
	retriesPerPackage = make(map[int64]int8)

	// If EchoSendingErrors flag is set to true, each sending error will be sent to event publisher channel
	// Depends on the number of retries set by configuration, one package can be tried to be sent several times.
	// Each of this times, if fails, it will sent a error message to the sending errors channel
	if EchoSendingErrors {
		SendingErrorsChan = make(chan *Package, messagesPackageSendingChanSize)
		go func() {
			var pkg *Package
			var buff bytes.Buffer
			for {
				pkg = <-SendingErrorsChan
				fmt.Printf("*jas Received from SendingErrorsChan: %#v\n", pkg)
				for k, _ := range pkg.SendingErrors {
					buff.WriteString(fmt.Sprintf("[retry: %d, error: %s]", pkg.SendingErrors[k], k))
				}
				msg, _ := EVENT.NotifyEvent("051-001", "", &[]string{buff.String()})
				EVENT_PUBLISHER.EventPublisherChannel <- msg
			}
		}()
	}

	// Initialize the mutexes
	mutex = &sync.Mutex{}
	mutexRetriesPerPackage = &sync.Mutex{}

	// Initialize the package of messages
	initializeLot()

	// Taking the messages collector runner
	var err error
	if runner, err = lib_gc_runner.GetRunner(100); err != nil {
		panic(err)
	} else {
		//Starting the messages collector runner
		if err = runner.Start(); err != nil {
			panic(err)
		}
	}

	// Fire up the packages sending loop
	if task, err = lib_gc_runner.GetTask(1, time.Duration(0), sendPackagesLoop); err != nil {
		panic(err)
	} else {
		// Waking up the task
		if err = runner.WakeUpTask(task); err != nil {
			panic(err)
		}
	}
	EVENT_PUBLISHER.EventPublisherChannel <- fmt.Sprintf("Messages package sending loop started. Sending errors %t", EchoSendingErrors)

	// Fire up the packages sending checking loop
	if task, err = lib_gc_runner.GetTask(2, time.Duration(0), checkLoop); err != nil {
		panic(err)
	} else {
		// Waking up the task
		if err = runner.WakeUpTask(task); err != nil {
			panic(err)
		}
	}
	EVENT_PUBLISHER.EventPublisherChannel <- "Messages package checking loop started !!!"

	// Fire up the packages retrying checking loop
	if task, err = lib_gc_runner.GetTask(22, time.Duration(0), retryingPackagesLoop); err != nil {
		panic(err)
	} else {
		// Waking up the task
		if err = runner.WakeUpTask(task); err != nil {
			panic(err)
		}
	}
	EVENT_PUBLISHER.EventPublisherChannel <- "Messages package retrying loop started !!!"

	// Fire up the main loop
	if task, err = lib_gc_runner.GetTask(3, time.Duration(0), mainLoop); err != nil {
		panic(err)
	} else {
		// Waking up the task
		if err = runner.WakeUpTask(task); err != nil {
			panic(err)
		}
	}
	EVENT_PUBLISHER.EventPublisherChannel <- "Messages collector main loop started !!!"

	EVENT_PUBLISHER.EventPublisherChannel <- "Messages sending collector started !!!"
}

// Initializing counters
var CURRENT_ON_BUILDING_PACKAGE_LENGTH = expvar.NewInt("CURRENT_ON_BUILDING_PACKAGE_LENGTH")
var PENDING_RETRYING_PAKCAGES = expvar.NewInt("PENDING_RETRYING_PAKCAGES")
var PACKAGES_SENDING_ERRORS = expvar.NewInt("PACKAGES_SENDING_ERRORS")
var DISCARDED_MESSAGES = expvar.NewInt("DISCARDED_MESSAGES")


var TIMEOUT_SENDING_CHECKING = 1000
var TIMEOUT_PACKAGE_CLOSING = 1000
var TIMEOUT_PUTTING_MESSAGE = 10000
var TIMEOUT_INITIALIZING_LOT = 1000

// runner task (this component ensures that the task is always up and running)
var runner lib_gc_runner.Runner_I
var task lib_gc_runner.Task_I

type Message struct {
	Id             int64
	CreationTime   int64
	Ttl            time.Duration
	SendingRetries int16
	Content        interface{}
	Sender         string
}

type Package struct {
	Sender        string
	Queue         *lib_gc_fifo.Fifo
	Id            int64
	SendingErrors []error
}

// Messages collector chan
var MessagesCollectorChan chan *Message

// Packages to be sent chan
var packagesToBeSentChan chan *Package

// Sending's errors chan. The failed sending tries, will be sent to this channel
// This channel must be consumed it EcoSendingErrors property is set to true
var SendingErrorsChan chan *Package

// Package sending retries chan
var packagesToBeRetiedChan chan *Package

// If is set to true, the sending errors will be sent to SendingErrorsChan channel
var EchoSendingErrors bool

// Retries per package
var retriesPerPackage map[int64]int8

//
// Packaging limits
//
// Maximum size of the messages package. If the package reaches this size, it will be sent.
var messagesPackageMaxSize int32

// Maximum wasting time (in ms) that the collector will wait for reaching the maximum size.
// Once this time is consumed, the package will be sent.
var messagesPackageMaxWaitingTime int64

// Number of retries for package sending to kafka if it fails
var messagesPackageMaxSendingRetries int8

// Size of the channel used to put the packages which will be sent
var messagesPackageSendingChanSize int16

// Interval of time (in ms) to check for package sending
var messagesPackageCheckTime int64

// Lot
var lot *lib_gc_fifo.Fifo

// Mutexes
var mutex *sync.Mutex
var mutexRetriesPerPackage *sync.Mutex

// Time out wrapper for package sending
type packageSendingChecker struct{}

func (psc *packageSendingChecker) Execute() (*[]interface{}, error) {
	if lot.Len() > 0 {
		checkPackageSending()
	}
	return nil, nil
}

// Time out wrapper for package closing
type packageCloser struct{}

func (pc *packageCloser) Execute() (*[]interface{}, error) {
	if lot.Len() > 0 {
		closePackage()
	}
	return nil, nil
}

// Timeout wrapper for package appending to retrying channel
type packageRetryingChannelAppender struct {
	pkg *Package
}

func (prca *packageRetryingChannelAppender) Execute() (*[]interface{}, error) {

	f := func(retriesPerPackage map[int64]int8, p []interface{}) interface{} {
		id := p[0].(int64)
		return retriesPerPackage[id]
	}
	pkg_retries := doOnRetriesPerpackage(f, prca.pkg.Id).(int8)

	msg, _ := EVENT.NotifyEvent("060-006", "", &[]string{fmt.Sprint(prca.pkg.Id), fmt.Sprint(prca.pkg.Queue.Len()), fmt.Sprint(pkg_retries), fmt.Sprint(messagesPackageMaxSendingRetries)})
	EVENT_PUBLISHER.EventPublisherChannel <- msg
	packagesToBeRetiedChan <- prca.pkg
	return nil, nil
}

// Checkig for package sending loop
var checkLoop lib_gc_runner.TASK_IMPLEMENTATION = func(eot chan struct{}, taskNanager lib_gc_runner.TaskManager_I, args []interface{}) ([]interface{}, error) {

	tickerCheck := time.NewTicker(time.Duration(messagesPackageCheckTime) * time.Millisecond)
	tickerSendingByTime := time.NewTicker(time.Duration(messagesPackageMaxWaitingTime) * time.Millisecond)

	for {
		select {

		case <-eot: // lib_gc_running logic
			fmt.Println("The task finalizes by EOT signal ...")
			return []interface{}{}, nil

		case <-tickerCheck.C:
			{
				if _, _, timeout := lib_gc_timeout_wrapper.TimeoutWrapper.Wrap(time.Duration(TIMEOUT_SENDING_CHECKING)*time.Millisecond, &packageSendingChecker{}); timeout != nil {
					msg, _ := EVENT.NotifyEvent("060-001", "", &[]string{fmt.Sprint(TIMEOUT_SENDING_CHECKING)})
					EVENT_PUBLISHER.EventPublisherChannel <- msg
				}
			}

		case <-tickerSendingByTime.C:
			{
				if _, _, timeout := lib_gc_timeout_wrapper.TimeoutWrapper.Wrap(time.Duration(TIMEOUT_PACKAGE_CLOSING)*time.Millisecond, &packageCloser{}); timeout != nil {
					msg, _ := EVENT.NotifyEvent("060-002", "", &[]string{fmt.Sprint(TIMEOUT_PACKAGE_CLOSING)})
					EVENT_PUBLISHER.EventPublisherChannel <- msg
				}
			}

		default:
			// None action by default
		}
	}
	fmt.Println("The task has finished by its self ....")
	return []interface{}{}, nil

}

// Lot sending loop
var sendPackagesLoop lib_gc_runner.TASK_IMPLEMENTATION = func(eot chan struct{}, taskNanager lib_gc_runner.TaskManager_I, args []interface{}) ([]interface{}, error) {
	for {
		select {

		case <-eot:
			{ // lib_gc_running logic
				fmt.Println("The task finalizes by EOT signal ...")
				return []interface{}{}, nil
			}

		case pkg := <-packagesToBeSentChan:
			{
				packageSender(pkg)
			}

		default:
			// None action by default
		}
	}

	fmt.Println("The task has finished by its self ....")
	return []interface{}{}, nil
}

// Lot retrying loop
var retryingPackagesLoop lib_gc_runner.TASK_IMPLEMENTATION = func(eot chan struct{}, taskNanager lib_gc_runner.TaskManager_I, args []interface{}) ([]interface{}, error) {
	for {
		select {

		case <-eot:
			{ // lib_gc_running logic
				fmt.Println("The task finalizes by EOT signal ...")
				return []interface{}{}, nil
			}

		case pkg := <-packagesToBeRetiedChan:
			{
				packageSender(pkg)
			}

		default:
			// None action by default
		}
	}

	fmt.Println("The task has finished by its self ....")
	return []interface{}{}, nil
}

// Main loop
var mainLoop lib_gc_runner.TASK_IMPLEMENTATION = func(eot chan struct{}, taskNanager lib_gc_runner.TaskManager_I, args []interface{}) ([]interface{}, error) {

	for {
		select {

		case <-eot: // lib_gc_runner logic
			fmt.Println("The task finalizes by EOT signal ...")
			return []interface{}{}, nil

		case message := <-MessagesCollectorChan:
			{
				processMessage(message)
			}

		default:
			// None action by default
		}
	}
	fmt.Println("The task has finished by its self ....")
	return []interface{}{}, nil
}

func packageSender(pkg *Package) {

	// Closure for increment the number of retries and take it
	fadd := func(retriesPerPackage map[int64]int8, p []interface{}) interface{} {
		id := p[0].(int64)
		retriesPerPackage[id] = retriesPerPackage[id] + 1

		// Note the package retrying
		PENDING_RETRYING_PAKCAGES.Add(1)

		return retriesPerPackage[id]
	}

	// Closure for removing the package from retries map
	frm := func(retriesPerPackage map[int64]int8, p []interface{}) interface{} {
		id := p[0].(int64)
		delete(retriesPerPackage, id)

		// Note the package retrying
		PENDING_RETRYING_PAKCAGES.Add(-1)

		return nil
	}

	if pkg.Queue.Len() > 0 {
		// Package sending logic
		if pkg, err := sendPackage(pkg); err != nil && len(err) > 0 {

			// Note the package sending error
			PACKAGES_SENDING_ERRORS.Add(1)

			// Increment the number of retries and take it
			retries := doOnRetriesPerpackage(fadd, pkg.Id).(int8)

			// Check if the nubmer of retries is greather than the max allowed value
			note_discarded := false
			if retries > messagesPackageMaxSendingRetries {

				msg, _ := EVENT.NotifyEvent("060-005", "", &[]string{fmt.Sprint(pkg.Id), fmt.Sprint(pkg.Queue.Len()), fmt.Sprint(messagesPackageMaxSendingRetries), fmt.Sprintf("%#v", err)})
				EVENT_PUBLISHER.EventPublisherChannel <- msg
				doOnRetriesPerpackage(frm, pkg.Id)
				note_discarded = true

			} else {

				// The package is re-sent to the sending packages channel for a new retry
				_packageRetryingChannelAppender := &packageRetryingChannelAppender{pkg: pkg}
				if _, _, timeout := lib_gc_timeout_wrapper.TimeoutWrapper.Wrap(time.Duration(TIMEOUT_PUTTING_MESSAGE)*time.Millisecond, _packageRetryingChannelAppender); timeout != nil {
					// It has result in timeout
					msg, _ := EVENT.NotifyEvent("060-008", "", &[]string{fmt.Sprint(pkg.Id)})
					EVENT_PUBLISHER.EventPublisherChannel <- msg
					doOnRetriesPerpackage(frm, pkg.Id)
					note_discarded = true
				} else {
					doOnRetriesPerpackage(fadd, pkg.Id)
				}
			}

			if note_discarded {
				// Note the discarded package
				DISCARDED_MESSAGES.Add(1)
			}

		} else {
			// Remove the sent package from sending retries registry
			doOnRetriesPerpackage(frm, pkg.Id)
		}
	} else {
		doOnRetriesPerpackage(frm, pkg.Id)
	}
}

// Package intializatino
func initializeLot() *lib_gc_fifo.Fifo {
	mutex.Lock()
	defer mutex.Unlock()
	oldLot := lot
	lot = lib_gc_fifo.GetFifo(messagesPackageMaxSize + 1)

	// Note the package length
	CURRENT_ON_BUILDING_PACKAGE_LENGTH.Set(0)

	return oldLot
}

// Message processing timeout wrapper
type processMessageTimeoutWrapper struct {
	lot     *lib_gc_fifo.Fifo
	message *Message
}

var putMessages map[int64]int64 = make(map[int64]int64)
func (pm *processMessageTimeoutWrapper) Execute() (*[]interface{}, error) {
	pm.lot.Put(pm.message)

	// Note the package length
	CURRENT_ON_BUILDING_PACKAGE_LENGTH.Add(1)

	return nil, nil
}

// This methos adds a new message to the
func processMessage(message *Message) {
	toWrapper := &processMessageTimeoutWrapper{lot: lot, message: message}
	if _, _, timeout := lib_gc_timeout_wrapper.TimeoutWrapper.Wrap(time.Duration(TIMEOUT_PUTTING_MESSAGE)*time.Millisecond, toWrapper); timeout != nil {
		msg, _ := EVENT.NotifyEvent("060-003", "", &[]string{fmt.Sprint(TIMEOUT_PUTTING_MESSAGE)})
		EVENT_PUBLISHER.EventPublisherChannel <- msg
	}

	// Check if the package is already full
	checkPackageSending()
}

// This checks if the messages package must be sent, by maximum size reached and/or by maximum waiting time reached
func checkPackageSending() {
	if lot.Len() >= messagesPackageMaxSize {
		closePackage()
	}
}

// Closing the package to be sent
func closePackage() {
	if lot.Len() > 0 {
		lot2Send := initializeLot()
		packageToSend := Package{Queue: lot2Send, Id: time.Now().UnixNano(), SendingErrors: make([]error, 0)}
		f := func(retriesPerPackage map[int64]int8, p []interface{}) interface{} {
			id := p[0].(int64)
			retriesPerPackage[id] = 0
			return nil
		}
		doOnRetriesPerpackage(f, packageToSend.Id)

		packagesToBeSentChan <- &packageToSend
	}
}

// Sending all of Messages of a package.
// If at end of the package's sending it is not empty (someone message sending has failed),
// it will be returned with an error. in this way, if the number of retries has not been exceeded,
// the package will be tried to be sent again with the remain messages.
func sendPackage(pkg *Package) (*Package, []error) {

	err := make([]error, 0)
	sender := pkg.Queue.Peek().(*Message).Sender
	pkg.Sender = sender
	failPkg := &Package{Id: pkg.Id, Queue: pkg.Queue, SendingErrors: pkg.SendingErrors, Sender: sender}

	if pkg, sending_err := sendMessage(pkg); sending_err != nil {

		msg, _ := EVENT.NotifyEvent("060-007", "", &[]string{fmt.Sprint(pkg.Id), fmt.Sprint(pkg.Queue.Len()), sending_err.Error()})
		EVENT_PUBLISHER.EventPublisherChannel <- msg

		failPkg.Queue = pkg.Queue
		failPkg.SendingErrors = append(pkg.SendingErrors, sending_err)
		err = append(err, sending_err)
		if EchoSendingErrors {
			SendingErrorsChan <- failPkg
		}
	}

	return failPkg, err
}

// Sending the message through the specialized sender
func sendMessage(pkg *Package) (*Package, error) {
	if sender_func, err := MessageSenderFactory.GetSender(pkg.Sender); err != nil {
		return pkg, err
	} else {
		return sender_func(pkg)
	}
}

// ---------------------------------------- helpers
func BuildMessageHelper(content interface{}, id int64, specializedSender string, Ttl time.Duration) *Message {
	if id == 0 {
		id = time.Now().UnixNano()
	}
	return &Message{Content: content, Id: id, CreationTime: time.Now().UnixNano(), Ttl: Ttl, SendingRetries: 0, Sender: specializedSender}
}

func doOnRetriesPerpackage(f func(map[int64]int8, []interface{}) interface{}, p ...interface{}) interface{} {
	mutexRetriesPerPackage.Lock()
	defer mutexRetriesPerPackage.Unlock()
	return f(retriesPerPackage, p)
}
