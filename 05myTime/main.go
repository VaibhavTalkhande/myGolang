package main

import (
	"fmt" //Package fmt implements formatted I/O with functions analogous to C's printf and scanf. The format 'verbs' are derived from C's but are simpler.

	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()//time.Now() returns the current local time.

	fmt.Println("Current time is: ", presentTime)	
	fmt.Println("Formatted time is: ", presentTime.Format("01-02-2006 15:04:05 Monday"))//Format returns a textual representation of the time value formatted according to layout, which defines the format by showing how the reference time, defined to be
     createdDate := time.Date(2018, time.May, 20, 23, 23, 0, 0, time.UTC)//Date returns the Time corresponding to yyyy-mm-dd hh:mm:ss + nsec nanoseconds in the appropriate zone for that time in the given location.
     fmt.Println("Created date is: ", createdDate.Format("01-02-2006 Monday"))	 
	
}