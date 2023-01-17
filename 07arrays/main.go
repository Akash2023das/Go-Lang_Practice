package main

import "fmt"

func main() {
	//arrays in Go Lang
	fmt.Println("Arrays in Go Lang")

	//declaring an array
	var fruit [5]string
	fruit[0] = "Apple"
	fruit[1] = "Orange"
	fruit[2] = "Banana"
	fruit[3] = "Grapes"
	fruit[4] = "Pineapple"

	//printing the array
	fmt.Println("Fruit Array is ", fruit)
	//size of the array
	fmt.Println("Size of the array is :", len(fruit))

	//declaring and initializing an array
	vegetables := [5]string{"Carrot", "Potato", "Tomato", "Onion", "Cabbage"}
	fmt.Println("Vegetables Array is :", vegetables)
	fmt.Println("Size of the array is :", len(vegetables))
}
