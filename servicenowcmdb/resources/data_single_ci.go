package resources

import (
	"context"
	"errors"
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

	// Finally, update the Terraform schema with Config Item data from ServiceNow
	if err := UpdateTFSchema(ciMetaDataModel, jsonBuf, resourceData); err != nil {
		return diag.FromErr(err)
	}
	return nil

}
