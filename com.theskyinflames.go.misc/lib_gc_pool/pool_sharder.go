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
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	"errors"
)

func init() {
	PoolSharderFactory = &poolSharderFactoryImpl{pool_sharders: make(map[string]PoolSharder_I)}
}

/*
   Pool sharder factory
*/
type PoolSharderFactory_I interface {
	AddPoolShardingAdapter(name string, pool_sharder PoolSharder_I) error
	GetPoolSharder(name string) (PoolSharder_I, error)
}
type poolSharderFactoryImpl struct {
	pool_sharders map[string]PoolSharder_I
}

func (psf *poolSharderFactoryImpl) AddPoolShardingAdapter(name string, pool_sharder PoolSharder_I) error {
	psf.pool_sharders[name] = pool_sharder
	return nil
}
func (psf *poolSharderFactoryImpl) GetPoolSharder(name string) (PoolSharder_I, error) {
	if sharder, ok := psf.pool_sharders[name]; !ok {
		_msg, _ := EVENT.NotifyEvent("037-002", "", &[]string{name})
		return nil, errors.New(_msg)
	} else {
		return sharder.GetInstance()
	}
}

var PoolSharderFactory PoolSharderFactory_I

/*
   Pool sharder implementation
*/
type ShardedPool struct {
	Name   string
	PKPool Pooler
	Weight int
}
type PoolPerKeys struct {
	Name   string
	PKPool Pooler
	Keys   *[][]byte
}
type PoolSharder_I interface {
	InitializeWithWeightedPools(pools []*ShardedPool) error
	GetPoolPerKey(key *[]byte) (*ShardedPool, error)
	GetPoolsPerKeys(keys *[][]byte) (map[string]*PoolPerKeys, error)
	GetInstance() (PoolSharder_I, error)
}
