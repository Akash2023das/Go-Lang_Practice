/*
Write a function that takes a slice of strings and returns a new slice with all strings that are at least a given length.
Input: (["cat", "dog", "elephant", "lion"], 5)
Output: ["elephant"]
*/

package main

import "fmt"

func matchingLength(str []string, ln int) []string {
	var ans []string
	for _, element := range str {
		if len(element) > ln {
			ans = append(ans, element)
		}
	}
	return ans
}
func main() {
	str := []string{"cat", "dog", "elephant", "lion"}

	ln := 5

	fmt.Println(matchingLength(str, ln))
}
