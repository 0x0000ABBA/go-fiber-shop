package logger

import (
	"log"
	"os"
)

func NewLogger(p string) *log.Logger {
	return log.New(os.Stdout, p, log.LstdFlags)
}
