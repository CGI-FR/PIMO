package selector_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"reflect"
	"strings"
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

func getExample() string {
	cmd := exec.Command("jq", "-cMS", ".")
	cmd.Stdin = strings.NewReader(example)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	var stderr bytes.Buffer
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if stderr.Len() > 0 {
		log.Fatal(stderr.String())
	}
	return stdout.String()
}

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

func toMap(inter interface{}) map[string]interface{} {
	dic := make(map[string]interface{})
	mapint := inter.(map[string]interface{})

	for k, v := range mapint {
		switch typedValue := v.(type) {
		case map[string]interface{}:
			dic[k] = toMap(v)
		case []interface{}:
			tab := []interface{}{}
			for _, item := range typedValue {
				_, dico := item.(map[string]interface{})

				if dico {
					tab = append(tab, toMap(item))
				} else {
					tab = append(tab, item)
				}
			}
			dic[k] = tab
		default:
			dic[k] = v
		}
	}
	return dic
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

func TestWriteContext(t *testing.T) {
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
