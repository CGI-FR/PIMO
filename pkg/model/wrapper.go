package model

import "iter"

type ProcessWrapper struct {
	input Iterable[Dictionary]
	p     Processor
	queue *QueueCollector
}

func (c *ProcessWrapper) Open() error {
	return nil
}

func (c *ProcessWrapper) Close() error {
	return nil
}

func (c *ProcessWrapper) Values() iter.Seq2[Dictionary, error] {
	return func(yield func(Dictionary, error) bool) {
		for item, err := range c.input.Values() {
			if err != nil {
				yield(item, err)
				return
			}

			if err := c.p.ProcessDictionary(item, c.queue); err != nil {
				yield(item, err)
				return
			}

			if c.queue.Next() {
				if !yield(c.queue.Value().(Dictionary), nil) {
					return
				}
			}
		}
	}
}

func NewProcessWrapper(input Iterable[Dictionary], p Processor) Iterable[Dictionary] {
	return &ProcessWrapper{input, p, NewCollector()}
}
