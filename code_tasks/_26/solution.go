package _26

import (
	"unicode"
	"unicode/utf8"
)

func IsStringOfUnique(s string) bool {
	seenRunes := make(map[rune]struct{})
	seenBytes := make(map[byte]struct{}) // for non-runes
	for len(s) > 0 {
		offset := 0
		if rn, sz := utf8.DecodeRuneInString(s); rn != utf8.RuneError {
			lower, upper := rn, rn
			if unicode.IsLower(rn) {
				upper = unicode.ToUpper(rn)
			} else {
				lower = unicode.ToLower(rn)
			}
			if _, found := seenRunes[lower]; found {
				return false
			}
			if _, found := seenRunes[upper]; found {
				return false
			}
			seenRunes[lower] = struct{}{}
			seenRunes[upper] = struct{}{}
			offset = sz
		} else {
			if _, found := seenBytes[s[0]]; found {
				return false
			}
			seenBytes[s[0]] = struct{}{}
			offset = 1
		}
		s = s[offset:]
	}
	return true
}
