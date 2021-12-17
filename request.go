package wygop

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type apiParams = []string
type namedParams = map[string]string
type RequestOptional func(*WykopRequest)

type WykopRequest struct {
	endpoint    string
	apiParams   apiParams
	namedParams namedParams
	body        *url.Values
}

func CreateRequest(endpoint string, options ...RequestOptional) *WykopRequest {
	request := &WykopRequest{
		endpoint: endpoint,
	}
	for _, applyOptionalFn := range options {
		applyOptionalFn(request)
	}

	return request
}

func SetApiParams(v apiParams) RequestOptional {
	return func(r *WykopRequest) {
		r.apiParams = v
	}
}

func SetNamedParams(v namedParams) RequestOptional {
	return func(r *WykopRequest) {
		r.namedParams = v
	}
}

func SetPostBody(v *url.Values) RequestOptional {
	return func(r *WykopRequest) {
		r.body = v
	}
}

func (req *WykopRequest) getRequestMethod() string {
	if req.body != nil {
		return "POST"
	}
	return "GET"
}

func (req *WykopRequest) createHTTPRequest() (*http.Request, error) {
	URL := fmt.Sprintf("https://a2.wykop.pl/%s/", req.endpoint)

	if req.apiParams != nil {
		URL += fmt.Sprintf("%s/", strings.Join(req.apiParams, "/"))
	}
	if req.namedParams != nil {
		for k, v := range req.namedParams {
			URL += fmt.Sprintf("%s/%s/,", k, v)
		}
	}

	method := req.getRequestMethod()

	var body io.Reader
	if req.body != nil {
		body = strings.NewReader(req.body.Encode())
	}

	return http.NewRequest(
		method,
		URL,
		body,
	)
}
