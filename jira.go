package jirascrumsprintinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"path"
)

type issue struct {
	ID     string                 `json:"id,omitempty"`
	Key    string                 `json:"key,omitempty"`
	Fields map[string]interface{} `json:"fields,omitempty"`
}

type IssueResponse struct {
	StartAt    int     `json:"startAt,omitempty"`
	MaxResults int     `json:"maxResults,omitempty"`
	Total      int     `json:"total,omitempty"`
	Issues     []issue `json:"issues,omitempty"`
}

// API: https://docs.atlassian.com/jira-software/REST/7.3.1/

func (c *jiraClient) getDataFromJira(namepoint string, queryParams map[string]string, data interface{}) (err error) {

	u, err := url.Parse(c.url)
	u.Path = path.Join(u.Path, namepoint)
	requestURL := u.String()

	if len(queryParams) > 0 {
		var urlA *url.URL
		urlA, err = url.Parse(requestURL)
		if err != nil {
			return
		}
		params := urlA.Query()
		for k, v := range queryParams {
			params.Add(k, v)
		}
		urlA.RawQuery = params.Encode()
		requestURL = urlA.String()
	}

	req, err := http.NewRequest("GET", requestURL, nil)
	if err != nil {
		return
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&data)

	return
}
