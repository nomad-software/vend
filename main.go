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

		json := cli.ReadModFile()
		mod := file.ParseModJSON(json)

		json = cli.ReadDownloadCache()
		deps := file.ParseDownloadJSON(json)

		file.CopyDependencies(mod, deps)
	}
}
