package resources

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
)

const DataCiNamecmdb_ci = "cmdb_ci"

func DataConfigurationItem() *schema.Resource {
	return &schema.Resource{
		ReadContext: readDataConfigurationItem,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"class_name": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func readDataConfigurationItem(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {
	servicenowClient := serviceNowClient.(*client.Client)
	SnowUrl := client.CMDBInstanceApi + CiNamecmdb_ci + "/" + resourceData.Id()
	var jsonData []byte
	jsonData, err := servicenowClient.RequestJSON("GET", SnowUrl, nil)
	if err != nil {
		resourceData.SetId("")
		return diag.FromErr(err)
	}

	if err := copyFromServiceNowToTerraformConfigurationItem(resourceData, jsonData); err != nil {
		return diag.FromErr(err)
	}
	return nil
}
