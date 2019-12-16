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

    "testing"
    "time"
    "fmt"
    //"errors"
)
func init(){
}


func sum(a,b int) int{
    return a+b
}

type wrappable_sum struct{
    a int
    b int
}

func (ws *wrappable_sum) Execute() (*[]interface{}, error){
    return &[]interface{}{sum(ws.a,ws.b)},nil
}

type wrappable_sum2 struct{
    a int
    b int
}

func (ws *wrappable_sum2) Execute() (*[]interface{}, error){
    time.Sleep(2 * time.Second)
    return &[]interface{}{sum(ws.a,ws.b)},nil
}

type wrappable3 struct {
    a,b int
}
func (ws *wrappable3) Execute() (*[]interface{}, error){
    slice := make([][]int,2)
    slice[0] = []int{ws.a}
    slice[1] = []int{ws.b}

    if ws.b == 3{
        time.Sleep(5 * time.Second)
    }

    return &[]interface{}{slice},nil
}




func Test_WithoutTimeout(t *testing.T) {

    ws := &wrappable_sum{a:1,b:2}
    if result,err, timeout := TimeoutWrapper.Wrap(time.Duration(1000)*time.Millisecond, ws); timeout != nil {
        fmt.Printf("Timeout: %v\n")
    }else{
        fmt.Printf("OK Result: %v, err: %v\n", result, err)
    }
}
func Test_WithTimeout(t *testing.T) {

    ws := &wrappable_sum2{a:1,b:2}
    if result,err,timeout := TimeoutWrapper.Wrap(time.Duration(1000)*time.Millisecond, ws); timeout != nil {
        fmt.Printf("Es timeout correcto: %v\n", timeout)
    }else{
        t.Error("Debía haber sido timeout, result: %v, err: %v\n", result,err)
    }
}



func Test_WithoutTimeoutWithSliceOnReturn(t *testing.T) {
    ws := &wrappable3{a:1,b:2}
    if result,err,timeout := TimeoutWrapper.Wrap(time.Duration(1000)*time.Millisecond, ws); timeout != nil {
        fmt.Printf("Timeout: %v\n")
    }else{
        fmt.Printf("OK Result: %v, err: %v\n", result, err)
    }
}

func Test_WithTimeoutWithSliceOnReturn(t *testing.T) {
    ws := &wrappable3{a:1,b:3}
    if result,err,timeout := TimeoutWrapper.Wrap(time.Duration(1000)*time.Millisecond, ws); timeout != nil {
        fmt.Printf("Es timeout correcto: %v\n", timeout)
    }else{
        fmt.Printf("OK Result: %v, err: %v\n", result, err)
    }
}