/*
Write a function that takes a string and returns a new string with all the words in reverse order, with the letters in each word reversed as well.
Input: "This is a test"
 Output: "tset a si sihT"
*/

package main

import "fmt"

func reverseEveryLetter(input string) string {
	ans := ""
	for i := len(input) - 1; i >= 0; i-- {
		ans = ans + string(input[i])
	}
	return ans
}
func main() {
	str := "This is a test"

	fmt.Println(reverseEveryLetter(str))
}
