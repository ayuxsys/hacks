package main

import "encoding/json"

type Org struct {
	Name  string       `json:"name"`
	Cloud GitCloudType `json:"cloud"`
}

func NewOrg(name string, cloud GitCloudType) Org {
	return Org{name, cloud}
}

type GitCloudType int

const (
	Github GitCloudType = iota + 1
)

func (g GitCloudType) String() string {
	switch g {
	case Github:
		return "github"
	default:
		return "unknown"
	}
}

func (g GitCloudType) MarshalJSON() ([]byte, error) {
	return json.Marshal(g.String())
}
