package main

import "fmt"

func main() {
	fmt.Println("If else in Go lang")

	var result string
	loginCount := 23
	if loginCount < 10 {
		result = "Regular user"
	}
	fmt.Println(result)

}
