package main

import (
	"fmt"
	"time"
)

func main() {
	//time now
	presentTime := time.Now() //it returns the current time
	fmt.Println("Present time is: ", presentTime)

	//printing the date in a specific format
	fmt.Println(presentTime.Format("01-02-2006 Monday"))

}
