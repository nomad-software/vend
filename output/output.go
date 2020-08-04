package output

import (
	"fmt"
	"os"
)

// OnError prints an error if err is not nil and exits the program.
func OnError(err error, text string) {
	if err != nil {
		Fatal("%s:, %s", text, err.Error())
	}
}

// Fatal prints an error and exits the program.
func Fatal(format string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

// Info prints information.
func Info(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
