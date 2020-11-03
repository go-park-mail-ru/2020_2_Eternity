package postgres

import (
	"context"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"time"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}

func (r *Repository) StoreNote(n *domain.Notification) error {
	err := r.dbConn.QueryRow(
		context.Background(),
		"INSERT INTO notifications " +
			"(from_user_id, to_user_id, type, encoded_data, creation_time) " +
			"VALUES ($1, $2, $3, $4, $5) " +
			"RETURNING id, creation_time",
		n.FromUserId, n.ToUserId, n.Type, n.EncodedData, time.Now()).Scan(&n.Id, &n.CreationTime)

	if err != nil {
		config.Lg("notes_repo", "StoreNote").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetNoteById(noteId int) (domain.Notification, error) {
	n := domain.Notification{}
	err := r.dbConn.QueryRow(
		context.Background(),
		"SELECT id, from_user_id, to_user_id, type, encoded_data, creation_time " +
			"FROM notifications " +
			"WHERE id = $1",
		noteId).Scan(&n.Id, &n.FromUserId, &n.ToUserId, &n.Type, &n.EncodedData, &n.CreationTime)

	if err != nil {
		config.Lg("notes_repo", "GetNoteById").Error(err.Error())
		return domain.Notification{}, err
	}

	return n, nil
}


func (r *Repository) GetNotesToUser(userId int) ([]domain.Notification, error) {
	rows, err := r.dbConn.Query(
		context.Background(),
		"SELECT id, from_user_id, to_user_id, type, encoded_data, creation_time " +
			"FROM notifications " +
			"WHERE to_user_id = $1",
		userId)

	if err != nil {
		config.Lg("notes_repo", "GetNotesToUser").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	notes := []domain.Notification{}
	for rows.Next() {
		n := domain.Notification{}
		err := rows.Scan(&n.Id, &n.FromUserId, &n.ToUserId, &n.Type, &n.EncodedData, &n.CreationTime)

		if err != nil {
			config.Lg("notes_repo", "GetNotesToUser").Error(err.Error())
			return nil, err
		}

		notes = append(notes, n)
	}

	if rows.Err() != nil {
		config.Lg("notes_repo", "GetNotesToUser").Error(rows.Err())
		return nil, rows.Err()
	}


	return notes, nil
}
