package file

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"github.com/nomad-software/vend/output"
)

// ParseDownloadJSON parses the dependency file into a data structure.
func ParseDownloadJSON(raw string) []Dep {
	decoder := json.NewDecoder(strings.NewReader(raw))
	data := make([]Dep, 0, 10)

	for {
		var dep Dep
		err := decoder.Decode(&dep)

		if err != nil {
			if err == io.EOF {
				break
			}
			output.OnError(err, "Error decoding dependency json")
		}

		data = append(data, dep)
	}

	return data
}

// Dep represents parsed dependency json data.
type Dep struct {
	Path     string
	Version  string
	Error    string
	Info     string
	GoMod    string
	Zip      string
	Dir      string
	Sum      string
	GoModSum string
}

// String returns a string representation.
func (d *Dep) String() string {
	return fmt.Sprintf("%s (%s)", d.Path, d.Version)
}
