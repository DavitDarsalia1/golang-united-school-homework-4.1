package reverse_string

import "fmt"

func ReverseString(input string) (output string) {
	runes := []rune(input)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	output = string(runes)
	fmt.Println(output)

	return
}
