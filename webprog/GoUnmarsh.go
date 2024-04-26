package main

import (
	"encoding/json"
	"fmt"
)

 type employee struct{
	Ename string
	Age int
	Salary int
	Address string
 }

 func main(){
	var vstruct []employee

	data1 := []byte(`
	[
		{"Ename" : "Rama Rao","Age":30, "Salary":89230,"Address":"mumbai"},
		{"Ename" : "Srinivas","Age":42, "Salary":32000,"Address":"Hyd"},
		{"Ename" : "venkatesh","Age":34, "Salary":45200,"Address":"Vizag"},
		{"Ename" : "Radha","Age":39, "Salary":28930,"Address":"guntur"}
	]`)

	err := json.Unmarshal(data1,&vstruct)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vstruct)

	for i,j := range vstruct{
		fmt.Println(i,j)
	}
 }