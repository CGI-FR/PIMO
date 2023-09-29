package pimo

import (
	"io"

	"github.com/youen/xixo/pkg/xixo"
)

func ParseXML(input io.Reader, output io.Writer) {
	xixo.NewXMLParser(input, output)
}
