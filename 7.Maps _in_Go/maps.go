package main

import "fmt"

func main() {
	fmt.Println("Maps in Go lang")

	languages := make(map[string]string) //creating an empty map
	languages["JS"] = "JavaScript"       //adding key,value in the map
	languages["PY"] = "Python"
	languages["RB"] = "Ruby"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS short form: ", languages["JS"])

	delete(languages, "RB") //delete from map
	fmt.Println("List of all languages: ", languages)

	//loops through map
	for key, value := range languages {
		fmt.Printf("For key: %v, value: %v\n", key, value)
	}

}
