package main

type ModelCollection struct {
	UID              string                     `json:"UID"`
	Name             string                     `json:"Name"`
	Description      string                     `json:"Description"`
	Source           string                     `json:"Source"`
	CreatedAt        string                     `json:"CreatedAt"`
	TestCaseOrder    map[int]string             `json:"TestCaseOrder"`
	Variables        map[string]string          `json:"Variables"`
	DynamicVariables map[string]string          `json:"DynamicVariables"`
	Servers          map[string]ModelServer     `json:"Servers"`
	TestIdList       []string                   `json:"TestIdList"`
	AuthOptions      map[string]string          `json:"AuthOptions"`
	TestCaseMap      map[string]ModelTestCase   `json:"TestCaseMap"`
	ResultMap        map[string]ModelTestResult `json:"ResultMap"`
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


func main(){
	// Build a function that takes ModelTestCase and returns a request object
	// reference https://pkg.go.dev/net/http#NewRequestWithContext
	// Goal is to send requests for each ModelTestCase in the order defined
	// in the ModelCollection:TestCaseOrder
	// Refer collection.json
}

