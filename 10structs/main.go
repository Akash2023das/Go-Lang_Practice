package main

import "fmt"

//declare a struct
type Student struct {
	Name  string
	Age   int
	Grade string
	Roll  int
}

func main() {
	fmt.Println("Structs in Go Lang")

	//student 1
	akash := Student{"Akash", 20, "A", 1}
	fmt.Println("Student Name:", akash.Name)
	fmt.Println("Student Age:", akash.Age)
	fmt.Println("Student Grade:", akash.Grade)
	fmt.Println("Student Roll:", akash.Roll)

	fmt.Printf("Details of student 1: %+v\n", akash)

	fmt.Println("-----------------------")
	//student 2
	mrinal := Student{Name: "Mrinal", Age: 20, Grade: "A", Roll: 2}
	fmt.Println("Student Name:", mrinal.Name)
	fmt.Println("Student Age:", mrinal.Age)
	fmt.Println("Student Grade:", mrinal.Grade)
	fmt.Println("Student Roll:", mrinal.Roll)

	fmt.Printf("Details of student 2: %+v\n", mrinal)

}
