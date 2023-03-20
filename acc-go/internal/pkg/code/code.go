package code

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

const separator = "."

func Level(code string) int {
	codes := strings.Split(code, separator)
	return len(codes)
}

func Parent(code string) string {
	pos := strings.LastIndex(code, separator)
	return code[:pos]
}

func Next(code string) (string, error) {
	if code == "" {
		return "", nil
	}
	pos := strings.LastIndex(code, separator)
	prefix := code[0 : pos+1]
	current := code[pos+1:]
	v, err := strconv.Atoi(current)
	if err != nil {
		return "", err
	}
	if v == 99 {
		return "", errors.New("exceeds maximum value")
	}
	str := strconv.Itoa(v + 1)
	return prefix + fmt.Sprintf("%0*s", 2, str), nil
}

func FirstChild(code string) string {
	return code + separator + "01"
}
