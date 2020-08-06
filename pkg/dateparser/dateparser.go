package dateparser

import (
	"fmt"
	"time"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is a type to change a date format
type MaskEngine struct {
	inputFormat  string
	outputFormat string
}

// NewMask create a MaskEngine
func NewMask(input, output string) MaskEngine {
	return MaskEngine{input, output}
}

// Mask change a time format
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	var t time.Time
	var err error
	if me.inputFormat != "" {
		timestring := fmt.Sprintf("%v", e)
		t, err = time.Parse(me.inputFormat, timestring)
		if err != nil {
			return nil, err
		}
	} else {
		switch v := e.(type) {
		case string:
			t, err = time.Parse(time.RFC3339, v)
		case time.Time:
			t = v
		default:
			return e, fmt.Errorf("Field to mask is not a time nor a string")
		}
		if err != nil {
			return nil, err
		}
	}
	if me.outputFormat != "" {
		return t.Format(me.outputFormat), nil
	}
	return t, nil
}

// Create a mask from a configuration
func RegistryMaskToConfiguration(conf model.Masking, config model.MaskConfiguration, seed int64) (model.MaskConfiguration, bool, error) {
	if len(conf.Mask.DateParser.InputFormat) != 0 || len(conf.Mask.DateParser.OutputFormat) != 0 {
		mask := NewMask(conf.Mask.DateParser.InputFormat, conf.Mask.DateParser.OutputFormat)
		return config.WithEntry(conf.Selector.Jsonpath, mask), true, nil
	}
	return nil, false, nil
}
