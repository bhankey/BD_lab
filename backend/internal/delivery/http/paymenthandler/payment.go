package paymentshandler

import (
	"context"
	"encoding/json"
	"net/http"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/models"
	"github.com/go-openapi/strfmt"
)

func (s *PaymentHandler) create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		defer func() { _ = r.Body.Close() }()
		var req models.Payment

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

		if err := s.paymentsService.Create(ctx, int(req.AccountID), req.Sum, req.Reason); err != nil {
			s.WriteErrorResponse(ctx, w, err, false)

			return
		}

		deliveryhttp.WriteResponse(w, models.BaseResponse{
			Error:   "",
			Success: true,
		})
	}
}

func (s *PaymentHandler) getAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		payments, err := s.paymentsService.GetAll(ctx)
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, false)

			return
		}

		response := make(models.PaymentsGetAllResponse, 0, len(payments))
		for _, payment := range payments {
			response = append(response, &models.Payment{
				AccountID: int64(payment.AccountID),
				Reason:    payment.Reason,
				Sum:       payment.Sum,
			})
		}

		deliveryhttp.WriteResponse(w, response)
	}
}

func (s *PaymentHandler) getClientPayments() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()

		defer func() { _ = r.Body.Close() }()
		var req models.GetClientPaymentsRequest

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

		payments, err := s.paymentsService.GetClientPayments(ctx, int(req.AccountID))
		if err != nil {
			s.WriteErrorResponse(ctx, w, err, false)

			return
		}

		response := make(models.PaymentsGetAllResponse, 0, len(payments))
		for _, payment := range payments {
			response = append(response, &models.Payment{
				AccountID: int64(payment.AccountID),
				Reason:    payment.Reason,
				Sum:       payment.Sum,
			})
		}

		deliveryhttp.WriteResponse(w, response)
	}
}
