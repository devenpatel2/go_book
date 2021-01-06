// demonstration of using json interface with http requests
// See https://developer.github.com/v3/search/#search-issues.

package github

import "time"

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
    TotalCount  int `json:"total_count`  // field tag to convert field TotalCount to total_count for json
    Items       []*Issue    
}

// struct names have to be capitalized if they are to be imported along with the package
type Issue struct {
    Number      int
    HTMLURL     string `json:"html_url"`
    Title       string
    State       string
    User        *User
    CreatedAt   time.Time   `json:"created_at"`
    Body        string  // in Markdown format
}

type User struct {
    Login   string
    HTMLURL string `json:"html_url"`
}
