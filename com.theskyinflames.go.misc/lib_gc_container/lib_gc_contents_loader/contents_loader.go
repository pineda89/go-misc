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

package lib_gc_contents_loader

import (
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	"sync"
)

func init() {
	initiators = make(map[string]func() error)
}

var initiators map[string]func() error
var mutex *sync.Mutex = &sync.Mutex{}

func AddInitiator(name string, initiator func() error) {
	mutex.Lock()
	initiators[name] = initiator
	lib_gc_log.Info.Printf("Adding %d initiators to the container \n", len(initiators))
	mutex.Unlock()
}

/*
   This method initialize all of container contents. Basically, these are added to the containers.
*/
func InitializeContainerContents() error {

	for name, initiator := range initiators {
		lib_gc_log.Info.Printf("Firing up initiator %s \n", name)
		if err := initiator(); err != nil {
			lib_gc_log.Error.Printf("ERROR at execute the containter contents initiator %s: %v\n", name, err)
		}
	}
	lib_gc_log.Info.Println("Container contents initated OK.")
	return nil
}

func GetIniatorsLen() int {
	return len(initiators)
}
