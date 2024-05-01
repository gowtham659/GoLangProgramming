package main

import (
	"fmt"
)

func main()  {
	var arr = [5]int{0:1,2:23,4};
	fmt.Println(arr)

	var arr2 = []float32{0.23,1.232,0.003,1.231,9.0101};
	fmt.Println(arr2);
	var arr3 = []string{"apple","cow"}
	fmt.Printf("values: %v type: %T",arr3,arr3)

	var a = [6]int{};
	a[0] = 23
	a[1] = 53
	a[3] = 90
	a[4] = 12
	a[5] = 38
	fmt.Println(a,a[3],a[5])

	arr4 :=[2][2]int{
		{0:10,1:20},
		{0:50,1:40}}
	fmt.Println(arr4)
}