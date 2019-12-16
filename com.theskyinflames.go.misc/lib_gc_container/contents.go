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
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_pool"
	"sync"
)

type Contents struct {
	Container      IGenericContainer
	Status         ContainerStatus
	StatusListener IContainerStatusListener
}

func (contents *Contents) Shutdown() error {
	contents.Status = CONTAINER_STATUS_SHUTDOWN
	return contents.StatusListener.Shutdown()
}

func (contents *Contents) Start() error {
	contents.Status = CONTAINER_STATUS_STARTED
	return contents.StatusListener.Start()
}

func (contents *Contents) Stop() error {
	contents.Status = CONTAINER_STATUS_STOPPED
	return contents.StatusListener.Stop()
}

func (contents *Contents) WaitForStatusChanges() {
	lib_gc_log.Info.Printf("Start WaitForStatusChanges on content %p, container %p [%s] \n", contents, contents.Container, contents.Container.GetName())

	// Synchronize the container Start with the goroutine
	wg := &sync.WaitGroup{}
	wg.Add(1)
	var once sync.Once

	go func(contents *Contents, wg *sync.WaitGroup) {
		var stop bool = false
		//        var first_loop bool = true
		for !stop {
			shutdown_chan, start_chan, stop_chan := contents.Container.GetContainerStatusChannels()
			lib_gc_log.Trace.Printf("Check for status, content %p [status %d], conainer %p [%s], shutdown chan %p, start chan %p, stop chan %p", contents, contents.Status, contents.Container, contents.Container.GetName(), shutdown_chan, start_chan, stop_chan)

			// Only on the first loop, the goroutine must be synchronized with the container. If not,
			// the container would close its start channel before the goroutine was running and the
			// start event would not be catched.
			once.Do(wg.Done)

			// Catch container events.
			select {
			case <-shutdown_chan:
				lib_gc_log.Trace.Printf("Received SHUTDOWN on content %p, container %p [%s] \n", contents, contents.Container, contents.Container.GetName())
				if err := contents.Shutdown(); err != nil {
					_msg, _ := lib_gc_event.NotifyEvent("003-002", "", &[]string{err.Error()})
					lib_gc_log.Error.Println(_msg)
				}
				stop = true
			case <-start_chan:
				lib_gc_log.Trace.Printf("Received START on content %p, container %p [%s] \n", contents, contents.Container, contents.Container.GetName())
				contents.Start()
			case <-stop_chan:
				lib_gc_log.Trace.Printf("Received STOP on content %p, container %p [%s] \n", contents, contents.Container, contents.Container.GetName())
				contents.Stop()
			}
		}
	}(contents, wg)
	wg.Wait()
}

type PooledContents struct {
	lib_gc_pool.Pooler
}

func (pc *PooledContents) WaitForStatusChanges() {
	n_pool := lib_gc_pool.POOLMaker.GetPool(pc.Pooler.GetSize())
	end := false

	for !end {
		select {
		case poolable := <-pc.Pooler.GetPool():
			(*poolable).Item.(IGenericContents).WaitForStatusChanges()
			n_pool.GetPool() <- poolable
		default:
			end = true
			break
		}
	}
	pc.Pooler = n_pool
}
