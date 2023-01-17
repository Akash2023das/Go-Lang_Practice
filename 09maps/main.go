package main

import "fmt"

func main() {
	fmt.Println("Maps in Go Lang")

	//declare a empty map
	courses := make(map[string]string)
	courses["Android"] = "Android Development"
	courses["iOS"] = "iOS Development"
	courses["Web"] = "Web Development"
	courses["Python"] = "Python Development"
	courses["Go"] = "Go Development"
	courses["Java"] = "Java Development"

	//print the map
	fmt.Println("List of the Courses:", courses)
	fmt.Println("full course name:", courses["Go"])

	//delete a map element
	delete(courses, "iOS")

	//print the map
	fmt.Println("List of the Courses:", courses)

	//loop through the map
	for key, value := range courses {
		fmt.Printf("For key %v value is %v\n", key, value)
	}

}
