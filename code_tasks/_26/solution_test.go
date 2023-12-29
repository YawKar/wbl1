package _26

import "testing"

func TestSamples(t *testing.T) {
	tcases := []struct {
		s      string
		unique bool
	}{
		{"abcd", true},
		{"abCdefAaf", false},
		{"aabcd", false},
	}
	for i, tcase := range tcases {
		if res := IsStringOfUnique(tcase.s); res != tcase.unique {
			t.Errorf("#%d test: results differ: %t != %t", i, res, tcase.unique)
		}
	}
}
