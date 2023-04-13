package transfer

import (
	"crypto/tls"
	"encoding/json"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"log"
	"os"
)

func (Go *Mango) Errs(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func (Go *Mango) Config(name string) Config {
	var config Config
	conf, err := os.Open(name)

	defer conf.Close()
	Go.Errs(err)

	xp := json.NewDecoder(conf)
	xp.Decode(&config)

	return config

}

func (Go *Mango) TransportCfg(conf *tls.Config) {
	Go.Cfg.Transport = conf
}

func (Go *Mango) NetConf() *HclConfig {
	var config HclConfig
	err := hclsimple.DecodeFile("config.hcl", nil, &config)
	if err != nil {
		log.Fatalf("Failed to load configuration: %s", err)
	}
	return &config
}
