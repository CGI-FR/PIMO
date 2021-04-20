// Copyright (C) 2021 CGI France
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

package command

import (
	"os/exec"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
)

// MaskEngine implements MaskEngine with a console command
type MaskEngine struct {
	Cmd string
}

// NewMask return a MaskEngine from a value
func NewMask(cmd string) MaskEngine {
	return MaskEngine{cmd}
}

// Mask delegate mask algorithm to an external program
func (cme MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	splitCommand := strings.Split(cme.Cmd, " ")
	/* #nosec */
	out, err := exec.Command(splitCommand[0], splitCommand[1:]...).Output()

	resulting := strings.Trim(string(out), "\n")
	if err != nil {
		return e, err
	}
	return resulting, nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64, caches map[string]model.Cache) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Command) != 0 {
		return NewMask(conf.Mask.Command), true, nil
	}
	return nil, false, nil
}
