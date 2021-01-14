package model

type Cache interface {
	get(key Entry) (Entry, bool)
	put(key Entry, value Entry)
}

type MemCache struct {
	cache map[Entry]Entry
}

func (mc MemCache) get(key Entry) (Entry, bool) {
	value, ok := mc.cache[key]
	return value, ok
}

func (mc MemCache) put(key Entry, value Entry) { mc.cache[key] = value }

func NewMemCache() MemCache {
	return MemCache{map[Entry]Entry{}}
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
