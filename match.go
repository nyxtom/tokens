package tokens

import "regexp"

// MatchAny will see if the given string matches any of the given regular expressions
func MatchAny(text string, patterns ...*regexp.Regexp) bool {
	for _, p := range patterns {
		if p.MatchString(text) {
			return true
		}
	}

	return false
}
