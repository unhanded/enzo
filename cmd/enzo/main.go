package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/unhanded/enzo/internal/apply"
	"github.com/unhanded/enzo/internal/del"
	"github.com/unhanded/enzo/internal/get"
	"github.com/unhanded/enzo/internal/probe"
)

func main() {
	var fPath string
	flag.StringVar(&fPath, "f", "", "The file to submit")
	var port int
	flag.IntVar(&port, "p", 8080, "enzod port")
	var host string
	flag.StringVar(&host, "h", "127.0.0.1", "enzod host addr")

	flag.Parse()

	command := flag.Arg(0)

	switch command {
	case "probe":
		probe.Run(fPath, host, port)
	case "apply":
		apply.Run(fPath, host, port)
	case "delete":
		del.Run(flag.Arg(1), host, port)
	case "get":
		get.Run(host, port)
	default:
		fmt.Printf("No command found, exiting..\n")
		os.Exit(0)
	}
}
