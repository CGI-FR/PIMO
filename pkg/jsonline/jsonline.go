package jsonline

import (
	"bufio"
	"encoding/json"
	"io"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// NewSource creates a new Source.
func NewSource(file io.Reader) model.ISource {
	scanner := bufio.NewScanner(file)
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, 1024*1024*100) // increase buffer up to 100 MB
	return &Source{scanner, nil, nil}
}

// Source export line to JSON format.
type Source struct {
	fscanner *bufio.Scanner
	value    model.Dictionary
	err      error
}

// JSONToDictionary return a model.Dictionary from a jsonline
func JSONToDictionary(jsonline []byte) (model.Dictionary, error) {
	var inter interface{}
	err := json.Unmarshal(jsonline, &inter)
	if err != nil {
		return nil, err
	}
	dic := model.InterfaceToDictionary(inter)
	return dic, nil
}

func (s *Source) Open() error {
	return nil
}

// Next convert next line to model.Dictionary
func (s *Source) Next() bool {
	if !s.fscanner.Scan() {
		s.err = s.fscanner.Err()
		return false
	}
	line := s.fscanner.Bytes()
	s.value, s.err = JSONToDictionary(line)
	return s.err == nil
}

func (s *Source) Value() model.Dictionary {
	return s.value
}

func (s *Source) Err() error {
	return s.err
}

// NewSink creates a new Sink.
func NewSink(file io.Writer) model.ISinkProcess {
	return Sink{file}
}

type Sink struct {
	file io.Writer
}

func (s Sink) Open() error {
	return nil
}

func (s Sink) ProcessDictionary(dictionary model.Dictionary) error {
	jsonline, err := json.Marshal(dictionary)
	if err != nil {
		return err
	}
	jsonline = append(jsonline, "\n"...)

	_, err = s.file.Write(jsonline)
	if err != nil {
		return err
	}
	return nil
}
