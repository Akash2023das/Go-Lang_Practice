/*
Write a function that takes a string and returns a map with the frequency count of each character in the string.
Input: "hello"
Output: {"h": 1, "e": 1, "l": 2, "o": 1}
*/

package main

import "fmt"

//creating a function that returns a map with the frequency of every element of the string
func freqMapCount(str string) map[string]int {
	//creating en empty map
	ans := make(map[string]int)

	/*
		A rune represents a Unicode Point. By converting a string to rune array basically it is same as creating an array of Unicode Points of that string. Therefore once the string is converted into the rune array, it can be used to iterate over all characters of the string.
	*/
	chars := []rune(str)

	//iterating over the chars array every element one by one
	for i := 0; i < len(chars); i++ {

		//storing the elements one by one
		el := string(chars[i])

		val, ok := ans[el] //checking if the element is present in the map or not
		if ok {            //if the element is already present then we increasing the val by +1
			ans[el] = val + 1
		} else { // if not present then we add the value 1
			ans[el] = 1
		}
	}
	return ans //return the map that we have created

}

func main() {
	//creating a variable that stores a string
	str := "hello"

	fmt.Println(freqMapCount(str))
}
