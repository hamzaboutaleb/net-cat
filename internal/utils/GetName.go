package utils

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

func flushStdin(c net.Conn) {
	// Create a new reader for standard input
	buffer := make([]byte, 1024)
	n, _ := c.Read(buffer)

	for n == 1024 {
		println("n: ", n)
	}
}

func GetLine(c net.Conn) (string, error) {
	size := 10
	buffer := make([]byte, size)
	res := ""
	var err error
	n, err := c.Read(buffer)
	if err != nil {
		return res, err
	}
	if n == 1 && buffer[0] == 0 {
		println("end")
		return res, err
	}
	fmt.Println(n, err)
	if n == 0 {
		c.Write(ToBytes("\n"))
		return res, err
	}
	if n == size {
		flushStdin(c)
	}
	if buffer[n-1] != '\n' {
		c.Write(ToBytes("\neee"))
	} else {
		n--
	}
	res = string(buffer[:n])

	return res, err
}

func GetName(conn net.Conn) (string, error) {
	name, err := GetLine(conn)
	fmt.Println("here", err)

	if name == "" || strings.TrimSpace(name) == "" {
		return "", errors.New("please write something")
	}
	if len(name) >= 10 {
		return "", errors.New("name too long")
	}
	return string(name), nil
}
