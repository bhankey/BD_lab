package accounthandler

import (
	deliveryhttp "finance/internal/delivery/http"
	accountservise "finance/internal/service/accountservice"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type AccountHandler struct {
	Router chi.Router
	*deliveryhttp.BaseHandler

	accountService *accountservise.AccountService
}

func NewAccountHandler(baseHandler *deliveryhttp.BaseHandler, accountService *accountservise.AccountService) *AccountHandler {
	router := chi.NewRouter()

	accountHandler := &AccountHandler{
		Router:         router,
		BaseHandler:    baseHandler,
		accountService: accountService,
	}

	router.Method(http.MethodPost, "/create", accountHandler.create())
	router.Method(http.MethodGet, "/get_all", accountHandler.getAll())
	router.Method(http.MethodPost, "/update", accountHandler.update())
	router.Method(http.MethodDelete, "/delete", accountHandler.delete())

	return accountHandler
}
