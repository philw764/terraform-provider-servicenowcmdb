package resources
///////////////////////////////////////////////////////////////////////////////////////////////////
//  ****WARNING***    : This is automatically generated code  using the generatprovidersource
//                      utility in this package, any edits to this file will be overwritten
//                      and lost.  Refer to the readme.md in the generateprovidersource directory
//                       for details on how to regenerate this code.
// Provider Version    : 0.01
// Generator Version   : 1.00
//
//  Description       :   This file is the resource provider for the cmdb_ci_business_app CMDB Class.  This code is executed
//                        when the servicenowcmdb_cmdb_ci_business_app keyword is used in a terraform script (*.tf) file
//
//                        This file is will need to be regenerated if the ServiceNow CMDB base CI Class
//                        "cmdb_ci" or of the cmdb_ci_business_app CI Class is modified.
///////////////////////////////////////////////////////////////////////////////////////////////////
import (

"encoding/json"
"fmt"
"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)
// This constant stores the name of the CI and is used to construct the URL to ServiceNow
const CiNamecmdb_ci_business_app  = "cmdb_ci_business_app"

// This is the structure to construct the JSON payload when POSTing to ServiceNow.  This is needed because
// ServiceNow has strict parsing on the JSON Data and will fail if the JSON format doesn't match exactly.
// This is essentially the same as CmdbCiBusinessAppGet but does not contain the "value", "link" and
// "display_value" fields for reference objects.
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiBusinessAppPost struct {
Result struct {
Attributes   struct {
            OperationalStatus    string `json:"operational_status,omitempty"`
            SysUpdatedOn    string `json:"sys_updated_on,omitempty"`
            InstallType    string `json:"install_type,omitempty"`
            Number    string `json:"number,omitempty"`
            DiscoverySource    string `json:"discovery_source,omitempty"`
            FirstDiscovered    string `json:"first_discovered,omitempty"`
            DueIn    string `json:"due_in,omitempty"`
            ItApplicationOwner    string `json:"it_application_owner,omitempty"`
            ApmBusinessProcess    string `json:"apm_business_process,omitempty"`
            InvoiceNumber    string `json:"invoice_number,omitempty"`
            GlAccount    string `json:"gl_account,omitempty"`
            SysCreatedBy    string `json:"sys_created_by,omitempty"`
            WarrantyExpiration    string `json:"warranty_expiration,omitempty"`
            OrganizationUnitCount    string `json:"organization_unit_count,omitempty"`
            Active    string `json:"active,omitempty"`
            OwnedBy    string `json:"owned_by,omitempty"`
            CheckedOut    string `json:"checked_out,omitempty"`
            SysDomainPath    string `json:"sys_domain_path,omitempty"`
            BusinessUnit    string `json:"business_unit,omitempty"`
            MaintenanceSchedule    string `json:"maintenance_schedule,omitempty"`
            ApplicationManager    string `json:"application_manager,omitempty"`
            CostCenter    string `json:"cost_center,omitempty"`
            DnsDomain    string `json:"dns_domain,omitempty"`
            Assigned    string `json:"assigned,omitempty"`
            PurchaseDate    string `json:"purchase_date,omitempty"`
            ShortDescription    string `json:"short_description,omitempty"`
            LastChangeDate    string `json:"last_change_date,omitempty"`
            ManagedBy    string `json:"managed_by,omitempty"`
            LastDiscovered    string `json:"last_discovered,omitempty"`
            CanPrint    string `json:"can_print,omitempty"`
            NextAssessmentDate    string `json:"next_assessment_date,omitempty"`
            Manufacturer    string `json:"manufacturer,omitempty"`
            ContractEndDate    string `json:"contract_end_date,omitempty"`
            DataClassification    string `json:"data_classification,omitempty"`
            Vendor    string `json:"vendor,omitempty"`
            Currency    string `json:"currency,omitempty"`
            Certified    string `json:"certified,omitempty"`
            ModelNumber    string `json:"model_number,omitempty"`
            AssignedTo    string `json:"assigned_to,omitempty"`
            StartDate    string `json:"start_date,omitempty"`
            BusinessCriticality    string `json:"business_criticality,omitempty"`
            SerialNumber    string `json:"serial_number,omitempty"`
            Url    string `json:"url,omitempty"`
            SupportGroup    string `json:"support_group,omitempty"`
            TechnologyStack    string `json:"technology_stack,omitempty"`
            AudienceType    string `json:"audience_type,omitempty"`
            Unverified    string `json:"unverified,omitempty"`
            CorrelationId    string `json:"correlation_id,omitempty"`
            Attributes    string `json:"attributes,omitempty"`
            Asset    string `json:"asset,omitempty"`
            SoftwareLicense    string `json:"software_license,omitempty"`
            SkipSync    string `json:"skip_sync,omitempty"`
            ActiveUserCount    string `json:"active_user_count,omitempty"`
            SupportVendor    string `json:"support_vendor,omitempty"`
            AppraisalFiscalType    string `json:"appraisal_fiscal_type,omitempty"`
            SysUpdatedBy    string `json:"sys_updated_by,omitempty"`
            SysCreatedOn    string `json:"sys_created_on,omitempty"`
            InstallDate    string `json:"install_date,omitempty"`
            AssetTag    string `json:"asset_tag,omitempty"`
            Fqdn    string `json:"fqdn,omitempty"`
            ChangeControl    string `json:"change_control,omitempty"`
            EmergencyTier    string `json:"emergency_tier,omitempty"`
            ProductSupportStatus    string `json:"product_support_status,omitempty"`
            DeliveryDate    string `json:"delivery_date,omitempty"`
            InstallStatus    string `json:"install_status,omitempty"`
            SupportedBy    string `json:"supported_by,omitempty"`
            Name    string `json:"name,omitempty"`
            Subcategory    string `json:"subcategory,omitempty"`
            WorkNotes    string `json:"work_notes,omitempty"`
            AssignmentGroup    string `json:"assignment_group,omitempty"`
            ApplicationType    string `json:"application_type,omitempty"`
            ArchitectureType    string `json:"architecture_type,omitempty"`
            Platform    string `json:"platform,omitempty"`
            PoNumber    string `json:"po_number,omitempty"`
            CheckedIn    string `json:"checked_in,omitempty"`
            MacAddress    string `json:"mac_address,omitempty"`
            Company    string `json:"company,omitempty"`
            Justification    string `json:"justification,omitempty"`
            Department    string `json:"department,omitempty"`
            Cost    string `json:"cost,omitempty"`
            Comments    string `json:"comments,omitempty"`
            CmdbSoftwareProductModel    string `json:"cmdb_software_product_model,omitempty"`
            SysModCount    string `json:"sys_mod_count,omitempty"`
            Monitor    string `json:"monitor,omitempty"`
            ModelId    string `json:"model_id,omitempty"`
            IpAddress    string `json:"ip_address,omitempty"`
            DuplicateOf    string `json:"duplicate_of,omitempty"`
            SysTags    string `json:"sys_tags,omitempty"`
            CostCc    string `json:"cost_cc,omitempty"`
            UserBase    string `json:"user_base,omitempty"`
            OrderDate    string `json:"order_date,omitempty"`
            Schedule    string `json:"schedule,omitempty"`
            Due    string `json:"due,omitempty"`
            PlatformHost    string `json:"platform_host,omitempty"`
            Location    string `json:"location,omitempty"`
            Category    string `json:"category,omitempty"`
            FaultCount    string `json:"fault_count,omitempty"`
            Age    string `json:"age,omitempty"`
            LeaseId    string `json:"lease_id,omitempty"`

} `json:"attributes"`
} `json:"result"`
}
// This is the structure where JSON data is Unmarshal'ed from ServiceNow.  This is different from the POST
// structure because it has to contain the "value", "link" and "display_value" fields for reference objects
//
// The generateprovidersource utility constructs this STRUCT from metadata pulled from ServiceNow.
//
type CmdbCiBusinessAppGet struct {
Result struct {
Attributes   struct {
                OperationalStatus    string `json:"operational_status,omitempty"`
                SysUpdatedOn    string `json:"sys_updated_on,omitempty"`
                InstallType    string `json:"install_type,omitempty"`
                Number    string `json:"number,omitempty"`
                DiscoverySource    string `json:"discovery_source,omitempty"`
                FirstDiscovered    string `json:"first_discovered,omitempty"`
                DueIn    string `json:"due_in,omitempty"`
                ItApplicationOwner    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"it_application_owner"`
                ApmBusinessProcess    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"apm_business_process"`
                InvoiceNumber    string `json:"invoice_number,omitempty"`
                GlAccount    string `json:"gl_account,omitempty"`
                SysCreatedBy    string `json:"sys_created_by,omitempty"`
                WarrantyExpiration    string `json:"warranty_expiration,omitempty"`
                OrganizationUnitCount    string `json:"organization_unit_count,omitempty"`
                Active    string `json:"active,omitempty"`
                OwnedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"owned_by"`
                CheckedOut    string `json:"checked_out,omitempty"`
                SysDomainPath    string `json:"sys_domain_path,omitempty"`
                BusinessUnit    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"business_unit"`
                MaintenanceSchedule    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"maintenance_schedule"`
                ApplicationManager    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"application_manager"`
                CostCenter    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"cost_center"`
                DnsDomain    string `json:"dns_domain,omitempty"`
                Assigned    string `json:"assigned,omitempty"`
                PurchaseDate    string `json:"purchase_date,omitempty"`
                ShortDescription    string `json:"short_description,omitempty"`
                LastChangeDate    string `json:"last_change_date,omitempty"`
                ManagedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"managed_by"`
                LastDiscovered    string `json:"last_discovered,omitempty"`
                CanPrint    string `json:"can_print,omitempty"`
                NextAssessmentDate    string `json:"next_assessment_date,omitempty"`
                Manufacturer    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"manufacturer"`
                ContractEndDate    string `json:"contract_end_date,omitempty"`
                DataClassification    string `json:"data_classification,omitempty"`
                Vendor    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"vendor"`
                Currency    string `json:"currency,omitempty"`
                Certified    string `json:"certified,omitempty"`
                ModelNumber    string `json:"model_number,omitempty"`
                AssignedTo    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"assigned_to"`
                StartDate    string `json:"start_date,omitempty"`
                BusinessCriticality    string `json:"business_criticality,omitempty"`
                SerialNumber    string `json:"serial_number,omitempty"`
                Url    string `json:"url,omitempty"`
                SupportGroup    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"support_group"`
                TechnologyStack    string `json:"technology_stack,omitempty"`
                AudienceType    string `json:"audience_type,omitempty"`
                Unverified    string `json:"unverified,omitempty"`
                CorrelationId    string `json:"correlation_id,omitempty"`
                Attributes    string `json:"attributes,omitempty"`
                Asset    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"asset"`
                SoftwareLicense    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"software_license"`
                SkipSync    string `json:"skip_sync,omitempty"`
                ActiveUserCount    string `json:"active_user_count,omitempty"`
                SupportVendor    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"support_vendor"`
                AppraisalFiscalType    string `json:"appraisal_fiscal_type,omitempty"`
                SysUpdatedBy    string `json:"sys_updated_by,omitempty"`
                SysCreatedOn    string `json:"sys_created_on,omitempty"`
                InstallDate    string `json:"install_date,omitempty"`
                AssetTag    string `json:"asset_tag,omitempty"`
                Fqdn    string `json:"fqdn,omitempty"`
                ChangeControl    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"change_control"`
                EmergencyTier    string `json:"emergency_tier,omitempty"`
                ProductSupportStatus    string `json:"product_support_status,omitempty"`
                DeliveryDate    string `json:"delivery_date,omitempty"`
                InstallStatus    string `json:"install_status,omitempty"`
                SupportedBy    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"supported_by"`
                Name    string `json:"name,omitempty"`
                Subcategory    string `json:"subcategory,omitempty"`
                WorkNotes    string `json:"work_notes,omitempty"`
                AssignmentGroup    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"assignment_group"`
                ApplicationType    string `json:"application_type,omitempty"`
                ArchitectureType    string `json:"architecture_type,omitempty"`
                Platform    string `json:"platform,omitempty"`
                PoNumber    string `json:"po_number,omitempty"`
                CheckedIn    string `json:"checked_in,omitempty"`
                MacAddress    string `json:"mac_address,omitempty"`
                Company    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"company"`
                Justification    string `json:"justification,omitempty"`
                Department    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"department"`
                Cost    string `json:"cost,omitempty"`
                Comments    string `json:"comments,omitempty"`
                CmdbSoftwareProductModel    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"cmdb_software_product_model"`
                SysModCount    string `json:"sys_mod_count,omitempty"`
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
                UserBase    string `json:"user_base,omitempty"`
                OrderDate    string `json:"order_date,omitempty"`
                Schedule    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"schedule"`
                Due    string `json:"due,omitempty"`
                PlatformHost    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"platform_host"`
                Location    struct {
                Value   string `json:"value,omitempty"`
                Link    string `json:"link,omitempty"`
                DisplayValue string `json:"display_value,omitempty"`
                } `json:"location"`
                Category    string `json:"category,omitempty"`
                FaultCount    string `json:"fault_count,omitempty"`
                Age    string `json:"age,omitempty"`
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

func ResourceCmdbCiBusinessApp() *schema.Resource {
return &schema.Resource{
Create: createResourceCmdbCiBusinessApp,
Read:   readResourceCmdbCiBusinessApp,
Update: updateResourceCmdbCiBusinessApp,
Delete: deleteResourceCmdbCiBusinessApp,

Importer: &schema.ResourceImporter{
State:schema.ImportStatePassthrough,
},

Schema: map[string]*schema.Schema{
                "operational_status" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_updated_on" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "install_type" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "number" : {
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
                "due_in" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "it_application_owner" : {
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
                "apm_business_process" : {
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
                "organization_unit_count" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "active" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "checked_out" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_domain_path" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "business_unit" : {
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
                "application_manager" : {
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
                "dns_domain" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "assigned" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "purchase_date" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "short_description" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "last_change_date" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "last_discovered" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "can_print" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "next_assessment_date" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "contract_end_date" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "data_classification" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "currency" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "certified" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "model_number" : {
                Type:     schema.TypeString,
                    Computed:true,
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
                "business_criticality" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "serial_number" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "url" : {
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
                "technology_stack" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "audience_type" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "software_license" : {
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
                "skip_sync" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "active_user_count" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "support_vendor" : {
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
                "appraisal_fiscal_type" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "sys_updated_by" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "sys_created_on" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "install_date" : {
                Type:     schema.TypeString,
                    Computed:true,
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
                "emergency_tier" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "product_support_status" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "delivery_date" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "install_status" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "name" : {
                Type:     schema.TypeString,
                    Required:true,
                },
                "subcategory" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "work_notes" : {
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
                "application_type" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "architecture_type" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "platform" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "po_number" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "checked_in" : {
                Type:     schema.TypeString,
                    Computed:true,
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
                "cost" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "comments" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "cmdb_software_product_model" : {
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
                "sys_mod_count" : {
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
                "user_base" : {
                Type:     schema.TypeString,
                    Optional:true,
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
                "platform_host" : {
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
                "category" : {
                Type:     schema.TypeString,
                    Computed:true,
                },
                "fault_count" : {
                Type:     schema.TypeString,
                    Optional:true,
                },
                "age" : {
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
func createResourceCmdbCiBusinessApp(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client) //Client Connection details
// Use common function to update base attributes
var ci CmdbCiBusinessAppPost
if err := copyFromTerraformToServiceNowCmdbCiBusinessApp(resourceData, &ci); err != nil {
return err
}

//t := time.Now()
//ci.Result.Attributes.Installdate = t.String()
//ci.Result.Attributes.SysCreatedOn = t.String()

// Using "sysparm fields=sys_id" reduces the amount of data received, we only need the sys_id for create action
SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_business_app + "?sysparm_fields=sys_id"

jsonData, err := servicenowClient.RequestJSON("POST", SnowUrl, ci.Result)
if err != nil {
resourceData.SetId("")
return err
}

resourceData.SetId(GetSysId(jsonData))
return readResourceCmdbCiBusinessApp(resourceData, serviceNowClient)
}

//  Read routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func readResourceCmdbCiBusinessApp(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client)
SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_business_app + "/" + resourceData.Id()
var jsonData []byte
jsonData, err := servicenowClient.RequestJSON("GET", SnowUrl, nil)
if err != nil {
resourceData.SetId("")
return err
}

if err := copyFromServiceNowToTerraformCmdbCiBusinessApp(resourceData,jsonData);  err != nil {
return err
}

return nil
}

//  Update routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func updateResourceCmdbCiBusinessApp(resourceData *schema.ResourceData, serviceNowClient interface{}) error {
servicenowClient := serviceNowClient.(*Client)

var ci CmdbCiBusinessAppPost
if err := copyFromTerraformToServiceNowCmdbCiBusinessApp(resourceData, &ci); err != nil {
return err
}

SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_business_app + "/" + resourceData.Id()
//jsonData, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
_, err := servicenowClient.RequestJSON("PATCH", SnowUrl, ci.Result)
if err != nil {
resourceData.SetId("")
return err
}
return readResourceCmdbCiBusinessApp(resourceData, serviceNowClient)
}

// TODO:  Need to work out what to do with deleting CIs. ServiceNow does not support deleting CIs via the API
//        CIs shouldn't be deleted because may be referenced by other objects in the ServiceNow database
//        (eg changes and incidents).  Therefore the best practice is to change their operational and
//        install status.  The tricky bit is working out how to make this flexible for everyone to use.

//  Delete routine - This function is called when Terraform wants to create a new CI in the ServiceNow CMDB
//
func deleteResourceCmdbCiBusinessApp(resourceData *schema.ResourceData, serviceNowClient interface{}) error {

servicenowClient := serviceNowClient.(*Client)
var ci CmdbCiBusinessAppPost
if err := copyFromTerraformToServiceNowCmdbCiBusinessApp(resourceData, &ci); err != nil {
return err
}

if err := resourceData.Set("install_status", "retired"); err != nil {
return fmt.Errorf("CmdbCiBusinessAppfailed to set install_status field during destroy action %s", err)
}

SnowUrl := CMDBInstanceApi + CiNamecmdb_ci_business_app + "/" + resourceData.Id()
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
func copyFromTerraformToServiceNowCmdbCiBusinessApp (resourceData *schema.ResourceData, ci *CmdbCiBusinessAppPost) error {

attrs := &ci.Result.Attributes
                attrs.OperationalStatus  = resourceData.Get("operational_status").(string)
                attrs.SysUpdatedOn  = resourceData.Get("sys_updated_on").(string)
                attrs.InstallType  = resourceData.Get("install_type").(string)
                attrs.Number  = resourceData.Get("number").(string)
                attrs.DiscoverySource  = resourceData.Get("discovery_source").(string)
                attrs.FirstDiscovered  = resourceData.Get("first_discovered").(string)
                attrs.DueIn  = resourceData.Get("due_in").(string)
                attrs.ItApplicationOwner  = resourceData.Get("it_application_owner.value").(string)
                attrs.ApmBusinessProcess  = resourceData.Get("apm_business_process.value").(string)
                attrs.InvoiceNumber  = resourceData.Get("invoice_number").(string)
                attrs.GlAccount  = resourceData.Get("gl_account").(string)
                attrs.SysCreatedBy  = resourceData.Get("sys_created_by").(string)
                attrs.WarrantyExpiration  = resourceData.Get("warranty_expiration").(string)
                attrs.OrganizationUnitCount  = resourceData.Get("organization_unit_count").(string)
                attrs.Active  = resourceData.Get("active").(string)
                attrs.OwnedBy  = resourceData.Get("owned_by.value").(string)
                attrs.CheckedOut  = resourceData.Get("checked_out").(string)
                attrs.SysDomainPath  = resourceData.Get("sys_domain_path").(string)
                attrs.BusinessUnit  = resourceData.Get("business_unit.value").(string)
                attrs.MaintenanceSchedule  = resourceData.Get("maintenance_schedule.value").(string)
                attrs.ApplicationManager  = resourceData.Get("application_manager.value").(string)
                attrs.CostCenter  = resourceData.Get("cost_center.value").(string)
                attrs.DnsDomain  = resourceData.Get("dns_domain").(string)
                attrs.Assigned  = resourceData.Get("assigned").(string)
                attrs.PurchaseDate  = resourceData.Get("purchase_date").(string)
                attrs.ShortDescription  = resourceData.Get("short_description").(string)
                attrs.LastChangeDate  = resourceData.Get("last_change_date").(string)
                attrs.ManagedBy  = resourceData.Get("managed_by.value").(string)
                attrs.LastDiscovered  = resourceData.Get("last_discovered").(string)
                attrs.CanPrint  = resourceData.Get("can_print").(string)
                attrs.NextAssessmentDate  = resourceData.Get("next_assessment_date").(string)
                attrs.Manufacturer  = resourceData.Get("manufacturer.value").(string)
                attrs.ContractEndDate  = resourceData.Get("contract_end_date").(string)
                attrs.DataClassification  = resourceData.Get("data_classification").(string)
                attrs.Vendor  = resourceData.Get("vendor.value").(string)
                attrs.Currency  = resourceData.Get("currency").(string)
                attrs.Certified  = resourceData.Get("certified").(string)
                attrs.ModelNumber  = resourceData.Get("model_number").(string)
                attrs.AssignedTo  = resourceData.Get("assigned_to.value").(string)
                attrs.StartDate  = resourceData.Get("start_date").(string)
                attrs.BusinessCriticality  = resourceData.Get("business_criticality").(string)
                attrs.SerialNumber  = resourceData.Get("serial_number").(string)
                attrs.Url  = resourceData.Get("url").(string)
                attrs.SupportGroup  = resourceData.Get("support_group.value").(string)
                attrs.TechnologyStack  = resourceData.Get("technology_stack").(string)
                attrs.AudienceType  = resourceData.Get("audience_type").(string)
                attrs.Unverified  = resourceData.Get("unverified").(string)
                attrs.CorrelationId  = resourceData.Get("correlation_id").(string)
                attrs.Attributes  = resourceData.Get("attributes").(string)
                attrs.Asset  = resourceData.Get("asset.value").(string)
                attrs.SoftwareLicense  = resourceData.Get("software_license.value").(string)
                attrs.SkipSync  = resourceData.Get("skip_sync").(string)
                attrs.ActiveUserCount  = resourceData.Get("active_user_count").(string)
                attrs.SupportVendor  = resourceData.Get("support_vendor.value").(string)
                attrs.AppraisalFiscalType  = resourceData.Get("appraisal_fiscal_type").(string)
                attrs.SysUpdatedBy  = resourceData.Get("sys_updated_by").(string)
                attrs.SysCreatedOn  = resourceData.Get("sys_created_on").(string)
                attrs.InstallDate  = resourceData.Get("install_date").(string)
                attrs.AssetTag  = resourceData.Get("asset_tag").(string)
                attrs.Fqdn  = resourceData.Get("fqdn").(string)
                attrs.ChangeControl  = resourceData.Get("change_control.value").(string)
                attrs.EmergencyTier  = resourceData.Get("emergency_tier").(string)
                attrs.ProductSupportStatus  = resourceData.Get("product_support_status").(string)
                attrs.DeliveryDate  = resourceData.Get("delivery_date").(string)
                attrs.InstallStatus  = resourceData.Get("install_status").(string)
                attrs.SupportedBy  = resourceData.Get("supported_by.value").(string)
                attrs.Name  = resourceData.Get("name").(string)
                attrs.Subcategory  = resourceData.Get("subcategory").(string)
                attrs.WorkNotes  = resourceData.Get("work_notes").(string)
                attrs.AssignmentGroup  = resourceData.Get("assignment_group.value").(string)
                attrs.ApplicationType  = resourceData.Get("application_type").(string)
                attrs.ArchitectureType  = resourceData.Get("architecture_type").(string)
                attrs.Platform  = resourceData.Get("platform").(string)
                attrs.PoNumber  = resourceData.Get("po_number").(string)
                attrs.CheckedIn  = resourceData.Get("checked_in").(string)
                attrs.MacAddress  = resourceData.Get("mac_address").(string)
                attrs.Company  = resourceData.Get("company.value").(string)
                attrs.Justification  = resourceData.Get("justification").(string)
                attrs.Department  = resourceData.Get("department.value").(string)
                attrs.Cost  = resourceData.Get("cost").(string)
                attrs.Comments  = resourceData.Get("comments").(string)
                attrs.CmdbSoftwareProductModel  = resourceData.Get("cmdb_software_product_model.value").(string)
                attrs.SysModCount  = resourceData.Get("sys_mod_count").(string)
                attrs.Monitor  = resourceData.Get("monitor").(string)
                attrs.ModelId  = resourceData.Get("model_id.value").(string)
                attrs.IpAddress  = resourceData.Get("ip_address").(string)
                attrs.DuplicateOf  = resourceData.Get("duplicate_of.value").(string)
                attrs.SysTags  = resourceData.Get("sys_tags").(string)
                attrs.CostCc  = resourceData.Get("cost_cc").(string)
                attrs.UserBase  = resourceData.Get("user_base").(string)
                attrs.OrderDate  = resourceData.Get("order_date").(string)
                attrs.Schedule  = resourceData.Get("schedule.value").(string)
                attrs.Due  = resourceData.Get("due").(string)
                attrs.PlatformHost  = resourceData.Get("platform_host.value").(string)
                attrs.Location  = resourceData.Get("location.value").(string)
                attrs.Category  = resourceData.Get("category").(string)
                attrs.FaultCount  = resourceData.Get("fault_count").(string)
                attrs.Age  = resourceData.Get("age").(string)
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
func copyFromServiceNowToTerraformCmdbCiBusinessApp (resourceData *schema.ResourceData, jsonData [] byte) error {
ci := CmdbCiBusinessAppGet{}
if err := json.Unmarshal(jsonData, &ci); err != nil {
//resourceData.SetId("")
//return err
}

resourceData.SetId(GetSysId(jsonData))
attrs := ci.Result.Attributes
                _ = resourceData.Set("operational_status", attrs.OperationalStatus)
                _ = resourceData.Set("sys_updated_on", attrs.SysUpdatedOn)
                _ = resourceData.Set("install_type", attrs.InstallType)
                _ = resourceData.Set("number", attrs.Number)
                _ = resourceData.Set("discovery_source", attrs.DiscoverySource)
                _ = resourceData.Set("first_discovered", attrs.FirstDiscovered)
                _ = resourceData.Set("due_in", attrs.DueIn)
                _ = resourceData.Set("it_application_owner", StructToMap(attrs.ItApplicationOwner))
                _ = resourceData.Set("apm_business_process", StructToMap(attrs.ApmBusinessProcess))
                _ = resourceData.Set("invoice_number", attrs.InvoiceNumber)
                _ = resourceData.Set("gl_account", attrs.GlAccount)
                _ = resourceData.Set("sys_created_by", attrs.SysCreatedBy)
                _ = resourceData.Set("warranty_expiration", attrs.WarrantyExpiration)
                _ = resourceData.Set("organization_unit_count", attrs.OrganizationUnitCount)
                _ = resourceData.Set("active", attrs.Active)
                _ = resourceData.Set("owned_by", StructToMap(attrs.OwnedBy))
                _ = resourceData.Set("checked_out", attrs.CheckedOut)
                _ = resourceData.Set("sys_domain_path", attrs.SysDomainPath)
                _ = resourceData.Set("business_unit", StructToMap(attrs.BusinessUnit))
                _ = resourceData.Set("maintenance_schedule", StructToMap(attrs.MaintenanceSchedule))
                _ = resourceData.Set("application_manager", StructToMap(attrs.ApplicationManager))
                _ = resourceData.Set("cost_center", StructToMap(attrs.CostCenter))
                _ = resourceData.Set("dns_domain", attrs.DnsDomain)
                _ = resourceData.Set("assigned", attrs.Assigned)
                _ = resourceData.Set("purchase_date", attrs.PurchaseDate)
                _ = resourceData.Set("short_description", attrs.ShortDescription)
                _ = resourceData.Set("last_change_date", attrs.LastChangeDate)
                _ = resourceData.Set("managed_by", StructToMap(attrs.ManagedBy))
                _ = resourceData.Set("last_discovered", attrs.LastDiscovered)
                _ = resourceData.Set("can_print", attrs.CanPrint)
                _ = resourceData.Set("next_assessment_date", attrs.NextAssessmentDate)
                _ = resourceData.Set("manufacturer", StructToMap(attrs.Manufacturer))
                _ = resourceData.Set("contract_end_date", attrs.ContractEndDate)
                _ = resourceData.Set("data_classification", attrs.DataClassification)
                _ = resourceData.Set("vendor", StructToMap(attrs.Vendor))
                _ = resourceData.Set("currency", attrs.Currency)
                _ = resourceData.Set("certified", attrs.Certified)
                _ = resourceData.Set("model_number", attrs.ModelNumber)
                _ = resourceData.Set("assigned_to", StructToMap(attrs.AssignedTo))
                _ = resourceData.Set("start_date", attrs.StartDate)
                _ = resourceData.Set("business_criticality", attrs.BusinessCriticality)
                _ = resourceData.Set("serial_number", attrs.SerialNumber)
                _ = resourceData.Set("url", attrs.Url)
                _ = resourceData.Set("support_group", StructToMap(attrs.SupportGroup))
                _ = resourceData.Set("technology_stack", attrs.TechnologyStack)
                _ = resourceData.Set("audience_type", attrs.AudienceType)
                _ = resourceData.Set("unverified", attrs.Unverified)
                _ = resourceData.Set("correlation_id", attrs.CorrelationId)
                _ = resourceData.Set("attributes", attrs.Attributes)
                _ = resourceData.Set("asset", StructToMap(attrs.Asset))
                _ = resourceData.Set("software_license", StructToMap(attrs.SoftwareLicense))
                _ = resourceData.Set("skip_sync", attrs.SkipSync)
                _ = resourceData.Set("active_user_count", attrs.ActiveUserCount)
                _ = resourceData.Set("support_vendor", StructToMap(attrs.SupportVendor))
                _ = resourceData.Set("appraisal_fiscal_type", attrs.AppraisalFiscalType)
                _ = resourceData.Set("sys_updated_by", attrs.SysUpdatedBy)
                _ = resourceData.Set("sys_created_on", attrs.SysCreatedOn)
                _ = resourceData.Set("install_date", attrs.InstallDate)
                _ = resourceData.Set("asset_tag", attrs.AssetTag)
                _ = resourceData.Set("fqdn", attrs.Fqdn)
                _ = resourceData.Set("change_control", StructToMap(attrs.ChangeControl))
                _ = resourceData.Set("emergency_tier", attrs.EmergencyTier)
                _ = resourceData.Set("product_support_status", attrs.ProductSupportStatus)
                _ = resourceData.Set("delivery_date", attrs.DeliveryDate)
                _ = resourceData.Set("install_status", attrs.InstallStatus)
                _ = resourceData.Set("supported_by", StructToMap(attrs.SupportedBy))
                _ = resourceData.Set("name", attrs.Name)
                _ = resourceData.Set("subcategory", attrs.Subcategory)
                _ = resourceData.Set("work_notes", attrs.WorkNotes)
                _ = resourceData.Set("assignment_group", StructToMap(attrs.AssignmentGroup))
                _ = resourceData.Set("application_type", attrs.ApplicationType)
                _ = resourceData.Set("architecture_type", attrs.ArchitectureType)
                _ = resourceData.Set("platform", attrs.Platform)
                _ = resourceData.Set("po_number", attrs.PoNumber)
                _ = resourceData.Set("checked_in", attrs.CheckedIn)
                _ = resourceData.Set("mac_address", attrs.MacAddress)
                _ = resourceData.Set("company", StructToMap(attrs.Company))
                _ = resourceData.Set("justification", attrs.Justification)
                _ = resourceData.Set("department", StructToMap(attrs.Department))
                _ = resourceData.Set("cost", attrs.Cost)
                _ = resourceData.Set("comments", attrs.Comments)
                _ = resourceData.Set("cmdb_software_product_model", StructToMap(attrs.CmdbSoftwareProductModel))
                _ = resourceData.Set("sys_mod_count", attrs.SysModCount)
                _ = resourceData.Set("monitor", attrs.Monitor)
                _ = resourceData.Set("model_id", StructToMap(attrs.ModelId))
                _ = resourceData.Set("ip_address", attrs.IpAddress)
                _ = resourceData.Set("duplicate_of", StructToMap(attrs.DuplicateOf))
                _ = resourceData.Set("sys_tags", attrs.SysTags)
                _ = resourceData.Set("cost_cc", attrs.CostCc)
                _ = resourceData.Set("user_base", attrs.UserBase)
                _ = resourceData.Set("order_date", attrs.OrderDate)
                _ = resourceData.Set("schedule", StructToMap(attrs.Schedule))
                _ = resourceData.Set("due", attrs.Due)
                _ = resourceData.Set("platform_host", StructToMap(attrs.PlatformHost))
                _ = resourceData.Set("location", StructToMap(attrs.Location))
                _ = resourceData.Set("category", attrs.Category)
                _ = resourceData.Set("fault_count", attrs.FaultCount)
                _ = resourceData.Set("age", attrs.Age)
                _ = resourceData.Set("lease_id", attrs.LeaseId)
return nil
}
