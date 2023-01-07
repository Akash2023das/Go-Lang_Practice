/*
Write a function that takes a string and returns a new string with all the words in reverse order.
Input: "This is a test"
Output: "test a is This"
*/
package main

import (
	"fmt"
	"strings"
)

func reverseString(str string) string {
	var ans []string
	parts := strings.Split(str, " ")
	for i := len(parts) - 1; i >= 0; i-- {
		if len(parts[i]) == 0 {
			continue
		}
		ans = append(ans, parts[i])
	}
	return strings.Join(ans, " ")
}
func main() {
	str := "This is a test"

	fmt.Println(reverseString(str))
}
