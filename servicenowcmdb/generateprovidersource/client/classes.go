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
		return ciClassList, nil
	}

	ci := CmdbCIMetaModel{}
	if err := GetCIMetaData(Api+Class, &ci, client); err != nil {
		//TODO:  Need to interrogate and return error message
		return ciClassList, err
	}

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

	// Update each attribute with additional information required for template processing.
	for Attribute := range ci.Result.Attributes {
		attrs := &ci.Result.Attributes[Attribute]
		attrs.AttributeCamelCase = CamelCaseString(attrs.Element)
		attrs.AttrFlags = flags.Get(ci.CiLabelCamelCase, attrs.AttributeCamelCase)
	}

	ciClassList = append(ciClassList, ci)

	fmt.Println("The current CI being looked at:" + Class)

	// If the current CI Class has children then recurse through each one to retrieve their records if
	// the command line options or configuration file have specified "recurse" processing.
	if len(ci.Result.Children) > 0 && recurse {
		for _, child := range ci.Result.Children {
			if IsValidCi(fmt.Sprintf("%v", child)) {
				ciClassList, _ = GetListOfClassesFromServiceNow(fmt.Sprintf("%v", child), recurse, ciClassList, client)
			}
		}
	} else {
		return ciClassList, nil
	}

	return ciClassList, nil
}

// This function sets up the connection to ServiceNow and then ranges through all of the Classes to process as specified
// in the options.ClassList command line option.
func ProcessClassList(options *cli.Options) ([]CmdbCIMetaModel, error) {

	var CiList []CmdbCIMetaModel
	var err error
	// Set up the connection to ServiceNow
	var snowClient *Client
	if env, err := cli.GetEnvVars(); err != nil {
		fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s\n", err)
		os.Exit(1)
	} else {
		snowClient = NewClient(env.Url, env.Userid, env.Password)
	}
	// Classlist is a map specified on the command line and/or config file identifying the classes to be included
	// in the provider.  Each ClassList entry will be set to either "recurse" or "norecurse".  If the map entry
	// is set to "recurse" then all of the child classes for the class are also retrieved and included.
	for class := range options.ClassList {
		recurse := RecurseClasses(options.ClassList[class])
		if classAlreadyInCiList(class, CiList) {
			continue
		} else {
			if CiList, err = GetListOfClassesFromServiceNow(class, recurse, CiList, snowClient); err != nil {
				return nil, err
			}
		}
	}
	return CiList, nil
}

// This function ranges through the list of CI Classes already retrieved from ServiceNow to confirm that the
// next CI Class to retrieve hasn't already been retrieved.  This is required to ensure that the class is not
// included multiple times as a resource when the "provider_baseline.go" file is generated.  Duplicate resources defined in
// the "provider_baseline.go" file will cause a compilation error when the provider is built.
// TODO: Need to generate a report or at least a warning message when this occurs, at the moment tho I'm just
//
//	ignoring it.
func classAlreadyInCiList(class string, ciList []CmdbCIMetaModel) bool {

	for c := range ciList {
		if ciList[c].CiName == class {
			fmt.Printf("This class has already been processed:%s", class)
			return true
		}
	}
	return false

}
func RecurseClasses(recurseOption string) bool {
	if strings.ToLower(recurseOption) == "recurse" {

		return true
	}
	return false
}

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

	request, _ := http.NewRequest(method, client.BaseURL+"/"+path, data)

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
