package main

import (
	"encoding/json"
	"fmt"
)

type student struct{
	Sname string
	Age int
	Sub1 int
	Sub2 int
	Sub3 int
	Total int
}

func main(){
	//marshalling
	s1 := student{"gowtham",23,70,60,90,220}
	sjson,err := json.Marshal(s1)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(sjson))

	
	//array of struct type instances
	var sarr []student = []student{
		{Sname: "harsha", Age: 24, Sub1: 89, Sub2: 70,Sub3: 98,Total: 257},
		{Sname: "sarath", Age:34, Sub1: 90, Sub2: 75,Sub3: 99,Total: 278},
		{
			Sname: "mohan",
			Age: 89,
			Sub1: 80,
			Sub2: 78,
			Sub3: 53,
			Total: 88,
		},
	}
	fmt.Println(sarr)

	sarrjson, error := json.Marshal(sarr)

	if error != nil {
		fmt.Println(error)
	}

	fmt.Println(string(sarrjson))

	//unmarshling
	var data []byte = []byte(`{"Sname":"Ramesh",
	"Age":12,
	"Sub1":90,
	"Sub2":87,
	"Sub3":56,
	"Total":252
	}`)
	fmt.Println(data)

	var unmarsh student
	eror :=json.Unmarshal(data,&unmarsh)
	if eror != nil {
		fmt.Println(eror)
	}
	fmt.Print(unmarsh.Sname,unmarsh.Age,unmarsh.Sub1,unmarsh.Sub2,unmarsh.Sub3,unmarsh.Total)
}