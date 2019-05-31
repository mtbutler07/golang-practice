package main

import "fmt"

// Function to add two number together and return their sum
func add(x, y float64) float64 {
	return x + y
}

func main() {

	var num1 = 5.6453
	var num2 = 34.3466
	fmt.Println(add(num1, num2))
}
