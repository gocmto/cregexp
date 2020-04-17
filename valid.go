package cregexp

import "regexp"

const langRus = "rus"
const langEng = "eng"

//
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
