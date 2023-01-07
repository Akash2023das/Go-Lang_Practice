/*
Write a function that takes a slice of integers and a number and returns the number of times the number appears in the slice.
Input: ([1, 2, 3, 2, 4, 5], 2)
 Output: 2
*/

package main

import "fmt"

//creating a function that returns a integer
func FreqCheck(numbers []int, num int) int {
	//taking a empty map
	mp := make(map[int]int)

	//adding all the numbers with its corresponding frequency
	for _, ele := range numbers {
		val, ok := mp[ele]
		if ok {
			mp[ele] = val + 1
		} else {
			mp[ele] = 1
		}
	}

	//checking the key is matching with the given number or not
	//if matching then return the value of the key
	ans := 0
	for key, v := range mp {
		if key == num {
			ans = v
		}
	}
	return ans //returns the value of the key
}
func main() {
	//slice of integers
	numbers := []int{1, 2, 3, 2, 4, 5}

	//number
	num := 2

	fmt.Println(FreqCheck(numbers, num))
}
