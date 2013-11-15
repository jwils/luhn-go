package luhn

import (
	"testing"
)

func TestValid(t *testing.T) {
	valid_tests := [...]string{"4242424242424242", "7992939871", "12345678903", "4444333322221111"}

	for _, test := range valid_tests {

		if !Valid(test) {
			t.Errorf("%q should have been a valid luhn number", test)
		}
	}

	invalid_tests := [...]string{"123456789", "222222222222", "9875693324", "323545454234432343243"}
	for _, test := range invalid_tests {

		if Valid(test) {
			t.Errorf("%q should have been an invalid luhn number", test)
		}
	}
}