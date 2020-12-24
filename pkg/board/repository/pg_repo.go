package repository

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/api"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}

func (r *Repository) CreateBoard(userId int, b *api.CreateBoard) (*domain.Board, error) {
	rb := &domain.Board{}
	if err := r.dbConn.QueryRow("insert into boards(title, content, user_id) values($1, $2, $3) returning id", b.Title, b.Content, userId).Scan(&rb.ID); err != nil {
		config.Lg("board", "CreateBoard").Error(err.Error())
		return rb, errors.New("bad board")
	}
	rb.Title = b.Title
	rb.Content = b.Content
	if err := r.dbConn.QueryRow("select username from users where id = $1", userId).Scan(&rb.Username); err != nil {
		config.Lg("board", "CreateBoard").Error(err.Error())
		return rb, errors.New("bad uid")
	}
	return rb, nil
}
func (r *Repository) GetBoard(id int) (*domain.Board, error) {
	b := &domain.Board{
		ID: id,
	}
	if err := r.dbConn.QueryRow("select title, content, username from boards join users "+
		"on users.id = boards.user_id where boards.id = $1", b.ID).Scan(&b.Title, &b.Content, &b.Username); err != nil {
		config.Lg("board", "GetBoard").Error(err.Error())
		return b, errors.New("bad id")
	}
	return b, nil
}

func (r *Repository) GetAllBoardsByUser(username string) ([]domain.Board, error) {
	var boards []domain.Board

	rows, err := r.dbConn.Query("select boards.id, title, content from boards join users on users.id = "+
		"boards.user_id where lower(username) = lower($1)", username)
	if err != nil {
		config.Lg("board", "GetAllBoardsByUserId").Error(err.Error())
		return boards, errors.New("bad id")
	}
	defer rows.Close()
	for rows.Next() {
		b := domain.Board{
			Username: username,
		}
		if err := rows.Scan(&b.ID, &b.Title, &b.Content); err != nil {
			config.Lg("board", "GetAllBoardsByUserId").Error(err.Error())
			return nil, err
		}
		boards = append(boards, b)
	}
	return boards, nil
}

func (r *Repository) GetBoardsPinNotAttach(userId, pinId int) ([]domain.Board, error) {
	var boards []domain.Board

	rows, err := r.dbConn.Query("select id, title from boards where boards.id not in (select id from boards "+
		"join boards_pins on boards.id = boards_pins.board_id where pin_id = $1) and user_id = $2", pinId, userId)
	if err != nil {
		config.Lg("board", "GetAllBoardsByUserId").Error(err.Error())
		return boards, errors.New("bad id")
	}
	defer rows.Close()
	for rows.Next() {
		b := domain.Board{}
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			config.Lg("board", "GetAllBoardsByUserId").Error(err.Error())
			return nil, err
		}
		boards = append(boards, b)
	}
	return boards, nil
}

func (r *Repository) CheckOwner(userId int, boardId int) error {
	var owner int
	if err := r.dbConn.QueryRow("select user_id from boards where id = $1", boardId).Scan(&owner); err != nil {
		config.Lg("board", "CheckOwner").Error(err.Error())
		return err
	}

	if owner != userId {
		err := errors.New("not an owner")
		config.Lg("board", "CheckOwner").Error(err.Error())
		return err
	}
	return nil
}

func (r *Repository) AttachPin(boardId int, pinId int) error {
	if _, err := r.dbConn.Exec("insert into boards_pins(board_id, pin_id) values($1, $2)", boardId, pinId); err != nil {
		config.Lg("board", "AttachPin").Error(err.Error())
		return err
	}
	return nil
}

func (r *Repository) DetachPin(boardId int, pinId int) error {
	if _, err := r.dbConn.Exec("delete from boards_pins where board_id = $1 and pin_id = $2", boardId, pinId); err != nil {
		config.Lg("board", "DetachPin").Error(err.Error())
		return err
	}
	return nil
}
