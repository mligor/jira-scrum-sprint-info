package jirascrumsprintinfo

import (
	"time"
)

type StoryPointInfo struct {
	CompletedIssuesInitialEstimate    float64 `json:"completed_issues_initial_estimate,omitempty"`
	CompletedIssuesEstimate           float64 `json:"completed_issues_estimate,omitempty"`
	IssuesNotCompletedInitialEstimate float64 `json:"issues_not_completed_initial_estimate,omitempty"`
	IssuesNotCompletedEstimate        float64 `json:"issues_not_completed_estimate,omitempty"`
	AllIssuesEstimate                 float64 `json:"all_issues_estimate,omitempty"`
}

type SprintInfo struct {
	ID           int            `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	CompleteDate *time.Time     `json:"completeDate,omitempty"`
	EndDate      *time.Time     `json:"endDate,omitempty"`
	StartDate    *time.Time     `json:"startDate,omitempty"`
	State        string         `json:"state,omitempty"`
	StoryPoint   StoryPointInfo `json:"story_point,omitempty"`
	BoardID      int            `json:"originBoardId,omitempty"`
}

type JiraClient interface {
	GetSprintInfo(sprintID int) (sprintInfo SprintInfo, err error)
}

func NewClient(url string, token string) (client JiraClient) {

	return &jiraClient{
		url:   url,
		token: token,
	}
}
