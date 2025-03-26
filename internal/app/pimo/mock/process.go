package mock

import (
	"fmt"
	"os"
	"sync"

	"github.com/adrienaury/zeromdc"
	"github.com/cgi-fr/pimo/pkg/jsonline"
	"github.com/cgi-fr/pimo/pkg/model"
)

type BlackHoleSink struct{}

func (s BlackHoleSink) Open() error {
	return nil
}

func (s BlackHoleSink) ProcessDictionary(value model.Entry) error {
	return nil
}

type Processor struct {
	mutex    *sync.Mutex
	source   *model.MutableSource
	pipeline model.SinkedPipeline
}

func NewProcessor(maskingFile string, globalSeed *int64, caches map[string]model.Cache, cachesToLoad map[string]string) (*Processor, map[string]model.Cache, error) {
	pdef, err := model.LoadPipelineDefinitionFromFile(maskingFile)
	if err != nil {
		return nil, caches, err
	}

	if globalSeed != nil {
		pdef.SetSeed(*globalSeed)
	}

	for cacheName, cacheDef := range pdef.Caches {
		if path, ok := cachesToLoad[cacheName]; ok {
			if err := loadCache(cacheName, caches, path, cacheDef.Reverse, cacheDef.Unique); err != nil {
				return nil, caches, err
			}
		}
	}

	source := model.NewMutableSource()

	pipeline, caches, err := model.BuildPipeline(model.NewPipeline(source), pdef, caches, nil)
	if err != nil {
		return nil, caches, err
	}

	return &Processor{
		mutex:    &sync.Mutex{},
		source:   source,
		pipeline: pipeline.AddSink(BlackHoleSink{}),
	}, caches, nil
}

func (p *Processor) Process(dict model.Dictionary) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	p.source.SetValues(dict)

	err := p.pipeline.Run()

	zeromdc.ClearGlobalFields()

	return err
}

func loadCache(name string, caches map[string]model.Cache, path string, reverse bool, unique bool) error {
	cache, exists := caches[name]
	if !exists {
		if unique {
			caches[name] = model.NewUniqueMemCache()
		} else {
			caches[name] = model.NewMemCache()
		}
		cache = caches[name]
	}

	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	defer file.Close()

	pipe := model.NewPipeline(jsonline.NewSource(file))

	if reverse {
		reverseFunc := func(d model.Dictionary) (model.Dictionary, error) {
			reverse := model.NewDictionary()
			reverse.Set("key", d.Get("value"))
			reverse.Set("value", d.Get("key"))
			return reverse, nil
		}

		pipe = pipe.Process(model.NewMapProcess(reverseFunc))
	}

	err = pipe.AddSink(model.NewSinkToCache(cache)).Run()
	if err != nil {
		return fmt.Errorf("Cache %s not loaded : %s", name, err.Error())
	}
	return nil
}
