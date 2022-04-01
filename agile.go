package jirascrumsprintinfo

import "fmt"

func (c *jiraClient) getSprintInfoFromJira(sprintID int) (sprintInfo SprintInfo, err error) {
	err = c.getDataFromJira(fmt.Sprintf("/rest/agile/1.0/sprint/%d", sprintID), nil, &sprintInfo)
	return
}
