package transfer

import (
	"crypto/tls"
)

var (
	Go     = X()
	Netcfg = Go.NetConf()
	buffer = make([]byte, 1024)
)

type Mango struct {
	Cfg Config
}

type Config struct {
	Transport *tls.Config
	Database  struct {
		Db     string `json:"dbname"`
		DbType string `json:"dbtype"`
		Port   int    `json:"port"`
	} `json:"DataBase"`

	DbAuth struct {
		ServerIP []string `json:"ServerIPv4"`
		User     string   `json:"username"`
		Pass     string   `json:"password"`
	} `json:"DbAuth"`
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
