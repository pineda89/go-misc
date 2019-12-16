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

package lib_gc_pool_sharder_adapter

import (
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	POOL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_pool"

	"github.com/serialx/hashring"

	"errors"
)

func init() {
	// Register the sharding adapter
	POOL.PoolSharderFactory.AddPoolShardingAdapter(HASHRING_ADAPTER_NAME, &poolSharderImpl{pools: make(map[string]*POOL.ShardedPool)})
}

const HASHRING_ADAPTER_NAME string = "HASHRING"

type poolSharderImpl struct {
	hashRing *hashring.HashRing
	pools    map[string]*POOL.ShardedPool
}

func (ps *poolSharderImpl) GetInstance() (POOL.PoolSharder_I, error) {
	return ps, nil
}

// Initialize the pool sharder
func (ps *poolSharderImpl) InitializeWithWeightedPools(pools []*POOL.ShardedPool) error {
	_m := make(map[string]int)
	for _, sp := range pools {
		_m[sp.Name] = sp.Weight
		ps.pools[sp.Name] = sp
	}
	ps.hashRing = hashring.NewWithWeights(_m)
	return nil
}

// Retrieve the pool for a given key
func (ps *poolSharderImpl) GetPoolPerKey(key *[]byte) (*POOL.ShardedPool, error) {
	if pooler_name, ok := ps.hashRing.GetNode(string(*key)); !ok {
		_msg, _ := EVENT.NotifyEvent("037-001", "", &[]string{})
		return nil, errors.New(_msg)
	} else {
		//println("*jas* key: ",string(*key),", pool: ",ps.pools[pooler_name].Name)
		return ps.pools[pooler_name], nil
	}
}

// Given a set of keys, it returns a map of instances of PoolPerKeys, where the keys appears
// grouped by its corresponding pool.
func (ps *poolSharderImpl) GetPoolsPerKeys(keys *[][]byte) (map[string]*POOL.PoolPerKeys, error) {
	_poolsPerKeys := make(map[string]*POOL.PoolPerKeys)
	var sharded_pool *POOL.ShardedPool
	var err error
	for _, key := range *keys {
		if sharded_pool, err = ps.GetPoolPerKey(&key); err != nil {
			return nil, err
		} else {
			if _, ok := _poolsPerKeys[sharded_pool.Name]; !ok {
				ppk := POOL.PoolPerKeys{}
				ppk.Name = sharded_pool.Name
				ppk.PKPool = sharded_pool.PKPool
				ppk.Keys = &[][]byte{key}
				_poolsPerKeys[sharded_pool.Name] = &ppk
			} else {
				*_poolsPerKeys[sharded_pool.Name].Keys = append((*_poolsPerKeys[sharded_pool.Name].Keys), key)
			}
		}
	}
	return _poolsPerKeys, nil
}
