package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	PerformGETRequest()
}



func PerformGETRequest(){
	const myurl = "http://localhost:8000/get"

	response,err := http.Get(myurl)
	if err!=nil {
		panic(err)
	}

	defer response.Body.Close()//this will close the connection after the function is executed
	fmt.Println("Status code is ", response.StatusCode)
	fmt.Println("Content length is ", response.ContentLength)
    var responseString strings.Builder
	databytes,err:= ioutil.ReadAll(response.Body)//this will read the response body
	if err!=nil {
		panic(err)
	}
	bytecount,_ :=responseString.Write(databytes)
	fmt.Println("Number of bytes read are ", bytecount)
	fmt.Println("Content is ", responseString.String())
	// content := string(databytes)
	// fmt.Println("Content is ", content)
}