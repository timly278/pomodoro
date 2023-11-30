package util

import "strings"

func IsErrorContain(err error, subString string) bool {
	return strings.Contains(err.Error(), subString)
}
