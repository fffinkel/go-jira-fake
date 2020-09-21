package main

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func searchIssues(cl *jira.Client) error {

	opts := jira.SearchOptions{
		MaxResults:   7,
		MetadataOnly: true,
	}

	issues, resp, err := cl.Issue.Search(`project=FART`, &opts)
	if err != nil {
		return err
	}
	fmt.Printf("NumIssues: %d, MaxResults: %d, Total: %d\n", len(issues), resp.MaxResults, resp.Total)

	return nil
}
