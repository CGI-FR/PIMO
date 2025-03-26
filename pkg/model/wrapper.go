package model

import "iter"

type ProcessWrapper struct {
	input     Source
	processor Processor
	queue     *QueueCollector
}

func (c *ProcessWrapper) Open() error {
	if err := c.input.Open(); err != nil {
		return err
	}

	if err := c.queue.Open(); err != nil {
		return err
	}

	return c.processor.Open()
}

func (c *ProcessWrapper) Close() error {
	if err := c.input.Close(); err != nil {
		return err
	}
	return nil
}

func (c *ProcessWrapper) Values() iter.Seq2[Entry, error] {
	return func(yield func(Entry, error) bool) {
		for item, err := range c.input.Values() {
			if err != nil {
				yield(item, err)
				return
			}

			if err := c.processor.ProcessDictionary(item.(Dictionary), c.queue); err != nil {
				yield(item, err)
				return
			}

			for c.queue.Next() {
				if !yield(c.queue.Value(), nil) {
					return
				}
			}
		}
	}
}

func NewProcessWrapper(input Source, p Processor) *ProcessWrapper {
	return &ProcessWrapper{input, p, NewCollector()}
}
