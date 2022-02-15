package reportshandler

import (
	"net/http"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/bhankey/BD_lab/backend/internal/service/reportservice"
	"github.com/go-chi/chi/v5"
)

type ReportsHandler struct {
	Router chi.Router
	*deliveryhttp.BaseHandler

	reportService *reportservice.ReportService
}

func NewReportsHandler(
	baseHandler *deliveryhttp.BaseHandler,
	reportsService *reportservice.ReportService) *ReportsHandler {
	router := chi.NewRouter()

	accountHandler := &ReportsHandler{
		Router:        router,
		BaseHandler:   baseHandler,
		reportService: reportsService,
	}

	router.Method(http.MethodPost, "/turnover_sheets", accountHandler.getTurnOverReport())
	router.Method(http.MethodPost, "/debtors", accountHandler.getDebtors())

	return accountHandler
}
