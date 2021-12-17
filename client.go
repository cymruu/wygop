package wygop

import (
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

	username   *string
	accountkey *string
}

func CreateClient(appkey, secret string, client *http.Client) *WykopClient {
	return &WykopClient{
		appkey:     appkey,
		secret:     secret,
		httpClient: client,
		username:   new(string),
		accountkey: new(string),
	}
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

func (c *WykopClient) CreateRequest(endpoint string, requestOptions ...RequestOptional) *WykopRequest {
	return CreateRequest("login/index", requestOptions...)
}

func (c *WykopClient) SendRequest(wykopRequest *WykopRequest) (*responses.APIResponse, error) {
	request, err := wykopRequest.createHTTPRequest()
	if err != nil {
		return nil, err
	}
	c.signRequest(request, wykopRequest.body)

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
