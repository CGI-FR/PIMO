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
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "1", "name": "Bob", "supervisor": "2"}),
		NewDictionaryFromMap(map[string]Entry{"id": "2", "name": "John", "supervisor": "4"}),
		NewDictionaryFromMap(map[string]Entry{"id": "3", "name": "Tom", "supervisor": "2"}),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "bob", "name": "Bob", "supervisor": "john"}),
		NewDictionaryFromMap(map[string]Entry{"id": "tom", "name": "Tom", "supervisor": "john"}),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitForLoopProvid(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "1", "name": "Bob", "supervisor": "2"}),
		NewDictionaryFromMap(map[string]Entry{"id": "2", "name": "John", "supervisor": "3"}),
		NewDictionaryFromMap(map[string]Entry{"id": "3", "name": "Tom", "supervisor": "1"}),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "bob", "name": "Bob", "supervisor": "john"}),
		NewDictionaryFromMap(map[string]Entry{"id": "john", "name": "John", "supervisor": "tom"}),
		NewDictionaryFromMap(map[string]Entry{"id": "tom", "name": "Tom", "supervisor": "bob"}),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldUsedPreviouslyCachedValue(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	cache.Put("1", "boby")

	mySlice := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "1", "name": "Bob", "supervisor": "2"}),
		NewDictionaryFromMap(map[string]Entry{"id": "2", "name": "John", "supervisor": "3"}),
		NewDictionaryFromMap(map[string]Entry{"id": "3", "name": "Tom", "supervisor": "1"}),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "boby", "name": "Bob", "supervisor": "john"}),
		NewDictionaryFromMap(map[string]Entry{"id": "john", "name": "John", "supervisor": "tom"}),
		NewDictionaryFromMap(map[string]Entry{"id": "tom", "name": "Tom", "supervisor": "boby"}),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldReorderList(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "1", "name": "Bob", "supervisor": "2"}),
		NewDictionaryFromMap(map[string]Entry{"id": "4", "name": "Alice", "supervisor": "5"}),
		NewDictionaryFromMap(map[string]Entry{"id": "5", "name": "Rabbit", "supervisor": "5"}),
		NewDictionaryFromMap(map[string]Entry{"id": "2", "name": "John", "supervisor": "3"}),
		NewDictionaryFromMap(map[string]Entry{"id": "3", "name": "Tom", "supervisor": "1"}),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "alice", "name": "Alice", "supervisor": "rabbit"}),
		NewDictionaryFromMap(map[string]Entry{"id": "rabbit", "name": "Rabbit", "supervisor": "rabbit"}),
		NewDictionaryFromMap(map[string]Entry{"id": "bob", "name": "Bob", "supervisor": "john"}),
		NewDictionaryFromMap(map[string]Entry{"id": "john", "name": "John", "supervisor": "tom"}),
		NewDictionaryFromMap(map[string]Entry{"id": "tom", "name": "Tom", "supervisor": "bob"}),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitWithUnique(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewUniqueMemCache()

	mySlice := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "1", "name": "Bob", "supervisor": "2"}),
		NewDictionaryFromMap(map[string]Entry{"id": "2", "name": "John", "supervisor": "3"}),
		NewDictionaryFromMap(map[string]Entry{"id": "3", "name": "Tom", "supervisor": "1"}),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewUniqueMaskCacheEngine(cache, idMasking))).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache)).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionaryFromMap(map[string]Entry{"id": "bob", "name": "Bob", "supervisor": "john"}),
		NewDictionaryFromMap(map[string]Entry{"id": "john", "name": "John", "supervisor": "tom"}),
		NewDictionaryFromMap(map[string]Entry{"id": "tom", "name": "Tom", "supervisor": "bob"}),
	}
	assert.Equal(t, wanted, result)
}
