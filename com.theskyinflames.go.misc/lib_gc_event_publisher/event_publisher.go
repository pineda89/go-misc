package lib_gc_event_publisher

import (
	"fmt"
	"strings"

	lib_config "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_config"
	lib_log "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_log"
	////L "bitbucket.org/itserhstourism/srv-logservice/com.serhs.util.logservice/domain"
)

func init() {
	/*
		LOG_SERVICE_ACTIVE = lib_config.PtrConfigProvider.GetPropertyBOOL("LOG_SERVICE", "LOG_SERVICE_ACTIVE")
		LOG_SERVICE_URL = lib_config.PtrConfigProvider.GetPropertySTRING("LOG_SERVICE", "LOG_SERVICE_URL")
		LOG_SERVICE_CHANNEL = lib_config.PtrConfigProvider.GetPropertySTRING("LOG_SERVICE", "LOG_SERVICE_CHANNEL_NAME")
		LOG_SERVICE_SENDER = lib_config.PtrConfigProvider.GetPropertySTRING("LOG_SERVICE", "LOG_SERVICE_SENDER")
	*/

	CHANNEL_BUFFER = lib_config.PtrConfigProvider.GetPropertyINT("LIB_EVENT_PUBLISHER", "CHANNEL_BUFFER")
	EVENT_LEVEL = lib_config.PtrConfigProvider.GetPropertySTRING("LIB_EVENT_PUBLISHER", "EVENT_LEVEL")

	// Start the logging messages sending
	go startPublisher()

	lib_log.Info.Println("lib_event_publisher package initialized. ( with log level ", EVENT_LEVEL, " )")
	println("Log level ", EVENT_LEVEL)
}

var EventPublisherChannel chan string = make(chan string, CHANNEL_BUFFER)

var LOG_SERVICE_ACTIVE bool
var LOG_SERVICE_URL string
var LOG_SERVICE_CHANNEL string
var LOG_SERVICE_SENDER string

var CHANNEL_BUFFER int
var cancel bool = false
var EVENT_LEVEL string

func ClosePublisherChannel() error {
	cancel = true
	return nil
}

func startPublisher() {
	defer func() {
		close(EventPublisherChannel)
	}()

	for !cancel {
		message, ok := <-EventPublisherChannel

		if !ok {
			cancel = true
		} else {

			if !strings.Contains(message, "Pulses") {

				go func(message string) {

					switch {
					case EVENT_LEVEL == "ERROR":
						lib_log.Error.Printf("%s\n", fmt.Sprint(message))
					case EVENT_LEVEL == "WARNING":
						lib_log.Warning.Printf("%s\n", fmt.Sprint(message))
					case EVENT_LEVEL == "INFO":
						lib_log.Info.Printf("%s\n", fmt.Sprint(message))
					case EVENT_LEVEL == "TRACE":
						lib_log.Trace.Printf("%s\n", fmt.Sprint(message))
					default:
						fmt.Printf("%s\n", fmt.Sprint(message))
					}

				}(message)
			}
		}
	}
}
