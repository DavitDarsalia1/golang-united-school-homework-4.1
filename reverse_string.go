package reverse_string

import "strings"

func ReverseString(input string) (output string) {
	var str []string

	for i := len(input) - 1; i >= 0; i-- {

		str = append(str, string(input[i]))
	}

	output = strings.Join(str, "")

	return
}
