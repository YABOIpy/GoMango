package main

import (
	t "GoMango/transfer/Server"
)

func main() {
	go t.HostHTML("8080")
	t.HostPHP("8888")
}
