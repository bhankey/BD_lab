package reportshandler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/models"
	"github.com/bhankey/BD_lab/backend/internal/delivery/mappers"
	"github.com/go-openapi/strfmt"
)

func (s *ReportsHandler) getTurnOverReport() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		defer func() { _ = r.Body.Close() }()
		var req models.ReportTurnOverSheetsRequest

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&req)
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, true)

			return
		}

		if err := req.Validate(strfmt.NewFormats()); err != nil {
			s.WriteErrorResponse(ctx, w, err, true)

			return
		}

		accountIDs := mappers.SliceInt64ToInt(req.AccountIds)
		reports, err := s.reportService.GetTurnOverSheets(ctx, accountIDs, int(req.Year))
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, false)

			return
		}

		resp := make([]*models.ReportTurnOverSheetsResponseItem, 0, len(reports))
		for _, report := range reports {
			details := make([]*models.ReportTurnOverByMonth, 0, len(report.MothDetails))
			for i := time.January; i <= time.December; i++ {
				details = append(details, &models.ReportTurnOverByMonth{
					Income: report.MothDetails[i].Income,
					Outgo:  report.MothDetails[i].Outgo,
					Sum:    report.MothDetails[i].Sum,
				})
			}

			resp = append(resp, &models.ReportTurnOverSheetsResponseItem{
				AccountID:   int64(report.AccountID),
				EndSum:      report.EndSum,
				MonthReport: details,
				StartSum:    report.StartingSum,
			})
		}

		deliveryhttp.WriteResponse(w, resp)
	}
}

func (s *ReportsHandler) getDebtors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		defer func() { _ = r.Body.Close() }()
		var req models.ReportDebtorsRequest

		decoder := json.NewDecoder(r.Body)

		err := decoder.Decode(&req)
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, true)

			return
		}

		if err := req.Validate(strfmt.NewFormats()); err != nil {
			s.WriteErrorResponse(ctx, w, err, true)

			return
		}

		accountIDs := mappers.SliceInt64ToInt(req.AccountIds)
		debtors, err := s.reportService.GetDebtors(ctx, accountIDs)
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, false)

			return
		}

		resp := make([]models.ReportDebtorsResponseItems0, 0, len(debtors))
		for _, debtor := range debtors {
			resp = append(resp, models.ReportDebtorsResponseItems0{
				AccountID: int64(debtor.AccountID),
				DebtSum:   debtor.Own,
				Income:    debtor.Income,
				Outgo:     debtor.Outgo,
			})
		}

		deliveryhttp.WriteResponse(w, resp)
	}
}
