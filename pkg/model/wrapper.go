package model

import "iter"

type ProcessWrapper struct {
	input     Iterable[Dictionary]
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

func (c *ProcessWrapper) Values() iter.Seq2[Dictionary, error] {
	return func(yield func(Dictionary, error) bool) {
		for item, err := range c.input.Values() {
			if err != nil {
				yield(item, err)
				return
			}

			if err := c.processor.ProcessDictionary(item, c.queue); err != nil {
				yield(item, err)
				return
			}

			for c.queue.Next() {
				if !yield(c.queue.Value().(Dictionary), nil) {
					return
				}
			}
		}
	}
}

func NewProcessWrapper(input Iterable[Dictionary], p Processor) *ProcessWrapper {
	return &ProcessWrapper{input, p, NewCollector()}
}
