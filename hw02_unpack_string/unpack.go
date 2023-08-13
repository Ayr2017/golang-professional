package hw02unpackstring

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var sb strings.Builder
	byteSlice := []rune(inputString)

	if matched, err := regexp.Match(`\d{2,}`, []byte(inputString)); matched || err != nil {
		return "", ErrInvalidString
	}

	if len(byteSlice) >0  && isDigit(byteSlice[0]) {
		return "", ErrInvalidString
	}

	for i, j := 0, 1; i < len(byteSlice); i++ {
		if isDigit(byteSlice[j]) {
			count, err := strconv.Atoi(string(byteSlice[j]))
			res := repeatCharacter(byteSlice[i], count)
			if err != nil {
				return "", ErrInvalidString
			}
			sb.WriteString(res)
		} else if !isDigit(byteSlice[i]) {
			sb.WriteString(string(byteSlice[i]))
		}

		if j < len(byteSlice)-1 {
			j++
		}
	}

	return sb.String(), nil
}

func repeatCharacter(character rune, count int) string {
	var repeatedResult string
	if count > 0 {
		str := string(character)
		repeatedResult = strings.Repeat(str, count)
	} else {
		repeatedResult = ""
	}

	return repeatedResult
}

func isDigit(ch rune) bool {
	if _, err := strconv.Atoi(string(ch)); err == nil {
		return true
	}
	return false
}
