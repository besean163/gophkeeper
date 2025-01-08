package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	changedetector "github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/change_detector"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/syncer"
	"github.com/besean163/gophkeeper/internal/logger"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	pencrypt "github.com/besean163/gophkeeper/internal/utils/password_encrypter"
	timecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller"
)

type ChangeDetector interface {
	GetAccountChanges(user models.User, items []models.Account, externalItems []models.ExternalAccount) (created []models.Account, updated []models.Account, deleted []models.Account)
	GetNoteChanges(user models.User, items []models.Note, externalItems []models.ExternalNote) (created []models.Note, updated []models.Note, deleted []models.Note)
	GetCardChanges(user models.User, items []models.Card, externalItems []models.ExternalCard) (created []models.Card, updated []models.Card, deleted []models.Card)
}

type Service struct {
	logger         logger.Logger
	storeService   interfaces.DataService
	apiClient      interfaces.ApiClient
	encrypter      pencrypt.Encrypter
	timeController timecontroller.TimeController
	syncer         syncer.Syncer
	changeDetector ChangeDetector
}

func NewService(storeService interfaces.DataService, apiClient interfaces.ApiClient, encrypter pencrypt.Encrypter, timeController timecontroller.TimeController, logger logger.Logger, syncer syncer.Syncer, changeDetector ChangeDetector) Service {
	if logger == nil {
		logger = defaultlogger.NewDefaultLogger()
	}

	service := Service{
		storeService:   storeService,
		apiClient:      apiClient,
		logger:         logger,
		encrypter:      encrypter,
		timeController: timeController,
		syncer:         syncer,
		changeDetector: changeDetector,
	}

	if service.changeDetector == nil {
		service.changeDetector = changedetector.NewChangeDetector()
	}

	if service.syncer == nil {
		service.syncer = service
	}

	return service
}

func (s Service) SyncAll(user models.User) error {
	var err error
	err = s.SyncAccounts(user)
	if err != nil {
		return err
	}

	err = s.SyncNotes(user)
	if err != nil {
		return err
	}

	err = s.SyncCards(user)
	if err != nil {
		return err
	}

	s.logger.Debug("sync all ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

func (s Service) SyncAccounts(user models.User) error {
	var err error
	err = s.syncAccountsOnServer(user)
	if err != nil {
		return err
	}
	err = s.syncAccountsOnClient(user)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) syncAccountsOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetAccounts(user)
	if err != nil {
		return err
	}

	apiItems := make([]entities.AccountSyncInput, 0)
	for _, a := range items {
		apiItem := entities.AccountSyncInput{
			UUID:      a.UUID,
			Name:      a.Name,
			Login:     a.Login,
			Password:  a.Password,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := entities.AccountsSyncInput{
		Accounts: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncAccounts(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync accounts on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

func (s Service) syncAccountsOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetAccounts(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetAccounts()
	if err != nil {
		return err
	}

	externalItems := make([]models.ExternalAccount, 0)
	for _, apiItem := range apiOutput.Accounts {
		externalItem := models.ExternalAccount{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Login:     apiItem.Login,
			Password:  apiItem.Password,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	created, updated, deleted := s.changeDetector.GetAccountChanges(user, items, externalItems)

	for _, item := range created {
		err := s.storeService.CreateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := s.storeService.UpdateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range deleted {
		err := s.storeService.DeleteAccount(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync accounts on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

func (s Service) SyncNotes(user models.User) error {
	var err error
	err = s.syncNotesOnServer(user)
	if err != nil {
		return err
	}
	err = s.syncNotesOnClient(user)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) syncNotesOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetNotes(user)
	if err != nil {
		return err
	}

	apiItems := make([]entities.NoteSyncInput, 0)
	for _, a := range items {
		apiItem := entities.NoteSyncInput{
			UUID:      a.UUID,
			Name:      a.Name,
			Content:   a.Content,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := entities.NotesSyncInput{
		Notes: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncNotes(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync notes on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

func (s Service) syncNotesOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetNotes(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetNotes()
	if err != nil {
		return err
	}

	externalItems := make([]models.ExternalNote, 0)
	for _, apiItem := range apiOutput.Notes {
		externalItem := models.ExternalNote{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Content:   apiItem.Content,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	created, updated, deleted := s.changeDetector.GetNoteChanges(user, items, externalItems)

	for _, item := range created {
		err := s.storeService.CreateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := s.storeService.UpdateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range deleted {
		err := s.storeService.DeleteNote(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync notes on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

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

func (s Service) syncCardsOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetCards(user)
	if err != nil {
		return err
	}

	apiItems := make([]entities.CardSyncInput, 0)
	for _, a := range items {
		apiItem := entities.CardSyncInput{
			UUID:      a.UUID,
			Name:      a.Name,
			Number:    a.Number,
			Exp:       a.Exp,
			CVV:       a.CVV,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := entities.CardsSyncInput{
		Cards: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncCards(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync cards on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}

func (s Service) syncCardsOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetCards(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetCards()
	if err != nil {
		return err
	}

	externalItems := make([]models.ExternalCard, 0)
	for _, apiItem := range apiOutput.Cards {
		externalItem := models.ExternalCard{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Number:    apiItem.Number,
			Exp:       apiItem.Exp,
			CVV:       apiItem.CVV,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	created, updated, deleted := s.changeDetector.GetCardChanges(user, items, externalItems)

	for _, item := range created {
		err := s.storeService.CreateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := s.storeService.UpdateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range deleted {
		err := s.storeService.DeleteCard(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync cards on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
