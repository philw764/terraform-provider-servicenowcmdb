package resources

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var CMDBConfigItemKeys = make(map[string]string)

func CMDBSchema() map[string]*schema.Schema {

	x := map[string]*schema.Schema{
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
		"extended_attrs": {
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"owned_by": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"delivery_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"maintenance_schedule": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"install_status": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"cost_center": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"attested_by": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"supported_by": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
		"subcategory": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"life_cycle_stage": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"short_description": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"assignment_group": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"managed_by": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"managed_by_group": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"life_cycle_stage_status": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"mac_address": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"company": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"assigned_to": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"ip_address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"duplicate_of": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"order_date": {
			Type:     schema.TypeString,
			Computed: true,
		},
		"schedule": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
		},
		"asset": {
			Type:     schema.TypeList,
			Optional: true,
			Computed: true,
			Elem:     &schema.Resource{Schema: referenceFieldSchema()},
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
	}
	return x
}

func referenceFieldSchema() map[string]*schema.Schema {
	sch := map[string]*schema.Schema{
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
	}
	return sch
}

func init() {
	for k := range CMDBSchema() {
		CMDBConfigItemKeys[k] = "base"
	}
}
