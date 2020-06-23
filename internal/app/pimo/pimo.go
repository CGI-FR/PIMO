package pimo

import (
	"bufio"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"time"

	"github.com/goccy/go-yaml"
	"makeit.imfr.cgi.com/makeit2/scm/lino/pimo/pkg/model"
)

// YAMLStructure of the file
type YAMLStructure struct {
	Version string          `yaml:"version"`
	Seed    int64           `yaml:"seed"`
	Masking []model.Masking `yaml:"masking"`
}

// YamlConfig create a MaskConfiguration from a file
func YamlConfig(filename string, factories []func(model.Masking, model.MaskConfiguration, int64) (model.MaskConfiguration, bool, error)) (model.MaskConfiguration, error) {
	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return model.NewMaskConfiguration(), err
	}
	var conf YAMLStructure
	err = yaml.Unmarshal(source, &conf)
	if err != nil {
		return model.NewMaskConfiguration(), err
	}
	if conf.Seed == 0 {
		conf.Seed = time.Now().UnixNano()
	}
	config := model.NewMaskConfiguration()
	for _, v := range conf.Masking {
		nbArg := 0
		for _, factory := range factories {
			newConfig, present, err := factory(v, config, conf.Seed)
			if err != nil {
				return nil, errors.New(err.Error() + " for " + v.Selector.Jsonpath)
			}
			if present {
				config = newConfig
				nbArg++
			}
		}
		if nbArg != 1 {
			return config, errors.New("Not the right number of argument for " + v.Selector.Jsonpath + ". There should be 1 and there is " + strconv.Itoa(nbArg))
		}
	}
	return config, nil
}

//InterfaceToDictionary returns a model.Dictionary from an interface
func InterfaceToDictionary(inter interface{}) model.Dictionary {
	dic := make(map[string]model.Entry)
	mapint := inter.(map[string]interface{})
	for k, v := range mapint {
		switch v.(type) {
		case map[string]interface{}:
			dic[k] = InterfaceToDictionary(v)
		default:
			dic[k] = v
		}
	}
	return dic
}

//JSONToDictionary return a model.Dictionary from a jsonline
func JSONToDictionary(jsonline []byte) (model.Dictionary, error) {
	var inter interface{}
	err := json.Unmarshal(jsonline, &inter)
	if err != nil {
		return nil, err
	}
	dic := InterfaceToDictionary(inter)
	return dic, nil
}

//StopIteratorError is an error for le JSONLineIterator
type StopIteratorError struct{}

func (e StopIteratorError) Error() string {
	return "Iterator couldn't find any value"
}

// JSONLineIterator export line to JSON format.
type JSONLineIterator struct {
	fscanner *bufio.Scanner
}

// NewJSONLineIterator creates a new JSONLineIterator.
func NewJSONLineIterator(file io.Reader) JSONLineIterator {
	return JSONLineIterator{bufio.NewScanner(file)}
}

// Next convert next line to model.Dictionary
func (jli *JSONLineIterator) Next() (model.Dictionary, error) {
	if !jli.fscanner.Scan() {
		return nil, StopIteratorError{}
	}
	line := jli.fscanner.Bytes()
	return JSONToDictionary(line)
}

//DictionaryToJSON create a jsonline from a model.Dictionary
func DictionaryToJSON(dic model.Dictionary) ([]byte, error) {
	jsonline, err := json.Marshal(dic)
	if err != nil {
		return nil, err
	}
	jsonline = append(jsonline, "\n"...)
	return jsonline, nil
}
