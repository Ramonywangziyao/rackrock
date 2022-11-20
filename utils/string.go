package utils

import "strings"

func IsEmptyStr(str string) bool {
	return strings.TrimSpace(str) == ""
}
