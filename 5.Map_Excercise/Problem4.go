/*
4.Write a function that takes a map and a slice of strings, and returns a new slice with all the strings from the input slice that are keys in the map. For example, given the input map[string]int{"a": 1, "b": 2, "c": 3} and []string{"a", "b", "d"}, the function should return a slice []string{"a", "b"}.
*/

package main

import "fmt"

func MatchingKey(mp map[string]int, slice []string) []string {
	//creating a empty slice to hold matching keys
	ansSlice := []string{}

	//iterating through the map pointing the keys only
	for _, val := range slice {
		_, ok := mp[val]
		if ok {
			ansSlice = append(ansSlice, val)
		}

	}
	return ansSlice //returning the new slice
}
func main() {
	//creating a map
	mp := make(map[string]int)
	mp["a"] = 1
	mp["b"] = 2
	mp["c"] = 3

	//creating a slice of strings
	slice := []string{"a", "b", "d"}

	//printing the returning value of the given function
	fmt.Println(MatchingKey(mp, slice))

}
