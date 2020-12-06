package github

import "time"

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}

const IssuesURL = "https://api.github.com/search/issues"
