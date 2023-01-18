package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://lco.dev"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Response Type: %T", response)
	defer response.Body.Close() //caller's responsibility to close the body

	dataBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(dataBytes)
	fmt.Println(content)
}
