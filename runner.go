package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)


type ModelCollection struct {
	UID              string            `json:"UID"`
	Name             string            `json:"Name"`
	Description      string            `json:"Description"`
	Source           string            `json:"Source"`
	CreatedAt        string            `json:"CreatedAt"`
	TestCaseOrder    map[int]string    `json:"TestCaseOrder"`
	Variables        map[string]string `json:"Variables"`
	DynamicVariables map[string]string `json:"DynamicVariables"`
	//Servers          map[string]ModelServer     `json:"Servers"`
	TestIdList  []string                 `json:"TestIdList"`
	AuthOptions map[string]string        `json:"AuthOptions"`
	TestCaseMap map[string]ModelTestCase `json:"TestCaseMap"`
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

func (mtc *ModelTestCase) GetURL(server string) *url.URL{
	furl:=fmt.Sprintf("%s%s",server,mtc.Path)
	URLA,_:= url.Parse(furl)
	if mtc.QueryParams != nil{
		qval:= URLA.Query()
		for _,j:= range mtc.QueryParams{
			qval.Add(j.ParamKey,j.ParamValue)
		}
		URLA.RawQuery = qval.Encode()
	}
	return URLA
}

func (mtc *ModelTestCase) GetHeader() http.Header{
	var head http.Header = http.Header{}
	if mtc.HeaderParams != nil { //check headerparams null or not
		for _, j := range mtc.HeaderParams { //iterate over map
			head.Add(j.ParamKey,j.ParamValue)
		}
	}
	fmt.Println(head)
	return head 
}

func (mtc *ModelTestCase) GetBody(req *http.Request){
	if mtc.Method != "GET" && mtc.IsBody{
		if mtc.BodyContentType == "Form"{
			for _,j:= range mtc.BodyParams{
				bar,er:=json.Marshal(j)
				if er!= nil{
					fmt.Println("eero",er)
				}
				bread:=bytes.NewReader(bar)
				req.Body = io.NopCloser(bread) //reader to readcloser
			}
		}else if mtc.BodyContentType == "json" {
			req.Body.Read([]byte(mtc.BodyJson))
		}
	}
}
func (mtc *ModelTestCase) CreateRequest(server string) (*http.Request,error){
	req,_:=http.NewRequest(mtc.Method,server,nil)
	req.URL = mtc.GetURL(server)
	req.Header = mtc.GetHeader()
	mtc.GetBody(req)
	if req ==nil {
		er:=errors.New("unable to request")
		return nil,er
	}
	return req,nil
}

func main(){
	var mc ModelCollection
	jdata,er:= os.ReadFile("./collection.json")
	if er != nil {
		fmt.Println("file not found",er)
	}
	json.Unmarshal(jdata,&mc)

	for _,j:= range mc.TestCaseOrder{ 
		mtc:=mc.TestCaseMap[j]
		req,er:=mtc.CreateRequest("https://geeksforgeeks.com/")
		if er != nil {
			fmt.Println(er)
		}
		fmt.Println(*req)
	}
}