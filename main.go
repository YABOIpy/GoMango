package main

import (
	"GoMango/src"
	"GoMango/transfer"
	"fmt"
)

var (
	c = src.X()
	d = transfer.X()
)

func main() {
	var choice int
	fmt.Print(`
		[1] Server [2] Client
	`)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		d.Server_Listen()
	case 2:
		var msg string
		fmt.Print("</>")
		fmt.Scanln(&msg)
		d.Dial(msg)
	default:
		main()
	}

}
