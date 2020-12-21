package usecase

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	domainWs "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain/ws"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/microservices/chat"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/notifications"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/pin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/user"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/ws"
)

const (
	NoteComment = iota
	NotePin
	NoteFollow
	NoteChat
	NoteMessage
)

type Usecase struct {
	noteRepo  notifications.IRepository
	pinRepo pin.IRepository
	userRepo user.IRepository
	chatRepo chat.IRepository
	wsServer ws.IServer
}

func NewUsecase(nr notifications.IRepository, pr pin.IRepository, ur user.IRepository, cr chat.IRepository, server ws.IServer) *Usecase {
	return &Usecase{
		noteRepo:  nr,
		pinRepo: pr,
		userRepo: ur,
		chatRepo: cr,
		wsServer: server,
	}
}

//func typeStrToInt(msgType string) (int, error) {
//
//}

func typeIntToStr(noteType int) (string, error) {
	noteTypeWs := ""
	switch noteType {
	case NoteComment:
		noteTypeWs = domain.NoteCommentRespType
	case NotePin:
		noteTypeWs = domain.NotePinRespType
	case NoteFollow:
		noteTypeWs = domain.NoteFollowRespType
	case NoteChat:
		noteTypeWs = domain.NoteChatRespType
	case NoteMessage:
		noteTypeWs = domain.NoteMessageRespType
	default:
		return "", errors.New(fmt.Sprintf("switch: Can't find needed note type, given: %d", noteType))
	}

	return noteTypeWs, nil
}

func (uc *Usecase) sendNotes(n interface{}, noteType int, toUsers []int)  {
	data, err := json.Marshal(n)
	if err != nil {
		config.Lg("notifications_usecase", "sendNotes").
			Error("Marshal: ", err.Error())
		return
	}

	noteTypeWs, err := typeIntToStr(noteType)
	if err != nil {
		config.Lg("notifications_usecase", "sendNotes").
			Error("typeIntToStr: ", err.Error())
		return
	}

	for _, id := range toUsers {
		wsNote := domainWs.MessageResp{
			UserId: id,
			Type: noteTypeWs,
			Status: 200,
			Data: data,
		}

		uc.wsServer.SendMessage(&wsNote)
	}
}

func (uc *Usecase) getDstNoteComment (n *domain.NoteComment) ([]int, error) {
	p, err := uc.pinRepo.GetPin(n.PinId)
	if err != nil {
		config.Lg("notifications_usecase", "getDstNoteComment").Error(err.Error())
		return nil, err
	}
	toUsers := []int{p.UserId}

	return toUsers, nil
}

func (uc *Usecase) getDstNoteChat (n *domain.NoteChat) ([]int, error) {
	collocId, err := uc.chatRepo.GetCollocutorId(n.CreatorId, n.Id)
	if err != nil {
		config.Lg("notifications_usecase", "getDstNoteChat").Error(err.Error())
		return nil, err
	}

	toUsers := []int{collocId}

	return toUsers, nil
}

func (uc *Usecase) getDstNoteMessage(n *domain.NoteMessage) ([]int, error) {
	collocId, err := uc.chatRepo.GetCollocutorId(n.CreatorId, n.ChatId)
	if err != nil {
		config.Lg("notifications_usecase", "getDstNoteMessage").Error(err.Error())
		return nil, err
	}

	toUsers := []int{collocId}

	return toUsers, nil
}

func (uc *Usecase) getDstNotePin(p *domain.NotePin) ([]int, error) {
	toUsers, err := uc.userRepo.GetFollowersIds(p.UserId)
	if err != nil {
		config.Lg("notifications_usecase", "getDstNotePin").Error(err.Error())
		return nil, err
	}

	return toUsers, nil
}

func (uc *Usecase) getDstNoteFollow(f *domain.NoteFollow) ([]int, error) {
	return []int{f.UserId}, nil
}

func (uc *Usecase) CreateNotes(iNote interface{}) error {
	var noteType int
	var toUsers []int
	var err error

	switch note := iNote.(type) {
	case *domain.NoteComment:
		noteType = NoteComment
		note.Username = uc.userRepo.GetUsername(note.UserId)
		toUsers, err = uc.getDstNoteComment(note)
		uc.sendNotes(note, noteType, toUsers)
	case *domain.NotePin:
		noteType = NotePin
		note.Username = uc.userRepo.GetUsername(note.UserId)
		toUsers, err = uc.getDstNotePin(note)
		uc.sendNotes(note, noteType, toUsers)
	case *domain.NoteFollow:
		noteType = NoteFollow
		note.Username = uc.userRepo.GetUsername(note.UserId)
		toUsers, err = uc.getDstNoteFollow(note)
		uc.sendNotes(note, noteType, toUsers)
	case *domain.NoteChat:
		noteType = NoteChat
		toUsers, err = uc.getDstNoteChat(note)
		uc.sendNotes(note, noteType, toUsers)
	case *domain.NoteMessage:
		noteType = NoteMessage
		toUsers, err = uc.getDstNoteMessage(note)
		uc.sendNotes(note, noteType, toUsers)
	default:
		config.Lg("notifications_usecase", "CreateNote").Error("Unknown notification type")
		return errors.New("Unknown notification type")
	}

	if err != nil {
		config.Lg("notifications_usecase", "CreateNote").Error(err.Error())
		return err
	}

	encoded, err := json.Marshal(iNote)
	if err != nil {
		config.Lg("notifications_usecase", "CreateNote").Error(err.Error())
		return err
	}


	note := domain.Notification{
		Type: noteType,
		EncodedData: encoded,
	}

	for _, id := range toUsers {
		note.ToUserId = id
		err = uc.noteRepo.StoreNote(&note)

		if err != nil {
			config.Lg("notifications_usecase", "CreateNote").Error(err.Error())
			return err
		}
	}

	return nil
}


func (uc *Usecase) GetUserNotes(userId int) ([]domain.NoteResp, error) {
	modelNotes, err :=  uc.noteRepo.GetNotesToUser(userId)

	respNotes := []domain.NoteResp{}
	for _, mNote := range modelNotes {
		respNotes = append(respNotes, domain.NoteResp{
			Id: mNote.Id,
			Type: mNote.Type,
			EncodedData: mNote.EncodedData,
			CreationTime: mNote.CreationTime,
			IsRead: mNote.IsRead,
		})
	}

	return respNotes, err
}

func (uc *Usecase) UpdateNote(noteId int) error {
	return uc.noteRepo.UpdateNoteIsRead(noteId)
}

func (uc *Usecase) UpdateUserNotes(userId int) error {
	return uc.noteRepo.UpdateUserNotes(userId)
}
