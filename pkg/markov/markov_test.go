package markov

import (
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
