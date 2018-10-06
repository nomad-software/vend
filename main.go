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
		mod := file.ParseModFile(json)

		json = cli.ReadCacheDir()
		deps := file.ParseDepFile(json)

		file.CopyDependencies(mod, deps)
	}
}
