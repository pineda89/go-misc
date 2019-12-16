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

package lib_gc_container

import (
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	"errors"
	"sync"
)

func init() {
	GenericContainerFactory = &genericContainerFactory{}
}

type ShutdownChan chan struct{}
type StartChan chan struct{}
type StopChan chan struct{}

type ContainerStatus int

const (
	CONTAINER_STATUS_STARTED  ContainerStatus = 1
	CONTAINER_STATUS_STOPPED  ContainerStatus = 2
	CONTAINER_STATUS_SHUTDOWN ContainerStatus = 3
)

type IContainerStatusListener interface {
	Start() error
	Stop() error
	Shutdown() error
}

type IGenericContainerGetter interface {
	GetGenericContainer() IGenericContainer
}

var GenericContainerFactory IGenericContainerFactory

type IGenericContainerFactory interface {
	GetContainer(activated_items []string, name string) (IGenericContainer, error)
}

type genericContainerFactory struct{}

func (factory *genericContainerFactory) GetContainer(activated_items []string, name string) (IGenericContainer, error) {
	container := &genericContainer{name: name, items: make(map[string]IGenericContents), parameters: make(map[string]interface{}), activated_items: activated_items, container_status: CONTAINER_STATUS_STOPPED, shutdown_chan: make(ShutdownChan, 2), start_chan: make(StartChan, 2), stop_chan: make(StopChan, 2), wg: &sync.WaitGroup{}, initialized: false}
	lib_gc_log.Info.Printf("Take container %p with name %s \n", container, name)
	return container, nil
}

type IGenericContents interface {
	WaitForStatusChanges()
}

type IGenericContainer interface {
	GetName() string
	GetWg() *sync.WaitGroup
	AddItem(string, IGenericContents) error
	GetItem(string) (IGenericContents, error)
	IsItemActivated(string) bool
	GetContainerStatusChannels() (ShutdownChan, StartChan, StopChan)
	AddParameter(string, interface{})
	GetParameter(string) (interface{}, error)
	IContainerStatusListener
}

type genericContainer struct {
	name             string
	items            map[string]IGenericContents
	parameters       map[string]interface{}
	activated_items  []string
	container_status ContainerStatus
	shutdown_chan    ShutdownChan
	start_chan       StartChan
	stop_chan        StopChan
	wg               *sync.WaitGroup
	initialized      bool
}

func (container *genericContainer) GetName() string {
	return container.name
}

func (container *genericContainer) GetWg() *sync.WaitGroup {
	return container.wg
}

func (container *genericContainer) AddItem(name string, item IGenericContents) error {
	container.items[name] = item
	return nil
}

func (container *genericContainer) GetItem(name string) (IGenericContents, error) {
	if item, ok := container.items[name]; !ok {
		_msg, _ := lib_gc_event.NotifyEvent("005-001", "", &[]string{name})
		return nil, errors.New(_msg)
	} else {
		return item, nil
	}
}

func (container *genericContainer) IsItemActivated(name string) bool {
	active := false
	for _, n := range container.activated_items {
		if n == name {
			active = true
			break
		}
	}
	return active
}

func (container *genericContainer) GetContainerStatusChannels() (ShutdownChan, StartChan, StopChan) {
	return container.shutdown_chan, container.start_chan, container.stop_chan
}

func (container *genericContainer) AddParameter(key string, value interface{}) {
	container.parameters[key] = value
}

func (container *genericContainer) GetParameter(key string) (interface{}, error) {
	if value, ok := container.parameters[key]; !ok {
		_msg, _ := lib_gc_event.NotifyEvent("005-002", "", &[]string{key})
		return nil, errors.New(_msg)
	} else {
		return value, nil
	}
}

func (container *genericContainer) Shutdown() error {
	close(container.shutdown_chan)
	container.container_status = CONTAINER_STATUS_SHUTDOWN
	return nil
}

func (container *genericContainer) Start() error {

	if !container.initialized {
		// Wait for all contents has been added to the container.
		container.wg.Wait()

		// Start the status listener of each content of each container
		for _, c := range container.items {
			c.WaitForStatusChanges()
		}

		// Container initialized
		container.initialized = true
		lib_gc_log.Info.Printf("Container %p [%s] initiated OK \n", container, container.name)
	}

	// Start the container
	container.container_status = CONTAINER_STATUS_STARTED
	close(container.start_chan)
	container.start_chan = make(StartChan)
	return nil
}

func (container *genericContainer) Stop() error {
	container.container_status = CONTAINER_STATUS_STOPPED
	close(container.stop_chan)
	container.stop_chan = make(StopChan)
	return nil
}
