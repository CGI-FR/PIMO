package markov

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"unicode"
)

const (
	StartToken = "^"
	EndToken   = "$"
)

type FrequencyMap map[string]int

type Link struct {
	CurrentState NGram
	NextState    string
}

type NGram []string

func (n NGram) String(separator string) string {
	return strings.Join(n, separator)
}

func NewChain(orderValue int, randValue *rand.Rand, separatorValue string) Chain {
	return Chain{
		order:     orderValue,
		rand:      randValue,
		graph:     make(map[string]FrequencyMap),
		separator: separatorValue,
	}
}

type Chain struct {
	order     int
	rand      *rand.Rand
	graph     map[string]FrequencyMap
	separator string
}

func (chain *Chain) Add(input []string) {
	var tokens []string
	tokens = append(tokens, array(StartToken, chain.order)...)
	tokens = append(tokens, splitPunct(input...)...)
	tokens = append(tokens, array(EndToken, chain.order)...)
	links := makeLinks(tokens, chain.order)
	for _, link := range links {
		c, n := link.CurrentState.String(chain.separator), link.NextState
		if _, ok := chain.graph[c]; !ok {
			chain.graph[c] = make(FrequencyMap)
		}
		chain.graph[c][n]++
	}
}

func (chain *Chain) Generate(current NGram) (string, error) {
	if len(current) != chain.order {
		return "", fmt.Errorf("N-gram length does not match chain order")
	}
	if current[len(current)-1] == EndToken {
		// Dont generate anything after the end token
		return "", nil
	}
	freqMap, currentExists := chain.graph[current.String(chain.separator)]
	if !currentExists {
		return "", fmt.Errorf("Unknown ngram %v", current)
	}

	keys := make([]string, len(freqMap))
	idx := 0
	for k := range freqMap {
		keys[idx] = k
		idx++
	}
	sort.Slice(keys, func(i, j int) bool {
		return freqMap[keys[i]] < freqMap[keys[j]]
	})

	totals := make([]int, len(keys))
	runningTotal := 0
	for i := range keys {
		runningTotal += int(freqMap[keys[i]])
		totals[i] = runningTotal
	}
	r := chain.rand.Intn(runningTotal) + 1
	i := sort.SearchInts(totals, r)

	return keys[i], nil
}

func makeLinks(tokens []string, order int) (links []Link) {
	for i := 0; i < len(tokens)-order; i++ {
		link := Link{
			CurrentState: tokens[i : i+order],
			NextState:    tokens[i+order],
		}
		links = append(links, link)
	}
	return
}

func array(input string, order int) (result []string) {
	result = make([]string, order)
	for i := 0; i < order; i++ {
		result[i] = input
	}
	return
}

func splitPunct(list ...string) []string {
	result := []string{}
	for _, el := range list {
		result = append(result, splitFunc(el, unicode.IsPunct)...)
	}
	return result
}

func splitFunc(s string, f func(rune) bool) []string {
	// A span is used to record a slice of s of the form s[start:end].
	// The start index is inclusive and the end index is exclusive.
	type span struct {
		start int
		end   int
	}
	spans := make([]span, 0, 32)

	// Find the field start and end indices.
	// Doing this in a separate pass (rather than slicing the string s
	// and collecting the result substrings right away) is significantly
	// more efficient, possibly due to cache effects.
	start := -1 // valid span start if >= 0
	for end, rune := range s {
		if f(rune) {
			if start >= 0 {
				spans = append(spans, span{start, end})
				spans = append(spans, span{end, end + 1})
				// Set start to a negative value.
				// Note: using -1 here consistently and reproducibly
				// slows down this code by a several percent on amd64.
				start = ^start
			}
		} else {
			if start < 0 {
				start = end
			}
		}
	}

	// Last field might end at EOF.
	if start >= 0 {
		spans = append(spans, span{start, len(s)})
	}

	// Create strings from recorded field indices.
	a := make([]string, len(spans))
	for i, span := range spans {
		a[i] = s[span.start:span.end]
	}

	return a
}
