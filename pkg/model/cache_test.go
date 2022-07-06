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
		NewDictionary().With("id", "1").With("name", "Bob").With("supervisor", "2"),
		NewDictionary().With("id", "2").With("name", "John").With("supervisor", "4"),
		NewDictionary().With("id", "3").With("name", "Tom").With("supervisor", "2"),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking), "")).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("id", "bob").With("name", "Bob").With("supervisor", "john"),
		NewDictionary().With("id", "tom").With("name", "Tom").With("supervisor", "john"),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitForLoopProvid(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionary().With("id", "1").With("name", "Bob").With("supervisor", "2"),
		NewDictionary().With("id", "2").With("name", "John").With("supervisor", "3"),
		NewDictionary().With("id", "3").With("name", "Tom").With("supervisor", "1"),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking), "")).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("id", "bob").With("name", "Bob").With("supervisor", "john"),
		NewDictionary().With("id", "john").With("name", "John").With("supervisor", "tom"),
		NewDictionary().With("id", "tom").With("name", "Tom").With("supervisor", "bob"),
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
		NewDictionary().With("id", "1").With("name", "Bob").With("supervisor", "2"),
		NewDictionary().With("id", "2").With("name", "John").With("supervisor", "3"),
		NewDictionary().With("id", "3").With("name", "Tom").With("supervisor", "1"),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking), "")).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("id", "boby").With("name", "Bob").With("supervisor", "john"),
		NewDictionary().With("id", "john").With("name", "John").With("supervisor", "tom"),
		NewDictionary().With("id", "tom").With("name", "Tom").With("supervisor", "boby"),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldReorderList(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewMemCache()

	mySlice := []Dictionary{
		NewDictionary().With("id", "1").With("name", "Bob").With("supervisor", "2"),
		NewDictionary().With("id", "4").With("name", "Alice").With("supervisor", "5"),
		NewDictionary().With("id", "5").With("name", "Rabbit").With("supervisor", "5"),
		NewDictionary().With("id", "2").With("name", "John").With("supervisor", "3"),
		NewDictionary().With("id", "3").With("name", "Tom").With("supervisor", "1"),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewMaskCacheEngine(cache, idMasking), "")).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("id", "alice").With("name", "Alice").With("supervisor", "rabbit"),
		NewDictionary().With("id", "rabbit").With("name", "Rabbit").With("supervisor", "rabbit"),
		NewDictionary().With("id", "bob").With("name", "Bob").With("supervisor", "john"),
		NewDictionary().With("id", "john").With("name", "John").With("supervisor", "tom"),
		NewDictionary().With("id", "tom").With("name", "Tom").With("supervisor", "bob"),
	}
	assert.Equal(t, wanted, result)
}

func TestFromCacheProcessShouldWaitWithUnique(t *testing.T) {
	idMasking := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) {
		return strings.ToLower(contexts[0].Get("name").(string)), nil
	}}

	cache := NewUniqueMemCache()

	mySlice := []Dictionary{
		NewDictionary().With("id", "1").With("name", "Bob").With("supervisor", "2"),
		NewDictionary().With("id", "2").With("name", "John").With("supervisor", "3"),
		NewDictionary().With("id", "3").With("name", "Tom").With("supervisor", "1"),
	}
	var result []Dictionary

	pipeline := NewPipelineFromSlice(mySlice).
		Process(NewMaskEngineProcess(NewPathSelector("id"), NewUniqueMaskCacheEngine(cache, idMasking), "")).
		Process(NewFromCacheProcess(NewPathSelector("supervisor"), cache, "")).
		AddSink(NewSinkToSlice(&result))
	err := pipeline.Run()

	assert.Nil(t, err)

	wanted := []Dictionary{
		NewDictionary().With("id", "bob").With("name", "Bob").With("supervisor", "john"),
		NewDictionary().With("id", "john").With("name", "John").With("supervisor", "tom"),
		NewDictionary().With("id", "tom").With("name", "Tom").With("supervisor", "bob"),
	}
	assert.Equal(t, wanted, result)
}
