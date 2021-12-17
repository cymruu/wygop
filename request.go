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
	body        url.Values
}

func CreateRequest(endpoint string, options ...RequestOptional) *WykopRequest {
	request := &WykopRequest{
		endpoint:    endpoint,
		apiParams:   make([]string, 0),
		namedParams: make(map[string]string),
		body:        url.Values{},
	}
	for _, applyOptionalFn := range options {
		applyOptionalFn(request)
	}

	return request
}

func SetApiParams(v apiParams) RequestOptional {
	return func(r *WykopRequest) {
		r.apiParams = append(r.apiParams, v...)
	}
}

func SetNamedParams(params namedParams) RequestOptional {
	return func(r *WykopRequest) {
		for k, v := range params {
			r.namedParams[k] = v
		}
	}
}

func SetPostBody(values *url.Values) RequestOptional {
	return func(r *WykopRequest) {
		for k, v := range *values {
			r.body[k] = v
		}
	}
}

func (req *WykopRequest) getRequestMethod() string {
	if req.body != nil {
		return "POST"
	}
	return "GET"
}

func (req *WykopRequest) createHTTPRequest() (*http.Request, error) {
	URL := fmt.Sprintf("https://a2.wykop.pl/%s", req.endpoint)

	URL += fmt.Sprintf("%s/", strings.Join(req.apiParams, "/"))

	for k, v := range req.namedParams {
		URL += fmt.Sprintf("%s/%s/", k, v)
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
