package lib_gc_conf

import (
	"fmt"
	"os"
	"strings"
)

func init() {
	//flag.StringVar(&CONF_PREFIX, "loaderConfigMisc", ".", "Is necessary to spcify the application configuration directory: -loaderConfigMisc ./myconf")
	//flag.Parse()
	setLoaderConf()
}

var Dummy struct{}

var CONF_PREFIX string = ""

const CONF_PREFIX_NAME = "loaderConfig"

const CONF_FILE_NAME = "LoaderConfiguration.ini"

func setLoaderConf() {
	for _, k := range os.Args {
		if strings.Contains(k, CONF_PREFIX_NAME) {
			CONF_PREFIX = k[strings.Index(k, CONF_PREFIX_NAME)+len(CONF_PREFIX_NAME)+1:]
		}
	}
	fmt.Println("Set loader go-misc library configuration directory to: ", CONF_PREFIX)
}
