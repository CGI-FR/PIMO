package model

import (
	"errors"
	"math/rand"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMemCache_putUnique(t *testing.T) {
	cache := NewUniqueMemCache()

	result := cache.PutUnique("A", "1")

	assert.True(t, result)

	result = cache.PutUnique("B", "1")
	assert.False(t, result)

	result = cache.PutUnique("B", "2")
	assert.True(t, result)
}

func TestNewUniqueMaskCacheEngine(t *testing.T) {
	rand.Seed(0)
	MyFunc := func(entry Entry, contexts ...Dictionary) (Entry, error) {
		// nolint:gosec
		return rand.Int() % 2, nil
	}

	cachedMask := NewUniqueMaskCacheEngine(NewUniqueMemCache(), FunctionMaskEngine{MyFunc})

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

func TestFromCacheProcessShouldWaitForValueProvide(t *testing.T) {
	var idMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0]["name"].(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		{"id": "1", "name": "Bob", "supervisor": "2"},
		{"id": "2", "name": "John", "supervisor": "4"},
		{"id": "3", "name": "Tom", "supervisor": "2"},
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewSimplePathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"id": "bob", "name": "Bob", "supervisor": "john"},
		{"id": "tom", "name": "Tom", "supervisor": "john"},
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitForLoopProvid(t *testing.T) {
	var idMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0]["name"].(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		{"id": "1", "name": "Bob", "supervisor": "2"},
		{"id": "2", "name": "John", "supervisor": "3"},
		{"id": "3", "name": "Tom", "supervisor": "1"},
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewSimplePathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"id": "bob", "name": "Bob", "supervisor": "john"},
		{"id": "john", "name": "John", "supervisor": "tom"},
		{"id": "tom", "name": "Tom", "supervisor": "bob"},
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldUsedPreviouslyCachedValue(t *testing.T) {
	var idMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0]["name"].(string)), nil
	}}

	cache := NewMemCache()

	cache.Put("1", "boby")

	mySlice := []Dictionary{
		{"id": "1", "name": "Bob", "supervisor": "2"},
		{"id": "2", "name": "John", "supervisor": "3"},
		{"id": "3", "name": "Tom", "supervisor": "1"},
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewSimplePathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"id": "boby", "name": "Bob", "supervisor": "john"},
		{"id": "john", "name": "John", "supervisor": "tom"},
		{"id": "tom", "name": "Tom", "supervisor": "boby"},
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldReorderList(t *testing.T) {
	var idMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0]["name"].(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		{"id": "1", "name": "Bob", "supervisor": "2"},
		{"id": "4", "name": "Alice", "supervisor": "5"},
		{"id": "5", "name": "Rabbit", "supervisor": "5"},
		{"id": "2", "name": "John", "supervisor": "3"},
		{"id": "3", "name": "Tom", "supervisor": "1"},
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewSimplePathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"id": "alice", "name": "Alice", "supervisor": "rabbit"},
		{"id": "rabbit", "name": "Rabbit", "supervisor": "rabbit"},
		{"id": "bob", "name": "Bob", "supervisor": "john"},
		{"id": "john", "name": "John", "supervisor": "tom"},
		{"id": "tom", "name": "Tom", "supervisor": "bob"},
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitWithUnique(t *testing.T) {
	var idMasking = FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0]["name"].(string)), nil
	}}

	cache := NewUniqueMemCache()

	mySlice := []Dictionary{
		{"id": "1", "name": "Bob", "supervisor": "2"},
		{"id": "2", "name": "John", "supervisor": "3"},
		{"id": "3", "name": "Tom", "supervisor": "1"},
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewSimplePathSelector("id"), NewUniqueMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewSimplePathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		{"id": "bob", "name": "Bob", "supervisor": "john"},
		{"id": "john", "name": "John", "supervisor": "tom"},
		{"id": "tom", "name": "Tom", "supervisor": "bob"},
	}
	assert.Equal(t, wanted, result)
}
