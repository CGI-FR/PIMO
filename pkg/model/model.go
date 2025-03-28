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

package model

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"iter"
	tmpl "text/template"
	"time"

	"github.com/cgi-fr/pimo/pkg/statistics"
	"github.com/cgi-fr/pimo/pkg/template"
	"github.com/rs/zerolog/log"
)

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry, ...Dictionary) (Entry, error)
}

// MaskContextEngine is a masking algorithm for dictionary
type MaskContextEngine interface {
	MaskContext(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// HasCleaner interface provides a function to apply on cleanup
type HasCleaner interface {
	GetCleaner() FunctionMaskContextEngine
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	Function func(Entry, ...Dictionary) (Entry, error)
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry, context ...Dictionary) (Entry, error) {
	return fme.Function(e, context...)
}

// FunctionMaskContextEngine implements MaskContextEngine with a simple function
type FunctionMaskContextEngine struct {
	Function func(Dictionary, string, ...Dictionary) (Dictionary, error)
}

// MaskContext delegate mask algorithm to the function
func (fme FunctionMaskContextEngine) MaskContext(e Dictionary, key string, context ...Dictionary) (Dictionary, error) {
	return fme.Function(e, key, context...)
}

type MaskFactoryConfiguration struct {
	Masking   Masking
	Seed      int64
	Cache     map[string]Cache
	Functions tmpl.FuncMap
}

type MaskFactory func(MaskFactoryConfiguration) (MaskEngine, bool, error)

type MaskContextFactory func(MaskFactoryConfiguration) (MaskContextEngine, bool, error)

type SelectorType struct {
	Jsonpath string `yaml:"jsonpath" json:"jsonpath" jsonschema_description:"Path of the target value to mask in the JSON input"`
}

type IncrementalType struct {
	Start     int `yaml:"start" json:"start" jsonschema_description:"First value in the sequence"`
	Increment int `yaml:"increment" json:"increment" jsonschema_description:"Increment to add to reach the next value in the sequence"`
}

type RandDateType struct {
	DateMin time.Time `yaml:"dateMin" json:"dateMin" jsonschema_description:"Lower bound of the date range"`
	DateMax time.Time `yaml:"dateMax" json:"dateMax" jsonschema_description:"Higher bound of the date range"`
}

type RandIntType struct {
	Min int `yaml:"min" json:"min" jsonschema_description:"Lower bound of the integer range"`
	Max int `yaml:"max" json:"max" jsonschema_description:"Lower bound of the integer range"`
}

type WeightedChoiceType struct {
	Choice Entry `yaml:"choice" json:"choice" jsonschema_description:"Value for this choice"`
	Weight uint  `yaml:"weight" json:"weight" jsonschema_description:"Weight of this choice, higher weights will be selected more frequently"`
}

type RandomDurationType struct {
	Min string `yaml:"min" json:"min" jsonschema_description:"Lower bound of the duration range (ISO 8601 notation)"`
	Max string `yaml:"max" json:"max" jsonschema_description:"Higher bound of the duration range (ISO 8601 notation)"`
}

type RandomDecimalType struct {
	Min       float64 `yaml:"min" json:"min" jsonschema_description:"Lower bound of the decimal range"`
	Max       float64 `yaml:"max" json:"max" jsonschema_description:"Lower bound of the decimal range"`
	Precision int     `yaml:"precision" json:"precision" jsonschema_description:"Precision of the generated value"`
}

type DateParserType struct {
	InputFormat  string `yaml:"inputFormat,omitempty" json:"inputFormat,omitempty" jsonschema_description:"Format of the input datetime, it should always display the following date : Mon Jan 2 15:04:05 -0700 MST 2006 or the constant value 'unixEpoch'"`
	OutputFormat string `yaml:"outputFormat,omitempty" json:"outputFormat,omitempty" jsonschema_description:"Format of the output datetime, it should always display the following date : Mon Jan 2 15:04:05 -0700 MST 2006 or the constant value 'unixEpoch'"`
}

type FF1Type struct {
	KeyFromEnv string  `yaml:"keyFromEnv" json:"keyFromEnv" jsonschema_description:"Name of the system environment variable that contains the private key"`
	TweakField string  `yaml:"tweakField,omitempty" json:"tweakField,omitempty" jsonschema_description:"Name of the field to use as 'tweak' value : reduce the attack surface by using a varying value on each record, it can be considered as an extension of the secret key that change on each record"`
	Radix      uint    `yaml:"radix,omitempty" json:"radix,omitempty" jsonschema_description:"determine which part of the fixed FF1 domain definition will actually be used, for example 10 will use the first 10 characters of 0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"`
	Domain     string  `yaml:"domain,omitempty" json:"domain,omitempty" jsonschema_description:"allowed characters domain that will be encrypted, do not use radix and domain at the same time"`
	Preserve   string  `yaml:"preserve,omitempty" json:"preserve,omitempty" jsonschema_description:"preserve characters that are not part of the allowed domain"`
	OnError    *string `yaml:"onError,omitempty" json:"onError,omitempty" jsonschema_description:"template to execute if there is an error while encrypting value"`
	Decrypt    bool    `yaml:"decrypt,omitempty" json:"decrypt,omitempty" jsonschema_description:"Decrypt the value instead of encrypt"`
}

type PipeType struct {
	Masking        []Masking `yaml:"masking,omitempty" json:"masking,omitempty" jsonschema_description:"Define a list of selector/mask couple to apply on the JSON stream, in this order"`
	InjectParent   string    `yaml:"injectParent,omitempty" json:"injectParent,omitempty" jsonschema_description:"Used in conjunction with the 'pipe' mask, inject the parent object with the given field name"`
	InjectRoot     string    `yaml:"injectRoot,omitempty" json:"injectRoot,omitempty" jsonschema_description:"Used in conjunction with the 'pipe' mask, inject the root object with the given field name"`
	DefinitionFile string    `yaml:"file,omitempty" json:"file,omitempty" jsonschema_description:"URI to an external resource to read the pipeline definition"`
}

type TemplateEachType struct {
	Item     string `yaml:"item,omitempty" json:"item,omitempty" jsonschema_description:"Inject the current element value under the given field name"`
	Index    string `yaml:"index,omitempty" json:"index,omitempty" jsonschema_description:"Inject the current element index under the given field name"`
	Template string `yaml:"template,omitempty" json:"template,omitempty" jsonschema_description:"Replace the current value with the result of executing this Golang template"`
}

type LuhnType struct {
	Universe string `yaml:"universe,omitempty" json:"universe,omitempty" jsonschema_description:"All possible characters that can be encountered as input value"`
}

type MarkovType struct {
	MaxSize   int    `yaml:"max-size,omitempty" json:"max-size,omitempty" jsonschema_description:"Maximum length for the generated text"`
	Sample    string `yaml:"sample,omitempty" json:"sample,omitempty" jsonschema_description:"URI to an external resource to train the Markiv model"`
	Separator string `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator to use to read tokens, leave empty to treat each character as a token"`
	Order     int    `yaml:"order,omitempty" json:"order,omitempty" jsonschema_description:"Number of tokens to consider, a higher value = more similar to sample, too high and the generated text will be completly similar to sample"`
}

type Class struct {
	Input  string `yaml:"input" json:"input" jsonschema_description:"Characters to replace in the input value"`
	Output string `yaml:"output" json:"output" jsonschema_description:"Characters to use to generate the output value"`
}

type TranscodeType struct {
	Classes []Class `yaml:"classes,omitempty" json:"classes,omitempty" jsonschema_description:"Each class will define a rule to replace a set of characters by another"`
}

type EmbeddedSha3Type struct {
	Sha3Type `yaml:",inline"`
	Field    string `yaml:"field" json:"field" jsonschema_description:"Name of the identifier"`
}

type ChoiceInCSVType struct {
	URI             string           `yaml:"uri" json:"uri" jsonschema_description:"URI of the CSV resource"`
	Header          bool             `yaml:"header,omitempty" json:"header,omitempty" jsonschema_description:"Does the CSV resource contains a header line"`
	Separator       string           `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator character"`
	Comment         string           `yaml:"comment,omitempty" json:"comment,omitempty" jsonschema_description:"Lines beginning with the comment character without preceding whitespace are ignored"`
	FieldsPerRecord int              `yaml:"fieldsPerRecord,omitempty" json:"fieldsPerRecord,omitempty" jsonschema_description:"FieldsPerRecord is the number of expected fields per record, 0 means the number of fields in the first record"`
	TrimSpace       bool             `yaml:"trim,omitempty" json:"trim,omitempty" jsonschema_description:"If true leading white space in a field is ignored"`
	Identifier      EmbeddedSha3Type `yaml:"identifier,omitempty" json:"identifier,omitempty" jsonschema_description:"Configure a collision resistant ID generator"`
}

type ExactMatchType struct {
	CSV   string `yaml:"csv" json:"csv" jsonschema_description:"Characters exact to match in the csv file"`
	Entry string `yaml:"entry" json:"entry" jsonschema_description:"Characters exact to find in the entry file"`
}

type FindInCSVType struct {
	URI             string         `yaml:"uri" json:"uri" jsonschema_description:"URI of the CSV resource"`
	ExactMatch      ExactMatchType `yaml:"exactMatch,omitempty" json:"exactMatch,omitempty" jsonschema_description:"Compare csv data and entry data, find the exected matched csv line"`
	JaccardMatch    ExactMatchType `yaml:"jaccard,omitempty" json:"jaccard,omitempty" jsonschema_description:"Compare csv data and entry data with jaccard, find the similarity matched csv line"`
	Expected        string         `yaml:"expected" json:"expected" jsonschema_description:"How much result return, 3 modes availables: only-one, at-least-one, many"`
	Header          bool           `yaml:"header,omitempty" json:"header,omitempty" jsonschema_description:"Does the CSV resource contains a header line"`
	Separator       string         `yaml:"separator,omitempty" json:"separator,omitempty" jsonschema_description:"Separator character"`
	Comment         string         `yaml:"comment,omitempty" json:"comment,omitempty" jsonschema_description:"Lines beginning with the comment character without preceding whitespace are ignored"`
	FieldsPerRecord int            `yaml:"fieldsPerRecord,omitempty" json:"fieldsPerRecord,omitempty" jsonschema_description:"FieldsPerRecord is the number of expected fields per record, 0 means the number of fields in the first record"`
	TrimSpace       bool           `yaml:"trim,omitempty" json:"trim,omitempty" jsonschema_description:"If true leading white space in a field is ignored"`
}

type XMLType struct {
	XPath        string    `yaml:"xpath" json:"xpath" jsonschema_description:"Specifies the XPath expression to target the XML"`
	InjectParent string    `yaml:"injectParent,omitempty" json:"injectParent,omitempty" jsonschema_description:"Injects the parent element into a variable for access in subsequent masks"`
	Masking      []Masking `yaml:"masking,omitempty" json:"masking,omitempty" jsonschema_description:"Sub-list of masks to be applied inside the selected XML element"`
}

type TimeLineConstraintType struct {
	Before  string `yaml:"before,omitempty" json:"before,omitempty" jsonschema:"oneof_required=Before,title=Before,description=Name the point which should serve as the upper limit"`
	After   string `yaml:"after,omitempty" json:"after,omitempty" jsonschema:"oneof_required=After,title=After,description=Name the point which should serve as the lower limit"`
	OnError string `yaml:"onError,omitempty" json:"onError,omitempty" jsonschema:"enum=default,enum=reject" jsonschema_description:"What to do if there is an error : default = use point default value (or null if not set), reject = fail masking for this line, default = use default value for the point"`
	Epsilon string `yaml:"epsilon,omitempty" json:"epsilon,omitempty" jsonschema_description:"Minimum period to consider 2 dates unequals for this constraint"`
}

type TimeLinePointType struct {
	Name        string                   `yaml:"name" json:"name" jsonschema_description:"Name of the point in the timeline"`
	From        string                   `yaml:"from,omitempty" json:"from,omitempty" jsonschema_description:"Name of the reference point in the timeline to create this point"`
	Min         string                   `yaml:"min" json:"min" jsonschema_description:"Minimum shift relative to the reference point (ISO 8601 notation)"`
	Max         string                   `yaml:"max" json:"max" jsonschema_description:"Maximum shift relative to the reference point (ISO 8601 notation)"`
	Constraints []TimeLineConstraintType `yaml:"constraints,omitempty" json:"constraints,omitempty" jsonschema_description:"List of constraints to fulfill"`
	Default     string                   `yaml:"default,omitempty" json:"default,omitempty" jsonschema_description:"Name a point of the timeline to use as a default value if a constraint fail (with onError=default)"`
}

type TimeLineStartType struct {
	Name  string `yaml:"name" json:"name" jsonschema_description:"Name of the starting point in the timeline"`
	Value string `yaml:"value,omitempty" json:"value,omitempty" jsonschema_description:"Value of the starting point in the timeline, in RFC3339 format (2006-01-02T15:04:05Z07:00), if omitted equals to now"`
}

type TimeLineType struct {
	Start    TimeLineStartType   `yaml:"start" json:"start" jsonschema_description:"Origin of the timeline"`
	Format   string              `yaml:"format,omitempty" json:"format,omitempty" jsonschema_description:"Format of datetimes, it should always display the following date : Mon Jan 2 15:04:05 -0700 MST 2006 or the constant value 'unixEpoch'"`
	Points   []TimeLinePointType `yaml:"points" json:"points" jsonschema_description:"List of points in the timeline"`
	MaxRetry *int                `yaml:"retry,omitempty" json:"retry,omitempty" jsonschema_description:"Maximum number of retry if constraint fail before error (default : 200)"`
	Epsilon  string              `yaml:"epsilon,omitempty" json:"epsilon,omitempty" jsonschema_description:"Minimum period to consider 2 dates unequals (default : P0)"`
}

type SequenceType struct {
	Format  string `yaml:"format" json:"format" jsonschema_description:"Format of the sequenced ID"`
	Varying string `yaml:"varying,omitempty" json:"varying,omitempty" jsonschema_description:"List of varying characters in the sequence (default : 0123456789)"`
}

type Sha3Type struct {
	Length     int    `yaml:"length,omitempty" json:"length,omitempty" jsonschema:"oneof_required=Length,title=Length,description=Length of the produced output in bytes"`
	Resistance int    `yaml:"resistance,omitempty" json:"resistance,omitempty" jsonschema:"oneof_required=Resistance,title=Resistance,description=Collision resistance of the produced hash"`
	Domain     string `yaml:"domain,omitempty" json:"domain,omitempty" jsonschema_description:"allowed characters domain in the output, default to hexadecimal (0123456789abcdef)"`
	MaxStrLen  int    `yaml:"maxstrlen,omitempty" json:"maxstrlen,omitempty" jsonschema_description:"an error will occur if the identifier can grow longer than the specified length"`
}

type ApplyType struct {
	URI string `yaml:"uri" json:"uri" jsonschema_description:"URI of the mask resource"`
}

type PartitionType struct {
	Name string     `yaml:"name" json:"name" jsonschema_description:"name of the partition"`
	When string     `yaml:"when,omitempty" json:"when,omitempty" jsonschema_description:"template to execute, if true the condition is active"`
	Then []MaskType `yaml:"then,omitempty" json:"then,omitempty" jsonschema_description:"list of masks to execute if the condition is active"`
}

type SegmentType struct {
	Regex   string                `yaml:"regex" json:"regex" jsonschema_description:"regex used to create segments using group captures, groups must be named"`
	Match   map[string][]MaskType `yaml:"match" json:"match" jsonschema_description:"list of masks to execute for each group if the regex matched"`
	NoMatch []MaskType            `yaml:"nomatch,omitempty" json:"nomatch,omitempty" jsonschema_description:"list of masks to execute for each group if the regex did not match"`
}

type LogType struct {
	Message string `yaml:"message,omitempty" json:"message,omitempty" jsonschema_description:"log message (this field is a template), if not set log the current selected value"`
	Level   string `yaml:"level,omitempty" json:"level,omitempty" jsonschema:"enum=trace,enum=debug,enum=info,enum=warn,enum=error" jsonschema_description:"log level, default to info"`
}

type MaskType struct {
	Add               Entry                `yaml:"add,omitempty" json:"add,omitempty" jsonschema:"oneof_required=Add,title=Add Mask,description=Add a new field in the JSON stream"`
	AddTransient      Entry                `yaml:"add-transient,omitempty" json:"add-transient,omitempty" jsonschema:"oneof_required=AddTransient,title=Add Transient Mask" jsonschema_description:"Add a new temporary field, that will not show in the JSON output"`
	Constant          Entry                `yaml:"constant,omitempty" json:"constant,omitempty" jsonschema:"oneof_required=Constant,title=Constant Mask" jsonschema_description:"Replace the input value with a constant field"`
	RandomChoice      []Entry              `yaml:"randomChoice,omitempty" json:"randomChoice,omitempty" jsonschema:"oneof_required=RandomChoice,title=Random Choice Mask" jsonschema_description:"Replace the input value with a value chosen randomly from a constant list"`
	RandomChoiceInURI string               `yaml:"randomChoiceInUri,omitempty" json:"randomChoiceInUri,omitempty" jsonschema:"oneof_required=RandomChoiceInURI,title=Random Choice in URI" jsonschema_description:"Replace the input value with a value chosen randomly from an external resource (1 line = 1 value)"`
	RandomChoiceInCSV ChoiceInCSVType      `yaml:"randomChoiceInCSV,omitempty" json:"randomChoiceInCSV,omitempty" jsonschema:"oneof_required=RandomChoiceInCSV,title=Random Choice in CSV" jsonschema_description:"Replace the input value with a record chosen randomly from an external CSV resource"`
	Command           string               `yaml:"command,omitempty" json:"command,omitempty" jsonschema:"oneof_required=Command,title=Command Mask" jsonschema_description:"Replace the input value with the result of the given system command"`
	RandomInt         RandIntType          `yaml:"randomInt,omitempty" json:"randomInt,omitempty" jsonschema:"oneof_required=RandomInt,title=Random Integer Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an integer range"`
	WeightedChoice    []WeightedChoiceType `yaml:"weightedChoice,omitempty" json:"weightedChoice,omitempty" jsonschema:"oneof_required=WeightedChoice,title=Weighted Choice Mask" jsonschema_description:"Replace the input value with a value chosen randomly from a constant list, each value is given a weight (higher weight value has higher chance to be selected)"`
	Regex             string               `yaml:"regex,omitempty" json:"regex,omitempty" jsonschema:"oneof_required=Regex,title=Regex Mask" jsonschema_description:"Replace the input value with a random generated value that match the given regular expression"`
	Hash              []Entry              `yaml:"hash,omitempty" json:"hash,omitempty" jsonschema:"oneof_required=Hash,title=Hash Mask" jsonschema_description:"Replace the input value with a value chosen deterministically from a constant list, the same input will always be replaced by the same output"`
	HashInURI         string               `yaml:"hashInUri,omitempty" json:"hashInUri,omitempty" jsonschema:"oneof_required=HashInURI,title=Hash in URI Mask" jsonschema_description:"Replace the input value with a value chosen deterministically from an external resource (1 line = 1 value), the same input will always be replaced by the same output"`
	HashInCSV         ChoiceInCSVType      `yaml:"hashInCSV,omitempty" json:"hashInCSV,omitempty" jsonschema:"oneof_required=HashInCSV,title=Hash in CSV" jsonschema_description:"Replace the input value with a record chosen deterministically from an external CSV resource, the same input will always be replaced by the same output"`
	RandDate          RandDateType         `yaml:"randDate,omitempty" json:"randDate,omitempty" jsonschema:"oneof_required=RandDate,title=Random Date Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an date range"`
	Incremental       IncrementalType      `yaml:"incremental,omitempty" json:"incremental,omitempty" jsonschema:"oneof_required=Incremental,title=Incremental Mask" jsonschema_description:"Replace the input value with an incrementing integer sequence"`
	Replacement       string               `yaml:"replacement,omitempty" json:"replacement,omitempty" jsonschema:"oneof_required=Replacement,title=Replacement Mask" jsonschema_description:"Replace the input value with the value of another field"`
	Template          string               `yaml:"template,omitempty" json:"template,omitempty" jsonschema:"oneof_required=Template,title=Template Mask" jsonschema_description:"Replace the input value with the result of executing the given Golang template"`
	TemplateEach      TemplateEachType     `yaml:"template-each,omitempty" json:"template-each,omitempty" jsonschema:"oneof_required=TemplateEach,title=Template Each Mask" jsonschema_description:"Replace all input values (selector must be an array field) with the result of executing the given Golang template on each value"`
	Duration          string               `yaml:"duration,omitempty" json:"duration,omitempty" jsonschema:"oneof_required=Duration,title=Duration Mask" jsonschema_description:"Modify the input value (selector must be a date field) increasing or decreasing by the given amount of time"`
	Remove            bool                 `yaml:"remove,omitempty" json:"remove,omitempty" jsonschema:"oneof_required=Remove,title=Remove Mask" jsonschema_description:"Remove the field from the JSON stream"`
	RangeMask         int                  `yaml:"range,omitempty" json:"range,omitempty" jsonschema:"oneof_required=RangeMask,title=Range Mask" jsonschema_description:"Replace the integer value with a range of the given size"`
	RandomDuration    RandomDurationType   `yaml:"randomDuration,omitempty" json:"randomDuration,omitempty" jsonschema:"oneof_required=RandomDuration,title=Random Duration Mask" jsonschema_description:"Modify the input value (selector must be a date field) increasing or decreasing by a random amount of time"`
	FluxURI           string               `yaml:"fluxUri,omitempty" json:"fluxUri,omitempty" jsonschema:"oneof_required=FluxURI,title=Flux in URI Mask" jsonschema_description:"Replace the input value with the next value in the sequence given by an external resource (1 line = 1 value)"`
	RandomDecimal     RandomDecimalType    `yaml:"randomDecimal,omitempty" json:"randomDecimal,omitempty" jsonschema:"oneof_required=RandomDecimal,title=Random Decimal Mask" jsonschema_description:"Replace the input value with a value chosen randomly from an decimal range"`
	DateParser        DateParserType       `yaml:"dateParser,omitempty" json:"dateParser,omitempty" jsonschema:"oneof_required=DateParser,title=Date Parser Mask" jsonschema_description:"Change the format of the input date"`
	FromCache         string               `yaml:"fromCache,omitempty" json:"fromCache,omitempty" jsonschema:"oneof_required=FromCache,title=From Cache Mask" jsonschema_description:"Replace the input value with the value stored at the corresponding key in the given cache"`
	FF1               FF1Type              `yaml:"ff1,omitempty" json:"ff1,omitempty" jsonschema:"oneof_required=FF1,title=FF1 Mask" jsonschema_description:"Encrypt the input value using the FF1 algorithm (format preserving encryption)"`
	Pipe              PipeType             `yaml:"pipe,omitempty" json:"pipe,omitempty" jsonschema:"oneof_required=Pipe,title=Pipe Mask" jsonschema_description:"If the input value contains an array of object, stream each object with the given masking pipeline definition, this mask exists to handle complex data structures"`
	FromJSON          string               `yaml:"fromjson,omitempty" json:"fromjson,omitempty" jsonschema:"oneof_required=FromJSON,title=From JSON Mask" jsonschema_description:"Parse the input value as raw JSON, and add the resulting structure to the JSON stream"`
	Luhn              *LuhnType            `yaml:"luhn,omitempty" json:"luhn,omitempty" jsonschema:"oneof_required=Luhn,title=Luhn Mask" jsonschema_description:"Concatenate a checksum key to the input value computed with the luhn algorithm"`
	Markov            MarkovType           `yaml:"markov,omitempty" json:"markov,omitempty" jsonschema:"oneof_required=Markov,title=Markov Mask" jsonschema_description:"Produces pseudo text based on sample text"`
	Transcode         *TranscodeType       `yaml:"transcode,omitempty" json:"transcode,omitempty" jsonschema:"oneof_required=Transcode,title=Transcode Mask" jsonschema_description:"Produce a random string by preserving character classes from the original value"`
	FindInCSV         FindInCSVType        `yaml:"findInCSV,omitempty" json:"findInCSV,omitempty" jsonschema:"oneof_required=FindInCSV,title=Find in CSV Mask" jsonschema_description:"Find matched values in a CSV file based on input json file and save the matched csv line as type objet"`
	XML               XMLType              `yaml:"xml,omitempty" json:"xml,omitempty" jsonschema:"oneof_required=xml,title=XML Mask" jsonschema_description:"Apply mask for XML content within JSON values"`
	TimeLine          TimeLineType         `yaml:"timeline,omitempty" json:"timeline,omitempty" jsonschema:"oneof_required=TimeLine,title=TimeLine Mask" jsonschema_description:"Generate a timeline under constraints and rules"`
	Sequence          SequenceType         `yaml:"sequence,omitempty" json:"sequence,omitempty" jsonschema:"oneof_required=Sequence,title=Sequence Mask" jsonschema_description:"Generate a sequenced ID that follows specified format"`
	Sha3              Sha3Type             `yaml:"sha3,omitempty" json:"sha3,omitempty" jsonschema:"oneof_required=Sha3,title=Sha3 Mask" jsonschema_description:"Generate a variable-length crytographic hash (collision resistant)"`
	Apply             ApplyType            `yaml:"apply,omitempty" json:"apply,omitempty" jsonschema:"oneof_required=Apply,title=Apply Mask" jsonschema_description:"Call external masking file"`
	Partition         []PartitionType      `yaml:"partitions,omitempty" json:"partitions,omitempty" jsonschema:"oneof_required=Partition,title=Partitions Mask" jsonschema_description:"Identify specific cases and apply a defined list of masks for each case"`
	Segment           SegmentType          `yaml:"segments,omitempty" json:"segments,omitempty" jsonschema:"oneof_required=Segment,title=Segments Mask" jsonschema_description:"Allow transformations on specific parts of a field's value"`
	Log               *LogType             `yaml:"log,omitempty" json:"log,omitempty" jsonschema:"oneof_required=Log,title=Log Mask" jsonschema_description:"Output a log message"`
}

type Masking struct {
	// Masking requires at least one Selector and one Mask definition.
	// Case1: One selector, One mask
	// Case2: One selector, Multiple masks
	// Case3: Multiple selectors, One mask
	// Case4: Multiple selectors, Multiple masks
	Selector  SelectorType   `yaml:"selector,omitempty" json:"selector,omitempty" jsonschema:"oneof_required=case1,oneof_required=case2" jsonschema_description:"A selector defines on which field the mask will be applied"`
	Selectors []SelectorType `yaml:"selectors,omitempty" json:"selectors,omitempty" jsonschema:"oneof_required=case3,oneof_required=case4" jsonschema_description:"Defines on which fields the mask will be applied"`
	Mask      MaskType       `yaml:"mask,omitempty" json:"mask,omitempty" jsonschema:"oneof_required=case1,oneof_required=case3" jsonschema_description:"Defines how the selected value(s) will be masked"`
	Masks     []MaskType     `yaml:"masks,omitempty" json:"masks,omitempty" jsonschema:"oneof_required=case2,oneof_required=case4" jsonschema_description:"Defines how the selected value(s) will be masked"`
	Cache     string         `yaml:"cache,omitempty" json:"cache,omitempty" jsonschema_description:"Use an in-memory cache to preserve coherence between original/masked values"`
	Preserve  string         `yaml:"preserve,omitempty" json:"preserve,omitempty" jsonschema:"enum=null,enum=empty,enum=blank,enum=notInCache" jsonschema_description:"Preserve (do not mask) some values : null = preserve null value, empty = preserve empty strings, blank = preserve both null and empty values, notInCache = preserve value even if not present in cache (fromCache mask)"`
	PreserveL []Entry        `yaml:"preserve-list,omitempty" json:"preserve-list,omitempty" jsonschema_description:"Preserve (do not mask) given values"`
	Seed      SeedType       `yaml:"seed,omitempty" json:"seed,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the value given field"`
}

type SeedType struct {
	Field string `yaml:"field,omitempty" json:"field,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the given field value, a Golang Template can be used here"`
}

type CacheDefinition struct {
	Unique  bool `yaml:"unique,omitempty" json:"unique,omitempty" jsonschema_description:"The cache will not allow a masked value to be used multiple times, the mask will be reapplied until a unique value is generated"`
	Reverse bool `yaml:"reverse,omitempty" json:"reverse,omitempty" jsonschema_description:"Reverse the cache, keys will be used as values, and values will be used as keys"`
}
type Function struct {
	Params []Param `yaml:"params" json:"params" jsonschema_description:"Declare parameters function"`
	Body   string  `yaml:"body" json:"body" jsonschema_description:"Declare body function"`
}

type Param struct {
	Name string `yaml:"name" json:"name" jsonschema_description:"Declare name parameters"`
}

type Definition struct {
	Version   string                     `yaml:"version" json:"version" jsonschema_description:"Version of the pipeline definition, use the value 1"`
	Seed      int64                      `yaml:"seed,omitempty" json:"seed,omitempty" jsonschema_description:"Initialize the Pseaudo-Random-Generator with the given value"`
	Functions map[string]Function        `yaml:"functions,omitempty" json:"functions,omitempty" jsonschema_description:"Declare functions to be used in the masking"`
	Masking   []Masking                  `yaml:"masking" json:"masking" jsonschema_description:"Masking pipeline definition"`
	Caches    map[string]CacheDefinition `yaml:"caches,omitempty" json:"caches,omitempty" jsonschema_description:"Declare in-memory caches"`
}

/***************
 * REFACTORING *
 ***************/

// Processor process Dictionary and none, one or many element
type Processor interface {
	Open() error
	ProcessDictionary(Dictionary, Collector) error
}

// Collector collect Dictionary generate by Process
type Collector interface {
	Collect(Entry)
}

// SinkProcess send Dictionary process by Pipeline to an output
type SinkProcess interface {
	Open() error
	ProcessDictionary(Entry) error
}

type Pipeline interface {
	Source() Source
	Process(Processor) Pipeline
	WithSource(Source) Pipeline
	AddSink(SinkProcess) SinkedPipeline
}

type SinkedPipeline interface {
	Run() error
}

// Source is an iterator over Entry
type Source Iterable[Entry]

type Mapper func(Dictionary) (Dictionary, error)

/******************
 * IMPLEMENTATION *
 ******************/

func NewPipelineFromSlice(dictionaries []Dictionary) Pipeline {
	return SimplePipeline{source: &SliceSource{dictionaries}}
}

func NewSourceFromSlice(dictionaries []Dictionary) Source {
	return &SliceSource{dictionaries}
}

func NewRepeaterUntilProcess(source *TempSource, text, mode string, skipLogFile string) (Processor, error) {
	eng, err := template.NewEngine(text, tmpl.FuncMap{}, 0, "")
	var errlogger *MsgLogger
	if len(skipLogFile) > 0 {
		errlogger = NewMsgLogger(skipLogFile)
	}
	return RepeaterUntilProcess{eng, source, mode, errlogger}, err
}

type RepeaterUntilProcess struct {
	tmpl      *template.Engine
	tmp       *TempSource
	mode      string
	errlogger *MsgLogger
}

func (p RepeaterUntilProcess) Open() error {
	return nil
}

func (p RepeaterUntilProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	var output bytes.Buffer
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Cannot execute template, error: %v", r)
		}
	}()
	err = p.tmpl.Execute(&output, dictionary.UnpackUntyped())

	if err != nil && skipLineOnError {
		log.Warn().AnErr("error", err).Msg("Line skipped")
		statistics.IncIgnoredLinesCount()
		if p.errlogger != nil {
			if msg, ok := dictionary.GetValue("original"); !ok {
				return nil
			} else if msgstr, ok := msg.(string); !ok {
				return nil
			} else if err := p.errlogger.Log(msgstr); err != nil {
				log.Err(err)
			}
		}
		return nil
	}

	if err == nil {
		b := output.String()
		switch p.mode {
		case "while":
			p.tmp.repeat = b == "true"
			if p.tmp.repeat {
				out.Collect(dictionary)
			}
		case "until":
			p.tmp.repeat = b == "false"
			out.Collect(dictionary)
		default:
			p.tmp.repeat = false
			out.Collect(dictionary)
		}
	}

	return err
}

func NewTempSource(sourceValue Source) *TempSource {
	return &TempSource{repeat: false, source: sourceValue, value: NewPackedDictionary()}
}

type TempSource struct {
	repeat bool
	value  Entry
	source Source
}

func (s *TempSource) Open() error {
	return s.source.Open()
}

func (s *TempSource) Close() error {
	return s.source.Open()
}

func (s *TempSource) Values() iter.Seq2[Entry, error] {
	return func(yield func(Entry, error) bool) {
		for value, err := range s.source.Values() {
			if err != nil {
				yield(value, err)
				return
			}

			if !yield(Copy(value), nil) {
				return
			}

			for s.repeat {
				if !yield(Copy(value), nil) {
					return
				}
			}
		}
	}
}

type MapProcess struct {
	mapper Mapper
}

func NewMapProcess(mapper Mapper) Processor {
	return MapProcess{mapper: mapper}
}

func (mp MapProcess) Open() error {
	return nil
}

func (mp MapProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	mappedValue, err := mp.mapper(dictionary)
	if err != nil {
		return err
	}
	out.Collect(mappedValue)
	return nil
}

func NewSinkToSlice(dictionaries *[]Entry) SinkProcess {
	return &SinkToSlice{dictionaries}
}

type SinkToSlice struct {
	dictionaries *[]Entry
}

func (sink *SinkToSlice) Open() error {
	*sink.dictionaries = []Entry{}
	return nil
}

func (sink *SinkToSlice) ProcessDictionary(dictionary Entry) error {
	*sink.dictionaries = append(*sink.dictionaries, dictionary)
	return nil
}

func NewSinkToCache(cache Cache) SinkProcess {
	return &SinkToCache{cache}
}

type SinkToCache struct {
	cache Cache
}

func (sink *SinkToCache) Open() error {
	return nil
}

func (sink *SinkToCache) ProcessDictionary(entry Entry) error {
	dictionary := entry.(Dictionary)
	sink.cache.Put(CleanTypes(dictionary.Get("key")), CleanTypes(dictionary.Get("value")))
	return nil
}

type SimpleSinkedPipeline struct {
	source Source
	sink   SinkProcess
}

type SimplePipeline struct {
	source Source
}

func NewPipeline(source Source) Pipeline {
	return SimplePipeline{source: source}
}

func (pipeline SimplePipeline) Source() Source {
	return pipeline.source
}

func (pipeline SimplePipeline) WithSource(source Source) Pipeline {
	return SimplePipeline{source: source}
}

func (pipeline SimplePipeline) Process(process Processor) Pipeline {
	return NewProcessPipeline(pipeline.source, process)
}

func (pipeline SimplePipeline) AddSink(sink SinkProcess) SinkedPipeline {
	return SimpleSinkedPipeline{pipeline, sink}
}

func (pipeline SimplePipeline) Open() error {
	return pipeline.source.Open()
}

func (pipeline SimplePipeline) Close() error {
	return pipeline.source.Close()
}

func (pipeline SimplePipeline) Values() iter.Seq2[Entry, error] {
	return pipeline.source.Values()
}

func NewCollector() *QueueCollector {
	return &QueueCollector{[]Entry{}, NewDictionary()}
}

type QueueCollector struct {
	queue []Entry
	value Entry
}

func (c *QueueCollector) Err() error {
	return nil
}

func (c *QueueCollector) Open() error {
	return nil
}

func (c *QueueCollector) Close() error {
	return nil
}

func (c *QueueCollector) Values() iter.Seq2[Entry, error] {
	return func(yield func(Entry, error) bool) {
		for c.Next() {
			if !yield(c.Value(), nil) {
				return
			}
		}
	}
}

func (c *QueueCollector) Collect(dictionary Entry) {
	c.queue = append(c.queue, dictionary)
}

func (c *QueueCollector) Next() bool {
	if len(c.queue) > 0 {
		c.value = c.queue[0]
		c.queue = c.queue[1:]
		return true
	}
	return false
}

func (c *QueueCollector) Value() Entry {
	return c.value
}

func NewProcessPipeline(source Source, process Processor) Pipeline {
	return &ProcessPipeline{NewProcessWrapper(source, process)}
}

type ProcessPipeline struct {
	*ProcessWrapper
}

func (p *ProcessPipeline) AddSink(sink SinkProcess) SinkedPipeline {
	return SimpleSinkedPipeline{p, sink}
}

func (p *ProcessPipeline) Process(process Processor) Pipeline {
	return NewProcessPipeline(p, process)
}

func (p *ProcessPipeline) Source() Source {
	return p.input
}

func (p *ProcessPipeline) WithSource(source Source) Pipeline {
	if s, ok := p.input.(*ProcessPipeline); ok {
		return NewProcessPipeline(s.WithSource(source).(Source), p.processor)
	}
	return NewProcessPipeline(source, p.processor)
}

func (pipeline SimpleSinkedPipeline) Run() (err error) {
	err = pipeline.source.Open()
	if err != nil {
		return err
	}

	err = pipeline.sink.Open()
	if err != nil {
		return err
	}

	for item, err := range pipeline.source.Values() {
		if err != nil {
			return err
		}

		err := pipeline.sink.ProcessDictionary(item)
		if err != nil {
			return err
		}
	}
	return nil
}

type Seeder func(Dictionary) (int64, bool, error)

func NewSeeder(sourceField string, seed int64) Seeder {
	var seeder Seeder

	if jpath := sourceField; jpath != "" {
		sel := NewPackedPathSelector(jpath)
		hash := fnv.New64a()
		seeder = func(context Dictionary) (int64, bool, error) {
			e, ok := sel.Read(context)
			if !ok {
				return 0, ok, nil
			}
			hash.Reset()
			_, err := hash.Write([]byte(fmt.Sprintf("%v", e)))
			return int64(hash.Sum64()) + seed, true, err //nolint:gosec
		}
	} else {
		seeder = func(context Dictionary) (int64, bool, error) {
			return seed, false, nil
		}
	}
	return seeder
}
