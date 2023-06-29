package main

// Model for the course - file
type Course struct {
    CourseId string `json:"courseid"`
	Coursename string `json:"coursename"`
	CoursePrice int `json:"price"`
	Author *Author `json:"author"`//this will change the name of the field in the json

}

type Author struct {
     Fullname string `json:"fullname"`
	 Website string `json:"website"`
}

//fake db
var courses []Course


//middleware
func (c *Course) IsEmpty() bool {
	return c.CourseId == "" && c.Coursename == "" && c.CoursePrice == 0 && c.Author == nil
	
}
func main() {
   
}