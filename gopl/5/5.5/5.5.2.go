package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const IssuesURL = "https://api.github.com/search/issues"

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
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

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

// SearchIssues queries the GitHub issue tracker.
func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}

	// We must close resp.Body on all execution paths.
	// (Chapter 5 presents 'defer', which makes this simpler.)
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	// fmt.Println(result.Items)
	resp.Body.Close()
	return &result, nil
}

func main() {
	result, err := SearchIssues([]string{"junit-team/junit5/issues"})
	if err != nil {
		log.Fatal(err)
	}

	issueMap := make(map[string][]*Issue)
	for _, item := range result.Items {
		d := time.Since(item.CreatedAt)

		if d.Hours() < 30*24 {
			issueMap["lessMonth"] = append(issueMap["lessMonth"], item)
		} else if d.Hours() < 365*30*24 {
			issueMap["lessYear"] = append(issueMap["lessYear"], item)
		} else {
			issueMap["moreYear"] = append(issueMap["moreYear"], item)
		}

		//練習 4.10： 脩改issues程序，根據問題的時間進行分類，比如不到一個月的、不到一年的、超過一年。
		// fmt.Printf("#%-5d %9.9s %.55s %v\n",
		// 	item.Number, item.User.Login, item.Title, d.Hours())
		// fmt.Printf("#%-5d %9.9s %.55s \n",
		// 	item.Number, item.User.Login, item.Title)
	}

	fmt.Println(issueMap)
}
