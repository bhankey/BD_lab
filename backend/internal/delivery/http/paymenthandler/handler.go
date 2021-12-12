package paymentshandler

import (
	deliveryhttp "finance/internal/delivery/http"
	"finance/internal/service/paymentsservice"
	"github.com/go-chi/chi/v5"
)

type PaymentHandler struct {
	Router chi.Router
	*deliveryhttp.BaseHandler

	paymentsService *paymentsservice.PaymentsService
}

func NewPaymentHandler(baseHandler *deliveryhttp.BaseHandler, paymentService *paymentsservice.PaymentsService) *PaymentHandler {
	router := chi.NewRouter()

	accountHandler := &PaymentHandler{
		Router:          router,
		BaseHandler:     baseHandler,
		paymentsService: paymentService,
	}

	router.Post("/create", accountHandler.create())

	return accountHandler
}
