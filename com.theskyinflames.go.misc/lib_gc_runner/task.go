/*
Copyright 2016 - Jaume Arús

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

package lib_gc_runner

import (
	"fmt"
	"runtime/debug"
	"time"

	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event_publisher"
)

// Task implementation type.
// The first parameter is a chan to catch the END signal. When this channel closes, the implementation task must to finish
// The second parameters is a slice of input parameters for the task implementation function
type TASK_IMPLEMENTATION func(chan struct{}, TaskManager_I, []interface{}) ([]interface{}, error)

//
// Task
//
type Task_I interface {
	GetID() int64
	GetDuration() time.Duration
	GetResponseChan() chan *TaskResponse
	Run(args []interface{}) error
	IsRunning() bool
	IsFinished() bool
	Finalize() error
	Flush()
}

type Task struct {
	Id             int64
	Duration       time.Duration
	Implementation TASK_IMPLEMENTATION
	isRunning      bool
	isFinished     bool
	Eot            chan struct{}
	Response       *TaskResponse
	ResponseChan   chan *TaskResponse
}

type TaskResponse struct {
	Result interface{}
	Err    error
}

func (t *Task) GetID() int64 {
	return t.Id
}

func (t *Task) GetDuration() time.Duration {
	return t.Duration
}

func (t *Task) Run(args []interface{}) error {
	if !t.isRunning {
		go func(t1 *Task) {
			// Panic catching
			defer func() {
				t.isRunning = false
				var err_msg string
				if r := recover(); r != nil {
					err_msg = fmt.Sprintf("Task %d has finished with PANIC by %#v at %s with stack trace: %s\n", t.Id, r, fmt.Sprint(time.Now()), debug.Stack())
					if msg, err := lib_gc_event.NotifyEvent("041-001", "", &[]string{err_msg}); err != nil {
						fmt.Println("<", err.Error(), ">", err_msg)
					} else {
						lib_gc_event_publisher.EventPublisherChannel <- msg
					}
				}
			}()

			t.isRunning = true
			t.Response.Result, t.Response.Err = t.Implementation(t.Eot, &taskManager{t}, args)
			t.Flush()
		}(t)
	}
	return nil
}

func (t *Task) IsRunning() bool {
	return t.isRunning
}

func (t *Task) IsFinished() bool {
	return t.isFinished
}

func (t *Task) Finalize() error {
	close(t.Eot)
	t.isFinished = true
	return nil
}

func (t *Task) GetResponseChan() chan *TaskResponse {
	return t.ResponseChan
}

func (t *Task) Flush() {
	t.ResponseChan <- t.Response
}

// TaskImplmentation_I interface
func (t *Task) Finish() {
	t.Flush()
	t.Finalize()
}

// taking a Task to be waked up by a runner instance
// Parameters:
// 		id: Task's identification. It must be unique per task
//		duration: Maximum duration of the task. The task's implementation function is responsible for catching
//				  the closing of the EOT channel and process these event as a finalization request. See the test
//				  cases for more details about it.
//                **special case**: 0 duration stands for infinite task
//		implementation: Implementation function of the task. It my be complaint with the TASK_IMPLEMENTATION type
func GetTask(id int64, duration time.Duration, implementation TASK_IMPLEMENTATION) (Task_I, error) {
	return &Task{Id: id,
		Duration:       duration,
		Implementation: implementation,
		isRunning:      false,
		isFinished:     false,
		Eot:            make(chan struct{}),
		ResponseChan:   make(chan *TaskResponse, 2),
		Response:       &TaskResponse{}}, nil
}

//
// Task finisher
//
type TaskManager_I interface {
	Finish()
	Flush()
}

type taskManager struct {
	task Task_I
}

// This finish the task
func (tf *taskManager) Finish() {
	tf.task.Flush()
	tf.task.Finalize()
}

// This flushes the task, by does not ends it
func (tf *taskManager) Flush() {
	tf.task.Flush()
}
