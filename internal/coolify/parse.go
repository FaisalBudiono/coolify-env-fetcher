package coolify

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EnvObject struct {
	UUID             string `json:"uuid"`
	ApplicationID    int    `json:"application_id"`
	IsBuildTime      bool   `json:"is_build_time"`
	IsBuildTimeOnly  bool   `json:"is_build_time_only"`
	IsLiteral        bool   `json:"is_literal"`
	IsMultiline      bool   `json:"is_multiline"`
	IsPreview        bool   `json:"is_preview"`
	IsReallyRequired bool   `json:"is_really_required"`
	IsRequired       bool   `json:"is_required"`
	IsShared         bool   `json:"is_shared"`
	IsShownOnce      bool   `json:"is_shown_once"`
	Key              string `json:"key"`
	Order            int    `json:"order"`
	RealValue        string `json:"real_value"`
	Value            string `json:"value"`
	Version          string `json:"version"`
	CreatedAt        string `json:"created_at"`
	UpdatedAt        string `json:"updated_at"`
}

func (e EnvObject) IsBuildENV() bool {
	return e.IsBuildTime || e.IsBuildTimeOnly
}

type errorResponse struct {
	Message string `json:"message"`
}

func ParseENV(baseURL, appID, accessToken string) ([]EnvObject, error) {
	req, err := http.NewRequest(http.MethodGet, url(baseURL, appID), nil)
	if err != nil {
		return []EnvObject{}, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return []EnvObject{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []EnvObject{}, err
	}

	if resp.StatusCode != http.StatusOK {
		var eR errorResponse
		err = json.Unmarshal(body, &eR)
		if err != nil {
			return []EnvObject{}, err
		}

		return []EnvObject{}, fmt.Errorf("Coolify error %d: %s", resp.StatusCode, eR.Message)
	}

	var res []EnvObject
	err = json.Unmarshal(body, &res)
	if err != nil {
		return []EnvObject{}, err
	}

	return res, nil
}

func url(baseURL, projectID string) string {
	return fmt.Sprintf(
		"%s/api/v1/applications/%s/envs",
		baseURL,
		projectID,
	)
}
