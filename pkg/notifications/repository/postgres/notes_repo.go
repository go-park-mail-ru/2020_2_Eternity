package postgres

import (
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
		"INSERT INTO notifications "+
			"(to_user_id, type, encoded_data, creation_time, is_read) "+
			"VALUES ($1, $2, $3, $4, $5) "+
			"RETURNING id, creation_time, is_read",
		n.ToUserId, n.Type, n.EncodedData, time.Now(), false).Scan(&n.Id, &n.CreationTime, &n.IsRead)

	if err != nil {
		config.Lg("notes_repo", "StoreNote").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetNoteById(noteId int) (domain.Notification, error) {
	n := domain.Notification{}
	err := r.dbConn.QueryRow(
		"SELECT id, to_user_id, type, encoded_data, creation_time, is_read "+
			"FROM notifications "+
			"WHERE id = $1",
		noteId).Scan(&n.Id, &n.ToUserId, &n.Type, &n.EncodedData, &n.CreationTime, &n.IsRead)

	if err != nil {
		config.Lg("notes_repo", "GetNoteById").Error(err.Error())
		return domain.Notification{}, err
	}

	return n, nil
}

func (r *Repository) GetNotesToUser(userId int) ([]domain.Notification, error) {
	rows, err := r.dbConn.Query(
		"SELECT id, to_user_id, type, encoded_data, creation_time, is_read "+
			"FROM notifications "+
			"WHERE to_user_id = $1 "+
			"ORDER BY creation_time ",
		userId)

	if err != nil {
		config.Lg("notes_repo", "GetNotesToUser").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	notes := []domain.Notification{}
	for rows.Next() {
		n := domain.Notification{}
		err := rows.Scan(&n.Id, &n.ToUserId, &n.Type, &n.EncodedData, &n.CreationTime, &n.IsRead)

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

func (r *Repository) UpdateNoteIsRead(noteId int) error {
	_, err := r.dbConn.Exec(
		"UPDATE notifications "+
			"SET is_read = true "+
			"WHERE id = $1 ",
		noteId)

	if err != nil {
		config.Lg("notes_repo", " UpdateNoteIsRead").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) UpdateUserNotes(userId int) error {
	_, err := r.dbConn.Exec(
		"UPDATE notifications "+
			"SET is_read = true "+
			"WHERE to_user_id = $1 ",
		userId)

	if err != nil {
		config.Lg("notes_repo", " UpdateUserNotes").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetUserNotesAmount(userId int) (int, error) {
	n := 0
	err := r.dbConn.QueryRow(
		"SELECT COUNT(1) FROM notifications " +
			"WHERE u.to_user_id = $1 AND is_read = false ",
		userId).Scan(&n)

	if err != nil {
		config.Lg("notes_repo", "GetUserNotesAmount").Error(err.Error())
		return 0, err
	}

	return n, nil
}

