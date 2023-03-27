package config

import (
	"reflect"
	"strings"

	"github.com/CharanDetDev/go-port-adapter-unit-test/util/converse"
	"github.com/CharanDetDev/go-port-adapter-unit-test/util/logg"

	"github.com/spf13/viper"
)

var Env Config

func ConfigInit() bool {

	viper.AddConfigPath(".")
	viper.AddConfigPath("./util/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		logg.Printlogger("\t\t ***** Initail Config With Viper ERROR :: viper.ReadInConfig() *****", err)
		return false
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		logg.Printlogger("\t\t ***** Initail Config With Viper ERROR :: viper.Unmarshal(&Env) *****", err)
		return false
	}

	fn := reflect.ValueOf(&Env).Elem()
	for i := 0; i < fn.NumField(); i++ {
		value := converse.ParseToString(fn.Field(i).Interface())
		reflect.ValueOf(&Env).Elem().FieldByName(fn.Type().Field(i).Name).SetString(value)
	}

	return true

}

func ConfigInitForTest() {
	viper.AddConfigPath(".")
	// viper.AddConfigPath("../config") // for Database unit test
	viper.AddConfigPath("../util/config")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	viper.GetViper().SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	err := viper.ReadInConfig()
	if err != nil {
		logg.Printlogger("\t\t ***** Initail Config With Viper ERROR :: viper.ReadInConfig() *****", err)
		panic(err)
	}

	err = viper.Unmarshal(&Env)
	if err != nil {
		logg.Printlogger("\t\t ***** Initail Config With Viper ERROR :: viper.Unmarshal(&Env) *****", err)
		panic(err)
	}

	fn := reflect.ValueOf(&Env).Elem()
	for i := 0; i < fn.NumField(); i++ {
		value := converse.ParseToString(fn.Field(i).Interface())
		reflect.ValueOf(&Env).Elem().FieldByName(fn.Type().Field(i).Name).SetString(value)
	}

	// logg.PrintloggerJsonMarshalIndentHasHeader("", "", Env)

}
