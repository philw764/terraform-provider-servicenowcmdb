package main

import (
	"fmt"
	"os"
	"terraform-provider-servicenowcmdb/servicenowcmdb/cli"
	"terraform-provider-servicenowcmdb/servicenowcmdb/generateprovidersource/client"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////
// GenerateProviderSource is a utility provided with the ServiceNow Terraform Provider to enable users to
// create a custom build of the provider using the Configuration Item classes they require.
//
// Basically this utility writes a bunch of GO sources files using metadata from ServiceNow and the GO
// template engine to generate the source code in GO required to build the Terraform to ServiceNow provider.
//
// This utility is required because ServiceNow provides the ability to add and change Configuration Items
// in the CMDB.  If we were to build a the Provider based only on the out of the box Service Now
// Configuration Items and their attributes then users of this provider would not be able to manage
// custom Configuration Item classes and their attributes.
//
// This utility works by connecting to the ServiceNow server using the ServiceNow provide API to retrieve the
// metadata for each CMDB CI Class.  This metadata describes the CI Class and its attributes, this meta
// data is required to enable the use of the API to write to the CMDB.
//
// Once the CI Class metadata is retrieved, each CI and its metadata is stored in a list.  This list is then
// used as input to a template file to produce the GO code required to support writing to the CMDB for each
// CI Class.  The "cmdb_ci_template.tmpl" file is the template file used to create the CMDB CI resource
// required by the Terraform provider to operate.  Any changes to the code for each CI resource file must
// be performed in the template file, any changes made to a resource file will be overwritten the next time
// this utility is executed.
//
// Once all of the resource files are produced (one for each CI Class) then the Terraform provider file is
// generated from a template.  The template is called "provider.tmpl", this template generator uses the
// same list of CI Class metadata as the resource template generator.
//
// Once this utility has been executed, the terraform-provider-servicenowcmdb provider must be re-compiled
// to do this, execute "go build" in the provider base directory, then execute "go install" to place
// the newly created provider binary file into its correct location.
//
// Once this provider matures, I'll create a default build of the provider which will work using the
// most commonly used CMDB CI Classes on an out of the box ServiceNow instance to enable users to get up
// and running quickly.
//////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	// Retrieve and process command line options passed in.  Refer to the documention for details on
	// the command line options.
	var options cli.Options
	err := cli.ProcessOptions(&options)
	if err != nil {
		fmt.Printf("failed to process command line options - exiting with failure code : %s\n", err)
		os.Exit(1)
	}

	// Set up the connection to ServiceNow
	var snowClient *client.Client
	if env, err := cli.GetEnvVars(); err != nil {
		fmt.Printf("Environment Variables not set cannot connect to ServiceNow:%s\n", err)
		os.Exit(1)
	} else {
		snowClient = client.NewClient(env.Url, env.Userid, env.Password)
	}

	// Create a list to store the metadata for each CMDB CI Class retrieved from ServiceNow.
	var CiList []client.CmdbCIMetaModel

	CiList, count, err := client.GetListOfClassesFromServiceNow("cmdb_ci_server", 0, CiList, snowClient, &options)
	if err != nil {
		fmt.Println("Failed to get data. Is ServiceNow running?  If it is, check credentials as correct")
	} else {
		// Now that the list of CI Classes have been retrieved, call WriteCIResourcesToFile to create a resource
		// file for each CI Class.  Each resource file is created using the CI Class name.
		for ci := range CiList {
			_ = snowClient.WriteCIResourcesToFile(CiList[ci])
		}
		// Call WriteProviderToFile is called to generate the "provider.go" file
		_ = snowClient.WriteProviderToFile(CiList)
		fmt.Printf("the number of records processed is: %v", count)
	}
}
