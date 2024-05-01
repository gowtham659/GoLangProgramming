package main

import "fmt"

func main(){
	var a int = 10
	var b = 20
	if(a>b){
		fmt.Println("a is larger")
	}else{
		fmt.Println("b is larger")
	}

	score := 78
	if (score>90) {
		fmt.Println("A grade")
	}else if(score>80 && score<=90){
		fmt.Println("B grade")
	}else if (score>70 && score<=80) {
		fmt.Println("C grade")
	}else if (score>60 && score<=70) {
		fmt.Println("D grade")
	}else if (score>50 && score<=60){
		fmt.Println("E grade")
	}else{
		fmt.Printf(" Fail")
	}

	n := -10
	if (n>=0) {
		if (n%2 == 0) {
			fmt.Printf("%v is +even number",n)
		}else {
			fmt.Printf("%v a -odd number",n)
		}
	}else{
		if (n%2 == 0) {
			fmt.Printf("%v is -even number",n)
		}else {
			fmt.Printf("%v is a -odd number",n)
		}
	}

	var val int = 20;
	for i:=1;i<=val;i++ {
		if i%3 == 0 && i%5 == 0{
			fmt.Println("Fizz Buzz")
		}else if (i%3 == 0) {
			fmt.Println("Fizz")
		}else if (i%5 == 0) {
			fmt.Println("Buzz")
		}
	}
}
