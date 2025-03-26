package model_test

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestMutableSource(t *testing.T) {
	source := model.NewMutableSource()
	assert.NotNil(t, source)
}

func TestMutableSourceNextShouldReturnFalseBeforeValueIsSetted(t *testing.T) {
	source := model.NewMutableSource()
	for range source.Values() {
		assert.Fail(t, "Should not have any value")
	}
}

func TestMutableSourceNextShouldReturnTrueAfterValueIsSetted(t *testing.T) {
	source := model.NewMutableSource()
	source.SetValues(model.NewDictionary())
	done := false
	for range source.Values() {
		if !done {
			done = true
		} else {
			assert.Fail(t, "Should have only one value")
		}
	}
}
