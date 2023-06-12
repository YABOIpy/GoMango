package transfer

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"os/exec"
	"strings"
)

func (Go *Mango) ServerListen(r *mux.Router) {
	listener, err := net.Listen(
		"tcp",
		Go.NetConf().Service.ListenAddr,
	)
	Go.Errs(err)

	defer listener.Close()
	for {
		conn, er := listener.Accept()
		if er != nil {
			fmt.Println(er)
			continue
		}
		go ServerRead(conn, r)
	}
}

func ServerRead(conn net.Conn, r *mux.Router) (output string) {

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if strings.Contains(err.Error(), "EOF") == false {
				fmt.Println(err)
			}
			return
		}
		TcpMsg := string(buffer[:n])
		if strings.Contains(TcpMsg, "--cmd") {
			cmd := exec.Command(TcpMsg[6:])
			data, er := cmd.Output()
			Go.Errs(er)

			output = string(data)
		} else {

			fmt.Println("Received message:", TcpMsg, "\n")
		}
	}

	return output
}
