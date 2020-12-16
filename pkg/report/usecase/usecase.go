package usecase

import (
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/report"
)

type Usecase struct {
	r report.IRepository
}

func New(r report.IRepository) *Usecase {
	return &Usecase{
		r: r,
	}
}

func (uc *Usecase) ReportPin(userId int, rep *domain.ReportReq) (int, error) {
	return uc.r.ReportPin(userId, rep)
}

func (uc *Usecase) GetReportsByPinId(pinId int) ([]domain.Report, error) {
	return uc.r.GetReportsByPinId(pinId)
}

func (uc *Usecase) GetReportsByUsername(username string) ([]domain.Report, error) {
	return uc.r.GetReportsByUsername(username)
}
