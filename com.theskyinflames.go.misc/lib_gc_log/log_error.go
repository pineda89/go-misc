// +build log_error

// +build !debug
// +build !trace

package lib_gc_log

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

func init() {
	//ioutil.Discard
	Init(ioutil.Discard, ioutil.Discard, ioutil.Discard, os.Stdout)
}

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

var LogLevel LOG_LEVEL = ERROR

func Init(
	traceHandle io.Writer,
	infoHandle io.Writer,
	warningHandle io.Writer,
	errorHandle io.Writer) {

	Trace = log.New(traceHandle,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	Info = log.New(infoHandle,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	Warning = log.New(warningHandle,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)

	Error = log.New(errorHandle,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
}
