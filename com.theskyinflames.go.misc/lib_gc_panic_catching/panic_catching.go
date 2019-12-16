package lib_gc_panic_catching

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func PanicCatching(functionName string, params ...string) {
	if r := recover(); r != nil {
		t := time.Now()
		if params == nil {
			params = []string{}
		}
		msg := fmt.Sprintf("%s PANIC at %s : PANIC Defered recover: %v. With params: %v.\n", t, functionName, r, params)
		fmt.Println(msg)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		msg = fmt.Sprintf("PANIC Stack Trace at %s : %s\n", functionName, string(buf))
		fmt.Println(msg)
	}
}

func PanicCatchingWithRecover(functionName string, r interface{}, params ...string) {
	if r != nil {
		t := time.Now()
		if params == nil {
			params = []string{}
		}
		msg := fmt.Sprintf("%s PANIC at %s : PANIC Defered recover: %v. With params: %v.\n", t, functionName, r, params)
		fmt.Println(msg)

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		msg = fmt.Sprintf("PANIC Stack Trace at %s : %s\n", functionName, string(buf))
		fmt.Println(msg)
	}
}

func PanicCatchingLogger(_logger *log.Logger, functionName string, params ...string) {
	if r := recover(); r != nil {
		t := time.Now()
		if params == nil {
			params = []string{}
		}
		msg := fmt.Sprintf("%s PANIC at %s : PANIC Defered recover: %v. With params: %v.\n", t, functionName, r, params)
		if _logger != nil {
			_logger.Println(msg)
		} else {
			fmt.Println(msg)
		}

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		msg = fmt.Sprintf("PANIC Stack Trace at %s : %s\n", functionName, string(buf))
		if _logger != nil {
			_logger.Println(msg)
		} else {
			fmt.Println(msg)
		}
	}
}

func PanicCatchingLoggerWithRecover(_logger *log.Logger, functionName string, r interface{}, params ...string) {
	if r != nil {
		t := time.Now()
		if params == nil {
			params = []string{}
		}
		msg := fmt.Sprintf("%s PANIC at %s : PANIC Defered recover: %v. With params: %v.\n", t, functionName, r, params)
		if _logger != nil {
			_logger.Println(msg)
		} else {
			fmt.Println(msg)
		}

		// Capture the stack trace
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)
		msg = fmt.Sprintf("PANIC Stack Trace at %s : %s\n", functionName, string(buf))
		if _logger != nil {
			_logger.Println(msg)
		} else {
			fmt.Println(msg)
		}
	}
}
