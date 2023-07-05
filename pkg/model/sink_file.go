package model

import (
	"encoding/base64"
	"io"
	"os"
)

type SinkToFile struct {
	file io.Writer
}

func NewSinkToFile(filename string) SinkProcess {
	file, _ := os.Create(filename)
	return &SinkToFile{file}
}

func (sink *SinkToFile) Open() error {
	return nil
}

func (sink *SinkToFile) ProcessDictionary(dictionary Entry) error {
	value, err := base64.StdEncoding.DecodeString(string(dictionary.([]byte)))
	if err != nil {
		return err
	}

	_, err = sink.file.Write(append(value, '\n'))
	if err != nil {
		return err
	}
	return nil
}
