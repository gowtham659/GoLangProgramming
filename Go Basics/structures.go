package main

import "fmt"

type employee struct{
	ename string
	age int
	salary float32
	location string
}

func display(emp employee){
	fmt.Println("Employee name:",emp.ename)
	fmt.Println("age:",emp.age)
	fmt.Println("salary:",emp.salary)
	fmt.Println("location:",emp.location)

}

func pdisplay(pemp *employee){
	fmt.Println(pemp.ename)
	fmt.Println(pemp.age)
	fmt.Println(pemp.location)

}
func main(){
 	var emp1 employee
	emp1.ename = "gowtham"
	emp1.age = 23
	emp1.salary = 32000
	emp1.location = "hyd"

	fmt.Println("Employee name:",emp1.ename)
	fmt.Println("age:",emp1.age)
	fmt.Println("salary:",emp1.salary)
	fmt.Println("location:",emp1.location)

	var emp2 employee = employee{"hasha",22,89200,"vizag"}
	fmt.Println(emp2,emp2.ename)

	emp3 := employee{"sarath",24,72322,"mumbai"}
	display(emp3)
	emp3.age = 40
	display(emp3)

	fmt.Scanln(&emp3.salary,&emp3.location)
	display(emp3)
	var pointer_emp employee
	pointer_emp = employee{
		ename: "manoj",
		age: 90,
		location: "manglore",
	}

	pdisplay(&pointer_emp)

	// definig array of a strruct instances
	employee := []employee{
		{ename : "gowtham", age: 23,salary: 90232,location: "errger"},
		{ename: "harsha",age:78,salary:21290,location: "ferrg"},
	}
	fmt.Println(employee)
}