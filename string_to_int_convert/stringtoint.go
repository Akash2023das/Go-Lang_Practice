package main

import "fmt"

func atoi(s string) int {
	num := 0
	n := len(s) //=2

	for i := 0; i < n; i++ {//i=0 0<2 ,  i=1 1<2 , i=2 2<2
		num = (num*10 + int(s[i]) - 48)//num=0+1=1 ,  num=10+2 =12
	}
	return num

}

func main() {
	s := "12"
	fmt.Println(2 * atoi(s))
}
