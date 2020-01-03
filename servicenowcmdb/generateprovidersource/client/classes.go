package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"terraform-provider-servicenowcmdb/servicenowcmdb/cli"
)

func GetListOfClassesFromServiceNow(Class string, count int, ciClassList []CmdbCIMetaModel, client *Client, options *cli.Options) ([]CmdbCIMetaModel, int, error) {

	// This is a recursive function, this condition is the stop condition for
	// the recursion.

	//TODO: This is where I start tomorrow.  Now that I have the command line list of classes to retrieve I need to see
	//		if the class retrieved is in the list to process.
	if Class == "" || count == 50 {
		//if Class == "" {
		return ciClassList, count, nil
	}
	fmt.Printf("This is the options passed: %s\n", options)
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
				ciClassList, count, _ = GetListOfClassesFromServiceNow(fmt.Sprintf("%v", child), count, ciClassList, client, options)
			}
		}
	} else {
		return ciClassList, count, nil
		//, _ = readCIs("", count)
	}

	return ciClassList, count, nil
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
