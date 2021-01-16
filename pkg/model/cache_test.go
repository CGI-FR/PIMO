package model

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemCache_putUnique(t *testing.T) {
	cache := NewMemCache()

	result := cache.putUnique("A", "1")

	assert.True(t, result)

	result = cache.putUnique("B", "1")
	assert.False(t, result)

	result = cache.putUnique("B", "2")
	assert.True(t, result)
}

func TestNewUniqueMaskCacheEngine(t *testing.T) {
	rand.Seed(0)
	MyFunc := func(entry Entry, contexts ...Dictionary) (Entry, error) {
		// nolint:gosec
		return rand.Int() % 2, nil
	}

	cachedMask := NewUniqueMaskCacheEngine(NewMemCache(), FunctionMaskEngine{MyFunc})

	masked, err := cachedMask.Mask("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, masked)

	masked, err = cachedMask.Mask("2")
	assert.Nil(t, err)
	assert.Equal(t, 0, masked)

	masked, err = cachedMask.Mask("1")
	assert.Nil(t, err)
	assert.Equal(t, 1, masked)

	masked, err = cachedMask.Mask("3")
	assert.Equal(t, err, errors.New("Unique value not found"))
	assert.Nil(t, masked)
}
