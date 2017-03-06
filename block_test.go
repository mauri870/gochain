package gochain

import "testing"

func TestMakeHash(t *testing.T) {
	for _, test := range []string{"Hello world", "dasdsad", "test"} {
		h, err := MakeHash(test)
		if err != nil {
			t.Error(err)
		}

		if len(h) != 64 {
			t.Error("Expect a hash with 64 characters, got", len(h))
		}
	}
}
