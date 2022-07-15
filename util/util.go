package util

import (
	"github.com/gijsbers/go-pcre"
)

// Match :
// Helper function for pattern matching.
func Match(pattern string, target string) bool {
	uspRgx := pcre.MustCompile(pattern, 0)
	uspMatcher := uspRgx.MatcherString(target, 0)
	return uspMatcher.Matches()
}
