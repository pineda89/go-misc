package lib_gc_fifo

import (
	"sync"
	"sync/atomic"
	"time"
)

var putMutex *sync.Mutex = &sync.Mutex{}

func GetFifo(sz int32) *Fifo {
	return &Fifo{c_fifo: make(chan interface{}, sz), size: 0, id: time.Now().UnixNano()}
}

type Fifo struct {
	id     int64
	c_fifo chan interface{}
	size   int32
}

func (s *Fifo) Empty() bool { return s.size == 0 }
func (s *Fifo) Peek() interface{} {
	putMutex.Lock()
	defer putMutex.Unlock()

	item := <-s.c_fifo
	s.c_fifo <- item
	return item
}
func (s *Fifo) Len() int32 { return s.size }

func (s *Fifo) Put(i interface{}) {
	putMutex.Lock()
	defer putMutex.Unlock()

	s.size = atomic.AddInt32(&s.size, int32(1))
	s.c_fifo <- i
}
func (s *Fifo) Pop() interface{} {
	putMutex.Lock()
	defer putMutex.Unlock()

	s.size = atomic.AddInt32(&s.size, int32(-1))
	item := <-s.c_fifo
	return item
}