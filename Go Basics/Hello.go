package main

import (
	"fmt"
	"New/Display"
)

func main(){
	fmt.Println("Hello gowtham thota");
	//emp := employee{"santosh",89,1230,"amp"}
	New.display()
	//pointers: it a variable that holds address of another variable
	var v int = 100
	var p *int = &v
	fmt.Println(v,p,&p)
	var np *int
	fmt.Println("nil pointer: ",np)
	fmt.Println(v,p,*p)
	p1 :=&p
	fmt.Println(v,&v,p,*p,p1,*p1,*&p)

	var arr1  = []int{1,2,3}
	for i := 0; i < 3; i++ {
		fmt.Println(arr1[i])
	}
	var parr1 [3]*int
	parr1[0] = &arr1[0]
	parr1[1] = &arr1[1]
	parr1[2] = &arr1[2]

	fmt.Println(parr1)
	for i:=0;i<3;i++{
		fmt.Println(*parr1[i],parr1[i])
	}

}