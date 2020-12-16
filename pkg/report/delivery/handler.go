package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/domain"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/jwthelper"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/report"
	"github.com/go-park-mail-ru/2020_2_Eternity/pkg/utils"
	"github.com/microcosm-cc/bluemonday"
	"net/http"
	"strconv"
)

type Handler struct {
	uc report.IUsecase
	p  *bluemonday.Policy
}

func New(uc report.IUsecase, p *bluemonday.Policy) *Handler {
	return &Handler{
		uc: uc,
		p:  p,
	}
}

func (h *Handler) ReportPin(c *gin.Context) {
	claimsId, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	r := domain.ReportReq{}
	if err := c.BindJSON(&r); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: "[BindJSON]: " + err.Error()})
		return
	}

	h.sanitize(&r)

	reportId, err := h.uc.ReportPin(claimsId, &r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, utils.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, reportId)
}

func (h *Handler) GetByPinId(c *gin.Context) {
	_, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	id, err := strconv.Atoi(c.Query("pin_id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, utils.Error{Error: err.Error()})
		return
	}

	reports, err := h.uc.GetReportsByPinId(id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, utils.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

func (h *Handler) GetByUsername(c *gin.Context) {
	_, ok := jwthelper.GetClaims(c)
	if !ok {
		c.AbortWithStatusJSON(http.StatusUnauthorized, utils.Error{Error: "invalid token"})
		return
	}

	reports, err := h.uc.GetReportsByUsername(c.Query("username"))

	if err != nil {
		c.AbortWithStatusJSON(http.StatusConflict, utils.Error{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, reports)
}

func (h *Handler) sanitize(r *domain.ReportReq) {
	r.Message = h.p.Sanitize(r.Message)
}
