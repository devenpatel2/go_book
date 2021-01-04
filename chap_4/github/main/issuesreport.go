// This code demonstrates use of template strings
// to run
// go run issuesreport.go is:open json decoder


// A template string is enclosed in double braces - actions
// actions can trigger other behaviours
// actions contains expressions in template language

// packages - text/template , html/template
package main
import (
	"log"
	"os"
	"text/template"
	"time"
	"github"
)

// Current Value is refered as "." i.e the template will be processing the 
// data structure executed with template will print the corresponding
// fields of the data-structure, which in this case is "github.IssuesSearchResult"

// {{ range .Items}} and {{end}} loop the indivutal Item elements of IssuesSearchResult
// '|' operator takes output of one action and passes it to the next, like Unix's pipe operator
// i.e output of current results (IssuesSearchResults.Item.Title) is passed to `printf`
// Go's ``printf`` command is buitl-in synonym for fmt.Sprintf
// for .Age. the output is passed to daysAgo func , since CreatedAt is not a string

const templ = `{{.TotalCount}} issues:
{{range .Items}}------------------------------------
Number: {{.Number}}
User: 	{{.User.Login}}
Title: 	{{.Title | printf "%.64s"}}
Age:	{{.CreatedAt |daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}


// To produce the ouput, parse the template into a suitable representation
// Then execute the representation with suitable input
/*
report, err := template.New("report").
	Funcs(template.FuncMap{"daysAgo": daysAgo}).
	Parse(templ)

if err != nil{
	log.Fatal(err)
}
*/

// the above two steps for creating a correct representation of the templated, 
// and then verifying if the step was successful, can be done in one step
// this is done so that template is created at compile time and possible errors are caught early
var report = template.Must(template.New("issuelist").   // create template
	Funcs(template.FuncMap{"daysAgo": daysAgo}).		// augmment with func daysAgo
	Parse(templ))										// parse template

func main(){
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	if err := report.Execute(os.Stdout, result); err != nil { //execute template with result as input
		log.Fatal(err)
	}
}

/* Partial+ output
0 issues:
------------------------------------
Number: 3
User: 	havardl
Title: 	json.decoder.JSONDecodeError
Age:	17 days
------------------------------------
Number: 24231
User: 	mshanak
Title: 	Desk translation page bug ( json.decoder)
Age:	6 days
------------------------------------
*/