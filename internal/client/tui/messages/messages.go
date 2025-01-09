package messages

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

type SignLoginMsg struct{}
type SignRegistrationMsg struct{}
type LoginSuccessMsg struct{}
type RegistrationSuccessMsg struct{}
type SignBackMsg struct{}

type SelectAccountMsg struct{}
type SelectNoteMsg struct{}
type SelectCardMsg struct{}
type SectionBackMsg struct{}

type AccountListBackMsg struct{}
type AccountEditMsg struct{ models.Account }
type AccountDeleteMsg struct{ models.Account }

type NoteListBackMsg struct{}
type NoteEditMsg struct{ models.Note }
type NoteDeleteMsg struct{ models.Note }

type CardListBackMsg struct{}
type CardEditMsg struct{ models.Card }
type CardDeleteMsg struct{ models.Card }

type ButtonSubmitMsg struct{}
type ButtonBackMsg struct{}
