package http

import (
	"context"
	"encoding/json"
	"finance/internal/delivery/http/models"
	"finance/pkg/logger"
	"github.com/sirupsen/logrus"
	"net/http"
)

type BaseHandler struct {
	Logger logger.Logger
}

func NewHandler(l logger.Logger) *BaseHandler {
	h := &BaseHandler{
		Logger: l,
	}

	return h
}

func (h *BaseHandler) WriteErrorResponse(ctx context.Context, w http.ResponseWriter, error error, isShown bool) {
	h.Logger.WithFields(logrus.Fields{
		"error":   error,
		"context": ctx,
	}).Errorf("response.error")

	w.WriteHeader(http.StatusBadRequest)

	resp := models.BaseResponse{}
	if error == nil || !isShown {
		resp = models.BaseResponse{
			Error:   "Something went wrong",
			Success: false,
		}
	} else {
		resp = models.BaseResponse{
			Error:   error.Error(),
			Success: false,
		}
	}

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}

func WriteResponse(w http.ResponseWriter, resp interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(resp)
	if err != nil {
		return
	}
}
