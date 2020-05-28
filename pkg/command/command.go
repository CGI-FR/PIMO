package command

import (
	"os/exec"
	"strings"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine implements MaskEngine with a console command
type MaskEngine struct {
	Cmd string
}

// NewMask return a CommandMask from a value
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
		return nil, err
	}
	return resulting, nil
}

func NewMaskFromConfig(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Command) != 0 {
		return NewMask(conf.Mask.Command), true, nil
	}
	return nil, false, nil
}
