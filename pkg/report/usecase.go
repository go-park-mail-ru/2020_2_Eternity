package report

import "github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"

type IUsecase interface {
	ReportPin(userId int, rep *domain.ReportReq) (int, error)
	GetReportsByPinId(pinId int) ([]domain.Report, error)
	GetReportsByUsername(username string) ([]domain.Report, error)
}
