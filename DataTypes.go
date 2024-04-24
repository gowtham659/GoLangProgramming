//Data types
//specifies the type and size of variable value
//statically typed means type checking is perfomed at compile-time and  once a variable type is defined, it can only store data of that type.
//built-in data types: int ,float32, float64, string

package main

import "fmt"

func main(){
	//for Numeric values : 	
	//two types of integer:
	//	signed : store both positive and negative values
	//	unsigned : store only positive values
	//default type for integer is int
	//signed integer:
	// 		declared with type int can store both + and -
	//	size:  dependends on platform. 
	//			4 bytes -> 32-bit
	//			8 byte -> 64-bit

	var a int = 10;
	var b int;
	b =a+1;
	fmt.Println(b,a/b,b/a);
	/*
	unsigned integer: stores only positive values
		data type : uint
		size: 32/64 bits or 4/8 bytes
	*/
	var usign uint = 100;
	var usingn2 uint = 200;
	usin3 := usign+usingn2;
	fmt.Println(usin3)


	//	for float values: float32 ->32-bit/4bytes, float64 ->64-bit/8 bytes
	// default type for float values : float64 
	var average,i,j,k float32;
	i = 84.5
	j = -56.5;
	k = 90.5;
	average = (i+j+k)/3;
	fmt.Println(".2f",average)

	//for char and string types
	//string 
	var Uname string = "gowtham667"
	var pass string = "1242"
	fmt.Println("username : "+Uname, "Password : "+pass)
	fmt.Println("username : "+Uname, "Password : "+pass)

}

