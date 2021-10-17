package types

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/michimani/gotwi/internal/util"
)

type BlocksMaxResult int

type BlocksBlockingGetParams struct {
	accessToken string

	// Path parameter
	ID string

	// Query parameters
	MaxResult       BlocksMaxResult
	PaginationToken string
	Expansions      []string
	TweetFields     []string
	UserFields      []string
}

var BlocksBlockingGetQueryParams = map[string]struct{}{
	"max_results":      {},
	"pagination_token": {},
	"expansions":       {},
	"tweet.fields":     {},
	"user.fields":      {},
}

func (m BlocksMaxResult) Valid() bool {
	return m > 0 && m <= 1000
}

func (m BlocksMaxResult) String() string {
	return strconv.Itoa(int(m))
}

func (p *BlocksBlockingGetParams) SetAccessToken(token string) {
	p.accessToken = token
}

func (p *BlocksBlockingGetParams) AccessToken() string {
	return p.accessToken
}

func (p *BlocksBlockingGetParams) ResolveEndpoint(endpointBase string) string {
	if p.ID == "" {
		return ""
	}

	encoded := url.QueryEscape(p.ID)
	endpoint := strings.Replace(endpointBase, ":id", encoded, 1)

	pm := p.ParameterMap()
	qs := util.QueryString(pm, BlocksBlockingGetQueryParams)

	if qs == "" {
		return endpoint
	}

	return endpoint + "?" + qs
}

func (p *BlocksBlockingGetParams) Body() io.Reader {
	return nil
}

func (p *BlocksBlockingGetParams) ParameterMap() map[string]string {
	m := map[string]string{}

	if p.MaxResult.Valid() {
		m["max_results"] = p.MaxResult.String()
	}

	if p.PaginationToken != "" {
		m["pagination_token"] = p.PaginationToken
	}

	if p.Expansions != nil && len(p.Expansions) > 0 {
		m["expansions"] = util.QueryValue(p.Expansions)
	}

	if p.TweetFields != nil && len(p.TweetFields) > 0 {
		m["tweet.fields"] = util.QueryValue(p.TweetFields)
	}

	if p.UserFields != nil && len(p.UserFields) > 0 {
		m["user.fields"] = util.QueryValue(p.UserFields)
	}

	return m
}
