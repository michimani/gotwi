package gotwi_test

import (
	"testing"

	"github.com/michimani/gotwi"
	"github.com/stretchr/testify/assert"
)

func Test_EndpointDetail(t *testing.T) {
	cases := []struct {
		name     string
		endpoint gotwi.Endpoint
		expect   *gotwi.EndpointInfo
	}{
		{
			name:     "ok",
			endpoint: "endpoint",
			expect: &gotwi.EndpointInfo{
				Raw:                      "endpoint",
				Base:                     "endpoint",
				EncodedQueryParameterMap: map[string]string{},
			},
		},
		{
			name:     "ok with some parameters",
			endpoint: "endpoint?key1=value1&key2=value2",
			expect: &gotwi.EndpointInfo{
				Raw:  "endpoint?key1=value1&key2=value2",
				Base: "endpoint",
				EncodedQueryParameterMap: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
		},
		{
			name:     "ok with encoded parameter",
			endpoint: "endpoint?key1=value1&key2=value2&key3=value%20value3",
			expect: &gotwi.EndpointInfo{
				Raw:  "endpoint?key1=value1&key2=value2&key3=value%20value3",
				Base: "endpoint",
				EncodedQueryParameterMap: map[string]string{
					"key1": "value1",
					"key2": "value2",
					"key3": "value value3",
				},
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			ed, err := c.endpoint.Detail()
			assert.NoError(tt, err)
			assert.Equal(tt, c.expect.Raw, ed.Raw)
			assert.Equal(tt, c.expect.Base, ed.Base)
			assert.Equal(tt, len(c.expect.EncodedQueryParameterMap), len(ed.EncodedQueryParameterMap))
			for k, v := range c.expect.EncodedQueryParameterMap {
				assert.Equal(tt, v, ed.EncodedQueryParameterMap[k])
			}
		})
	}
}

func Test_CreateOAuthSignature(t *testing.T) {
	cases := []struct {
		name    string
		in      *gotwi.CreateOAuthSignatureInput
		wantErr bool
	}{
		{
			name: "normal",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     map[string]string{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "normal: parameter map is nil",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     nil,
			},
			wantErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			out, err := gotwi.CreateOAuthSignature(c.in)
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.NoError(tt, err)
			assert.NotEmpty(tt, out.OAuthNonce)
			assert.Equal(tt, "HMAC-SHA1", out.OAuthSignatureMethod)
			assert.Equal(tt, "1.0", out.OAuthVersion)
			assert.NotEmpty(tt, out.OAuthTimestamp)
			assert.NotEmpty(tt, out.OAuthSignature)
		})
	}
}

func Test_generateOAthNonce(t *testing.T) {
	cases := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "normal",
			wantErr: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			n, err := gotwi.ExportGenerateOAthNonce()
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.NoError(tt, err)
			assert.NotEmpty(tt, n)
		})
	}
}
func Test_endpointBase(t *testing.T) {
	cases := []struct {
		name     string
		endpoint string
		expect   string
	}{
		{
			name:     "normal: has no query parameters",
			endpoint: "endpoint",
			expect:   "endpoint",
		},
		{
			name:     "normal: has query parameters",
			endpoint: "endpoint?p1=v1&p2=v2",
			expect:   "endpoint",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			eb := gotwi.ExportEndpointBase(c.endpoint)

			assert.Equal(tt, c.expect, eb)
		})
	}
}

func Test_createParameterString(t *testing.T) {
	cases := []struct {
		name   string
		in     *gotwi.CreateOAuthSignatureInput
		nonce  string
		ts     string
		expect string
	}{
		{
			name: "normal: parameter map is nil",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     nil,
			},
			nonce:  "nonce",
			ts:     "1234567890",
			expect: "oauth_consumer_key=o-auth-consumer-key&oauth_nonce=nonce&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1234567890&oauth_token=o-auth-token&oauth_version=1.0",
		},
		{
			name: "normal: one parameter",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     map[string]string{"key1": "value1"},
			},
			nonce:  "nonce",
			ts:     "1234567890",
			expect: "key1=value1&oauth_consumer_key=o-auth-consumer-key&oauth_nonce=nonce&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1234567890&oauth_token=o-auth-token&oauth_version=1.0",
		},
		{
			name: "normal: two parameters",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     map[string]string{"key1": "value1", "key2": "value2"},
			},
			nonce:  "nonce",
			ts:     "1234567890",
			expect: "key1=value1&key2=value2&oauth_consumer_key=o-auth-consumer-key&oauth_nonce=nonce&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1234567890&oauth_token=o-auth-token&oauth_version=1.0",
		},
		{
			name: "normal: parameter with white space",
			in: &gotwi.CreateOAuthSignatureInput{
				HTTPMethod:       "POST",
				RawEndpoint:      "raw-endpoint",
				OAuthConsumerKey: "o-auth-consumer-key",
				OAuthToken:       "o-auth-token",
				SigningKey:       "signing-key",
				ParameterMap:     map[string]string{"key1": "value 1", "key2": "value 2"},
			},
			nonce:  "nonce",
			ts:     "1234567890",
			expect: "key1=value%201&key2=value%202&oauth_consumer_key=o-auth-consumer-key&oauth_nonce=nonce&oauth_signature_method=HMAC-SHA1&oauth_timestamp=1234567890&oauth_token=o-auth-token&oauth_version=1.0",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			s := gotwi.ExportCreateParameterString(c.nonce, c.ts, c.in)
			assert.Equal(tt, c.expect, s)
		})
	}
}

func Test_createSignatureBase(t *testing.T) {
	cases := []struct {
		name                 string
		method               string
		endpointBase         string
		queryParameterString string
		expect               string
	}{
		{
			name:                 "normal: lower method",
			method:               "get",
			endpointBase:         "endpoint",
			queryParameterString: "",
			expect:               "GET&endpoint&",
		},
		{
			name:                 "normal: query parameter is empty",
			method:               "GET",
			endpointBase:         "endpoint",
			queryParameterString: "",
			expect:               "GET&endpoint&",
		},
		{
			name:                 "normal: has parameters",
			method:               "GET",
			endpointBase:         "endpoint",
			queryParameterString: "key1=value1&key2=value2",
			expect:               "GET&endpoint&key1%3Dvalue1%26key2%3Dvalue2",
		},
		{
			name:                 "normal: has parameters with white space",
			method:               "GET",
			endpointBase:         "endpoint",
			queryParameterString: "key1=value%201&key2=value%202",
			expect:               "GET&endpoint&key1%3Dvalue%25201%26key2%3Dvalue%25202",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			b := gotwi.ExportCreateSignatureBase(c.method, c.endpointBase, c.queryParameterString)
			assert.Equal(tt, c.expect, b)
		})
	}
}

func Test_calculateSignature(t *testing.T) {
	cases := []struct {
		name    string
		base    string
		key     string
		wantErr bool
		expect  string
	}{
		{
			name:    "normal",
			base:    "base",
			key:     "key",
			wantErr: false,
			expect:  "LdQ0mqLyDXoda6+8WAf8tcglIME=",
		},
		{
			name:    "normal: base is empty",
			base:    "",
			key:     "key",
			wantErr: false,
			expect:  "9Cuw7rAY671Fl65yE3EexgdghD8=",
		},
		{
			name:    "normal: key is empty",
			base:    "base",
			key:     "",
			wantErr: false,
			expect:  "Frqq9jhCJCjivrWB8RgjZ3hwiHY=",
		},
		{
			name:    "normal: base and key are empty",
			base:    "",
			key:     "",
			wantErr: false,
			expect:  "+9sdGxiqbAgyS31ktx+3Y3BpDh0=",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(tt *testing.T) {
			s, err := gotwi.ExportCalculateSignature(c.base, c.key)
			if c.wantErr {
				assert.Error(tt, err)
				return
			}

			assert.NoError(tt, err)
			assert.Equal(tt, c.expect, s)
		})
	}
}
