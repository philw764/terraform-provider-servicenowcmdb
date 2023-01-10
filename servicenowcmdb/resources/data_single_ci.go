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
		ReadContext: readDataSingleCI,

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"class": {
				Type:     schema.TypeString,
				Required: true,
			},
			"sys_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"filter": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"attributes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"attr_name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_value": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_is_reference": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_is_inherited": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_label": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"attr_element": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func readDataSingleCI(ctx context.Context, resourceData *schema.ResourceData, clientInterface interface{}) diag.Diagnostics {
	serviceNowClient := clientInterface.(*client.Client)
	className := resourceData.Get("class").(string)
	sysId := resourceData.Get("sys_id").(string)
	ciname := resourceData.Get("name").(string)
	filter := resourceData.Get("filter").(string)

	if className == "" {
		return diag.FromErr(errors.New("classname is missing, this must be specified"))
	}

	//If we don't have a name or a sys_id then return with an error
	if sysId == "" && ciname == "" {
		//TODO:return error, name or sys_id must be passed in
		return diag.FromErr(errors.New("missing sysid and CI Name"))
	}

	//First we have to get the metadata model for the Class of the CI being queried
	ciMetaDataModel := client.CmdbCIMetaModel{}
	if err := client.GetCIMetaData(className, &ciMetaDataModel, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	//if we have a name and no sys_id then get the sys_id using the ciname as the sysparm_query
	var err error
	if sysId == "" {
		sysId, err = client.GetSysIDForCI(className, filter, ciname, serviceNowClient)
		if err != nil || sysId == "" {
			return diag.FromErr(err)
		}
	}

	//Now get the actual CI data
	var jsonBuf = make(map[string]interface{})
	if jsonBuf, err = client.GetDataForCI(className, sysId, serviceNowClient); err != nil {
		return diag.FromErr(err)
	}

	if err := client.UpdateTFSchema(ciMetaDataModel, jsonBuf, resourceData); err != nil {
		return diag.FromErr(err)
	}

	//return diag.FromErr(errors.New(("this is returned:" + fmt.Sprintf("%s", jsonBuf["cost"]))))
	return nil

}
