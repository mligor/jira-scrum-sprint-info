package jirascrumsprintinfo

import (
	"fmt"
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

func (s StoryPointInfo) String() string {

	return fmt.Sprintf("Comp-Init: %.2f, Comp: %.2f, Not-Comp-Init: %.2f, Not-Comp: %.2f, All: %.2f", s.CompletedIssuesInitialEstimate, s.CompletedIssuesEstimate, s.IssuesNotCompletedInitialEstimate, s.IssuesNotCompletedEstimate, s.AllIssuesEstimate)
}

func (s SprintInfo) String() string {

	const timeFormat = "2006-01-02 15:04:05"
	start := " ----------------- "
	end := " ----------------- "
	complete := " ----------------- "

	if s.StartDate != nil {
		start = s.StartDate.Format(timeFormat)
	}
	if s.EndDate != nil {
		end = s.EndDate.Format(timeFormat)
	}
	if s.CompleteDate != nil {
		complete = s.CompleteDate.Format(timeFormat)
	}

	return fmt.Sprintf("Sprint: %d (%s) - Timeline: %v - %v (%v) - StoryPoints: (%v) - '%s'", s.ID, s.State, start, end, complete, s.StoryPoint, s.Name)
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
