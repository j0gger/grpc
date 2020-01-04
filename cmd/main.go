package main

import (
	"fmt"
	"log"
	"os"

	"github.com/j0gger/grpc/pkg/client"
	"github.com/j0gger/grpc/pkg/server"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %v <server|client>", os.Args[0])
	os.Exit(1)
}

func main() {
	var allowdCmds = map[string]bool{
		"server": true,
		"client": true,
	}
	if len(os.Args) < 2 {
		usage()
	}

	if _, ok := allowdCmds[os.Args[1]]; !ok {
		usage()
	}

	if os.Args[1] == "server" {
		if err := server.RunServer(); err != nil {
			log.Fatal(err)
		}
	} else {
		resp, err := client.MakeRequest()
		if err != nil {
			log.Fatalf("Request failed: %v", err)
		} else {
			fmt.Println(resp)
		}
	}
}
