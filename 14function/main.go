package main

import (
	
	"fmt"

)


func main() {

	fmt.Println("Welcome to functions in golang")
     var result int = add(2, 3)
	fmt.Println("Result is ", result)
	fmt.Println("Result is ", proAdder(2, 3, 4, 5, 6, 7, 8, 9, 10))

}

func add(a int, b int) int {
	return a + b
}

func proAdder(values ...int) int {
	finalValue := 0
	for _, value := range values {//range returns index and value
		finalValue += value
	}
	return finalValue
}
