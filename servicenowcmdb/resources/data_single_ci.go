package resources

import (
	"context"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
)

func DataSingleCI() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataCI,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: CMDBSchema(),
	}
}

func readDataCI(ctx context.Context, resourceData *schema.ResourceData, clientInterface interface{}) diag.Diagnostics {
	serviceNowClient := clientInterface.(*client.Client)
	className := resourceData.Get("class").(string)
	sysId := resourceData.Get("sys_id").(string)
	name := resourceData.Get("name").(string)
	filter := resourceData.Get("filter").(string)

	if className == "" {
		return diag.FromErr(errors.New("classname is missing, this must be specified"))
	}

	if sysId == "" && name == "" {
		//TODO:return error, name or sys_id must be passed in
		return diag.FromErr(errors.New("missing sysid and CI Name"))
	}

	//Get the Metadata model for CI class, this is used to determine the CI fields
	ciMetaDataModel := client.CmdbCIMetaModel{}
	if err := client.GetCIMetaData(className, &ciMetaDataModel, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	//if we have a name and no sys_id then get the sys_id using the name as the sysparm_query
	var err error
	if sysId == "" {
		sysId, err = client.GetSysIDForCI(className, filter, name, serviceNowClient)
		if err != nil || sysId == "" {
			return diag.FromErr(err)
		}
	}

	//Now get the actual CI data
	var jsonBuf = make(map[string]interface{})
	if jsonBuf, err = client.GetDataForCI(className, sysId, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	if err := UpdateTFSchema(ciMetaDataModel, jsonBuf, resourceData); err != nil {
		return diag.FromErr(err)
	}
	return nil

}

// UpdateTFSchema This function updates the Terraform schema using the ciMetaModel to determine the map key and ciData to set the actual values.
// Still need to handle reference attributes.
func UpdateTFSchema(ciMetaModel client.CmdbCIMetaModel, ciData map[string]interface{}, resourceData *schema.ResourceData) error {

	resourceData.SetId(ciData["sys_id"].(string))
	resourceData.Set("sys_id", ciData["sys_id"].(string))
	resourceData.Set("name", ciData["name"].(string))
	resourceData.Set("filter", resourceData.Get("filter"))

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
			attr["attr_value"] = fmt.Sprintf("%s", ciData[rec.Element])
			attr["attr_type"] = fmt.Sprintf(rec.Type)
			attr["attr_label"] = fmt.Sprintf("%s", rec.Label)

			attrs = append(attrs, attr)
		}
	}
	if err := resourceData.Set("extended_attrs", &attrs); err != nil {
		return err
	}

	return nil
}
