/*
Write a function that takes a slice of strings and returns a map with the frequency of each string in the slice. For example, given the input []string{"a", "b", "c", "a", "b"}, the function should return a map { "a": 2, "b": 2, "c": 1 }.
*/

package main

import "fmt"

//creating a function that returns a map of strings as key and int as value
func FrequencyMap(slice []string) map[string]int {

	//creating  a new map initially it is empty
	hMap := make(map[string]int)

	//iterating through the slice
	for _, element := range slice {

		//ok is boolean type it check true or false
		val, ok := hMap[element]
		if ok {
			hMap[element] = val + 1
		} else {
			hMap[element] = 1
		}
	}
	return hMap

}
func main() {
	slice := []string{"a", "b", "c", "a", "b"}
	ans := FrequencyMap(slice)
	fmt.Println(ans)
}
