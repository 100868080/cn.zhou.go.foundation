package main

import (
	"html/template"
	"log"
	"os"
)

func main() {
	const temp1 = `<p>A:{{.A}}</p><p>B:{{.B}}</p>`
	t := template.Mult(template.New("escape").Parse(temp1))
	var data struct {
		A string        //untrusted plain text
		B template.HTML //trusted HTML
	}
	data.A = "<b>HELLO!</b>"
	data.B = "<b>HELLO!</b>"
	if err := t.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

var issueList = template.Must(template.New("issuelist").Parse(`
	<h1>{{.TotalCount}} issues</h1>
	<table>
	<tr style='text-align:left'>
	<th>#</th>
	<th>State</th>
	<th>User</th>
	<th>Title</th>
	</tr>
	{{range.Items}}
	<tr>
	<td><a href='{{.HTMLURL}}'>{{.Number}}</a></td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.HTMLURL}}'>{{.Title}}</a></td>
	</td>
	{{end}}
	</table>
`))
