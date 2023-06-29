package main

import "fmt"

func main() {
	defer fmt.Println("This is defer")//last in first out means this will be executed last
	defer fmt.Println("This is defer 2")
	defer fmt.Println("This is defer 3")
	fmt.Println("Welcome to functions in golang")
	myDefer()

}

func myDefer() {
	for i:=0; i<5; i++ {
		defer fmt.Println(i)
		//4 3 2 1 0
	}
}

