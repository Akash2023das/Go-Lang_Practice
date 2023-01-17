package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app:")
	fmt.Println("Please give a rating for our pizza between 1 and 5: ")

	//taking user input
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	//converting the input to a float64
	rating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating : ", rating+1)
	}

}
