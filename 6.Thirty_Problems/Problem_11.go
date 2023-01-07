/*
Write a function that takes a slice of integers and returns the second-largest number in the slice.
Input: [1, 2, 3, 4, 5]
Output: 4
*/

package main

import "fmt"

//creating a function that return a integer
func secondLargestNumber(numbers []int) int {

	//taking a temporary variable initially it is zero
	temp := 0

	//looping over the slice of integer pointing the first element
	for i := 0; i < len(numbers); i++ {

		////looping over the slice of integer pointing the second element
		for j := i + 1; j < len(numbers); j++ {

			//checking if the first element is greater than the second then swap them in sorted order
			if numbers[i] > numbers[j] {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	// returning the second largest element
	return numbers[len(numbers)-2]
}

func main() {

	//slice of integers
	numbers := []int{1, 2, 3, 4, 5}

	//passing the slice to the function and printing the result
	fmt.Println(secondLargestNumber(numbers))

}
