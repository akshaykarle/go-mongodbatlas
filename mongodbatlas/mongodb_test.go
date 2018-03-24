package mongodbatlas

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var defaultTestTimeout = time.Second * 1

// testServer returns an http Client, ServeMux, and Server. The client proxies
// requests to the server and handlers can be registered on the mux to handle
// requests. The caller must close the test server.
func testServer() (*http.Client, *http.ServeMux, *httptest.Server) {
	mux := http.NewServeMux()
	server := httptest.NewServer(mux)
	transport := &RewriteTransport{&http.Transport{
		Proxy: func(req *http.Request) (*url.URL, error) {
			return url.Parse(server.URL)
		},
	}}
	client := &http.Client{Transport: transport}
	return client, mux, server
}

// RewriteTransport rewrites https requests to http to avoid TLS cert issues
// during testing.
type RewriteTransport struct {
	Transport http.RoundTripper
}

// RoundTrip rewrites the request scheme to http and calls through to the
// composed RoundTripper or if it is nil, to the http.DefaultTransport.
func (t *RewriteTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.URL.Scheme = "http"
	if t.Transport == nil {
		return http.DefaultTransport.RoundTrip(req)
	}
	return t.Transport.RoundTrip(req)
}

func assertMethod(t *testing.T, expectedMethod string, req *http.Request) {
	assert.Equal(t, expectedMethod, req.Method)
}

// assertReqJSON tests that the Request has the expected key values pairs json
// encoded in its Body
func assertReqJSON(t *testing.T, expected map[string]interface{}, req *http.Request) {
	var reqJSON interface{}
	err := json.NewDecoder(req.Body).Decode(&reqJSON)
	if err != nil {
		t.Errorf("error decoding request JSON %v", err)
	}
	assert.Equal(t, expected, reqJSON)
}

// assertReqJSONList tests that the Request has the expected list of key values
// pairs json encoded in its Body
func assertReqJSONList(t *testing.T, expected []interface{}, req *http.Request) {
	var reqJSON interface{}
	err := json.NewDecoder(req.Body).Decode(&reqJSON)
	if err != nil {
		t.Errorf("error decoding request JSON %v", err)
	}
	assert.Equal(t, expected, reqJSON)
}
