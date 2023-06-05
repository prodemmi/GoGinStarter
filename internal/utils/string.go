package utils

import (
	"strings"
)

func SIsString(value interface{}) bool {
	_, ok := value.(string)
	return ok
}

func SFirstLetterUpper(str string) string {
	if str == "" {
		return str
	}
	firstLetter := strings.ToUpper(string(str[0]))

	return firstLetter + str[1:]
}
