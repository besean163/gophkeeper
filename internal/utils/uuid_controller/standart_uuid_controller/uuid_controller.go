package standartuuidcontroller

import "github.com/google/uuid"

type UUIDController struct{}

func NewUUIDController() UUIDController {
	return UUIDController{}
}

func (c UUIDController) GetUUID() string {
	return uuid.NewString()
}
