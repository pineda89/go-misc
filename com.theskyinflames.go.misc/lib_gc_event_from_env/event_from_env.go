package lib_gc_event_from_env

import (
	"bytes"
	"fmt"
	"sync"
	"text/template"

	ENV "github.com/pineda89/go-misc/com.theskyinflames.go.misc/lib_gc_configurable_from_env"
)

func init() {
	templates = make(map[string]*template.Template)
	mtx = &sync.RWMutex{}
}

var templates map[string]*template.Template
var mtx *sync.RWMutex

func GetMessageFromEnvTemplate(prefix, code string, paremters []string) (string, error) {

	t_key := fmt.Sprintf("%s_%s", prefix, code)
	var t *template.Template
	var _t string
	var err error
	var ok bool

	// Check for a previously loaded template
	if t, ok = templates[t_key]; !ok {
		// Try to retrieve the template from environment
		if _t, err = ENV.GetEnvVariable(code); err != nil {
			return "", err
		}
		if t, err = template.New(t_key).Parse(_t); err != nil {
			return "", err
		}

		// Adding the new template to the template's map
		mtx.Lock()
		templates[t_key] = t
		mtx.Unlock()
	}

	buff := &bytes.Buffer{}
	if err = t.ExecuteTemplate(buff, t_key, paremters); err != nil {
		return "", err
	}

	return buff.String(), nil
}
