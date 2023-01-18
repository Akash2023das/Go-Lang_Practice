package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go Lang")

	//content of a file
	content := "My Name is Akash Ruidas, I am a Computer Science Student"

	//create a file
	file, err := os.Create("17files/name.txt") //pass the path and name of the file to create
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err)

	//write to file
	length, err := io.WriteString(file, content) //pass the file and content to write
	checkNilError(err)

	//print the length of the file
	fmt.Println("Length of the file is", length)

	//close the file
	defer file.Close()

	//call the function
	readFile("17files/name.txt")
}

func readFile(filename string) {
	dataByte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file:\n", string(dataByte))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
