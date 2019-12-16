package client

type CmdbCIMetaModel struct {
	//CmdbBaseObject
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
	//Attributes struct {
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
	//} `json:"attributes"`
}

func (c CmdbCIMetaModel) HasExtendedAttributes() bool {

	for attr := range c.Result.Attributes {
		if !c.Result.Attributes[attr].IsBaseAttr {
			return true
		}
	}
	return false
}

func (c CmdbCIMetaModel) IncludeAttribute(element string) bool {
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
