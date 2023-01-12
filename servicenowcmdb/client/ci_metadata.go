package client

import (
	"encoding/json"
	"fmt"
)

const MetaDataAPI = "/api/now/cmdb/meta/"

// CmdbCIMetaModel - This structure is filled using the CMDB Meta API.  The CMDB Class defines
// the CMDB table in ServiceNow, this structure defines the fields that make up this table.
type CmdbCIMetaModel struct {
	CiName           string
	CiNameCamelCase  string
	CiLabel          string
	CiLabelCamelCase string
	Version          string
	GeneratorVersion string
	Result           struct {
		IconUrl      string        `json:"icon_url"`
		IsExtendable bool          `json:"is_extendable"`
		Parent       string        `json:"parent"`
		Children     []interface{} `json:"children"`
		Name         string        `json:"name"`
		Icon         string        `json:"icon"`

		Attributes []struct {
			Reference          string `json:"reference"`
			IsInherited        string `json:"is_inherited"`
			IsMandatory        string `json:"is_mandatory"`
			IsReadOnly         string `json:"is_read_only"`
			DefaultValue       string `json:"default_value"`
			AttrFlags          string
			Label              string `json:"label"`
			AttributeCamelCase string
			Type               string `json:"type"`
			Element            string `json:"element"`
			MaxLength          string `json:"max_length"`
			IsDisplay          string `json:"is_display"`
			IsBaseAttr         bool
		} `json:"attributes"`
		Label             string `json:"label"`
		RelationshipRules []struct {
			Parent  string `json:"parent"`
			RelType string `json:"relation_type"`
			Child   string `json:"child"`
		} `json:"relationship_rules"`
	} `json:"result"`
}

type CmdbCIAttribute struct {
	Reference    string `json:"reference"`
	IsInherited  string `json:"is_inherited"`
	IsMandatory  string `json:"is_mandatory"`
	IsReadOnly   string `json:"is_read_only"`
	DefaultValue string `json:"default_value"`
	Label        string `json:"label"`
	Type         string `json:"type"`
	Element      string `json:"element"`
	MaxLength    string `json:"max_length"`
	IsDisplay    string `json:"is_display"`
}

func (c *CmdbCIMetaModel) IncludeAttribute(element string) bool {
	var excludeList = []string{
		"sys_id",
		"sys_class_name",
		"sys_class_path",
		"sys_domain",
	}
	for excluded := range excludeList {
		if element == excludeList[excluded] {

			return false
		}
	}
	return true
}

func (c *CmdbCIMetaModel) IsReferenceAttr(reference string) bool {
	if reference != "" {
		return true
	}
	return false
}

func GetCIMetaData(class string, responseObjectOut interface{}, client *Client) error {
	endpoint := MetaDataAPI + class
	jsonResponse, err := client.RequestJSON("GET", endpoint, nil)
	if err != nil {
		return fmt.Errorf("Failed to get Metadata for class:%s  error: %s", class, err)
	}
	if err := json.Unmarshal(jsonResponse, responseObjectOut); err != nil {
		return fmt.Errorf("Failed to Unmarshal json response for class:%s   error:", class, err)
	}
	return nil
}
