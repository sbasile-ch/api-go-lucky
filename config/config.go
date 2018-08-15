package config

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/companieshouse/chs.go/log"
	"github.com/ian-kent/gofigure"
)

var mtx sync.Mutex
var missing []string

//Config Struct
type config struct {
	gofigure         interface{} `order:"env,flag"`
	ApiKey           string      `env:"MY_CH_API" flag:"api-key" flagDesc:"API key to access CH API"`
	ServerIpAddr     string      `env:"AGL_SERVER_ADDR" flag:"server-addr" flagDesc:"Server IP Address"`
	ServerPort       uint        `env:"AGL_SERVER_PORT" flag:"server-port" flagDesc:"Server Port"`
	HttpReqTimeout   uint        `env:"AGL_HTTP_TIMEOUT" flag:"http-timeout" flagDesc:"seconds Timeout for HTTP requests"`
	HttpGetUseragent string      `env:"AGL_HTTP_UAGENT" flag:"user-agent" flagDesc:"User agent for HTTP requests"`
}

var cfg *config

//______________________________________
// Get configures the application and returns the configuration
func Get() (*config, error) {
	mtx.Lock()
	defer mtx.Unlock()

	if cfg != nil {
		return cfg, nil
	}

	cfg = &config{
		ServerIpAddr:   "127.0.0.1",
		ServerPort:     8080,
		HttpReqTimeout: 5,
	}

	err := gofigure.Gofigure(cfg)

	if err != nil {
		return nil, err
	}

	cfg.print()

	return cfg, nil
}

//______________________________________
func (c *config) print() {
	log.Trace("Configuration:")

	s := reflect.ValueOf(c).Elem()
	t := s.Type()

	ml := 0
	for i := 0; i < s.NumField(); i++ {
		if !s.Field(i).CanSet() {
			continue
		}
		if l := len(t.Field(i).Name); l > ml {
			ml = l
		}
	}

	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if !s.Field(i).CanSet() {
			continue
		}
		log.Trace(fmt.Sprintf("%s: %s", t.Field(i).Name, f.Interface()))
	}
}
