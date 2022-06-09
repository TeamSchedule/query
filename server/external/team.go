package external

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

type Team struct {
	Id           int64   `json:"id"`
	Name         string  `json:"name"`
	CreationDate string  `json:"creationDate"`
	AdminId      int64   `json:"adminId"`
	MembersIds   []int64 `json:"membersIds"`
	Color        string  `json:"color"`
}

type TeamJson struct {
	Team Team
}

func GetTeamById(teamId string, client *resty.Client) (*Team, error) {
	resp, err := client.R().
		SetPathParams(map[string]string{
			"teamId": teamId,
		}).
		Get("/schedule/team/{teamId}")

	if err != nil {
		return nil, err
	}

	// TODO: check status code resp.StatusCode()

	var team TeamJson
	err = json.Unmarshal(resp.Body(), &team)

	if err != nil {
		return nil, err
	}

	return &team.Team, nil
}
