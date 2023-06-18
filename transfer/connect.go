package transfer

import (
	"fmt"
	"net"
)

func (Go *Mango) Dial(msg string) {
	data := Go.Config("dbconfig.json").DbAuth
	host, err := net.Dial(
		"tcp",
		data.ServerIP[0]+":"+data.ServerIP[1],
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer host.Close()
	Go.Send(host, msg, err)
}

func (Go *Mango) Send(host net.Conn, msg string, err error) {
	_, err = host.Write([]byte(msg))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func X() Mango {
	var x Mango
	return x
}
