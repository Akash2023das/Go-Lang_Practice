/*
Write a function that takes a slice of strings and returns a new slice with all strings that contain a given letter.
 Input: (["cat", "dog", "elephant", "lion"], "e")
 Output: ["elephant"]
*/

package main

import "fmt"

func containsLetter(str []string, letter string) (ans []string) {
	for _, element := range str {
		for i := 0; i < len(element); i++ {
			if letter == string(element[i]) {
				ans = append(ans, element)
				break
			}
		}
	}
	return ans

}
func main() {
	str := []string{"cat", "dog", "elephant", "lion", "eagle"}

	letter := "e"

	fmt.Println(containsLetter(str, letter))
}
