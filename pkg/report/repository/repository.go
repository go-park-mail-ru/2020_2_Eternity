package repository

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/configs/config"
	"github.com/go-park-mail-ru/2020_2_Eternity/internal/app/database"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
)

type Repository struct {
	db database.IDbConn
}

func New(db database.IDbConn) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) ReportPin(userId int, rep *domain.ReportReq) (int, error) {
	var repId int
	if err := r.db.QueryRow("insert into reports (pin_id, user_id, message, owner, type) values ($1, $2, $3, "+
		"(select username from users join pins on pins.user_id = users.id where pins.id = $1), $4) "+
		"returning reports.id", rep.PinId, userId, rep.Message, rep.Type).Scan(&repId); err != nil {
		config.Lg("report", "ReportPin").Error(err.Error())
		return 0, err
	}
	return repId, nil
}

func (r *Repository) GetReportsByPinId(pinId int) ([]domain.Report, error) {
	rows, err := r.db.Query("select id, pin_id, user_id, message, owner, type from reports where pin_id = $1", pinId)
	if err != nil {
		config.Lg("report", "GetReportByPinId").Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	reports := make([]domain.Report, 0)

	for rows.Next() {
		rep := domain.Report{}
		if err := rows.Scan(&rep.Id, &rep.PinId, &rep.OwnerId, &rep.Message, &rep.PinOwner, &rep.Type); err != nil {
			config.Lg("report", "GetReportIdScan").Error(err.Error())
			return reports, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}

func (r *Repository) GetReportsByUsername(username string) ([]domain.Report, error) {
	rows, err := r.db.Query("select id, pin_id, user_id, message, owner, type from reports where owner = $1", username)
	if err != nil {
		config.Lg("report", "GetReportByPinId").Error(err.Error())
		return nil, err
	}
	defer rows.Close()

	reports := make([]domain.Report, 0)

	for rows.Next() {
		rep := domain.Report{}
		if err := rows.Scan(&rep.Id, &rep.PinId, &rep.OwnerId, &rep.Message, &rep.PinOwner, &rep.Type); err != nil {
			config.Lg("report", "GetReportIdScan").Error(err.Error())
			return reports, err
		}
		reports = append(reports, rep)
	}
	return reports, nil
}
