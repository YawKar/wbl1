package _20

import "testing"

var (
	tcases []string
	tvalid []string
)

func init() {
	tcases = []string{
		"Hello\t\nWorld",
		"\n\t\r\n\t",
		"hello",
		"				главрыба  ",
		"snow dog sun",
		"Red Wine",
		"Закат Рассвет",
		"Владивосток\n\t\t 2000",
		"Pourvu   qu'elles soient douces			",
		string([]byte{1, 2, 3}),
		"\xe2\x28\xa1",         // Invalid utf8 3 Octet Sequence (in 2nd Octet)
		"\xf8\xa1\xa1\xa1\xa1", // Valid 5 Octet Sequence (but not Unicode!)
	}
	tvalid = []string{
		"World\n\tHello",
		"\t\n\r\t\n",
		"hello",
		"  главрыба				",
		"sun dog snow",
		"Wine Red",
		"Рассвет Закат",
		"2000 \t\t\nВладивосток",
		"			douces soient qu'elles   Pourvu",
		string([]byte{1, 2, 3}),
		"\xe2\x28\xa1",         // Invalid utf8 3 Octet Sequence (in 2nd Octet)
		"\xf8\xa1\xa1\xa1\xa1", // Valid 5 Octet Sequence (but not Unicode!)
	}
}

func TestSamples(t *testing.T) {
	for i, tcase := range tcases {
		reversed := ReverseWords(tcase)
		if reversed != tvalid[i] {
			t.Errorf("#%d test: %q != %q", i, reversed, tvalid[i])
		}
	}
}
