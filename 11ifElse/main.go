package main

func main() {
	a := 10
	b := 20
	if a > b {
		println("a > b")

	} else if a == b {
		println("a == b")

	} else {
		println("a <= b")
	}
	//
	loginCount := 23
	var result string

	if loginCount > 10 {
		result = "Login count > 10"
	} else if loginCount > 20 {
		result = "Login count > 20"
	} else {
		result = "Login count <= 10"
	}
	println(result)
	//
	if no := 21; no%2 == 0 {
		println("no is even")
	} else {
		println("no is odd")
	}
	

}