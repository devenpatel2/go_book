// implement SearchIssue function
package github

import (
    "encoding/json"
    "fmt"
    "net/http"
    "net/url"
    "strings"
)

// SearchIssues  queries the Github issue tracker
// Makes an HTTP request and decodes the result as JSON
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	// handle escape chars like ?, & 
    q := url.QueryEscape(strings.Join(terms, " "))
    resp, err := http.Get(IssuesURL + "?q=" + q)
    if err != nil {
        return nil, err
    }

    // We must close resp.Body on all execution paths
    if resp.StatusCode != http.StatusOK {
        resp.Body.Close()
        return nil, fmt.Errorf("search query failed :%s", resp.Status)
    }

    var result IssuesSearchResult
	// json.NewDecoder returns a streaming decoder
	// Decode populates the result
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
        resp.Body.Close()
        return nil, err
    }
    resp.Body.Close()
    return &result, nil
}
