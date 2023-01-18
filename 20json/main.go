package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type student struct {
	Name   string `json."name,omitempty"`
	Course string `json."course,omitempty"`
}

func main() {
	//json format data
	// db := []student{
	// 	{"Akash", "Golang"},
	// 	{"Rahul", "Python"},
	// 	{"Raj", "Java"},
	// 	{"Ravi", "C++"},
	// 	{"Rajesh", "C#"},
	// 	{"Rakesh", "C"},
	// 	{"Ramesh", "Ruby"},
	// 	{"Raju", "PHP"},
	// }

	//ask user to enter student  data and store it in slice db
	db := make([]student, 0, 3)
	for {
		var s student
		fmt.Println("Enter student name: ")
		fmt.Scanln(&s.Name)
		fmt.Println("Enter student course: ")
		fmt.Scanln(&s.Course)
		db = append(db, s)
		fmt.Println("Do you want to add more student data? (y/n)")
		var ans string
		fmt.Scanln(&ans)
		if ans == "n" {
			break
		}

	}

	//store data in memory
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(db) //db is my data

	//creating a file
	file, err := os.Create("20json/studentdb.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(file, buf)
}
