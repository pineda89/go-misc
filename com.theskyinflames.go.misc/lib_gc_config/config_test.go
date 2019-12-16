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

package lib_gc_config

import (
	"fmt"
	"strconv"
	"testing"
)

func Test_loadConfigSTRING(t *testing.T) {
	var value string
	value = ConfigurationProvider.GetPropertySTRING("PULSE_STORAGE", "PULSE_STORAGE_ACTIVE_STORAGES")
	fmt.Println("STRING:", value)
}

func Test_loadConfigBOOL(t *testing.T) {
	var value bool
	value = ConfigurationProvider.GetPropertyBOOL("DEBUG", "DEBUG")
	fmt.Println("BOOL:", strconv.FormatBool(value))
}

func Test_loadConfigINT(t *testing.T) {
	var value int
	value = ConfigurationProvider.GetPropertyINT("PULSE_AGGREGATOR_DEFAULT", "PULSE_AGGREGATOR_DEFAULT_BLOCK_SIZE")
	fmt.Println("INT:", value)
}

func Test_loadConfigINT64(t *testing.T) {
	var value int64
	value = ConfigurationProvider.GetPropertyINT64("PULSE_AGGREGATOR_DEFAULT", "PULSE_AGGREGATOR_DEFAULT_BLOCK_SIZE")
	fmt.Println("INT64:", value)
}

//func Test_loadConfigFLOAT64(t *testing.T) {
//	var value float64
//	value = PtrConfigProvider.GetPropertyFLOAT64("TEST", "MAX_CONCURRENT_PENTING_LOCK_REQUESTS")
//	fmt.Println("FLOAT64:", value)
//}
