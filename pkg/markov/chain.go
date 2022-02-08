package markov

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
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
	tokens = append(tokens, input...)
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
	for k, i := range freqMap {
		keys[i] = k
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
