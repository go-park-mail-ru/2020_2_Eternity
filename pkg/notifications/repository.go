package notifications

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IRepository interface {
	StoreNote(n *domain.Notification) error
	GetNoteById(noteId int) (domain.Notification, error)
	GetNotesToUser(userId int) ([]domain.Notification, error)
}
