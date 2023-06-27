package main

import (
	"fmt"
	"reflect"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")
	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Type of fruitList is: ", reflect.TypeOf(fruitList))
	fmt.Println("Fruit list length is: ", len(fruitList))
	fruitList =append(fruitList, "Mango", "Banana")
	fmt.Println("Fruit list after adding Mango and Banana is: ", fruitList)

	fruitList = append(fruitList[1:3])//This is called slicing a slice :)
	
	fmt.Println("Fruit list after slicing is: ", fruitList)

	//make function
	//make([]T, length, capacity)
	//capacity is optional and defaults to length
	//var myFruitList = make([]string, 5, 10)//5 is length and 10 is capacity
	highScores := make([]int, 4)//4 is length and capacity
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	fmt.Println("Contents of highScores: ", highScores)
	highScores = append(highScores, 555, 666, 777)
	fmt.Println("Contents of highScores after addition: ", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Contents of highScores after sorting: ", highScores)
    //how to remove a item from slice based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("Courses: ", courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)//... is used to append the slice
	fmt.Println("Courses after removing: ", courses)

}