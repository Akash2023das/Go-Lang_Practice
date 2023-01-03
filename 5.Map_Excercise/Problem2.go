/* 2.Write a function that takes a map and returns a slice of the keys in the map, sorted in alphabetical order. For example, given the input map[string]int{"a": 1, "c": 2, "b": 3}, the function should return a slice []string{"a", "b", "c"}.
 */

package main

import (
	"fmt"
	"sort"
)

// creating a function that returns a slice of strings
func AlphabaticalOrder(a map[string]int) []string {

	//creating a new slice of strings with the size of the  given map length
	mk := make([]string, len(a)) //          k=a            mk : //c b a 
	                             //i=0  i=1  i=2  i=3

	//iterating through the map
	i := 0
	for k, _ := range a {
		mk[i] = k //add the key to the slice with the corresponing indexes
		i++
	}

	//now sorting the slice of string

	sort.Strings(mk) //a,b,c
	return mk
}
func main() {

	//creating a map
	a := make(map[string]int)
	a["a"] = 1 //adding the key,values in the map
	a["c"] = 2
	a["b"] = 3

	//printing the returning value of the given function
	ans := AlphabaticalOrder(a)
	fmt.Println(ans)

}
