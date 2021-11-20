package task

import (
	"fmt"
	"io/ioutil"
	"os"
	"text/template"

	"github.com/aldorperez1/workutils/utils"
)

type Task struct {
	Name string
	CWD  string
}

func (r *Task) CreateTask() error {
	prefix, err := utils.GeneratePrefix(".")
	if err != nil {
		return err
	}
	r.Name = fmt.Sprintf("%s-%s", prefix, r.Name)
	if err := os.Mkdir(r.Name, 0755); err != nil {
		return err
	}

	// README
	if err = os.Chdir(r.Name); err != nil {
		return err
	}
	r.CWD, err = os.Getwd()
	const readmeTmpl = `
# {{.Name}}

---

[TOC]

---

## About


**Background**
<!-- Background as to why this project exist. e.g. I was told to do this because blah -->

**Request**
<!-- What was asked of you to do -->

**Project Status**: :red_circle:

**Work Directory**: {{.CWD}}


## Tasks
<!--Plan out your task here e.g. -[ ] Take over the world ;)  -->


---


## Worklogs

`
	tmp, err := template.New("Readme").Parse(readmeTmpl)
	if err != nil {
		fmt.Println("Unable to parse template", err)
		return err
	}
	// Create
	f, err := os.OpenFile("README.md", os.O_CREATE|os.O_WRONLY, 0644)
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

	// Makefile
	const makefileTmpl = `
DOCFILES := $(shell find . -name '*.md' -type f)
DOCNAME = docs.md 
.PHONY: docs clean-docs get-all-tasks effort
docs:
	cat $(DOCFILES) > $(DOCNAME)
clean-docs:
	rm $(DOCNAME)
get-all-tasks:
	find . -name '*worklog.md' -type f  -exec grep -e '- \[ \]' {} \; | sort | uniq | sed  's/- \[ \] //'
`

	if err := ioutil.WriteFile("Makefile", []byte(makefileTmpl), 0644); err != nil {
		fmt.Printf("Unable to write file: %v", err)
		return err
	}

	//err = ioutil.WriteFile("filename.txt", []byte("Hello"), 0755)
	//if err != nil {
	//	fmt.Printf("Unable to write file: %v", err)
	//	return err
	//}
	return nil
}
