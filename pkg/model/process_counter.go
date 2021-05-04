package model

import (
	over "github.com/Trendyol/overlog"
)

type CounterProcess struct {
	contextName string
	updater     func(int)
}

func NewCounterProcess(contextName string, initValue int) Processor {
	return NewCounterProcessWithCallback(contextName, initValue, nil)
}

func NewCounterProcessWithCallback(contextName string, initValue int, updater func(int)) Processor {
	over.MDC().Set(contextName, initValue)
	return CounterProcess{contextName, updater}
}

func (p CounterProcess) Open() error {
	// we don't do that, because sometime we want a hidden counter
	// over.AddGlobalFields(p.contextName)
	return nil
}

func (p CounterProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	value, exists := over.MDC().Get(p.contextName)
	if !exists {
		err := p.Open()
		if err != nil {
			return err
		}
		return p.ProcessDictionary(dictionary, out)
	}

	if counter, ok := value.(int); ok {
		counter++
		over.MDC().Set(p.contextName, counter)
		if p.updater != nil {
			p.updater(counter)
		}
	}

	out.Collect(dictionary)
	return nil
}
