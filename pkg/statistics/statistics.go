package statistics

import (
	"encoding/json"

	over "github.com/Trendyol/overlog"
	"github.com/rs/zerolog/log"
)

type ExecutionStats interface {
	GetIgnoredPathsCount() int  // counter for path not found in data
	GetIgnoredLinesCount() int  // counter for line skipped (flag --skip-line-on-error)
	GetIgnoredFieldsCount() int // counter for field skipped (flag --skip-field-on-error)

	ToJSON() []byte
}

type stats struct {
	IgnoredPathsCounter  int `json:"ignoredPaths"`
	IgnoredLinesCounter  int `json:"skippedLines"`
	IgnoredFieldsCounter int `json:"skippedFields"`
}

// Reset all statistics to zero
func Reset() {
	over.MDC().Set("stats", &stats{})
}

// Compute current statistics and give a snapshot
func Compute() ExecutionStats {
	value, exists := over.MDC().Get("stats")
	if stats, ok := value.(ExecutionStats); exists && ok {
		return stats
	}
	log.Warn().Msg("Unable to compute statistics")
	return &stats{}
}

func (s *stats) ToJSON() []byte {
	b, err := json.Marshal(s)
	if err != nil {
		log.Warn().Msg("Unable to read statistics")
	}
	return b
}

func (s *stats) GetIgnoredPathsCount() int {
	return s.IgnoredPathsCounter
}

func (s *stats) GetIgnoredLinesCount() int {
	return s.IgnoredLinesCounter
}

func (s *stats) GetIgnoredFieldsCount() int {
	return s.IgnoredFieldsCounter
}

func IncIgnoredPathsCount() {
	stats := getStats()
	stats.IgnoredPathsCounter++
}

func IncIgnoredLinesCount() {
	stats := getStats()
	stats.IgnoredLinesCounter++
}

func IncIgnoredFieldsCount() {
	stats := getStats()
	stats.IgnoredFieldsCounter++
}

// Compute current statistics and give a snapshot
func getStats() *stats {
	value, exists := over.MDC().Get("stats")
	if stats, ok := value.(*stats); exists && ok {
		return stats
	}
	log.Warn().Msg("Statistics uncorrectly initialized")
	return &stats{}
}
