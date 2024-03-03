package esr

import "testing"

// Test some valid numbers
func TestValidNums(t *testing.T) {
	validNums := []string{"639400", "666408", "166507", "310202", "165504", "0", "197511", "924605"}
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

// Test generate invalid numbers
func TestGenerateNums(t *testing.T) {
	validNums := []string{"63940", "66640", "16650", "31020", "16550", "0", "19751"}
	for _, item := range validNums {
		code := Generate(item)
		if err := Validate(code); err == ErrValueNotCommon {
			t.Error("Invalid number generated", code)
		}
	}
}
