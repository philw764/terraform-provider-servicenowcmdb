package resources

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
)

func ResourceSingleCI() *schema.Resource {
	return &schema.Resource{
		CreateContext: createResourceSingleCI,
		ReadContext:   readResourceSingleCI,
		UpdateContext: updateResourceSingleCI,
		DeleteContext: deleteResourceSingleCI,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: CMDBSchema(),
	}
}

// createResourceSingleCI - This function perform the create processing for apply.
func createResourceSingleCI(ctx context.Context, resourceData *schema.ResourceData, clientInterface interface{}) diag.Diagnostics {
	var err error
	serviceNowClient := clientInterface.(*client.Client)
	className := resourceData.Get("class").(string)
	name := resourceData.Get("name").(string)

	if ok, err := IsValidClass(className); !ok {
		return diag.FromErr(err)
	}
	if ok, err := IsValidName(name); !ok {
		return diag.FromErr(err)
	}
	//Query the Metadata model for the class from ServiceNow.  This is used to build the json
	//structures for input and output from ServiceNow.  This is required because we need to
	//know the field structure of the CI Class dynamically.
	ciMetaDataModel := client.CmdbCIMetaModel{}
	if err = client.GetCIMetaData(className, &ciMetaDataModel, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	//Load in the CI from the Terraform schema resource, which will be used to create the CI
	//in ServiceNow.
	var ci = make(map[string]interface{})
	if ci, err = ReadTFSchema(resourceData, &ciMetaDataModel); err != nil {
		return diag.FromErr(err)
	}

	//Construct the JSON payload and call the ServiceNow cmdb web API to create the CI
	delete(ci, "sys_id")
	delete(ci, "sys_class_name")
	var x = make(map[string]interface{})
	x["attributes"] = ci

	var data json.RawMessage //Wow - this is really important or jsom.Marshall will base64 encode the var.
	data, _ = json.MarshalIndent(x, "\r", "\t")

	var jsonData []byte
	SnowUrl := client.CMDBInstanceApi + className + "?sysparm_fields=sys_id"
	jsonData, err = serviceNowClient.RequestJSON("POST", SnowUrl, data)
	if err != nil {
		resourceData.SetId("")
		return diag.FromErr(fmt.Errorf("This is jsonData: %s    error:%s", data, err))
	}

	//Set the resource ID and read back the CI from ServiceNow to pass the detail back
	//to Terraform.
	resourceData.SetId(client.GetSysId(jsonData))
	return readResourceConfigurationItem(ctx, resourceData, serviceNowClient)
}

func readResourceSingleCI(ctx context.Context, resourceData *schema.ResourceData, clientInterface interface{}) diag.Diagnostics {
	serviceNowClient := clientInterface.(*client.Client)
	className := resourceData.Get("class").(string)
	sysId := resourceData.Id()

	name := resourceData.Get("name").(string)
	//filter := resourceData.Get("filter").(string)

	if className == "" {
		return diag.FromErr(errors.New("classname is missing, this must be specified"))
	}

	if name == "" {
		//TODO:return error, name or sys_id must be passed in
		return diag.FromErr(errors.New("missing CI Name"))
	}

	//Get the Metadata model for CI class, this is used to determine the CI fields
	ciMetaDataModel := client.CmdbCIMetaModel{}
	if err := client.GetCIMetaData(className, &ciMetaDataModel, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	var err error

	//Now get the actual CI data
	var jsonBuf = make(map[string]interface{})
	if jsonBuf, err = client.GetDataForCI(className, sysId, serviceNowClient); err != nil {

		return diag.FromErr(errors.New("Failed to get data for ci:" + fmt.Sprintf("%s", className) + "  SysID:" + fmt.Sprintf("%s", sysId)))
		return diag.FromErr(err)
	}

	//Finally, update the Terraform schema with Config Item data from ServiceNow
	if err := UpdateTFSchema(ciMetaDataModel, jsonBuf, resourceData); err != nil {
		return diag.FromErr(errors.New("Failed to get data for ci:" + fmt.Sprintf("%s", className) + "  SysID:" + fmt.Sprintf("%s", sysId)))
		return diag.FromErr(err)
	}
	return nil

}

func updateResourceSingleCI(ctx context.Context, resourceData *schema.ResourceData, clientInterface interface{}) diag.Diagnostics {
	var err error
	serviceNowClient := clientInterface.(*client.Client)
	className := resourceData.Get("class").(string)
	//sysId := resourceData.Get("sys_id").(string)
	name := resourceData.Get("name").(string)

	if className == "" {
		return diag.FromErr(errors.New("classname is missing, this must be specified"))
	}

	if name == "" {
		//TODO:return error, name or sys_id must be passed in
		return diag.FromErr(errors.New("missing CI Name"))
	}

	//Get the Metadata model for CI class, this is used to determine the CI fields
	ciMetaDataModel := client.CmdbCIMetaModel{}
	if err = client.GetCIMetaData(className, &ciMetaDataModel, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	var ci = make(map[string]interface{})
	if ci, err = ReadTFSchema(resourceData, &ciMetaDataModel); err != nil {
		return diag.FromErr(err)
	}
	delete(ci, "sys_id")
	delete(ci, "sys_class_name")
	var x = make(map[string]interface{})
	x["attributes"] = ci
	var data json.RawMessage //Wow - this is really important or jsom.Marshall will base64 encode the var.
	data, _ = json.MarshalIndent(x, "\r", "\t")
	SnowUrl := client.CMDBInstanceApi + className + "/" + resourceData.Id()
	var jsonData []byte
	jsonData, err = serviceNowClient.RequestJSON("PATCH", SnowUrl, data)
	if err != nil {
		resourceData.SetId("")
		return diag.FromErr(fmt.Errorf("This is jsonData: %s", err))
	}
	resourceData.SetId(client.GetSysId(jsonData))
	return readResourceConfigurationItem(ctx, resourceData, serviceNowClient)
}

func deleteResourceSingleCI(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {
	//servicenowClient := serviceNowClient.(*Client)
	var ci ConfigurationItemPost
	if err := copyFromTerraformToServiceNowConfigurationItem(resourceData, &ci); err != nil {
		return diag.FromErr(err)
	}

	if err := resourceData.Set("install_status", "7"); err != nil {
		//return fmt.Errorf("ConfigurationItemfailed to set install_status field during destroy action %s", err)
		return diag.FromErr(err)
	}

	if err := resourceData.Set("operational_status", "6"); err != nil {
		//return fmt.Errorf("ConfigurationItemfailed to set install_status field during destroy action %s", err)
		return diag.FromErr(err)
	}
	//SnowUrl := CMDBInstanceApi + CiNamecmdb_ci + "/" + resourceData.Id()
	//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	//_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	if err := updateResourceConfigurationItem(ctx, resourceData, serviceNowClient); err != nil {
		return err
	}

	return nil
}
