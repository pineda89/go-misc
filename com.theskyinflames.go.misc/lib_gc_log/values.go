package lib_gc_log

import (
	"errors"
	"fmt"
)

type LOG_LEVEL int

const LOG_ITEM_SEPARATOR_PIPE = "|"

const (
	_ = iota
	TRACE
	INFO
	WARNING
	ERROR
)

func Log(level int, itemSeparator string, items ...string) error {
	msg := ""
	for z, item := range items {
		if z > 0 {
			msg = fmt.Sprintf("%s%s", LOG_ITEM_SEPARATOR_PIPE, item)
		} else {
			msg = item
		}
	}

	switch level {
	case TRACE:
		Trace.Println(msg)
	case INFO:
		Info.Println(msg)
	case WARNING:
		Warning.Println(msg)
	case ERROR:
		Error.Println(msg)
	default:
		return errors.New("Log level not supported !!!")
	}
	return nil
}
