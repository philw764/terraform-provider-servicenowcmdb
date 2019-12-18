package resources

///////////////////////////////////////////////////////////////////////////////////////////////////
//  ****WARNING***    : This is automatically generated code  using the generatprovidersource
//                      utility in this package, any edits to this file will be overwritten
//                      and lost.  Refer to the readme.md in the generateprovidersource directory
//                       for details on how to regenerate this code.
// Provider Version    : 0.01
// Generator Version   : 1.00
//
//  Description       :   This file is the resource provider for the cmdb_ci_solaris_server CMDB Class.  This code is executed
//                        when the servicenowcmdb_cmdb_ci_solaris_server keyword is used in a terraform script (*.tf) file
//
//                        This file is will need to be regenerated if the ServiceNow CMDB base CI Class
//                        "cmdb_ci" or of the cmdb_ci_solaris_server CI Class is modified.
///////////////////////////////////////////////////////////////////////////////////////////////////
import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// This constant stores the name of the CI and is used to construct the URL to ServiceNow
const CiNamecmdb_ci_solaris_server = "cmdb_ci_solaris_server"

// This is the structure to construct the JSON payload when POSTing to ServiceNow.  This is needed because
// ServiceNow has strict parsing on the JSON Data and will fail if the JSON format doesn't match exactly.
// This is essentially the same as CmdbCiSolarisServerGet but does not contain the "value", "link" and
// "display_value" fields for reference objects.
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiSolarisServerPost struct {
	Result struct {
		Attributes struct {
			OsAddressWidth      string `json:"os_address_width,omitempty"`
			FirewallStatus      string `json:"firewall_status,omitempty"`
			OperationalStatus   string `json:"operational_status,omitempty"`
			OsServicePack       string `json:"os_service_pack,omitempty"`
			CpuCoreThread       string `json:"cpu_core_thread,omitempty"`
			CpuManufacturer     string `json:"cpu_manufacturer,omitempty"`
			SysUpdatedOn        string `json:"sys_updated_on,omitempty"`
			DiscoverySource     string `json:"discovery_source,omitempty"`
			FirstDiscovered     string `json:"first_discovered,omitempty"`
			DueIn               string `json:"due_in,omitempty"`
			UsedFor             string `json:"used_for,omitempty"`
			InvoiceNumber       string `json:"invoice_number,omitempty"`
			GlAccount           string `json:"gl_account,omitempty"`
			SysCreatedBy        string `json:"sys_created_by,omitempty"`
			WarrantyExpiration  string `json:"warranty_expiration,omitempty"`
			Ram                 string `json:"ram,omitempty"`
			CpuName             string `json:"cpu_name,omitempty"`
			CpuSpeed            string `json:"cpu_speed,omitempty"`
			OwnedBy             string `json:"owned_by,omitempty"`
			CheckedOut          string `json:"checked_out,omitempty"`
			SysDomainPath       string `json:"sys_domain_path,omitempty"`
			DiskSpace           string `json:"disk_space,omitempty"`
			Classification      string `json:"classification,omitempty"`
			ObjectId            string `json:"object_id,omitempty"`
			MaintenanceSchedule string `json:"maintenance_schedule,omitempty"`
			CostCenter          string `json:"cost_center,omitempty"`
			DnsDomain           string `json:"dns_domain,omitempty"`
			Assigned            string `json:"assigned,omitempty"`
			PurchaseDate        string `json:"purchase_date,omitempty"`
			ShortDescription    string `json:"short_description,omitempty"`
			CdSpeed             string `json:"cd_speed,omitempty"`
			Floppy              string `json:"floppy,omitempty"`
			ManagedBy           string `json:"managed_by,omitempty"`
			OsDomain            string `json:"os_domain,omitempty"`
			LastDiscovered      string `json:"last_discovered,omitempty"`
			CanPrint            string `json:"can_print,omitempty"`
			Manufacturer        string `json:"manufacturer,omitempty"`
			CpuCount            string `json:"cpu_count,omitempty"`
			Vendor              string `json:"vendor,omitempty"`
			ModelNumber         string `json:"model_number,omitempty"`
			AssignedTo          string `json:"assigned_to,omitempty"`
			StartDate           string `json:"start_date,omitempty"`
			OsVersion           string `json:"os_version,omitempty"`
			SerialNumber        string `json:"serial_number,omitempty"`
			CdRom               string `json:"cd_rom,omitempty"`
			SupportGroup        string `json:"support_group,omitempty"`
			Unverified          string `json:"unverified,omitempty"`
			CorrelationId       string `json:"correlation_id,omitempty"`
			Attributes          string `json:"attributes,omitempty"`
			Asset               string `json:"asset,omitempty"`
			FormFactor          string `json:"form_factor,omitempty"`
			CpuCoreCount        string `json:"cpu_core_count,omitempty"`
			SkipSync            string `json:"skip_sync,omitempty"`
			SysUpdatedBy        string `json:"sys_updated_by,omitempty"`
			SysCreatedOn        string `json:"sys_created_on,omitempty"`
			CpuType             string `json:"cpu_type,omitempty"`
			InstallDate         string `json:"install_date,omitempty"`
			AssetTag            string `json:"asset_tag,omitempty"`
			DrBackup            string `json:"dr_backup,omitempty"`
			HardwareSubstatus   string `json:"hardware_substatus,omitempty"`
			Fqdn                string `json:"fqdn,omitempty"`
			ChangeControl       string `json:"change_control,omitempty"`
			DeliveryDate        string `json:"delivery_date,omitempty"`
			HardwareStatus      string `json:"hardware_status,omitempty"`
			InstallStatus       string `json:"install_status,omitempty"`
			SupportedBy         string `json:"supported_by,omitempty"`
			Name                string `json:"name,omitempty"`
			Subcategory         string `json:"subcategory,omitempty"`
			DefaultGateway      string `json:"default_gateway,omitempty"`
			ChassisType         string `json:"chassis_type,omitempty"`
			Virtual             string `json:"virtual,omitempty"`
			AssignmentGroup     string `json:"assignment_group,omitempty"`
			PoNumber            string `json:"po_number,omitempty"`
			CheckedIn           string `json:"checked_in,omitempty"`
			MacAddress          string `json:"mac_address,omitempty"`
			Company             string `json:"company,omitempty"`
			Justification       string `json:"justification,omitempty"`
			Department          string `json:"department,omitempty"`
			Cost                string `json:"cost,omitempty"`
			Comments            string `json:"comments,omitempty"`
			Os                  string `json:"os,omitempty"`
			SysModCount         string `json:"sys_mod_count,omitempty"`
			Monitor             string `json:"monitor,omitempty"`
			ModelId             string `json:"model_id,omitempty"`
			IpAddress           string `json:"ip_address,omitempty"`
			DuplicateOf         string `json:"duplicate_of,omitempty"`
			SysTags             string `json:"sys_tags,omitempty"`
			CostCc              string `json:"cost_cc,omitempty"`
			OrderDate           string `json:"order_date,omitempty"`
			Schedule            string `json:"schedule,omitempty"`
			Due                 string `json:"due,omitempty"`
			Location            string `json:"location,omitempty"`
			Category            string `json:"category,omitempty"`
			FaultCount          string `json:"fault_count,omitempty"`
			LeaseId             string `json:"lease_id,omitempty"`
			HostName            string `json:"host_name,omitempty"`
		} `json:"attributes"`
	} `json:"result"`
}

// This is the structure where JSON data is Unmarshal'ed from ServiceNow.  This is different from the POST
// structure because it has to contain the "value", "link" and "display_value" fields for reference objects
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiSolarisServerGet struct {
	Result struct {
		Attributes struct {
			OsAddressWidth    string `json:"os_address_width,omitempty"`
			FirewallStatus    string `json:"firewall_status,omitempty"`
			OperationalStatus string `json:"operational_status,omitempty"`
			OsServicePack     string `json:"os_service_pack,omitempty"`
			CpuCoreThread     string `json:"cpu_core_thread,omitempty"`
			CpuManufacturer   struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"cpu_manufacturer"`
			SysUpdatedOn       string `json:"sys_updated_on,omitempty"`
			DiscoverySource    string `json:"discovery_source,omitempty"`
			FirstDiscovered    string `json:"first_discovered,omitempty"`
			DueIn              string `json:"due_in,omitempty"`
			UsedFor            string `json:"used_for,omitempty"`
			InvoiceNumber      string `json:"invoice_number,omitempty"`
			GlAccount          string `json:"gl_account,omitempty"`
			SysCreatedBy       string `json:"sys_created_by,omitempty"`
			WarrantyExpiration string `json:"warranty_expiration,omitempty"`
			Ram                string `json:"ram,omitempty"`
			CpuName            string `json:"cpu_name,omitempty"`
			CpuSpeed           string `json:"cpu_speed,omitempty"`
			OwnedBy            struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"owned_by"`
			CheckedOut          string `json:"checked_out,omitempty"`
			SysDomainPath       string `json:"sys_domain_path,omitempty"`
			DiskSpace           string `json:"disk_space,omitempty"`
			Classification      string `json:"classification,omitempty"`
			ObjectId            string `json:"object_id,omitempty"`
			MaintenanceSchedule struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"maintenance_schedule"`
			CostCenter struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"cost_center"`
			DnsDomain        string `json:"dns_domain,omitempty"`
			Assigned         string `json:"assigned,omitempty"`
			PurchaseDate     string `json:"purchase_date,omitempty"`
			ShortDescription string `json:"short_description,omitempty"`
			CdSpeed          string `json:"cd_speed,omitempty"`
			Floppy           string `json:"floppy,omitempty"`
			ManagedBy        struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"managed_by"`
			OsDomain       string `json:"os_domain,omitempty"`
			LastDiscovered string `json:"last_discovered,omitempty"`
			CanPrint       string `json:"can_print,omitempty"`
			Manufacturer   struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"manufacturer"`
			CpuCount string `json:"cpu_count,omitempty"`
			Vendor   struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"vendor"`
			ModelNumber string `json:"model_number,omitempty"`
			AssignedTo  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"assigned_to"`
			StartDate    string `json:"start_date,omitempty"`
			OsVersion    string `json:"os_version,omitempty"`
			SerialNumber string `json:"serial_number,omitempty"`
			CdRom        string `json:"cd_rom,omitempty"`
			SupportGroup struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"support_group"`
			Unverified    string `json:"unverified,omitempty"`
			CorrelationId string `json:"correlation_id,omitempty"`
			Attributes    string `json:"attributes,omitempty"`
			Asset         struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"asset"`
			FormFactor   string `json:"form_factor,omitempty"`
			CpuCoreCount string `json:"cpu_core_count,omitempty"`
			SkipSync     string `json:"skip_sync,omitempty"`
			SysUpdatedBy string `json:"sys_updated_by,omitempty"`
			SysCreatedOn string `json:"sys_created_on,omitempty"`
			CpuType      string `json:"cpu_type,omitempty"`
			InstallDate  string `json:"install_date,omitempty"`
			AssetTag     string `json:"asset_tag,omitempty"`
			DrBackup     struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"dr_backup"`
			HardwareSubstatus string `json:"hardware_substatus,omitempty"`
			Fqdn              string `json:"fqdn,omitempty"`
			ChangeControl     struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"change_control"`
			DeliveryDate   string `json:"delivery_date,omitempty"`
			HardwareStatus string `json:"hardware_status,omitempty"`
			InstallStatus  string `json:"install_status,omitempty"`
			SupportedBy    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"supported_by"`
			Name            string `json:"name,omitempty"`
			Subcategory     string `json:"subcategory,omitempty"`
			DefaultGateway  string `json:"default_gateway,omitempty"`
			ChassisType     string `json:"chassis_type,omitempty"`
			Virtual         string `json:"virtual,omitempty"`
			AssignmentGroup struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"assignment_group"`
			PoNumber   string `json:"po_number,omitempty"`
			CheckedIn  string `json:"checked_in,omitempty"`
			MacAddress string `json:"mac_address,omitempty"`
			Company    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"company"`
			Justification string `json:"justification,omitempty"`
			Department    struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"department"`
			Cost        string `json:"cost,omitempty"`
			Comments    string `json:"comments,omitempty"`
			Os          string `json:"os,omitempty"`
			SysModCount string `json:"sys_mod_count,omitempty"`
			Monitor     string `json:"monitor,omitempty"`
			ModelId     struct {
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
			SysTags   string `json:"sys_tags,omitempty"`
			CostCc    string `json:"cost_cc,omitempty"`
			OrderDate string `json:"order_date,omitempty"`
			Schedule  struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"schedule"`
			Due      string `json:"due,omitempty"`
			Location struct {
				Value        string `json:"value,omitempty"`
				Link         string `json:"link,omitempty"`
				DisplayValue string `json:"display_value,omitempty"`
			} `json:"location"`
			Category   string `json:"category,omitempty"`
			FaultCount string `json:"fault_count,omitempty"`
			LeaseId    string `json:"lease_id,omitempty"`
			HostName   string `json:"host_name,omitempty"`
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

func ResourceCmdbCiSolarisServer() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCmdbCiSolarisServer,
		Read:   readResourceCmdbCiSolarisServer,
		Update: updateResourceCmdbCiSolarisServer,
		Delete: deleteResourceCmdbCiSolarisServer,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"os_address_width": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"firewall_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"operational_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"os_service_pack": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_core_thread": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_manufacturer": {
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
			"due_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"used_for": {
				Type:     schema.TypeString,
				Optional: true,
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
			"ram": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_speed": {
				Type:     schema.TypeString,
				Optional: true,
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
			"disk_space": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"classification": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_id": {
				Type:     schema.TypeString,
				Optional: true,
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
			"dns_domain": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"assigned": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"purchase_date": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"short_description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cd_speed": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"floppy": {
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
			"os_domain": {
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
			"cpu_count": {
				Type:     schema.TypeString,
				Optional: true,
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
			"model_number": {
				Type:     schema.TypeString,
				Computed: true,
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
			"os_version": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cd_rom": {
				Type:     schema.TypeString,
				Optional: true, Default: false,
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
			"form_factor": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"cpu_core_count": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"skip_sync": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_updated_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sys_created_on": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cpu_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"install_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"asset_tag": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"dr_backup": {
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
			"delivery_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardware_status": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"install_status": {
				Type:     schema.TypeString,
				Optional: true,
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
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"subcategory": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_gateway": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"chassis_type": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual": {
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
			"po_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"checked_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
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
			"cost": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"comments": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"os": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"sys_mod_count": {
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
			"host_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

//  Create routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func createResourceCmdbCiSolarisServer(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client) //Client Connection details
	// Use common function to update base attributes
	var ci CmdbCiSolarisServerPost
	if err := copyFromTerraformToServiceNowCmdbCiSolarisServer(resourceData, &ci); err != nil {
		return err
	}

	//t := time.Now()
	//ci.Result.Attributes.Installdate = t.String()
	//ci.Result.Attributes.SysCreatedOn = t.String()

	// Using "sysparm fields=sys_id" reduces the amount of data received, we only need the sys_id for create action
	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_solaris_server + "?sysparm_fields=sys_id"

	jsonData, err := servicenowClient.RequestJSON("POST", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return err
	}

	resourceData.SetId(GetSysId(jsonData))
	return readResourceCmdbCiSolarisServer(resourceData, serviceNowClient)
}

//  Read routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func readResourceCmdbCiSolarisServer(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client)
	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_solaris_server + "/" + resourceData.Id()
	var jsonData []byte
	jsonData, err := servicenowClient.RequestJSON("GET", SnowUrl, nil)
	if err != nil {
		resourceData.SetId("")
		return err
	}

	if err := copyFromServiceNowToTerraformCmdbCiSolarisServer(resourceData, jsonData); err != nil {
		return err
	}

	return nil
}

//  Update routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func updateResourceCmdbCiSolarisServer(resourceData *schema.ResourceData, serviceNowClient interface{}) error {
	servicenowClient := serviceNowClient.(*Client)

	var ci CmdbCiSolarisServerPost
	if err := copyFromTerraformToServiceNowCmdbCiSolarisServer(resourceData, &ci); err != nil {
		return err
	}

	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_solaris_server + "/" + resourceData.Id()
	//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
	if err != nil {
		resourceData.SetId("")
		return err
	}
	return readResourceCmdbCiSolarisServer(resourceData, serviceNowClient)
}

// TODO:  Need to work out what to do with deleting CIs. ServiceNow does not support deleting CIs via the API
//        CIs shouldn't be deleted because may be referenced by other objects in the ServiceNow database
//        (eg changes and incidents).  Therefore the best practice is to change their operational and
//        install status.  The tricky bit is working out how to make this flexible for everyone to use.

//  Delete routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func deleteResourceCmdbCiSolarisServer(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

	servicenowClient := serviceNowClient.(*Client)
	var ci CmdbCiSolarisServerPost
	if err := copyFromTerraformToServiceNowCmdbCiSolarisServer(resourceData, &ci); err != nil {
		return err
	}

	if err := resourceData.Set("install_status", "retired"); err != nil {
		return fmt.Errorf("CmdbCiSolarisServerfailed to set install_status field during destroy action %s", err)
	}

	SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_solaris_server + "/" + resourceData.Id()
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
func copyFromTerraformToServiceNowCmdbCiSolarisServer(resourceData *schema.ResourceData, ci *CmdbCiSolarisServerPost) error {

	attrs := &ci.Result.Attributes
	attrs.OsAddressWidth = resourceData.Get("os_address_width").(string)
	attrs.FirewallStatus = resourceData.Get("firewall_status").(string)
	attrs.OperationalStatus = resourceData.Get("operational_status").(string)
	attrs.OsServicePack = resourceData.Get("os_service_pack").(string)
	attrs.CpuCoreThread = resourceData.Get("cpu_core_thread").(string)
	attrs.CpuManufacturer = resourceData.Get("cpu_manufacturer.value").(string)
	attrs.SysUpdatedOn = resourceData.Get("sys_updated_on").(string)
	attrs.DiscoverySource = resourceData.Get("discovery_source").(string)
	attrs.FirstDiscovered = resourceData.Get("first_discovered").(string)
	attrs.DueIn = resourceData.Get("due_in").(string)
	attrs.UsedFor = resourceData.Get("used_for").(string)
	attrs.InvoiceNumber = resourceData.Get("invoice_number").(string)
	attrs.GlAccount = resourceData.Get("gl_account").(string)
	attrs.SysCreatedBy = resourceData.Get("sys_created_by").(string)
	attrs.WarrantyExpiration = resourceData.Get("warranty_expiration").(string)
	attrs.Ram = resourceData.Get("ram").(string)
	attrs.CpuName = resourceData.Get("cpu_name").(string)
	attrs.CpuSpeed = resourceData.Get("cpu_speed").(string)
	attrs.OwnedBy = resourceData.Get("owned_by.value").(string)
	attrs.CheckedOut = resourceData.Get("checked_out").(string)
	attrs.SysDomainPath = resourceData.Get("sys_domain_path").(string)
	attrs.DiskSpace = resourceData.Get("disk_space").(string)
	attrs.Classification = resourceData.Get("classification").(string)
	attrs.ObjectId = resourceData.Get("object_id").(string)
	attrs.MaintenanceSchedule = resourceData.Get("maintenance_schedule.value").(string)
	attrs.CostCenter = resourceData.Get("cost_center.value").(string)
	attrs.DnsDomain = resourceData.Get("dns_domain").(string)
	attrs.Assigned = resourceData.Get("assigned").(string)
	attrs.PurchaseDate = resourceData.Get("purchase_date").(string)
	attrs.ShortDescription = resourceData.Get("short_description").(string)
	attrs.CdSpeed = resourceData.Get("cd_speed").(string)
	attrs.Floppy = resourceData.Get("floppy").(string)
	attrs.ManagedBy = resourceData.Get("managed_by.value").(string)
	attrs.OsDomain = resourceData.Get("os_domain").(string)
	attrs.LastDiscovered = resourceData.Get("last_discovered").(string)
	attrs.CanPrint = resourceData.Get("can_print").(string)
	attrs.Manufacturer = resourceData.Get("manufacturer.value").(string)
	attrs.CpuCount = resourceData.Get("cpu_count").(string)
	attrs.Vendor = resourceData.Get("vendor.value").(string)
	attrs.ModelNumber = resourceData.Get("model_number").(string)
	attrs.AssignedTo = resourceData.Get("assigned_to.value").(string)
	attrs.StartDate = resourceData.Get("start_date").(string)
	attrs.OsVersion = resourceData.Get("os_version").(string)
	attrs.SerialNumber = resourceData.Get("serial_number").(string)
	attrs.CdRom = resourceData.Get("cd_rom").(string)
	attrs.SupportGroup = resourceData.Get("support_group.value").(string)
	attrs.Unverified = resourceData.Get("unverified").(string)
	attrs.CorrelationId = resourceData.Get("correlation_id").(string)
	attrs.Attributes = resourceData.Get("attributes").(string)
	attrs.Asset = resourceData.Get("asset.value").(string)
	attrs.FormFactor = resourceData.Get("form_factor").(string)
	attrs.CpuCoreCount = resourceData.Get("cpu_core_count").(string)
	attrs.SkipSync = resourceData.Get("skip_sync").(string)
	attrs.SysUpdatedBy = resourceData.Get("sys_updated_by").(string)
	attrs.SysCreatedOn = resourceData.Get("sys_created_on").(string)
	attrs.CpuType = resourceData.Get("cpu_type").(string)
	attrs.InstallDate = resourceData.Get("install_date").(string)
	attrs.AssetTag = resourceData.Get("asset_tag").(string)
	attrs.DrBackup = resourceData.Get("dr_backup.value").(string)
	attrs.HardwareSubstatus = resourceData.Get("hardware_substatus").(string)
	attrs.Fqdn = resourceData.Get("fqdn").(string)
	attrs.ChangeControl = resourceData.Get("change_control.value").(string)
	attrs.DeliveryDate = resourceData.Get("delivery_date").(string)
	attrs.HardwareStatus = resourceData.Get("hardware_status").(string)
	attrs.InstallStatus = resourceData.Get("install_status").(string)
	attrs.SupportedBy = resourceData.Get("supported_by.value").(string)
	attrs.Name = resourceData.Get("name").(string)
	attrs.Subcategory = resourceData.Get("subcategory").(string)
	attrs.DefaultGateway = resourceData.Get("default_gateway").(string)
	attrs.ChassisType = resourceData.Get("chassis_type").(string)
	attrs.Virtual = resourceData.Get("virtual").(string)
	attrs.AssignmentGroup = resourceData.Get("assignment_group.value").(string)
	attrs.PoNumber = resourceData.Get("po_number").(string)
	attrs.CheckedIn = resourceData.Get("checked_in").(string)
	attrs.MacAddress = resourceData.Get("mac_address").(string)
	attrs.Company = resourceData.Get("company.value").(string)
	attrs.Justification = resourceData.Get("justification").(string)
	attrs.Department = resourceData.Get("department.value").(string)
	attrs.Cost = resourceData.Get("cost").(string)
	attrs.Comments = resourceData.Get("comments").(string)
	attrs.Os = resourceData.Get("os").(string)
	attrs.SysModCount = resourceData.Get("sys_mod_count").(string)
	attrs.Monitor = resourceData.Get("monitor").(string)
	attrs.ModelId = resourceData.Get("model_id.value").(string)
	attrs.IpAddress = resourceData.Get("ip_address").(string)
	attrs.DuplicateOf = resourceData.Get("duplicate_of.value").(string)
	attrs.SysTags = resourceData.Get("sys_tags").(string)
	attrs.CostCc = resourceData.Get("cost_cc").(string)
	attrs.OrderDate = resourceData.Get("order_date").(string)
	attrs.Schedule = resourceData.Get("schedule.value").(string)
	attrs.Due = resourceData.Get("due").(string)
	attrs.Location = resourceData.Get("location.value").(string)
	attrs.Category = resourceData.Get("category").(string)
	attrs.FaultCount = resourceData.Get("fault_count").(string)
	attrs.LeaseId = resourceData.Get("lease_id").(string)
	attrs.HostName = resourceData.Get("host_name").(string)
	return nil
}

// This function Unmarshal's the JSON payload from ServiceNow into a corresponding STRUCT for the CI.  It then
// Sets each field in the Terraform schema.ResourceData object with the corrsponding field from the CI STRUCT.
//
// NOTE:  Terraform expects a MAP for fields that represent "reference fields" so that the "value", "link" and
//        "display_value" are decoded correctly by Terraform.  The map is constructed for each reference field
//        using a common function called "StructToMap" in the "client_base.go" file.
//
func copyFromServiceNowToTerraformCmdbCiSolarisServer(resourceData *schema.ResourceData, jsonData []byte) error {
	ci := CmdbCiSolarisServerGet{}
	if err := json.Unmarshal(jsonData, &ci); err != nil {
		//resourceData.SetId("")
		//return err
	}

	resourceData.SetId(GetSysId(jsonData))
	attrs := ci.Result.Attributes
	_ = resourceData.Set("os_address_width", attrs.OsAddressWidth)
	_ = resourceData.Set("firewall_status", attrs.FirewallStatus)
	_ = resourceData.Set("operational_status", attrs.OperationalStatus)
	_ = resourceData.Set("os_service_pack", attrs.OsServicePack)
	_ = resourceData.Set("cpu_core_thread", attrs.CpuCoreThread)
	_ = resourceData.Set("cpu_manufacturer", StructToMap(attrs.CpuManufacturer))
	_ = resourceData.Set("sys_updated_on", attrs.SysUpdatedOn)
	_ = resourceData.Set("discovery_source", attrs.DiscoverySource)
	_ = resourceData.Set("first_discovered", attrs.FirstDiscovered)
	_ = resourceData.Set("due_in", attrs.DueIn)
	_ = resourceData.Set("used_for", attrs.UsedFor)
	_ = resourceData.Set("invoice_number", attrs.InvoiceNumber)
	_ = resourceData.Set("gl_account", attrs.GlAccount)
	_ = resourceData.Set("sys_created_by", attrs.SysCreatedBy)
	_ = resourceData.Set("warranty_expiration", attrs.WarrantyExpiration)
	_ = resourceData.Set("ram", attrs.Ram)
	_ = resourceData.Set("cpu_name", attrs.CpuName)
	_ = resourceData.Set("cpu_speed", attrs.CpuSpeed)
	_ = resourceData.Set("owned_by", StructToMap(attrs.OwnedBy))
	_ = resourceData.Set("checked_out", attrs.CheckedOut)
	_ = resourceData.Set("sys_domain_path", attrs.SysDomainPath)
	_ = resourceData.Set("disk_space", attrs.DiskSpace)
	_ = resourceData.Set("classification", attrs.Classification)
	_ = resourceData.Set("object_id", attrs.ObjectId)
	_ = resourceData.Set("maintenance_schedule", StructToMap(attrs.MaintenanceSchedule))
	_ = resourceData.Set("cost_center", StructToMap(attrs.CostCenter))
	_ = resourceData.Set("dns_domain", attrs.DnsDomain)
	_ = resourceData.Set("assigned", attrs.Assigned)
	_ = resourceData.Set("purchase_date", attrs.PurchaseDate)
	_ = resourceData.Set("short_description", attrs.ShortDescription)
	_ = resourceData.Set("cd_speed", attrs.CdSpeed)
	_ = resourceData.Set("floppy", attrs.Floppy)
	_ = resourceData.Set("managed_by", StructToMap(attrs.ManagedBy))
	_ = resourceData.Set("os_domain", attrs.OsDomain)
	_ = resourceData.Set("last_discovered", attrs.LastDiscovered)
	_ = resourceData.Set("can_print", attrs.CanPrint)
	_ = resourceData.Set("manufacturer", StructToMap(attrs.Manufacturer))
	_ = resourceData.Set("cpu_count", attrs.CpuCount)
	_ = resourceData.Set("vendor", StructToMap(attrs.Vendor))
	_ = resourceData.Set("model_number", attrs.ModelNumber)
	_ = resourceData.Set("assigned_to", StructToMap(attrs.AssignedTo))
	_ = resourceData.Set("start_date", attrs.StartDate)
	_ = resourceData.Set("os_version", attrs.OsVersion)
	_ = resourceData.Set("serial_number", attrs.SerialNumber)
	_ = resourceData.Set("cd_rom", attrs.CdRom)
	_ = resourceData.Set("support_group", StructToMap(attrs.SupportGroup))
	_ = resourceData.Set("unverified", attrs.Unverified)
	_ = resourceData.Set("correlation_id", attrs.CorrelationId)
	_ = resourceData.Set("attributes", attrs.Attributes)
	_ = resourceData.Set("asset", StructToMap(attrs.Asset))
	_ = resourceData.Set("form_factor", attrs.FormFactor)
	_ = resourceData.Set("cpu_core_count", attrs.CpuCoreCount)
	_ = resourceData.Set("skip_sync", attrs.SkipSync)
	_ = resourceData.Set("sys_updated_by", attrs.SysUpdatedBy)
	_ = resourceData.Set("sys_created_on", attrs.SysCreatedOn)
	_ = resourceData.Set("cpu_type", attrs.CpuType)
	_ = resourceData.Set("install_date", attrs.InstallDate)
	_ = resourceData.Set("asset_tag", attrs.AssetTag)
	_ = resourceData.Set("dr_backup", StructToMap(attrs.DrBackup))
	_ = resourceData.Set("hardware_substatus", attrs.HardwareSubstatus)
	_ = resourceData.Set("fqdn", attrs.Fqdn)
	_ = resourceData.Set("change_control", StructToMap(attrs.ChangeControl))
	_ = resourceData.Set("delivery_date", attrs.DeliveryDate)
	_ = resourceData.Set("hardware_status", attrs.HardwareStatus)
	_ = resourceData.Set("install_status", attrs.InstallStatus)
	_ = resourceData.Set("supported_by", StructToMap(attrs.SupportedBy))
	_ = resourceData.Set("name", attrs.Name)
	_ = resourceData.Set("subcategory", attrs.Subcategory)
	_ = resourceData.Set("default_gateway", attrs.DefaultGateway)
	_ = resourceData.Set("chassis_type", attrs.ChassisType)
	_ = resourceData.Set("virtual", attrs.Virtual)
	_ = resourceData.Set("assignment_group", StructToMap(attrs.AssignmentGroup))
	_ = resourceData.Set("po_number", attrs.PoNumber)
	_ = resourceData.Set("checked_in", attrs.CheckedIn)
	_ = resourceData.Set("mac_address", attrs.MacAddress)
	_ = resourceData.Set("company", StructToMap(attrs.Company))
	_ = resourceData.Set("justification", attrs.Justification)
	_ = resourceData.Set("department", StructToMap(attrs.Department))
	_ = resourceData.Set("cost", attrs.Cost)
	_ = resourceData.Set("comments", attrs.Comments)
	_ = resourceData.Set("os", attrs.Os)
	_ = resourceData.Set("sys_mod_count", attrs.SysModCount)
	_ = resourceData.Set("monitor", attrs.Monitor)
	_ = resourceData.Set("model_id", StructToMap(attrs.ModelId))
	_ = resourceData.Set("ip_address", attrs.IpAddress)
	_ = resourceData.Set("duplicate_of", StructToMap(attrs.DuplicateOf))
	_ = resourceData.Set("sys_tags", attrs.SysTags)
	_ = resourceData.Set("cost_cc", attrs.CostCc)
	_ = resourceData.Set("order_date", attrs.OrderDate)
	_ = resourceData.Set("schedule", StructToMap(attrs.Schedule))
	_ = resourceData.Set("due", attrs.Due)
	_ = resourceData.Set("location", StructToMap(attrs.Location))
	_ = resourceData.Set("category", attrs.Category)
	_ = resourceData.Set("fault_count", attrs.FaultCount)
	_ = resourceData.Set("lease_id", attrs.LeaseId)
	_ = resourceData.Set("host_name", attrs.HostName)
	return nil
}
