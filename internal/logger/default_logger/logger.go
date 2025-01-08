package defaultlogger

import "github.com/besean163/gophkeeper/internal/logger"

type Logger struct{}

func NewDefaultLogger() *Logger {
	return &Logger{}
}

func (n *Logger) Info(msg string, fields ...logger.Field)  {}
func (n *Logger) Error(msg string, fields ...logger.Field) {}
func (n *Logger) Debug(msg string, fields ...logger.Field) {}
