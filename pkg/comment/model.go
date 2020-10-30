package comment

import (
	"errors"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"strconv"
)

type Comment struct {
	Id      int
	Path    []int32
	Content string
	PinId   int
	UserId  int
}

type RepoComment interface {
	CreateChildComment(c *Comment, parentId int) error
	CreateRootComment(c *Comment) error
	GetComment(id int) (Comment, error)
	GetAllComments(pinId int) ([]Comment, error)
}

type RepoCommentInstance struct{}

func (rc *RepoCommentInstance) CreateChildComment(c *Comment, parentId int) error {
	err := config.Db.QueryRow(
		"insert into comments (path, content, pin_id, user_id) "+
			"values("+
			"(select path from comments where id = $1) || (select currval('comments_id_seq')::integer), "+
			"$2, $3, $4 "+
			") returning id, path",
		parentId, c.Content, c.PinId, c.UserId).Scan(&c.Id, &c.Path)

	if len(c.Path) < 2 {
		if _, err := config.Db.Exec("delete from comments where id = $1", c.Id); err != nil {
			config.Lg("comment", "comment.CreateComment").
				Error("Can't delete wrongly created comment")

			return errors.New("Can't delete wrongly created comment")
		}

		config.Lg("comment", "comment.CreateComment").
			Error("Given parent id not found in table")
		return errors.New("Given parent id not found in table")
	}

	if err != nil {
		config.Lg("comment", "comment.CreateChildComment").Error(err.Error())
		return err
	}

	return nil
}

func (rc *RepoCommentInstance) CreateRootComment(c *Comment) error {
	err := config.Db.QueryRow(
		"insert into comments (path, content, pin_id, user_id) "+
			"values("+
			"ARRAY(select currval('comments_id_seq')::integer), "+
			"$1, $2, $3 "+
			") returning id, path",
		c.Content, c.PinId, c.UserId).Scan(&c.Id, &c.Path)

	if err != nil {
		config.Lg("comment", "comment.CreateRootComment").Error(err.Error())
		return err
	}

	return nil
}

func (rc *RepoCommentInstance) GetComment(id int) (Comment, error) {
	c := Comment{}
	err := config.Db.QueryRow(
		"select id, path, content, pin_id, user_id "+
			"from comments "+
			"where id = $1",
		id).Scan(&c.Id, &c.Path, &c.Content, &c.PinId, &c.UserId)

	if err != nil {
		config.Lg("comment", "comment.GetComment").Error(err.Error())
		return Comment{}, err
	}

	return c, nil
}

func (rc *RepoCommentInstance) GetAllComments(pinId int) ([]Comment, error) {
	rows, err := config.Db.Query(
		"select id, path, content, pin_id, user_id "+
			"from comments "+
			"where pin_id = $1 "+
			"order by path",
		pinId)

	if err != nil {
		config.Lg("comment", "comment.GetAllComments").Error(err.Error())
		return nil, err
	}

	defer rows.Close()

	comments := []Comment{}
	for rows.Next() {
		c := Comment{}
		err := rows.Scan(&c.Id, &c.Path, &c.Content, &c.PinId, &c.UserId)

		if err != nil {
			config.Lg("comment", "comment.GetAllComments").Error(err.Error())
			return nil, err
		}

		comments = append(comments, c)
	}

	if rows.Err() != nil {
		config.Lg("comment", "comment.GetAllComments").Error(rows.Err())
		return nil, rows.Err()
	}

	// TODO (Pavel S) Is it an error?
	if len(comments) == 0 {
		config.Lg("comment", "comment.GetAllComments").
			Error("Comments not found for given id ", pinId)
		return nil, errors.New("Comments not found for given id " + strconv.Itoa(pinId))
	}

	return comments, nil
}
