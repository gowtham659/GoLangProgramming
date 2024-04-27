package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//users struct with a memeber array of type user to hold multiple json array of user objects
type users struct{
	U []user `json:"users"`
}

//user struct with members based on name/value parirs of json objects
type user struct{
	Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Social social `json:"social"`
}

//social struct which contains a
// list of links
type social struct{
	Facebook string `json:"facebook"`
    Twitter  string `json:"twitter"`
}

func main(){

	//ReadFile() reads json data from file and returns a byte array 
	content, err := ioutil.ReadFile("webprog/users.json")

	if err != nil{
		fmt.Println(err)
	}

	//fmt.Println(string(content))

	var us users
	//unmarshal our byteArray which contains our
	//jsonFile's content into 'users' which we defined above
	err1 := json.Unmarshal(content,&us)
	if err1 != nil{
		fmt.Println(err1)
	}

	//fmt.Println(us.U[0].Name)
	//fmt.Println(us.U[1])
	//iterate through every user within our users array and
	//print out the user Type, their name, and their facebook url
	//as just an example
	for _,j:= range us.U{
		fmt.Printf("Name: %v\n",j.Name)
		fmt.Printf("Age: %v\n",j.Age)
		fmt.Printf("Type: %v\n",j.Type)
		fmt.Printf("FaceBook: %v\n",j.Social.Facebook)
		fmt.Printf("Twitter: %v\n",j.Social.Twitter)

	}

}