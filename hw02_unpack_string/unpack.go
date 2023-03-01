package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder

	//unicode := utf8.RuneCountInString(input)
	length := len(input)

	for i := 0; i < length; i++ {

		atoiInt, err := strconv.Atoi(string(input[i]))

		if err != nil {
			b.WriteString(string(input[i]))
		} else {
			if atoiInt != 0 {
				b.WriteString(strings.Repeat(string(input[i-1]), atoiInt-1))
			}
		}
	}

	result := b.String()
	b.Reset()
	return result, nil
}
