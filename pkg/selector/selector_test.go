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

package selector_test

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/cgi-fr/pimo/pkg/selector"
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

func getExampleAsInterface() interface{} {
	var inter interface{}
	err := json.Unmarshal(([]byte)(example), &inter)
	if err != nil {
		log.Fatal(err)
	}
	return inter
}

func getExampleAsDictionary() selector.Dictionary {
	return selector.InterfaceToDictionary(getExampleAsInterface())
}

func TestNotFound(t *testing.T) {
	sut := selector.NewSelector("will.not.be.found")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Fail(t, "this should not be called")
		return selector.NOTHING, nil
	})
	assert.False(t, found)
}

func TestReadSingle(t *testing.T) {
	sut := selector.NewSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "id", key)
		assert.Equal(t, "123456789", value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSingle(t *testing.T) {
	sut := selector.NewSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "123456789", value)
		return selector.WRITE, "0"
	})
	assert.True(t, found)
	assert.Equal(t, "0", dictionary["id"])
}

func TestDeleteSingle(t *testing.T) {
	sut := selector.NewSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "123456789", value)
		return selector.DELETE, nil
	})
	assert.True(t, found)
	assert.NotContains(t, dictionary, "id")
}

func TestReadSub(t *testing.T) {
	sut := selector.NewSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "name", key)
		assert.Equal(t, "test", value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSub(t *testing.T) {
	sut := selector.NewSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "test", value)
		return selector.WRITE, "write"
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "write", "tags": []selector.Entry{"red", "blue", "yellow"}}, dictionary["summary"])
}

func TestDeleteSub(t *testing.T) {
	sut := selector.NewSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "test", value)
		return selector.DELETE, nil
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "tags": []selector.Entry{"red", "blue", "yellow"}}, dictionary["summary"])
}

func TestReadSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "blue", "yellow"},
		}, parentContext)
		assert.Equal(t, "tags", key)
		collect = append(collect, value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{"red", "blue", "yellow"}, collect)
}

func TestWriteSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	cnt := 0
	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "blue", "yellow"},
		}, parentContext)
		collect = append(collect, value)
		cnt++
		return selector.WRITE, fmt.Sprintf("%v", cnt)
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{"red", "blue", "yellow"}, collect)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "test", "tags": []selector.Entry{"1", "2", "3"}}, dictionary["summary"])
}

func TestDeleteSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "blue", "yellow"},
		}, parentContext)
		collect = append(collect, value)
		if value == "blue" {
			return selector.DELETE, nil
		}
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{"red", "blue", "yellow"}, collect)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "test", "tags": []selector.Entry{"red", "yellow"}}, dictionary["summary"])
}

func TestReadComplexArray(t *testing.T) {
	sut := selector.NewSelector("organizations")

	dictionary := getExampleAsDictionary()

	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		collect = append(collect, value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, dictionary["organizations"], collect)
}

func TestWriteComplexArray(t *testing.T) {
	sut := selector.NewSelector("organizations")

	dictionary := getExampleAsDictionary()

	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		collect = append(collect, value)
		return selector.WRITE, selector.Dictionary{"domain": "masked", "persons": []selector.Entry{}}
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{selector.Dictionary{"domain": "masked", "persons": []selector.Entry{}}, selector.Dictionary{"domain": "masked", "persons": []selector.Entry{}}}, dictionary["organizations"])
}

func TestDeleteComplexArray(t *testing.T) {
	sut := selector.NewSelector("organizations")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		v := reflect.ValueOf(value)
		if v.MapIndex(reflect.ValueOf("domain")).Interface() == "company.com" {
			return selector.DELETE, nil
		}
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{selector.Dictionary{"domain": "company.fr", "persons": []selector.Entry{selector.Dictionary{"email": "", "name": "jean-baptiste", "surname": "renet"}, selector.Dictionary{"email": "", "name": "paul", "surname": "crouzeau"}}}}, dictionary["organizations"])
}

func TestReadComplexNestedArray(t *testing.T) {
	sut := selector.NewSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	collectParents := []selector.Entry{}
	collectValues := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		collectParents = append(collectParents, parentContext)
		collectValues = append(collectValues, value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	// company.com is seen twice, then company.fr is seen twice
	assert.Equal(t, []selector.Entry{selector.Dictionary{
		"domain": "company.com",
		"persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "leona", "surname": "miller"},
			selector.Dictionary{"email": "", "name": "joe", "surname": "davis"},
		},
	}, selector.Dictionary{
		"domain": "company.com",
		"persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "leona", "surname": "miller"},
			selector.Dictionary{"email": "", "name": "joe", "surname": "davis"},
		},
	}, selector.Dictionary{
		"domain": "company.fr",
		"persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "jean-baptiste", "surname": "renet"},
			selector.Dictionary{"email": "", "name": "paul", "surname": "crouzeau"},
		},
	}, selector.Dictionary{
		"domain": "company.fr",
		"persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "jean-baptiste", "surname": "renet"},
			selector.Dictionary{"email": "", "name": "paul", "surname": "crouzeau"},
		},
	}}, collectParents)
	// the four persons are seen in this order
	assert.Equal(t, []selector.Entry{
		selector.Dictionary{"email": "", "name": "leona", "surname": "miller"},
		selector.Dictionary{"email": "", "name": "joe", "surname": "davis"},
		selector.Dictionary{"email": "", "name": "jean-baptiste", "surname": "renet"},
		selector.Dictionary{"email": "", "name": "paul", "surname": "crouzeau"},
	}, collectValues)
}

func TestWriteComplexNestedArray(t *testing.T) {
	sut := selector.NewSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		return selector.WRITE, selector.Dictionary{"email": "", "name": "rick", "surname": "roll"}
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{
		selector.Dictionary{"domain": "company.com", "persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "rick", "surname": "roll"},
			selector.Dictionary{"email": "", "name": "rick", "surname": "roll"},
		}},
		selector.Dictionary{"domain": "company.fr", "persons": []selector.Entry{
			selector.Dictionary{"email": "", "name": "rick", "surname": "roll"},
			selector.Dictionary{"email": "", "name": "rick", "surname": "roll"},
		}}}, dictionary["organizations"])
}

func TestReadContextSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	found := sut.ApplyContext(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, parentContext, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "blue", "yellow"},
		})
		assert.Equal(t, []selector.Entry{"red", "blue", "yellow"}, value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "test", "tags": []selector.Entry{"red", "blue", "yellow"}}, dictionary["summary"])
}

func TestWriteContextSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	found := sut.ApplyContext(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, parentContext, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "blue", "yellow"},
		})
		assert.Equal(t, []selector.Entry{"red", "blue", "yellow"}, value)
		return selector.WRITE, []selector.Entry{"pink", "cyan", "magenta"}
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "test", "tags": []selector.Entry{"pink", "cyan", "magenta"}}, dictionary["summary"])
}

func TestDeleteComplexNestedArray(t *testing.T) {
	sut := selector.NewSelector("organizations.persons")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, key string, value selector.Entry) (selector.Action, selector.Entry) {
		v := reflect.ValueOf(value)
		if v.MapIndex(reflect.ValueOf("name")).Interface() == "leona" {
			return selector.DELETE, nil
		}
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{
		selector.Dictionary{
			"domain": "company.com",
			"persons": []selector.Entry{
				// leona is missing
				selector.Dictionary{"email": "", "name": "joe", "surname": "davis"},
			},
		},
		selector.Dictionary{
			"domain": "company.fr",
			"persons": []selector.Entry{
				selector.Dictionary{"email": "", "name": "jean-baptiste", "surname": "renet"},
				selector.Dictionary{"email": "", "name": "paul", "surname": "crouzeau"},
			},
		}}, dictionary["organizations"])
}

// Tests to copy the selector V1 API

/* func TestReadV1(t *testing.T) {
	sut := selector.NewSelector("elements.persons.phonenumber")
	dictionary := selector.Dictionary{"elements": []selector.Entry{selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": "001"}, selector.Dictionary{"phonenumber": "002"}}}, selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": "003"}, selector.Dictionary{"phonenumber": "004"}}}}}

	entry, found := sut.Read(dictionary)

	expected := []selector.Entry{[]selector.Entry{"001", "002"}, []selector.Entry{"003", "004"}}

	assert.Equal(t, true, found)
	assert.Equal(t, expected, entry)
}

func TestWriteV1(t *testing.T) {
	sut := selector.NewSelector("elements.persons.phonenumber")
	dictionary := selector.Dictionary{"elements": []selector.Entry{selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": "001"}, selector.Dictionary{"phonenumber": "002"}}}, selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": "003"}, selector.Dictionary{"phonenumber": "004"}}}}}
	entry := []selector.Entry{[]selector.Entry{1, 2}, []selector.Entry{3, 4}}

	result := sut.Write(dictionary, entry)

	expected := selector.Dictionary{"elements": []selector.Entry{selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": 1}, selector.Dictionary{"phonenumber": 2}}}, selector.Dictionary{"persons": []selector.Entry{selector.Dictionary{"phonenumber": 3}, selector.Dictionary{"phonenumber": 4}}}}}

	assert.Equal(t, expected, result)
} */
