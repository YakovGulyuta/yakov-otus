package hw02unpackstring

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var b strings.Builder

	unicode := utf8.RuneCountInString(input)
	length := len(input)

	for i := 0; i < len(input); i++ {

		atoiInt, _ := strconv.Atoi(string(input[i]))
		_, err := strconv.Atoi(string(input[0]))
		if err == nil {
			return "", ErrInvalidString
		}
		if atoiInt == 0 {
			b.WriteString(string(input[i]))
		} else {
			b.WriteString(strings.Repeat(string(input[i-1]), atoiInt-1))
		}
		//fmt.Println(prev, atoiInt)
		//if (prev % 2) == 0 {
		//	b.WriteString(string(input[i]))
		//}

	}

	result := b.String()
	b.Reset()

	fmt.Println(result, unicode, length)

	//strings.Repeat
	//strconv.Atoi
	return result, nil
}
