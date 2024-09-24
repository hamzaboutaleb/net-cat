package service

import (
	"fmt"
	"time"

	"netcat/internal/utils"
)

type Message struct {
	Date     time.Time
	Msg      string
	Username string
}

func (m Message) String() string {
	formattedTime := utils.FormatDate(m.Date)
	return fmt.Sprintf("[%s][%s]:%s", formattedTime, m.Username, m.Msg)
}
