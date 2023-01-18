package main

import "fmt"

func main() {

	defer fmt.Println("Hello") //this is executed last
	fmt.Println("World")       //this is executed first

	//calling the function below
	myDefer()
}

func myDefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
