package lib_gc_event

import (
	"errors"
	"fmt"
	"testing"
)

func init() {

}

// Test for events loading from file
func Test_EventsLoaded(t *testing.T) {
	if ptrEH == nil {
		t.Error(errors.New("Events has not been loaded"))
	} else {
		fmt.Println(ptrEH.Configuration)
		for _, v := range ptrEH.Configuration.Events {
			t.Log("Loaded event: " + v.Code)
		}
		for k, v := range ptrEH.Configuration.EventMap {
			t.Log("Loaded event by code: " + k + " -> " + v.Code)
		}
	}
}

// Test events notification
func Test_NotifyEvent(t *testing.T) {

	// Test for an existing event
	if msg, err := NotifyEvent("001-000", "", &[]string{"ItDoesExist"}); err != nil {
		t.Error(err)
	} else {
		t.Log(msg)
	}

	// Test for an existing event with customized message
	if msg, err := NotifyEvent("001-000", "Customized message for the event #  .", &[]string{"ItDoesExist"}); err != nil {
		t.Error(err)
	} else {
		t.Log(msg)
	}

	// Test for a not existing event
	if _, err := NotifyEvent("XXX-001", "", &[]string{"ItDoesNotExist"}); err != nil {
		t.Log("The event does not exist, the error it's ok: " + err.Error())
	} else {
		t.Error("The event dos not exist. It must fail !!!")
	}
}
