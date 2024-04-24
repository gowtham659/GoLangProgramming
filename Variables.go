package main

import (
	"fmt"
)

var company string = "TCS"
//sal := 10000;
func main() {
	//case1 : declaring variables using var keyword 
	var name string = "gowtham";
	fmt.Println("Hello "+name)
	//we have to specify either type or value or both 
	// declaring vars by specifying type
	var marks int;
	marks = 92;
	fmt.Println(marks)
	// declaring vars by specifying value
	var lang = "telugu";
	fmt.Println("language: "+lang);

	//we cannot decalre a varible without mentioning type or value while using var keyword
	//var loc;
	//loc = "fr";
	//fmt.Println(loc)


	//case2: declaring variables using := operator
	//Here type of variable is inferred from value assigned means compiler decides the type of variable based on type of value.
	//here value assign is mandatory
	age  := 30;
	fmt.Println(age)

	//default values of datatypes
	var a int;
	var b float32;
	var c string;
	var d bool;
	fmt.Println(a);
	fmt.Println(b);
	fmt.Println(c);
	fmt.Println(d);

	//assignment after declaration
	var x float32;
	var y float32;
	x=10.23;
	y=32.42;
	fmt.Println(x/y)

	//diff b/w var and := 
	//var
	//	Can be used inside and outside of functions	
	//:=
	//	Can only be used inside functions
	//var
	//	Variable declaration and value assignment can be done separately	
	//:=
	//	Variable declaration and value assignment cannot be done separately (must be done in the same line)
	fmt.Println(company)


	//declaring multiple variables on same line
	//if type is used then only possible to declare one type of variables per line.
	var m,n,o int = 10,20,40;
	fmt.Println(m)
	fmt.Println(n)
	fmt.Println(o)

	//if type is not specified then any no of variables and any type of values can be declared, assigned
	var i,j,k  = 10,23.343, "gowtham";
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(k)

	
	//Go Variable Declaration in a Block
	//Multiple variable declarations can also be grouped together into a block for greater readability:
	var(
		ename string ="harsha";
		sal int;
		
		eloc string = "hyd";
	)
	sal = 23243;
	fmt.Println(ename)
	fmt.Println(sal)
	fmt.Println(eloc)
	
	//output functions
	//Print() : prints in a single line
	var email string = "gowtham123@yahoo.com"
	uid := "gowtham893"
	fmt.Print(email, uid)
	fmt.Print(email," ",uid)
	fmt.Print("\n",email," ",uid,"\n")
	fmt.Print(i,j)
	//Println() : adds white space b/w args and a newline at end
	fmt.Println(email,uid)

	//Printf() : formats args based on given fromat verbs and prints it 
	fmt.Printf("value is %v",email)
	fmt.Printf("type is %T",email)

	
}