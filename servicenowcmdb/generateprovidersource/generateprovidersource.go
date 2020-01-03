package main

import (
	"fmt"
	"terraform-provider-servicenowcmdb/servicenowcmdb/cli"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/client"
)

// TODO:  Need to create command line arguments:
//		1. diff ci definitions
//		2. Read ServiceNow url, user and password
//		3. Spit out error handling messages
//		4. Provide ability to check classes
//		5. Provide ability to update named classes

func main() {
	if err := cli.ProcessFlags(); err != nil {
		fmt.Println("Failed to process command line parameters.")
		return
	}

	var snowClient *client.Client
	if env, err := cli.GetEnvVars(); err != nil {
		fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s", err)
		return
	} else {
		fmt.Printf("This is the userid:%s\n", env.Userid)
		fmt.Printf("This is the pwd:%s\n", env.Password)
		fmt.Printf("This is the url:%s\n", env.Url)
		snowClient = client.NewClient(env.Url, env.Userid, env.Password)
	}

	// Slice of records to store the Meta Data for each CI in the ServiceNow CMDB
	var CiList []client.CmdbCIMetaModel
	//CiList, count, err := snowClient.ReadCIs(client.BaseClass, 0, CiList)

	CiList, count, err := client.ReadCIs("cmdb_ci_server", 0, CiList, snowClient)
	if err != nil {
		fmt.Println("Failed to get data. Is ServiceNow running?  If it is, check credentials as correct")
	} else {
		//_ = flags.GenerateControlFlagFile(CiList)
		for ci := range CiList {
			_ = snowClient.WriteCIStructToFile(CiList[ci])
		}
		_ = snowClient.WriteProviderToFile(CiList)
		fmt.Printf("the number of records processed is: %v", count)
	}
}
