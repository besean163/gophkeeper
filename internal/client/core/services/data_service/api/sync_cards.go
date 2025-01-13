package api

import models "github.com/besean163/gophkeeper/internal/models/client"

func (s Service) SyncCards(user models.User) error {
	var err error
	err = s.syncCardsOnServer(user)
	if err != nil {
		return err
	}
	err = s.syncCardsOnClient(user)
	if err != nil {
		return err
	}
	return nil
}
