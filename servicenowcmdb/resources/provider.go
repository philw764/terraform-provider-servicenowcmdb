package resources

import (
"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
"github.com/hashicorp/terraform-plugin-sdk/terraform"
"log"
)

//Terraform Provider for CI Classes in a ServiceNow instance.
func Provider() terraform.ResourceProvider {
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
   "servicenowcmdb_cmdb_ci":   ResourceCmdbCi(),
   "servicenowcmdb_cmdb_ci_oslv_local_image":   ResourceCmdbCiOslvLocalImage(),
   "servicenowcmdb_cmdb_ci_docker_local_image":   ResourceCmdbCiDockerLocalImage(),
   "servicenowcmdb_cmdb_ci_business_app":   ResourceCmdbCiBusinessApp(),
   "servicenowcmdb_cmdb_ci_acc":   ResourceCmdbCiAcc(),
   "servicenowcmdb_cmdb_ci_port":   ResourceCmdbCiPort(),
   "servicenowcmdb_cmdb_ci_information_object":   ResourceCmdbCiInformationObject(),
   "servicenowcmdb_cmdb_ci_cluster_node":   ResourceCmdbCiClusterNode(),
   "servicenowcmdb_cmdb_ci_unix_cluster_node":   ResourceCmdbCiUnixClusterNode(),
   "servicenowcmdb_cmdb_ci_storage_cluster_node":   ResourceCmdbCiStorageClusterNode(),
   "servicenowcmdb_cmdb_ci_win_cluster_node":   ResourceCmdbCiWinClusterNode(),
   "servicenowcmdb_cmdb_ci_san_zone":   ResourceCmdbCiSanZone(),
   "servicenowcmdb_cmdb_ci_storage_export":   ResourceCmdbCiStorageExport(),
   "servicenowcmdb_cmdb_ci_san_export":   ResourceCmdbCiSanExport(),
   "servicenowcmdb_cmdb_ci_iscsi_export":   ResourceCmdbCiIscsiExport(),
   "servicenowcmdb_cmdb_ci_fc_export":   ResourceCmdbCiFcExport(),
   "servicenowcmdb_cmdb_ci_hardware":   ResourceCmdbCiHardware(),
   "servicenowcmdb_cmdb_ci_printer":   ResourceCmdbCiPrinter(),
   "servicenowcmdb_cmdb_ci_mfp_printer":   ResourceCmdbCiMfpPrinter(),
   "servicenowcmdb_cmdb_ci_personal_printer":   ResourceCmdbCiPersonalPrinter(),
   "servicenowcmdb_cmdb_ci_computer":   ResourceCmdbCiComputer(),
   "servicenowcmdb_cmdb_ci_ucs_rack_unit":   ResourceCmdbCiUcsRackUnit(),
   "servicenowcmdb_cmdb_ci_server":   ResourceCmdbCiServer(),
   "servicenowcmdb_cmdb_ci_storage_node_element":   ResourceCmdbCiStorageNodeElement(),
   "servicenowcmdb_cmdb_ci_datapower_server":   ResourceCmdbCiDatapowerServer(),
   "servicenowcmdb_cmdb_ci_unix_server":   ResourceCmdbCiUnixServer(),
   "servicenowcmdb_cmdb_ci_solaris_server":   ResourceCmdbCiSolarisServer(),
   "servicenowcmdb_cmdb_ci_aix_server":   ResourceCmdbCiAixServer(),
   "servicenowcmdb_cmdb_ci_hpux_server":   ResourceCmdbCiHpuxServer(),
   "servicenowcmdb_cmdb_ci_netware_server":   ResourceCmdbCiNetwareServer(),
   "servicenowcmdb_cmdb_ci_mainframe":   ResourceCmdbCiMainframe(),
   "servicenowcmdb_cmdb_ci_cim_server":   ResourceCmdbCiCimServer(),
   "servicenowcmdb_cmdb_ci_mainframe_lpar":   ResourceCmdbCiMainframeLpar(),
   "servicenowcmdb_cmdb_ci_lb":   ResourceCmdbCiLb(),
   "servicenowcmdb_cmdb_ci_lb_a10":   ResourceCmdbCiLbA10(),
   "servicenowcmdb_cmdb_ci_lb_network":   ResourceCmdbCiLbNetwork(),
   "servicenowcmdb_cmdb_ci_lb_alteon":   ResourceCmdbCiLbAlteon(),
   "servicenowcmdb_cmdb_ci_lb_cisco_css":   ResourceCmdbCiLbCiscoCss(),
   "servicenowcmdb_cmdb_ci_lb_f5_gtm":   ResourceCmdbCiLbF5Gtm(),
   "servicenowcmdb_cmdb_ci_lb_netscaler":   ResourceCmdbCiLbNetscaler(),
   "servicenowcmdb_cmdb_ci_lb_ace":   ResourceCmdbCiLbAce(),
   "servicenowcmdb_cmdb_ci_lb_f5_ltm":   ResourceCmdbCiLbF5Ltm(),
   "servicenowcmdb_cmdb_ci_lb_isa":   ResourceCmdbCiLbIsa(),
   "servicenowcmdb_cmdb_ci_lb_cisco_csm":   ResourceCmdbCiLbCiscoCsm(),
   "servicenowcmdb_cmdb_ci_lb_cisco_gss":   ResourceCmdbCiLbCiscoGss(),
   "servicenowcmdb_cmdb_ci_lb_radware":   ResourceCmdbCiLbRadware(),
   "servicenowcmdb_cmdb_ci_lb_bigip":   ResourceCmdbCiLbBigip(),
   "servicenowcmdb_cmdb_ci_server_hardware":   ResourceCmdbCiServerHardware(),
   "servicenowcmdb_cmdb_ci_chassis_server":   ResourceCmdbCiChassisServer(),
   "servicenowcmdb_cmdb_ci_tape_server":   ResourceCmdbCiTapeServer(),
},
DataSourcesMap: map[string]*schema.Resource{

},
ConfigureFunc: configure,
}
}

func configure(data *schema.ResourceData) (interface{}, error) {
// Create a new client to talk to the instance.
log.Printf("[DEBUG] In configure")
servicenowClient := NewClient(
data.Get("instance_url").(string),
data.Get("username").(string),
data.Get("password").(string))
return servicenowClient, nil
}
