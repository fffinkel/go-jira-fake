package main

import (
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
