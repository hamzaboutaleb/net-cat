package main

import (
	"time"

	"netcat/internal/service"
)

func main() {
	h := service.History{}
	h.Push("hello")
	h.Push(service.Message{
		Date:     time.Now(),
		Msg:      "hello",
		Username: "test",
	})
}
