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

		generator.SetPoint(point.Name, point.From, pointMin, pointMax)
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

	timestamps := me.Generate(me.rand)

	result := model.NewDictionary()
	for _, point := range me.points {
		result.Set(point, me.formatDate(*(timestamps[point])))
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

func durationToSeconds(ISO8601 string) (int64, error) {
	dur, err := duration.ParseDuration(ISO8601)
	if err != nil {
		return 0, err
	}

	return int64(dur.Seconds()), nil
}
