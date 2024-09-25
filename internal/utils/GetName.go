package utils

import (
	"bufio"
	"errors"
	"net"
	"strings"
)

func GetName(conn net.Conn) (string, error) {
	scanner := bufio.NewScanner(conn)
	scanner.Scan()
	name := scanner.Text()
	if name == "" || strings.TrimSpace(name) == "" {
		return "", errors.New("please write something")
	}
	if len(name) >= 10 {
		return "", errors.New("name too long")
	}
	return string(name), nil
}
