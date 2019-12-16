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

package lib_gc_pool

import (
	"fmt"
	"testing"
)

func init() {

}

var tpool Pooler

const pool_size int32 = 5

func Test_pool_creation(t *testing.T) {
	tpool = POOLMaker.GetPool(pool_size)
	fmt.Println("Pool created.")
}

func Test_pool_filling(t *testing.T) {
	for z := 0; z < int(pool_size); z++ {
		fmt.Println("Pooling: ", z)
		tpool.GetPool() <- &Poolable{Item: z}
	}
	fmt.Println("Pool is full")
}

func Test_pool_item_extraction(t *testing.T) {
	z := <-tpool.GetPool()
	fmt.Println("Z: ", z.Item.(int))
	z1 := <-tpool.GetPool()
	fmt.Println("Z1: ", z1.Item.(int))

	tpool.GetPool() <- z
	tpool.GetPool() <- z1
}
