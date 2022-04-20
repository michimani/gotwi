package types

import (
	"encoding/json"
	"io"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/fields"
	"github.com/michimani/gotwi/internal/util"
)

type ListRulesInput struct {
	accessToken string

	// Query parameters
	IDs []string
}

var listRulesQueryParameters = map[string]struct{}{
	"ids": {},
}

func (p *ListRulesInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *ListRulesInput) AccessToken() string {
	return p.accessToken
}

func (p *ListRulesInput) ResolveEndpoint(endpointBase string) string {
	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, listRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *ListRulesInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *ListRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.IDs != nil && len(p.IDs) > 0 {
		m["ids"] = util.QueryValue(p.IDs)
	}

	return m
}

type AddingRules []AddingRule

type AddingRule struct {
	Value *string `json:"value,omitempty"`
	Tag   *string `json:"tag,omitempty"`
}

type DeletingRules struct {
	IDs []string `json:"ids"`
}

type CreateRulesInput struct {
	accessToken string

	// Query parameters
	DryRun bool `json:"-"` // default false

	// JSON body parameter
	Add AddingRules `json:"add,omitempty"`
}

var createOrDeleteRulesQueryParameters = map[string]struct{}{
	"dry_run": {},
}

func (p *CreateRulesInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *CreateRulesInput) AccessToken() string {
	return p.accessToken
}

func (p *CreateRulesInput) ResolveEndpoint(endpointBase string) string {
	if len(p.Add) == 0 {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, createOrDeleteRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *CreateRulesInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *CreateRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m["dry_run"] = strconv.FormatBool(p.DryRun)
	return m
}

type DeleteRulesInput struct {
	accessToken string

	// Query parameters
	DryRun bool `json:"-"` // default false

	// JSON body parameter
	Delete *DeletingRules `json:"delete,omitempty"`
}

func (p *DeleteRulesInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *DeleteRulesInput) AccessToken() string {
	return p.accessToken
}

func (p *DeleteRulesInput) ResolveEndpoint(endpointBase string) string {
	if p.Delete == nil {
		return ""
	}

	endpoint := endpointBase
	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, createOrDeleteRulesQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *DeleteRulesInput) Body() (io.Reader, error) {
	json, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	return strings.NewReader(string(json)), nil
}

func (p *DeleteRulesInput) ParameterMap() map[string]string {
	m := map[string]string{}
	m["dry_run"] = strconv.FormatBool(p.DryRun)
	return m
}

type SearchStreamBackfillMinutes int

func (s SearchStreamBackfillMinutes) Valid() bool {
	return int(s) > 0
}

func (s SearchStreamBackfillMinutes) String() string {
	return strconv.Itoa(int(s))
}

type SearchStreamInput struct {
	accessToken string

	// Query parameters
	BackfillMinutes SearchStreamBackfillMinutes
	Expansions      fields.ExpansionList
	MediaFields     fields.MediaFieldList
	PlaceFields     fields.PlaceFieldList
	PollFields      fields.PollFieldList
	TweetFields     fields.TweetFieldList
	UserFields      fields.UserFieldList
}

var searchStreamQueryParameters = map[string]struct{}{
	"backfill_minutes": {},
	"expansions":       {},
	"media.fields":     {},
	"place.fields":     {},
	"poll.fields":      {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (p *SearchStreamInput) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *SearchStreamInput) AccessToken() string {
	return p.accessToken
}

func (p *SearchStreamInput) ResolveEndpoint(endpointBase string) string {
	if p == nil {
		return ""
	}

	endpoint := endpointBase

	pm := p.ParameterMap()
	if len(pm) > 0 {
		qs := util.QueryString(pm, searchStreamQueryParameters)
		endpoint += "?" + qs
	}

	return endpoint
}

func (p *SearchStreamInput) Body() (io.Reader, error) {
	return nil, nil
}

func (p *SearchStreamInput) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.BackfillMinutes.Valid() {
		m["backfill_minutes"] = p.BackfillMinutes.String()
	}

	m = fields.SetFieldsParams(m, p.Expansions, p.MediaFields, p.PlaceFields, p.PollFields, p.TweetFields, p.UserFields)

	return m
}
