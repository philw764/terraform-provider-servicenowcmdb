package flags

type Key struct {
	CiName      string
	CiAttribute string
}

type FlagMap map[Key]string

var flagAttrMap = FlagMap{
	{"ALL", "skip_sync"}:           "Computed:true",
	{"ALL", "name"}:                "Required:true",
	{"ALL", "SkipSync"}:            "Computed:true",
	{"ALL", "OperationalStatus"}:   "Optional:true",
	{"ALL", "SysUpdatedOn"}:        "Computed:true",
	{"ALL", "SysUpdatedBy"}:        "Computed:true",
	{"ALL", "DiscoverySource"}:     "Optional:true,Default:\"Terraform\"",
	{"ALL", "SysUpdatedOn"}:        "Computed:true",
	{"ALL", "FirstDiscovered"}:     "Computed:true",
	{"ALL", "DueIn"}:               "Computed:true",
	{"ALL", "SysCreatedOn"}:        "Computed:true",
	{"ALL", "InstallDate"}:         "Computed:true",
	{"ALL", "InvoiceNumber"}:       "Optional:true",
	{"ALL", "GlAccount"}:           "Optional:true",
	{"ALL", "SysCreatedBy"}:        "Computed:true",
	{"ALL", "WarrantyExpiration"}:  "Optional:true",
	{"ALL", "AssetTag"}:            "Optional:true",
	{"ALL", "Fqdn"}:                "Optional:true",
	{"ALL", "SysUpdatedOn"}:        "Computed:true",
	{"ALL", "OwnedBy"}:             "Optional:true",
	{"ALL", "ChangeControl"}:       "Optional:true",
	{"ALL", "CheckedOut"}:          "Optional:true",
	{"ALL", "SysDomainPath"}:       "Optional:true",
	{"ALL", "DeliveryDate"}:        "Computed:true",
	{"ALL", "MaintenanceSchedule"}: "Optional:true",
	{"ALL", "CostCenter"}:          "Optional:true",
	{"ALL", "SupportedBy"}:         "Optional:true",
	{"ALL", "DnsDomain"}:           "Optional:true",
	{"ALL", "Name"}:                "Required:true",
	{"ALL", "Assigned"}:            "Optional:true",
	{"ALL", "Subcategory"}:         "Computed:true",
	{"ALL", "ShortDescription"}:    "Optional:true",
	{"ALL", "AssignmentGroup"}:     "Optional:true",
	{"ALL", "ManagedBy"}:           "Optional:true",
	{"ALL", "LastDiscovered"}:      "Computed:true",
	{"ALL", "CanPrint"}:            "Computed:true",
	{"ALL", "Manufacturer"}:        "Optional:true",
	{"ALL", "PoNumber"}:            "Computed:true",
	{"ALL", "CheckedIn"}:           "Computed:true",
	{"ALL", "Vendor"}:              "Computed:true",
	{"ALL", "MacAddress"}:          "Computed:true",
	{"ALL", "Company"}:             "Computed:true",
	{"ALL", "ModelNumber"}:         "Computed:true",
	{"ALL", "Department"}:          "Computed:true",
	{"ALL", "AssignedTo"}:          "Computed:true",
	{"ALL", "StartDate"}:           "Computed:true",
	{"ALL", "Cost"}:                "Computed:true",
	{"ALL", "Comments"}:            "Computed:true",
	{"ALL", "SysModCount"}:         "Computed:true",
	{"ALL", "SerialNumber"}:        "Computed:true",
	{"ALL", "Monitor"}:             "Optional:true",
	{"ALL", "ModelId"}:             "Optional:true",
	{"ALL", "IpAddress"}:           "Optional:true",
	{"ALL", "DuplicateOf"}:         "Computed:true",
	{"ALL", "SysTags"}:             "Computed:true",
	{"ALL", "CostCc"}:              "Optional:true",
	{"ALL", "SysTags"}:             "Computed:true",
	{"ALL", "SupportGroup"}:        "Optional:true",
	{"ALL", "OrderDate"}:           "Computed:true",
	{"ALL", "Schedule"}:            "Computed:true",
	{"ALL", "Due"}:                 "Computed:true",
	{"ALL", "Unverified"}:          "Computed:true",
	{"ALL", "CorrelationId"}:       "Computed:true",
	{"ALL", "Attributes"}:          "Computed:true",
	{"ALL", "Location"}:            "Optional:true",
	{"ALL", "Asset"}:               "Optional:true",
	{"ALL", "Category"}:            "Computed:true",
	{"ALL", "LeaseId"}:             "Optional:true",
	{"ALL", "CdRom"}:             "Optional:true,Default:false",
	{"ALL", "ALL"}:                 "Optional:true",
}

func (FlagAttrMap *FlagMap) Exists(ciName string, ciAttr string) bool {
	var key Key
	key.CiName = ciName
	key.CiAttribute = ciAttr

	if _, ok := flagAttrMap[key]; ok {
		return true
	}

	return false
}

type Flags interface {
	Get(string, string) string
	Exists(string, string) string
}

func Get(CiName string, ciAttr string) string {
	var key Key
	key.CiName = CiName
	key.CiAttribute = ciAttr
	if val, ok := flagAttrMap[key]; ok {
		return val
	}
	key.CiName = "ALL"
	if val, ok := flagAttrMap[key]; ok {
		return val
	}
	key.CiAttribute = "ALL"
	if val, ok := flagAttrMap[key]; ok {
		return val
	}
	return ""
}
