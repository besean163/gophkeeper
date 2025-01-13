#!/usr/bin/bash

MOCKGEN_PATH="/home/besean/go/bin/mockgen"
PROJECT_PATH="/opt/gophkeeper"
SERVER_INTERFACE_PATH="internal/server/interfaces"
SERVER_TEST_PATH="internal/server/tests"
CLIENT_INTERFACE_PATH="internal/client/core/interfaces"
CLIENT_MOCK_PATH="internal/client/tests/mocks"
UTILS_INTERFACE_PATH="internal/utils"
UTILS_MOCK_PATH="internal/tests/mocks/utils"

# моки интерфейсов сервера
$MOCKGEN_PATH -source=$PROJECT_PATH/$SERVER_INTERFACE_PATH/auth_service.go -destination=$PROJECT_PATH/$SERVER_TEST_PATH/mocks/auth_service.go -package=mock AuthService 
$MOCKGEN_PATH -source=$PROJECT_PATH/$SERVER_INTERFACE_PATH/bucket_service.go -destination=$PROJECT_PATH/$SERVER_TEST_PATH/mocks/bucket_service.go -package=mock BucketService

# моки интерфейсов клиента
$MOCKGEN_PATH -source=$PROJECT_PATH/$CLIENT_INTERFACE_PATH/data_service.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/data_service.go -package=mock DataService
$MOCKGEN_PATH -source=$PROJECT_PATH/$CLIENT_INTERFACE_PATH/api_client.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/api_client.go -package=mock ApiClient

# моки утилит
$MOCKGEN_PATH -source=$PROJECT_PATH/$UTILS_INTERFACE_PATH/password_encrypter/encrypter.go -destination=$PROJECT_PATH/$UTILS_MOCK_PATH/encrypter.go -package=mock Encrypter
$MOCKGEN_PATH -source=$PROJECT_PATH/$UTILS_INTERFACE_PATH/api_token/tokener.go -destination=$PROJECT_PATH/$UTILS_MOCK_PATH/tokener.go -package=mock Tokener
$MOCKGEN_PATH -source=$PROJECT_PATH/$UTILS_INTERFACE_PATH/time_controller/time_controller.go -destination=$PROJECT_PATH/$UTILS_MOCK_PATH/time_controller.go -package=mock TimeController
$MOCKGEN_PATH -source=$PROJECT_PATH/$UTILS_INTERFACE_PATH/uuid_controller/uuid_controller.go -destination=$PROJECT_PATH/$UTILS_MOCK_PATH/uuid_controller.go -package=mock UUIDController

# вспомогательные
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/server/services/auth/service.go -destination=$PROJECT_PATH/$SERVER_TEST_PATH/mocks/services/auth/repository.go -package=mock Repository 
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/server/services/bucket/service.go -destination=$PROJECT_PATH/$SERVER_TEST_PATH/mocks/services/bucket/service_mocks.go -package=mock Repository ChangeDetector
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/client/core/services/data_service/api/syncer/syncer.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/syncer.go -package=mock Syncer
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/client/core/services/data_service/database/service.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/repository.go -package=mock Repository
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/client/interfaces/core_interface.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/core.go -package=mock Core
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/client/core/services/data_service/api/service.go -destination=$PROJECT_PATH/$CLIENT_MOCK_PATH/change_detector.go -package=mock ChangeDetector


# для клиента либы сервера
$MOCKGEN_PATH -source=$PROJECT_PATH/internal/server/api/client/http_client/http_client.go -destination=$PROJECT_PATH/$SERVER_TEST_PATH/mocks/api/client/client_mocks.go -package=mock Response HTTPClient
