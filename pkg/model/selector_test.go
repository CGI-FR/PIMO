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

package model_test

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"github.com/cgi-fr/pimo/pkg/model"
	"github.com/stretchr/testify/assert"
)

// nolint: gochecknoglobals
var example = `
{
    "id": "123456789",
	"summary":
		{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": ["red","blue","yellow"]
		},
    "organizations": [
        {
            "domain": "company.com",
            "persons": [
                {
                    "name": "leona",
                    "surname": "miller",
                    "email": ""
                },
                {
                    "name": "joe",
                    "surname": "davis",
                    "email": ""
                }
            ]
        },
        {
            "domain": "company.fr",
            "persons": [
                {
                    "name": "jean-baptiste",
                    "surname": "renet",
                    "email": ""
                },
                {
                    "name": "paul",
                    "surname": "crouzeau",
                    "email": ""
                }
            ]
        }
    ]
}
`

func getExampleAsDictionary() model.Dictionary {
	dict := model.NewDictionary()
	err := json.Unmarshal(([]byte)(example), &dict)
	if err != nil {
		log.Fatal(err)
	}
	return model.CleanDictionary(dict)
}

func TestNotFound(t *testing.T) {
	sut := model.NewPathSelector("will.not.be.found")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Fail(t, "this should not be called")
		return model.NOTHING, nil
	})
	assert.False(t, found)
}

func TestReadSingle(t *testing.T) {
	sut := model.NewPathSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "id", key)
		assert.Equal(t, "123456789", value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSingle(t *testing.T) {
	sut := model.NewPathSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "123456789", value)
		return model.WRITE, "0"
	})
	assert.True(t, found)
	assert.Equal(t, "0", dictionary.Get("id"))
}

func TestDeleteSingle(t *testing.T) {
	sut := model.NewPathSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "123456789", value)
		return model.DELETE, nil
	})
	assert.True(t, found)
	_, exist := dictionary.OrderedMap.Get("id")
	assert.False(t, exist)
}

func TestReadSub(t *testing.T) {
	sut := model.NewPathSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary.Get("summary"), parentContext)
		assert.Equal(t, "name", key)
		assert.Equal(t, "test", value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSub(t *testing.T) {
	sut := model.NewPathSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary.Get("summary"), parentContext)
		assert.Equal(t, "test", value)
		return model.WRITE, "write"
	})
	assert.True(t, found)
	assert.Equal(t,
		model.NewDictionary().
			With("name", "write").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
		dictionary.Get("summary"))
}

func TestDeleteSub(t *testing.T) {
	sut := model.NewPathSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary.Get("summary"), parentContext)
		assert.Equal(t, "test", value)
		return model.DELETE, nil
	})
	assert.True(t, found)
	assert.Equal(t,
		model.NewDictionary().
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
		dictionary.Get("summary"))
}

func TestReadSimpleArray(t *testing.T) {
	sut := model.NewPathSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	collect := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
			parentContext)
		assert.Equal(t, "tags", key)
		collect = append(collect, value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []model.Entry{"red", "blue", "yellow"}, collect)
}

func TestWriteSimpleArray(t *testing.T) {
	sut := model.NewPathSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	cnt := 0
	collect := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
			parentContext)
		collect = append(collect, value)
		cnt++
		return model.WRITE, fmt.Sprintf("%v", cnt)
	})
	assert.True(t, found)
	assert.Equal(t, []model.Entry{"red", "blue", "yellow"}, collect)
	assert.Equal(t,
		model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"1", "2", "3"}),
		dictionary.Get("summary"))
}

func TestDeleteSimpleArray(t *testing.T) {
	sut := model.NewPathSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	collect := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
			parentContext)
		collect = append(collect, value)
		if value == "blue" {
			return model.DELETE, nil
		}
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []model.Entry{"red", "blue", "yellow"}, collect)
	assert.Equal(t,
		model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "yellow"}),
		dictionary.Get("summary"))
}

func TestReadComplexArray(t *testing.T) {
	sut := model.NewPathSelector("organizations")

	dictionary := getExampleAsDictionary()

	collect := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		collect = append(collect, value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, dictionary.Get("organizations"), collect)
}

func TestWriteComplexArray(t *testing.T) {
	sut := model.NewPathSelector("organizations")

	dictionary := getExampleAsDictionary()

	collect := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		collect = append(collect, value)
		return model.WRITE, model.NewDictionary().
			With("domain", "masked").
			With("persons", []model.Entry{})
	})
	assert.True(t, found)
	assert.Equal(t,
		[]model.Entry{
			model.NewDictionary().
				With("domain", "masked").
				With("persons", []model.Entry{}),
			model.NewDictionary().
				With("domain", "masked").
				With("persons", []model.Entry{}),
		}, dictionary.Get("organizations"))
}

func TestDeleteComplexArray(t *testing.T) {
	sut := model.NewPathSelector("organizations")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		dict := value.(model.Dictionary)
		if dict.Get("domain") == "company.com" {
			return model.DELETE, nil
		}
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []model.Entry{model.NewDictionary().
		With("domain", "company.fr").
		With("persons", []model.Entry{
			model.NewDictionary().
				With("name", "jean-baptiste").
				With("surname", "renet").
				With("email", ""),
			model.NewDictionary().
				With("name", "paul").
				With("surname", "crouzeau").
				With("email", ""),
		})}, dictionary.Get("organizations"))
}

func TestReadComplexNestedArray(t *testing.T) {
	sut := model.NewPathSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	collectParents := []model.Entry{}
	collectValues := []model.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		collectParents = append(collectParents, parentContext)
		collectValues = append(collectValues, value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
	// company.com is seen twice, then company.fr is seen twice
	assert.Equal(t, []model.Entry{
		model.NewDictionary().
			With("domain", "company.com").
			With("persons", []model.Entry{
				model.NewDictionary().With("name", "leona").With("surname", "miller").With("email", ""),
				model.NewDictionary().With("name", "joe").With("surname", "davis").With("email", ""),
			}),
		model.NewDictionary().
			With("domain", "company.com").
			With("persons", []model.Entry{
				model.NewDictionary().With("name", "leona").With("surname", "miller").With("email", ""),
				model.NewDictionary().With("name", "joe").With("surname", "davis").With("email", ""),
			}),
		model.NewDictionary().
			With("domain", "company.fr").
			With("persons", []model.Entry{
				model.NewDictionary().With("name", "jean-baptiste").With("surname", "renet").With("email", ""),
				model.NewDictionary().With("name", "paul").With("surname", "crouzeau").With("email", ""),
			}),
		model.NewDictionary().
			With("domain", "company.fr").
			With("persons", []model.Entry{
				model.NewDictionary().With("name", "jean-baptiste").With("surname", "renet").With("email", ""),
				model.NewDictionary().With("name", "paul").With("surname", "crouzeau").With("email", ""),
			}),
	}, collectParents)
	// the four persons are seen in this order
	assert.Equal(t, []model.Entry{
		model.NewDictionary().With("name", "leona").With("surname", "miller").With("email", ""),
		model.NewDictionary().With("name", "joe").With("surname", "davis").With("email", ""),
		model.NewDictionary().With("name", "jean-baptiste").With("surname", "renet").With("email", ""),
		model.NewDictionary().With("name", "paul").With("surname", "crouzeau").With("email", ""),
	}, collectValues)
}

func TestWriteComplexNestedArray(t *testing.T) {
	sut := model.NewPathSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		return model.WRITE, model.NewDictionary().With("email", "").With("name", "rick").With("surname", "roll")
	})
	assert.True(t, found)
	assert.Equal(t,
		[]model.Entry{
			model.NewDictionary().
				With("domain", "company.com").
				With("persons", []model.Entry{
					model.NewDictionary().With("email", "").With("name", "rick").With("surname", "roll"),
					model.NewDictionary().With("email", "").With("name", "rick").With("surname", "roll"),
				}),
			model.NewDictionary().
				With("domain", "company.fr").
				With("persons", []model.Entry{
					model.NewDictionary().With("email", "").With("name", "rick").With("surname", "roll"),
					model.NewDictionary().With("email", "").With("name", "rick").With("surname", "roll"),
				}),
		},
		dictionary.Get("organizations"))
}

func TestReadContextSimpleArray(t *testing.T) {
	sut := model.NewPathSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	found := sut.ApplyContext(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, parentContext, model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
		)
		assert.Equal(t, []model.Entry{"red", "blue", "yellow"}, value)
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, model.NewDictionary().
		With("name", "test").
		With("date", "2012-04-23T18:25:43.511Z").
		With("tags", []model.Entry{"red", "blue", "yellow"}), dictionary.Get("summary"))
}

func TestWriteContextSimpleArray(t *testing.T) {
	sut := model.NewPathSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	found := sut.ApplyContext(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, parentContext, model.NewDictionary().
			With("name", "test").
			With("date", "2012-04-23T18:25:43.511Z").
			With("tags", []model.Entry{"red", "blue", "yellow"}),
		)
		assert.Equal(t, []model.Entry{"red", "blue", "yellow"}, value)
		return model.WRITE, []model.Entry{"pink", "cyan", "magenta"}
	})
	assert.True(t, found)
	assert.Equal(t, model.NewDictionary().
		With("name", "test").
		With("date", "2012-04-23T18:25:43.511Z").
		With("tags", []model.Entry{"pink", "cyan", "magenta"}), dictionary.Get("summary"))
}

func TestDeleteComplexNestedArray(t *testing.T) {
	sut := model.NewPathSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext model.Dictionary, key string, value model.Entry) (model.Action, model.Entry) {
		dict := value.(model.Dictionary)
		if dict.Get("name") == "leona" {
			return model.DELETE, nil
		}
		return model.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []model.Entry{
		model.NewDictionary().
			With("domain", "company.com").
			With("persons", []model.Entry{
				// leona is missing
				model.NewDictionary().With("name", "joe").With("surname", "davis").With("email", ""),
			},
			),
		model.NewDictionary().
			With("domain", "company.fr").
			With("persons", []model.Entry{
				model.NewDictionary().With("name", "jean-baptiste").With("surname", "renet").With("email", ""),
				model.NewDictionary().With("name", "paul").With("surname", "crouzeau").With("email", ""),
			},
			),
	}, dictionary.Get("organizations"))
}

// Tests to copy the selector V1 API

/* func TestReadV1(t *testing.T) {
	sut := model.NewPathSelector("elements.persons.phonenumber")
	dictionary := model.Dictionary{"elements": []model.Entry{model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": "001"}, model.Dictionary{"phonenumber": "002"}}}, model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": "003"}, model.Dictionary{"phonenumber": "004"}}}}}

	entry, found := sut.Read(dictionary)

	expected := []model.Entry{[]model.Entry{"001", "002"}, []model.Entry{"003", "004"}}

	assert.Equal(t, true, found)
	assert.Equal(t, expected, entry)
}

func TestWriteV1(t *testing.T) {
	sut := model.NewPathSelector("elements.persons.phonenumber")
	dictionary := model.Dictionary{"elements": []model.Entry{model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": "001"}, model.Dictionary{"phonenumber": "002"}}}, model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": "003"}, model.Dictionary{"phonenumber": "004"}}}}}
	entry := []model.Entry{[]model.Entry{1, 2}, []model.Entry{3, 4}}

	result := sut.Write(dictionary, entry)

	expected := model.Dictionary{"elements": []model.Entry{model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": 1}, model.Dictionary{"phonenumber": 2}}}, model.Dictionary{"persons": []model.Entry{model.Dictionary{"phonenumber": 3}, model.Dictionary{"phonenumber": 4}}}}}

	assert.Equal(t, expected, result)
} */
