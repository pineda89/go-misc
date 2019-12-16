package lib_gc_restful

import (
	"net/http"
	"strconv"
)

var Dummy dummy

type dummy struct{}

type Administered struct{}

type AdministrationCommand struct {
	Path    string
	Handler *func(w http.ResponseWriter,r *http.Request)
}

type Administrable interface {
	StartAPI(port int, handlers *[]AdministrationCommand) error
}

func (administred *Administered) StartAPI(port int, handlers *[]AdministrationCommand) error {

	// Set handlers
	for _, handler := range *handlers {
		http.HandleFunc(handler.Path, *handler.Handler)
	}

	// Start the administration service
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		return err
	}
	return nil
}
