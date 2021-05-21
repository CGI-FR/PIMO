package model_test

import (
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSimple(t *testing.T) {
	jsonl := `{"b":0,"a":1}`

	dict := model.NewDictionary()
	err := dict.UnmarshalJSON([]byte(jsonl))
	assert.Nil(t, err)

	b, err := dict.MarshalJSON()
	assert.Nil(t, err)

	assert.Equal(t, jsonl, string(b))
}

func TestComplex(t *testing.T) {
	jsonl := `{"c":"","b":"","a":{"3":null,"2":null,"1":null}}`

	dict := model.NewDictionary()
	err := dict.UnmarshalJSON([]byte(jsonl))
	assert.Nil(t, err)

	b, err := dict.MarshalJSON()
	assert.Nil(t, err)

	assert.Equal(t, jsonl, string(b))
}
