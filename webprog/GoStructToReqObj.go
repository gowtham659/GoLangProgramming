package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)
type ModelCollection struct {
	UID              string                     `json:"UID"`
	Name             string                     `json:"Name"`
	Description      string                     `json:"Description"`
	Source           string                     `json:"Source"`
	CreatedAt        string                     `json:"CreatedAt"`
	TestCaseOrder    map[int]string             `json:"TestCaseOrder"`
	Variables        map[string]string          `json:"Variables"`
	DynamicVariables map[string]string          `json:"DynamicVariables"`
	//Servers          map[string]ModelServer     `json:"Servers"`
	TestIdList       []string                   `json:"TestIdList"`
	AuthOptions      map[string]string          `json:"AuthOptions"`
	TestCaseMap      map[string]ModelTestCase   `json:"TestCaseMap"`
	//ResultMap        map[string]ModelTestResult `json:"ResultMap"`
}
type ModelTestCase struct {
	Order           int                               `json:"Order"`
	TestId          string                            `json:"TestId"`
	Method          string                            `json:"Method"`
	Path            string                            `json:"Path"`
	IsBody          bool                              `json:"IsBody"`
	BodyContentType string                            `json:"BodyContentType"`
	RequestParams   map[string]ModelTestCaseParameter `json:"RequestParams"`
	BodyParams      map[string]ModelTestCaseParameter `json:"BodyParams"`
	HeaderParams    map[string]ModelTestCaseParameter `json:"HeaderParams"`
	QueryParams     map[string]ModelTestCaseParameter `json:"QueryParams"`
	PathParams      map[string]ModelTestCaseParameter `json:"PathParams"`
	CreatedAt       string                            `json:"CreatedAt"`
	UpdatedAt       string                            `json:"UpdatedAt"`
	BodyJson        string                            `json:"BodyJson"`
}

type ModelTestCaseParameter struct {
	Key         string `json:"Key"`
	Name        string `json:"Name"`
	DataType    string `json:"DataType"`
	VariableKey string `json:"VariableKey"`
	ParamKey    string `json:"ParamKey"`
	ParamValue  string `json:"ParamValue"`
}

func (mtc *ModelTestCase) requestObj() *http.Request{
	

	barr,_:=json.Marshal(mtc)
	req, err := http.NewRequest("POST", "https://httpbin.org/get", bytes.NewReader(barr))
	if err != nil {
    log.Fatalf("impossible to build request: %s", err)
	}
	//fmt.Println(*req.URL)
	req.Header.Set("User-Agent", "MyUserAgent/1.0")
	return req

}
func main(){
	
	//fmt.Println(string(bary))
	var mc ModelCollection
	var mtc ModelTestCase

	//fmt.Println(mtc.requestObj())	
	bary,err := ioutil.ReadFile("WebAppPrac/collection.json")
	if err != nil{
		fmt.Println(err)
	}
	er:= json.Unmarshal(bary,&mc)
	if er != nil{
		fmt.Println(er)
	}
	
	for _,j :=range mc.TestCaseMap{
		req:=j.requestObj()
		fmt.Println(req)
		json.NewDecoder(req.Body).Decode(&mtc)
		//fmt.Println(mtc)
		//body,_:=req.GetBody()
		//fmt.Println(body)
	}
}