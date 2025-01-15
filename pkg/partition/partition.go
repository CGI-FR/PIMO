package partition

import (
	"bytes"
	"hash/fnv"
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	partitions []Partition
	seed       int64
	seeder     model.Seeder
}

type Partition struct {
	name string
	when *template.Engine
	exec model.Pipeline
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
func NewMask(partitions map[string]model.PartitionType, caches map[string]model.Cache, fns tmpl.FuncMap, seed int64, seeder model.Seeder, seedField string) (MaskEngine, error) {
	parts := []Partition{}

	// Build partitions pipelines
	for name, partition := range partitions {
		template, err := template.NewEngine(partition.When, fns, seed, seedField)
		if err != nil {
			return MaskEngine{}, err
		}

		if partition.When == "" {
			template = nil
		}

		definition := buildDefinition(partition.Then, seed)
		pipeline := model.NewPipeline(nil)
		pipeline, _, err = model.BuildPipeline(pipeline, definition, caches, fns, "", "")
		if err != nil {
			return MaskEngine{}, err
		}

		parts = append(parts, Partition{
			name: name,
			when: template,
			exec: pipeline,
		})
	}

	return MaskEngine{parts, seed, seeder}, nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask partition")

	// exec all partitions
	for _, partition := range me.partitions {
		var output bytes.Buffer

		if partition.when != nil {
			if err := partition.when.Execute(&output, context[0].UnpackUnordered()); err != nil {
				return nil, err
			}
		} else {
			output.WriteString("true")
		}

		if output.String() == "true" {
			log.Info().Msgf("Mask partition - executing partition %s", partition.name)

			break
		}
	}

	return e, nil
}

// Factory create a mask from a configuration
func Factory(conf model.MaskFactoryConfiguration) (model.MaskEngine, bool, error) {
	if conf.Masking.Mask.Partition != nil {
		seeder := model.NewSeeder(conf.Masking.Seed.Field, conf.Seed)

		// set differents seeds for differents jsonpath
		h := fnv.New64a()
		h.Write([]byte(conf.Masking.Selector.Jsonpath))
		conf.Seed += int64(h.Sum64()) //nolint:gosec
		mask, err := NewMask(conf.Masking.Mask.Partition, conf.Cache, conf.Functions, conf.Seed, seeder, conf.Masking.Seed.Field)
		if err != nil {
			return mask, true, err
		}
		return mask, true, nil
	}
	return nil, false, nil
}
