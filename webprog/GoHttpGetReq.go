package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	var res *http.Response
	var err error
	res,err = http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))



}