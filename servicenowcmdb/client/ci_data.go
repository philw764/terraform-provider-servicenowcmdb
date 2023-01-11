package client

import (
	"encoding/json"
	"errors"
	"fmt"
)

type attribute struct {
	AttrName    string
	AttrValue   string
	AttrType    string
	AttrLabel   string
	AttrElement string
}

type GetSysID struct {
	Result []struct {
		SysId string `json:"sys_id"`
		Name  string `json:"name"`
	} `json:"result"`
}

type DataResp struct {
	Result struct {
		Attributes struct {
			Attribute interface{}
		} `json:"attributes"`
	} `json:"result"`
}

type CIGet struct {
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

func GetSysIDForCI(class string, filter string, name string, client *Client) (string, error) {
	getSysID := GetSysID{}

	endpoint := CMDBInstanceApi + class + "?sysparm_query=name=" + name + "^" + filter

	if jsonResponse, err := client.RequestJSON("GET", endpoint, nil); err != nil {
		return "", errors.New("Failed to get sysid for Class:" + class + " CI:" + name)
	} else {
		if err := json.Unmarshal(jsonResponse, &getSysID); err != nil {
			return "", errors.New("Failed to unmarshal while getting sysid for Class:" + class + " CI:" + name)
		}
	}

	if len(getSysID.Result) == 1 {
		sysid := getSysID.Result[0].SysId
		return sysid, nil
	} else if len(getSysID.Result) > 1 {
		return "", errors.New("More than one record found, refine your search criteria :" + fmt.Sprintf("%d getSysID found", len(getSysID.Result)))
	} else {
		return "", errors.New("Record for class " + class + " not found")
	}
}

func GetDataForCI(class string, sysid string, client *Client) (map[string]interface{}, error) {

	endpoint := CMDBInstanceApi + class + "/" + sysid

	jsonData, err := client.RequestJSON("GET", endpoint, nil)
	if err != nil {
		return nil, errors.New("Failed to get data for ci:" + class + "  sysid:" + sysid)
	}

	var mapBuf = make(map[string]interface{})
	var jsonBuf []byte

	if err = json.Unmarshal(jsonData, &mapBuf); err != nil {
		return nil, err
	}
	if jsonBuf, err = json.Marshal(mapBuf["result"]); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(jsonBuf, &mapBuf); err != nil {
		return nil, err
	}
	if jsonBuf, err = json.Marshal(mapBuf["attributes"]); err != nil {
		return nil, err
	}
	if err = json.Unmarshal(jsonBuf, &mapBuf); err != nil {
		return nil, err
	}

	return mapBuf, nil
}
