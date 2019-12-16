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
	"errors"
	"expvar"
	"fmt"

	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_conf"

	"github.com/Unknwon/goconfig"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
)

func init() {
	PtrConfigProvider = &configuration{}
	if err := PtrConfigProvider.loadConfiguration(fmt.Sprintf("%s/%s", lib_gc_conf.CONF_PREFIX, lib_gc_conf.CONF_FILE_NAME)); err != nil {
		panic(err)
	}
}

var PtrConfigProvider IConfigurationProvider

const COFNIG_MAP_NAME = "CONFIG_MAP"

// Config map for exposing the configuratino throught expvar metrics
var configMap *expvar.Map

type configuration struct {
	file *goconfig.ConfigFile
}

type IConfigurationProvider interface {
	loadConfiguration(string) error
	GetPropertySTRING(section string, property string) string
	GetPropertyBOOL(section string, property string) bool
	GetPropertyINT(section string, property string) int
	GetPropertyINT64(section string, property string) int64
	GetPropertyFLOAT64(section string, property string) float64
}

func (configuration *configuration) loadConfiguration(CONFIGURATION_FILE string) error {
	file, err := goconfig.LoadConfigFile(CONFIGURATION_FILE)
	if err != nil {
		return errors.New("Error cargando el fichero de configuracion global: " + err.Error())
	} else {
		configuration.file = file
		configMap = expvar.NewMap(COFNIG_MAP_NAME)
		lib_gc_log.Info.Println(fmt.Sprintf("lig_bc_config package initialized from file %s \n", CONFIGURATION_FILE))
		return nil
	}
}

func (configuration *configuration) GetPropertySTRING(section string, property string) string {
	if value, err := configuration.file.GetValue(section, property); err != nil {
		panic("An error has happened on retrieve the property (sec/prop) " + section + "/" + property + " from properties file: " + err.Error())
	} else {
		addPropertyToConfig(section, property, value)
		return value
	}
}

func (configuration *configuration) GetPropertyBOOL(section string, property string) bool {
	if value, err := configuration.file.Bool(section, property); err != nil {
		panic("An error has happened on retrieve the property (sec/prop) " + section + "/" + property + " from properties file: " + err.Error())
	} else {
		addPropertyToConfig(section, property, value)
		return value
	}
}

func (configuration *configuration) GetPropertyINT(section string, property string) int {
	if value, err := configuration.file.Int(section, property); err != nil {
		panic("An error has happened on retrieve the property (sec/prop) " + section + "/" + property + " from properties file: " + err.Error())
	} else {
		addPropertyToConfig(section, property, value)
		return value
	}
}

func (configuration *configuration) GetPropertyINT64(section string, property string) int64 {
	if value, err := configuration.file.Int64(section, property); err != nil {
		panic("An error has happened on retrieve the property (sec/prop) " + section + "/" + property + " from properties file: " + err.Error())
	} else {
		addPropertyToConfig(section, property, value)
		return value
	}
}

func (configuration *configuration) GetPropertyFLOAT64(section string, property string) float64 {
	if value, err := configuration.file.Float64(section, property); err != nil {
		panic("An error has happened on retrieve the property (sec/prop) " + section + "/" + property + " from properties file: " + err.Error())
	} else {
		addPropertyToConfig(section, property, value)
		return value
	}
}

func addPropertyToConfig(section string, property string, value interface{}) {
	key := fmt.Sprintf("%s>>%s", section, property)
	//fmt.Printf("*jas* expvar.Get(%s): %s", key, expvar.Get(key))
	if v := expvar.Get(key); v == nil {
		z := expvar.NewString(key)
		z.Set(fmt.Sprint(value))
		configMap.Set(fmt.Sprintf("%s>>%s", section, property), z)
	}
}
