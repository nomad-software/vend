package cli

import (
	"bufio"
	"os"
	"os/exec"
	"strings"

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

// ReadModFile reads the mod file and returns a Json string.
func ReadModFile() string {
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

// ReadCacheDir reads the cache directory for dependency information.
func ReadCacheDir() string {
	cmd := exec.Command("go", "mod", "download", "-json")

	cmd.Env = buildEnv()
	cmd.Stderr = os.Stderr

	stdout, err := cmd.StdoutPipe()
	output.OnError(err, "Error connecting to 'go mod edit' stdout")

	err = cmd.Start()
	output.OnError(err, "Error starting 'go mod edit'")

	scanner := bufio.NewScanner(stdout)
	json := ""
	for scanner.Scan() {
		json += scanner.Text()
	}

	err = cmd.Wait()
	output.OnError(err, "Error while waiting for 'go mod edit'")

	// Knarly workaround because the above command doesn't return valid Json ffs!
	json = "[" + strings.Replace(json, "}{", "},{", -1) + "]"

	return json
}

// BuildEnv creates the environment in which to run the commands.
func buildEnv() []string {
	env := os.Environ()
	env = append(env, "GO111MODULE=on")
	return env
}
