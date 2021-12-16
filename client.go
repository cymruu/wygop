package wygop

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

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

func (c *WykopClient) Post(endpoint string, body *url.Values) (*responses.APIResponse, error) {
	url := c.createURL(endpoint)
	request, _ := http.NewRequest("POST", url, strings.NewReader(body.Encode()))
	c.signRequest(request, body)

	return c.sendRequest(request)
}

func (c *WykopClient) Get(endpoint string) (*responses.APIResponse, error) {
	url := c.createURL(endpoint)
	request, _ := http.NewRequest("GET", url, nil)
	c.signRequest(request, nil)

	return c.sendRequest(request)
}

func (c *WykopClient) signRequest(request *http.Request, body *url.Values) {
	hashPayload := c.secret + request.URL.String()
	if body != nil {
		for _, v := range *body {
			hashPayload += v[0] + ","
		}
		hashPayload = hashPayload[:len(hashPayload)-1]
	}
	apisign := md5.Sum([]byte(hashPayload))
	request.Header.Add("apisign", fmt.Sprintf("%x", apisign))
}

func (c *WykopClient) createURL(endpoint string) string {
	url := fmt.Sprintf("https://a2.wykop.pl/%s/appkey/%s", endpoint, c.appkey)

	return url
}

func (c *WykopClient) sendRequest(request *http.Request) (*responses.APIResponse, error) {
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
