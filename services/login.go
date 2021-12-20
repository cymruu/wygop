package services

import (
	"context"
	"encoding/json"
	"net/url"

	"github.com/cymruu/wygop"
	"github.com/cymruu/wygop/responses"
)

type LoginService struct {
	client *wygop.WykopClient
}

func (s *LoginService) Index(ctx context.Context, accountkey string) (*responses.LoginResult, error) {
	body := url.Values{}
	body.Add("accountkey", accountkey)
	request := s.client.CreateRequest("login/index", wygop.WithPostBody(&body))
	response, err := s.client.SendRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	var loginResult responses.LoginResult

	err = json.Unmarshal(response.Data, &loginResult)
	if err != nil {
		return nil, err
	}

	s.client.SetUserkey(loginResult.UserKey)
	return &loginResult, nil
}
