package services

import "github.com/cymruu/wygop"

type WykopService struct {
	Entries EntryService
	Login   LoginService
}

func CreateWykopService(client *wygop.WykopClient) *WykopService {
	return &WykopService{
		Entries: EntryService{client},
		Login:   LoginService{client},
	}
}
