// Copyright (C) 2024 CGI France
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

package timeline

import (
	"hash/fnv"
	"math/rand"
	"time"

	"github.com/cgi-fr/pimo/pkg/axis"
	"github.com/cgi-fr/pimo/pkg/duration"
	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	*axis.Generator
	format string
	points []string
	rand   *rand.Rand
	seeder model.Seeder
}

// NewMask create a MaskEngine
func NewMask(model model.TimeLineType, seed int64, seeder model.Seeder) (MaskEngine, error) {
	var err error

	origin := time.Now()

	if len(model.Start.Value) > 0 {
		origin, err = time.Parse(time.RFC3339, model.Start.Value)
		if err != nil {
			return MaskEngine{}, err
		}
	}

	generator := axis.NewGenerator(model.Start.Name, origin.Unix())

	if model.MaxRetry != nil {
		generator.SetMaxRetry(*model.MaxRetry)
	}

	points := make([]string, len(model.Points))
	for i, point := range model.Points {
		points[i] = point.Name

		pointMin, err := durationToSeconds(point.Min)
		if err != nil {
			return MaskEngine{}, err
		}

		pointMax, err := durationToSeconds(point.Max)
		if err != nil {
			return MaskEngine{}, err
		}

		constraints := []axis.Constraint{}

		var epsilon int64

		if model.Epsilon != "" {
			epsilon, err = durationToSeconds(model.Epsilon)
			if err != nil {
				return MaskEngine{}, err
			}
		}

		for _, constraint := range point.Constraints {
			behavior := axis.Replace // we will try to use the default value of the point
			if constraint.OnError == "reject" {
				behavior = axis.Reject
			}
			localEpsilon := epsilon
			if constraint.Epsilon != "" {
				localEpsilon, err = durationToSeconds(constraint.Epsilon)
				if err != nil {
					return MaskEngine{}, err
				}
			}
			if len(constraint.Before) > 0 {
				constraints = append(constraints, axis.LowerThan(constraint.Before, localEpsilon, behavior))
			}
			if len(constraint.After) > 0 {
				constraints = append(constraints, axis.GreaterThan(constraint.After, localEpsilon, behavior))
			}
		}

		generator.SetPoint(point.Name, point.From, pointMin, pointMax, point.Default, constraints...)
	}

	return MaskEngine{
		Generator: generator,
		format:    model.Format,
		points:    points,
		rand:      rand.New(rand.NewSource(seed)), //nolint:gosec
		seeder:    seeder,
	}, nil
}

// Mask generate a timeline
func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask timeline")
	if e == nil {
		return e, nil
	}

	if len(context) > 0 {
		seed, ok, err := me.seeder(context[0])
		if err != nil {
			return nil, err
		}
		if ok {
			me.rand.Seed(seed)
		}
	}

	initialState := map[string]*int64{}
	if dict, ok := e.(model.Dictionary); ok {
		for _, key := range dict.Keys() {
			initialState[key] = me.formatTimestamp(dict.Get(key))
		}
	}

	timestamps, err := me.Generate(me.rand, initialState)
	if err != nil {
		return nil, err
	}

	result := model.NewDictionary()
	result.Set(me.Generator.Origin(), me.formatDate(*(timestamps[me.Generator.Origin()])))
	for _, point := range me.points {
		value := timestamps[point]
		if value != nil {
			result.Set(point, me.formatDate(*value))
		} else {
			result.Set(point, nil)
		}
	}

	return result, nil
}

// Create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.TimeLine.Start.Name) > 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64())
		mask, err := NewMask(conf.Masking.Mask.TimeLine, conf.Seed, seeder)
		if err != nil {
			return nil, false, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}

func (me MaskEngine) formatDate(timestamp int64) model.Entry {
	if me.format == "unixEpoch" {
		return timestamp
	} else if me.format != "" {
		return time.Unix(timestamp, 0).Format(me.format)
	}

	return time.Unix(timestamp, 0).Format(time.RFC3339)
}

func (me MaskEngine) formatTimestamp(date model.Entry) *int64 {
	switch typedDate := date.(type) {
	case nil:
		return nil
	case string:
		t, err := time.Parse(me.format, typedDate)
		if err != nil {
			return nil
		}
		timestamp := t.Unix()
		return &timestamp
	default:
		return nil
	}
}

func durationToSeconds(iso8601 string) (int64, error) {
	dur, err := duration.ParseDuration(iso8601)
	if err != nil {
		return 0, err
	}

	return int64(dur.Seconds()), nil
}
