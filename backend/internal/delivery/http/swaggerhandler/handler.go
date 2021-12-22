package swaggerhandler

import (
	deliveryhttp "finance/internal/delivery/http"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type SwaggerHandler struct {
	Router chi.Router

	*deliveryhttp.BaseHandler
}

func NewSwaggerHandler(baseHandler *deliveryhttp.BaseHandler) *SwaggerHandler {
	router := chi.NewRouter()

	initRoutes(router)
	return &SwaggerHandler{
		Router:      router,
		BaseHandler: baseHandler,
	}

}

func initRoutes(router chi.Router) {
	fs := http.FileServer(http.Dir("./docs/"))
	router.Handle("/*", http.StripPrefix("/docs/", fs))
}
