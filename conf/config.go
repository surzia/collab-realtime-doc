package conf

import (
	"fmt"
	"reflect"

	"github.com/go-gcfg/gcfg"
)

type ConfigItem interface {
	Check() error
}

type Config struct {
	NewsConf NewsConfig
}

func Load(confPath string) (*Config, error) {
	conf, err := load(confPath)
	if err != nil {
		return conf, err
	}

	err = conf.check()
	if err != nil {
		return conf, err
	}

	return conf, nil
}

func load(confPath string) (*Config, error) {
	conf := Config{}
	err := gcfg.ReadFileInto(&conf, confPath)

	return &conf, err
}

func (a *Config) check() error {
	t := reflect.TypeOf(a)
	v := reflect.ValueOf(a)

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		v = v.Elem()
	}

	fieldNum := t.NumField()
	for i := 0; i < fieldNum; i++ {
		if err := v.Field(i).Addr().Interface().(ConfigItem).Check(); err != nil {
			return fmt.Errorf("%v Check(): %v", t.Field(i).Name, err)
		}
	}

	return nil
}
