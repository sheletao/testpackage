package utils

import "testing"

type In struct {
	in, want string
}

func TestReverse(t *testing.T) {
	testcases := []In{
		{in: "Hello, world", want: "dlrow ,olleH"},
		{in: " ", want: " "},
		{in: "!12345", want: "54321!"},
	}

	for _, tc := range testcases {
		rev := Reverse(tc.in)

		if rev != tc.want {
			t.Errorf("Reverse：%q, want: %q", rev, tc.want)
		}
	}
}
