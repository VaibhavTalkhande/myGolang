package main

import "fmt"


func main() {
      fmt.Println("Maps in Golang")
      languages := make(map[string]string)
      languages["JS"] = "Javascript"
      languages["PY"] = "Python"
      languages["RB"] = "Ruby"
      languages["GO"] = "Golang"
     
      fmt.Println(languages)

      fmt.Println(languages["PY"])
      fmt.Println(languages["RB"])
      fmt.Println(languages["GO"])
      fmt.Println(languages["JS"])

      delete(languages, "JS")
      fmt.Println(languages)
      // := is used to declare and initialize a variable in golang
      // loops are interesting in golang
      for key, value := range languages {// for each key, value pair in languages
        	fmt.Printf("For key %v, value is %v\n", key, value)
      }
      // if you want to iterate over only keys
      for key := range languages {
          	fmt.Println("Key is", key)
      }
      // if you want to iterate over only values
      for _, value := range languages {
            	fmt.Println("Value is", value)
      }
          
  //
}