package model

import (
	over "github.com/Trendyol/overlog"
	"github.com/rs/zerolog/log"
)

type CounterProcess struct {
	contextName string
}

func NewCounterProcess(contextName string, initValue int) Processor {
	process := CounterProcess{contextName}
	err := process.Open()
	if err != nil {
		log.Warn().AnErr("err", err).Msg("Should not happen")
	}
	over.MDC().Set(contextName, initValue)
	return CounterProcess{contextName}
}

func (p CounterProcess) Open() error {
	over.AddGlobalFields(p.contextName)
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
		over.MDC().Set(p.contextName, counter+1)
	}

	out.Collect(dictionary)
	return nil
}
