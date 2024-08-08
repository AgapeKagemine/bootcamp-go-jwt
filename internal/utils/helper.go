package utils

import "strings"

func IsValid(str string) bool {
	return strings.Trim(str, " ") != "" && len(str) > 2
}
