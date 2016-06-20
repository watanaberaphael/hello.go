package stringutil

import "testing"

func TestReverse(t *testing.T) {

	cases := []struct { in, want string	}{
		{"Hello, world", "dlrow ,olleH"},
		{"Hello, 世界", "界世 ,olleH"},
		{"", ""},
	}

	for _, c := range cases {

		got := Reverse(c.in)

		if got != c.want {
			t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
		}		
	}
}

func TestReverseAMessage(t *testing.T) {

	message := Reverse("dlrow ,olleH")

	if message != "Hello, world" {
		t.Errorf("Reverse doesn't work")
	}
}