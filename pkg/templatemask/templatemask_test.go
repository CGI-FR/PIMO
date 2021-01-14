package templatemask

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

func TestMaskingShouldReplaceSensitiveValueByTemplate(t *testing.T) {
	template := "{{.name}}.{{.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	config := model.NewMaskConfiguration().
		WithEntry("mail", tempMask)
	maskingEngine := model.MaskingEngineFactory(config, true)

	data := model.Dictionary{"name": "Jean", "surname": "Bonbeur", "mail": "jean44@outlook.com"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"name": "Jean", "surname": "Bonbeur", "mail": "Jean.Bonbeur@gmail.com"}
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestMaskingShouldReplaceSensitiveValueByTemplateInNested(t *testing.T) {
	template := "{{.customer.identity.name}}.{{.customer.identity.surname}}@gmail.com"
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	config := model.NewMaskConfiguration().
		WithEntry("mail", tempMask)
	maskingEngine := model.MaskingEngineFactory(config, true)

	data := model.Dictionary{"customer": model.Dictionary{"identity": model.Dictionary{"name": "Jean", "surname": "Bonbeur"}}, "mail": "jean44@outlook.com"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"customer": model.Dictionary{"identity": model.Dictionary{"name": "Jean", "surname": "Bonbeur"}}, "mail": "Jean.Bonbeur@gmail.com"}
	assert.Equal(t, waited, result, "Should create the right mail")
}

func TestFactoryShouldCreateAMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "mail"}, Mask: model.MaskType{Template: "{{.name}}.{{.surname}}@gmail.com"}}
	config, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, err, "error should be nil")
	maskingEngine, _ := NewMask("{{.name}}.{{.surname}}@gmail.com")
	assert.IsType(t, maskingEngine, config, "should be equal")
	assert.True(t, present, "should be true")
	assert.Nil(t, err, "error should be nil")
	data := model.Dictionary{"name": "Toto", "surname": "Tata", "mail": ""}
	result, err := maskingEngine.Mask(data, data)
	assert.Nil(t, err, "error should be nil")
	waitedResult := "Toto.Tata@gmail.com"
	assert.Equal(t, waitedResult, result, "Should create a right mail")
}

func TestFactoryShouldNotCreateAMaskFromAnEmptyConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be false")
	assert.Nil(t, err, "error should be nil")
}

func TestFactoryShouldReturnAnErrorInWrongConfig(t *testing.T) {
	maskingConfig := model.Masking{Mask: model.MaskType{Template: "{{.name}.{{.surname}}@gmail.com"}}
	mask, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, mask, "should be nil")
	assert.False(t, present, "should be true")
	assert.NotNil(t, err, "error shouldn't be nil")
}

func TestMaskingTemplateShouldFormat(t *testing.T) {
	template := `{{"hello!" | upper | repeat 2}}`
	tempMask, err := NewMask(template)
	if err != nil {
		assert.Fail(t, err.Error())
	}
	config := model.NewMaskConfiguration().
		WithEntry("field", tempMask)
	maskingEngine := model.MaskingEngineFactory(config, true)

	data := model.Dictionary{"field": "anything"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	waited := model.Dictionary{"field": "HELLO!HELLO!"}
	assert.Equal(t, waited, result, "Should create the right field")
}

func TestRFactoryShouldCreateANestedContextMask(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "foo.bar"}, Mask: model.MaskType{Template: "{{.baz}}"}}
	config := model.NewMaskConfiguration()
	maskEngine, present, err := Factory(maskingConfig, 0)
	assert.Nil(t, err, "should be nil")
	assert.True(t, present, "should be true")
	data := model.Dictionary{"baz": "BAZ", "foo": model.Dictionary{"bar": "BAR"}}

	config = config.WithEntry("foo.bar", maskEngine)

	masked, _ := model.MaskingEngineFactory(config, true).Mask(data, data)

	waited := model.Dictionary{"baz": "BAZ", "foo": model.Dictionary{"bar": "BAZ"}}
	assert.Equal(t, waited, masked, "Should replace foo.bar with BAZ")
}

func TestFactoryShouldReuseMaskedData(t *testing.T) {
	maskingConfig := model.Masking{Selector: model.SelectorType{Jsonpath: "foo.bar"}, Mask: model.MaskType{Template: "{{.baz}}"}}
	config := model.NewMaskConfiguration()

	maskEngine, present, err := Factory(maskingConfig, 0)
	config = config.WithEntry("foo.bar", maskEngine)

	assert.Nil(t, err, "should be nil")
	assert.True(t, present, "should be true")

	maskingConfig2 := model.Masking{Selector: model.SelectorType{Jsonpath: "foo.bar"}, Mask: model.MaskType{Template: "{{.foo.bar}}"}}
	maskEngine2, present, err := Factory(maskingConfig2, 0)
	config = config.WithEntry("foo.bar", maskEngine2)

	assert.Nil(t, err, "should be nil")
	assert.True(t, present, "should be true")

	data := model.Dictionary{"baz": "BAZ", "foo": model.Dictionary{"bar": "BAR"}}

	masked, _ := model.MaskingEngineFactory(config, true).Mask(data, data)

	waited := model.Dictionary{"baz": "BAZ", "foo": model.Dictionary{"bar": "BAZ"}}
	assert.Equal(t, waited, masked, "Should replace foo.bar with BAZ")
}
