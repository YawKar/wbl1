package _20

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// it reverses order of words and order of whitespace sequences
//
// e.g.: "Hello\t\nWorld" -> "World\n\tHello"
func ReverseWords(s string) string {
	result := make([]byte, 0, len(s))
	fields := strings.Fields(s)
	for len(s) > 0 {
		sOffset := 0
		fOffset := 0
		// if it is a space then just push it
		if rn, sz := utf8.DecodeLastRuneInString(s); rn != utf8.RuneError && unicode.IsSpace(rn) {
			result = utf8.AppendRune(result, rn)

			sOffset = sz
		} else { // otherwise we got a word
			word := fields[len(fields)-1]
			result = append(result, word...)

			sOffset = len(word)
			fOffset = 1
		}
		s = s[:len(s)-sOffset]
		fields = fields[:len(fields)-fOffset]
	}
	return string(result)
}
