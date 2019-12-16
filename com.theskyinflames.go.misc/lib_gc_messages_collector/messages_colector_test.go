package lib_gc_messages_collector

import (
	"errors"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_event"
)

/*
	 TEST INSTRUCTIONS !!!

	  BEFORE EXECUTING THE TEST, YOU MUST:

	    * LOAD THE ENVIRONMENT MESSAGES FROM THE SCRIPT ../../set_env_messages_eng.sh

		* LOAD THE ENVIRONMENT VARIABLES FROM THE SCRIPT ../../set_env_variables.sh

		* SET THE ENVIRONMENT VARIABLES:

           GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__PACKAGE_MAX_SIZE          (p.e to 10)
	       GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__PACKAGE_MAX_WAITING_TIME  (p.e to 100)
	       GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__PACKAGE_CHECK_INTERVAL    (p.e to 10)
	       GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__MAX_SENDING_RETRIES       (p.e to 3)
	       GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__SENDING_RETRIES_CHAN_SIZE (p.e to 100)
		   GO_COMMON__LIB_GC_MESSAGES_COLLECTOR__ECHO_SENDING_ERRORS       (p.e to true)
*/

var initFlag *sync.Once = &sync.Once{}

const SENDER_NAME = "MySender"
const SENDER_RETRIES_NAME = "MySender"

// Sender functions
var MySenderFunc MessageSenderFunc = func(pkg *Package) (*Package, error) {

	fmt.Printf("Received a package with %d messages (%t) \n", pkg.Queue.Len(), pkg.Queue.Empty())

	for c := 0; c < int(pkg.Queue.Len()); c++ {
		fmt.Printf("LOLO\n")
		msg := pkg.Queue.Pop().(*Message)
		fmt.Println("Message %d: %s", c, msg.Content.(string))
	}

	return nil, nil
}

var MySenderFunc_with_retries MessageSenderFunc = func(pkg *Package) (*Package, error) {

	fmt.Printf("Received a package with %d messages (%t) \n", pkg.Queue.Len(), pkg.Queue.Empty())

	c := 0
	for !pkg.Queue.Empty() {
		msg := pkg.Queue.Pop().(*Message)
		if msg.SendingRetries == 0 {
			msg.SendingRetries += 1
			pkg.Queue.Put(msg)
			return pkg, errors.New("Testing retries")
		} else {
			fmt.Println("Message %d: %s\n", c, msg.Content.(string))
		}
		c += 1
	}

	return nil, nil
}

func Init() {

	fmt.Println("Test initialization ...")

	// Load event messages
	lib_gc_event.LoadConfigurationFromEnv("GO_COMMON_LIB")

	// Adding the customized messages's sender to messages collector
	MessageSenderFactory.AddMessageSender(SENDER_NAME, MySenderFunc)
	MessageSenderFactory.AddMessageSender(SENDER_RETRIES_NAME, MySenderFunc_with_retries)
}

func MessageBuilder(text string, id int, sender string) *Message {
	return &Message{CreationTime: time.Now().UnixNano(), Sender: SENDER_NAME, Id: int64(id), Content: text}
}

func Test_Messages_Sending_OK(t *testing.T) {

	initFlag.Do(Init)

	t.Log("Sending messages")
	for c := 0; c < 100; c++ {
		MessagesCollectorChan <- MessageBuilder(fmt.Sprintf("Message %d", c), c, SENDER_NAME)
	}
	t.Log("Messages sent")

	time.Sleep(time.Duration(2000) * time.Millisecond)
}

func Test_Messages_Sending_with_retries(t *testing.T) {

	initFlag.Do(Init)

	t.Log("Sending messages")
	for c := 0; c < 100; c++ {
		MessagesCollectorChan <- MessageBuilder(fmt.Sprintf("Message %d", c), c, SENDER_RETRIES_NAME)
	}
	t.Log("Messages sent")

	time.Sleep(time.Duration(2000) * time.Millisecond)
}
