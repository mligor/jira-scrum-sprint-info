package jirascrumsprintinfo

import "fmt"

type boardConfig struct {
	Estimation boardEstimation `json:"estimation,omitempty"`
}

type boardEstimation struct {
	Type  string               `json:"type,omitempty"`
	Field boardEstimationField `json:"field,omitempty"`
}

type boardEstimationField struct {
	FieldID     string `json:"fieldId,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

func (c *jiraClient) getEstimationField(boardID int) (fieldName string, err error) {

	requestURL := fmt.Sprintf("%s/rest/agile/1.0/board/%d/configuration", c.url, boardID)

	var e boardConfig

	err = c.getDataFromJira(requestURL, nil, &e)

	if err != nil {
		return
	}

	fieldName = e.Estimation.Field.FieldID

	return
}
