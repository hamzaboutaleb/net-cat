package service

import (
	"fmt"
	"os"

	"netcat/internal/config"
)

type Logger struct {
	filename string
	file     *os.File
	isEnable bool
}

func InitLogger() *Logger {
	logger := &Logger{
		filename: config.LOGS_FILE,
		isEnable: false,
	}
	file, err := os.OpenFile("./logs/"+logger.filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0o600)
	if err != nil {
		fmt.Println("Open logs file failed")
		return nil
	}
	fmt.Println("Logger is Enable")
	logger.file = file
	logger.isEnable = true
	return logger
}

func (l *Logger) Append(text string) {
	if l.isEnable {
		l.file.WriteString(text)
	}
}
