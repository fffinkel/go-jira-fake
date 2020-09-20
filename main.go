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

	jiraClient, _ := jira.NewClient(tp.Client(), "https://go-jira-finkel.atlassian.net/")
	issue, _, _ := jiraClient.Issue.Get("FART-1", nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
}
