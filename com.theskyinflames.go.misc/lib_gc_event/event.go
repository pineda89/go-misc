/*
Copyright 2015 - Serhstourism, S.A

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

package lib_gc_event

import (
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_configurable_from_json"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"

	"bytes"
	"encoding/json"
	"errors"
	"os"
	"strings"
	"time"
)

func init() {
	ptrEH = nil
	conf_file := os.Getenv("GO_COMMON_LIB_EVENTS_FILE") // Config folder
	if conf_file == "" {
		panic(errors.New("GO_COMMON_LIB_EVENTS_FILE environment variable is not defined !!!"))
	} else {
		if err := loadConfigurationFromFile(conf_file); err != nil {
			panic(err)
		}
	}
	lib_gc_log.Info.Println("lib_gc_event package initiated")
}

type EventLevel byte

const EVENTS_FILE_NAME = "EventsConfiguration.json"

const (
	INFO EventLevel = iota
	WARNING
	CRITICAL
	FATAL
)

type Event struct {
	Code      string
	Level     string
	Message   string
	Timestamp int64
	Hostname  string
}

type eventsConf struct {
	Events   []Event
	EventMap map[string]*Event
}

type eventsHandler struct {
	Configuration eventsConf

	// Makes it configurable
	lib_gc_configurable_from_json.Configurable
}

// Event message to publish on a kafka topic
type eventMessage struct {
	Timestamp int64
	Event
}

type NotifyRequirement struct {
	EventCode  string
	Message    string
	Parameters *[]string
}

var ptrEH *eventsHandler

func loadConfigurationFromFile(configuration_path string) error {
	var err error
	ptrEH = new(eventsHandler)
	if _, err = ptrEH.LoadConfigurationFromFile(configuration_path, ptrEH.GetRecord()); err == nil {
		ptrEH.Configuration.EventMap = make(map[string]*Event, len(ptrEH.Configuration.Events))
		var ptrEvent *Event
		for i, _ := range ptrEH.Configuration.Events {
			ptrEvent = &ptrEH.Configuration.Events[i]
			ptrEH.Configuration.EventMap[ptrEvent.Code] = ptrEvent
		}
		lib_gc_log.Info.Printf("Loaded %d event codes.", len(ptrEH.Configuration.Events))
	} else {
		lib_gc_log.Error.Printf("Something went wrong on loading event codes : %s", err.Error())
	}
	return err
}

func NotifyEvent(eventCode string, pmessage string, parameters *[]string) (string, error) {
	//TO DO : Use go templates

	// Retrieve the event message
	var err error
	var message string
	if _event, ok := ptrEH.Configuration.EventMap[eventCode]; ok {

		event := &Event{Code: _event.Code, Level: _event.Level, Message: _event.Message}

		// Set the event message
		if len(pmessage) > 0 {
			event.Message = pmessage
		}

		// Insert customized message parts
		if parameters != nil {
			for _, param := range *parameters {
				event.Message = strings.Replace(event.Message, "#", param, 1)
			}
		}

		// Set timestamp
		event.Timestamp = time.Now().Unix()

		// Set hostname
		if event.Hostname, err = os.Hostname(); err != nil {
			event.Hostname = "NOT_RETRIEVED"
		}

		var _bytes bytes.Buffer
		if _b, err := json.Marshal(&event); err != nil {
			_bytes.WriteString("Some when wrong on marshal the message: " + err.Error())
		} else {
			_bytes.Write(_b)
		}
		message = _bytes.String()

	} else {
		error_msg := "It has not been possible to retrieve the event code " + eventCode
		lib_gc_log.Error.Println(error_msg)
		err = errors.New(error_msg)
	}

	return message, err
}

// -------------------------------------------------------------------------------

// MakeConfigurable interface implementation -------------------------------------
func (events *eventsHandler) GetRecord() interface{} {
	var record eventsConf
	return &record
}

func (events *eventsHandler) LoadConfigurationFromFile(filePath string, record interface{}) (interface{}, error) {
	if conf, err := events.Configurable.LoadConfigurationFromFile(filePath, record); err == nil {
		if res, ok := conf.(*eventsConf); ok {
			events.Configuration = *res
		} else {
			return nil, errors.New("Error on retrieve configuration...")
		}
	} else {
		return nil, err
	}
	return nil, nil
}

// --------------------------------------------------------------------------------
