package main

import ("fmt")

type person struct{
	Name string
}

type num int;

//value receiver
//a copy is instance has passed but doesnt change original value
func (per *person) display(){ 
	fmt.Println(per.Name)
	per.Name = "harsha" //name updated
}


func main(){
	var p person
	p.Name = "gowtham"
	fmt.Println("original value: "+p.Name)
	p.display()
	fmt.Println("updated value: "+p.Name)//original instance doesnt change

	var n num = 10
	var u num = 300
	var m num = 500
	fmt.Printf("%v ,%T\n",n,n)
	fmt.Println(u,m)
}