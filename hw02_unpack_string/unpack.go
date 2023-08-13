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
	runesSlice := []rune(inputString)

	if matched, err := regexp.Match(`\d{2,}`, []byte(inputString)); matched || err != nil {
		return "", ErrInvalidString
	}

	if len(runesSlice) > 0 && isDigit(runesSlice[0]) {
		return "", ErrInvalidString
	}

	for i, j := 0, 1; i < len(runesSlice); i++ {
		if isDigit(runesSlice[j]) {
			count, err := strconv.Atoi(string(runesSlice[j]))
			if err != nil {
				return "", ErrInvalidString
			}
			res := repeatCharacter(runesSlice[i], count)
			sb.WriteString(res)
		} else if !isDigit(runesSlice[i]) {
			sb.WriteRune(runesSlice[i])
		}

		if j < len(runesSlice)-1 {
			j++
		}
	}

	return sb.String(), nil
}

func repeatCharacter(character rune, count int) string {
	var repeatedResult string
	if count > 0 && !isDigit(character) {
		str := string(character)
		repeatedResult = strings.Repeat(str, count)
	} else {
		repeatedResult = ""
	}

	return repeatedResult
}

func isDigit(char rune) bool {
	return char >= '0' && char <= '9'
}
