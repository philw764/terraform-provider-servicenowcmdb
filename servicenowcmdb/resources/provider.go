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
			"servicenowcmdb_server":                       ResourceServer(),
			"servicenowcmdb_cim_server":                   ResourceCimServer(),
			"servicenowcmdb_server_hardware":              ResourceServerHardware(),
			"servicenowcmdb_server_chassis":               ResourceServerChassis(),
			"servicenowcmdb_network_appliance_hardware":   ResourceNetworkApplianceHardware(),
			"servicenowcmdb_server_tape_unit":             ResourceServerTapeUnit(),
			"servicenowcmdb_virtualization_server":        ResourceVirtualizationServer(),
			"servicenowcmdb_hyperv_server":                ResourceHypervServer(),
			"servicenowcmdb_vmware_vcenter_server_object": ResourceVmwareVcenterServerObject(),
			"servicenowcmdb_esx_server":                   ResourceEsxServer(),
			"servicenowcmdb_nutanix_host":                 ResourceNutanixHost(),
			"servicenowcmdb_osx_server":                   ResourceOsxServer(),
			"servicenowcmdb_unix_server":                  ResourceUnixServer(),
			"servicenowcmdb_solaris_server":               ResourceSolarisServer(),
			"servicenowcmdb_hpux_server":                  ResourceHpuxServer(),
			"servicenowcmdb_aix_server":                   ResourceAixServer(),
			"servicenowcmdb_isam_server":                  ResourceIsamServer(),
			"servicenowcmdb_ibm_zos_server":               ResourceIbmZosServer(),
			"servicenowcmdb_load_balancer":                ResourceLoadBalancer(),
			"servicenowcmdb_alteon":                       ResourceAlteon(),
			"servicenowcmdb_f5_bigip_gtm":                 ResourceF5BigipGtm(),
			"servicenowcmdb_network_load_balancer":        ResourceNetworkLoadBalancer(),
			"servicenowcmdb_cisco_css":                    ResourceCiscoCss(),
			"servicenowcmdb_a10_load_balancer":            ResourceA10LoadBalancer(),
			"servicenowcmdb_f5_bigip_ltm":                 ResourceF5BigipLtm(),
			"servicenowcmdb_cisco_csm":                    ResourceCiscoCsm(),
			"servicenowcmdb_f5_bigip":                     ResourceF5Bigip(),
			"servicenowcmdb_isa_server":                   ResourceIsaServer(),
			"servicenowcmdb_radware_load_balancer":        ResourceRadwareLoadBalancer(),
			"servicenowcmdb_ace":                          ResourceAce(),
			"servicenowcmdb_cisco_gss":                    ResourceCiscoGss(),
			"servicenowcmdb_citrix_netscaler":             ResourceCitrixNetscaler(),
			"servicenowcmdb_storage_server":               ResourceStorageServer(),
			"servicenowcmdb_linux_server":                 ResourceLinuxServer(),
			"servicenowcmdb_ibm_hmc_server":               ResourceIbmHmcServer(),
			"servicenowcmdb_ibm_mainframe":                ResourceIbmMainframe(),
			"servicenowcmdb_ibm_frame":                    ResourceIbmFrame(),
			"servicenowcmdb_storage_node_element":         ResourceStorageNodeElement(),
			"servicenowcmdb_data_power_hosting_server":    ResourceDataPowerHostingServer(),
			"servicenowcmdb_windows_server":               ResourceWindowsServer(),
			"servicenowcmdb_netware_server":               ResourceNetwareServer(),
			"servicenowcmdb_ibm_mainframe_lpar":           ResourceIbmMainframeLpar(),
		},
		DataSourcesMap: map[string]*schema.Resource{},
		ConfigureFunc:  configure,
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
