package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")
    // var ptr *int 
	// fmt.Println("Value of ptr is: ", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("Value of myNumber is: ", ptr)
	fmt.Println("Address of myNumber is: ", &myNumber)
	fmt.Println("Value of myNumber is: ", *ptr)
	*ptr = *ptr *2
	fmt.Println("Value of myNumber is: ", myNumber)
}