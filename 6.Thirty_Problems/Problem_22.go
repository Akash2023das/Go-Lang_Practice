/*
Write a function that takes a slice of integers and returns a new slice with all elements that are divisible by a given number.
Input: ([1, 2, 3, 4, 5], 2)
Output: [2, 4]
*/

package main

import "fmt"

func divisibleByTwo(numbers []int, num int) []int {
	var ans []int
	for _, v := range numbers {
		if v%2 == 0 {
			ans = append(ans, v)
		}
	}
	return ans
}
func main() {
	numbers := []int{1, 2, 3, 4, 5}

	num := 2
	fmt.Println(divisibleByTwo(numbers, num))
}
