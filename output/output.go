package output

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/mattn/go-colorable"
)

var (
	stdout = colorable.NewColorableStdout()
	stderr = colorable.NewColorableStderr()
)

// OnError prints an error if err is not nil and exits the program.
func OnError(err error, text string) {
	if err != nil {
		Error("%s:, %s", text, err.Error())
		os.Exit(1)
	}
}

// Error prints an error and exits the program.
func Error(format string, args ...interface{}) {
	fmt.Fprintf(stderr, color.RedString(format)+"\n", args...)
	os.Exit(1)
}

// Info prints information.
func Info(format string, args ...interface{}) {
	fmt.Fprintf(stdout, color.GreenString(format)+"\n", args...)
}
