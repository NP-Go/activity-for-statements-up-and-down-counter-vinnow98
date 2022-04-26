package main

import "fmt"

func main() {
	var first int
	var second int
	//var tempVal int
	fmt.Print("Enter first ")
	fmt.Scanln(&first)
	fmt.Print("Enter second ")
	fmt.Scanln(&second)

	if second < first {
		tempVal := second
		second = first
		first = tempVal
	}
	for i := first; i < second; i++ {
		fmt.Println(i)
	}
	for i := second; i >= first; i-- {
		fmt.Println(i)
	}
}
