package model

import (
	"fmt"
)

type Cache interface {
	get(key Entry) (Entry, bool)
	put(key Entry, value Entry)
	putUnique(key Entry, value Entry) bool
}

// MemCache is a cache in memory
type MemCache struct {
	cache  map[Entry]Entry
	values map[Entry]struct{}
}

func (mc MemCache) get(key Entry) (Entry, bool) {
	value, ok := mc.cache[key]
	return value, ok
}

func (mc MemCache) put(key Entry, value Entry) { mc.cache[key] = value }

func (mc MemCache) putUnique(key Entry, value Entry) bool {
	_, valueAlerdyUse := mc.values[value]

	if valueAlerdyUse {
		return false
	}
	mc.cache[key] = value
	mc.values[value] = struct{}{}

	return true
}

func NewMemCache() MemCache {
	return MemCache{map[Entry]Entry{}, map[Entry]struct{}{}}
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
	cachedValue, isInCache := mce.Cache.get(e)
	if isInCache {
		return cachedValue, nil
	}
	value, err := mce.OriginalEngine.Mask(e, context...)
	if err == nil {
		mce.Cache.put(e, value)
	}
	return value, err
}

type UniqueMaskCacheEngine struct {
	MaskCacheEngine
	maxRetries int
}

func NewUniqueMaskCacheEngine(cache Cache, original MaskEngine) UniqueMaskCacheEngine {
	return UniqueMaskCacheEngine{MaskCacheEngine{cache, original}, 1000}
}

// Mask masks run mask with cache
func (umce UniqueMaskCacheEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	cachedValue, isInCache := umce.Cache.get(e)
	if isInCache {
		return cachedValue, nil
	}
	for retry := 0; retry < umce.maxRetries; retry++ {
		value, err := umce.OriginalEngine.Mask(e, context...)
		if err == nil {
			ok := umce.Cache.putUnique(e, value)
			if ok {
				return value, nil
			}
		} else {
			return value, err
		}
	}
	return nil, fmt.Errorf("Unique value not found")
}
