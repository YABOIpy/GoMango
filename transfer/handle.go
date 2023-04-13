package transfer

import (
	"fmt"
	"net"
	"strings"
)

func (Go *Mango) Server_Listen() {

	listener, err := net.Listen(
		"tcp",
		Go.NetConf().Service.ListenAddr,
	)
	Go.Errs(err)

	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go Go.Server_Read(conn)
	}
}

func (Go *Mango) Server_Read(conn net.Conn) (output string) {

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") == false {
				fmt.Println(err)
			}
			return
		}
		TcpMsg := string(buffer[:n])

		fmt.Println("Message:", TcpMsg, "\n")
		// send this data to db handler

	}

	return output
}
