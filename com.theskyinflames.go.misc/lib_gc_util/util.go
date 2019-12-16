package lib_gc_util

import (
	"sync"

	C "github.com/glycerine/go-capnproto"

	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"expvar"
)

// Get a value of type interface{} as a string
func GetAsString(value interface{}) string {

	switch v := value.(type) {
	case string:
		return value.(string)
	case fmt.Stringer:
		return value.(fmt.Stringer).String()
	default:
		return fmt.Sprint(v)
	}
}

// Render time
func RenderTimeManual(t time.Time) string {
	//	return fmt.Sprintf("%d%02d%02d-%02d:%02d:%02d.%03d",
	//		t.Year(),
	//		t.Month(),
	//		t.Day(),
	//		t.Hour(),
	//		t.Minute(),
	//		t.Second(),
	//		t.Nanosecond()/100000)
	return fmt.Sprint(t)
}

// Render time
func RenderDurationManual(d time.Duration) string {
	//    return fmt.Sprintf("%02d:%02d:%02d.%03d",
	//    d.Hours(),
	//    d.Minutes(),
	//    d.Seconds(),
	//    d.Nanoseconds()/100000)
	return fmt.Sprint(d)
}

//
// Cap'N'Proto utilities
//
func GetCAPNPTextList(texts *[]string, segment *C.Segment) *C.TextList {
	tl := segment.NewTextList(len(*texts))
	for i, t := range *texts {
		tl.Set(i, t)
	}
	return &tl
}
func GetCAPNPDataList(data *[][]byte, segment *C.Segment) *C.DataList {
	dl := segment.NewDataList(len(*data))
	for i, d := range *data {
		dl.Set(i, d)
	}
	return &dl
}

func IsSliceContainsString(item string, slice *[]string) bool {
	contains := false
	for _, slice_item := range *slice {
		if item == slice_item {
			contains = true
			break
		}
	}
	return contains
}

//
// ---------------  Take environment variables -------------------------
//

const (
	ENVIRONMENT_VARIABLE_TYPE_STRING = iota
	ENVIRONMENT_VARIABLE_TYPE_BOOL
	ENVIRONMENT_VARIABLE_TYPE_INT64
	ENVIRONMENT_VARIABLE_TYPE_UINT64
	ENVIRONMENT_VARIABLE_TYPE_FLOAT64
)

type EnvironmentVariable struct {
	Var_name  string
	Var_type  int
	Var_value interface{}
}

func GetEnvironmentVariables(vars []*EnvironmentVariable) ([]*EnvironmentVariable, error) {
	for _, ev := range vars {
		if val := os.Getenv(ev.Var_name); val == "" {
			return nil, errors.New(fmt.Sprintf("Environment variable %s is not defined !!!", ev.Var_name))
		} else {
			if val != "" {
				var err error
				switch {
				case ev.Var_type == ENVIRONMENT_VARIABLE_TYPE_STRING:
					ev.Var_value = val
				case ev.Var_type == ENVIRONMENT_VARIABLE_TYPE_BOOL:
					ev.Var_value, err = strconv.ParseBool(val)
				case ev.Var_type == ENVIRONMENT_VARIABLE_TYPE_INT64:
					ev.Var_value, err = strconv.ParseInt(val, 10, 64)
				case ev.Var_type == ENVIRONMENT_VARIABLE_TYPE_UINT64:
					ev.Var_value, err = strconv.ParseUint(val, 10, 64)
				case ev.Var_type == ENVIRONMENT_VARIABLE_TYPE_FLOAT64:
					ev.Var_value, err = strconv.ParseFloat(val, 64)
				default:
					return nil, errors.New(fmt.Sprintf("Type %d not supported for environment variable %s !!!", ev.Var_type, ev.Var_name))
				}

				// Add the var as metric value
				z := expvar.NewString(ev.Var_name)
				z.Set(fmt.Sprint(ev.Var_value))

				if err != nil {
					return nil, errors.New(fmt.Sprintf("Something went wrong at try to retrieve the environment variable %s : %s", ev.Var_name, err.Error()))
				}
			}
		}
	}

	return vars, nil
}

//
// ---------------  locker -------------------------
//
var l_mutex *sync.Mutex = &sync.Mutex{}

type LockerFactory_I interface {
	GetLocker(name string) *Locker
}

var LockerFactory LockerFactory_I = &LockerFactoryImpl{make(map[string]*Locker)}

type LockerFactoryImpl struct {
	lockers map[string]*Locker
}

func (lf *LockerFactoryImpl) GetLocker(name string) *Locker {
	if locker, ok := lf.lockers[name]; !ok {
		locker = &Locker{make(map[string]string)}
		lf.lockers[name] = locker
		return locker
	} else {
		return locker
	}
}

type Locker struct {
	Locks map[string]string
}

func (_locker *Locker) Lock(key string) bool {
	l_mutex.Lock()
	defer l_mutex.Unlock()

	locked := false
	if _, ok := _locker.Locks[key]; !ok {
		_locker.Locks[key] = key
		locked = true
	}

	return locked
}

func (_locker *Locker) Unlock(key string) {
	l_mutex.Lock()
	defer l_mutex.Unlock()

	delete(_locker.Locks, key)
}
