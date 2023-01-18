package main

import "fmt"

//declare a struct
type Student struct {
	Name  string
	Age   int
	Grade string
	Roll  int
}

//declare a method
func (s Student) printDetails() {
	fmt.Println("Student Name:", s.Name)
	fmt.Println("Student Age:", s.Age)
	fmt.Println("Student Grade:", s.Grade)
	fmt.Println("Student Roll:", s.Roll)
}

//changing the age of a student
func (s Student) newAge() { //if we don't use pointer to (s *Student), the changes will not be reflected
	s.Age = 21
	fmt.Println("New Age:", s.Age)
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

	fmt.Println("-----------------------")

	//call the method
	fmt.Println("Calling the method printDetails()")
	akash.printDetails()
	fmt.Println("Calling the method printDetails()")
	mrinal.printDetails()

	//call the method newAge()
	fmt.Println("Calling the method newAge()")
	akash.newAge()
	akash.printDetails()
	fmt.Println("Calling the method newAge()")
	mrinal.newAge()

}
