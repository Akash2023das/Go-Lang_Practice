/*
Write a function that takes a slice of strings and a string, and returns the index of the first occurrence of the string in the slice.
Input: (["cat", "dog", "elephant"], "dog")
Output: 1
*/

package main

import "fmt"

//creating a function that return a integer
func indexOccur(str []string, s string) int {

	//iterating over the slice with key and value of the slice
	for k, v := range str {

		//checking if the value is matching with the given string
		if v == s {
			return k
		}
	}
	return -1 //if not found return -1
}
func main() {
	//slice of string
	str := []string{"cat", "dog", "elephant"}

	s := "dog"

	fmt.Println(indexOccur(str, s))
}
