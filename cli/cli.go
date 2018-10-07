package cli

import (
	"flag"
	"fmt"

	"github.com/fatih/color"
)

// Options contains CLI arguments passed to the program.
type Options struct {
	Help bool
}

// ParseOptions parses the command line options and returns a struct filled with
// the relevant options.
func ParseOptions() Options {
	var opt Options

	flag.BoolVar(&opt.Help, "help", false, "Show help.")
	flag.Parse()

	return opt
}

// PrintUsage prints the usage of this tool.
func (opt *Options) PrintUsage() {
	const banner string = `                     _
__   _____ _ __   __| |
\ \ / / _ \ '_ \ / _' |
 \ V /  __/ | | | (_| |
  \_/ \___|_| |_|\__,_|

`

	color.Cyan(banner)
	fmt.Printf("Properly vendor all dependencies.\n\n")

	flag.Usage()
}
