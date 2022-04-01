package jirascrumsprintinfo

import (
	"strconv"
)

type sprintReport struct {
	Contents sprintReportContents `json:"contents,omitempty"`
}

type sprintReportContents struct {
	CompletedIssuesInitialEstimateSum    sprintReportSPValue `json:"completedIssuesInitialEstimateSum,omitempty"`
	CompletedIssuesEstimateSum           sprintReportSPValue `json:"completedIssuesEstimateSum,omitempty"`
	IssuesNotCompletedInitialEstimateSum sprintReportSPValue `json:"issuesNotCompletedInitialEstimateSum,omitempty"`
	IssuesNotCompletedEstimateSum        sprintReportSPValue `json:"issuesNotCompletedEstimateSum,omitempty"`
	AllIssuesEstimateSum                 sprintReportSPValue `json:"allIssuesEstimateSum,omitempty"`
}

type sprintReportSPValue struct {
	Value float64 `json:"value,omitempty"`
	Text  string  `json:"text,omitempty"`
}

func (c *jiraClient) getSprintReport(boardID int, sprinID int) (report sprintReport, err error) {

	err = c.getDataFromJira("/rest/greenhopper/1.0/rapid/charts/sprintreport", map[string]string{
		"rapidViewId": strconv.Itoa(boardID),
		"sprintId":    strconv.Itoa(sprinID),
	}, &report)
	return
}
