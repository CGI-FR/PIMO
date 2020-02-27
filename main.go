package main

type MaskFunction = func(interface{}) interface{}
type MaskConfiguration = map[string]MaskFunction
type Dictionnary = map[string]interface{}

// MaskingFactory return Masking function data without private information
func MaskingFactory(config MaskConfiguration) func(Dictionnary) Dictionnary {
	return func(input Dictionnary) Dictionnary {
		output := Dictionnary{}
		for k, v := range input {
			mask, ok := config[k]
			if ok {
				output[k] = mask(v)
			} else {
				output[k] = v
			}
		}
		return output
	}
}
