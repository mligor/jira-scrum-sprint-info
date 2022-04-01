package jirascrumsprintinfo

type jiraClient struct {
	url   string
	token string
}

func (c *jiraClient) GetSprintInfo(sprintID int) (sprintInfo SprintInfo, err error) {

	sprintInfo, err = c.getSprintInfoFromJira(sprintID)
	if err != nil {
		return
	}

	report, err := c.getSprintReport(sprintInfo.BoardID, sprintID)
	if err != nil {
		return
	}

	sprintInfo.StoryPoint.CompletedIssuesEstimate = report.Contents.CompletedIssuesEstimateSum.Value
	sprintInfo.StoryPoint.AllIssuesEstimate = report.Contents.AllIssuesEstimateSum.Value
	sprintInfo.StoryPoint.CompletedIssuesInitialEstimate = report.Contents.CompletedIssuesInitialEstimateSum.Value
	sprintInfo.StoryPoint.IssuesNotCompletedEstimate = report.Contents.IssuesNotCompletedEstimateSum.Value
	sprintInfo.StoryPoint.IssuesNotCompletedInitialEstimate = report.Contents.IssuesNotCompletedInitialEstimateSum.Value
	return
}
