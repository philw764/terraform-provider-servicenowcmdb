package client

import (
	"encoding/base64"
	"io/ioutil"
	"os"
	"strings"

	"text/template"
)

const Api = "api/now/cmdb/meta/"
const Version = "0.01"
const BuildNumber = "1.00"

// Client is the client used to interact with ServiceNow API.
type Client struct {
	BaseURL string
	Auth    string
}

// NewClient returns a new ServiceNowClient.
func NewClient(baseURL string, username string, password string) *Client {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &Client{
		BaseURL: baseURL,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}

func WriteCIResourcesToFile(ci CmdbCIMetaModel) error {

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

func WriteProviderToFile(ciList []CmdbCIMetaModel) error {

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
