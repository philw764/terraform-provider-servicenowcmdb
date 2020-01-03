package cli

import (
	"fmt"
	"github.com/jessevdk/go-flags"
)

//type Cli struct {
//	ClassList  *string
//	DryRun     *bool
//	Rebuild    *bool
//	ConfigFile *string
//	Silent	*bool
//	Debug 	*bool
//}
//
//func ProcessFlags(CliArgs *Cli) error {
//	//var CliArgs Cli
//	CliArgs.DryRun = flag.Bool("dryrun", false, "Show what would change without making changes")
//	CliArgs.Rebuild = flag.Bool("rebuild", false, "Delete and rebuild all classes")
//	CliArgs.Silent	= flag.Bool("silent", false, "Run silently")
//	CliArgs.Debug	= flag.Bool( "debug", false, "Enable debug mode")
//	CliArgs.ClassList = flag.String("classlist", "", "list of classes to include in provider from Servicenow")
//	CliArgs.ConfigFile = flag.String("configuration file", "", "use a config file to store build config")
//	return nil
//}

type Options struct {
	ClassList map[string]string `short:"c" long:"classlist" description:"this is the list of classes to include in the provider" default:"cmdb_ci_server:recurse"`
}

func ProcessFlags(options *Options) error {
	var parser = flags.NewParser(options, flags.Default)
	test, err := parser.Parse()
	if err != nil {
		fmt.Printf("Invalid command line passed:%s .. %s", test, err)
		return err
	}
	return nil
}
