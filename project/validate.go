package project

import "regexp"

func ValidateName(name string) (valid bool) {
	re := regexp.MustCompile(`^[A-Za-z][A-Za-z0-9_]*$`)
	return re.MatchString(name)
}
