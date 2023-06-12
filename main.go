package main

import (
	t "GoMango/transfer/Server"

	"fmt"
)

var ()

// make a project tab to see all the available projects
// write documentation for it
func main() {
	t.Run()
	var choice int
	fmt.Print(`
		[1] Server [2] Client
	`)
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		//go d.ServerListen(mux.NewRouter())
	case 2:
		var msg string
		fmt.Print("</>")
		fmt.Scanln(&msg)
		//d.Dial(msg)
	default:
		main()
	}

}
