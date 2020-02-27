package main

// Dictionary is a Map with string as key and Entry as value
type Dictionary = map[string]Entry

// Entry is a dictionary value
type Entry interface{}

// MaskEngine is a masking algorithm
type MaskEngine interface {
	Mask(Entry) Entry
}

// MaskConfiguration is a configuration to mask dictionaries content
type MaskConfiguration interface {
	GetMaskingEngine(string) (MaskEngine, bool)
}

// MapMaskConfiguration Implements MaskConfiguration with a map
type MapMaskConfiguration struct {
	config map[string]MaskEngine
}

// GetMaskingEngine return the MaskEngine configured for that string
func (mmc MapMaskConfiguration) GetMaskingEngine(key string) (MaskEngine, bool) {

	engine, test := mmc.config[key]
	return engine, test
}

// FunctionMaskEngine implements MaskEngine with a simple function
type FunctionMaskEngine struct {
	function func(Entry) Entry
}

// Mask delegate mask algorithm to the function
func (fme FunctionMaskEngine) Mask(e Entry) Entry { return fme.function(e) }

// MaskingFactory return Masking function data without private information
func MaskingFactory(config MaskConfiguration) MaskEngine {
	return FunctionMaskEngine{func(input Entry) Entry {
		output := Dictionary{}

		dico, ok := input.(Dictionary)
		if !ok {
			return input
		}

		for k, v := range dico {
			engine, ok := config.GetMaskingEngine(k)
			if ok {
				output[k] = engine.Mask(v)
			} else {
				output[k] = v
			}
		}
		return output
	}}
}
