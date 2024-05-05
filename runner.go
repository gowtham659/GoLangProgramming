package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
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

// creates and returns a new request object for every modeltestcase
func newRequestObj(endpath string) *http.Request {
	//creates a new request object
	//first parameter: constant which is default method type "GET"
	//second para: URL concatenated with path received from collection.json
	//third: current request method is "GET" so body is nil
	httpreq, er := http.NewRequest(http.MethodGet, "https://httpbin.org/get"+endpath, nil)
	if er != nil {
		log.Fatal("unable to create request")
	}
	return httpreq
}
func (mtc *ModelTestCase) requestBody(req *http.Request) *http.Request {

	if mtc.Method != "GET" {
		req.Method = mtc.Method //adding request type from collection.json to request object
		if mtc.IsBody == true {
			if mtc.BodyContentType == "json" {
				//request body is of type readcloser so the json format must be converted to readcloser type
				bread := strings.NewReader(mtc.BodyJson) //converting json format data to reader tye
				req.Body = ioutil.NopCloser(bread)       // reader type to readcloser type

			} else if (mtc.BodyContentType == "form") {
				if mtc.BodyParams != nil {
					for _, j := range mtc.BodyParams {
						mtcpara, _ := json.Marshal(j)     //converting struct type to byte array format
						mread := bytes.NewReader(mtcpara) //byte array to reader type
						req.Body = io.NopCloser(mread)    //reader to readcloser
					}
				}
			}
		} else {
			nobody := strings.NewReader("no body for current post request")
			req.Body = io.NopCloser(nobody)
		}
	} else {
		req.Body = nil
	}
	return req

}

// returns a request object with data in header section
func (mtc *ModelTestCase) requestHeader(req *http.Request) *http.Request {
	if mtc.HeaderParams != nil { //check headerparams null or not
		for _, j := range mtc.HeaderParams { //iterate over map
			req.Header.Set(j.ParamKey,j.ParamValue)
			
			//mtcpara, _ := json.Marshal(j)      //marshalling modeltypeparameters to byte[]
			//req.Header.Set(i, string(mtcpara)) //adding json data to http header in the form of key/value pairs
		}
	}
	return req
}

func main() {
	// Build a function that takes ModelTestCase and returns a request object
	// reference https://pkg.go.dev/net/http#NewRequestWithContext
	// Goal is to send requests for each ModelTestCase in the order defined
	// in the ModelCollection:TestCaseOrder
	// Refer collection.json

	var mc ModelCollection

	bary, err := ioutil.ReadFile("WebAppPrac/collection.json")
	if err != nil {
		fmt.Println(err)
	}
	er := json.Unmarshal(bary, &mc)
	if er != nil {
		fmt.Println(er)
	}

	for _, j := range mc.TestCaseMap {
		httpreq := newRequestObj(j.Path)
		j.requestBody(httpreq)
		j.requestHeader(httpreq)

		//for displaying request body
		if httpreq.Body != nil { //checking body is null or not
			body, err := ioutil.ReadAll(httpreq.Body)
			if err != nil {
				fmt.Println(err)
				return
			}

			// Display the response body
			fmt.Println(string(body))

		}
		fmt.Println(httpreq)

	}
}
