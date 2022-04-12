package types

import (
	"encoding/json"
	"io"
	"net/url"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type ComplianceType string

const (
	ComplianceTypeTweets ComplianceType = "tweets"
	ComplianceTypeUsers  ComplianceType = "users"
)

type ComplianceStatus string

const (
	ComplianceStatusCreated    ComplianceStatus = "created"
	ComplianceStatusInProgress ComplianceStatus = "in_progress"
	ComplianceStatusFailed     ComplianceStatus = "failed"
	ComplianceStatusComplete   ComplianceStatus = "complete"
)

type ListJobsInput struct {
	accessToken string

	// Query Parameters
	Type   ComplianceType
	Status ComplianceStatus
}

var listJobsQueryParameters = map[string]struct{}{
	"type":   {},
	"status": {},
}

func (p *ListJobsInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListJobsInput) AccessToken() string {
	return p.accessToken
}

func (p *ListJobsInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Type == "" {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listJobsQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListJobsInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListJobsInput) ParameterMap() map[string]string {
	m := map[string]string{}

	m["type"] = string(p.Type)

	if p.Status != "" {
		m["status"] = string(p.Status)
	}

	return m
}

type GetJobInput struct {
	accessToken string

	// Path parameters
	ID string
}

func (p *GetJobInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *GetJobInput) AccessToken() string {
	return p.accessToken
}

func (p *GetJobInput) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *GetJobInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *GetJobInput) ParameterMap() map[string]string {
	return map[string]string{}
}

type CreateJobInput struct {
	accessToken string

	// JSON body parameter
	Type      ComplianceType `json:"type,omitempty"` // required
	Name      *string        `json:"name,omitempty"`
	Resumable *bool          `json:"resumable,omitempty"`
}

func (p *CreateJobInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateJobInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateJobInput) ResolveEndpoint(endpointBase string) string {
	return endpointBase
}

func (p *CreateJobInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateJobInput) ParameterMap() map[string]string {
	return map[string]string{}
}
