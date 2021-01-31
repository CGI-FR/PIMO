package model

import (
	"fmt"
)

type ICache interface {
	Get(key Entry) (Entry, bool)
	Put(key Entry, value Entry)
	Subscribe(key Entry, observer IObserver)
	Iterate() ISource
}

type IUniqueCache interface {
	ICache
	PutUnique(key Entry, value Entry) bool
}

type IObserver interface {
	Notify(key Entry, value Entry)
}

// MemCache is a cache in memory
type MemCache struct {
	cache     map[Entry]Entry
	observers map[Entry][]IObserver
}

func (mc *MemCache) Iterate() ISource {
	collector := NewCollector()
	for k, v := range mc.cache {
		collector.Collect(Dictionary{"key": k, "value": v})
	}
	return collector
}

func (mc *MemCache) Get(key Entry) (Entry, bool) {
	value, ok := mc.cache[key]
	return value, ok
}

func (mc *MemCache) Put(key Entry, value Entry) {
	observers, ok := mc.observers[key]
	if ok {
		for _, observer := range observers {
			observer.Notify(key, value)
		}
	}
	mc.cache[key] = value
}

func (mc *MemCache) Subscribe(key Entry, observer IObserver) {
	observers, ok := mc.observers[key]
	if !ok {
		observers = []IObserver{}
	}
	mc.observers[key] = append(observers, observer)
}

func NewUniqueMemCache() IUniqueCache {
	return &UniqueMemCache{MemCache{map[Entry]Entry{}, map[Entry][]IObserver{}}, map[Entry]struct{}{}}
}

type UniqueMemCache struct {
	MemCache
	values map[Entry]struct{}
}

func (mc *UniqueMemCache) PutUnique(key Entry, value Entry) bool {
	_, valueAlerdyUse := mc.values[value]

	if valueAlerdyUse {
		return false
	}
	mc.Put(key, value)
	mc.values[value] = struct{}{}

	return true
}

func NewMemCache() ICache {
	return &MemCache{map[Entry]Entry{}, map[Entry][]IObserver{}}
}

// MaskCacheEngine is a struct to create a cahed mask
type MaskCacheEngine struct {
	Cache          ICache
	OriginalEngine MaskEngine
}

// NewMaskCacheEngine create an MaskCacheEngine
func NewMaskCacheEngine(cache ICache, original MaskEngine) MaskCacheEngine {
	return MaskCacheEngine{cache, original}
}

// Mask masks run mask with cache
func (mce MaskCacheEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	cachedValue, isInCache := mce.Cache.Get(e)
	if isInCache {
		return cachedValue, nil
	}
	value, err := mce.OriginalEngine.Mask(e, context...)
	if err == nil {
		mce.Cache.Put(e, value)
	}
	return value, err
}

type UniqueMaskCacheEngine struct {
	cache          IUniqueCache
	originalEngine MaskEngine
	maxRetries     int
}

func NewUniqueMaskCacheEngine(cache IUniqueCache, original MaskEngine) UniqueMaskCacheEngine {
	return UniqueMaskCacheEngine{cache, original, 1000}
}

// Mask masks run mask with cache
func (umce UniqueMaskCacheEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	cachedValue, isInCache := umce.cache.Get(e)
	if isInCache {
		return cachedValue, nil
	}
	for retry := 0; retry < umce.maxRetries; retry++ {
		value, err := umce.originalEngine.Mask(e, context...)
		if err == nil {
			ok := umce.cache.PutUnique(e, value)
			if ok {
				return value, nil
			}
		} else {
			return value, err
		}
	}
	return nil, fmt.Errorf("Unique value not found")
}

func NewFromCacheProcess(selector ISelector, cache ICache) IProcess {
	return &FromCacheProcess{selector, cache, &Collector{}, map[Entry]*Collector{}}
}

type FromCacheProcess struct {
	selector  ISelector
	cache     ICache
	readiness *Collector
	waiting   map[Entry]*Collector
}

func (p *FromCacheProcess) Open() error {
	return nil
}

func (p *FromCacheProcess) ProcessDictionary(dictionary Dictionary, out ICollector) error {
	for p.readiness.Next() {
		p.processDictionary(p.readiness.Value(), out)
	}

	p.processDictionary(dictionary, out)
	return nil
}

func (p *FromCacheProcess) processDictionary(dictionary Dictionary, out ICollector) {
	key, ok := p.selector.Read(dictionary)
	if !ok {
		return
	}

	value, inCache := p.cache.Get(key)
	if inCache {
		out.Collect(p.selector.Write(dictionary, value))
	} else {
		collector, exist := p.waiting[key]
		if !exist {
			collector = &Collector{}
			p.waiting[key] = collector
		}
		collector.Collect(dictionary)
		p.cache.Subscribe(key, p)
	}
}

func (p *FromCacheProcess) Notify(key Entry, value Entry) {
	collector, exist := p.waiting[key]
	if exist {
		for collector.Next() {
			p.readiness.Collect(collector.Value())
		}
	}
}
