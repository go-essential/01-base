package main

import "fmt"

func main() {
	var number int
	number = 4

	var number1 int = 5
	var (
		number2 int = 6
		number3 int = 7
	)
	
	number4 := 8
	number5, number6 := 9, 10

	fmt.Println((number))
	fmt.Println(number1)
	fmt.Println(number2)
	fmt.Println(number3)
	fmt.Println(number4)
	fmt.Println(number5)
	fmt.Println(number6)
	
}