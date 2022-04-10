package types

import (
	"encoding/json"
	"io"
	"strings"
)

type ComplianceType string

const (
	ComplianceTypeTweets ComplianceType = "tweets"
	ComplianceTypeUsers  ComplianceType = "users"
)

type BatchComplianceJobsPostParams struct {
	accessToken string

	// JSON body parameter
	Type      ComplianceType `json:"type,omitempty"`
	Name      *string        `json:"name,omitempty"`
	Resumable *bool          `json:"resumable,omitempty"`
}

func (p *BatchComplianceJobsPostParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BatchComplianceJobsPostParams) AccessToken() string {
	return p.accessToken
}

func (p *BatchComplianceJobsPostParams) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *BatchComplianceJobsPostParams) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *BatchComplianceJobsPostParams) ParameterMap() map[string]string {
	return map[string]string{}
}
