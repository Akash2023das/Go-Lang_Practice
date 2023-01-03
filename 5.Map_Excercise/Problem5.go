/*
5.Write a function that takes a map and returns a new map with all the keys and values reversed. For example, given the input map[string]int{"a": 1, "b": 2, "c": 3}, the function should return a new map map[int]string{1: "a", 2: "b", 3: "c"}.
*/

package main

import "fmt"

//creating a function that will return the key and values in reverse order
func ReversedMap(mp map[string]int) map[int]string {

	//creating a new map that will stores integer as a key and string as a value in pairs
	newMap := make(map[int]string)

	//iterating over the map
	for k, v := range mp {
		newMap[v] = k //storing the value as a key in the new map
	}

	//returning the new map
	return newMap

}
func main() {

	//creating a map
	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	//printing the returning value of the given function
	fmt.Println(ReversedMap(mp))

}
