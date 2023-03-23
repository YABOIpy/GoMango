package transfer

import (
	"crypto/tls"
	"fmt"
)

func (Go *Mango) dial() {
	Go.TransportCfg(
		&tls.Config{
			InsecureSkipVerify: true,
		})

	host, err := tls.Dial(
		"tcp",
		*Address,
		Go.Cfg.Transport,
	)
	fmt.Print(host, err)

}

func X() Mango {
	var x Mango
	return x
}
