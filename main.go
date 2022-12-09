package main

import (
	"flag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"terraform-provider-servicenowcmdb/servicenowcmdb/resources"
)

//	func main() {
//		plugin.Serve(&plugin.ServeOpts{
//			ProviderFunc: func() terraform.ResourceProvider {
//				return resources.Provider()
//			},
//		})
//	}
var (
	// these will be set by the goreleaser configuration
	// to appropriate values for the compiled binary
	version string = "dev"

	// goreleaser can also pass the specific commit if you want
	// commit  string = ""
)

func main() {
	var debugMode bool

	flag.BoolVar(&debugMode, "debug", false, "set to true to run the provider with support for debuggers like delve")
	flag.Parse()

	plugin.Serve(&plugin.ServeOpts{ProviderFunc: resources.Provider})
}
