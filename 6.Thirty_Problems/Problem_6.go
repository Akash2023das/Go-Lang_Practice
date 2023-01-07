/*
Write a function that takes a slice of strings and returns a new slice with all strings that have more than 5 characters.
Input: ["cat", "dog", "elephant", "lion"]
Output: ["elephant"]
*/

package main

import "fmt"

//creating a function that returns a new slice with all strings that have more than 5 characters
func countCharacters(str []string) []string {

	//creating a empty slice of strings
	var ans []string

	//iterating over the given slice of strings by only taking the elements
	for _, element := range str {

		//checking the length of every string of elements  if more than 5 then
		if len(element) > 5 {
			ans = append(ans, element) // then append the element to the empty slice of string
		}
	}
	return ans //return the result
}
func main() {
	//slice of strings
	str := []string{"cat", "dog", "elephant", "lion", "godzilla"}

	fmt.Println(countCharacters(str))

}
