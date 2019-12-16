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
	"text/template"
)

const BaseUrl = "https://dev75324.service-now.com/"
const Api = "api/now/cmdb/meta/"
const BaseClass = "cmdb_ci"
const Userid = "admin"
const Password = "T1v0l1uu"
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
type ServiceNowClient interface {
	GetCIMetaData(string, CmdbCIMetaModel) error
	WriteCIStruct(model CmdbCIMetaModel) error
	ReadCIs(string, int, []CmdbCIMetaModel) ([]CmdbCIMetaModel, int, error)

	//func (client *Client) ReadCIs(Class string, count int, ciClassList []CmdbCIMetaModel)  ([]CmdbCIMetaModel,int, error) {

}

// NewClient is a factory method used to return a new ServiceNowClient.
func NewClient(baseURL string, username string, password string) *Client {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &Client{
		BaseURL: baseURL,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func (client *Client) GetCIMetaData(endpoint string, responseObjectOut interface{}) error {
	jsonResponse, err := client.requestJSON("GET", endpoint, nil)
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
func (client *Client) requestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
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

	//ci.CiName = CamelCaseString(ci.Result.Name)
	//ci.CiLabel = strings.ToLower(strings.ReplaceAll(ci.Result.Label, " ", "_"))
	//ci.CiLabelCamelCase = CamelCaseString(ci.Result.Label)
	for Attribute := range ci.Result.Attributes {
		attrs := &ci.Result.Attributes[Attribute]
		attrs.AttributeCamelCase = CamelCaseString(attrs.Element)

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

	file, err := os.OpenFile("../resources/"+ci.Result.Name+".go", os.O_TRUNC|os.O_CREATE, os.FileMode(0644))
	if err != nil {
		return err
	}
	err = t.Execute(file, ci)
	if err != nil {
		panic(err)
	}
	return nil
}
// This function connects to ServiceNow using the MetaData API to pull the details for every CI in
// the CMDB.
func (client *Client) ReadCIs(Class string, count int, ciClassList []CmdbCIMetaModel) ([]CmdbCIMetaModel, int, error) {

	// This is a recursive function, this condition is the stop condition for
	// the recursion.

	//TODO: hard code only getting 10 classes during dev and test
	if Class == "" || count == 4 {
		return ciClassList, count, nil
	}

	snowClient := NewClient(BaseUrl, Userid, Password)
	ci := CmdbCIMetaModel{}
	if err := snowClient.GetCIMetaData(Api+Class, &ci); err != nil {
		//TODO:  Need to interrogate and return error message
		return ciClassList, count, err
	}
	count++
	ci.CiName = Class
	ci.Version = Version
	ci.GeneratorVersion = BuildNumber
	ci.CiNameCamelCase = CamelCaseString(Class)
	ci.CiLabel = strings.ToLower(strings.ReplaceAll(ci.Result.Label, " ", "_"))
	ci.CiLabelCamelCase = CamelCaseString(ci.Result.Label)

	ciClassList = append(ciClassList, ci)

	fmt.Println("The current CI being looked at:" + Class + "\tCount is:" + strconv.Itoa(count))

	if Class == BaseClass {
		BaseCI = ci
	}

	if len(ci.Result.Children) > 0 {
		for _, child := range ci.Result.Children {
			ciClassList, count, _ = client.ReadCIs(fmt.Sprintf("%v", child), count, ciClassList)
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

	file, err := os.OpenFile("../resources/provider.go", os.O_TRUNC|os.O_CREATE, os.FileMode(0644))
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
