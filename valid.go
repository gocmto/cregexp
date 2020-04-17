package cregexp

import "regexp"

const langRus = "rus"
const langEng = "eng"

// Only numbers are allowed in a string
// ^[0-9]+$
func OnlyNumbers(someString string) bool {
	var valid = regexp.MustCompile(`^[0-9]+$`)
	return valid.MatchString(someString)
}

// Only latin characters are allowed in a string
// ^[a-zA-Z]+$
func OnlyEngCharacters(someString string) bool {
	var valid = regexp.MustCompile(`^[a-zA-Z]+$`)
	return valid.MatchString(someString)
}

// Only cyrillic characters allowed in the string
// ^[а-яА-ЯёЁ]+$
func OnlyRusCharacters(someString string) bool {
	var valid = regexp.MustCompile(`^[а-яА-ЯёЁ]+$`)
	return valid.MatchString(someString)
}
