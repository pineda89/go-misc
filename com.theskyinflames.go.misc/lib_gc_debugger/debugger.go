package lib_gc_debugger

import (
	"fmt"
	"io"
	"log"
	lib_config "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_config"
	"os"
	"time"
)

func init() {
	DEBUG = lib_config.PtrConfigProvider.GetPropertyBOOL("DEBUG", "DEBUG")
	DebuggerFactory = &debuggerFactory{}
}

var DEBUG bool
var DebuggerFactory DebbugerFactory_I

type DebbugerFactory_I interface {
	GetDebuggerInstance(chan_length int, file_path string) (Debugger_I, error)
}
type debuggerFactory struct{}

func (dbf *debuggerFactory) GetDebuggerInstance(chan_length int, file_path string) (Debugger_I, error) {
	debugger := &debugger{}
	if err := debugger.StartDebugger(chan_length, file_path); err != nil {
		return nil, err
	} else {
		return debugger, nil
	}
}

type Debugger_I interface {
	StartDebugger(chan_length int, file_path string) error
	Debug(msg string) error
	Shutdown()
}
type debugger struct {
	_file         io.Writer
	debug_chan    chan string
	shutdown_chan chan int
	started       bool
}

func (db *debugger) StartDebugger(chan_length int, file_path string) error {

	db.shutdown_chan = make(chan int)
	db.debug_chan = make(chan string, chan_length)

	if f, err := os.OpenFile(file_path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600); err != nil {
		return err
	} else {
		f.Chmod(os.FileMode(os.ModePerm))
		l := log.New(f, "DEBUG: ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
		go func(l *log.Logger, c <-chan string, shutdown <-chan int) {
			defer f.Close()
			for {
				select {
				case msg := <-c:
					l.Println(fmt.Sprint(time.Now()) + " >> " + msg + "\n")
				case <-shutdown:
					return
				}

			}
		}(l, db.debug_chan, db.shutdown_chan)
		db.started = true
	}
	return nil
}

func (db *debugger) Debug(msg string) error {
	db.debug_chan <- msg
	return nil
}
func (db *debugger) Shutdown() {
	close(db.shutdown_chan)
}
