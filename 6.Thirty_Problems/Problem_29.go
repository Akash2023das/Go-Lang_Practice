/*
Write a function that takes a slice of integers and returns a new slice with all elements that are palindromes (i.e., read the same forwards as backwards).
Input: [1, 11, 121, 12321]
Output: [11, 121, 12321]
*/

package main

import "fmt"

func isPalinDrome(n int) bool {
	oldNum := n
	var rem int
	var rev int
	for n != 0 {
		rem = n % 10
		rev = (10 * rev) + rem
		n = n / 10
	}
	if rev == oldNum {
		return true
	} else {
		return false
	}
}
func main() {
	numbers := []int{1, 11, 121, 12321}

	var ans []int
	for _, element := range numbers {
		if isPalinDrome(element) == true {
			ans = append(ans, element)
		}
	}
	fmt.Println(ans)

}
