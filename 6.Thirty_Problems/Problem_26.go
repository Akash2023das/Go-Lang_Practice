/*
Write a function that takes a slice of strings and returns a new slice with all strings that are palindromes (i.e., read the same backwards as forwards).
 Input: ["racecar", "level", "hello"]
 Output: ["racecar", "level"]
*/

package main

import "fmt"

func reverseString(element string) (result string) {
	for _, v := range element {
		result = string(v) + result
	}
	return
}
func main() {
	str := []string{"racecar", "level", "hello"}

	var ans []string
	for _, element := range str {
		oldString := element
		newString := reverseString(element)

		if oldString == newString {
			ans = append(ans, element)
		}
	}
	fmt.Println(ans)

}
