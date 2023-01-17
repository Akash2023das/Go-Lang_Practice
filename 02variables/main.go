package main

import "fmt"

//constants in Go lang

const pi = 3.14

func main() {
	//string
	var username string = "Akash Ruidas"
	fmt.Println(username)
	fmt.Printf("username is a type of %T\n", username)

	//bool
	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("isLoggedIn is a type of %T\n", isLoggedIn)

	//int
	var smallValue uint16 = 256
	fmt.Println(smallValue)
	fmt.Printf("smallValue is a type of %T\n", smallValue)

	//float
	var valueFloat float64 = 34.5
	fmt.Println(valueFloat)
	fmt.Printf("valueFloat is a type of %T\n", valueFloat)

	//defalut values and some alias
	var defalutValue int
	fmt.Println(defalutValue)

	//implicit type
	var name = "www.google.com"
	fmt.Println(name)

	//no var style
	age := 20 //this type of declaration is only possible inside a function
	fmt.Println(age)
	fmt.Printf("age is a type of %T\n", age)

	//printing the constant
	fmt.Printf("pi is = %T\n", pi)

}
