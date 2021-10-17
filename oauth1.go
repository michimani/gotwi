package gotwi

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	OAuthVersion10               = "1.0"
	OAuthSignatureMethodHMACSHA1 = "HMAC-SHA1"
)

type Endpoint string

type EndpointInfo struct {
	Raw                      string
	Base                     string
	EncodedQueryParameterMap map[string]string
}

type CreateOAthSignatureInput struct {
	HTTPMethod       string
	RawEndpoint      string
	OAuthConsumerKey string
	OAuthToken       string
	SigningKey       string
	ParameterMap     map[string]string
}

type CreateOAthSignatureOutput struct {
	OAuthNonce           string
	OAuthSignatureMethod string
	OAuthTimestamp       string
	OAuthVersion         string
	OAuthSignature       string
}

const oauth1header = `OAuth oauth_consumer_key="%s",oauth_nonce="%s",oauth_signature="%s",oauth_signature_method="%s",oauth_timestamp="%s",oauth_token="%s",oauth_version="%s"`

// SetOAuth1Header returns http.Request with the header information required for OAuth1.0a authentication.
func (c *TwitterClient) SetOAuth1Header(r *http.Request, paramsMap map[string]string) (*http.Request, error) {
	in := &CreateOAthSignatureInput{
		HTTPMethod:       r.Method,
		RawEndpoint:      r.URL.String(),
		OAuthConsumerKey: c.OAuthConsumerKey,
		OAuthToken:       c.OAuthToken,
		SigningKey:       c.SigningKey,
		ParameterMap:     paramsMap,
	}

	out, err := createOAuthSignature(in)
	if err != nil {
		return nil, err
	}

	r.Header.Add("Authorization", fmt.Sprintf(oauth1header,
		url.QueryEscape(c.OAuthConsumerKey),
		url.QueryEscape(out.OAuthNonce),
		url.QueryEscape(out.OAuthSignature),
		url.QueryEscape(out.OAuthSignatureMethod),
		url.QueryEscape(out.OAuthTimestamp),
		url.QueryEscape(c.OAuthToken),
		url.QueryEscape(out.OAuthVersion),
	))

	return r, nil
}

func createOAuthSignature(in *CreateOAthSignatureInput) (*CreateOAthSignatureOutput, error) {
	out := CreateOAthSignatureOutput{
		OAuthSignatureMethod: OAuthSignatureMethodHMACSHA1,
		OAuthVersion:         OAuthVersion10,
	}
	nonce := generateOAthNonce()
	out.OAuthNonce = nonce
	ts := fmt.Sprintf("%d", time.Now().Unix())
	out.OAuthTimestamp = ts
	endpointBase := endpointBase(in.RawEndpoint)

	parameterString := createParameterString(in.ParameterMap, nonce, ts, in)
	sigBase := createSignatureBase(in.HTTPMethod, endpointBase, parameterString)
	sig := calculateSignature(sigBase, in.SigningKey)
	out.OAuthSignature = sig
	return &out, nil
}

func generateOAthNonce() string {
	key := make([]byte, 32)
	rand.Read(key)
	nonce := base64.StdEncoding.EncodeToString(key)
	symbols := []string{"+", "/", "="}
	for _, s := range symbols {
		nonce = strings.Replace(nonce, s, "", -1)
	}
	return nonce
}

func endpointBase(e string) string {
	queryIdx := strings.Index(e, "?")
	if queryIdx < 0 {
		return e
	}

	return e[:queryIdx]
}

func (e Endpoint) String() string {
	return string(e)
}

func (e Endpoint) Detail() (*EndpointInfo, error) {
	d := EndpointInfo{
		Raw:                      e.String(),
		EncodedQueryParameterMap: map[string]string{},
	}

	queryIdx := strings.Index(e.String(), "?")
	if queryIdx < 0 {
		d.Base = string(e)
		return &d, nil
	}

	d.Base = e.String()[:queryIdx]
	queryPart := e.String()[queryIdx+1:]
	paramsPairs := strings.Split(queryPart, "&")
	for _, pp := range paramsPairs {
		keyValue := strings.Split(pp, "=")
		var err error
		v := ""
		if len(keyValue) == 2 {
			v, err = url.QueryUnescape(keyValue[1])
			if err != nil {
				return nil, err
			}
		}
		d.EncodedQueryParameterMap[keyValue[0]] = v
	}

	return &d, nil
}

func createParameterString(paramsMap map[string]string, nonce, ts string, in *CreateOAthSignatureInput) string {
	qv := url.Values{}
	for k, v := range paramsMap {
		qv.Add(k, v)
	}

	qv.Add("oauth_consumer_key", in.OAuthConsumerKey)
	qv.Add("oauth_nonce", nonce)
	qv.Add("oauth_signature_method", OAuthSignatureMethodHMACSHA1)
	qv.Add("oauth_timestamp", ts)
	qv.Add("oauth_token", in.OAuthToken)
	qv.Add("oauth_version", OAuthVersion10)

	return qv.Encode()
}

func createSignatureBase(method, endpointBase, parameterString string) string {
	return fmt.Sprintf(
		"%s&%s&%s",
		url.QueryEscape(strings.ToUpper(method)),
		url.QueryEscape(endpointBase),
		url.QueryEscape(parameterString),
	)
}

func calculateSignature(base, key string) string {
	b := []byte(key)
	h := hmac.New(sha1.New, b)
	io.WriteString(h, base)
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
