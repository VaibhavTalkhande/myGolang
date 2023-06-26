package main

import "fmt"

const LoginToken string = "asdasdasdasdasdasdasdasdasd"//public

func main() {
	var username string = "vaibhav"
	fmt.Println("Hello, ", username)
	fmt.Printf("Variable is of type: %T \n", username)
	var isLoggedin bool = true
	fmt.Println("Logged in status: ", isLoggedin)
	var smallVal uint8 = 255
	fmt.Println("Max value of uint8: ", smallVal)
	var smallFloat float32 = 255.455
	fmt.Println("Small float: ", smallFloat)


	// Default values and some aliases
	var anotherVariable int
	fmt.Println("Default value of int: ", anotherVariable)
	//implicit type
	var website = "learncodeonline.in"
	fmt.Println("Website: ", website)
	// no var style
	numberOfUsers := 300000.00
	fmt.Println("Number of users: ", numberOfUsers)
	fmt.Println("Login Token: ", LoginToken)
	//%T is used to print the type of the variable
	fmt.Printf("Variable is of type: %T \n", numberOfUsers)
}