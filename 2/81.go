package main

import (
	"html/template"
	"log"
	"os"
)

/*
var report = template.Must(template.New("issuelist").
	funcs(template.FuncMap{"daysAgo": dayAgo}).
	Parse(temp1))
func daysAgo(t time.Time) int {
    return int(time.Since(t).Hours() / 24)
}
report, err := template.New("report").
    Funcs(template.FuncMap{"daysAgo": daysAgo}).
    Parse(templ)
if err != nil {
    log.Fatal(err)
}
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	if err := report.Execute(os.Stdout, result); err != nil {
		log.Fatal(err)
	}
}
