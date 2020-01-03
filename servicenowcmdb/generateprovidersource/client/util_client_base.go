package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"terraform-provider-servicenowcmdb/servicenowcmdb/cli"

	//"terraform-provider-servicenowcmdb/servicenowcmdb/cli"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/flags"
	"text/template"
)

//const BaseUrl = "https://dev75324.service-now.com/"
const Api = "api/now/cmdb/meta/"
const BaseClass = "cmdb_ci"

//const Userid = "admin"
//const Password = "XXXXXX"
const Version = "0.01"
const BuildNumber = "1.00"

var BaseCI CmdbCIMetaModel

// Client is the client used to interact with ServiceNow API.
type Client struct {
	BaseURL string
	Auth    string
}

type ListofCisFound struct {
	CiName           string
	CiNameCamelCase  string
	CiLabel          string
	CiLabelCamelCase string
}

func CamelCaseString(str string) string {
	x := strings.ReplaceAll(str, "_", " ")
	return strings.Join(strings.Split(strings.Title(x), " "), "")
}

// ServiceNowClient defines possible methods to call on the ServiceNowClient.
//type ServiceNowClient interface {
//GetCIMetaData(string, CmdbCIMetaModel) error
//WriteCIStruct(model CmdbCIMetaModel) error
//ReadCIs(string, int, []CmdbCIMetaModel) ([]CmdbCIMetaModel, int, error)

//func (client *Client) ReadCIs(Class string, count int, ciClassList []CmdbCIMetaModel)  ([]CmdbCIMetaModel,int, error) {

//}

// NewClient returns a new ServiceNowClient.
func NewClient(baseURL string, username string, password string) *Client {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &Client{
		BaseURL: baseURL,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func GetCIMetaData(endpoint string, responseObjectOut interface{}, client *Client) error {
	jsonResponse, err := requestJSON("GET", endpoint, nil, client)
	if err != nil {
		return err
	}
	// DEBUG:  Use an objmap to unmarshal the JSON data to see format
	// var objmap map[string]interface{}
	//if err := json.Unmarshal(jsonResponse, &objmap); err != nil {
	//	return err
	//}
	if err := json.Unmarshal(jsonResponse, responseObjectOut); err != nil {
		return err
	}
	return nil
}

// requestJSON execute an HTTP request and returns the raw response data.
func requestJSON(method string, path string, jsonData interface{}, client *Client) ([]byte, error) {
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	request, _ := http.NewRequest(method, client.BaseURL+path, data)

	// Add the needed headers.
	request.Header.Set("Authorization", client.Auth)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)

	if err != nil {
		return nil, err
	}

	responseData, _ := ioutil.ReadAll(response.Body)

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return nil, fmt.Errorf("HTTP response status %s, %s", response.Status, responseData)
	}

	return responseData, nil
}

func (client *Client) WriteCIStructToFile(ci CmdbCIMetaModel) error {

	for Attribute := range ci.Result.Attributes {
		attrs := &ci.Result.Attributes[Attribute]
		attrs.AttributeCamelCase = CamelCaseString(attrs.Element)
		attrs.AttrFlags = flags.Get(ci.CiLabelCamelCase, attrs.AttributeCamelCase)

		if isBaseCIAttribute(attrs.Element, BaseCI) || ci.Result.Name == "cmdb_ci" {
			attrs.IsBaseAttr = true
		}
	}

	resourceTemplate, err := ioutil.ReadFile("templates/cmdb_ci_template.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("CmdbCi").Parse(string(resourceTemplate))
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("../resources/"+ci.Result.Name+".go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}
	err = t.Execute(file, ci)
	if err != nil {
		panic(err)
	}
	return nil
}

func IsValidCi(ci string) bool {
	if strings.Contains(ci, "cmdb_ci_endpoint") {
		return false
	}
	if strings.Contains(ci, "cmdb_ci") {
		return true
	}
	return false

}

// This function connects to ServiceNow using the MetaData API to pull the details for every CI in
// the CMDB.
func ReadCIs(Class string, count int, ciClassList []CmdbCIMetaModel, client *Client, options *cli.Options) ([]CmdbCIMetaModel, int, error) {

	// This is a recursive function, this condition is the stop condition for
	// the recursion.

	//TODO: This is where I start tomorrow.  Now that I have the command line list of classes to retrieve I need to see
	//		if the class retrieved is in the list to process.
	if Class == "" || count == 50 {
		//if Class == "" {
		return ciClassList, count, nil
	}
	fmt.Printf("This is the options passed: %s", options)
	//snowClient := NewClient(BaseUrl, Userid, Password)
	//var snowClient *Client
	//if env, err := cli.GetEnvVars(); err != nil {
	//	fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s", err)
	//	return nil, count, err
	//} else {
	//	fmt.Printf("This is the userid:%s\n", env.Userid)
	//	fmt.Printf("This is the pwd:%s\n", env.Password)
	//	fmt.Printf("This is the url:%s\n", env.Url)
	//	snowClient = NewClient(env.Url, env.Userid, env.Password)
	//}
	ci := CmdbCIMetaModel{}
	if err := GetCIMetaData(Api+Class, &ci, client); err != nil {
		//TODO:  Need to interrogate and return error message
		return ciClassList, count, err
	}
	//if err := GetCIMetaData(Api+Class, &ci, client); err != nil {
	//	//TODO:  Need to interrogate and return error message
	//	return ciClassList, count, err
	//}
	count++

	ci.CiName = Class
	ci.Version = Version
	ci.GeneratorVersion = BuildNumber

	ci.CiLabel = strings.ToLower(strings.ReplaceAll(ci.Result.Label, " ", "_"))

	ci.CiLabel = strings.ReplaceAll(ci.CiLabel, "-", "")
	ci.CiLabel = strings.ReplaceAll(ci.CiLabel, "/", "")
	ci.CiLabel = strings.ReplaceAll(ci.CiLabel, ".", "Dot")
	ci.CiName = strings.ReplaceAll(ci.CiName, "-", "")
	ci.CiName = strings.ReplaceAll(ci.CiName, "-", "")
	ci.CiName = strings.ReplaceAll(ci.CiName, ".", "Dot")

	ci.CiLabelCamelCase = CamelCaseString(ci.CiLabel)
	ci.CiNameCamelCase = CamelCaseString(ci.CiName)

	ciClassList = append(ciClassList, ci)

	fmt.Println("The current CI being looked at:" + Class + "\tCount is:" + strconv.Itoa(count))

	if Class == BaseClass {
		BaseCI = ci
	}

	if len(ci.Result.Children) > 0 {
		for _, child := range ci.Result.Children {
			if IsValidCi(fmt.Sprintf("%v", child)) {
				ciClassList, count, _ = ReadCIs(fmt.Sprintf("%v", child), count, ciClassList, client)
			}
		}
	} else {
		return ciClassList, count, nil
		//, _ = readCIs("", count)
	}

	return ciClassList, count, nil
}

func (client *Client) WriteProviderToFile(ciList []CmdbCIMetaModel) error {

	resourceTemplate, err := ioutil.ReadFile("templates/provider.tmpl")
	if err != nil {
		panic(err)
	}

	t, err := template.New("Provider").Parse(string(resourceTemplate))
	if err != nil {
		panic(err)
	}

	file, err := os.OpenFile("../resources/provider.go", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		return err
	}
	err = t.Execute(file, ciList)
	if err != nil {
		panic(err)
	}
	return nil
}

func (client *Client) WriteCiLookupToFile(ciList []CmdbCIMetaModel) error {

	file, _ := json.MarshalIndent(ciList, "", "\t")
	_ = ioutil.WriteFile("../resources/cilookup.json", file, 0644)
	fmt.Print("this is a test")
	return nil
}

func isBaseCIAttribute(element string, BaseCI CmdbCIMetaModel) bool {
	for attribute := range BaseCI.Result.Attributes {
		attrs := BaseCI.Result.Attributes[attribute]
		if element == attrs.Element {
			return true
		}
	}

	return false
}
