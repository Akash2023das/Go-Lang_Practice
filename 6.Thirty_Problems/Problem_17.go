/*
Write a function that takes a slice of strings and returns a new slice with all strings that start with a given letter.
Input: (["cat", "dog", "elephant"], "e")
 Output: ["elephant"]
*/

package main

import "fmt"

func matchingLetter(str []string, s string) []string {

	//creating a new slice initially it is empty
	var ans []string

	//iterating over the slice of string by taking the elements directly
	for _, letter := range str {

		//iterating over the elements of the slice
		for i := 0; i < len(letter); i++ {
			//converting the first element and storing it then checking the condition if matching with the given string or not
			el := string(letter[i])
			if el == s {
				ans = append(ans, letter)
				break
			}
		}
	}
	return ans

}

func main() {
	//slice of string
	str := []string{"cat", "dog", "elephant", "eagle"}

	s := "e"

	fmt.Println(matchingLetter(str, s))
}
