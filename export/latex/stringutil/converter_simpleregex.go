package stringutil

import "regexp"

type simpleRegexConverter struct {
	regex       *regexp.Regexp
	replacement string
}

func NewSimpleRegexConverter(regexStr string, replacementStr string) *simpleRegexConverter {
	return &simpleRegexConverter{
		regex:       regexp.MustCompile(regexStr),
		replacement: replacementStr,
	}
}

func (conv *simpleRegexConverter) Process(text string) string {
	return conv.regex.ReplaceAllString(text, conv.replacement)
}
