package esr

import "testing"

// Test some valid numbers
func TestValidNums(t *testing.T) {
	validNums := []string{"639400", "666408", "166507", "310202", "165504", "0"}
	for _, item := range validNums {
		if err := Validate(item); err == ErrValueNotCommon {
			t.Error("Valid number validated as invalid", item)
		}
	}
}

// Test some invalid numbers
func TestInvalidNums(t *testing.T) {
	invalidNums := []string{"639401", "666409", "166506", "310200", "165509", ""}
	for _, item := range invalidNums {
		if err := Validate(item); err != ErrValueNotCommon {
			t.Error("Invalid number validated as valid", item)
		}
	}
}
