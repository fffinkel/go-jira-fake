package main

import (
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

func main() {
	fmt.Println("go-jira-fake: My fake project that uses go-jira")

	testIssueID := "FART-1"

	tp := jira.BasicAuthTransport{
		Username: "finkel.matt@gmail.com",
		Password: os.Getenv("_FAKE_"),
	}

	jiraClient, _ := jira.NewClient(tp.Client(), "https://go-jira-finkel.atlassian.net/")

	issueService := jiraClient.Issue
	issue, _, _ := issueService.Get(testIssueID, nil)

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	fmt.Printf("Status: %s\n", issue.Fields.Status.Name)

	possibleTransitions, _, _ := issueService.GetTransitions("FART-1")
	currentStatus := issue.Fields.Status.Name
	var transitionID string
	for _, v := range possibleTransitions {
		if v.Name != currentStatus {
			transitionID = v.ID
			break
		}
	}

	issueService.DoTransition("FART-1", transitionID)
	issue, _, _ = issueService.Get(testIssueID, nil)

	fmt.Printf("Status after transition: %+v\n", issue.Fields.Status.Name)
}
