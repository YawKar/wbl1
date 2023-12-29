package _19

import (
	"unicode/utf8"
)

// if `s` is a valid utf8 string it reverses the string rune-wise,
// otherwise it does it byte-wise
func ReverseString(s string) string {
	reversed := make([]byte, 0, len(s))
	if utf8.ValidString(s) {
		for len(s) > 0 {
			rn, size := utf8.DecodeLastRuneInString(s)
			reversed = utf8.AppendRune(reversed, rn)
			s = s[:len(s)-size]
		}
	} else {
		for len(s) > 0 {
			reversed = append(reversed, s[len(s)-1])
			s = s[:len(s)-1]
		}
	}
	return string(reversed)
}
