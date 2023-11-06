// Copyright (C) 2022 CGI France
//
// This file is part of PIMO.
//
// PIMO is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// PIMO is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with PIMO.  If not, see <http://www.gnu.org/licenses/>.

package pimo

import (
	"io"

	"github.com/CGI-FR/xixo/pkg/xixo"
)

func ParseXML(input io.Reader, output io.Writer) *xixo.XMLParser {
	return xixo.NewXMLParser(input, output).EnableXpath()
}
