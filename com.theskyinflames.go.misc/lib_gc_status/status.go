package lib_gc_status

/*
	Messages building using Go templates: https://play.golang.org/p/d_rGgTndmb
	[EM] GO_COMMON_LIB.500-001 = {\"code"\:\"500-001\",\"message\":\"Some went wrong at start the expvar http listener on [{{index .M 0}}] {{index .M 1}}: {{index .M 2}}\"}
*/

import (
	EVENT "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
	LOG "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	UTIL "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_util"

	"errors"
	"net"
	"net/http"
	"sync"
)

func init() {
	_network := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_STATUS__NETWORK", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}
	_addr := &UTIL.EnvironmentVariable{Var_name: "GO_COMMON__LIB_GC_STATUS__ADDR", Var_type: UTIL.ENVIRONMENT_VARIABLE_TYPE_STRING}

	evs := []*UTIL.EnvironmentVariable{_network, _addr}
	if _, err := UTIL.GetEnvironmentVariables(evs); err != nil {
		panic(err)
	} else {
		network = _network.Var_value.(string)
		addr = _addr.Var_value.(string)
	}
}

// Configuration variables
var network string
var addr string

// Status manager started flag
var started bool = false
var start_mutex *sync.Mutex = &sync.Mutex{}

func StartStatusExposer() error {
	start_mutex.Lock()
	defer start_mutex.Unlock()
	if !started {
		if sock, err := net.Listen(network, addr); err != nil {
			if msg, err := EVENT.NotifyEvent("500-001", "", &[]string{network, addr, err.Error()}); err != nil {
				return err
			} else {
				return errors.New(msg)
			}
		} else {
			go http.Serve(sock, nil)
			LOG.Info.Println("GO_COMMON.LIB_GC_STATUS LIB STARTED.")
			started = true
			return nil
		}
	} else {
		return nil
	}
}
