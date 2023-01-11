package resources

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
	//"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
)

// Terraform Provider for CI Classes in a ServiceNow instance.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"instance_url": {
				Type:        schema.TypeString,
				Description: "The Url of the ServiceNow instance to work with.",
				Required:    true,
			},
			"username": {
				Type:        schema.TypeString,
				Description: "Username used to manage resources in the ServiceNow instance using Basic authentication.",
				Required:    true,
			},
			"password": {
				Type:        schema.TypeString,
				Description: "Password of the user to manage resources.",
				Required:    true,
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"servicenowcmdb_windows_server": ResourceWindowsServer(),
			"servicenowcmdb_ci":             ResourceConfigurationItem(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"servicenowcmdb_single_ci": DataSingleCI(),
		},

		ConfigureContextFunc: configure,
	}
}

func configure(ctx context.Context, data *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// Create a new client to talk to the instance.
	log.Printf("[DEBUG] In configure")
	servicenowClient := client.NewClient(
		data.Get("instance_url").(string),
		data.Get("username").(string),
		data.Get("password").(string))
	return servicenowClient, nil
}
