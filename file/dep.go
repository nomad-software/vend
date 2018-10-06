package file

import (
	"encoding/json"

	"github.com/nomad-software/vend/output"
)

// ParseDepFile parses the dependency file into a data structure.
func ParseDepFile(raw string) []Dep {
	data := make([]Dep, 10)
	err := json.Unmarshal([]byte(raw), &data)
	output.OnError(err, "Error parsing json")

	return data
}

// Dep represents parsed module json data.
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
