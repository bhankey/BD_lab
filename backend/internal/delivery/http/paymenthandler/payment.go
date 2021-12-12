package paymentshandler

import (
	"context"
	"encoding/json"
	deliveryhttp "finance/internal/delivery/http"
	"finance/internal/delivery/http/models"
	"github.com/go-openapi/strfmt"
	"net/http"
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
