package cli

import (
	"bufio"
	"os"
	"os/exec"

	"github.com/nomad-software/vend/output"
)

// UpdateModule makes sure the module is updated ready to vendor the
// dependencies.
func UpdateModule() {
	var commands = []string{"tidy", "download"}

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

	stdout, err := cmd.StdoutPipe()
	output.OnError(err, "Can't connect to 'go mod edit' stdout")

	err = cmd.Start()
	output.OnError(err, "Error starting 'go mod edit'")

	scanner := bufio.NewScanner(stdout)
	json := ""
	for scanner.Scan() {
		json += scanner.Text()
	}

	err = cmd.Wait()
	output.OnError(err, "Error waiting for 'go mod edit'")

	return json
}

// ReadDownloadJSON reads dependency information and returns a Json string.
func ReadDownloadJSON() string {
	cmd := exec.Command("go", "mod", "download", "-json")

	cmd.Env = buildEnv()
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	output.OnError(err, "Error connecting to 'go mod download' stdout")

	err = cmd.Start()
	output.OnError(err, "Error starting 'go mod download'")

	scanner := bufio.NewScanner(stdout)
	json := ""
	for scanner.Scan() {
		json += scanner.Text()
	}

	err = cmd.Wait()
	output.OnError(err, "Error while waiting for 'go download edit'")

	return json
}

// BuildEnv creates the environment in which to run the commands.
func buildEnv() []string {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")
	return env
}
