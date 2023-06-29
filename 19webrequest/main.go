package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//lco request
const url = "https://lco.dev"

func main(){
	fmt.Println("Welcome to web requests in golang")
	
	response,err := http.Get(url)
	if err!=nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n", response)
	defer response.Body.Close()//this will close the connection after the function is executed
	databytes,err:= ioutil.ReadAll(response.Body)//this will read the response body
	if err!=nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println("Content is ", content)
}
