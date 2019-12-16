package lib_gc_fifo

import (
	"testing"
	"time"
)

func Test_fifo(t *testing.T) {

	fifo := GetFifo(int32(2))

	go fifo.Put(4)
	go fifo.Put(3)
	go fifo.Put(2)
	go fifo.Put(1)

	go println(fifo.Pop().(int))
	go println(fifo.Pop().(int))
	go println(fifo.Pop().(int))
	go println(fifo.Pop().(int))

	time.Sleep(time.Duration(1) * time.Second)
}
