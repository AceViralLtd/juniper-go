package test

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

// RequestOpts can be used to modify the request behavior
type RequestOpts struct {
	Body        *url.Values
	BearerToken string
	Method      string
}

var defaultRequestOptions = &RequestOpts{
	Method: "GET",
}

// Client for testing http requests
type Client struct {
	router *gin.Engine
}

// NewClient for making test http requests
func NewClient(router *gin.Engine) *Client {
	return &Client{router: router}
}

// Request a page from the router
func (client *Client) Request(path string, opts ...*RequestOpts) *httptest.ResponseRecorder {
	var (
		reader  io.Reader
		options *RequestOpts
	)

	if len(opts) >= 1 {
		options = opts[0]
		if options.Method == "" {
			options.Method = defaultRequestOptions.Method
		}
	} else {
		options = defaultRequestOptions
	}

	if options.Body != nil {
		reader = strings.NewReader(options.Body.Encode())
	}

	recorder := httptest.NewRecorder()
	req, _ := http.NewRequest(options.Method, path, reader)

	if options.Method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if options.BearerToken != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", options.BearerToken))
	}

	client.router.ServeHTTP(recorder, req)

	return recorder
}

// outComeResponse is used to verify the outcome field in the response json
type outComeResponse struct {
	Outcome bool
}

// ParseOpts can be used to modify the parse logic
type ParseOpts struct {
	Code    int
	Outcome bool
}

var defaultParseOpts = &ParseOpts{
	Outcome: true,
	Code:    http.StatusOK,
}

// Respons wrapper for the faked request
type Response struct {
	Reader *httptest.ResponseRecorder
}

// Parse the response into a struct and verify the expected outcome
func (response *Response) Parse(t *testing.T, target any, opts ...*ParseOpts) {
	var options *ParseOpts

	if len(opts) >= 1 {
		options = opts[0]
		if options.Code == 0 {
			options.Code = defaultParseOpts.Code
		}
	} else {
		options = defaultParseOpts
	}

	if response.Reader.Code != options.Code {
		t.Fatalf(
			"didnt return %d (%d)\n%s",
			options.Code,
			response.Reader.Code,
			response.Reader.Body.String(),
		)
	}

	if target == nil {
		t.Fatal("You must provide a target struct to unmarshal the response into")
	}

	outcomeTarget := outComeResponse{}
	if err := json.Unmarshal(response.Reader.Body.Bytes(), &outcomeTarget); err != nil {
		t.Fatal("response json did not have an outcome field")
	}
	if outcomeTarget.Outcome != options.Outcome {
		t.Fatalf("did not return outcome: %t", options.Outcome)
	}

	if err := json.Unmarshal(response.Reader.Body.Bytes(), target); err != nil {
		t.Fatal("response json could not be unmarshaled into the target struct")
	}
}
