package resources

///////////////////////////////////////////////////////////////////////////////////////////////////
//  ****WARNING***    : This is automatically generated code  using the generatprovidersource
//                      utility in this package, any edits to this file will be overwritten
//                      and lost.  Refer to the readme.md in the generateprovidersource directory
//                       for details on how to regenerate this code.
// Provider Version    : 0.01
// Generator Version   : 1.00
//
//  Description       :   This file is the resource provider for the cmdb_ci_personal_printer CMDB Class.  This code is executed
//                        when the servicenowcmdb_cmdb_ci_personal_printer keyword is used in a terraform script (*.tf) file
//
//                        This file is will need to be regenerated if the ServiceNow CMDB base CI Class
//                        "cmdb_ci" or of the cmdb_ci_personal_printer CI Class is modified.
///////////////////////////////////////////////////////////////////////////////////////////////////
import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// This constant stores the name of the CI and is used to construct the URL to ServiceNow
const CiNamecmdb_ci_personal_printer = "cmdb_ci_personal_printer"

// This is the structure to construct the JSON payload when POSTing to ServiceNow.  This is needed because
// ServiceNow has strict parsing on the JSON Data and will fail if the JSON format doesn't match exactly.
// This is essentially the same as CmdbCiPersonalPrinterGet but does not contain the "value", "link" and
// "display_value" fields for reference objects.
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiPersonalPrinterPost struct {
	Result struct {
		Attributes struct {
			SkipSync             string `json:"skip_sync,omitempty"`
			OperationalStatus    string `json:"operational_status,omitempty"`
			ResolutionUnits      string `json:"resolution_units,omitempty"`
			SysUpdatedOn         string `json:"sys_updated_on,omitempty"`
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
			HardwareSubstatus    string `json:"hardware_substatus,omitempty"`
			Fqdn                 string `json:"fqdn,omitempty"`
			ChangeControl        string `json:"change_control,omitempty"`
			OwnedBy              string `json:"owned_by,omitempty"`
			CheckedOut           string `json:"checked_out,omitempty"`
			SysDomainPath        string `json:"sys_domain_path,omitempty"`
			HorizontalResolution string `json:"horizontal_resolution,omitempty"`
			DeliveryDate         string `json:"delivery_date,omitempty"`
			MaintenanceSchedule  string `json:"maintenance_schedule,omitempty"`
			HardwareStatus       string `json:"hardware_status,omitempty"`
			UseUnits             string `json:"use_units,omitempty"`
			InstallStatus        string `json:"install_status,omitempty"`
			CostCenter           string `json:"cost_center,omitempty"`
			SupportedBy          string `json:"supported_by,omitempty"`
			DnsDomain            string `json:"dns_domain,omitempty"`
			Name                 string `json:"name,omitempty"`
			Assigned             string `json:"assigned,omitempty"`
			VerticalResolution   string `json:"vertical_resolution,omitempty"`
			PurchaseDate         string `json:"purchase_date,omitempty"`
			Subcategory          string `json:"subcategory,omitempty"`
			DefaultGateway       string `json:"default_gateway,omitempty"`
			ShortDescription     string `json:"short_description,omitempty"`
			AssignmentGroup      string `json:"assignment_group,omitempty"`
			Color                string `json:"color,omitempty"`
			ManagedBy            string `json:"managed_by,omitempty"`
			Postscript           string `json:"postscript,omitempty"`
			PrintType            string `json:"print_type,omitempty"`
			LastDiscovered       string `json:"last_discovered,omitempty"`
			CanPrint             string `json:"can_print,omitempty"`
			Colors               string `json:"colors,omitempty"`
			Manufacturer         string `json:"manufacturer,omitempty"`
			UseCount             string `json:"use_count,omitempty"`
			PoNumber             string `json:"po_number,omitempty"`
			Paper                string `json:"paper,omitempty"`
			CheckedIn            string `json:"checked_in,omitempty"`
			Vendor               string `json:"vendor,omitempty"`
			MacAddress           string `json:"mac_address,omitempty"`
			Pcl                  string `json:"pcl,omitempty"`
			Company              string `json:"company,omitempty"`
			ModelNumber          string `json:"model_number,omitempty"`
			Justification        string `json:"justification,omitempty"`
			Department           string `json:"department,omitempty"`
			AssignedTo           string `json:"assigned_to,omitempty"`
			StartDate            string `json:"start_date,omitempty"`
			Cost                 string `json:"cost,omitempty"`
			Comments             string `json:"comments,omitempty"`
			SysModCount          string `json:"sys_mod_count,omitempty"`
			SerialNumber         string `json:"serial_number,omitempty"`
			Monitor              string `json:"monitor,omitempty"`
			ModelId              string `json:"model_id,omitempty"`
			IpAddress            string `json:"ip_address,omitempty"`
			Ppm                  string `json:"ppm,omitempty"`
			DuplicateOf          string `json:"duplicate_of,omitempty"`
			SysTags              string `json:"sys_tags,omitempty"`
			CostCc               string `json:"cost_cc,omitempty"`
			SupportGroup         string `json:"support_group,omitempty"`
			OrderDate            string `json:"order_date,omitempty"`
			Schedule             string `json:"schedule,omitempty"`
			Due                  string `json:"due,omitempty"`
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
//
type CmdbCiPersonalPrinterGet struct {
	Result struct {
		Attributes struct {
			SkipSync           string `json:"skip_sync,omitempty"`
			OperationalStatus  string `json:"operational_status,omitempty"`
			ResolutionUnits    string `json:"resolution_units,omitempty"`
			SysUpdatedOn       string `json:"sys_updated_on,omitempty"`
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
			HardwareSubstatus  string `json:"hardware_substatus,omitempty"`
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
			CheckedOut           string `json:"checked_out,omitempty"`
			SysDomainPath        string `json:"sys_domain_path,omitempty"`
			HorizontalResolution string `json:"horizontal_resolution,omitempty"`
			DeliveryDate         string `json:"delivery_date,omitempty"`
			MaintenanceSchedule  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"maintenance_schedule"`
			HardwareStatus string `json:"hardware_status,omitempty"`
			UseUnits       string `json:"use_units,omitempty"`
			InstallStatus  string `json:"install_status,omitempty"`
			CostCenter     struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"cost_center"`
			SupportedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"supported_by"`
			DnsDomain          string `json:"dns_domain,omitempty"`
			Name               string `json:"name,omitempty"`
			Assigned           string `json:"assigned,omitempty"`
			VerticalResolution string `json:"vertical_resolution,omitempty"`
			PurchaseDate       string `json:"purchase_date,omitempty"`
			Subcategory        string `json:"subcategory,omitempty"`
			DefaultGateway     string `json:"default_gateway,omitempty"`
			ShortDescription   string `json:"short_description,omitempty"`
			AssignmentGroup    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"assignment_group"`
			Color     string `json:"color,omitempty"`
			ManagedBy struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"managed_by"`
			Postscript     string `json:"postscript,omitempty"`
			PrintType      string `json:"print_type,omitempty"`
			LastDiscovered string `json:"last_discovered,omitempty"`
			CanPrint       string `json:"can_print,omitempty"`
			Colors         string `json:"colors,omitempty"`
			Manufacturer   struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"manufacturer"`
			UseCount  string `json:"use_count,omitempty"`
			PoNumber  string `json:"po_number,omitempty"`
			Paper     string `json:"paper,omitempty"`
			CheckedIn string `json:"checked_in,omitempty"`
			Vendor    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"vendor"`
			MacAddress string `json:"mac_address,omitempty"`
			Pcl        string `json:"pcl,omitempty"`
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
			StartDate    string `json:"start_date,omitempty"`
			Cost         string `json:"cost,omitempty"`
			Comments     string `json:"comments,omitempty"`
			SysModCount  string `json:"sys_mod_count,omitempty"`
			SerialNumber string `json:"serial_number,omitempty"`
			Monitor      string `json:"monitor,omitempty"`
			ModelId      struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"model_id"`
			IpAddress   string `json:"ip_address,omitempty"`
			Ppm         string `json:"ppm,omitempty"`
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
			Due           string `json:"due,omitempty"`
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

func ResourceCmdbCiPersonalPrinter() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCmdbCiPersonalPrinter,
		Read:   readResourceCmdbCiPersonalPrinter,
		Update: updateResourceCmdbCiPersonalPrinter,
		Delete: deleteResourceCmdbCiPersonalPrinter,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"skip_sync": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operational_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"resolution_units": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sys_updated_on": {
				Type:     schema.TypeString,
				Computed: true,
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
			"hardware_substatus": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"change_control": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Optional: true,
			},
			"horizontal_resolution": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"delivery_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_schedule": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"hardware_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_units": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"install_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cost_center": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"vertical_resolution": {
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
			"default_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"short_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assignment_group": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"color": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"managed_by": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"postscript": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"print_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"last_discovered": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"can_print": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"colors": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"manufacturer": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"use_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"po_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"paper": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"checked_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vendor": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"pcl": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"company": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Optional: true,
			},
			"model_id": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"ppm": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"duplicate_of": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Optional: true,
			},
			"support_group": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
			"due": {
				Type:     schema.TypeString,
				Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"display_value": {
							Type:     schema.TypeString,
							Computed: true,
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
				Optional: true,
			},
			"lease_id": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//  Create routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func createResourceCmdbCiPersonalPrinter(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client) //Client Connection details
	// Use common function to update base attributes
	var ci CmdbCiPersonalPrinterPost
	if err := copyFromTerraformToServiceNowCmdbCiPersonalPrinter(resourceData, &ci); err != nil {
		return err
	}

	//t := time.Now()
	//ci.Result.Attributes.Installdate = t.String()
	//ci.Result.Attributes.SysCreatedOn = t.String()

	// Using "sysparm fields=sys_id" reduces the amount of data received, we only need the sys_id for create action
	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_personal_printer + "?sysparm_fields=sys_id"

	jsonData, err := servicenowClient.RequestJSON("POST", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return err
	}

	resourceData.SetId(GetSysId(jsonData))
	return readResourceCmdbCiPersonalPrinter(resourceData, serviceNowClient)
}

//  Read routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func readResourceCmdbCiPersonalPrinter(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client)
	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_personal_printer + "/" + resourceData.Id()
	var jsonData []byte
	jsonData, err := servicenowClient.RequestJSON("GET", SnowUrl, nil)
	if err != nil {
		resourceData.SetId("")
		return err
	}

	if err := copyFromServiceNowToTerraformCmdbCiPersonalPrinter(resourceData, jsonData); err != nil {
		return err
	}

	return nil
}

//  Update routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func updateResourceCmdbCiPersonalPrinter(resourceData *schema.ResourceData, serviceNowClient interface{}) error {
	servicenowClient := serviceNowClient.(*Client)

	var ci CmdbCiPersonalPrinterPost
	if err := copyFromTerraformToServiceNowCmdbCiPersonalPrinter(resourceData, &ci); err != nil {
		return err
	}

	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_personal_printer + "/" + resourceData.Id()
	//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return err
	}
	return readResourceCmdbCiPersonalPrinter(resourceData, serviceNowClient)
}

// TODO:  Need to work out what to do with deleting CIs. ServiceNow does not support deleting CIs via the API
//        CIs shouldn't be deleted because may be referenced by other objects in the ServiceNow database
//        (eg changes and incidents).  Therefore the best practice is to change their operational and
//        install status.  The tricky bit is working out how to make this flexible for everyone to use.

//  Delete routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func deleteResourceCmdbCiPersonalPrinter(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client)
	var ci CmdbCiPersonalPrinterPost
	if err := copyFromTerraformToServiceNowCmdbCiPersonalPrinter(resourceData, &ci); err != nil {
		return err
	}

	if err := resourceData.Set("install_status", "retired"); err != nil {
		return fmt.Errorf("CmdbCiPersonalPrinterfailed to set install_status field during destroy action %s", err)
	}

	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_personal_printer + "/" + resourceData.Id()
	//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return err
	}
	return nil
}

// This routine is called to copy data from the Terraform schema.ResourceData object into the CMDB CI object
// It would be nice if Terraform implemented a funciton to return a list of field names in a slice, this would
// make it easier to loop through the structure instead of doing a "Get" per field.
//
func copyFromTerraformToServiceNowCmdbCiPersonalPrinter(resourceData *schema.ResourceData, ci *CmdbCiPersonalPrinterPost) error {

	attrs := &ci.Result.Attributes
	attrs.SkipSync = resourceData.Get("skip_sync").(string)
	attrs.OperationalStatus = resourceData.Get("operational_status").(string)
	attrs.ResolutionUnits = resourceData.Get("resolution_units").(string)
	attrs.SysUpdatedOn = resourceData.Get("sys_updated_on").(string)
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
	attrs.HardwareSubstatus = resourceData.Get("hardware_substatus").(string)
	attrs.Fqdn = resourceData.Get("fqdn").(string)
	attrs.ChangeControl = resourceData.Get("change_control.value").(string)
	attrs.OwnedBy = resourceData.Get("owned_by.value").(string)
	attrs.CheckedOut = resourceData.Get("checked_out").(string)
	attrs.SysDomainPath = resourceData.Get("sys_domain_path").(string)
	attrs.HorizontalResolution = resourceData.Get("horizontal_resolution").(string)
	attrs.DeliveryDate = resourceData.Get("delivery_date").(string)
	attrs.MaintenanceSchedule = resourceData.Get("maintenance_schedule.value").(string)
	attrs.HardwareStatus = resourceData.Get("hardware_status").(string)
	attrs.UseUnits = resourceData.Get("use_units").(string)
	attrs.InstallStatus = resourceData.Get("install_status").(string)
	attrs.CostCenter = resourceData.Get("cost_center.value").(string)
	attrs.SupportedBy = resourceData.Get("supported_by.value").(string)
	attrs.DnsDomain = resourceData.Get("dns_domain").(string)
	attrs.Name = resourceData.Get("name").(string)
	attrs.Assigned = resourceData.Get("assigned").(string)
	attrs.VerticalResolution = resourceData.Get("vertical_resolution").(string)
	attrs.PurchaseDate = resourceData.Get("purchase_date").(string)
	attrs.Subcategory = resourceData.Get("subcategory").(string)
	attrs.DefaultGateway = resourceData.Get("default_gateway").(string)
	attrs.ShortDescription = resourceData.Get("short_description").(string)
	attrs.AssignmentGroup = resourceData.Get("assignment_group.value").(string)
	attrs.Color = resourceData.Get("color").(string)
	attrs.ManagedBy = resourceData.Get("managed_by.value").(string)
	attrs.Postscript = resourceData.Get("postscript").(string)
	attrs.PrintType = resourceData.Get("print_type").(string)
	attrs.LastDiscovered = resourceData.Get("last_discovered").(string)
	attrs.CanPrint = resourceData.Get("can_print").(string)
	attrs.Colors = resourceData.Get("colors").(string)
	attrs.Manufacturer = resourceData.Get("manufacturer.value").(string)
	attrs.UseCount = resourceData.Get("use_count").(string)
	attrs.PoNumber = resourceData.Get("po_number").(string)
	attrs.Paper = resourceData.Get("paper").(string)
	attrs.CheckedIn = resourceData.Get("checked_in").(string)
	attrs.Vendor = resourceData.Get("vendor.value").(string)
	attrs.MacAddress = resourceData.Get("mac_address").(string)
	attrs.Pcl = resourceData.Get("pcl").(string)
	attrs.Company = resourceData.Get("company.value").(string)
	attrs.ModelNumber = resourceData.Get("model_number").(string)
	attrs.Justification = resourceData.Get("justification").(string)
	attrs.Department = resourceData.Get("department.value").(string)
	attrs.AssignedTo = resourceData.Get("assigned_to.value").(string)
	attrs.StartDate = resourceData.Get("start_date").(string)
	attrs.Cost = resourceData.Get("cost").(string)
	attrs.Comments = resourceData.Get("comments").(string)
	attrs.SysModCount = resourceData.Get("sys_mod_count").(string)
	attrs.SerialNumber = resourceData.Get("serial_number").(string)
	attrs.Monitor = resourceData.Get("monitor").(string)
	attrs.ModelId = resourceData.Get("model_id.value").(string)
	attrs.IpAddress = resourceData.Get("ip_address").(string)
	attrs.Ppm = resourceData.Get("ppm").(string)
	attrs.DuplicateOf = resourceData.Get("duplicate_of.value").(string)
	attrs.SysTags = resourceData.Get("sys_tags").(string)
	attrs.CostCc = resourceData.Get("cost_cc").(string)
	attrs.SupportGroup = resourceData.Get("support_group.value").(string)
	attrs.OrderDate = resourceData.Get("order_date").(string)
	attrs.Schedule = resourceData.Get("schedule.value").(string)
	attrs.Due = resourceData.Get("due").(string)
	attrs.Unverified = resourceData.Get("unverified").(string)
	attrs.CorrelationId = resourceData.Get("correlation_id").(string)
	attrs.Attributes = resourceData.Get("attributes").(string)
	attrs.Location = resourceData.Get("location.value").(string)
	attrs.Asset = resourceData.Get("asset.value").(string)
	attrs.Category = resourceData.Get("category").(string)
	attrs.FaultCount = resourceData.Get("fault_count").(string)
	attrs.LeaseId = resourceData.Get("lease_id").(string)
	return nil
}

// This function Unmarshal's the JSON payload from ServiceNow into a corresponding STRUCT for the CI.  It then
// Sets each field in the Terraform schema.ResourceData object with the corrsponding field from the CI STRUCT.
//
// NOTE:  Terraform expects a MAP for fields that represent "reference fields" so that the "value", "link" and
//        "display_value" are decoded correctly by Terraform.  The map is constructed for each reference field
//        using a common function called "StructToMap" in the "client_base.go" file.
//
func copyFromServiceNowToTerraformCmdbCiPersonalPrinter(resourceData *schema.ResourceData, jsonData []byte) error {
	ci := CmdbCiPersonalPrinterGet{}
	if err := json.Unmarshal(jsonData, &ci); err != nil {
		//resourceData.SetId("")
		//return err
	}

	resourceData.SetId(GetSysId(jsonData))
	attrs := ci.Result.Attributes
	_ = resourceData.Set("skip_sync", attrs.SkipSync)
	_ = resourceData.Set("operational_status", attrs.OperationalStatus)
	_ = resourceData.Set("resolution_units", attrs.ResolutionUnits)
	_ = resourceData.Set("sys_updated_on", attrs.SysUpdatedOn)
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
	_ = resourceData.Set("hardware_substatus", attrs.HardwareSubstatus)
	_ = resourceData.Set("fqdn", attrs.Fqdn)
	_ = resourceData.Set("change_control", StructToMap(attrs.ChangeControl))
	_ = resourceData.Set("owned_by", StructToMap(attrs.OwnedBy))
	_ = resourceData.Set("checked_out", attrs.CheckedOut)
	_ = resourceData.Set("sys_domain_path", attrs.SysDomainPath)
	_ = resourceData.Set("horizontal_resolution", attrs.HorizontalResolution)
	_ = resourceData.Set("delivery_date", attrs.DeliveryDate)
	_ = resourceData.Set("maintenance_schedule", StructToMap(attrs.MaintenanceSchedule))
	_ = resourceData.Set("hardware_status", attrs.HardwareStatus)
	_ = resourceData.Set("use_units", attrs.UseUnits)
	_ = resourceData.Set("install_status", attrs.InstallStatus)
	_ = resourceData.Set("cost_center", StructToMap(attrs.CostCenter))
	_ = resourceData.Set("supported_by", StructToMap(attrs.SupportedBy))
	_ = resourceData.Set("dns_domain", attrs.DnsDomain)
	_ = resourceData.Set("name", attrs.Name)
	_ = resourceData.Set("assigned", attrs.Assigned)
	_ = resourceData.Set("vertical_resolution", attrs.VerticalResolution)
	_ = resourceData.Set("purchase_date", attrs.PurchaseDate)
	_ = resourceData.Set("subcategory", attrs.Subcategory)
	_ = resourceData.Set("default_gateway", attrs.DefaultGateway)
	_ = resourceData.Set("short_description", attrs.ShortDescription)
	_ = resourceData.Set("assignment_group", StructToMap(attrs.AssignmentGroup))
	_ = resourceData.Set("color", attrs.Color)
	_ = resourceData.Set("managed_by", StructToMap(attrs.ManagedBy))
	_ = resourceData.Set("postscript", attrs.Postscript)
	_ = resourceData.Set("print_type", attrs.PrintType)
	_ = resourceData.Set("last_discovered", attrs.LastDiscovered)
	_ = resourceData.Set("can_print", attrs.CanPrint)
	_ = resourceData.Set("colors", attrs.Colors)
	_ = resourceData.Set("manufacturer", StructToMap(attrs.Manufacturer))
	_ = resourceData.Set("use_count", attrs.UseCount)
	_ = resourceData.Set("po_number", attrs.PoNumber)
	_ = resourceData.Set("paper", attrs.Paper)
	_ = resourceData.Set("checked_in", attrs.CheckedIn)
	_ = resourceData.Set("vendor", StructToMap(attrs.Vendor))
	_ = resourceData.Set("mac_address", attrs.MacAddress)
	_ = resourceData.Set("pcl", attrs.Pcl)
	_ = resourceData.Set("company", StructToMap(attrs.Company))
	_ = resourceData.Set("model_number", attrs.ModelNumber)
	_ = resourceData.Set("justification", attrs.Justification)
	_ = resourceData.Set("department", StructToMap(attrs.Department))
	_ = resourceData.Set("assigned_to", StructToMap(attrs.AssignedTo))
	_ = resourceData.Set("start_date", attrs.StartDate)
	_ = resourceData.Set("cost", attrs.Cost)
	_ = resourceData.Set("comments", attrs.Comments)
	_ = resourceData.Set("sys_mod_count", attrs.SysModCount)
	_ = resourceData.Set("serial_number", attrs.SerialNumber)
	_ = resourceData.Set("monitor", attrs.Monitor)
	_ = resourceData.Set("model_id", StructToMap(attrs.ModelId))
	_ = resourceData.Set("ip_address", attrs.IpAddress)
	_ = resourceData.Set("ppm", attrs.Ppm)
	_ = resourceData.Set("duplicate_of", StructToMap(attrs.DuplicateOf))
	_ = resourceData.Set("sys_tags", attrs.SysTags)
	_ = resourceData.Set("cost_cc", attrs.CostCc)
	_ = resourceData.Set("support_group", StructToMap(attrs.SupportGroup))
	_ = resourceData.Set("order_date", attrs.OrderDate)
	_ = resourceData.Set("schedule", StructToMap(attrs.Schedule))
	_ = resourceData.Set("due", attrs.Due)
	_ = resourceData.Set("unverified", attrs.Unverified)
	_ = resourceData.Set("correlation_id", attrs.CorrelationId)
	_ = resourceData.Set("attributes", attrs.Attributes)
	_ = resourceData.Set("location", StructToMap(attrs.Location))
	_ = resourceData.Set("asset", StructToMap(attrs.Asset))
	_ = resourceData.Set("category", attrs.Category)
	_ = resourceData.Set("fault_count", attrs.FaultCount)
	_ = resourceData.Set("lease_id", attrs.LeaseId)
	return nil
}
