package main

import (
	"fmt"
	"os"

	"netcat/internal/service"
	"netcat/internal/utils"
)

func main() {
	args := os.Args[1:]
	server := service.NewServer()
	if len(args) > 1 {
		fmt.Println("USAGE: go run . [PORT]")
		os.Exit(1)
	}
	if len(args) == 1 {
		port := args[0]
		if !utils.IsValidPort(port) {
			fmt.Println("Invalid Port")
			os.Exit(1)
		}
		server.SetPort(port)
	}
	if server.Start() != nil {
		fmt.Println("Error server initialization failed")
		os.Exit(1)
	}
	defer server.Listener.Close()
	server.Accept()
}
