package main

import (
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var nameMasking = FunctionMaskEngine{func(name Entry) Entry { return "Toto" }}
var nameProgramMasking = CommandMaskEngine{"echo Toto"}
var nameList = []Entry{"Michel", "Marc", "Matthias", "Youen", "Alexis"}

func TestMaskingShouldReturnEmptyWhenInputISEmpty(t *testing.T) {
	maskingEngine := NewMaskConfiguration().AsEngine()
	data := Dictionary{}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be empty")
}

func TestMaskingShouldNoReplaceInsensitiveValue(t *testing.T) {
	maskingEngine := NewMaskConfiguration().AsEngine()

	data := Dictionary{"project": "ER456"}
	result := maskingEngine.Mask(data)

	assert.Equal(t, data, result, "should be equal")
}

func TestMaskingShouldReplaceSensitiveValue(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", nameMasking)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	assert.NotEqual(t, data, result, "should be masked")
}

func TestMaskingShouldReplaceSensitiveValueByCommand(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", nameProgramMasking)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	waited := Dictionary{"name": "Toto"}

	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReplaceSensitiveValueByRandomInList(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", NewMaskRandomList(nameList))

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)

	assert.NotEqual(t, data, result, "should be masked")

	namemap := result.(map[string]Entry)
	assert.Contains(t, nameList, namemap["name"], "Should be in the list")
}

func TestMaskingShouldReplaceSensitiveValueByRandomAndDifferent(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", NewMaskRandomList(nameList))
	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	diff := 0
	for i := 0; i < 1000; i++ {
		result := maskingEngine.Mask(data)
		resultBis := maskingEngine.Mask(data)
		if result.(map[string]Entry)["name"] != resultBis.(map[string]Entry)["name"] {
			diff++
		}
	}
	assert.True(t, diff >= 750, "Should be the same less than 250 times")
}

func TestMaskingShouldReplaceSensitiveValueByHashing(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("name", MaskHashList{nameList})

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Alexis"}
	result := maskingEngine.Mask(data)
	resultBis := maskingEngine.Mask(data)

	assert.Equal(t, result, resultBis, "Should be hashed the same way")
}

func TestMaskingShouldReplaceValueInNestedDictionary(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("customer", NewMaskConfiguration().
			WithEntry("name", nameMasking).AsEngine(),
		)

	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"customer": Dictionary{"name": "Benjamin"}, "project": "MyProject"}
	result := maskingEngine.Mask(data)
	want := Dictionary{"customer": Dictionary{"name": "Toto"}, "project": "MyProject"}
	assert.Equal(t, want, result, "should be masked")
}

func TestWithEntryShouldBuildNestedConfigurationWhenKeyContainsDot(t *testing.T) {
	config := NewMaskConfiguration().
		WithEntry("customer.name", nameMasking)
	_, ok := config.GetMaskingEngine("customer")
	assert.True(t, ok, "should be same configuration")

	_, okbis := config.GetMaskingEngine("customer.name")
	assert.False(t, okbis, "should be same configuration")
}

func TestMaskingShouldReplaceSensitiveValueByWeightedRandom(t *testing.T) {
	NameWeighted := []WeightedChoice{{data: "Michel", weight: 4}, {data: "Marc", weight: 1}}
	weightMask := NewWeightedMaskList(NameWeighted)
	config := NewMaskConfiguration().
		WithEntry("name", weightMask)

	maskingEngine := MaskingEngineFactory(config)

	wait := Dictionary{"name": "Michel"}
	equal := 0
	data := Dictionary{"name": "Benjamin"}
	for i := 0; i < 1000; i++ {
		result := maskingEngine.Mask(data).(map[string]Entry)
		if result["name"].(string) == wait["name"] {
			equal++
		}
	}

	assert.True(t, equal >= 750, "Should be more than 750 Michel")
	assert.True(t, equal <= 850, "Should be less than 150 Marc")
}

func TestMaskingShouldReplaceSensitiveValueByConstantValue(t *testing.T) {
	maskingConstant := "Toto"
	constMask := NewConstMask(maskingConstant)
	config := NewMaskConfiguration().
		WithEntry("name", constMask)
	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"name": "Benjamin"}
	result := maskingEngine.Mask(data)
	waited := Dictionary{"name": "Toto"}

	assert.NotEqual(t, data, result, "should be masked")
	assert.Equal(t, waited, result, "should be Toto")
}

func TestMaskingShouldReplaceSensitiveValueByRandomNumber(t *testing.T) {
	min := 7
	max := 77
	ageMask := RandomIntMask{min, max}
	config := NewMaskConfiguration().
		WithEntry("age", ageMask)

	maskingEngine := MaskingEngineFactory(config)
	data := Dictionary{"age": 83}
	result := maskingEngine.Mask(data).(map[string]Entry)

	assert.NotEqual(t, data, result, "Should be masked")
	assert.True(t, result["age"].(int) >= min, "Should be more than min")
	assert.True(t, result["age"].(int) <= max, "Should be less than max")
}

func TestMaskingShouldReplaceSensitiveValueByRegex(t *testing.T) {
	regex := "0[1-7]( ([0-9]){2}){4}"
	regmask := NewRegexMask(regex)
	config := NewMaskConfiguration().
		WithEntry("phone", regmask)
	maskingEngine := MaskingEngineFactory(config)

	data := Dictionary{"phone": "00 00 00 00 00"}
	result := maskingEngine.Mask(data).(map[string]Entry)
	match, _ := regexp.MatchString(regex, result["phone"].(string))
	assert.NotEqual(t, data, result, "should be masked")
	assert.True(t, match, "should match the regexp")
}

func TestMaskingShouldReplaceDateByRandom(t *testing.T) {
	dateMin := time.Date(1970, 1, 1, 0, 0, 0, 0, time.UTC)
	dateMax := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)
	dateRange := NewDateMask(dateMin, dateMax)

	config := NewMaskConfiguration().
		WithEntry("date", dateRange)

	maskingEngine := MaskingEngineFactory(config)
	data := Dictionary{"date": time.Date(2019, 3, 2, 0, 0, 0, 0, time.UTC)}

	result := maskingEngine.Mask(data).(map[string]Entry)

	assert.Greater(t, dateMax.Sub(result["date"].(time.Time)).Microseconds(), int64(0),
		"%v should be before max date %v", result["date"].(time.Time), dateMax)

	assert.Less(t, dateMin.Sub(result["date"].(time.Time)).Microseconds(), int64(0),
		"%v should be after min date %v", result["date"].(time.Time), dateMin)

	equal := 0
	for i := 0; i < 1000; i++ {
		result := maskingEngine.Mask(data).(map[string]Entry)
		if result["date"] == data["date"] {
			equal++
		}
	}

	assert.True(t, equal <= 1, "Shouldn't be the same date less than 0.1% of time")
}

func A(me MaskEngine, boo bool) MaskEngine {
	return me
}

func TestShouldCreateAMaskConfigurationFromAFile(t *testing.T) {
	filename := "masking.yml"
	config, _ := YamlConfig(filename)
	regex := "0[1-7]( ([0-9]){2}){4}"
	constMask := NewConstMask("Toto")
	regmask := NewRegexMask(regex)
	randList := NewMaskRandomListSeeded([]Entry{"Mickael", "Mathieu", "Marcelle"}, int64(42))
	nameWeighted := []WeightedChoice{{data: "Dupont", weight: 9}, {data: "Dupond", weight: 1}}
	waited := NewMaskConfiguration().
		WithEntry("customer.phone", regmask).
		WithEntry("name1", constMask).
		WithEntry("name2", randList).
		WithEntry("age", RandomIntMask{25, 32}).
		WithEntry("name3", CommandMaskEngine{"echo Dorothy"}).
		WithEntry("surname", NewWeightedMaskList(nameWeighted)).
		WithEntry("address.town", MaskHashList{[]Entry{"Emerald City", "Ruby City", "Sapphire City"}})

	assert.Equal(t, A(config.GetMaskingEngine("customer.phone")), A(waited.GetMaskingEngine("customer.phone")), "customer.phone not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name1")), A(waited.GetMaskingEngine("name1")), "name1 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name2")), A(waited.GetMaskingEngine("name2")), "name2 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("age")), A(waited.GetMaskingEngine("age")), "age not equal")
	assert.Equal(t, A(config.GetMaskingEngine("name3")), A(waited.GetMaskingEngine("name3")), "name3 not equal")
	assert.Equal(t, A(config.GetMaskingEngine("surname")), A(waited.GetMaskingEngine("surname")), "surname not equal")
	assert.Equal(t, A(config.GetMaskingEngine("address.town")), A(waited.GetMaskingEngine("address.town")), "surname not equal")
}

func TestShouldReturnAnErrorWithMultipleArguments(t *testing.T) {
	filename := "wrongMasking.yml"
	conf, err := YamlConfig(filename)
	t.Log(conf)
	t.Log(err)
	assert.NotEqual(t, err, nil, "Should not be nil")
}

func TestShouldCreateDictionaryFromJsonLine(t *testing.T) {
	jsonline := []byte("{\"name\":\"Benjamin\",\"age\":35}")
	dic, _ := JSONToDictionary(jsonline)
	waited := Dictionary{"name": "Benjamin", "age": float64(35)}

	assert.Equal(t, dic, waited, "Should create the right Dictionary")
}
