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

package lib_gc_container

import (
    "testing"
    "fmt"
)

func init(){
    a := []string{}
    if container,err := GenericContainerFactory.GetContainer(a);err!=nil{
        panic(err)
    }else{
        testContainer = &TestContainer{container}
    }

    testContents = &TestContents{name:"testContent"}
    testContents.Contents = Contents{testContainer.cnt,CONTAINER_STATUS_STOPPED, testContents}
    testContents.Contents.WaitForStatusChanges()

    testContents2 = &TestContents{name:"testContent2"}
    testContents2.Contents = Contents{testContainer.cnt,CONTAINER_STATUS_STOPPED, testContents2}
    testContents2.Contents.WaitForStatusChanges()

}

var testContainer *TestContainer

var testContents *TestContents
var testContents2 *TestContents

const ITEM = "ITEM"

type TestContainer struct{
    cnt IGenericContainer
}

type TestContents struct{
    Contents
    name string
}

func (c *TestContents) Start() error{
    fmt.Println("The testContest ",c.name ," is started.")
    return  nil
}

func (c *TestContents) Stop() error{
    fmt.Println("The testContest ",c.name ," is stopped.")
    return nil
}
func (c *TestContents) Shutdown() error{
    fmt.Println("The testContest ",c.name ," is shutdown.")
    return nil
}

func Test_Container_Item_Add(t *testing.T){
    if err:=testContainer.cnt.AddItem(ITEM, testContents);err!=nil{
        t.Error(err)
    }
}

func Test_Container_Take_Item(t *testing.T){
    if item,err := testContainer.cnt.GetItem(ITEM);err!=nil{
        t.Error(err)
    }else{
        fmt.Println("Retrieved item with name: ",item.(*TestContents).name)
    }
}

func Test_Container_Start(t *testing.T) {
    if err:=testContainer.cnt.Start();err!=nil{
        t.Error(err)
    }
}

func Test_Container_Stop(t *testing.T) {
    if err:=testContainer.cnt.Stop();err!=nil{
        t.Error(err)
    }
}

func Test_Container_Start2(t *testing.T) {
    if err:=testContainer.cnt.Start();err!=nil{
        t.Error(err)
    }
}

func Test_Container_Shutdown(t *testing.T) {
    if err:=testContainer.cnt.Shutdown();err!=nil{
        t.Error(err)
    }
}
