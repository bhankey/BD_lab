package accounthandler

import (
	"net/http"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	accountservise "github.com/bhankey/BD_lab/backend/internal/service/accountservice"
	"github.com/go-chi/chi/v5"
)

type AccountHandler struct {
	Router chi.Router
	*deliveryhttp.BaseHandler

	accountService *accountservise.AccountService
}

func NewAccountHandler(
	baseHandler *deliveryhttp.BaseHandler,
	accountService *accountservise.AccountService) *AccountHandler {
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
