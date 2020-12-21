package repository

import (
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/stretchr/testify/assert"
	"testing"
)

var _ = func() bool {
	testing.Init()
	config.Conf = config.NewConfigTst()
	return true
}()

var rep = &domain.ReportReq{
	Message: "spam",
	Type:    1,
	PinId:   2,
}

var repR = &domain.Report{
	Id:       1,
	Message:  "spam",
	OwnerId:  "2",
	Type:     1,
	PinId:    "2",
	PinOwner: "21savage",
}

func TestRepository_ReportPin(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

	mock.ExpectQuery("insert into reports").WithArgs(rep.PinId, 1, rep.Message, rep.Type).WillReturnRows(rows)

	id, err := r.ReportPin(1, rep)
	assert.NoError(t, err)
	assert.Equal(t, 1, id)

	mock.ExpectQuery("insert into reports").WithArgs(rep.PinId, 2, rep.Message, rep.Type).
		WillReturnError(errors.New("not found user"))
	_, err = r.ReportPin(2, rep)
	assert.Error(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetReportsByPinId(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	rows := sqlmock.NewRows([]string{"id", "pin_id", "user_id", "message", "owner", "type"}).
		AddRow(repR.Id, repR.PinId, repR.OwnerId, repR.Message, repR.PinOwner, repR.Type)

	mock.ExpectQuery("select ").WithArgs(2).WillReturnRows(rows)

	reps, err := r.GetReportsByPinId(2)
	assert.NoError(t, err)
	assert.Equal(t, len(reps), 1)

	mock.ExpectQuery("select ").WithArgs(3).WillReturnError(errors.New("not found"))
	_, err = r.GetReportsByPinId(3)
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestRepository_GetReportsByUsername(t *testing.T) {
	t.Helper()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	r := New(db)

	rows := sqlmock.NewRows([]string{"id", "pin_id", "user_id", "message", "owner", "type"}).
		AddRow(repR.Id, repR.PinId, repR.OwnerId, repR.Message, repR.PinOwner, repR.Type)

	mock.ExpectQuery("select ").WithArgs(repR.PinOwner).WillReturnRows(rows)

	reps, err := r.GetReportsByUsername(repR.PinOwner)
	assert.NoError(t, err)
	assert.Equal(t, len(reps), 1)

	mock.ExpectQuery("select ").WithArgs("22sav").WillReturnError(errors.New("not found"))
	_, err = r.GetReportsByUsername("22sav")
	assert.Error(t, err)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
