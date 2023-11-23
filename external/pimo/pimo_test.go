package pimo_test

import (
	"testing"

	"github.com/cgi-fr/pimo/external/pimo"
	"github.com/stretchr/testify/assert"
)

func TestNewContext(t *testing.T) {
	ctx, err := pimo.NewContext("{version: '1', seed: 42, masking: [{selector: {jsonpath: a}, mask: {regex: '0[1-7]( ([0-9]){2}){4}'}}]}")
	assert.Nil(t, err)

	result, err := ctx.ExecuteMap(map[string]string{"a": "1"})
	assert.Nil(t, err)

	assert.Equal(t, "02 03 37 35 83", result["a"])

	result, err = ctx.ExecuteMap(map[string]string{"a": "1"})
	assert.Nil(t, err)

	assert.Equal(t, "02 65 93 29 27", result["a"])
}

func TestExecuteJson(t *testing.T) {
	ctx, err := pimo.NewContext("{version: '1', seed: 40, masking: [{selector: {jsonpath: a}, mask: {regex: '0[1-7]( ([0-9]){2}){0,10}'}}]}")
	assert.Nil(t, err)

	input := []byte("{ \"a\": 1 }")
	output := make([]byte, 1024)

	jsonSize, err := pimo.ExecuteJSON(ctx, &input, &output)
	assert.Nil(t, err)

	assert.Equal(t, "{\"a\":\"02 78\"}", string(output[:jsonSize]))

	jsonSize, err = pimo.ExecuteJSON(ctx, &input, &output)
	assert.Nil(t, err)

	assert.Equal(t, "{\"a\":\"02\"}", string(output[:jsonSize]))
}
