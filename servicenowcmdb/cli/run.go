package cli

import "flag"

type Cli struct {
	classList  []string
	dryRun     *bool
	Rebuild    *bool
	ConfigFile string
}

func ProcessFlags() error {
	var CliArgs Cli
	CliArgs.dryRun = flag.Bool("dryrun", false, "Show what would change without making changes")
	CliArgs.Rebuild = flag.Bool("rebuild", false, "Delete and rebuild all classes")
	return nil
}
