package services

import (
	"encoding/json"

	"github.com/cymruu/wygop"
	"github.com/cymruu/wygop/responses"
)

type EntryService struct {
	client *wygop.WykopClient
}

func (s *EntryService) Stream() ([]responses.Entry, error) {
	response, err := s.client.Get("entries/stream")
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

func (s *EntryService) Entry(entryId int64) (*responses.Entry, error) {
	response, err := s.client.Get("entries/stream")
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
