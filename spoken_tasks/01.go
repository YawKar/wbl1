package main

import (
	"bytes"
	"strings"
)

// Old way: using buffers in *amortized* O(n)
// After Reset() keeps buffer allocated but sets length to 0
// Allows accessing raw bytes through Bytes()
func efficientlyConcatenateStrings(strs ...string) string {
	buf := bytes.NewBufferString("")
	for _, str := range strs {
		buf.WriteString(str)
	}
	return buf.String()
}

// Modern way: using string.Builder in *amortized* O(n)
// After Reset() detaches buffer
// Panics if copied by value
// Prevents from accessing raw bytes
func efficientlyConcatenateStrings2(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(str)
	}
	return sb.String()
}
