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
	"math/rand"
)

func NewPermuteProcess(seed int64, selector Selector) Processor {
	return &PermuteProcess{selector, NewPermuteMemory(seed), &QueueCollector{}, &QueueCollector{}}
}

func NewPermuteMemory(seed int64) *permuteMemory {
	// nolint: gosec
	rand := rand.New(rand.NewSource(seed))
	return &permuteMemory{[]Entry{}, rand}
}

type permuteMemory struct {
	memory []Entry
	rand   *rand.Rand
}

func (p permuteMemory) Len() int {
	return len(p.memory)
}

func (p *permuteMemory) pick() (Entry, bool) {
	if p.Len() > 0 {
		idx := p.rand.Int31n(int32(p.Len()))
		result := p.memory[idx]
		p.memory = append(p.memory[:idx], p.memory[idx+1:]...)
		return result, true
	}
	return nil, false
}

func (p *permuteMemory) add(entry Entry) {
	p.memory = append(p.memory, entry)
}

type PermuteProcess struct {
	selector  Selector
	shuffler  *permuteMemory
	readiness *QueueCollector
	waiting   *QueueCollector
}

func (p *PermuteProcess) Open() error {
	return nil
}

func (p *PermuteProcess) ProcessDictionary(dictionary Dictionary, out Collector) error {
	for p.readiness.Next() {
		p.processDictionary(p.readiness.value, out)
	}
	p.processDictionary(dictionary, out)
	return nil
}

func (p *PermuteProcess) processDictionary(dictionary Dictionary, out Collector) {
	key, ok := p.selector.Read(dictionary)
	if !ok {
		out.Collect(dictionary)
		return
	}
	p.waiting.Collect(dictionary)

	p.shuffler.add(key)
}

func (p *PermuteProcess) ShufflePick(out Collector) {
	for p.waiting.Next() {
		dictionary := p.waiting.Value()
		value, isPicked := p.shuffler.pick()
		if isPicked {
			out.Collect(p.selector.Write(dictionary, value))
		} else {
			p.waiting.Collect(dictionary)
		}
	}
}
