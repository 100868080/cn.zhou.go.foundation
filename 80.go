const temp1='{{.TotalCount}} issues:
{{range .Items}}-----------------------------
Number:{{.Number}}
User:{{.User.Login}}
Title:{{.Title|printf "%.64s"}}
Age:{{.CreatedAt|dayAgo}}days
{{end}}'


func dayAgo(t time.Time)int{
	return int(time.Since(t).Hours()/24)
	  
}

report,err:=template.New("report").
Funcs(template.FuncMap{"daysAgo":dayAgo}).
Parse(temp1)
if err!=nil{
	log.Fatal(err)
}

