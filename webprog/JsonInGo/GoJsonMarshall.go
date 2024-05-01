package main

import ("fmt"
		"encoding/json"
	)

type details struct{
	NAME string 
}
func main(){

	s := `{"name":"gowtham"}`
	fmt.Println(s)

	var op details;
	er := json.Unmarshal([]byte(s),&op)
	if er!=nil{
		fmt.Println(er)
	}
	fmt.Printf("Name : %v",op.NAME)
}