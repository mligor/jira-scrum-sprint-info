package main

import (
	"fmt"

	jirascrumsprintinfo "github.com/mligor/jira-scrum-sprint-info"
)

func main() {

	jiraURL := "https://jira.mycompany.com"
	jiraToken := "MYTOKEN"

	client := jirascrumsprintinfo.NewClient(jiraURL, jiraToken)

	sprintInfo, err := client.GetSprintInfo(862) // S107
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v", sprintInfo)

}
