package main

import (
	"fmt"

	jirascrumsprintinfo "github.com/mligor/jira-scrum-sprint-info"
)

func main() {

	jiraURL := "https://jira.mycompany.com"
	jiraToken := "MYTOKEN"

	client := jirascrumsprintinfo.NewClient(jiraURL, jiraToken)

	sprints := []int{
		728, // S100
		746, // S101
		784, // S102
		798, // S103
		812, // S104
		824, // S105
		839, // S106
		862, // S107
		865, // S108
		881, // S109
		897, // S110
		911, // S111
	}

	for _, s := range sprints {
		sprintInfo, err := client.GetSprintInfo(s) // S107
		if err != nil {
			panic(err)
		}

		fmt.Println(sprintInfo)
	}

}
