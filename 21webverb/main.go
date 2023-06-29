package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verb video")
	//PerformGETRequest()
	//PerformPostJsonRequest() 
	PerformPostFormRequest()
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

func PerformPostJsonRequest(){
	const myurl = "http://localhost:8000/post"

	//fake json data
    requestBody := strings.NewReader(`
	{
		"name":"Rahul",
		 "job":"Developer",
		 "city":"Delhi"
	}
	`)
	response,err := http.Post(myurl, "application/json", requestBody)//this will post the data to the server
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()//this will close the connection after the function is executed

	databytes,err:= ioutil.ReadAll(response.Body)//this will read the response body
	if err!=nil {
		panic(err)
	}

	content := string(databytes)
	fmt.Println("Content is ", content)

}

func PerformPostFormRequest(){
	const myurl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("name", "Rahul")
	data.Add("job", "Developer")
	data.Add("city", "Delhi")
	data.Add("age", "22")


	response,err := http.PostForm(myurl, data)//this will post the data to the server
	if err!=nil {
		panic(err)
	}
	defer response.Body.Close()//this will close the connection after the function is executed

	databytes,err:= ioutil.ReadAll(response.Body)//this will read the response body
	if err!=nil {
		panic(err)
	}
     
	content := string(databytes)
	fmt.Println("Content is ", content)

}