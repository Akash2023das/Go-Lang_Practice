/*
Write a function that takes a slice of strings and returns a new slice with all strings that are palindromes (i.e., read the same backwards as forwards).
Input: ["racecar", "level", "hello"]
Output: ["racecar", "level"]
*/

package main

import "fmt"

func isPalinDrome(str []string) []string {
	mySlice2 := []string{}
	var s string
	for i := 0; i < len(str); i++ {
		s = reverse(str[i])
		if s == str[i] {
			mySlice2 = append(mySlice2, str[i])
		}
	}
	return mySlice2

}
func reverse(s string) string {
	ans := ""
	for i := len(s) - 1; i >= 0; i-- {
		ans = ans + string(s[i])
	}
	return ans
}
func main() {
	str := []string{"racecar", "level", "hello"}

	fmt.Println(isPalinDrome(str))
}
