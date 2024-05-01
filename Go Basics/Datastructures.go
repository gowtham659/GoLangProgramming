package main

import "fmt"

func main(){
	var weekDays = make(map[int]string)
	fmt.Printf("%v and %t\n",weekDays,weekDays)

	weekDays[1] = "sunday"
	weekDays[2] = "monday"
	weekDays[3] = "tuesday"
	weekDays[4] = "wednesday"

	fmt.Println(weekDays)
	fmt.Println(weekDays[3])

	for sno,day := range(weekDays) { //range is an iterator, first value returns key and second value returns value.
							// if only one variable is used then keys will be returned and second variable is used then relevant value will be returned
		fmt.Println(sno,day)
	}
	sno, sweek := weekDays[3]
	fmt.Println(sno,sweek)

	moon :=map[string]string{"earth" :"moon","jupiter":"europa","saturn":"titan"}
	fmt.Println(moon)

	for planet,satellite := range moon{
		fmt.Println(planet," : ",satellite)
	}
	var pname string ;
	fmt.Scan(&pname)
	fmt.Println(pname)
}