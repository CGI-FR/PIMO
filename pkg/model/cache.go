// Copyright (C) 2021 CGI France
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

package model

import (
	"fmt"
)

type Cache interface {
	Get(key Entry) (Entry, bool)
	Put(key Entry, value Entry)
	Subscribe(key Entry, observer Observer)
	Iterate() Source
}

type UniqueCache interface {
	Cache
	PutUnique(key Entry, value Entry) bool
}

type Observer interface {
	Notify(key Entry, value Entry)
}

// MemCache is a cache in memory
type MemCache struct {
	cache     map[Entry]Entry
	observers map[Entry][]Observer
}

func (mc *MemCache) Iterate() Source {
	collector := NewCollector()
	for k, v := range mc.cache {
		entry := NewDictionary()
		entry.Set("key", k)
		entry.Set("value", v)
		collector.Collect(entry)
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

func (mc *MemCache) Subscribe(key Entry, observer Observer) {
	observers, ok := mc.observers[key]
	if !ok {
		observers = []Observer{}
	}
	mc.observers[key] = append(observers, observer)
}

func NewUniqueMemCache() UniqueCache {
	return &UniqueMemCache{MemCache{map[Entry]Entry{}, map[Entry][]Observer{}}, map[Entry]struct{}{}}
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

func NewMemCache() Cache {
	return &MemCache{map[Entry]Entry{}, map[Entry][]Observer{}}
}

// MaskCacheEngine is a struct to create a cahed mask
type MaskCacheEngine struct {
	Cache          Cache
	OriginalEngine MaskEngine
}

// NewMaskCacheEngine create an MaskCacheEngine
func NewMaskCacheEngine(cache Cache, original MaskEngine) MaskCacheEngine {
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

func (mce MaskCacheEngine) Name() string {
	return fmt.Sprintf("cached mask of %s", mce.OriginalEngine.Name())
}

type UniqueMaskCacheEngine struct {
	cache          UniqueCache
	originalEngine MaskEngine
	maxRetries     int
}

func NewUniqueMaskCacheEngine(cache UniqueCache, original MaskEngine) UniqueMaskCacheEngine {
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

func (umce UniqueMaskCacheEngine) Name() string {
	return fmt.Sprintf("unique cache of %s", umce.originalEngine.Name())
}

func NewFromCacheProcess(selector Selector, cache Cache) Processor {
	return &FromCacheProcess{selector, cache, &QueueCollector{}, map[Entry]*QueueCollector{}}
}

type FromCacheProcess struct {
	selector  Selector
	cache     Cache
	readiness *QueueCollector
	waiting   map[Entry]*QueueCollector
}

func (p *FromCacheProcess) Open() error {
	return nil
}

func (p *FromCacheProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	for p.readiness.Next() {
		p.processDictionary(p.readiness.Value(), out)
	}

	p.processDictionary(dictionary, out)
	return nil
}

func (p *FromCacheProcess) processDictionary(dictionary Dictionary, out Collector) {
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
			collector = &QueueCollector{}
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
