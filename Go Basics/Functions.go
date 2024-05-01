package main

import "fmt"

//no parameters no return type
func hello(){
	fmt.Println("Welcome to Golang PL")
}
//having parameters no return type
func hi(name string, age int)  {
	fmt.Printf("name: %v \nage: %v",name,age);
}

//having parameters with return type
func solve(name string, s1,s2,s3 int)float32{
	total := s1+s2+s3;
	var pc float32= float32(total*(100/300))
	fmt.Println("\nhi "+name);
	return pc;
}
func main(){
	hello()
	hi("gowtham",23)
	fmt.Println(" your percentage is : %f",solve("gowtham",67,94,77))
}