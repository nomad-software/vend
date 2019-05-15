package file

import (
	"encoding/json"

	"github.com/nomad-software/vend/output"
)

// ParseModJSON parses the mode file into a data structure.
func ParseModJSON(raw string) GoMod {
	data := GoMod{
		Module:  module{},
		Require: make([]require, 0, 10),
		Exclude: make([]module, 0, 10),
		Replace: make([]replace, 0, 10),
	}

	err := json.Unmarshal([]byte(raw), &data)
	output.OnError(err, "Error parsing module json")

	return data
}

// GoMod represents parsed module json data.
type GoMod struct {
	Module  module
	Require []require
	Exclude []module
	Replace []replace
}

// module represents parsed module json data.
type module struct {
	Path    string
	Version string
}

// require represents parsed module json data.
type require struct {
	Path     string
	Version  string
	Indirect bool
}

// replace represents parsed module json data.
type replace struct {
	Old module
	New module
}
