package segment

import (
	"hash/fnv"
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	pipelines map[string]model.Pipeline
	seed      int64
	seeder    model.Seeder
}

func buildDefinition(masks []model.MaskType, globalSeed int64) model.Definition {
	definition := model.Definition{
		Version:   "1",
		Seed:      globalSeed,
		Functions: nil,
		Masking:   []model.Masking{},
		Caches:    nil,
	}

	for _, mask := range masks {
		definition.Masking = append(definition.Masking, model.Masking{
			Selector: model.SelectorType{Jsonpath: "."},
			Mask:     mask,
		})
	}

	return definition
}

// NewMask return a MaskEngine from a value
func NewMask(segment model.SegmentType, caches map[string]model.Cache, fns tmpl.FuncMap, seed int64, seeder model.Seeder, seedField string) (MaskEngine, error) {
	var err error

	pipelines := map[string]model.Pipeline{}

	for groupname, masks := range segment.Replace {
		definition := buildDefinition(masks, seed)
		pipeline := model.NewPipeline(nil)
		pipeline, _, err = model.BuildPipeline(pipeline, definition, caches, fns, "", "")
		if err != nil {
			return MaskEngine{}, err
		}

		pipelines[groupname] = pipeline
	}

	return MaskEngine{pipelines, seed, seeder}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask segments")

	return e, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if len(conf.Masking.Mask.Segment.Regex) > 0 {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64()) //nolint:gosec
		mask, err := NewMask(conf.Masking.Mask.Segment, conf.Cache, conf.Functions, conf.Seed, seeder, conf.Masking.Seed.Field)
		if err != nil {
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
