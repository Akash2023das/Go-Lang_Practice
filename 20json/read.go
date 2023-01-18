package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type student struct {
	Name   string `json."name,omitempty"`
	Course string `json."course,omitempty"`
}

func main() {
	db := []student{}

	file, err := os.Open("20json/studentdb.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&db)
	fmt.Println(db)
}
