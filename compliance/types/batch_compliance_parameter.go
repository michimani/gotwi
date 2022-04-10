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

type BatchComplianceJobsParams struct {
	accessToken string

	// Query Parameters
	Type   ComplianceType
	Status ComplianceStatus
}

var BatchComplianceJobsQueryParams = map[string]struct{}{
	"type":   {},
	"status": {},
}

func (p *BatchComplianceJobsParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BatchComplianceJobsParams) AccessToken() string {
	return p.accessToken
}

func (p *BatchComplianceJobsParams) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase

	if p.Type == "" {
		return ""
	}

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, BatchComplianceJobsQueryParams)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *BatchComplianceJobsParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BatchComplianceJobsParams) ParameterMap() map[string]string {
	m := map[string]string{}

	m["type"] = string(p.Type)

	if p.Status != "" {
		m["status"] = string(p.Status)
	}

	return m
}

type BatchComplianceJobsIDParams struct {
	accessToken string

	// Path parameters
	ID string
}

func (p *BatchComplianceJobsIDParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BatchComplianceJobsIDParams) AccessToken() string {
	return p.accessToken
}

func (p *BatchComplianceJobsIDParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	escaped := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", escaped, 1)

	return endpoint
}

func (p *BatchComplianceJobsIDParams) Body() (io.Reader, error) {
	return nil, nil
}

func (p *BatchComplianceJobsIDParams) ParameterMap() map[string]string {
	return map[string]string{}
}

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
