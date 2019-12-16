package lib_gc_event_from_env

import (
	"fmt"
	"os"
	"testing"

	"github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_configurable_from_env"
)

func init() {

	// Setting a temp
	s_t := "{\"code\": \"001_000\",\"level\": \"INFO\",\"message\": \"default message ** {{index . 0}}\"}"
	os.Setenv("PREFIX_001_001", s_t)
	fmt.Printf("Set env variable: 'PREFIX_001_001' %s\n", os.Getenv("PREFIX_001_001"))

	// Loding environment variables
	if err := lib_gc_configurable_from_env.LoadVariablesFromEnv(PREFIX); err != nil {
		panic(err)
	}
}

const PREFIX = "PREFIX"

func Test_TakeMsg(t *testing.T) {

	param := []string{"Bartolo"}
	if msg, err := GetMessageFromEnvTemplate(PREFIX, "001_001", param); err != nil {
		t.Errorf(err.Error())
		t.Fail()
	} else {
		t.Log(fmt.Sprintf("Retrieved message: %s", msg))
	}
}
