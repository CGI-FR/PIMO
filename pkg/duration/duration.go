package duration

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// MaskEngine is to mask a value thanks to a duration
type MaskEngine struct {
	duration time.Duration
}

// NewMask create a MaskEngine with a ISO8601 duration string
func NewMask(text string) (MaskEngine, error) {
	dur, err := ParseDuration(text)
	return MaskEngine{dur}, err
}

// ParseDuration parse a ISO8601 into time.Duration
// Modified from : https://stackoverflow.com/questions/28125963/golang-parse-time-duration
func ParseDuration(str string) (time.Duration, error) {
	durationRegex := regexp.MustCompile(`(-)?P(?P<years>\d+Y)?(?P<months>\d+M)?(?P<days>\d+D)?T?(?P<hours>\d+H)?(?P<minutes>\d+M)?(?P<seconds>\d+S)?`)
	matches := durationRegex.FindStringSubmatch(str)
	if len(matches) == 0 {
		return time.Duration(0), fmt.Errorf("%s isn't ISO 8601 duration", str)
	}

	years := ParseInt64(matches[2])
	months := ParseInt64(matches[3])
	days := ParseInt64(matches[4])
	hours := ParseInt64(matches[5])
	minutes := ParseInt64(matches[6])
	seconds := ParseInt64(matches[7])

	hour := int64(time.Hour)
	minute := int64(time.Minute)
	second := int64(time.Second)

	if matches[1] == "-" {
		return time.Duration(-1 * (years*24*365*hour + months*30*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second)), nil
	}
	return time.Duration(years*24*365*hour + months*30*24*hour + days*24*hour + hours*hour + minutes*minute + seconds*second), nil
}

func ParseInt64(value string) int64 {
	if len(value) == 0 {
		return 0
	}
	parsed, err := strconv.Atoi(value[:len(value)-1])
	if err != nil {
		return 0
	}
	return int64(parsed)
}

// Mask masks a time value with a duration
func (dura MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	var t time.Time
	var err error
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
	return t.Add(dura.duration), nil
}

// Create a mask from a configuration
func Factory(conf model.Masking, seed int64) (model.MaskEngine, bool, error) {
	if len(conf.Mask.Duration) != 0 {
		mask, err := NewMask(conf.Mask.Duration)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
