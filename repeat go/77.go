package github
import "time"
const IssuesURL="https://api.github.com/search/issues"
type IssuesSear chResult struct{
	TotalCount int 'json:"total_count"'
	Items []*Issue 
}
type Issue struct{
	Number int
	HTMLURL string 'json:"html_url"'
	Title string
	State string
	User *User
	CreatedAt time.Time 'json:"created_at"'
	Body string  //in Markdown format
}
type User struct{
	Login string 
	HTML string 'json:"html_url'
}