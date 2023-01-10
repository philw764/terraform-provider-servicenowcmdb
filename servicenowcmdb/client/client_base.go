package client

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io"
	"net/http"
)

const CMDBInstanceApi = "/api/now/cmdb/instance/"

// Client HTTP Connector
type Client struct {
	BaseURL string
	Auth    string
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

var AttributesToExclude = []string{
	"sys_updated_on",
	"sys_updated_by",
	"sys_created_on",
	"sys_created_by",
	"sys_domain_path", //system related field, TODO: Might need to consider this for application Domains
	"sys_class_name",  //
	"sys_id",          //This field CANNOT be in a http POST request or the "create" action fails
	"sys_class_path",
	"sys_mod_count",
	"sys_tags",
	"",
}

type LinkRelStruct struct {
	DisplayValue string `json:"display_value"`
	Link         string `json:"link"`
	Value        string `json:"value"`
}

type ModelIDRel struct {
	Result struct {
		Attributes struct {
			ModelId struct {
				DisplayValue string `json:"display_value"`
				Link         string `json:"link"`
				Value        string `json:"value"`
			} `json:"model_id"`
		} `json:"Attributes"`
	} `json:"result"`
}

type AssetList struct {
	Result struct {
		Attributes struct {
			Asset struct {
				DisplayValue string `json:"display_value"`
				Link         string `json:"link"`
				Value        string `json:"value"`
			} `json:"asset"`
		} `json:"Attributes"`
	} `json:"result"`
}

type OutboundRelations struct {
	OutboundRel []struct {
		SysID string `json:"sys_id"`
		Type  struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"type"`
		Target struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"target"`
	} `json:"outbound_relations"`
}

type InboundRelations []struct {
	InboundRel []struct {
		SysID string `json:"sys_id"`
		Type  struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"type"`
		Target struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"target"`
	} `json:"inbound_relations"`
}

type SysId struct {
	Result struct {
		Attributes struct {
			SysId string `json:"sys_id"`
		}
	} `json:"result"`
}

func GetSysId(jsonData []byte) string {
	sysId := SysId{}
	if err := json.Unmarshal(jsonData, &sysId); err != nil {
		return ""
	}
	return sysId.Result.Attributes.SysId
}

// This function is used when updating the Terraform schema to unpack and update the fields that are
// reference attributes in ServiceNow.  For example "location" is a reference field.  A reference
// field consists of 3 fields, "link", "value" and "display_value", "value" is the important field
// the other fields are for aesthetics.
func StructToMap(in interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(in)
	json.Unmarshal(j, &m)

	return m
}

func StructToList(value string, link string, display_value string) []interface{} {
	result := make([]interface{}, 1, 1)
	r := make(map[string]interface{})
	r["display_value"] = display_value
	r["value"] = value
	r["link"] = link
	result[0] = r

	return result

}

func StructToList2(value string) []interface{} {
	result := make([]interface{}, 1, 1)
	r := make(map[string]interface{})
	r["value"] = value
	result[0] = r

	return result

}

func FlattenLinkRel(linkRel string, resourceData *schema.ResourceData) string {
	relsRaw := resourceData.Get(linkRel).([]interface{})
	for _, raw := range relsRaw {
		i := raw.(map[string]interface{})
		return i["value"].(string)
	}
	return ""
}

// requestJSON execute an HTTP request and returns the raw response data.
func (client *Client) RequestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	request, _ := http.NewRequest(method, client.BaseURL+path, data)

	// Add headers.
	request.Header.Set("Authorization", client.Auth)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	responseData, _ := io.ReadAll(response.Body)

	if response.StatusCode >= 300 || response.StatusCode < 200 {
		return nil, fmt.Errorf("HTTP response status %s, %s", response.Status, responseData)
	}
	return responseData, nil
}
