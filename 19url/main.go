package main

import (
	"fmt"
	"net/url"
)

const myurl string = "http://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("URL Handling in Go Lang")
	fmt.Println(myurl)

	// parse the url
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Println(qparams)
	fmt.Println(qparams["coursename"])
	fmt.Printf("The type of qparams is %T", qparams)

	//looping over the map
	for key, value := range qparams {
		fmt.Println(key, value)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitest",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)

}
