package services

import (
	"encoding/json"
	"net/url"

	"github.com/cymruu/wygop"
	"github.com/cymruu/wygop/responses"
)

type LoginService struct {
	client *wygop.WykopClient
}

func CreateLoginService(client *wygop.WykopClient) *LoginService {
	return &LoginService{client}
}

func (s *LoginService) Index(accountkey string) (*responses.Profile, error) {
	body := url.Values{}
	body.Add("accountkey", accountkey)
	response, err := s.client.Post("login/index", &body)
	if err != nil {
		return nil, err
	}

	var loginResult responses.Profile

	err = json.Unmarshal(response.Data, &loginResult)
	if err != nil {
		return nil, err
	}

	return &loginResult, nil
}
