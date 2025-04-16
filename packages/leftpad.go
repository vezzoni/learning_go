package leftpad

import (
	"strings"
	"unicode/utf8"
)

// variable declared at the package level
// visible to everything within the package
// cannot be used with the shorted hand declaration and assignment (:=)
var default_char = ' '

func Format(s string, size int) string {
	return FormatRune(s, size, default_char)
}

func FormatRune(s string, size int, r rune) string {
	l := utf8.RuneCountInString(s)
	if l >= size {
		return s
	}
	out := strings.Repeat(string(r), size-l) + s
	return out
}
