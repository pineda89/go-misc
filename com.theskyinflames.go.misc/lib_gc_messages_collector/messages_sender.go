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
	"errors"
	"fmt"
	"sync"
)

func init() {
	// Initializing the message's senders factory
	MessageSenderFactory = &messageSenderFactory{messageSenders: make(map[string]MessageSenderFunc), Mutex: sync.Mutex{}}
}

var MessageSenderFactory MessageSenderFactory_I

type MessageSenderFunc func(message *Package) (*Package, error)

// List of available senders. Each one of them, specialized in an specific sending target.
type MessageSenderFactory_I interface {
	AddMessageSender(string, MessageSenderFunc)
	RemoveMessageSender(string) (MessageSenderFunc, error)
	GetSender(string) (MessageSenderFunc, error)
}

// ------------------------------------------------
//	MessageFactory_I implementation
// ------------------------------------------------
type messageSenderFactory struct {
	sync.Mutex
	messageSenders map[string]MessageSenderFunc
}

func (msf *messageSenderFactory) AddMessageSender(name string, sender MessageSenderFunc) {
	msf.Lock()
	defer msf.Unlock()
	msf.messageSenders[name] = sender
}

func (msf *messageSenderFactory) RemoveMessageSender(name string) (MessageSenderFunc, error) {
	msf.Lock()
	defer msf.Unlock()

	var sender MessageSenderFunc
	var err error

	if sender, err = msf.GetSender(name); err == nil {
		delete(msf.messageSenders, name)
	}
	return sender, err
}

func (msf *messageSenderFactory) GetSender(name string) (MessageSenderFunc, error) {
	msf.Lock()
	defer msf.Unlock()

	if sender, ok := msf.messageSenders[name]; !ok {
		return nil, errors.New(fmt.Sprintf("The message's sender [%s] does not exist !!", name))
	} else {
		return sender, nil
	}
}
