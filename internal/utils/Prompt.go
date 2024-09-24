package utils

import (
	"fmt"
	"net"
	"time"
)

func Prompt(conn net.Conn, name string) {
	now := FormatDate(time.Now())
	prompt := fmt.Sprintf("[%v][%v]:", now, name)
	fmt.Fprint(conn, prompt)
}
