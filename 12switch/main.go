package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Println("Enter a number between 1-10: ")
	// input, _ := reader.ReadString('\n')
	// fmt.Println("Thanks for number ", input)
	// conversion, _ := strconv.ParseFloat(strings.TrimSpace(input), 64)

	// fmt.Printf("Type of this number is %T", conversion)
	// fmt.Println("Conversion to int ", conversion)

	// //generate a random number between 0 and 1 and multiply it by the number entered by the user
	// result := (rand.Float64() * conversion )
	// fmt.Println("Result is ", result)

	// //use a switch statement with a more descriptive label
	// prize := ""
	// switch  {
	// case result>9:
	// 	prize = "nothing"
	// case result>8:
	// 	prize = "car"
	// case result>6:
	// 	prize = "bike"
	// case result>3:
	// 	prize = "TV"
	// case result>0:
	// 	prize = "You won a pen"
	// default:
	// 	prize = "nothing"
	// }

	// fmt.Println("You won a", prize)
   fmt.Println("switch and case in golang")
   fmt.Println("-------------------------")
    rand.Seed(time.Now().UnixNano())

   
   dice := rand.Intn(6) + 1
   fmt.Println("Value of dice is ", dice)
   switch dice {
	   case 1:
		   fmt.Println("Dice value is 1 and you can open")

	   case 2:
		   fmt.Println("You can move 2 spaces")
		case 3:
			fmt.Println("You can move 3 spaces")
		case 4:
			fmt.Println("You can move 4 spaces")
		case 5:
			fmt.Println("You can move 5 spaces")
		case 6:
			fmt.Println("You can move 6 spaces and roll again")
		default:
			fmt.Println("What was that?")
   
}
}