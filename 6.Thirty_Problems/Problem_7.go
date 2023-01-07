/*
Write a function that takes a string and returns a new string with all vowels removed.
Input: "hello"
Output: "hll"
*/

package main

import (
	"fmt"
	"strings"
)

// creating a function that takes a string and returns a new string with all vowels removed from the string
func vowelRemoved(str string) string {

	// chars := []rune(str)
	// for i := 0; i < len(chars); i++ {
	// el := string(chars[i])

	for _, c := range []string{"a", "e", "i", "o", "u"} {
		str = strings.ReplaceAll(str, c, "")
	}

	return str
}

func main() {
	str := "hello"
	fmt.Println(vowelRemoved(str))
}
