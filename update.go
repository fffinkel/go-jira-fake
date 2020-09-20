package main

import (
	"fmt"

	"github.com/andygrunwald/go-jira"
)

func updateIssueStatus(svc *jira.IssueService) error {
	testIssueID := "FART-1"

	issue, _, err := svc.Get(testIssueID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	fmt.Printf("Status: %s\n", issue.Fields.Status.Name)

	allowedTransitions, _, err := svc.GetTransitions("FART-1")
	if err != nil {
		return err
	}

	currentStatus := issue.Fields.Status.Name
	var transitionID string
	for _, v := range allowedTransitions {
		if v.Name != currentStatus {
			transitionID = v.ID
			break
		}
	}

	svc.DoTransition("FART-1", transitionID)
	issue, _, err = svc.Get(testIssueID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("Status after transition: %+v\n", issue.Fields.Status.Name)
	return nil
}
