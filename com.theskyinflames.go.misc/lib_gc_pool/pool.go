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

import ()

func init() {
	POOLMaker = &poolFactory{}
}

/*
	Pool factory --------------------------------------------------------
*/
var POOLMaker PoolMaker

type PoolMaker interface {
	GetPool(size int32) Pooler
}

type poolFactory struct{}

func (pf *poolFactory) GetPool(size int32) Pooler {
	var p Pooler = &pool{size, make(chan (*Poolable), size)}
	return p
}

/*
   Pool ----------------------------------------------------------------
*/

type Pooler interface {
	GetSize() int32
	GetPool() chan (*Poolable)
}

type Poolable struct {
	Item interface{}
}

type pool struct {
	size   int32
	_pool_ chan (*Poolable)
}

func (p *pool) GetSize() int32 {
	return p.size
}

func (p *pool) GetPool() chan (*Poolable) {
	return p._pool_
}

func (p *pool) Stream(f func(interface{}) interface{}) {

	end := false
	n_chan := make(chan (*Poolable), p.size)
	for !end {
		select {
		case poolable := <-p._pool_:
			poolable.Item = f(poolable.Item)
			n_chan <- poolable
		default:
			end = true
			break
		}
	}
	p._pool_ = n_chan
}
