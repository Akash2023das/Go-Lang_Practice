/*
Write a function that takes a slice of strings and returns a new slice with all strings sorted in alphabetical order.
Input: ["zebra", "monkey", "aardvark"]
Output: ["aardvark", "monkey", "zebra"]
*/

package main

import (
	"fmt"
	"sort"
)

// creating a function that returns the slice in alphabetical order
func AlphabeticalOrder(str []string) []string {

	//using inbuilt sort function
	sort.Strings(str)

	return str
}

func main() {

	//slice of string
	str := []string{"zebra", "monkey", "aardvark"}

	//passing the slice through the function and printing the result
	fmt.Println(AlphabeticalOrder(str))
}
