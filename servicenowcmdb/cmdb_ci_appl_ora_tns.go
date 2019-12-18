package resources
///////////////////////////////////////////////////////////////////////////////////////////////////
//  ****WARNING***    : This is automatically generated code  using the generatprovidersource
//                      utility in this package, any edits to this file will be overwritten
//                      and lost.  Refer to the readme.md in the generateprovidersource directory
//                       for details on how to regenerate this code.
// Provider Version    : 0.01
// Generator Version   : 1.00
//
//  Description       :   This file is the resource provider for the cmdb_ci_appl_ora_tns CMDB Class.  This code is executed
//                        when the servicenowcmdb_cmdb_ci_appl_ora_tns keyword is used in a terraform script (*.tf) file
//
//                        This file is will need to be regenerated if the ServiceNow CMDB base CI Class
//                        "cmdb_ci" or of the cmdb_ci_appl_ora_tns CI Class is modified.
///////////////////////////////////////////////////////////////////////////////////////////////////
import (

"encoding/json"
"fmt"
"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)
// This constant stores the name of the CI and is used to construct the URL to ServiceNow
const CiNamecmdb_ci_appl_ora_tns  = "cmdb_ci_appl_ora_tns"

// This is the structure to construct the JSON payload when POSTing to ServiceNow.  This is needed because
// ServiceNow has strict parsing on the JSON Data and will fail if the JSON format doesn't match exactly.
// This is essentially the same as CmdbCiApplOraTnsGet but does not contain the "value", "link" and
// "display_value" fields for reference objects.
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiApplOraTnsPost struct {
Result struct {
Attributes   struct {
            TcpPort    string `json:"tcp_port,omitempty"`
            SkipSync    string `json:"skip_sync,omitempty"`
            OperationalStatus    string `json:"operational_status,omitempty"`
            RunningProcessCommand    string `json:"running_process_command,omitempty"`
            Pid    string `json:"pid,omitempty"`
            SysUpdatedOn    string `json:"sys_updated_on,omitempty"`
            RunningProcessKeyParameters    string `json:"running_process_key_parameters,omitempty"`
            RpCommandHash    string `json:"rp_command_hash,omitempty"`
            DiscoverySource    string `json:"discovery_source,omitempty"`
            FirstDiscovered    string `json:"first_discovered,omitempty"`
            SysUpdatedBy    string `json:"sys_updated_by,omitempty"`
            DueIn    string `json:"due_in,omitempty"`
            InstallDirectory    string `json:"install_directory,omitempty"`
            SysCreatedOn    string `json:"sys_created_on,omitempty"`
            Classifier    string `json:"classifier,omitempty"`
            UsedFor    string `json:"used_for,omitempty"`
            InstallDate    string `json:"install_date,omitempty"`
            IsClustered    string `json:"is_clustered,omitempty"`
            InvoiceNumber    string `json:"invoice_number,omitempty"`
            GlAccount    string `json:"gl_account,omitempty"`
            SysCreatedBy    string `json:"sys_created_by,omitempty"`
            WarrantyExpiration    string `json:"warranty_expiration,omitempty"`
            AssetTag    string `json:"asset_tag,omitempty"`
            Fqdn    string `json:"fqdn,omitempty"`
            ChangeControl    string `json:"change_control,omitempty"`
            OwnedBy    string `json:"owned_by,omitempty"`
            RpKeyParametersHash    string `json:"rp_key_parameters_hash,omitempty"`
            CheckedOut    string `json:"checked_out,omitempty"`
            SysDomainPath    string `json:"sys_domain_path,omitempty"`
            Version    string `json:"version,omitempty"`
            DeliveryDate    string `json:"delivery_date,omitempty"`
            MaintenanceSchedule    string `json:"maintenance_schedule,omitempty"`
            InstallStatus    string `json:"install_status,omitempty"`
            CostCenter    string `json:"cost_center,omitempty"`
            SupportedBy    string `json:"supported_by,omitempty"`
            DnsDomain    string `json:"dns_domain,omitempty"`
            Name    string `json:"name,omitempty"`
            Assigned    string `json:"assigned,omitempty"`
            PurchaseDate    string `json:"purchase_date,omitempty"`
            Subcategory    string `json:"subcategory,omitempty"`
            ShortDescription    string `json:"short_description,omitempty"`
            AssignmentGroup    string `json:"assignment_group,omitempty"`
            ManagedBy    string `json:"managed_by,omitempty"`
            Edition    string `json:"edition,omitempty"`
            LastDiscovered    string `json:"last_discovered,omitempty"`
            CanPrint    string `json:"can_print,omitempty"`
            Manufacturer    string `json:"manufacturer,omitempty"`
            PoNumber    string `json:"po_number,omitempty"`
            ClPort    string `json:"cl_port,omitempty"`
            CheckedIn    string `json:"checked_in,omitempty"`
            Vendor    string `json:"vendor,omitempty"`
            MacAddress    string `json:"mac_address,omitempty"`
            Company    string `json:"company,omitempty"`
            RunningProcess    string `json:"running_process,omitempty"`
            ModelNumber    string `json:"model_number,omitempty"`
            Justification    string `json:"justification,omitempty"`
            Department    string `json:"department,omitempty"`
            ConfigFile    string `json:"config_file,omitempty"`
            AssignedTo    string `json:"assigned_to,omitempty"`
            StartDate    string `json:"start_date,omitempty"`
            Cost    string `json:"cost,omitempty"`
            Comments    string `json:"comments,omitempty"`
            SysModCount    string `json:"sys_mod_count,omitempty"`
            SerialNumber    string `json:"serial_number,omitempty"`
            Monitor    string `json:"monitor,omitempty"`
            ModelId    string `json:"model_id,omitempty"`
            IpAddress    string `json:"ip_address,omitempty"`
            DuplicateOf    string `json:"duplicate_of,omitempty"`
            SysTags    string `json:"sys_tags,omitempty"`
            CostCc    string `json:"cost_cc,omitempty"`
            SupportGroup    string `json:"support_group,omitempty"`
            OrderDate    string `json:"order_date,omitempty"`
            Schedule    string `json:"schedule,omitempty"`
            Due    string `json:"due,omitempty"`
            Unverified    string `json:"unverified,omitempty"`
            CorrelationId    string `json:"correlation_id,omitempty"`
            Attributes    string `json:"attributes,omitempty"`
            Location    string `json:"location,omitempty"`
            Asset    string `json:"asset,omitempty"`
            Category    string `json:"category,omitempty"`
            FaultCount    string `json:"fault_count,omitempty"`
            ConfigDirectory    string `json:"config_directory,omitempty"`
            LeaseId    string `json:"lease_id,omitempty"`

} `json:"attributes"`
} `json:"result"`
}
// This is the structure where JSON data is Unmarshal'ed from ServiceNow.  This is different from the POST
// structure because it has to contain the "value", "link" and "display_value" fields for reference objects
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiApplOraTnsGet struct {
Result struct {
Attributes   struct {
                TcpPort    string `json:"tcp_port,omitempty"`
                SkipSync    string `json:"skip_sync,omitempty"`
                OperationalStatus    string `json:"operational_status,omitempty"`
                RunningProcessCommand    string `json:"running_process_command,omitempty"`
                Pid    string `json:"pid,omitempty"`
                SysUpdatedOn    string `json:"sys_updated_on,omitempty"`
                RunningProcessKeyParameters    string `json:"running_process_key_parameters,omitempty"`
                RpCommandHash    string `json:"rp_command_hash,omitempty"`
                DiscoverySource    string `json:"discovery_source,omitempty"`
                FirstDiscovered    string `json:"first_discovered,omitempty"`
                SysUpdatedBy    string `json:"sys_updated_by,omitempty"`
                DueIn    string `json:"due_in,omitempty"`
                InstallDirectory    string `json:"install_directory,omitempty"`
                SysCreatedOn    string `json:"sys_created_on,omitempty"`
                Classifier    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"classifier"`
                UsedFor    string `json:"used_for,omitempty"`
                InstallDate    string `json:"install_date,omitempty"`
                IsClustered    string `json:"is_clustered,omitempty"`
                InvoiceNumber    string `json:"invoice_number,omitempty"`
                GlAccount    string `json:"gl_account,omitempty"`
                SysCreatedBy    string `json:"sys_created_by,omitempty"`
                WarrantyExpiration    string `json:"warranty_expiration,omitempty"`
                AssetTag    string `json:"asset_tag,omitempty"`
                Fqdn    string `json:"fqdn,omitempty"`
                ChangeControl    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"change_control"`
                OwnedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"owned_by"`
                RpKeyParametersHash    string `json:"rp_key_parameters_hash,omitempty"`
                CheckedOut    string `json:"checked_out,omitempty"`
                SysDomainPath    string `json:"sys_domain_path,omitempty"`
                Version    string `json:"version,omitempty"`
                DeliveryDate    string `json:"delivery_date,omitempty"`
                MaintenanceSchedule    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"maintenance_schedule"`
                InstallStatus    string `json:"install_status,omitempty"`
                CostCenter    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"cost_center"`
                SupportedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"supported_by"`
                DnsDomain    string `json:"dns_domain,omitempty"`
                Name    string `json:"name,omitempty"`
                Assigned    string `json:"assigned,omitempty"`
                PurchaseDate    string `json:"purchase_date,omitempty"`
                Subcategory    string `json:"subcategory,omitempty"`
                ShortDescription    string `json:"short_description,omitempty"`
                AssignmentGroup    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"assignment_group"`
                ManagedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"managed_by"`
                Edition    string `json:"edition,omitempty"`
                LastDiscovered    string `json:"last_discovered,omitempty"`
                CanPrint    string `json:"can_print,omitempty"`
                Manufacturer    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"manufacturer"`
                PoNumber    string `json:"po_number,omitempty"`
                ClPort    string `json:"cl_port,omitempty"`
                CheckedIn    string `json:"checked_in,omitempty"`
                Vendor    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"vendor"`
                MacAddress    string `json:"mac_address,omitempty"`
                Company    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"company"`
                RunningProcess    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"running_process"`
                ModelNumber    string `json:"model_number,omitempty"`
                Justification    string `json:"justification,omitempty"`
                Department    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"department"`
                ConfigFile    string `json:"config_file,omitempty"`
                AssignedTo    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"assigned_to"`
                StartDate    string `json:"start_date,omitempty"`
                Cost    string `json:"cost,omitempty"`
                Comments    string `json:"comments,omitempty"`
                SysModCount    string `json:"sys_mod_count,omitempty"`
                SerialNumber    string `json:"serial_number,omitempty"`
                Monitor    string `json:"monitor,omitempty"`
                ModelId    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"model_id"`
                IpAddress    string `json:"ip_address,omitempty"`
                DuplicateOf    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"duplicate_of"`
                SysTags    string `json:"sys_tags,omitempty"`
                CostCc    string `json:"cost_cc,omitempty"`
                SupportGroup    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"support_group"`
                OrderDate    string `json:"order_date,omitempty"`
                Schedule    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"schedule"`
                Due    string `json:"due,omitempty"`
                Unverified    string `json:"unverified,omitempty"`
                CorrelationId    string `json:"correlation_id,omitempty"`
                Attributes    string `json:"attributes,omitempty"`
                Location    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"location"`
                Asset    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"asset"`
                Category    string `json:"category,omitempty"`
                FaultCount    string `json:"fault_count,omitempty"`
                ConfigDirectory    string `json:"config_directory,omitempty"`
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

func ResourceCmdbCiApplOraTns() *schema.Resource {
return &schema.Resource{
Create: createResourceCmdbCiApplOraTns,
Read:   readResourceCmdbCiApplOraTns,
Update: updateResourceCmdbCiApplOraTns,
Delete: deleteResourceCmdbCiApplOraTns,

Importer: &schema.ResourceImporter{
State:schema.ImportStatePassthrough,
},

Schema: map[string]*schema.Schema{
                "tcp_port" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "skip_sync" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "operational_status" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "running_process_command" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "pid" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_updated_on" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "running_process_key_parameters" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "rp_command_hash" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "discovery_source" : {
                Type:     schema.TypeString,
                    Optional:true,Default:"Terraform",
                },
                "first_discovered" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "sys_updated_by" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "due_in" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "install_directory" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_created_on" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "classifier" : {
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
                "used_for" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "install_date" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "is_clustered" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "invoice_number" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "gl_account" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_created_by" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "warranty_expiration" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "asset_tag" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "fqdn" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "change_control" : {
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
                "owned_by" : {
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
                "rp_key_parameters_hash" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "checked_out" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_domain_path" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "version" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "delivery_date" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "maintenance_schedule" : {
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
                "install_status" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "cost_center" : {
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
                "supported_by" : {
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
                "dns_domain" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "name" : {
                Type:     schema.TypeString,
                    Required:true,
                },
                "assigned" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "purchase_date" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "subcategory" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "short_description" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "assignment_group" : {
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
                "managed_by" : {
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
                "edition" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "last_discovered" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "can_print" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "manufacturer" : {
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
                "po_number" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "cl_port" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "checked_in" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "vendor" : {
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
                "mac_address" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "company" : {
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
                "running_process" : {
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
                "model_number" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "justification" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "department" : {
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
                "config_file" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "assigned_to" : {
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
                "start_date" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "cost" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "comments" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "sys_mod_count" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "serial_number" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "monitor" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "model_id" : {
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
                "ip_address" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "duplicate_of" : {
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
                "sys_tags" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "cost_cc" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "support_group" : {
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
                "order_date" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "schedule" : {
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
                "due" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "unverified" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "correlation_id" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "attributes" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "location" : {
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
                "asset" : {
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
                "category" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "fault_count" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "config_directory" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "lease_id" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
},
}
}

//  Create routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func createResourceCmdbCiApplOraTns(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client) //Client Connection details
// Use common function to update base attributes
var ci CmdbCiApplOraTnsPost
if err := copyFromTerraformToServiceNowCmdbCiApplOraTns(resourceData, &ci); err != nil {
return err
}

//t := time.Now()
//ci.Result.Attributes.Installdate = t.String()
//ci.Result.Attributes.SysCreatedOn = t.String()

// Using "sysparm fields=sys_id" reduces the amount of data received, we only need the sys_id for create action
SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_appl_ora_tns + "?sysparm_fields=sys_id"

jsonData, err := servicenowClient.RequestJSON("POST", SnowUrl, ci.Result)
if err != nil {
resourceData.SetId("")
return err
}

resourceData.SetId(GetSysId(jsonData))
return readResourceCmdbCiApplOraTns(resourceData, serviceNowClient)
}

//  Read routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func readResourceCmdbCiApplOraTns(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client)
SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_appl_ora_tns + "/" + resourceData.Id()
var jsonData []byte
jsonData, err := servicenowClient.RequestJSON("GET", SnowUrl, nil)
if err != nil {
resourceData.SetId("")
return err
}

if err := copyFromServiceNowToTerraformCmdbCiApplOraTns(resourceData,jsonData);  err != nil {
return err
}

return nil
}

//  Update routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func updateResourceCmdbCiApplOraTns(resourceData *schema.ResourceData, serviceNowClient interface{}) error {
servicenowClient := serviceNowClient.(*Client)

var ci CmdbCiApplOraTnsPost
if err := copyFromTerraformToServiceNowCmdbCiApplOraTns(resourceData, &ci); err != nil {
return err
}

SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_appl_ora_tns + "/" + resourceData.Id()
//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
if err != nil {
resourceData.SetId("")
return err
}
return readResourceCmdbCiApplOraTns(resourceData, serviceNowClient)
}

// TODO:  Need to work out what to do with deleting CIs. ServiceNow does not support deleting CIs via the API
//        CIs shouldn't be deleted because may be referenced by other objects in the ServiceNow database
//        (eg changes and incidents).  Therefore the best practice is to change their operational and
//        install status.  The tricky bit is working out how to make this flexible for everyone to use.

//  Delete routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func deleteResourceCmdbCiApplOraTns(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client)
var ci CmdbCiApplOraTnsPost
if err := copyFromTerraformToServiceNowCmdbCiApplOraTns(resourceData, &ci); err != nil {
return err
}

if err := resourceData.Set("install_status", "retired"); err != nil {
return fmt.Errorf("CmdbCiApplOraTnsfailed to set install_status field during destroy action %s", err)
}

SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_appl_ora_tns + "/" + resourceData.Id()
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
func copyFromTerraformToServiceNowCmdbCiApplOraTns (resourceData *schema.ResourceData, ci *CmdbCiApplOraTnsPost) error {

attrs := &ci.Result.Attributes
                attrs.TcpPort  = resourceData.Get("tcp_port").(string)
                attrs.SkipSync  = resourceData.Get("skip_sync").(string)
                attrs.OperationalStatus  = resourceData.Get("operational_status").(string)
                attrs.RunningProcessCommand  = resourceData.Get("running_process_command").(string)
                attrs.Pid  = resourceData.Get("pid").(string)
                attrs.SysUpdatedOn  = resourceData.Get("sys_updated_on").(string)
                attrs.RunningProcessKeyParameters  = resourceData.Get("running_process_key_parameters").(string)
                attrs.RpCommandHash  = resourceData.Get("rp_command_hash").(string)
                attrs.DiscoverySource  = resourceData.Get("discovery_source").(string)
                attrs.FirstDiscovered  = resourceData.Get("first_discovered").(string)
                attrs.SysUpdatedBy  = resourceData.Get("sys_updated_by").(string)
                attrs.DueIn  = resourceData.Get("due_in").(string)
                attrs.InstallDirectory  = resourceData.Get("install_directory").(string)
                attrs.SysCreatedOn  = resourceData.Get("sys_created_on").(string)
                attrs.Classifier  = resourceData.Get("classifier.value").(string)
                attrs.UsedFor  = resourceData.Get("used_for").(string)
                attrs.InstallDate  = resourceData.Get("install_date").(string)
                attrs.IsClustered  = resourceData.Get("is_clustered").(string)
                attrs.InvoiceNumber  = resourceData.Get("invoice_number").(string)
                attrs.GlAccount  = resourceData.Get("gl_account").(string)
                attrs.SysCreatedBy  = resourceData.Get("sys_created_by").(string)
                attrs.WarrantyExpiration  = resourceData.Get("warranty_expiration").(string)
                attrs.AssetTag  = resourceData.Get("asset_tag").(string)
                attrs.Fqdn  = resourceData.Get("fqdn").(string)
                attrs.ChangeControl  = resourceData.Get("change_control.value").(string)
                attrs.OwnedBy  = resourceData.Get("owned_by.value").(string)
                attrs.RpKeyParametersHash  = resourceData.Get("rp_key_parameters_hash").(string)
                attrs.CheckedOut  = resourceData.Get("checked_out").(string)
                attrs.SysDomainPath  = resourceData.Get("sys_domain_path").(string)
                attrs.Version  = resourceData.Get("version").(string)
                attrs.DeliveryDate  = resourceData.Get("delivery_date").(string)
                attrs.MaintenanceSchedule  = resourceData.Get("maintenance_schedule.value").(string)
                attrs.InstallStatus  = resourceData.Get("install_status").(string)
                attrs.CostCenter  = resourceData.Get("cost_center.value").(string)
                attrs.SupportedBy  = resourceData.Get("supported_by.value").(string)
                attrs.DnsDomain  = resourceData.Get("dns_domain").(string)
                attrs.Name  = resourceData.Get("name").(string)
                attrs.Assigned  = resourceData.Get("assigned").(string)
                attrs.PurchaseDate  = resourceData.Get("purchase_date").(string)
                attrs.Subcategory  = resourceData.Get("subcategory").(string)
                attrs.ShortDescription  = resourceData.Get("short_description").(string)
                attrs.AssignmentGroup  = resourceData.Get("assignment_group.value").(string)
                attrs.ManagedBy  = resourceData.Get("managed_by.value").(string)
                attrs.Edition  = resourceData.Get("edition").(string)
                attrs.LastDiscovered  = resourceData.Get("last_discovered").(string)
                attrs.CanPrint  = resourceData.Get("can_print").(string)
                attrs.Manufacturer  = resourceData.Get("manufacturer.value").(string)
                attrs.PoNumber  = resourceData.Get("po_number").(string)
                attrs.ClPort  = resourceData.Get("cl_port").(string)
                attrs.CheckedIn  = resourceData.Get("checked_in").(string)
                attrs.Vendor  = resourceData.Get("vendor.value").(string)
                attrs.MacAddress  = resourceData.Get("mac_address").(string)
                attrs.Company  = resourceData.Get("company.value").(string)
                attrs.RunningProcess  = resourceData.Get("running_process.value").(string)
                attrs.ModelNumber  = resourceData.Get("model_number").(string)
                attrs.Justification  = resourceData.Get("justification").(string)
                attrs.Department  = resourceData.Get("department.value").(string)
                attrs.ConfigFile  = resourceData.Get("config_file").(string)
                attrs.AssignedTo  = resourceData.Get("assigned_to.value").(string)
                attrs.StartDate  = resourceData.Get("start_date").(string)
                attrs.Cost  = resourceData.Get("cost").(string)
                attrs.Comments  = resourceData.Get("comments").(string)
                attrs.SysModCount  = resourceData.Get("sys_mod_count").(string)
                attrs.SerialNumber  = resourceData.Get("serial_number").(string)
                attrs.Monitor  = resourceData.Get("monitor").(string)
                attrs.ModelId  = resourceData.Get("model_id.value").(string)
                attrs.IpAddress  = resourceData.Get("ip_address").(string)
                attrs.DuplicateOf  = resourceData.Get("duplicate_of.value").(string)
                attrs.SysTags  = resourceData.Get("sys_tags").(string)
                attrs.CostCc  = resourceData.Get("cost_cc").(string)
                attrs.SupportGroup  = resourceData.Get("support_group.value").(string)
                attrs.OrderDate  = resourceData.Get("order_date").(string)
                attrs.Schedule  = resourceData.Get("schedule.value").(string)
                attrs.Due  = resourceData.Get("due").(string)
                attrs.Unverified  = resourceData.Get("unverified").(string)
                attrs.CorrelationId  = resourceData.Get("correlation_id").(string)
                attrs.Attributes  = resourceData.Get("attributes").(string)
                attrs.Location  = resourceData.Get("location.value").(string)
                attrs.Asset  = resourceData.Get("asset.value").(string)
                attrs.Category  = resourceData.Get("category").(string)
                attrs.FaultCount  = resourceData.Get("fault_count").(string)
                attrs.ConfigDirectory  = resourceData.Get("config_directory").(string)
                attrs.LeaseId  = resourceData.Get("lease_id").(string)
return nil
}

// This function Unmarshal's the JSON payload from ServiceNow into a corresponding STRUCT for the CI.  It then
// Sets each field in the Terraform schema.ResourceData object with the corrsponding field from the CI STRUCT.
//
// NOTE:  Terraform expects a MAP for fields that represent "reference fields" so that the "value", "link" and
//        "display_value" are decoded correctly by Terraform.  The map is constructed for each reference field
//        using a common function called "StructToMap" in the "client_base.go" file.
//
func copyFromServiceNowToTerraformCmdbCiApplOraTns (resourceData *schema.ResourceData, jsonData [] byte) error {
ci := CmdbCiApplOraTnsGet{}
if err := json.Unmarshal(jsonData, &ci); err != nil {
//resourceData.SetId("")
//return err
}

resourceData.SetId(GetSysId(jsonData))
attrs := ci.Result.Attributes
                _ = resourceData.Set("tcp_port", attrs.TcpPort)
                _ = resourceData.Set("skip_sync", attrs.SkipSync)
                _ = resourceData.Set("operational_status", attrs.OperationalStatus)
                _ = resourceData.Set("running_process_command", attrs.RunningProcessCommand)
                _ = resourceData.Set("pid", attrs.Pid)
                _ = resourceData.Set("sys_updated_on", attrs.SysUpdatedOn)
                _ = resourceData.Set("running_process_key_parameters", attrs.RunningProcessKeyParameters)
                _ = resourceData.Set("rp_command_hash", attrs.RpCommandHash)
                _ = resourceData.Set("discovery_source", attrs.DiscoverySource)
                _ = resourceData.Set("first_discovered", attrs.FirstDiscovered)
                _ = resourceData.Set("sys_updated_by", attrs.SysUpdatedBy)
                _ = resourceData.Set("due_in", attrs.DueIn)
                _ = resourceData.Set("install_directory", attrs.InstallDirectory)
                _ = resourceData.Set("sys_created_on", attrs.SysCreatedOn)
                _ = resourceData.Set("classifier", StructToMap(attrs.Classifier))
                _ = resourceData.Set("used_for", attrs.UsedFor)
                _ = resourceData.Set("install_date", attrs.InstallDate)
                _ = resourceData.Set("is_clustered", attrs.IsClustered)
                _ = resourceData.Set("invoice_number", attrs.InvoiceNumber)
                _ = resourceData.Set("gl_account", attrs.GlAccount)
                _ = resourceData.Set("sys_created_by", attrs.SysCreatedBy)
                _ = resourceData.Set("warranty_expiration", attrs.WarrantyExpiration)
                _ = resourceData.Set("asset_tag", attrs.AssetTag)
                _ = resourceData.Set("fqdn", attrs.Fqdn)
                _ = resourceData.Set("change_control", StructToMap(attrs.ChangeControl))
                _ = resourceData.Set("owned_by", StructToMap(attrs.OwnedBy))
                _ = resourceData.Set("rp_key_parameters_hash", attrs.RpKeyParametersHash)
                _ = resourceData.Set("checked_out", attrs.CheckedOut)
                _ = resourceData.Set("sys_domain_path", attrs.SysDomainPath)
                _ = resourceData.Set("version", attrs.Version)
                _ = resourceData.Set("delivery_date", attrs.DeliveryDate)
                _ = resourceData.Set("maintenance_schedule", StructToMap(attrs.MaintenanceSchedule))
                _ = resourceData.Set("install_status", attrs.InstallStatus)
                _ = resourceData.Set("cost_center", StructToMap(attrs.CostCenter))
                _ = resourceData.Set("supported_by", StructToMap(attrs.SupportedBy))
                _ = resourceData.Set("dns_domain", attrs.DnsDomain)
                _ = resourceData.Set("name", attrs.Name)
                _ = resourceData.Set("assigned", attrs.Assigned)
                _ = resourceData.Set("purchase_date", attrs.PurchaseDate)
                _ = resourceData.Set("subcategory", attrs.Subcategory)
                _ = resourceData.Set("short_description", attrs.ShortDescription)
                _ = resourceData.Set("assignment_group", StructToMap(attrs.AssignmentGroup))
                _ = resourceData.Set("managed_by", StructToMap(attrs.ManagedBy))
                _ = resourceData.Set("edition", attrs.Edition)
                _ = resourceData.Set("last_discovered", attrs.LastDiscovered)
                _ = resourceData.Set("can_print", attrs.CanPrint)
                _ = resourceData.Set("manufacturer", StructToMap(attrs.Manufacturer))
                _ = resourceData.Set("po_number", attrs.PoNumber)
                _ = resourceData.Set("cl_port", attrs.ClPort)
                _ = resourceData.Set("checked_in", attrs.CheckedIn)
                _ = resourceData.Set("vendor", StructToMap(attrs.Vendor))
                _ = resourceData.Set("mac_address", attrs.MacAddress)
                _ = resourceData.Set("company", StructToMap(attrs.Company))
                _ = resourceData.Set("running_process", StructToMap(attrs.RunningProcess))
                _ = resourceData.Set("model_number", attrs.ModelNumber)
                _ = resourceData.Set("justification", attrs.Justification)
                _ = resourceData.Set("department", StructToMap(attrs.Department))
                _ = resourceData.Set("config_file", attrs.ConfigFile)
                _ = resourceData.Set("assigned_to", StructToMap(attrs.AssignedTo))
                _ = resourceData.Set("start_date", attrs.StartDate)
                _ = resourceData.Set("cost", attrs.Cost)
                _ = resourceData.Set("comments", attrs.Comments)
                _ = resourceData.Set("sys_mod_count", attrs.SysModCount)
                _ = resourceData.Set("serial_number", attrs.SerialNumber)
                _ = resourceData.Set("monitor", attrs.Monitor)
                _ = resourceData.Set("model_id", StructToMap(attrs.ModelId))
                _ = resourceData.Set("ip_address", attrs.IpAddress)
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
                _ = resourceData.Set("config_directory", attrs.ConfigDirectory)
                _ = resourceData.Set("lease_id", attrs.LeaseId)
return nil
}
