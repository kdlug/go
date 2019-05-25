package strings

import "testing"

func TestReverse(t *testing.T) {
	var tests = []struct {
		s    string
		want string
	}{
		{"Backward", "drawkcaB"},
		{"What", "tahW"},
		{"", ""},
		{"Hi, 你好", "好你 ,iH"},
	}

	for _, test := range tests {
		got := Reverse(test.s)
		if got != test.want {
			t.Errorf("want '%s' but got '%s'", test.want, got)
		}
	}

}
