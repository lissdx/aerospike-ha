package utils

import (
	"regexp"
	"strings"
)

func IsEmptyString(str string) bool {
	return len(strings.TrimSpace(str)) == 0
}

func IsAlphaNumericString(str string) bool {
	r, _ := regexp.Compile("^.*[a-zA-Z0-9_].*$")
	return r.MatchString(str)
}

func IsNumericString(str string) bool {
	r, _ := regexp.Compile("^.*[0-9].*$")
	return r.MatchString(str)
}

func NormalizeStringToLower(s string) string {
	return strings.ToLower(NormalizeString(s))
}

func NormalizeStringToUpper(s string) string {
	return strings.ToUpper(NormalizeString(s))
}

// NormalizeString  returns a slice of the string s, with all leading
// and trailing white space removed, as defined by Unicode.
func NormalizeString(s string) string {
	return strings.TrimSpace(s)
}

// StrOrDefault returns firs found Not Empty string
// if all args are "empty string" the empty string will be returned
func StrOrDefault(args ...string) string {
	for i := 0; i < len(args); i++ {
		if !IsEmptyString(args[i]) {
			return args[i]
		}
	}
	return ""
}
