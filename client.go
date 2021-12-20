package wygop

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"

	"github.com/cymruu/wygop/responses"
)

type WykopClient struct {
	appkey string
	secret string

	httpClient *http.Client

	userkey *string

	RequestOptions []WykopRequestOption
}

func CreateClient(appkey, secret string, client *http.Client) *WykopClient {
	return &WykopClient{
		appkey:     appkey,
		secret:     secret,
		httpClient: client,
	}
}

func (c *WykopClient) SetUserkey(userkey string) {
	c.userkey = &userkey
}

func (c *WykopClient) signRequest(request *http.Request, body *url.Values) {
	hashPayload := c.secret + request.URL.String()
	if body != nil {
		sortedKeys := make([]string, 0)
		for k := range *body {
			sortedKeys = append(sortedKeys, k)
		}
		sort.Strings(sortedKeys)
		for _, k := range sortedKeys {
			hashPayload += body.Get(k) + ","
		}
		if len(sortedKeys) > 0 {
			hashPayload = hashPayload[:len(hashPayload)-1]
		}
	}

	signBytes := md5.Sum([]byte(hashPayload))
	request.Header.Add("apisign", fmt.Sprintf("%x", signBytes))
}

func (c *WykopClient) CreateRequest(endpoint string, options ...WykopRequestOption) *WykopRequest {
	authParams := make(namedParams)
	if c.userkey != nil {
		authParams["userkey"] = *c.userkey
	}
	authParams["appkey"] = c.appkey

	authOptions := WithNamedParams(authParams)
	requestOptions := c.RequestOptions
	requestOptions = append(requestOptions, options...)
	requestOptions = append(requestOptions, authOptions)

	return CreateRequest(endpoint, requestOptions...)
}

func (c *WykopClient) SendRequest(ctx context.Context, wykopRequest *WykopRequest) (*responses.APIResponse, error) {
	request, err := wykopRequest.toHTTPRequest(ctx)
	if err != nil {
		return nil, err
	}
	c.signRequest(request, &wykopRequest.body)
	if wykopRequest.body != nil {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	res, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var response = responses.APIResponse{}
	err = json.Unmarshal(data, &response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, response.Error
	}

	return &response, nil
}
