package lib_loader_command

import "net/http"

type LoaderCommand_I interface {
    Do(r *http.Request, w http.ResponseWriter) error
}
