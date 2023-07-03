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
	if _, ok := dictionary.GetValue("."); ok {
		out.Collect(dictionary)
	} else {
		packed := NewDictionary().With(".", dictionary)
		out.Collect(packed)
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
	unpacked := dictionary.(Dictionary).Get(".")
	out.Collect(unpacked)
	return nil
}
