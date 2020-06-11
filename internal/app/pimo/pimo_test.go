package pimo

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/command"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/constant"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/duration"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/hash"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randdate"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomint"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/randomlist"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/regex"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/replacement"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/templatemask"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/weightedchoice"
)

var nameMasking = model.FunctionMaskEngine{Function: func(name model.Entry) (model.Entry, error) { return "Toto", nil }}

func TestMaskingShouldReturnEmptyWhenInputISEmpty(t *testing.T) {
	maskingEngine := model.NewMaskConfiguration().AsEngine()
	data := model.Dictionary{}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, data, result, "should be empty")
}

func TestMaskingShouldNoReplaceInsensitiveValue(t *testing.T) {
	maskingEngine := model.NewMaskConfiguration().AsEngine()

	data := model.Dictionary{"project": "ER456"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.Equal(t, data, result, "should be equal")
}

func TestMaskingShouldReplaceSensitiveValue(t *testing.T) {
	config := model.NewMaskConfiguration().
		WithEntry("name", nameMasking)

	maskingEngine := model.MaskingEngineFactory(config)

	data := model.Dictionary{"name": "Benjamin"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	assert.NotEqual(t, data, result, "should be masked")
}

func TestMaskingShouldReplaceValueInNestedDictionary(t *testing.T) {
	config := model.NewMaskConfiguration().
		WithEntry("customer", model.NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := model.MaskingEngineFactory(config)
	data := model.Dictionary{"customer": model.Dictionary{"name": "Benjamin"}, "project": "MyProject"}
	result, err := maskingEngine.Mask(data)
	assert.Equal(t, nil, err, "error should be nil")
	want := model.Dictionary{"customer": model.Dictionary{"name": "Toto"}, "project": "MyProject"}
	assert.Equal(t, want, result, "should be masked")
}

func TestWithEntryShouldBuildNestedConfigurationWhenKeyContainsDot(t *testing.T) {
	config := model.NewMaskConfiguration().
		WithEntry("customer.name", nameMasking)
	_, ok := config.GetMaskingEngine("customer")
	assert.True(t, ok, "should be same configuration")

	_, okbis := config.GetMaskingEngine("customer.name")
	assert.False(t, okbis, "should be same configuration")
}

func TestYamlConfigShouldCreateEntriesInTheYamlOrder(t *testing.T) {
	maskingfile := "../../../test/maskingTest.yml"
	config, err := YamlConfig(maskingfile, []func(model.Masking, int64) (model.MaskEngine, bool, error){constant.NewMaskFromConfig, command.NewMaskFromConfig, regex.NewMaskFromConfig})
	if err != nil {
		t.Log(err)
	}
	actualEntries := config.Entries()
	waitedEntries := []string{"name", "phone"}
	assert.Equal(t, actualEntries, waitedEntries, "Should be the same")
}

func A(me model.MaskContextEngine, boo bool) model.MaskContextEngine {
	return me
}

func Must(me model.MaskEngine, err error) model.MaskEngine {
	return me
}

func TestShouldCreateAMaskConfigurationFromAFile(t *testing.T) {
	filename := "../../../test/masking.yml"
	factory := []func(model.Masking, int64) (model.MaskEngine, bool, error){
		constant.NewMaskFromConfig,
		command.NewMaskFromConfig,
		randomlist.NewMaskFromConfig,
		randomint.NewMaskFromConfig,
		weightedchoice.NewMaskFromConfig,
		regex.NewMaskFromConfig,
		hash.NewMaskFromConfig,
		randdate.NewMaskFromConfig,
		replacement.NewMaskFromConfig,
		duration.NewMaskFromConfig,
		templatemask.NewMaskFromConfig,
	}
	config, err := YamlConfig(filename, factory)
	if err != nil {
		t.Log(err)
	}
	regexp := "0[1-7]( ([0-9]){2}){4}"
	constMask := constant.NewMask("Toto")
	regmask, err := regex.NewMask(regexp, int64(42))
	if err != nil {
		assert.Fail(t, "regex could not initialise")
	}
	randList := randomlist.NewMask([]model.Entry{"Mickael", "Mathieu", "Marcelle"}, int64(42))
	nameWeighted := []model.WeightedChoiceType{{Choice: "Dupont", Weight: 9}, {Choice: "Dupond", Weight: 1}}
	dateMin := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	dateMax := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	templ, err := templatemask.NewMask("{{.surname}}.{{.name}}@gmail.com")
	if err != nil {
		assert.Fail(t, err.Error())
	}
	waited := model.NewMaskConfiguration().
		WithEntry("customer.phone", regmask).
		WithEntry("name", constMask).
		WithEntry("name2", randList).
		WithEntry("age", randomint.NewMask(25, 32, int64(42))).
		WithEntry("name3", command.MaskEngine{Cmd: "echo Dorothy"}).
		WithEntry("surname", weightedchoice.NewMask(nameWeighted, int64(42))).
		WithEntry("address.town", hash.MaskEngine{List: []model.Entry{"Emerald City", "Ruby City", "Sapphire City"}}).
		WithEntry("date", randdate.NewMask(dateMin, dateMax, 42)).
		WithEntry("name4", replacement.NewMask("name1")).
		WithEntry("mail", templ).
		WithEntry("last_contact", Must(duration.NewMask("-P60D")))

	waitedOrder := []string{"customer", "name", "name2", "age", "name3", "surname", "address", "date", "name4", "mail", "last_contact"}
	actualOrder := config.Entries()
	assert.Equal(t, waitedOrder, actualOrder, "Should be the same order")
	assert.Equal(t, A(config.GetMaskingEngine("customer.phone")), A(waited.GetMaskingEngine("customer.phone")), "customer.phone not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name")), A(waited.GetMaskingEngine("name")), "name1 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name2")), A(waited.GetMaskingEngine("name2")), "name2 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("age")), A(waited.GetMaskingEngine("age")), "age not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name3")), A(waited.GetMaskingEngine("name3")), "name3 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("surname")), A(waited.GetMaskingEngine("surname")), "surname not equal")
	assert.Equal(t, A(config.GetMaskingEngine("address.town")), A(waited.GetMaskingEngine("address.town")), "surname not equal")
	assert.Equal(t, A(config.GetMaskingEngine("date")), A(waited.GetMaskingEngine("date")), "date not equal")
	assert.Equal(t, A(config.GetMaskingEngine("mail")), A(waited.GetMaskingEngine("mail")), "mail not equal")
	assert.Equal(t, A(config.GetMaskingEngine("last_contact")), A(waited.GetMaskingEngine("last_contact")), "last_contact not equal")
}

func TestShouldReturnAnErrorWithMultipleArguments(t *testing.T) {
	filename := "../../../test/wrongMasking.yml"
	conf, err := YamlConfig(filename, []func(model.Masking, int64) (model.MaskEngine, bool, error){constant.NewMaskFromConfig, command.NewMaskFromConfig})
	t.Log(conf)
	t.Log(err)
	assert.NotEqual(t, err, nil, "Should not be nil")
}

func TestShouldCreateDictionaryFromJsonLine(t *testing.T) {
	jsonline := []byte("{\"name\":\"Benjamin\",\"age\":35}")
	dic, _ := JSONToDictionary(jsonline)
	waited := model.Dictionary{"name": "Benjamin", "age": float64(35)}

	assert.Equal(t, dic, waited, "Should create the right model.Dictionary")
}

func TestShouldIterateOverEntryIterator(t *testing.T) {
	tmpfile, _ := ioutil.TempFile("", "example")
	defer os.Remove(tmpfile.Name())
	_, errwrite := tmpfile.Write([]byte(`{"age":2, "name":"Toto"}`))
	if errwrite != nil {
		assert.Fail(t, "can't write")
	}

	_, errwrite = tmpfile.Write([]byte("\n"))
	if errwrite != nil {
		assert.Fail(t, "can't write")
	}

	_, errwrite = tmpfile.Write([]byte(`{"age":15, "name":"Benjamin"}`))
	if errwrite != nil {
		assert.Fail(t, "can't write")
	}

	_, errwrite = tmpfile.Seek(0, 0)
	if errwrite != nil {
		assert.Fail(t, "can't seek")
	}

	jsonlineIterator := NewJSONLineIterator(tmpfile)
	entry, err := jsonlineIterator.Next()
	if err != nil {
		assert.Fail(t, err.Error())
	}

	waited := model.Dictionary{"age": float64(2), "name": "Toto"}
	assert.Equal(t, entry, waited, "First read should be equal")

	entry, err = jsonlineIterator.Next()
	if err != nil {
		assert.Fail(t, "ERROR OPENING JSONLINEITERATOR")
	}

	waited = model.Dictionary{"age": float64(15), "name": "Benjamin"}
	assert.Equal(t, entry, waited, "Second read should be equal")

	_, err = jsonlineIterator.Next()
	assert.Equal(t, err, StopIteratorError{}, "Should return an end of iterator error")
}

func TestShouldCreateJsonLineFromDictionary(t *testing.T) {
	dic := model.Dictionary{"age": 35, "name": "Benjamin"}
	jsonline, _ := DictionaryToJSON(dic)
	waited := []byte("{\"age\":35,\"name\":\"Benjamin\"}\n")

	assert.Equal(t, jsonline, waited, "Should create the right JsonLine")
}

func TestNextShouldUnmarshalNestedJSON(t *testing.T) {
	jsonline := []byte(`{"personne":{"name": "Benjamin", "age" : 35}}`)
	jsonlineIterator := NewJSONLineIterator(bytes.NewReader(jsonline))
	dic, _ := jsonlineIterator.Next()
	waited := model.Dictionary{"personne": model.Dictionary{"name": "Benjamin", "age": float64(35)}}
	assert.Equal(t, dic, waited, "Should create a nested model.Dictionary")
}
