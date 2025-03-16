package main

import (
	"os"
)

// Найти проблемы в коде, дописать функции которые ниже 33 строки
type Logger interface {
	Log(message string) error
	Close() error
}

type FileLogger struct{ file *os.File }

func NewFileLogger(fileName string) (*FileLogger, error) {
	f, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	return &FileLogger{f}, nil
}

func (f FileLogger) Log(message string) error {
	_, err := f.file.WriteString(message + "\n")
	return err
}

func (f FileLogger) Close() error {
	return f.file.Close()
}

///////

type SequentialLogger struct {
	wrppedLogger Logger
}

func NewSequentialLogger(wrppedLogger Logger) *SequentialLogger {

}

func (sl *SequentialLogger) Log(message string) error {

}
