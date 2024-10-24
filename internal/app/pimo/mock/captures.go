package mock

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/cgi-fr/pimo/pkg/model"
)

type Matcher struct {
	expr  string
	regex *regexp.Regexp
}

func NewMatcher(expr string) *Matcher {
	placeholderRegex := regexp.MustCompile(`\{([^\}]*)\}`)
	regexPattern := placeholderRegex.ReplaceAllString(expr, `(?P<$1>[^/]*)`)
	regexPattern = "^" + regexPattern + "$" // append only if not present
	return &Matcher{
		expr:  expr,
		regex: regexp.MustCompile(regexPattern),
	}
}

type MatchResult struct {
	model.Dictionary
	matched  bool
	origin   string
	template string
}

func (m *Matcher) Match(path string) *MatchResult {
	matches := m.regex.FindStringSubmatchIndex(path)
	names := m.regex.SubexpNames()

	if matches == nil {
		return &MatchResult{
			Dictionary: model.NewDictionary(),
			matched:    false,
			origin:     path,
			template:   path,
		}
	}

	template := &strings.Builder{}

	c := 0
	for i := 2; i < len(matches); i += 2 {
		template.WriteString(path[c:matches[i]])
		template.WriteRune('{')
		template.WriteString(names[i/2])
		template.WriteRune('}')
		c = matches[i+1]
	}
	template.WriteString(path[c:])

	captures := model.NewDictionary()
	submatches := m.regex.FindStringSubmatch(path)
	for i, submatch := range submatches {
		captures.Set(names[i], submatch)
	}
	captures.Delete("")

	return &MatchResult{
		Dictionary: captures,
		matched:    true,
		origin:     path,
		template:   template.String(),
	}
}

func (mr *MatchResult) Build(new model.Dictionary) string {
	result := mr.template
	for _, key := range mr.Keys() {
		val := new.Get(key)
		placeholder := "{" + key + "}"
		result = strings.ReplaceAll(result, placeholder, entryToString(val))
	}
	return result
}

func entryToString(e model.Entry) string {
	switch typed := e.(type) {
	case string:
		return typed
	case int:
		return strconv.FormatInt(int64(typed), 10)
	case int64:
		return strconv.FormatInt(typed, 10)
	case int32:
		return strconv.FormatInt(int64(typed), 10)
	case int16:
		return strconv.FormatInt(int64(typed), 10)
	case int8:
		return strconv.FormatInt(int64(typed), 10)
	case uint:
		return strconv.FormatInt(int64(typed), 10) //nolint:gosec
	case uint64:
		return strconv.FormatInt(int64(typed), 10) //nolint:gosec
	case uint32:
		return strconv.FormatInt(int64(typed), 10)
	case uint16:
		return strconv.FormatInt(int64(typed), 10)
	case uint8:
		return strconv.FormatInt(int64(typed), 10)
	case float32:
		return strconv.FormatInt(int64(typed), 10)
	case float64:
		return strconv.FormatInt(int64(typed), 10)
	}
	return fmt.Sprintf("%s", e)
}
