package file

import (
	"encoding/json"

	"github.com/nomad-software/vend/output"
)

// ParseModJSON parses the mode file into a data structure.
func ParseModJSON(raw string) GoMod {
	data := GoMod{
		Module:  Module{},
		Require: make([]Require, 0, 10),
		Exclude: make([]Module, 0, 10),
		Replace: make([]Replace, 0, 10),
		Retract: make([]Retract, 0, 10),
	}

	err := json.Unmarshal([]byte(raw), &data)
	output.OnError(err, "Error parsing module json")

	return data
}

// GoMod represents parsed module json data.
type GoMod struct {
	Module  Module
	Go      string
	Require []Require
	Exclude []Module
	Replace []Replace
	Retract []Retract
}

// Module represents parsed module json data.
type Module struct {
	Path    string
	Version string
}

// Require represents parsed module json data.
type Require struct {
	Path     string
	Version  string
	Indirect bool
}

// Replace represents parsed module json data.
type Replace struct {
	Old Module
	New Module
}

// Retract represents dependency version that are retracted.
type Retract struct {
	Low       string
	High      string
	Rationale string
}
