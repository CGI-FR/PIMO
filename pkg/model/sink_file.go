package model

import (
	"encoding/json"
	"fmt"
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

func (sink *SinkToFile) ProcessDictionary(dictionary Dictionary) error {
	fmt.Println(dictionary)
	jsonline, err := json.Marshal(dictionary)
	if err != nil {
		return err
	}
	jsonline = append(jsonline, "\n"...)

	_, err = sink.file.Write(jsonline)
	if err != nil {
		return err
	}
	return nil
}
