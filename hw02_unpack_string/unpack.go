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

	for i, j := 0, 1; i < len(byteSlice); i++ {
		if i == 0 && isDigit(byteSlice[i]) {
			return "", ErrInvalidString
		}

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

		if i < len(byteSlice)-2 {
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
	return ch >= '0' && ch <= '9'
}
