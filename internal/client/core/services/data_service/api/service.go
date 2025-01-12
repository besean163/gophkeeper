package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"

	models "github.com/besean163/gophkeeper/internal/models/client"

	changedetector "github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/change_detector"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/syncer"
	"github.com/besean163/gophkeeper/internal/logger"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"
	pencrypt "github.com/besean163/gophkeeper/internal/utils/password_encrypter"
	timecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller"
)

const (
	SyncNodeAccount = iota
	SyncNodeNote
	SyncNodeCard
)

type ChangeDetector interface {
	GetAccountChanges(user models.User, compare changes.AccountCompare) changes.AccountChanges
	GetNoteChanges(user models.User, compare changes.NoteCompare) changes.NoteChanges
	GetCardChanges(user models.User, compare changes.CardCompare) changes.CardChanges
}

type ServiceOptions struct {
	interfaces.DataService
	interfaces.ApiClient
	pencrypt.Encrypter
	logger.Logger
	syncer.Syncer
	ChangeDetector
	timecontroller.TimeController
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

func NewService(options ServiceOptions) Service {
	if options.Logger == nil {
		options.Logger = defaultlogger.NewDefaultLogger()
	}

	service := Service{
		storeService:   options.DataService,
		apiClient:      options.ApiClient,
		logger:         options.Logger,
		encrypter:      options.Encrypter,
		timeController: options.TimeController,
		syncer:         options.Syncer,
		changeDetector: options.ChangeDetector,
	}

	if service.changeDetector == nil {
		service.changeDetector = changedetector.NewChangeDetector()
	}

	if service.syncer == nil {
		service.syncer = service
	}

	return service
}
