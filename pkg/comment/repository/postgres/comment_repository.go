package postgres

import (
	"context"
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"strconv"
)

type Repository struct {
	dbConn database.IDbConn
}

func NewRepo(d database.IDbConn) *Repository {
	return &Repository{
		dbConn: d,
	}
}


func (r *Repository) StoreChildComment(c *domain.Comment, parentId int) error {
	// TODO (Pavel S) query should check if parent comment exists
	err := r.dbConn.QueryRow(
		context.Background(),
		"insert into comments (path, content, pin_id, user_id) "+
			"values("+
			"(select path from comments where id = $1) || (select currval('comments_id_seq')::integer), "+
			"$2, $3, $4 "+
			") returning id, path",
		parentId, c.Content, c.PinId, c.UserId).Scan(&c.Id, &c.Path)

	if len(c.Path) < 2 {
		if _, err := r.dbConn.Exec(context.Background(),"delete from comments where id = $1", c.Id); err != nil {
			config.Lg("comment_postgres", "StoreChildComment").
				Error("Can't delete wrongly created comment")

			return errors.New("Can't delete wrongly created comment")
		}

		config.Lg("comment_postgres", "StoreChildComment").
			Error("Given parent id not found in table")
		return errors.New("Given parent id not found in table")
	}

	if err != nil {
		config.Lg("comment_postgres", "StoreChildComment").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) StoreRootComment(c *domain.Comment) error {
	err := r.dbConn.QueryRow(
		context.Background(),
		"insert into comments (path, content, pin_id, user_id) "+
			"values("+
			"ARRAY(select currval('comments_id_seq')::integer), "+
			"$1, $2, $3 "+
			") returning id, path",
		c.Content, c.PinId, c.UserId).Scan(&c.Id, &c.Path)

	if err != nil {
		config.Lg("comment_postgres", "StoreRootComment").Error(err.Error())
		return err
	}

	return nil
}

func (r *Repository) GetComment(commentId int) (domain.Comment, error) {
	c := domain.Comment{}
	err := r.dbConn.QueryRow(
		context.Background(),
		"select id, path, content, pin_id, user_id "+
			"from comments "+
			"where id = $1",
		commentId).Scan(&c.Id, &c.Path, &c.Content, &c.PinId, &c.UserId)

	if err != nil {
		config.Lg("comment_postgres", "GetComment").Error(err.Error())
		return domain.Comment{}, err
	}

	return c, nil
}

func (r *Repository) GetPinComments(pinId int) ([]domain.Comment, error) {
	rows, err := r.dbConn.Query(
		context.Background(),
		"select id, path, content, pin_id, user_id "+
			"from comments "+
			"where pin_id = $1 "+
			"order by path",
		pinId)

	if err != nil {
		config.Lg("comment_postgres", "GetPinComments").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	comments := []domain.Comment{}
	for rows.Next() {
		c := domain.Comment{}
		err := rows.Scan(&c.Id, &c.Path, &c.Content, &c.PinId, &c.UserId)

		if err != nil {
			config.Lg("comment_postgres", "GetPinComments").Error(err.Error())
			return nil, err
		}

		comments = append(comments, c)
	}

	if rows.Err() != nil {
		config.Lg("comment_postgres", "GetPinComments").Error(rows.Err())
		return nil, rows.Err()
	}

	// TODO (Pavel S) Is it an error?
	if len(comments) == 0 {
		config.Lg("comment_postgres", "GetPinComments").
			Error("Comments not found for given id ", pinId)
		return nil, errors.New("Comments not found for given id " + strconv.Itoa(pinId))
	}

	return comments, nil
}
