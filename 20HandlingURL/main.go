package main

import (
	"fmt"
	"net/url"
)

const myurl string= "https://lco.dev:3000/learn?coursename=reactjs&paymentid=1234"
func main() {

	fmt.Println("Welcome to url handling in golang")
	fmt.Println("Url is ", myurl)
    //parse url
	result,_ := url.Parse(myurl)
	qparams := result.Query()
	fmt.Println("Scheme is ", result.Scheme)//https
	fmt.Println("Host is ", result.Host)//lco.dev:3000
	fmt.Println("Path is ", result.Path)//learn
	fmt.Println("Port is ", result.Port())//3000
	fmt.Println("RawQuery is ", result.RawQuery)//coursename=reactjs&paymentid=1234
	fmt.Println("Query is ", result.Query())//map[coursename:[reactjs] paymentid:[1234]]
	fmt.Println("Fragment is ", result.Fragment)//nothing
	fmt.Println("type of query params is ", qparams)//map[string][]string
	fmt.Println("type of query params is ", qparams["coursename"])//[reactjs]

	for _,value := range qparams {
		fmt.Println("value is ", value)
	}

	partsOfUrl := &url.URL{//this is a pointer to url
		Scheme: "https",
		Host: "lco.dev",
		Path: "/tutcss",
		RawPath: "user=abc",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println("Another url is ", anotherUrl)


}