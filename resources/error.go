package resources

import (
	"fmt"

	"github.com/michimani/gotwi/internal/util"
)

type Non2XXError struct {
	Errors        []ErrorInformation         `json:"errors"`
	Title         string                     `json:"title,omitempty"`
	Detail        string                     `json:"detail,omitempty"`
	Type          string                     `json:"type,omitempty"`
	Status        string                     `json:"-"`
	StatusCode    int                        `json:"-"`
	RateLimitInfo *util.RateLimitInformation `json:"-"`
}

type ErrorInformation struct {
	Message    string              `json:"message"`
	Code       ErrorCode           `json:"code,omitempty"`
	Label      string              `json:"label,omitempty"`
	Parameters map[string][]string `json:"parameters,omitempty"`
}

func (e *Non2XXError) Summary() string {
	if e == nil {
		return ""
	}

	summary := ""
	if e.Status != "" {
		summary = summary + fmt.Sprintf("httpStatus=\"%s\" ", e.Status)
	}
	if e.StatusCode > 0 {
		summary = summary + fmt.Sprintf("httpStatusCode=%d ", e.StatusCode)
	}
	if e.Title != "" {
		summary = summary + fmt.Sprintf("title=\"%s\" ", e.Title)
	}
	if e.Detail != "" {
		summary = summary + fmt.Sprintf("detail=\"%s\" ", e.Detail)
	}
	for _, er := range e.Errors {
		summary = summary + fmt.Sprintf("message=\"%s\" ", er.Message)
		if er.Code > 0 {
			detail := er.Code.Detail()
			summary = summary + fmt.Sprintf("errorCode=%d errorText=\"%s\" errorDescription=\"%s\" ", er.Code, detail.Text, detail.Description)
		}
	}
	if e.RateLimitInfo != nil {
		summary = summary + fmt.Sprintf("rateLimit=%d rateLimitRemaining=%d rateLimitReset=\"%s\"", e.RateLimitInfo.Limit, e.RateLimitInfo.Remaining, e.RateLimitInfo.ResetAt)
	}

	return summary
}

type PartialError struct {
	ResourceType string `json:"resource_type"`
	Field        string `json:"field"`
	Parameter    string `json:"parameter"`
	ResourceId   string `json:"resource_id"`
	Title        string `json:"title"`
	Section      string `json:"section"`
	Detail       string `json:"detail"`
	Value        string `json:"value"`
	Type         string `json:"type"`
}
