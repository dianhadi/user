package utils

import (
	"regexp"
)

var (
	alphaNumeric *regexp.Regexp
	emailAddress *regexp.Regexp
)

func init() {
	alphaNumeric = regexp.MustCompile("^[a-zA-Z0-9]+$")
	emailAddress = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
}

func IsAlphaNumeric(input string) bool {
	return alphaNumeric.MatchString(input)
}

func IsValidEmail(input string) bool {
	return emailAddress.MatchString(input)
}
