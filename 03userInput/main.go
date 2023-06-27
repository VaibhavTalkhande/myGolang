package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
    //reader is a variable of type *Reader which is a struct
	//NewReader returns a new Reader whose buffer has the default size.
	//bufio is a package which implements buffered I/O. It wraps an io.Reader or io.Writer object, creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
	reader := bufio.NewReader(os.Stdin)//os.Stdin is a file type which implements the io.Reader interface and reads from standard input. 
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok syntax to check if there is an error or not 
	// ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.
	//ReadString returns the string read from the buffer. If ReadString encounters an error before finding a delimiter, it returns the data read before the error and the error itself (often io.EOF). ReadString returns err != nil if and only if the returned data does not end in delim.
	//ReadString returns err != nil if and only if the returned data does not end in delim.
	input, _ := reader.ReadString('\n')//ReadString reads until the first occurrence of delim in the input, returning a string containing the data up to and including the delimiter.
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
	
}