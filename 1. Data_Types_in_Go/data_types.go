package main

import "fmt"

const LoginToken string = "ard" //this constant is public

func main() {
	var userName string = "Akash Ruidas" //declaration of a variable in go lang
	fmt.Println(userName)
	fmt.Printf("Variable is a type of: %T\n", userName)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is a type of: %T\n", isLoggedIn)

	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is a type of: %T\n", smallVal)

	var smallFloat float32 = 256.5
	fmt.Println(smallFloat)
	fmt.Printf("Variable is a type of: %T\n", smallFloat)

	//implecit type to declare a variable
	var website = "google.com"
	fmt.Println(website)

	//no var style
	numberOfUser := 3000 //not allow to use outside te main func
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is a type of: %T\n", LoginToken)

}
