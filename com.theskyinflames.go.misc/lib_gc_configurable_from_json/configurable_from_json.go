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

package lib_gc_configurable_from_json

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func init() {
}

/*
/ It makes an component configurable
*/
type Configurable struct{}

type MakeConfigurable interface {
	LoadConfigurationFromFile(filePath string, record interface{}) (interface{}, error)
	GetRecord() interface{}
}

// Do Resolver MakeConfigurable interface implementation
func (configurable *Configurable) LoadConfigurationFromFile(filePath string, record interface{}) (interface{}, error) {

	if b, err := ioutil.ReadFile(filePath); err != nil {
		return nil, err
	} else if len(b) > 0 {
		err = json.Unmarshal(b, &record)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("Empty file: " + filePath)
	}
	return record, nil
}

func (configurable *Configurable) GetRecord() interface{} {
	return errors.New("Method GetRecord not implemented !!!")
}
