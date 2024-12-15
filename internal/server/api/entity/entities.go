package entity

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/models"
)

type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i LoginInput) Validate(failCode int) *apierrors.Error {
	if i.Login == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidateLoginEmpty.Error())
	}

	if i.Password == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidatePasswordEmpty.Error())
	}

	return nil
}

type RegisterInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i RegisterInput) Validate(failCode int) *apierrors.Error {
	if i.Login == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidateLoginEmpty.Error())
	}

	if i.Password == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidatePasswordEmpty.Error())
	}

	return nil
}

type TokenOutput struct {
	Token string `json:"token"`
}

type GetAccountsOutput struct {
	Accounts []*models.Account `json:"accounts"`
}

type PostAccountInput struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type PutAccountInput struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DeleteAccountInput struct {
	ID int `json:"id"`
}
