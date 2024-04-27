package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type httpresp struct{
	Args arg        `json:"args"`
	Headers header  `json:"headers"`
	Org string		`json:"origin"`
	Url string		`json:"url"`
}
type header struct{
	Aenc string    `json:"Accept-Encoding"`
	Host string		`json:"Host"`
	Uagent string	`json:"User-Agent"`
	Tid string		`json:"X-Amzn-Trace-Id"`
}
type arg struct{

}
func main(){
	var res *http.Response
	var err error
	res,err = http.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println(err)
	}
	body,err := ioutil.ReadAll(res.Body)
	fmt.Println(string(body))

	var hres httpresp
	err1 := json.Unmarshal(body,&hres)
	if err1 != nil {
		fmt.Println(err1)
	}	

	fmt.Println(hres.Args)
	fmt.Println(hres.Headers.Aenc,hres.Headers.Host,hres.Headers.Tid,hres.Headers.Uagent)
	fmt.Println(hres.Org)
	fmt.Println(hres.Url)
}