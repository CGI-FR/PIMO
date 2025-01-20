package logmask

import (
	"bytes"
	"fmt"
	"hash/fnv"
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	level   string
	message *template.Engine
}

// NewMask return a MaskEngine from a value
func NewMask(message, level string, caches map[string]model.Cache, fns tmpl.FuncMap, seed int64, seeder model.Seeder, seedField string) (MaskEngine, error) {
	var tmp *template.Engine
	var err error

	if message != "" {
		tmp, err = template.NewEngine(message, fns, seed, seedField)
		if err != nil {
			return MaskEngine{}, err
		}
	}

	return MaskEngine{
		level:   level,
		message: tmp,
	}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask log")

	var output bytes.Buffer

	if me.message != nil {
		if err := me.message.Execute(&output, context[0].UnpackUnordered()); err != nil {
			return nil, err
		}
	} else {
		output.WriteString(fmt.Sprintf("%v", e))
	}

	switch me.level {
	case "trace":
		log.Trace().Msg(output.String())
	case "debug":
		log.Debug().Msg(output.String())
	case "info":
		log.Info().Msg(output.String())
	case "warn":
		log.Warn().Msg(output.String())
	case "error":
		log.Error().Msg(output.String())
	default:
		log.Info().Msg(output.String())
	}

	return e, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.Log != nil {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64()) //nolint:gosec
		mask, err := NewMask(conf.Masking.Mask.Log.Message, conf.Masking.Mask.Log.Level, conf.Cache, conf.Functions, conf.Seed, seeder, conf.Masking.Seed.Field)
		if err != nil {
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
