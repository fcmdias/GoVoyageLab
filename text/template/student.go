package playtemplate

import (
	"os"
	"text/template"
)

type Student struct {
	ID     string
	Name   string
	Class  string
	Course string
}

func JoeSmith() *Student {
	return &Student{
		ID:    "1234",
		Name:  "Joe Smith",
		Class: "Physics",
	}
}

func (s Student) String() {
	tmpl, err := template.New("test").Parse("{{.Name}} is in the {{.Class}} course.\n")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, s)
	if err != nil {
		panic(err)
	}
}
