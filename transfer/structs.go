package transfer

import (
	"crypto/tls"
	"flag"
)

var (
	Go      = X()
	Netcfg  = Go.NetConf()
	Address = flag.String(
		Netcfg.Service.Processes.Command,
		Go.NetConf().Service.ListenAddr,
		"",
	)
)

type Mango struct {
	Cfg Config
}

type Config struct {
	Transport *tls.Config
}

type HclConfig struct {
	IOMode  string        `hcl:"io_mode"`
	Service ServiceConfig `hcl:"service,block"`
}

type ServiceConfig struct {
	Protocol   string        `hcl:"protocol,label"`
	Type       string        `hcl:"type,label"`
	ListenAddr string        `hcl:"listen_addr"`
	Processes  ProcessConfig `hcl:"process,block"`
}

type ProcessConfig struct {
	Type    string `hcl:"type,label"`
	Command string `hcl:"command"`
}
