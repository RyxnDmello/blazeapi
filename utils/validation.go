package utils

import (
	"regexp"
	"strconv"
	"strings"
)

const symbols = `[!@#\$%\^&\*\(\)\+\-=\{\}\[\]:;"'<>,\./\?\\|` + "`~" + `]`

func ValidateIdentifier(name string) (success bool) {
	if len(name) == 0 {
		return true
	}

	if isDigit(name[0:1]) {
		return false
	}

	if isDigit(name[len(name)-1:]) {
		return false
	}

	if strings.Contains(name, " ") {
		return false
	}

	return !regexp.MustCompile(symbols).MatchString(name)
}

func isDigit(letter string) (isDigit bool) {
	_, err := strconv.ParseInt(letter, 10, 64)
	return err == nil
}
