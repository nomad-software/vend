package file

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/nomad-software/vend/cli"
	"github.com/nomad-software/vend/output"
)

// VendorDir represents a vendor directory.
type VendorDir struct {
	basePath       string
	modFileContent []byte
	mod            GoMod
	deps           []Dep
}

// InitVendorDir creates a new vendor directory.
func InitVendorDir() VendorDir {
	wd, err := os.Getwd()
	output.OnError(err, "Error getting the current directory")

	v := VendorDir{
		basePath: path.Join(wd, "vendor"),
		mod:      ParseModJSON(cli.ReadModJSON()),
		deps:     ParseDownloadJSON(cli.ReadDownloadJSON()),
	}

	if !v.exists(v.basePath) {
		output.Error("No dependencies vendored")
	}

	return v
}

// CopyDependencies copies remote module level dependencies transitively.
func (v *VendorDir) CopyDependencies() {
	v.clear()

	for _, d := range v.deps {
		fmt.Fprintf(output.Stdout, "vend: copying %s (%s)\n", d.Path, d.Version)
		v.copy(d.Dir, v.vendPath(d.Path))
	}

	for _, r := range v.mod.Replace {
		if r.Old.Path != r.New.Path {
			fmt.Fprintf(output.Stdout, "vend: replacing %s with %s\n", r.Old.Path, r.New.Path)
			newPath := v.vendPath(r.New.Path)
			oldPath := v.vendPath(r.Old.Path)
			// If the directory is in the vendor folder it was copied from the
			// module cache so we can just rename it. Otherwise it's a local
			// directory located somewhere else that needs copying in.
			if v.exists(newPath) {
				v.copy(newPath, oldPath)
				v.remove(newPath)
			} else {
				v.copy(r.New.Path, oldPath)
			}
		}
	}
}

// exists checks if a file exists.
func (v *VendorDir) exists(file string) bool {
	_, err := os.Stat(file)
	return !os.IsNotExist(err)
}

// remove removes a path.
func (v *VendorDir) remove(p string) {
	err := os.RemoveAll(p)
	output.OnError(err, "Error removing path")
}

// vendPath creates a vendor directory path.
func (v *VendorDir) vendPath(p string) string {
	return path.Join(v.basePath, p)
}

// copyModFile internally copies and saves the modules.txt file.
func (v *VendorDir) copyModFile() {
	var err error
	v.modFileContent, err = ioutil.ReadFile(v.vendPath("modules.txt"))
	output.OnError(err, "Error reading modules.txt")
}

// writeModFile writes the modules.txt file into the vendor directory.
func (v *VendorDir) writeModFile() {
	err := ioutil.WriteFile(v.vendPath("modules.txt"), v.modFileContent, 0644)
	output.OnError(err, "Error saving modules.txt")
}

// clear removes all dependencies from the vendor directory.
func (v *VendorDir) clear() {
	v.copyModFile()
	v.remove(v.basePath)

	err := os.MkdirAll(v.basePath, 0755)
	output.OnError(err, "Error creating vendor directory")

	v.writeModFile()
}

// copy will copy files and directories.
func (v *VendorDir) copy(src string, dest string) {
	info, err := os.Lstat(src)
	output.OnError(err, "Error getting information about source")

	if info.Mode()&os.ModeSymlink != 0 {
		return // Completely ignore symlinks.
	}

	if info.IsDir() {
		v.copyDirectory(src, dest)
	} else {
		v.copyFile(src, dest)
	}
}

// copyDirectory will copy directories.
func (v *VendorDir) copyDirectory(src string, dest string) {
	err := os.MkdirAll(dest, 0755)
	output.OnError(err, "Error creating directories")

	contents, err := ioutil.ReadDir(src)
	output.OnError(err, "Error reading source directory")

	for _, content := range contents {
		s := filepath.Join(src, content.Name())
		d := filepath.Join(dest, content.Name())
		v.copy(s, d)
	}
}

// copyFile will copy files.
func (v *VendorDir) copyFile(src string, dest string) {
	err := os.MkdirAll(filepath.Dir(dest), 0755)
	output.OnError(err, "Error creating directories")

	d, err := os.Create(dest)
	output.OnError(err, "Error creating file")
	defer d.Close()

	s, err := os.Open(src)
	output.OnError(err, "Error opening file")
	defer s.Close()

	_, err = io.Copy(d, s)
	output.OnError(err, "Error copying file")
}
