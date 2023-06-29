package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	fmt.Println("--------------------------------")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {//classic for loop
	// 	fmt.Println(days[d])
	// }
//    for i:= range days{//range returns index
// 	   fmt.Println(days[i])
//    }
    // for index, day := range days{	//range returns index and value
	// 	fmt.Printf("Index %v, Day %v\n", index, day)
	// }
	
	// rougueValue := 1
	// for rougueValue < 10 {//while loop
	// 	fmt.Println("Value of rougueValue is ", rougueValue)
	// 	if rougueValue == 2 {
	// 		goto lco
	// 	}
	// 	rougueValue++
	// 	if rougueValue == 5 {
	// 		rougueValue++
	// 		continue
	// 	}
	// }

	// lco:
	//   fmt.Println("Welcome to lco")
	

}