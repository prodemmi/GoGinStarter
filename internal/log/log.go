package log

import (
	"log"
	"os"
)

type Level int

const (
	DEBUG Level = iota
	INFO
	WARNING
	ERROR
)

type Log interface {
	Debug(message string)
	Info(message string)
	Warning(message string)
	Error(message string)
}

type lg struct {
	logger *log.Logger
}

func (log *lg) Debug(message string) {
	log.logger.Printf("[DEBUG] %s", message)
}

func (log *lg) Info(message string) {
	log.logger.Printf("[INFO] %s", message)
}

func (log *lg) Warning(message string) {
	log.logger.Printf("[WARNING] %s", message)
}

func (log *lg) Error(message string) {
	log.logger.Printf("[ERROR] %s", message)
}

func ProvideLog() Log {
	file, _ := os.OpenFile("storage/logs/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	return &lg{
		logger: log.New(file, "", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}
