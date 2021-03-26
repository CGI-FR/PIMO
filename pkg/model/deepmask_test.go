package model

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	i := 0
	fme := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}
	entry := "A"

	result, err := deepmask(fme, entry)

	assert.Nil(t, err)
	assert.Equal(t, "1", result)
}

func TestSlice(t *testing.T) {
	i := 0
	fme := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}
	entry := []Entry{"A", "B"}

	result, err := deepmask(fme, entry)

	assert.Nil(t, err)
	assert.Equal(t, []Entry{"1", "2"}, result)
}

func TestNestedSlice(t *testing.T) {
	i := 0
	fme := FunctionMaskEngine{Function: func(name Entry, contexts ...Dictionary) (Entry, error) { i++; return fmt.Sprintf("%d", i), nil }}
	entry := []Entry{[]Entry{"A", "B"}, []Entry{"C", "D"}}

	result, err := deepmask(fme, entry)

	assert.Nil(t, err)
	assert.Equal(t, []Entry{[]Entry{"1", "2"}, []Entry{"3", "4"}}, result)
}
