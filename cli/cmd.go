package cli

import (
	"os"
	"os/exec"

	"github.com/nomad-software/vend/output"
)

// UpdateModule makes sure the module is updated ready to vendor the
// dependencies.
func UpdateModule() {
	var commands = []string{"tidy", "download", "vendor"}

	for _, command := range commands {
		cmd := exec.Command("go", "mod", command)

		cmd.Env = buildEnv()
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		output.OnError(err, command)
	}
}

// ReadModJSON reads the module file and returns a Json string.
func ReadModJSON() string {
	cmd := exec.Command("go", "mod", "edit", "-json")

	cmd.Env = buildEnv()
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()
	output.OnError(err, "Error running 'go mod edit'")

	return string(b)
}

// ReadDownloadJSON reads dependency information and returns a Json string.
func ReadDownloadJSON() string {
	cmd := exec.Command("go", "mod", "download", "-json")

	cmd.Env = buildEnv()
	cmd.Stderr = os.Stderr

	b, err := cmd.Output()
	output.OnError(err, "Error running 'go mod download'")

	return string(b)
}

// BuildEnv creates the environment in which to run the commands.
func buildEnv() []string {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")
	return env
}
