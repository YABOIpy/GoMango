package transfer

import (
	"crypto/tls"
	"github.com/hashicorp/hcl/v2/hclsimple"
	"log"
)

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
