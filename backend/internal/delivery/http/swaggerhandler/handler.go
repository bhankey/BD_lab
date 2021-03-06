package swaggerhandler

import (
	"net/http"

	deliveryhttp "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/go-chi/chi/v5"
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
