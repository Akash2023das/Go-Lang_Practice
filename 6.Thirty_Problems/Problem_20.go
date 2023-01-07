/*
Write a function that takes a slice of integers and returns a new slice with all the prime numbers.
Input: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10]
 Output: [2, 3, 5, 7]
*/

package main

import "fmt"

func isPrime(n int) bool {
	var flg int
	if n == 0 || n == 1 {
		flg = 1
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			flg = 1
			break
		}
	}
	if flg == 0 {
		return true
	}
	return false

}
func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	var ans []int
	for _, v := range numbers {
		if isPrime(v) {
			ans = append(ans, v)
		}
	}
	fmt.Println(ans)
}
