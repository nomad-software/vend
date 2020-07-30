package main

import (
	"github.com/nomad-software/vend/cli"
	"github.com/nomad-software/vend/file"
)

func main() {

	options := cli.ParseOptions()

	if options.Help {
		options.PrintUsage()

	} else {
		cli.UpdateModule()

		deps := file.ParseDownloadJSON(cli.ReadDownloadJSON())
		mod := file.ParseModJSON(cli.ReadModJSON())

		file.CopyModuleDependencies(mod, deps)
	}
}
