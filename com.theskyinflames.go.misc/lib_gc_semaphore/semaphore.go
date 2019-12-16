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

package lib_gc_semaphore

import (
	"sync"
    "errors"
)

func init() {
	SemaphoreFactory = &semaphoreFactory{}
}

var SemaphoreFactory SemaphoreFacory_I

var mutex1 *sync.Mutex = &sync.Mutex{}
var mutex2 *sync.Mutex = &sync.Mutex{}

type empty struct{}
type semaphore_chan chan empty

type SemaphoreFacory_I interface {
	GetSemaphore(capacity int) Semaphore_I
}

type semaphoreFactory struct{}

func (sf *semaphoreFactory) GetSemaphore(capacity int) Semaphore_I {
    _s := semaphore{make(semaphore_chan, capacity)}
	return &_s
}

type Semaphore_I interface {
	Lock(n int)
    TryToUnLock(n int) error
    ForceToUnLock(n int) error
}

type semaphore struct {
	capacity semaphore_chan
}

func (s *semaphore) Lock(n int) {
	e := empty{}
	for i := 0; i < n; i++ {
		s.capacity <- e
	}
}

func (s *semaphore) TryToUnLock(n int) error{
	for i := 0; i < n; i++ {
        select {
        case <-s.capacity:
        default:
            return errors.New("There are not enought locks to take.")
        }
	}
    return nil
}

func (s *semaphore) ForceToUnLock(n int) error{
    for i := 0; i < n; i++ {
        <-s.capacity
    }
    return nil
}
