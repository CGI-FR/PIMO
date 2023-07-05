package model

import (
	"io"
	"os"
)

type MsgLogger struct {
	file io.Writer
}

func NewMsgLogger(filename string) *MsgLogger {
	file, _ := os.Create(filename)
	return &MsgLogger{file}
}

func (logger *MsgLogger) Log(msg string) error {
	if _, err := logger.file.Write(append([]byte(msg), '\n')); err != nil {
		return err
	}
	return nil
}
