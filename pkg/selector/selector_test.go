package selector_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
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
			"tags": ["red", "yellow"]
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
                    "name": "jean-baptise",
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

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Fail(t, "this should not be called")
		return selector.NOTHING, nil
	})
	assert.False(t, found)
}

func TestReadSingle(t *testing.T) {
	sut := selector.NewSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary, parentContext)
		assert.Equal(t, "123456789", value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSingle(t *testing.T) {
	sut := selector.NewSelector("id")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
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

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
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

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "test", value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
}

func TestWriteSub(t *testing.T) {
	sut := selector.NewSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "test", value)
		return selector.WRITE, "write"
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "name": "write", "tags": []selector.Entry{"red", "yellow"}}, dictionary["summary"])
}

func TestDeleteSub(t *testing.T) {
	sut := selector.NewSelector("summary.name")

	dictionary := getExampleAsDictionary()

	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, dictionary["summary"], parentContext)
		assert.Equal(t, "test", value)
		return selector.DELETE, nil
	})
	assert.True(t, found)
	assert.Equal(t, selector.Dictionary{"date": "2012-04-23T18:25:43.511Z", "tags": []selector.Entry{"red", "yellow"}}, dictionary["summary"])
}

func TestReadSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "yellow"},
		}, parentContext)
		collect = append(collect, value)
		return selector.NOTHING, nil
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{"red", "yellow"}, collect)
}

func TestWriteSimpleArray(t *testing.T) {
	sut := selector.NewSelector("summary.tags")

	dictionary := getExampleAsDictionary()

	cnt := 0
	collect := []selector.Entry{}
	found := sut.Apply(dictionary, func(rootContext, parentContext selector.Dictionary, value selector.Entry) (selector.Action, selector.Entry) {
		assert.Equal(t, dictionary, rootContext)
		assert.Equal(t, selector.Dictionary{
			"name": "test",
			"date": "2012-04-23T18:25:43.511Z",
			"tags": []selector.Entry{"red", "yellow"},
		}, parentContext)
		collect = append(collect, value)
		cnt++
		return selector.WRITE, fmt.Sprintf("%v", cnt)
	})
	assert.True(t, found)
	assert.Equal(t, []selector.Entry{"red", "yellow"}, collect)
}
