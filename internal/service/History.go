package service

import "fmt"

type History struct {
	Messages []Message
}

func (h *History) Push(m Message) {
	h.Messages = append(h.Messages, m)
}

func (h *History) PrintHistory() {
	for _, message := range h.Messages {
		fmt.Println(message)
	}
}
