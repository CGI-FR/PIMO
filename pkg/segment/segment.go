package segment

import (
	"hash/fnv"
	"regexp"
	"strings"
	tmpl "text/template"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/rs/zerolog/log"
)

type MaskEngine struct {
	re        *regexp.Regexp
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

	return MaskEngine{
		re:        regexp.MustCompile(segment.Regex),
		pipelines: pipelines,
		seed:      seed,
		seeder:    seeder,
	}, nil
}

// replace captured groups named in the `value` string using the values ​​calculated by the `replacements` map
func replace(value string, re *regexp.Regexp, replacements map[string]func(string) (string, error)) (string, error) {
	result := &strings.Builder{}

	matchIndexes := re.FindStringSubmatchIndex(value)
	groupNames := re.SubexpNames()

	writeCount := 0
	for i := 2; i < len(matchIndexes); i += 2 {
		groupNumber := i / 2
		groupName := groupNames[groupNumber]
		startIndex := matchIndexes[i]
		endIndex := matchIndexes[i+1]
		capturedValue := value[startIndex:endIndex]

		result.WriteString(value[writeCount:startIndex])
		writeCount = endIndex

		if replacement, exists := replacements[groupName]; exists {
			if masked, err := replacement(capturedValue); err != nil {
				return value, err
			} else {
				result.WriteString(masked)
			}
		}
	}
	result.WriteString(value[writeCount:])

	return result.String(), nil
}

func (me MaskEngine) Mask(e model.Entry, context ...model.Dictionary) (model.Entry, error) {
	log.Info().Msg("Mask segments")

	replacements := map[string]func(string) (string, error){}

	for groupname, pipeline := range me.pipelines {
		replacements[groupname] = func(match string) (string, error) {
			var result []model.Entry
			err := pipeline.
				WithSource(model.NewSourceFromSlice([]model.Dictionary{model.NewDictionary().With(".", match)})).
				AddSink(model.NewSinkToSlice(&result)).
				Run()
			if err != nil {
				return match, err
			}
			return result[0].(string), nil
		}
	}

	result, err := replace(e.(string), me.re, replacements)
	if err != nil {
		return e, err
	}

	return result, nil
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
