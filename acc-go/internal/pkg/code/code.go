package code

import "strings"

const separator = "."

func Level(code string) int {
	codes := strings.Split(code, separator)
	return len(codes)
}

func Parent(code string) string {
	pos := strings.LastIndex(code, separator)
	return code[:pos]
}
