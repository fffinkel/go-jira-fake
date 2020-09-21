package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

func main() {
	fmt.Println("go-jira-fake: My fake project that uses go-jira")

	tp := jira.BasicAuthTransport{
		Username: "finkel.matt@gmail.com",
		Password: os.Getenv("_FAKE_"),
	}

	jiraClient, _ := jira.NewClient(
		tp.Client(),
		"https://go-jira-finkel.atlassian.net/",
	)

	//	err := updateIssueStatus(jiraClient.Issue)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	err = createIssueWithEpicLink(jiraClient)
	//	if err != nil {
	//		panic(err)
	//	}
	//
	//	err = updateIssueWithEpicLink(jiraClient)
	//	if err != nil {
	//		panic(err)
	//	}

	err := searchIssues(jiraClient)
	if err != nil {
		panic(err)
	}
}

func searchIssues(cl *jira.Client) error {

	opts := jira.SearchOptions{MaxResults: 0}

	issues, resp, err := cl.Issue.Search(`project=FART`, &opts)

	aaaahhhh, _ := json.MarshalIndent(resp, "", "\t")
	fmt.Printf("\n\n----------> here's the bidness:\n%s\n", aaaahhhh)

	if err != nil {
		return err
	}
	fmt.Printf("NumIssues: %d, MaxResults: %d, Total: %d\n", len(issues), resp.MaxResults, resp.Total)

	return nil
}
