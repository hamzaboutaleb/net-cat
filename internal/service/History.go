package service

import (
	"fmt"
	"net"
	"sync"
)

type History struct {
	mu       sync.Mutex
	Messages []any
}

func (h *History) Push(m any) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.Messages = append(h.Messages, m)
}

func (h *History) PrintHistory(c net.Conn) {
	h.mu.Lock()
	defer h.mu.Unlock()
	for _, message := range h.Messages {
		fmt.Fprint(c, message)
	}
}
