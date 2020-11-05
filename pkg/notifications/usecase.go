package notifications

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IUsecase interface {
	CreateNotes(iNote interface{}) error
	GetUserNotes(userId int) ([]domain.NoteResp, error)
	UpdateNote(noteId int) error
	UpdateUserNotes(userId int) error
}
