package _19

import "testing"

var (
	tcases []string
	tvalid []string
)

func init() {
	tcases = []string{
		"hello",
		"главрыба",
		"snow dog sun",
		"Red Wine",
		"Закат Рассвет",
		"Владивосток 2000",
		"Pourvu qu'elles soient douces",
		string([]byte{1, 2, 3}),
		"\xe2\x28\xa1",         // Invalid utf8 3 Octet Sequence (in 2nd Octet)
		"\xf8\xa1\xa1\xa1\xa1", // Valid 5 Octet Sequence (but not Unicode!)
	}
	tvalid = []string{
		"olleh",
		"абырвалг",
		"nus god wons",
		"eniW deR",
		"тевссаР такаЗ",
		"0002 котсовидалВ",
		"secuod tneios selle'uq uvruoP",
		string([]byte{3, 2, 1}),
		"\xa1\x28\xe2",         // Invalid utf8 3 Octet Sequence (in 2nd Octet)
		"\xa1\xa1\xa1\xa1\xf8", // Valid 5 Octet Sequence (but not Unicode!)
	}
}

func TestSamples(t *testing.T) {
	for i, tcase := range tcases {
		reversed := ReverseString(tcase)
		if reversed != tvalid[i] {
			t.Errorf("#%d test: %q != %q", i, reversed, tvalid[i])
		}
	}
}
