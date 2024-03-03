package esr

import "errors"

var ErrValueNotCommon = errors.New("esr: must be valid number by Common algorithm")

// Validate check number is valid or not based on Common algorithm
func Validate(value interface{}) error {
	s, _ := value.(string)
	l := len(s)
	if l < 1 {
		return ErrValueNotCommon
	}
	cDigit := checkDigitCalculate(s)
	checkDigit := int(s[l-1] - 48)
	result := cDigit == checkDigit
	if !result {
		return ErrValueNotCommon
	}
	return nil
}

// Generate return the number with check digit
func Generate(code string) string {
	checkDigit := checkDigitCalculate(code + "0")
	return code + string(rune(checkDigit+48))
}

// Digit generate return the check digit
func Digit(code string) int {
	return checkDigitCalculate(code + "0")
}

func checkDigitCalculate(code string) int {
	w1 := []int{1, 2, 3, 4, 5}
	w2 := []int{3, 4, 5, 6, 7}
	sum, _ := checkSum(code, w1)
	cDigit := sum % 11
	if cDigit == 10 {
		sum, _ = checkSum(code, w2)
		cDigit = sum % 11
		if cDigit == 10 {
			cDigit = 0
		}
	}

	return cDigit
}

func checkSum(code string, w []int) (int, error) {
	sum := 0
	strArr := []rune(code)
	nDigits := len(code)

	for i := 0; i < nDigits-1; i++ {
		digit := int(strArr[i]-48) * w[i]
		sum += digit
	}

	return sum, nil
}
