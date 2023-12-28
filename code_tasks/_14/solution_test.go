package _14

import "testing"

func TestSamples(t *testing.T) {
	tests := []struct {
		v any
		k Kind
	}{
		{4, Int},
		{"str", String},
		{true, Bool},
		{make(chan struct{}), Chan},
		{nil, Undefined},
	}
	for i, tcase := range tests {
		if kind := KindOf(tcase.v); kind != tcase.k {
			t.Errorf("#%d test: kind of v is %v but %v was returned", i, tcase.k, kind)
		}
	}
}
