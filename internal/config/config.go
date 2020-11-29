package config

import (
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strconv"

	"gopkg.in/yaml.v2"
)

type HandlerConfig struct {
	Prefix        string `yaml:"prefix" default:"/"`
	DumpHeaders   bool   `yaml:"dumpHeaders"`
	DumpBody      bool   `yaml:"dumpBody"`
	Code          int    `yaml:"code" default:"200"`
	ResponseByLua bool   `yaml:"responseByLua"`
	ResponseLua   string `yaml:"responseLua"`
	Type          string `yaml:"type" default:"http"`
	Root          string `yaml:"root" default:""`
}

type Config struct {
	Port     string          `yaml:"port" default:"8080"`
	Listen   string          `yaml:"listen" default:"0.0.0.0"`
	Handlers []HandlerConfig `yaml:"handlers"`
}

func validateConfig(cfg *Config) {
	rv := reflect.ValueOf(cfg).Elem()
	setStructValue(rv)
}

func setStringValue(value reflect.Value, defvalue string) {
	if value.String() == "" {
		value.SetString(defvalue)
	}
}

func setIntValue(value reflect.Value, defvaule string) {
	dv, _ := strconv.Atoi(defvaule)
	if value.Int() == 0 {
		value.SetInt(int64(dv))
	}
}

func setStructValue(value reflect.Value) {
	rt := value.Type()
	for i := 0; i < rt.NumField(); i++ {
		switch rt.Field(i).Type.Kind() {
		case reflect.String:
			setStringValue(value.Field(i), rt.Field(i).Tag.Get("default"))
		case reflect.Int:
			setIntValue(value.Field(i), rt.Field(i).Tag.Get("default"))
		case reflect.Slice:
			for j := 0; j < value.Field(i).Len(); j++ {
				v := value.Field(i).Index(j)
				setStructValue(v)
			}
		}
	}
}

func LoadConfig(filename string) (*Config, error) {
	var (
		cfg   *Config
		file  *os.File
		err   error
		bytes []byte
	)

	if file, err = os.Open(filename); err != nil {
		log.Println("Load Config: ", err)
		return nil, err
	}
	defer file.Close()

	if bytes, err = ioutil.ReadAll(file); err != nil {
		log.Println("Load Config: ", err)
		return nil, err
	}

	if err = yaml.Unmarshal(bytes, &cfg); err != nil {
		log.Println("Load Config: ", err)
		return nil, err
	}

	validateConfig(cfg)
	return cfg, nil
}
