package resources

///////////////////////////////////////////////////////////////////////////////////////////////////
//  ****WARNING***    : This is automatically generated code  using the generatprovidersource
//                      utility in this package, any edits to this file will be overwritten
//                      and lost.  Refer to the readme.md in the generateprovidersource directory
//                       for details on how to regenerate this code.
// Provider Version    : 0.01
// Generator Version   : 1.00
//
//  Description       :   This file is the resource provider for the cmdb_ci CMDB Class.  This code is executed
//                        when the servicenowcmdb_configuration_item keyword is used in a terraform script (*.tf) file
//
//                        This file is will need to be regenerated if the ServiceNow CMDB base CI Class
//                        "cmdb_ci" or of the cmdb_ci CI Class is modified.
///////////////////////////////////////////////////////////////////////////////////////////////////
import (
	"context"
	"encoding/json"
	"terraform-provider-servicenowcmdb/servicenowcmdb/client"
	//"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// This constant stores the name of the CI and is used to construct the URL to ServiceNow
const CiNamecmdb_ci = "cmdb_ci"

// This is the structure to construct the JSON payload when POSTing to ServiceNow.  This is needed because
// ServiceNow has strict parsing on the JSON Data and will fail if the JSON format doesn't match exactly.
// This is essentially the same as ConfigurationItemGet but does not contain the "value", "link" and
// "display_value" fields for reference objects.
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
type ConfigurationItemPost struct {
	Result struct {
		Attributes struct {
			AttestedDate         string `json:"attested_date,omitempty"`
			SkipSync             string `json:"skip_sync,omitempty"`
			OperationalStatus    string `json:"operational_status,omitempty"`
			SysUpdatedOn         string `json:"sys_updated_on,omitempty"`
			AttestationScore     string `json:"attestation_score,omitempty"`
			DiscoverySource      string `json:"discovery_source,omitempty"`
			FirstDiscovered      string `json:"first_discovered,omitempty"`
			SysUpdatedBy         string `json:"sys_updated_by,omitempty"`
			DueIn                string `json:"due_in,omitempty"`
			SysCreatedOn         string `json:"sys_created_on,omitempty"`
			InstallDate          string `json:"install_date,omitempty"`
			InvoiceNumber        string `json:"invoice_number,omitempty"`
			GlAccount            string `json:"gl_account,omitempty"`
			SysCreatedBy         string `json:"sys_created_by,omitempty"`
			WarrantyExpiration   string `json:"warranty_expiration,omitempty"`
			AssetTag             string `json:"asset_tag,omitempty"`
			Fqdn                 string `json:"fqdn,omitempty"`
			ChangeControl        string `json:"change_control,omitempty"`
			OwnedBy              string `json:"owned_by,omitempty"`
			CheckedOut           string `json:"checked_out,omitempty"`
			SysDomainPath        string `json:"sys_domain_path,omitempty"`
			BusinessUnit         string `json:"business_unit,omitempty"`
			DeliveryDate         string `json:"delivery_date,omitempty"`
			MaintenanceSchedule  string `json:"maintenance_schedule,omitempty"`
			InstallStatus        string `json:"install_status,omitempty"`
			CostCenter           string `json:"cost_center,omitempty"`
			AttestedBy           string `json:"attested_by,omitempty"`
			SupportedBy          string `json:"supported_by,omitempty"`
			DnsDomain            string `json:"dns_domain,omitempty"`
			Name                 string `json:"name,omitempty"`
			Assigned             string `json:"assigned,omitempty"`
			PurchaseDate         string `json:"purchase_date,omitempty"`
			Subcategory          string `json:"subcategory,omitempty"`
			LifeCycleStage       string `json:"life_cycle_stage,omitempty"`
			ShortDescription     string `json:"short_description,omitempty"`
			AssignmentGroup      string `json:"assignment_group,omitempty"`
			ManagedBy            string `json:"managed_by,omitempty"`
			ManagedByGroup       string `json:"managed_by_group,omitempty"`
			LastDiscovered       string `json:"last_discovered,omitempty"`
			CanPrint             string `json:"can_print,omitempty"`
			Manufacturer         string `json:"manufacturer,omitempty"`
			PoNumber             string `json:"po_number,omitempty"`
			CheckedIn            string `json:"checked_in,omitempty"`
			Vendor               string `json:"vendor,omitempty"`
			LifeCycleStageStatus string `json:"life_cycle_stage_status,omitempty"`
			MacAddress           string `json:"mac_address,omitempty"`
			Company              string `json:"company,omitempty"`
			ModelNumber          string `json:"model_number,omitempty"`
			Justification        string `json:"justification,omitempty"`
			Department           string `json:"department,omitempty"`
			AssignedTo           string `json:"assigned_to,omitempty"`
			StartDate            string `json:"start_date,omitempty"`
			Cost                 string `json:"cost,omitempty"`
			Comments             string `json:"comments,omitempty"`
			AttestationStatus    string `json:"attestation_status,omitempty"`
			CmdbOtEntity         string `json:"cmdb_ot_entity,omitempty"`
			SysModCount          string `json:"sys_mod_count,omitempty"`
			SerialNumber         string `json:"serial_number,omitempty"`
			Monitor              string `json:"monitor,omitempty"`
			ModelId              string `json:"model_id,omitempty"`
			IpAddress            string `json:"ip_address,omitempty"`
			DuplicateOf          string `json:"duplicate_of,omitempty"`
			SysTags              string `json:"sys_tags,omitempty"`
			CostCc               string `json:"cost_cc,omitempty"`
			SupportGroup         string `json:"support_group,omitempty"`
			OrderDate            string `json:"order_date,omitempty"`
			Schedule             string `json:"schedule,omitempty"`
			Environment          string `json:"environment,omitempty"`
			Due                  string `json:"due,omitempty"`
			Attested             string `json:"attested,omitempty"`
			Unverified           string `json:"unverified,omitempty"`
			CorrelationId        string `json:"correlation_id,omitempty"`
			Attributes           string `json:"attributes,omitempty"`
			Location             string `json:"location,omitempty"`
			Asset                string `json:"asset,omitempty"`
			Category             string `json:"category,omitempty"`
			FaultCount           string `json:"fault_count,omitempty"`
			LeaseId              string `json:"lease_id,omitempty"`
		} `json:"attributes"`
	} `json:"result"`
}

// This is the structure where JSON data is Unmarshal'ed from ServiceNow.  This is different from the POST
// structure because it has to contain the "value", "link" and "display_value" fields for reference objects
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
type ConfigurationItemGet struct {
	Result struct {
		Attributes struct {
			AttestedDate       string `json:"attested_date,omitempty"`
			SkipSync           string `json:"skip_sync,omitempty"`
			OperationalStatus  string `json:"operational_status,omitempty"`
			SysUpdatedOn       string `json:"sys_updated_on,omitempty"`
			AttestationScore   string `json:"attestation_score,omitempty"`
			DiscoverySource    string `json:"discovery_source,omitempty"`
			FirstDiscovered    string `json:"first_discovered,omitempty"`
			SysUpdatedBy       string `json:"sys_updated_by,omitempty"`
			DueIn              string `json:"due_in,omitempty"`
			SysCreatedOn       string `json:"sys_created_on,omitempty"`
			InstallDate        string `json:"install_date,omitempty"`
			InvoiceNumber      string `json:"invoice_number,omitempty"`
			GlAccount          string `json:"gl_account,omitempty"`
			SysCreatedBy       string `json:"sys_created_by,omitempty"`
			WarrantyExpiration string `json:"warranty_expiration,omitempty"`
			AssetTag           string `json:"asset_tag,omitempty"`
			Fqdn               string `json:"fqdn,omitempty"`
			ChangeControl      struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"change_control"`
			OwnedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"owned_by"`
			CheckedOut    string `json:"checked_out,omitempty"`
			SysDomainPath string `json:"sys_domain_path,omitempty"`
			BusinessUnit  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"business_unit"`
			DeliveryDate        string `json:"delivery_date,omitempty"`
			MaintenanceSchedule struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"maintenance_schedule"`
			InstallStatus string `json:"install_status,omitempty"`
			CostCenter    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"cost_center"`
			AttestedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"attested_by"`
			SupportedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"supported_by"`
			DnsDomain      string `json:"dns_domain,omitempty"`
			Name           string `json:"name,omitempty"`
			Assigned       string `json:"assigned,omitempty"`
			PurchaseDate   string `json:"purchase_date,omitempty"`
			Subcategory    string `json:"subcategory,omitempty"`
			LifeCycleStage struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"life_cycle_stage"`
			ShortDescription string `json:"short_description,omitempty"`
			AssignmentGroup  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"assignment_group"`
			ManagedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"managed_by"`
			ManagedByGroup struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"managed_by_group"`
			LastDiscovered string `json:"last_discovered,omitempty"`
			CanPrint       string `json:"can_print,omitempty"`
			Manufacturer   struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"manufacturer"`
			PoNumber  string `json:"po_number,omitempty"`
			CheckedIn string `json:"checked_in,omitempty"`
			Vendor    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"vendor"`
			LifeCycleStageStatus struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"life_cycle_stage_status"`
			MacAddress string `json:"mac_address,omitempty"`
			Company    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"company"`
			ModelNumber   string `json:"model_number,omitempty"`
			Justification string `json:"justification,omitempty"`
			Department    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"department"`
			AssignedTo struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"assigned_to"`
			StartDate         string `json:"start_date,omitempty"`
			Cost              string `json:"cost,omitempty"`
			Comments          string `json:"comments,omitempty"`
			AttestationStatus string `json:"attestation_status,omitempty"`
			CmdbOtEntity      struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"cmdb_ot_entity"`
			SysModCount  string `json:"sys_mod_count,omitempty"`
			SerialNumber string `json:"serial_number,omitempty"`
			Monitor      string `json:"monitor,omitempty"`
			ModelId      struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"model_id"`
			IpAddress   string `json:"ip_address,omitempty"`
			DuplicateOf struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"duplicate_of"`
			SysTags      string `json:"sys_tags,omitempty"`
			CostCc       string `json:"cost_cc,omitempty"`
			SupportGroup struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"support_group"`
			OrderDate string `json:"order_date,omitempty"`
			Schedule  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"schedule"`
			Environment   string `json:"environment,omitempty"`
			Due           string `json:"due,omitempty"`
			Attested      string `json:"attested,omitempty"`
			Unverified    string `json:"unverified,omitempty"`
			CorrelationId string `json:"correlation_id,omitempty"`
			Attributes    string `json:"attributes,omitempty"`
			Location      struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"location"`
			Asset struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"asset"`
			Category   string `json:"category,omitempty"`
			FaultCount string `json:"fault_count,omitempty"`
			LeaseId    string `json:"lease_id,omitempty"`
		} `json:"attributes"`
	} `json:"result"`
}

// This is schema.resource for the ServiceNow Provider.  The resource name is appended
// to the name when the function is automatically generated using generateprovidersource
// to ensure uniqueness of the function within the package.
//
// There are two types of fields for ServiceNow:
//    1. basic types: string, int etc
//    2. Reference types: These are fields that contain a reference pointer to another object within
//       the ServiceNow Database.  For example "Location" is a reference field which stores a pointer
//       ("sys_id") to the corresponding Location record.  The "sys_id" is stored in the "value" field
//       returned from ServiceNow in the JSON payload.
//
// TODO: Need to construct a way of setting the "Computed", "Optional" and "Mandatory" switches.  I
//       think use a lookup stored in a separate file.  The user can edit this field to suite their
//       requirements for managing attributes and provides a method for setting these switches for
//       custom attributes.

func ResourceConfigurationItem() *schema.Resource {
	return &schema.Resource{
		CreateContext: createResourceConfigurationItem,
		ReadContext:   readResourceConfigurationItem,
		UpdateContext: updateResourceConfigurationItem,
		DeleteContext: deleteResourceConfigurationItem,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"attested_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_sync": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operational_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sys_updated_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attestation_score": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovery_source": {
				Type:     schema.TypeString,
				Optional: true, Default: "Terraform",
			},
			"first_discovered": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"due_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"install_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"invoice_number": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"gl_account": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sys_created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"warranty_expiration": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"change_control": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"owned_by": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"checked_out": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sys_domain_path": {
				Type:     schema.TypeString,
				Optional: true, Default: "/",
			},
			"business_unit": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"delivery_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"install_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost_center": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"attested_by": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"supported_by": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"dns_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"assigned": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"purchase_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"subcategory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"life_cycle_stage": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"short_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assignment_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"managed_by": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"managed_by_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"last_discovered": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"can_print": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"manufacturer": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"po_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checked_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"life_cycle_stage_status": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"company": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"model_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"justification": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"department": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"assigned_to": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"start_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attestation_status": {
				Type:     schema.TypeString,
				Optional: true, Default: "Not Yet Reviewed",
			},
			"cmdb_ot_entity": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sys_mod_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"monitor": {
				Type:     schema.TypeString,
				Optional: true, Default: "false",
			},
			"model_id": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"ip_address": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"duplicate_of": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"sys_tags": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cost_cc": {
				Type:     schema.TypeString,
				Optional: true, Default: "USD",
			},
			"support_group": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"order_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"schedule": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"environment": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"due": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attested": {
				Type:     schema.TypeString,
				Optional: true, Default: "false",
			},
			"unverified": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"correlation_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"attributes": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"asset": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"value": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"link": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"category": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fault_count": {
				Type:     schema.TypeString,
				Optional: true, Default: "0",
			},
			"lease_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

// Create routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
func createResourceConfigurationItem(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {

	servicenowClient := serviceNowClient.(*client.Client) //Client Connection details
	// Use common function to update base attributes
	var ci ConfigurationItemPost
	if err := copyFromTerraformToServiceNowConfigurationItem(resourceData, &ci); err != nil {
		return diag.FromErr(err)
	}

	//t := time.Now()
	//ci.Result.Attributes.Installdate = t.String()
	//ci.Result.Attributes.SysCreatedOn = t.String()

	// Using "sysparm fields=sys_id" reduces the amount of data received, we only need the sys_id for create action
	SnowUrl := client.CMDBInstanceApi + CiNamecmdb_ci + "?sysparm_fields=sys_id"

	jsonData, err := servicenowClient.RequestJSON("POST", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return diag.FromErr(err)
	}

	resourceData.SetId(client.GetSysId(jsonData))
	return readResourceConfigurationItem(ctx, resourceData, serviceNowClient)
}

// Read routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
func readResourceConfigurationItem(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {

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

// Update routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
func updateResourceConfigurationItem(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {
	servicenowClient := serviceNowClient.(*client.Client)

	var ci ConfigurationItemPost
	if err := copyFromTerraformToServiceNowConfigurationItem(resourceData, &ci); err != nil {
		return diag.FromErr(err)
	}

	SnowUrl := client.CMDBInstanceApi + CiNamecmdb_ci + "/" + resourceData.Id()
	//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return diag.FromErr(err)
	}
	return readResourceConfigurationItem(ctx, resourceData, serviceNowClient)
}

// TODO:  Need to work out what to do with deleting CIs. ServiceNow does not support deleting CIs via the API
//        CIs shouldn't be deleted because may be referenced by other objects in the ServiceNow database
//        (eg changes and incidents).  Therefore the best practice is to change their operational and
//        install status.  The tricky bit is working out how to make this flexible for everyone to use.

// Delete routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
func deleteResourceConfigurationItem(ctx context.Context, resourceData *schema.ResourceData, serviceNowClient interface{}) diag.Diagnostics {

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
	//if err != nil {
	//resourceData.SetId("")
	//return err
	//}
	return nil
}

// This routine is called to copy data from the Terraform schema.ResourceData object into the CMDB CI object
// It would be nice if Terraform implemented a funciton to return a list of field names in a slice, this would
// make it easier to loop through the structure instead of doing a "Get" per field.
func copyFromTerraformToServiceNowConfigurationItem(resourceData *schema.ResourceData, ci *ConfigurationItemPost) error {

	attrs := &ci.Result.Attributes
	attrs.AttestedDate = resourceData.Get("attested_date").(string)
	attrs.SkipSync = resourceData.Get("skip_sync").(string)
	attrs.OperationalStatus = resourceData.Get("operational_status").(string)
	attrs.SysUpdatedOn = resourceData.Get("sys_updated_on").(string)
	attrs.AttestationScore = resourceData.Get("attestation_score").(string)
	attrs.DiscoverySource = resourceData.Get("discovery_source").(string)
	attrs.FirstDiscovered = resourceData.Get("first_discovered").(string)
	attrs.SysUpdatedBy = resourceData.Get("sys_updated_by").(string)
	attrs.DueIn = resourceData.Get("due_in").(string)
	attrs.SysCreatedOn = resourceData.Get("sys_created_on").(string)
	attrs.InstallDate = resourceData.Get("install_date").(string)
	attrs.InvoiceNumber = resourceData.Get("invoice_number").(string)
	attrs.GlAccount = resourceData.Get("gl_account").(string)
	attrs.SysCreatedBy = resourceData.Get("sys_created_by").(string)
	attrs.WarrantyExpiration = resourceData.Get("warranty_expiration").(string)
	attrs.AssetTag = resourceData.Get("asset_tag").(string)
	attrs.Fqdn = resourceData.Get("fqdn").(string)
	//attrs.ChangeControl  = resourceData.Get("change_control.value").(string)
	attrs.ChangeControl = client.FlattenLinkRel("change_control", resourceData)
	//attrs.OwnedBy  = resourceData.Get("owned_by.value").(string)
	attrs.OwnedBy = client.FlattenLinkRel("owned_by", resourceData)
	attrs.CheckedOut = resourceData.Get("checked_out").(string)
	attrs.SysDomainPath = resourceData.Get("sys_domain_path").(string)
	//attrs.BusinessUnit  = resourceData.Get("business_unit.value").(string)
	attrs.BusinessUnit = client.FlattenLinkRel("business_unit", resourceData)
	attrs.DeliveryDate = resourceData.Get("delivery_date").(string)
	//attrs.MaintenanceSchedule  = resourceData.Get("maintenance_schedule.value").(string)
	attrs.MaintenanceSchedule = client.FlattenLinkRel("maintenance_schedule", resourceData)
	attrs.InstallStatus = resourceData.Get("install_status").(string)
	//attrs.CostCenter  = resourceData.Get("cost_center.value").(string)
	attrs.CostCenter = client.FlattenLinkRel("cost_center", resourceData)
	//attrs.AttestedBy  = resourceData.Get("attested_by.value").(string)
	attrs.AttestedBy = client.FlattenLinkRel("attested_by", resourceData)
	//attrs.SupportedBy  = resourceData.Get("supported_by.value").(string)
	attrs.SupportedBy = client.FlattenLinkRel("supported_by", resourceData)
	attrs.DnsDomain = resourceData.Get("dns_domain").(string)
	attrs.Name = resourceData.Get("name").(string)
	attrs.Assigned = resourceData.Get("assigned").(string)
	attrs.PurchaseDate = resourceData.Get("purchase_date").(string)
	attrs.Subcategory = resourceData.Get("subcategory").(string)
	//attrs.LifeCycleStage  = resourceData.Get("life_cycle_stage.value").(string)
	attrs.LifeCycleStage = client.FlattenLinkRel("life_cycle_stage", resourceData)
	attrs.ShortDescription = resourceData.Get("short_description").(string)
	//attrs.AssignmentGroup  = resourceData.Get("assignment_group.value").(string)
	attrs.AssignmentGroup = client.FlattenLinkRel("assignment_group", resourceData)
	//attrs.ManagedBy  = resourceData.Get("managed_by.value").(string)
	attrs.ManagedBy = client.FlattenLinkRel("managed_by", resourceData)
	//attrs.ManagedByGroup  = resourceData.Get("managed_by_group.value").(string)
	attrs.ManagedByGroup = client.FlattenLinkRel("managed_by_group", resourceData)
	attrs.LastDiscovered = resourceData.Get("last_discovered").(string)
	attrs.CanPrint = resourceData.Get("can_print").(string)
	//attrs.Manufacturer  = resourceData.Get("manufacturer.value").(string)
	attrs.Manufacturer = client.FlattenLinkRel("manufacturer", resourceData)
	attrs.PoNumber = resourceData.Get("po_number").(string)
	attrs.CheckedIn = resourceData.Get("checked_in").(string)
	//attrs.Vendor  = resourceData.Get("vendor.value").(string)
	attrs.Vendor = client.FlattenLinkRel("vendor", resourceData)
	//attrs.LifeCycleStageStatus  = resourceData.Get("life_cycle_stage_status.value").(string)
	attrs.LifeCycleStageStatus = client.FlattenLinkRel("life_cycle_stage_status", resourceData)
	attrs.MacAddress = resourceData.Get("mac_address").(string)
	//attrs.Company  = resourceData.Get("company.value").(string)
	attrs.Company = client.FlattenLinkRel("company", resourceData)
	attrs.ModelNumber = resourceData.Get("model_number").(string)
	attrs.Justification = resourceData.Get("justification").(string)
	//attrs.Department  = resourceData.Get("department.value").(string)
	attrs.Department = client.FlattenLinkRel("department", resourceData)
	//attrs.AssignedTo  = resourceData.Get("assigned_to.value").(string)
	attrs.AssignedTo = client.FlattenLinkRel("assigned_to", resourceData)
	attrs.StartDate = resourceData.Get("start_date").(string)
	attrs.Cost = resourceData.Get("cost").(string)
	attrs.Comments = resourceData.Get("comments").(string)
	attrs.AttestationStatus = resourceData.Get("attestation_status").(string)
	//attrs.CmdbOtEntity  = resourceData.Get("cmdb_ot_entity.value").(string)
	attrs.CmdbOtEntity = client.FlattenLinkRel("cmdb_ot_entity", resourceData)
	attrs.SysModCount = resourceData.Get("sys_mod_count").(string)
	attrs.SerialNumber = resourceData.Get("serial_number").(string)
	attrs.Monitor = resourceData.Get("monitor").(string)
	//attrs.ModelId  = resourceData.Get("model_id.value").(string)
	attrs.ModelId = client.FlattenLinkRel("model_id", resourceData)
	attrs.IpAddress = resourceData.Get("ip_address").(string)
	//attrs.DuplicateOf  = resourceData.Get("duplicate_of.value").(string)
	attrs.DuplicateOf = client.FlattenLinkRel("duplicate_of", resourceData)
	attrs.SysTags = resourceData.Get("sys_tags").(string)
	attrs.CostCc = resourceData.Get("cost_cc").(string)
	//attrs.SupportGroup  = resourceData.Get("support_group.value").(string)
	attrs.SupportGroup = client.FlattenLinkRel("support_group", resourceData)
	attrs.OrderDate = resourceData.Get("order_date").(string)
	//attrs.Schedule  = resourceData.Get("schedule.value").(string)
	attrs.Schedule = client.FlattenLinkRel("schedule", resourceData)
	attrs.Environment = resourceData.Get("environment").(string)
	attrs.Due = resourceData.Get("due").(string)
	attrs.Attested = resourceData.Get("attested").(string)
	attrs.Unverified = resourceData.Get("unverified").(string)
	attrs.CorrelationId = resourceData.Get("correlation_id").(string)
	attrs.Attributes = resourceData.Get("attributes").(string)
	//attrs.Location  = resourceData.Get("location.value").(string)
	attrs.Location = client.FlattenLinkRel("location", resourceData)
	//attrs.Asset  = resourceData.Get("asset.value").(string)
	attrs.Asset = client.FlattenLinkRel("asset", resourceData)
	attrs.Category = resourceData.Get("category").(string)
	attrs.FaultCount = resourceData.Get("fault_count").(string)
	attrs.LeaseId = resourceData.Get("lease_id").(string)
	return nil
}

// This function Unmarshal's the JSON payload from ServiceNow into a corresponding STRUCT for the CI.  It then
// Sets each field in the Terraform schema.ResourceData object with the corrsponding field from the CI STRUCT.
//
// NOTE:  Terraform expects a MAP for fields that represent "reference fields" so that the "value", "link" and
//
//	"display_value" are decoded correctly by Terraform.  The map is constructed for each reference field
//	using a common function called "StructToMap" in the "client_base.go" file.
func copyFromServiceNowToTerraformConfigurationItem(resourceData *schema.ResourceData, jsonData []byte) error {
	ci := ConfigurationItemGet{}
	if err := json.Unmarshal(jsonData, &ci); err != nil {
		//resourceData.SetId("")
		//return err
	}

	resourceData.SetId(client.GetSysId(jsonData))
	attrs := ci.Result.Attributes
	_ = resourceData.Set("attested_date", attrs.AttestedDate)
	_ = resourceData.Set("skip_sync", attrs.SkipSync)
	_ = resourceData.Set("operational_status", attrs.OperationalStatus)
	_ = resourceData.Set("sys_updated_on", attrs.SysUpdatedOn)
	_ = resourceData.Set("attestation_score", attrs.AttestationScore)
	_ = resourceData.Set("discovery_source", attrs.DiscoverySource)
	_ = resourceData.Set("first_discovered", attrs.FirstDiscovered)
	_ = resourceData.Set("sys_updated_by", attrs.SysUpdatedBy)
	_ = resourceData.Set("due_in", attrs.DueIn)
	_ = resourceData.Set("sys_created_on", attrs.SysCreatedOn)
	_ = resourceData.Set("install_date", attrs.InstallDate)
	_ = resourceData.Set("invoice_number", attrs.InvoiceNumber)
	_ = resourceData.Set("gl_account", attrs.GlAccount)
	_ = resourceData.Set("sys_created_by", attrs.SysCreatedBy)
	_ = resourceData.Set("warranty_expiration", attrs.WarrantyExpiration)
	_ = resourceData.Set("asset_tag", attrs.AssetTag)
	_ = resourceData.Set("fqdn", attrs.Fqdn)
	//_ = resourceData.Set("change_control", StructToMap(attrs.ChangeControl))
	_ = resourceData.Set("change_control", client.StructToList(attrs.ChangeControl.Value, attrs.ChangeControl.Link, attrs.ChangeControl.DisplayValue))
	//_ = resourceData.Set("owned_by", StructToMap(attrs.OwnedBy))
	_ = resourceData.Set("owned_by", client.StructToList(attrs.OwnedBy.Value, attrs.OwnedBy.Link, attrs.OwnedBy.DisplayValue))
	_ = resourceData.Set("checked_out", attrs.CheckedOut)
	_ = resourceData.Set("sys_domain_path", attrs.SysDomainPath)
	//_ = resourceData.Set("business_unit", StructToMap(attrs.BusinessUnit))
	_ = resourceData.Set("business_unit", client.StructToList(attrs.BusinessUnit.Value, attrs.BusinessUnit.Link, attrs.BusinessUnit.DisplayValue))
	_ = resourceData.Set("delivery_date", attrs.DeliveryDate)
	//_ = resourceData.Set("maintenance_schedule", StructToMap(attrs.MaintenanceSchedule))
	_ = resourceData.Set("maintenance_schedule", client.StructToList(attrs.MaintenanceSchedule.Value, attrs.MaintenanceSchedule.Link, attrs.MaintenanceSchedule.DisplayValue))
	_ = resourceData.Set("install_status", attrs.InstallStatus)
	//_ = resourceData.Set("cost_center", StructToMap(attrs.CostCenter))
	_ = resourceData.Set("cost_center", client.StructToList(attrs.CostCenter.Value, attrs.CostCenter.Link, attrs.CostCenter.DisplayValue))
	//_ = resourceData.Set("attested_by", StructToMap(attrs.AttestedBy))
	_ = resourceData.Set("attested_by", client.StructToList(attrs.AttestedBy.Value, attrs.AttestedBy.Link, attrs.AttestedBy.DisplayValue))
	//_ = resourceData.Set("supported_by", StructToMap(attrs.SupportedBy))
	_ = resourceData.Set("supported_by", client.StructToList(attrs.SupportedBy.Value, attrs.SupportedBy.Link, attrs.SupportedBy.DisplayValue))
	_ = resourceData.Set("dns_domain", attrs.DnsDomain)
	_ = resourceData.Set("name", attrs.Name)
	_ = resourceData.Set("assigned", attrs.Assigned)
	_ = resourceData.Set("purchase_date", attrs.PurchaseDate)
	_ = resourceData.Set("subcategory", attrs.Subcategory)
	//_ = resourceData.Set("life_cycle_stage", StructToMap(attrs.LifeCycleStage))
	_ = resourceData.Set("life_cycle_stage", client.StructToList(attrs.LifeCycleStage.Value, attrs.LifeCycleStage.Link, attrs.LifeCycleStage.DisplayValue))
	_ = resourceData.Set("short_description", attrs.ShortDescription)
	//_ = resourceData.Set("assignment_group", StructToMap(attrs.AssignmentGroup))
	_ = resourceData.Set("assignment_group", client.StructToList(attrs.AssignmentGroup.Value, attrs.AssignmentGroup.Link, attrs.AssignmentGroup.DisplayValue))
	//_ = resourceData.Set("managed_by", StructToMap(attrs.ManagedBy))
	_ = resourceData.Set("managed_by", client.StructToList(attrs.ManagedBy.Value, attrs.ManagedBy.Link, attrs.ManagedBy.DisplayValue))
	//_ = resourceData.Set("managed_by_group", StructToMap(attrs.ManagedByGroup))
	_ = resourceData.Set("managed_by_group", client.StructToList(attrs.ManagedByGroup.Value, attrs.ManagedByGroup.Link, attrs.ManagedByGroup.DisplayValue))
	_ = resourceData.Set("last_discovered", attrs.LastDiscovered)
	_ = resourceData.Set("can_print", attrs.CanPrint)
	//_ = resourceData.Set("manufacturer", StructToMap(attrs.Manufacturer))
	_ = resourceData.Set("manufacturer", client.StructToList(attrs.Manufacturer.Value, attrs.Manufacturer.Link, attrs.Manufacturer.DisplayValue))
	_ = resourceData.Set("po_number", attrs.PoNumber)
	_ = resourceData.Set("checked_in", attrs.CheckedIn)
	//_ = resourceData.Set("vendor", StructToMap(attrs.Vendor))
	_ = resourceData.Set("vendor", client.StructToList(attrs.Vendor.Value, attrs.Vendor.Link, attrs.Vendor.DisplayValue))
	//_ = resourceData.Set("life_cycle_stage_status", StructToMap(attrs.LifeCycleStageStatus))
	_ = resourceData.Set("life_cycle_stage_status", client.StructToList(attrs.LifeCycleStageStatus.Value, attrs.LifeCycleStageStatus.Link, attrs.LifeCycleStageStatus.DisplayValue))
	_ = resourceData.Set("mac_address", attrs.MacAddress)
	//_ = resourceData.Set("company", StructToMap(attrs.Company))
	_ = resourceData.Set("company", client.StructToList(attrs.Company.Value, attrs.Company.Link, attrs.Company.DisplayValue))
	_ = resourceData.Set("model_number", attrs.ModelNumber)
	_ = resourceData.Set("justification", attrs.Justification)
	//_ = resourceData.Set("department", StructToMap(attrs.Department))
	_ = resourceData.Set("department", client.StructToList(attrs.Department.Value, attrs.Department.Link, attrs.Department.DisplayValue))
	//_ = resourceData.Set("assigned_to", StructToMap(attrs.AssignedTo))
	_ = resourceData.Set("assigned_to", client.StructToList(attrs.AssignedTo.Value, attrs.AssignedTo.Link, attrs.AssignedTo.DisplayValue))
	_ = resourceData.Set("start_date", attrs.StartDate)
	_ = resourceData.Set("cost", attrs.Cost)
	_ = resourceData.Set("comments", attrs.Comments)
	_ = resourceData.Set("attestation_status", attrs.AttestationStatus)
	//_ = resourceData.Set("cmdb_ot_entity", StructToMap(attrs.CmdbOtEntity))
	_ = resourceData.Set("cmdb_ot_entity", client.StructToList(attrs.CmdbOtEntity.Value, attrs.CmdbOtEntity.Link, attrs.CmdbOtEntity.DisplayValue))
	_ = resourceData.Set("sys_mod_count", attrs.SysModCount)
	_ = resourceData.Set("serial_number", attrs.SerialNumber)
	_ = resourceData.Set("monitor", attrs.Monitor)
	//_ = resourceData.Set("model_id", StructToMap(attrs.ModelId))
	_ = resourceData.Set("model_id", client.StructToList(attrs.ModelId.Value, attrs.ModelId.Link, attrs.ModelId.DisplayValue))
	_ = resourceData.Set("ip_address", attrs.IpAddress)
	//_ = resourceData.Set("duplicate_of", StructToMap(attrs.DuplicateOf))
	_ = resourceData.Set("duplicate_of", client.StructToList(attrs.DuplicateOf.Value, attrs.DuplicateOf.Link, attrs.DuplicateOf.DisplayValue))
	_ = resourceData.Set("sys_tags", attrs.SysTags)
	_ = resourceData.Set("cost_cc", attrs.CostCc)
	//_ = resourceData.Set("support_group", StructToMap(attrs.SupportGroup))
	_ = resourceData.Set("support_group", client.StructToList(attrs.SupportGroup.Value, attrs.SupportGroup.Link, attrs.SupportGroup.DisplayValue))
	_ = resourceData.Set("order_date", attrs.OrderDate)
	//_ = resourceData.Set("schedule", StructToMap(attrs.Schedule))
	_ = resourceData.Set("schedule", client.StructToList(attrs.Schedule.Value, attrs.Schedule.Link, attrs.Schedule.DisplayValue))
	_ = resourceData.Set("environment", attrs.Environment)
	_ = resourceData.Set("due", attrs.Due)
	_ = resourceData.Set("attested", attrs.Attested)
	_ = resourceData.Set("unverified", attrs.Unverified)
	_ = resourceData.Set("correlation_id", attrs.CorrelationId)
	_ = resourceData.Set("attributes", attrs.Attributes)
	//_ = resourceData.Set("location", StructToMap(attrs.Location))
	_ = resourceData.Set("location", client.StructToList(attrs.Location.Value, attrs.Location.Link, attrs.Location.DisplayValue))
	//_ = resourceData.Set("asset", StructToMap(attrs.Asset))
	_ = resourceData.Set("asset", client.StructToList(attrs.Asset.Value, attrs.Asset.Link, attrs.Asset.DisplayValue))
	_ = resourceData.Set("category", attrs.Category)
	_ = resourceData.Set("fault_count", attrs.FaultCount)
	_ = resourceData.Set("lease_id", attrs.LeaseId)
	return nil
}
