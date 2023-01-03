package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Handling in Go Lang")
	presentTime := time.Now()
	fmt.Println("Today date: ", presentTime)

	//change the date in this format
	fmt.Println(presentTime.Format("01-02-2003 Monday"))

}
