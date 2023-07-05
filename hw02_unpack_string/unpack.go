package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func repeat(v rune, r []rune) strings.Builder {
	var result strings.Builder
	cnt, _ := strconv.Atoi(string(r))
	for i := 1; i < cnt; i++ {
		result.WriteRune(v)
	}
	return result
}

func remove(a []rune, i int) []rune {
	copy(a[i:], a[i+1:])
	return a[:len(a)-1]
}

func Unpack(s string) (string, error) {
	var result strings.Builder
	var resultStr string
	v := []rune(s)
	index := 0

	if utf8.RuneCountInString(s) == 0 {
		return "", nil
	}
	if unicode.IsDigit(v[0]) {
		return "", ErrInvalidString
	}
	for i := 0; i < len(v); i++ {
		if i > 0 {
			if unicode.IsDigit(v[i-1]) && unicode.IsDigit(v[i]) {
				return "", ErrInvalidString
			}
		}
	}
	for i := 0; i < len(v); i++ {
		r := v[i]
		if r == rune('0') {
			v = remove(v, i-1)
		}
	}
	for i := 0; i < len(v); i++ {
		r := v[i]
		if unicode.IsDigit(r) && index == 0 {
			index = i
		}
		if !unicode.IsDigit(r) && index > 0 {
			s := repeat(v[index-1], v[index:i])
			result.WriteString(s.String())
			index = 0
		}
		if i == (len(v)-1) && index > 0 {
			s := repeat(v[index-1], v[index:i+1])
			result.WriteString(s.String())
			index = 0
		}

		if index > 0 {
			continue
		}
		result.WriteRune(r)
	}
	resultStr = result.String()
	if unicode.IsDigit(v[len(v)-1]) {
		resultStr = resultStr[:len(resultStr)-1]
	}
	return resultStr, nil
}
