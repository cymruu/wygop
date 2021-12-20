package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/cymruu/wygop"
	"github.com/cymruu/wygop/responses"
)

type EntryService struct {
	client *wygop.WykopClient
}

func (s *EntryService) Stream(ctx context.Context) ([]responses.Entry, error) {
	response, err := s.client.SendRequest(ctx, s.client.CreateRequest("entries/stream"))
	if err != nil {
		return nil, err
	}

	var entries []responses.Entry = make([]responses.Entry, 0)

	err = json.Unmarshal(response.Data, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (s *EntryService) Entry(ctx context.Context, entryId uint64) (*responses.Entry, error) {
	response, err := s.client.SendRequest(ctx, s.client.CreateRequest(fmt.Sprintf("entries/entry/%d", entryId)))
	if err != nil {
		return nil, err
	}

	var entry responses.Entry

	err = json.Unmarshal(response.Data, &entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (s *EntryService) Add(ctx context.Context, body *url.Values) (*responses.Entry, error) {
	request := s.client.CreateRequest("entries/add", wygop.WithPostBody(body))

	response, err := s.client.SendRequest(ctx, request)
	if err != nil {
		return nil, err
	}

	var entry responses.Entry

	err = json.Unmarshal(response.Data, &entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}
