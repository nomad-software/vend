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
		json := cli.ReadDownloadJSON()
		deps := file.ParseDownloadJSON(json)
		json = cli.ReadModJSON()
		mod := file.ParseModJSON(json)

		if options.PkgOnly {
			file.CopyPkgDependencies(mod, deps)
		} else {
			file.CopyModuleDependencies(mod, deps)
		}
	}
}
