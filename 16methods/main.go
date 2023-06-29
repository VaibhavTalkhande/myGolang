package main

import "fmt"

func main() {
	fmt.Println("Structs in Golang")
	// no inheritance in golang
	// no constructors in golang
	// no super in golang
	// no this in golang
	// no self in golang
	// no parent in golang
	// no child in golang
	// no polymorphism in golang
	// no overloading in golang
	// no overriding in golang

	hitesh := User{"Hitesh", "hitemlm09sh@go.dev",true,16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are %+v\n", hitesh)// %+v will print field names also in the output
	fmt.Printf("Name is %v and Email is %v\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail() // this will not change the email of hitesh
	fmt.Println("Email of hitesh is ", hitesh.Email)
    
}

type User struct {// struct is a collection of fields
	// fields are like variables of a struct
	// fields are like attributes of a struct

	Name string// field name is Name and type is string
	Email string
	Status bool
	Age int
}
func (u User) GetStatus() string {
	if u.Status == true {
		fmt.Println("User is active")
		return "Active"

	}
	fmt.Println("User is inactive")
	return "Inactive"

}

func (u User) NewMail() {
	u.Email = "sdfdsfdsfd@gssma.dev"
	fmt.Println("Email of this user is ", u.Email)
}

