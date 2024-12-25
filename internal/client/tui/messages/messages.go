package messages

import "github.com/besean163/gophkeeper/internal/client/core/models"

type SignLoginMsg struct{}
type SignRegistrationMsg struct{}
type LoginSuccessMsg struct{}
type SignBackMsg struct{}

type SelectAccountMsg struct{}
type SelectNoteMsg struct{}
type SelectCardMsg struct{}
type SectionBackMsg struct{}

type AccountListBackMsg struct{}
type AccountEditMsg struct{ models.Account }
type AccountDeleteMsg struct{ models.Account }

type ButtonSubmitMsg struct{}
type ButtonBackMsg struct{}
