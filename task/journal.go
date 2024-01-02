package task

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const journalTmpl = `
# 
**{{.Date}}**

<!-- Write down what you are doing -->

TODO: 

- [ ]

## ⚙️ 
`

type Journal struct {
	Date string
}

func CreateJournal() error {
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	r := Journal{date}
	tmp, err := template.New("journal").Parse(journalTmpl)
	if err != nil {
		fmt.Println("Unable to parse template", err)
		return err
	}
	// Create
	f, err := os.OpenFile(date+".md", os.O_CREATE|os.O_WRONLY, 0644)
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
