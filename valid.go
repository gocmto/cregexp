package cregexp

import "regexp"

const langRus = "rus"
const langEng = "eng"

// Checking the correctness of the telephone number
// Pattern: ^(\s*)?(\+)?([- _():=+]?\d[- _():=+]?){10,14}(\s*)?$
// Examples: +7(903)888-88-88, +79161234567, 8(999)99-999-99, +380(67)777-7-777
func CheckPhone(someString string) bool {
	var valid = regexp.MustCompile(`^(\s*)?(\+)?([- _():=+]?\d[- _():=+]?){10,14}(\s*)?$`)
	return valid.MatchString(someString)
}

// Checking a Credit Card. 16 numbers. With spaces.
// Pattern: (\d{4}\s([-]|)\d{4}\s([-]|)\d{4}\s([-]|)\d{4})
func Card16Space(someString string) bool {
	var valid = regexp.MustCompile(`(\d{4}\s([-]|)\d{4}\s([-]|)\d{4}\s([-]|)\d{4})`)
	return valid.MatchString(someString)
}

// Checking a Credit Card. 16 numbers. No spaces.
// Pattern: (\d{4}([-]|)\d{4}([-]|)\d{4}([-]|)\d{4})
func Card16NoSpace(someString string) bool {
	var valid = regexp.MustCompile(`(\d{4}([-]|)\d{4}([-]|)\d{4}([-]|)\d{4})`)
	return valid.MatchString(someString)
}

// Checking the correctness of the domain
// Pattern: ^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$
func CheckDomain(someString string) bool {
	var valid = regexp.MustCompile(`^([a-zA-Z0-9]([a-zA-Z0-9\-]{0,61}[a-zA-Z0-9])?\.)+[a-zA-Z]{2,6}$`)
	return valid.MatchString(someString)
}

// Checking the correctness of the email address
// Pattern: ^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$
func CheckEmail(someString string) bool {
	var valid = regexp.MustCompile(`^([a-z0-9_-]+\.)*[a-z0-9_-]+@[a-z0-9_-]+(\.[a-z0-9_-]+)*\.[a-z]{2,6}$`)
	return valid.MatchString(someString)
}

// Only numbers are allowed in a string
// Pattern: ^[0-9]+$
func OnlyNumbers(someString string) bool {
	var valid = regexp.MustCompile(`^[0-9]+$`)
	return valid.MatchString(someString)
}

// Only latin characters are allowed in a string
// Pattern: ^[a-zA-Z]+$
func OnlyEngCharacters(someString string) bool {
	var valid = regexp.MustCompile(`^[a-zA-Z]+$`)
	return valid.MatchString(someString)
}

// Only cyrillic characters allowed in the string
// Pattern: ^[а-яА-ЯёЁ]+$
func OnlyRusCharacters(someString string) bool {
	var valid = regexp.MustCompile(`^[а-яА-ЯёЁ]+$`)
	return valid.MatchString(someString)
}
