package main

import (
	"bytes"
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
	//updateIssueStatus(jiraClient.Issue)
	err := createIssueWithEpicLink(jiraClient)
	if err != nil {
		panic(err)
	}
	err = updateIssueWithEpicLink(jiraClient)
	if err != nil {
		panic(err)
	}
}

func createIssueWithEpicLink(cl *jira.Client) error {
	testIssueID := "FART-4"
	issueService := cl.Issue
	issue, _, err := issueService.Get(testIssueID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	fmt.Printf("Status: %s\n", issue.Fields.Status.Name)
	fmt.Printf("Unknowns: %+v\n", issue.Fields.Unknowns)

	aaaahhhh, _ := json.MarshalIndent(issue.Fields.Unknowns, "", "\t")
	fmt.Printf("\n\n----------> here's the bidness:\n%s\n", aaaahhhh)

	userService := cl.User
	me, _, _ := userService.GetSelf()

	fieldList, _, _ := cl.Field.GetList()

	var customFieldID string
	for _, v := range fieldList {
		if v.Name == "Epic Link" {
			customFieldID = v.ID
			break
		}
	}

	var unknowns map[string]interface{}
	unknowns = map[string]interface{}{
		customFieldID: "FART-2",
	}

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Description: "Test Issue",
			Summary:     "Beep boop.",
			Assignee:    me,
			Reporter:    me,
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "FART",
			},
			Unknowns: unknowns,
		},
	}

	newIssue, _, _ := issueService.Create(&i)
	aaaahhhh, _ = json.MarshalIndent(newIssue, "", "\t")
	fmt.Printf("\n\n----------> here's the bidness:\n%s\n", aaaahhhh)

	return nil
}

func updateIssueWithEpicLink(cl *jira.Client) error {
	testIssueID := "FART-4"
	issueService := cl.Issue
	issue, _, err := issueService.Get(testIssueID, nil)
	if err != nil {
		return err
	}

	fmt.Printf("%s: %+v\n", issue.Key, issue.Fields.Summary)
	fmt.Printf("Type: %s\n", issue.Fields.Type.Name)
	fmt.Printf("Priority: %s\n", issue.Fields.Priority.Name)
	fmt.Printf("Status: %s\n", issue.Fields.Status.Name)
	fmt.Printf("Unknowns: %+v\n", issue.Fields.Unknowns)

	aaaahhhh, _ := json.MarshalIndent(issue.Fields.Unknowns, "", "\t")
	fmt.Printf("\n\n----------> here's the bidness:\n%s\n", aaaahhhh)

	userService := cl.User
	me, _, _ := userService.GetSelf()

	i := jira.Issue{
		Fields: &jira.IssueFields{
			Description: "Test Issue",
			Summary:     "Beep boop.",
			Assignee:    me,
			Reporter:    me,
			Type: jira.IssueType{
				Name: "Bug",
			},
			Project: jira.Project{
				Key: "FART",
			},
		},
	}

	newIssue, _, _ := issueService.Create(&i)
	fieldList, _, _ := cl.Field.GetList()

	var customFieldID string
	for _, v := range fieldList {
		if v.Name == "Epic Link" {
			customFieldID = v.ID
			break
		}
	}

	is := cl.Issue
	var update map[string]interface{}
	update = map[string]interface{}{
		"fields": map[string]string{
			customFieldID: "FART-2",
		},
	}
	response, err := is.UpdateIssue(newIssue.ID, update)

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	aaaahhhh, _ = json.MarshalIndent(buf.String(), "", "\t")
	fmt.Printf("\n\n----------> here's the bidness:\n%s\n", aaaahhhh)

	return nil
}
