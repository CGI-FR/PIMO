package markov

import (
	"math/rand"
	"strings"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Markov: model.MarkovType{MaxSize: 20, Sample: "pimo://nameFR"}}}
	_, present, err := Factory(maskingConfig, 0, nil)
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldCreateAChain(t *testing.T) {
	//nolint: gosec
	chain := NewChain(2, rand.New(rand.NewSource(42)), "")
	data := []string{"aaa", "bbb", "aba"}
	for _, el := range data {
		chain.Add(strings.Split(el, ""))
	}
	newState, err := chain.Generate([]string{"a", "b"})
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, "a", newState)
	_, err = chain.Generate([]string{"a"})
	assert.Error(t, err, "N-gram length does not match chain order")
}
