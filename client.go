package wygop

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
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

func (c *WykopClient) Login(accountkey string) (*responses.APIResponse, error) {
	body := url.Values{}
	body.Add("accountkey", accountkey)
	response, err := c.Post("login/index", &body)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *WykopClient) Post(endpoint string, body *url.Values) (*responses.APIResponse, error) {
	endpointUrl := c.createURL(endpoint)
	request, err := http.NewRequest("POST", endpointUrl, strings.NewReader(body.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, err
	}

	c.signRequest(request, body)

	return c.sendRequest(request)
}

func (c *WykopClient) Get(endpoint string) (*responses.APIResponse, error) {
	endpointUrl := c.createURL(endpoint)
	request, _ := http.NewRequest("GET", endpointUrl, nil)
	c.signRequest(request, nil)

	return c.sendRequest(request)
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
	}

	signBytes := md5.Sum([]byte(hashPayload)[:len(hashPayload)-1])
	request.Header.Add("apisign", fmt.Sprintf("%x", signBytes))
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
