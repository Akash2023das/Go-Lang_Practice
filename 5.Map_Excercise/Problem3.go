/*
3.Write a function that takes a map and returns a new map with all the values multiplied by 2. For example, given the input map[string]int{"a": 1, "b": 2, "c": 3}, the function should return a new map map[string]int{"a": 2, "b": 4, "c": 6}.
*/

package main

import "fmt"

//creating-a function that will returns a map
func Multiply2(a map[string]int) map[string]int {

	//taking a  empty map
	//ans := make(map[string]int)

	//iterating over the given map 'a'
	for key, val := range a {
		a[key] = val * 2 //storing the key into the new map and also the value multiply by 2
	}
	return a// returning the  new map 'ans'  that we have created
}

func main() {

	//creating a map that will take string and integer in key, value pairs
	a := make(map[string]int)

	//putting the key , values into the map 'a'
	a["a"] = 1
	a["b"] = 2
	a["c"] = 3

	//printing the returning value of the given function
	fmt.Println(Multiply2(a))
}
