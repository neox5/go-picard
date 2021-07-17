package fileutil

import (
	"html/template"
	"os"
	"path/filepath"
)

// FileTemplate is the base from which a file gets created.
type FileTemplate struct {
	Name string      // filename or path/filename
	Tmpl string      // go template in string format
	Data interface{} // data used for template generation
}

// Create creates a file from a FileTemplate
func Create(file FileTemplate) error {
	t, err := template.New("file").Parse(file.Tmpl)
	if err != nil {
		return err
	}

	if !pathExists(filepath.Dir(file.Name)) {
		if err = os.MkdirAll(filepath.Dir(file.Name), 0777); err != nil {
			return err
		}
	}

	f, err := os.Create(file.Name)
	if err != nil {
		return err
	}
	defer f.Close()

	return t.Execute(f, file.Data)
}

// CreateMultiple creates multiple files from a FileTemplate array
func CreateMultiple(files []FileTemplate) error {
	for _, f := range files {
		if err := Create(f); err != nil {
			return err
		}
	}
	return nil
}

// pathExists checks if a path exists
func pathExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
