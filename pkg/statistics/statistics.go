package statistics

import (
	"encoding/json"
	"time"

	over "github.com/adrienaury/zeromdc"
	"github.com/rs/zerolog/log"
)

type ExecutionStats interface {
	WithErrorCode(ec int) ExecutionStats

	GetErrorCode() int
	GetIgnoredPathsCount() int  // counter for path not found in data
	GetIgnoredLinesCount() int  // counter for line skipped (flag --skip-line-on-error)
	GetIgnoredFieldsCount() int // counter for field skipped (flag --skip-field-on-error)
	GetDuration() time.Duration

	ToJSON() []byte
}

type stats struct {
	errorCode            int
	IgnoredPathsCounter  int           `json:"ignoredPaths"`
	IgnoredLinesCounter  int           `json:"skippedLines"`
	IgnoredFieldsCounter int           `json:"skippedFields"`
	Duration             time.Duration `json:"duration"`
}

// Reset all statistics to zero
func Reset() {
	over.MDC().Set("stats", &stats{})
}

func WithErrorCode(ec int) ExecutionStats {
	return &stats{errorCode: ec}
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

func (s *stats) WithErrorCode(ec int) ExecutionStats {
	s.SetErrorCode(ec)
	return s
}

func (s *stats) SetErrorCode(ec int) {
	s.errorCode = ec
}

func (s *stats) GetErrorCode() int {
	return s.errorCode
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

func SetDuration(duration time.Duration) {
	stats := getStats()
	stats.Duration = duration
}

func (s *stats) GetDuration() time.Duration {
	return s.Duration
}
