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
		fmt.Println("[USAGE]: ./TCPChat $port")
		os.Exit(1)
	}
	if len(args) == 1 {
		port := args[0]
		if !utils.IsValidPort(port) {
			fmt.Println("Invalid Port\n[USAGE]: ./TCPChat $port")
			os.Exit(1)
		}
		server.SetPort(port)
	}
	if err := server.Start(); err != nil {
		fmt.Println("Error server initialization failed", err)
		os.Exit(1)
	}
	defer server.Listener.Close()
	server.Accept()
}
