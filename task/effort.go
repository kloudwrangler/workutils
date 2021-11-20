package task

import (
	"errors"
	"fmt"
	"os"
	"text/template"

	"github.com/aldorperez1/workutils/utils"
)

type Effort struct {
	Name   string
	CWD    string
	Prefix string
}

func (r *Effort) CreateEffort() error {
	// Create effort directory

	fmt.Println("Creating a project called: ", r.Name)

	// TODO: validate directory name

	// Make sure that home directory exists
	effortDir := "./Efforts"
	_, err := os.Stat(effortDir)
	if err != nil {

		if errors.Is(err, os.ErrNotExist) {
			// path/to/whatever does *not* exist
			fmt.Println("An Efforts folder does not exist. Creating at ", effortDir)
			if err := os.Mkdir(effortDir, 0755); err != nil {
				fmt.Println("Something Happened: directory unable to be created")
				return err
			}
		} else {
			// file may or may not exist. See err for details
			fmt.Println("Something Happened: file may or may not exist")
			return err
		}
	}
	fmt.Println("Efforts exists at: ", effortDir)
	// Get the name of the project based on the contents of the home directory

	// Create Effort name
	prefix, err := utils.GeneratePrefix(effortDir)
	r.Prefix = prefix
	if err != nil {
		return err
	}
	r.Name = fmt.Sprintf("%s-%s", prefix, r.Name)
	if err := os.Chdir(effortDir); err != nil {
		return err
	}
	if err := os.Mkdir(r.Name, 0755); err != nil {
		return err
	}
	fmt.Println(r.Name)

	// README
	if err = os.Chdir(r.Name); err != nil {
		return err
	}
	r.CWD, err = os.Getwd()
	const readmeTmpl = `
### {{.Name}}
<!-- Write down what you are doing -->
`
	tmp, err := template.New("worklog").Parse(readmeTmpl)
	if err != nil {
		fmt.Println("Unable to parse template", err)
		return err
	}
	// Create
	f, err := os.OpenFile(r.Prefix+"-worklog.md", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Couldn't create the file")
	}
	// Write readme
	if err := tmp.Execute(f, r); err != nil {
		fmt.Println("Unable to execute template")
		return err
	}
	if err := f.Close(); err != nil {
		fmt.Println("Unable to close file", err)
		return err
	}

	return nil
}
