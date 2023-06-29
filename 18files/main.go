package main

import (
	"fmt"
	"io"        //input output
	"io/ioutil" //input output utility makes it easy to read and write files
	"os"        //operating system
)

func main() {
	fmt.Println("Welcome to functions in golang")
	content := "This is a sample text"

	file,err := os.Create("sample.txt")
	checkNilErr(err)
	length, err := io.WriteString(file, content)
    checkNilErr(err)
	fmt.Println("Length is ", length)
	defer file.Close()
	readFile("sample.txt")

}


func readFile(filname string){
	databyte, err := ioutil.ReadFile(filname)
	checkNilErr(err)
	//fmt.Println("Data inside file is ", databyte)//this will print in bytes
	fmt.Println("Text data inside file is ", string(databyte))
} 

func checkNilErr(err error){
	if err!=nil {
		panic(err)
	}
}