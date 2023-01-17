package main

import "fmt"

func main() {
	fmt.Println("Structs in Go lang")

	akash := User{"Akash", "akash@gmail.com", true, 21}
	fmt.Println(akash)
	fmt.Printf("Akash details are: %+v\n", akash)
	fmt.Printf("Name is: %v and Email is: %v", akash.Name, akash.Email)
}

//define Structures in Go lang
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
