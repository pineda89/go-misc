package com_theskyinflames_go_misc

import(
	"github.com/spf13/viper"
	"errors"
	"fmt"
)

const (
	ENV_VARIABLES_PREFIX = "SERHSPULSE"
)

func init(){
}

func GetEnvVariable(name string) (string, error){

	if value:=viper.Get(name);value==nil{
		return "",errors.New(fmt.Sprintf("Environment variable %s is not defined !!!",name))
	}else{
		return value.(string),nil
	}
}

func LoadVariablesFromEnv() error{

	viper.SetEnvPrefix(ENV_VARIABLES_PREFIX) // will be uppercased automatically
	viper.AutomaticEnv()

	return nil
}
