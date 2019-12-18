package resources

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
)

const CMDBInstanceApi = "/api/now/cmdb/instance/"

// Client HTTP Connector
type Client struct {
	BaseURL string
	Auth    string
}

var AttributesToExclude = []string{
	"sys_updated_on",
	"sys_updated_by",
	"sys_created_on",
	"sys_created_by",
	"sys_domain_path", //system related field, TODO: Might need to consider this for application Domains
	"sys_class_name",  //
	"sys_id",          //This field CANNOT be in a http POST request or the "create" action fails
	"sys_class_path",
	"sys_mod_count",
	"sys_tags",
	"",
}

type ModelIDRel struct {
	Result struct {
		Attributes struct {
			ModelId struct {
				DisplayValue string `json:"display_value"`
				Link         string `json:"link"`
				Value        string `json:"value"`
			} `json:"model_id"`
		} `json:"Attributes"`
	} `json:"result"`
}

type AssetList struct {
	Result struct {
		Attributes struct {
			Asset struct {
				DisplayValue string `json:"display_value"`
				Link         string `json:"link"`
				Value        string `json:"value"`
			} `json:"asset"`
		} `json:"Attributes"`
	} `json:"result"`
}

type OutboundRelations struct {
	OutboundRel []struct {
		SysID string `json:"sys_id"`
		Type  struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"type"`
		Target struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"target"`
	} `json:"outbound_relations"`
}

type InboundRelations []struct {
	InboundRel []struct {
		SysID string `json:"sys_id"`
		Type  struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"type"`
		Target struct {
			DisplayValue string `json:"display_value"`
			Link         string `json:"link"`
			Value        string `json:"value"`
		} `json:"target"`
	} `json:"inbound_relations"`
}

type SysId struct {
	Result struct {
		Attributes struct {
			SysId string `json:"sys_id"`
		}
	} `json:"result"`
}

func UpdateResourceDataWithBaseAttrs(resourceData *schema.ResourceData, jsonData []byte) error {
	baseCIAttributes := CmdbCiPost{}
	if err := json.Unmarshal(jsonData, &baseCIAttributes); err != nil {
		resourceData.SetId("")
		return err
	}
	resourceData.SetId(GetSysId(jsonData))

	attrs := baseCIAttributes.Result.Attributes
	log.Printf("***************************In UpdateResourceDataWithBaseAttrs:1")
	_ = resourceData.Set("name", attrs.Name)
	_ = resourceData.Set("short_description", attrs.ShortDescription)
	_ = resourceData.Set("ip_address", attrs.IpAddress)
	_ = resourceData.Set("discovery_source", attrs.DiscoverySource)
	_ = resourceData.Set("operational_status", attrs.OperationalStatus)
	_ = resourceData.Set("install_status", attrs.InstallStatus)
	_ = resourceData.Set("sys_updated_on", attrs.SysUpdatedBy)
	_ = resourceData.Set("first_discovered", attrs.FirstDiscovered)
	_ = resourceData.Set("sys_updated_by", attrs.SysUpdatedBy)
	_ = resourceData.Set("due_in", attrs.DueIn)
	_ = resourceData.Set("sys_created_on", attrs.SysCreatedOn)
	//_ = resourceData.Set("sys_domain",  			attrs.Sysdomain)
	_ = resourceData.Set("install_date", attrs.InstallDate)
	_ = resourceData.Set("invoice_number", attrs.InvoiceNumber)
	_ = resourceData.Set("gl_account", attrs.GlAccount)
	_ = resourceData.Set("sys_created_by", attrs.SysCreatedBy)
	_ = resourceData.Set("warranty_expiration", attrs.WarrantyExpiration)
	_ = resourceData.Set("asset_tag", attrs.AssetTag)
	_ = resourceData.Set("fqdn", attrs.Fqdn)
	//_ = resourceData.Set("change_control",  		attrs.ApprovalGroup)  //This is a reference to the group table
	//_ = resourceData.Set("owned_by", attrs.OwnedBy)
	_ = resourceData.Set("checked_out", attrs.CheckedOut)
	_ = resourceData.Set("delivery_date", attrs.DeliveryDate)
	//_ = resourceData.Set("maintenance_schedule",  	attrs.Maintenanceschedule)
	//_ = resourceData.Set("cost_center",  			attrs.Costcenter)
	//_ = resourceData.Set("supported_by",  			attrs.Supportedby)
	_ = resourceData.Set("dns_domain", attrs.DnsDomain)
	//_ = resourceData.Set("assigned",  				attrs.Assigned)
	_ = resourceData.Set("purchase_date", attrs.PurchaseDate)
	_ = resourceData.Set("sub_category", attrs.Subcategory)
	//_ = resourceData.Set("assignment_group",  		attrs.Assignmentgroup)
	//_ = resourceData.Set("managed_by",  			attrs.Managedby)
	_ = resourceData.Set("last_discovered", attrs.LastDiscovered)
	_ = resourceData.Set("can_print", attrs.CanPrint)
	//_ = resourceData.Set("manufacturer",  			attrs.Manufacturer)
	//_ = resourceData.Set("po_number",  			attrs.Ponumber)
	_ = resourceData.Set("checked_in", attrs.CheckedIn)
	//_ = resourceData.Set("vendor",  				attrs.Vendor)
	_ = resourceData.Set("mac_address", attrs.MacAddress)
	//_ = resourceData.Set("company",  				attrs.Company)
	//_ = resourceData.Set("model_number",  			attrs.Modelnumber)
	_ = resourceData.Set("justification", attrs.Justification)
	//_ = resourceData.Set("department", 		 	attrs.Department)
	//_ = resourceData.Set("assigned_to",  			attrs.Assignedto)  // needs special processing
	_ = resourceData.Set("start_date", attrs.StartDate)
	//_ = resourceData.Set("cost",  					attrs.Cost)
	_ = resourceData.Set("comments", attrs.Comments)
	_ = resourceData.Set("serialnumber", attrs.SerialNumber)
	_ = resourceData.Set("monitor", attrs.Monitor)
	//_ = resourceData.Set("model_number",	  		attrs.Modelnumber)
	//_ = resourceData.Set("duplicate_of",  			attrs.Duplicateof)
	//_ = resourceData.Set("cost_cc",			  	attrs.Costcc)
	//_ = resourceData.Set("support_group",  		attrs.Supportgroup)
	_ = resourceData.Set("order_date", attrs.OrderDate)
	//_ = resourceData.Set("schedule",  				attrs.Schedule)
	_ = resourceData.Set("due", attrs.Due)
	_ = resourceData.Set("unverified", attrs.Unverified)
	//_ = resourceData.Set("correlation_id",  		attrs.Correlationid)
	//_ = resourceData.Set("location",  				attrs.Location)
	_ = resourceData.Set("category", attrs.Category)
	_ = resourceData.Set("fault_count", attrs.FaultCount)
	_ = resourceData.Set("lease_id", attrs.LeaseId)
	log.Printf("In UpdateResourceDataWithBaseAttrs:10")

	//err := UpdateAssetInResourceData(resourceData, )
	return nil
}

func UpdateBaseAttrsFromResourceData(resourceData *schema.ResourceData, baseCIAttributes *CmdbCiPost) error {

	//baseCIAttributes 	:= CmdbCi{}

	attrs := &baseCIAttributes.Result.Attributes
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:1")
	attrs.Name = resourceData.Get("name").(string)
	attrs.ShortDescription = resourceData.Get("short_description").(string)
	attrs.IpAddress = resourceData.Get("ip_address").(string)
	attrs.DiscoverySource = resourceData.Get("discovery_source").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:2")
	attrs.OperationalStatus = resourceData.Get("operational_status").(string)
	attrs.InstallStatus = resourceData.Get("install_status").(string)
	//attrs.Sysupdatedon			=	resourceData.Get("sys_updated_on"		).(string)
	attrs.FirstDiscovered = resourceData.Get("first_discovered").(string)
	attrs.SysUpdatedBy = resourceData.Get("sys_updated_by").(string)
	attrs.DueIn = resourceData.Get("due_in").(string)
	attrs.SysCreatedOn = resourceData.Get("sys_created_on").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:3")
	//attrs.Sysdomain				=	resourceData.Get("sys_domain"			).(string)
	//attrs.Installdate			=	resourceData.Get("install_date"		).(string)
	attrs.InvoiceNumber = resourceData.Get("invoice_number").(string)
	attrs.GlAccount = resourceData.Get("gl_account").(string)
	attrs.SysCreatedBy = resourceData.Get("sys_created_by").(string)
	attrs.WarrantyExpiration = resourceData.Get("warranty_expiration").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:4")
	attrs.AssetTag = resourceData.Get("asset_tag").(string)
	attrs.Fqdn = resourceData.Get("fqdn").(string)
	//attrs.ApprovalGroup		=	resourceData.Get("change_control"		).(string)
	//attrs.OwnedBy = resourceData.Get("owned_by").(string)
	attrs.CheckedOut = resourceData.Get("checked_out").(string)
	attrs.DeliveryDate = resourceData.Get("delivery_date").(string)
	//attrs.Maintenanceschedule	=	resourceData.Get("maintenance_schedule").(string)
	//attrs.Costcenter			=	resourceData.Get("cost_center"			).(string)
	//attrs.Supportedby			=	resourceData.Get("supported_by"		).(string)
	attrs.DnsDomain = resourceData.Get("dns_domain").(string)
	//attrs.Assigned				=	resourceData.Get("assigned"			).(string)
	//attrs.Purchasedate			=	resourceData.Get("purchase_date"		).(string)
	//attrs.Subcategory = resourceData.Get("subcategory").(string)
	//attrs.Assignmentgroup		=	resourceData.Get("assignment_group"	).(string)
	//attrs.Managedby				=	resourceData.Get("managed_by"			).(string)
	//attrs.Lastdiscovered		=	resourceData.Get("last_discovered"		).(string)
	attrs.CanPrint = resourceData.Get("can_print").(string)
	//attrs.Manufacturer			=	resourceData.Get("manufacturer"		).(string)
	//attrs.Ponumber				=	resourceData.Get("po_number"			).(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:5")
	attrs.CheckedIn = resourceData.Get("checked_in").(string)
	//attrs.Vendor				=	resourceData.Get("vendor"				).(string)
	attrs.MacAddress = resourceData.Get("mac_address").(string)
	//attrs.Company				=	resourceData.Get("company"				).(string)
	//attrs.Modelnumber			=	resourceData.Get("model_number"		).(string)
	attrs.Justification = resourceData.Get("justification").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:6")
	//attrs.Department			=	resourceData.Get("department"			).(string)
	//attrs.Assignedto			=	resourceData.Get("assigned_to"			).(string)  Needs special processing
	attrs.StartDate = resourceData.Get("start_date").(string)
	//attrs.Cost					=	resourceData.Get("cost"				).(string)
	//attrs.Comments = resourceData.Get("comments").(string)
	//attrs.SerialNumber = resourceData.Get("serial number").(string)
	//attrs.Monitor = resourceData.Get("monitor").(string)
	//attrs.Modelnumber			=	resourceData.Get("model_number"		).(string)
	//attrs.Duplicateof			=	resourceData.Get("duplicate_of	"		).(string)
	//attrs.CostCc = resourceData.Get("cost_cc").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:7")

	//attrs.Supportgroup			=	resourceData.Get("support_group"		).(string)
	attrs.OrderDate = resourceData.Get("order_date").(string)
	//attrs.Schedule				=	resourceData.Get("schedule"			).(string)
	attrs.Due = resourceData.Get("due").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:8")
	attrs.Unverified = resourceData.Get("unverified").(string)
	//attrs.Correlationid			=	resourceData.Get("correlation_id"		).(string)
	//attrs.Location				=	resourceData.Get("location"			).(string)
	attrs.Category = resourceData.Get("category").(string)
	attrs.FaultCount = resourceData.Get("fault_count").(string)
	attrs.LeaseId = resourceData.Get("lease_id").(string)
	log.Printf("***************************In UpdateBaseAttrsFromResourceData:10")

	return nil
}

func GetSysId(jsonData []byte) string {
	sysId := SysId{}
	if err := json.Unmarshal(jsonData, &sysId); err != nil {
		return ""
	}
	return sysId.Result.Attributes.SysId
}

// NewClient is a factory method used to return a new ServiceNowClient.
func NewClient(baseURL string, username string, password string) *Client {
	// Concatenate username + password to create a basic authorization header.
	credentials := username + ":" + password
	return &Client{
		BaseURL: baseURL,
		Auth:    "Basic " + base64.StdEncoding.EncodeToString([]byte(credentials)),
	}
}
func StructToMap(in interface{}) map[string]interface{} {
	m := make(map[string]interface{})
	j, _ := json.Marshal(in)
	json.Unmarshal(j, &m)
	return m
}

// requestJSON execute an HTTP request and returns the raw response data.
func (client *Client) RequestJSON(method string, path string, jsonData interface{}) ([]byte, error) {
	log.Printf("[DEBUG] ************************** Entering RequestJSON")
	var data *bytes.Buffer

	if jsonData != nil {
		jsonValue, _ := json.Marshal(jsonData)
		log.Printf("[DEBUG] ********************* This is jsonValue :%s", jsonValue)
		data = bytes.NewBuffer(jsonValue)
	} else {
		data = bytes.NewBuffer(nil)
	}

	log.Printf("[DEBUG] ********* This is inJSON data :%s", data)

	log.Printf("[DEBUG] ********* This is in RequestJSON path value:%s", client.BaseURL+path)
	request, _ := http.NewRequest(method, client.BaseURL+path, data)

	// Add the needed headers.
	request.Header.Set("Authorization", client.Auth)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")

	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	log.Printf("[DEBUG] ***** this is just after httpClient.Do")
	if err != nil {
		log.Printf("[DEBUG] ***** This is in err after httpClient.Do:%s", err)
		return nil, err
	}
	responseData, _ := ioutil.ReadAll(response.Body)

	log.Printf("[DEBUG]********** this is the response data:%s", response.Header)
	if response.StatusCode >= 300 || response.StatusCode < 200 {
		log.Printf("[DEBUG] ************* This is in err after after statusCode check:%s", response.StatusCode)
		return nil, fmt.Errorf("HTTP response status %s, %s", response.Status, responseData)
	}

	return responseData, nil
}
