package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/cli"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/flags"
)

func GetListOfClassesFromServiceNow(Class string, recurse bool, ciClassList []CmdbCIMetaModel, client *Client) ([]CmdbCIMetaModel, error) {

	// This is a recursive function, this condition is the stop condition for
	// the recursion.

	//TODO: This is where I start tomorrow.  Now that I have the command line list of classes to retrieve I need to see
	//		if the class retrieved is in the list to process.
	if Class == "" {
		//if Class == "" {
		return ciClassList, nil
	}
	//fmt.Printf("This is the options passed: %s\n", options)
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
		return ciClassList, err
	}
	//if err := GetCIMetaData(Api+Class, &ci, client); err != nil {
	//	//TODO:  Need to interrogate and return error message
	//	return ciClassList, count, err
	//}
	//count++

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

	for Attribute := range ci.Result.Attributes {
		attrs := &ci.Result.Attributes[Attribute]
		attrs.AttributeCamelCase = CamelCaseString(attrs.Element)
		attrs.AttrFlags = flags.Get(ci.CiLabelCamelCase, attrs.AttributeCamelCase)

		//		if isBaseCIAttribute(attrs.Element, BaseCI) || ci.Result.Name == "cmdb_ci" {
		//			attrs.IsBaseAttr = true
		//		}
	}

	ciClassList = append(ciClassList, ci)

	fmt.Println("The current CI being looked at:" + Class)

	//if Class == BaseClass {
	//	BaseCI = ci
	//}

	if len(ci.Result.Children) > 0 {
		for _, child := range ci.Result.Children {
			if IsValidCi(fmt.Sprintf("%v", child)) {
				ciClassList, _ = GetListOfClassesFromServiceNow(fmt.Sprintf("%v", child), recurse, ciClassList, client)
			}
		}
	} else {
		return ciClassList, nil
		//, _ = readCIs("", count)
	}

	return ciClassList, nil
}

func ProcessClassList(options *cli.Options) ([]CmdbCIMetaModel, error) {

	var ciList []CmdbCIMetaModel
	var err error
	// Set up the connection to ServiceNow
	var snowClient *Client
	if env, err := cli.GetEnvVars(); err != nil {
		fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s\n", err)
		os.Exit(1)
	} else {
		snowClient = NewClient(env.Url, env.Userid, env.Password)
	}
	for class := range options.ClassList {
		if ciList, err = GetListOfClassesFromServiceNow(class, true, ciList, snowClient); err != nil {
			return nil, err
		}
	}
	return ciList, nil

}

//type ListofCisFound struct {
//	CiName           string
//	CiNameCamelCase  string
//	CiLabel          string
//	CiLabelCamelCase string
//}

func CamelCaseString(str string) string {
	x := strings.ReplaceAll(str, "_", " ")
	return strings.Join(strings.Split(strings.Title(x), " "), "")
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
