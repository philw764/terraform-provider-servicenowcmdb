package main

import (
	"fmt"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/client"
)

// TODO:  Need to create command line arguments:
//		1. diff ci definitions
//		2. Read ServiceNow url, user and password
//		3. Spit out error handling messages
//		4. Provide ability to check classes
//		5. Provide ability to update named classes

func main() {

	snowClient := client.NewClient(client.BaseUrl, client.Userid, client.Password)

	// Slice of records to store the Meta Data for each CI in the ServiceNow CMDB
	var CiList []client.CmdbCIMetaModel
	//CiList, count, err := snowClient.ReadCIs(client.BaseClass, 0, CiList)

	CiList, count, err := snowClient.ReadCIs("cmdb_ci_server", 0, CiList)
	if err != nil {
		fmt.Println("Failed to get data. Is ServiceNow running?  If it is, check credentials as correct")
	}
	_ = client.GenerateControlFlagFile(CiList)
	for ci := range CiList {
		snowClient.WriteCIStructToFile(CiList[ci])
	}
	_ = snowClient.WriteProviderToFile(CiList)
	fmt.Printf("the number of records processed is: %v", count)

}
