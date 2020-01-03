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
	var options cli.Options
	err := cli.ProcessFlags(&options)
	if err != nil {
		fmt.Printf("Exiting with fail code : %s", err)
	}

	// Set up the connection to ServiceNow
	var snowClient *client.Client
	if env, err := cli.GetEnvVars(); err != nil {
		fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s", err)
		return
	} else {
		snowClient = client.NewClient(env.Url, env.Userid, env.Password)
	}

	// Create a slice to store the records we retrieve from ServiceNow
	var CiList []client.CmdbCIMetaModel
	//CiList, count, err := snowClient.ReadCIs(client.BaseClass, 0, CiList)

	CiList, count, err := client.ReadCIs("cmdb_ci_server", 0, CiList, snowClient, &options)
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
