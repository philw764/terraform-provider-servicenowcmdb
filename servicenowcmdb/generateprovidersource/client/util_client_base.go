package client

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"text/template"
)

//const BaseUrl = "https://dev75324.service-now.com/"
const Api = "api/now/cmdb/meta/"

//const BaseClass = "cmdb_ci"

//const Userid = "admin"
//const Password = "XXXXXX"
const Version = "0.01"
const BuildNumber = "1.00"

//var BaseCI CmdbCIMetaModel

// Client is the client used to interact with ServiceNow API.
type Client struct {
	BaseURL string
	Auth    string
}

// ServiceNowClient defines possible methods to call on the ServiceNowClient.
//type ServiceNowClient interface {
//GetCIMetaData(string, CmdbCIMetaModel) error
//WriteCIStruct(model CmdbCIMetaModel) error
//GetListOfClassesFromServiceNow(string, int, []CmdbCIMetaModel) ([]CmdbCIMetaModel, int, error)

//func (client *Client) GetListOfClassesFromServiceNow(Class string, count int, ciClassList []CmdbCIMetaModel)  ([]CmdbCIMetaModel,int, error) {

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

func (client *Client) WriteCIResourcesToFile(ci CmdbCIMetaModel) error {

	//for Attribute := range ci.Result.Attributes {
	//	attrs := &ci.Result.Attributes[Attribute]
	//	attrs.AttributeCamelCase = CamelCaseString(attrs.Element)
	//	attrs.AttrFlags = flags.Get(ci.CiLabelCamelCase, attrs.AttributeCamelCase)

	//		if isBaseCIAttribute(attrs.Element, BaseCI) || ci.Result.Name == "cmdb_ci" {
	//			attrs.IsBaseAttr = true
	//		}
	//}

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

//func (client *Client) WriteCiLookupToFile(ciList []CmdbCIMetaModel) error {

//	file, _ := json.MarshalIndent(ciList, "", "\t")
//	_ = ioutil.WriteFile("../resources/cilookup.json", file, 0644)
//	fmt.Print("this is a test")
//	return nil
//}

//func isBaseCIAttribute(element string, BaseCI CmdbCIMetaModel) bool {
//	for attribute := range BaseCI.Result.Attributes {
//		attrs := BaseCI.Result.Attributes[attribute]
//		if element == attrs.Element {
//			return true
//		}
//	}
//
//	return false
//}
