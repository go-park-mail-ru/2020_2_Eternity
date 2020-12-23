package postgres

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

var (
	parentId int = 0

	childPath = []int{parentId, 1}

	childComment = domain.Comment{
		Id: 1,
		Content: "content",
		PinId: 2,
		UserId: 3,
		Username: "username",
	}
)

func TestMain(m *testing.M) {
	config.Conf = config.NewConfigTst()
	code := m.Run()
	os.Exit(code)
}


func TestStoreChildComment(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Query error

	c := childComment
	mock.ExpectQuery("insert into comments").
		WithArgs(parentId, c.Content, c.PinId, c.UserId).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "path", "user_id"}))


	r := NewRepo(db)
	err := r.StoreChildComment(&c, int(parentId))

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}

func TestStoreRootComment(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Query error

	c := childComment
	mock.ExpectQuery("insert into comments").
		WithArgs(c.Content, c.PinId, c.UserId).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "path", "user_id"}))


	r := NewRepo(db)
	err := r.StoreRootComment(&c)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}


func TestGetComment(t *testing.T) {
	db, mock, e := sqlmock.New()
	if e != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", e)
	}
	defer db.Close()

	// Error

	c := childComment
	mock.ExpectQuery("select").
		WithArgs(c.Id).
		WillReturnRows(
			sqlmock.NewRows([]string{"id", "path", "content", "pin_id", "user_id", "username"}))


	r := NewRepo(db)
	_, err := r.GetComment(c.Id)

	assert.NotNil(t, err)
	assert.Nil(t, mock.ExpectationsWereMet())
}


//var db *pgxpool.Pool
//
//func TestMain(m *testing.M) {
//	config.Conf = config.NewConfig()
//	conf, err := pgxpool.ParseConfig(fmt.Sprintf(
//		"user=%s password=%s host=%s dbname=%s sslmode=%s pool_max_conns=%s",
//		config.Conf.Db.Postgres.Username,
//		config.Conf.Db.Postgres.Password,
//		config.Conf.Db.Postgres.Host,
//		config.Conf.Db.Postgres.DbName,
//		config.Conf.Db.Postgres.SslMode,
//		config.Conf.Db.Postgres.MaxConn,
//	))
//	if err != nil {
//		fmt.Println("Error ", err.Error())
//	}
//
//	db, err = pgxpool.ConnectConfig(context.Background(), conf)
//
//	code := m.Run()
//
//	db.Close()
//	os.Exit(code)
//}


//func TestStore(t *testing.T) {
//	r := NewRepo(db)
//
//	c := domain.Comment{
//		Content: "conte",
//		PinId: 4,
//		UserId: 2,
//	}
//
//	assert.Nil(t, r.StoreChildComment(&c, 42))
//	fmt.Println(c)
//
//}

//func TestStoreRoot(t *testing.T) {
//	r := NewRepo(db)
//
//	c := domain.Comment{
//		Content: "content",
//		PinId: 4,
//		UserId: 2,
//	}
//
//	assert.Nil(t, r.StoreRootComment(&c))
//	fmt.Println(c)
//}


//func TestGetId(t *testing.T) {
//	r := NewRepo(db)
//
//
//	c, err := r.GetComment(53)
//	assert.Nil(t, err)
//	fmt.Println(c)
//}

//func TestPinComments(t *testing.T) {
//	r := NewRepo(db)
//
//
//	c, err := r.GetPinComments(311)
//	assert.Nil(t, err)
//	fmt.Println(c)
//}

