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

package lib_gc_timeout_wrapper

import(

    "time"
    "errors"
)

func init(){
    TimeoutWrapper = &timeoutWrapper{}
}

type TIMEOUT_ERROR error                    // Timeout error constant
var TimeoutWrapper TimeoutWrapper_I = nil   // Timeout wrapper

type Wrappable_I interface{
    Execute() (*[]interface{},error)
}

type TimeoutWrapper_I interface{
    Wrap(timeout time.Duration, wrappable Wrappable_I) (*[]interface{}, error, TIMEOUT_ERROR)
}

type f_result struct{
    result *[]interface{}
    err error
}

type timeoutWrapper struct{}

// Timeout wrapper implementation
func (tw *timeoutWrapper) Wrap(timeout time.Duration, wrappable Wrappable_I) (*[]interface{}, error, TIMEOUT_ERROR) {

    var c_func chan *f_result = make(chan *f_result)  // Wrappable function signal channel
    var c_timeout chan int = make(chan int)           // Timeout signal channel
    var err_timeout string                            // Timetout error message
    var ret *f_result                                 // Timeout wrapped action result

    go func() {   // Fire up the wrrapped function
        result,err := wrappable.Execute()
        c_func <- &f_result{result,err}
    }()

    go func(){  // Fire up the timeout watcher
        timer := time.NewTimer(timeout)
        <-timer.C
        close(c_timeout)
    }()

    cancel := false
    for !cancel {           // Wait for function execution finish, or for timeout signal
        select {
        case <-c_timeout:
            err_timeout = "*Timeout has been reached*"
            cancel = true
        case ret = <-c_func:
            cancel = true
        }
    }

    if err_timeout != "" { // The function execution has result in a timeout
        return nil, nil, errors.New(err_timeout).(TIMEOUT_ERROR)
    }else { // The function execution has finised in time
        return ret.result, ret.err, nil
    }
}
