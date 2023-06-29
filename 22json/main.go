package main

import (
	"encoding/json"
	"fmt"
)
type course struct {
	Name string `json:"coursename"`//this will change the name of the field in the json
	Price int `json:"price"`//this will change the name of the field in the json
	Platform string `json:"website"`//this will change the name of the field in the json 
	Password string `json:"-"`//this will not be shown in the json
	Tags []string `json:"tags",omitempty`//this will not be shown in the json if it is empty
	// Tags []string `json:"tags"`

}
func main() {
	fmt.Println("Welcome to json in golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson(){
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Youtube", "abc", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Youtube", "a1bc", []string{"web-dev", "js", "react"}},
		{"Angular Bootcamp", 299, "Youtube", "23abc",nil},

	}

	//package json

	//Marshal
	finalJson,err := json.MarshalIndent(lcoCourses,"lco","\t")//this will convert the struct to json
	if err!=nil {
		panic(err)
	}
	fmt.Printf("%s\n",finalJson)
}

func DecodeJson(){
    jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"price": 299,
		"website": "Youtube",
		"tags": [
			"web-dev",
			"js"
		]
	}
	`)
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)
	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)//this will convert the json to struct
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("Json is not valid")
	}

	// some cases where we don't know the structure of the json
	var unknownData map[string]interface{}//this will create a map of string and interface(unknown type) 
	json.Unmarshal(jsonDataFromWeb, &unknownData)//this will convert the json to map
	fmt.Printf("%#v\n", unknownData)
    for k,v := range unknownData {
		fmt.Printf("Key is %v and value is %v and type is %T\n", k,v,v)

	}

}