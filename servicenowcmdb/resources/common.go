package resources

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"strings"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
)

// UpdateTFSchema This function updates the Terraform schema using the ciMetaModel to determine the map key and ciData to set the actual values.
// Still need to handle reference attributes.
func UpdateTFSchema(ciMetaModel client.CmdbCIMetaModel, ciData map[string]interface{}, resourceData *schema.ResourceData) error {

	resourceData.SetId(ciData["sys_id"].(string))
	if err := resourceData.Set("sys_id", ciData["sys_id"].(string)); err != nil {
		return fmt.Errorf("failed to set sys_id field:%s", err)
	}
	if err := resourceData.Set("name", ciData["name"].(string)); err != nil {
		return fmt.Errorf("failed to set name field:%s", err)
	}
	if err := resourceData.Set("filter", resourceData.Get("filter")); err != nil {
		return fmt.Errorf("failed to set filter field:%s", err)
	}

	attrs := make([]interface{}, 0, 0) //len(ciMetaModel.Result.Attributes), len(ciMetaModel.Result.Attributes))
	for _, rec := range ciMetaModel.Result.Attributes {

		if val, ok := CMDBConfigItemKeys[rec.Element]; ok && val == "base" {
			if rec.Type == "reference" {

				if ciData[rec.Element] != "" {

					x := client.StructToMap(ciData[rec.Element])
					if err := resourceData.Set(fmt.Sprintf("%s", rec.Element), client.StructToList(x["value"].(string), x["display_value"].(string), x["link"].(string))); err != nil {
						return err
					}
				} else {
					if err := resourceData.Set(fmt.Sprintf("%s", rec.Element), client.StructToList("", "", "")); err != nil {
						return err
					}
				}
				continue
			}
			if err := resourceData.Set(fmt.Sprintf("%s", rec.Element), fmt.Sprintf("%s", ciData[rec.Element])); err != nil {
				return err
			}
		} else {
			attr := make(map[string]interface{})
			attr["attr_name"] = fmt.Sprintf("%s", rec.Element)
			attr["attr_value"] = fmt.Sprintf("%s", fmt.Sprintf("%s", ciData[rec.Element]))
			attr["attr_type"] = fmt.Sprintf(rec.Type)
			attr["attr_label"] = fmt.Sprintf("%s", rec.Label)

			attrs = append(attrs, attr)
		}
	}
	if err := resourceData.Set("extended_attrs", attrs); err != nil {
		return err
	}

	return nil
}

// ReadTFSchema read CI field data from the Terraform resource schema into a map.  All standard base CI
// fields are stored in the CMDBConfigItemKeys map, range over this map to read each field from the
// resourceData structure.
//
//	 	resourceData 	- This is passed in from Terraform and contains the fields set up in the terraform module.
//			ciMetaDataModel - This struct contains the field retrieved from ServiceNow using the CMDB Meta API for
//							  class.  This is required because all reads and writes to ServiceNow CI are dynamic JSON
//							  and we don't know the CI fields.
//
// TODO: Add read of extended attributes
func ReadTFSchema(resourceData *schema.ResourceData, ciMetaDataModel *client.CmdbCIMetaModel) (map[string]interface{}, error) {
	//attrs := make([]interface{}, 0, 0)
	var ci = make(map[string]interface{})
	for _, rec := range ciMetaDataModel.Result.Attributes {

		if val, ok := CMDBConfigItemKeys[rec.Element]; ok && val == "base" {
			if rec.Type != "reference" {
				ci[rec.Element] = resourceData.Get(rec.Element).(string)
			} else if rec.Type == "reference" {
				ci[rec.Element] = fmt.Sprintf("%s", client.FlattenLinkRel(rec.Element, resourceData))
			}
		} else {
			//r := make(map[string]interface{})
			//r["attr_name"] = fmt.Sprintf("%s", rec.Element)
			//r["attr_value"] = flattenExtendedAttrs(rec.Element, resourceData)

			//attrs = append(attrs, r)
			ci[rec.Element] = flattenExtendedAttrs(rec.Element, resourceData)

		}
	}
	//ci["extended_attrs"] = attrs

	return ci, nil
}

func flattenExtendedAttrs(value string, resourceData *schema.ResourceData) string {
	attrs := resourceData.Get("extended_attrs").([]interface{})
	for _, raw := range attrs {
		i := raw.(map[string]interface{})
		if i["attr_name"] == value {
			return i["attr_value"].(string)
		}
	}
	return ""

}

func IsValidClass(class string) (bool, error) {
	if class == "" || len(class) == 0 {
		return false, fmt.Errorf("class is missing or invalid. Class field must exist")
	}
	if strings.TrimSpace(class) == "" {
		return false, fmt.Errorf("class is missing or invalid. Class field must exist")
	}
	return true, nil
}

func IsValidName(name string) (bool, error) {
	if name == "" || len(name) == 0 {
		return false, fmt.Errorf("class is missing or invalid. Name field must exist")
	}
	if strings.TrimSpace(name) == "" {
		return false, fmt.Errorf("class is missing or invalid. Name field must exist")
	}
	return true, nil
}
