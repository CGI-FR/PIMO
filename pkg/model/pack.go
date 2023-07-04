package model

// PackProcess will pack the jsonline data into a wrapper {"": <jsonline>}
type PackProcess struct{}

func NewPackProcess() *PackProcess {
	return &PackProcess{}
}

func (pp *PackProcess) Open() error {
	return nil
}

func (pp *PackProcess) ProcessDictionary(entry Entry, out Collector) error {
	dictionary := entry.(Dictionary)
	if dictionary.IsPacked() {
		out.Collect(dictionary)
	} else {
		out.Collect(dictionary.Pack())
	}
	return nil
}

// UnpackProcess will unpack the jsonline data from a wrapper {"": <jsonline>}
type UnpackProcess struct{}

func NewUnpackProcess() *UnpackProcess {
	return &UnpackProcess{}
}

func (up *UnpackProcess) Open() error {
	return nil
}

func (up *UnpackProcess) ProcessDictionary(dictionary Entry, out Collector) error {
	out.Collect(dictionary.(Dictionary).Unpack())
	return nil
}
