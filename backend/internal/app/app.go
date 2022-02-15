package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	configinternal "github.com/bhankey/BD_lab/backend/internal/config"
	httphandler "github.com/bhankey/BD_lab/backend/internal/delivery/http"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/accounthandler"
	paymentshandler "github.com/bhankey/BD_lab/backend/internal/delivery/http/paymenthandler"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/reportshandler"
	"github.com/bhankey/BD_lab/backend/internal/delivery/http/swaggerhandler"
	"github.com/bhankey/BD_lab/backend/internal/repository"
	"github.com/bhankey/BD_lab/backend/internal/repository/accountrepo"
	"github.com/bhankey/BD_lab/backend/internal/repository/paymenthistroryrepo"
	"github.com/bhankey/BD_lab/backend/internal/repository/paymentrepo"
	"github.com/bhankey/BD_lab/backend/internal/service"
	accountservise "github.com/bhankey/BD_lab/backend/internal/service/accountservice"
	"github.com/bhankey/BD_lab/backend/internal/service/paymentsservice"
	"github.com/bhankey/BD_lab/backend/internal/service/reportservice"
	"github.com/bhankey/BD_lab/backend/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

type App struct {
	server *http.Server
	ds     *dataSources
	logger logger.Logger
}

const shutDownTimeoutSeconds = 10

func NewApp(configPath string) (*App, error) {
	config := configinternal.GetConfig(configPath)

	log := logger.GetLogger()

	dataSources, err := newDataSource(config)
	if err != nil {
		return nil, err
	}

	baseRepository := repository.NewRepository(log)

	accountRepo := accountrepo.NewAccountRepo(baseRepository, dataSources.db)
	paymentsRepo := paymentrepo.NewPaymentsRepo(baseRepository, dataSources.db)
	paymentsHistoryRepo := paymenthistroryrepo.NewPaymentsHistoryRepo(baseRepository, dataSources.db)

	baseService := service.NewService(log)

	accountService := accountservise.NewAccountService(baseService, accountRepo)
	paymentsService := paymentsservice.NewPaymentsService(baseService, paymentsRepo, paymentsHistoryRepo, accountRepo)
	reportsService := reportservice.NewReportService(baseService, paymentsRepo, paymentsHistoryRepo, accountRepo)

	baseHandler := httphandler.NewHandler(log)

	accountHandler := accounthandler.NewAccountHandler(baseHandler, accountService)
	paymentsHandler := paymentshandler.NewPaymentHandler(baseHandler, paymentsService)
	reportsHandler := reportshandler.NewReportsHandler(baseHandler, reportsService)

	swaggerHandler := swaggerhandler.NewSwaggerHandler(baseHandler)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
	}))

	router.Use(func(handler http.Handler) http.Handler {
		return httphandler.LoggingMiddleware(log)(handler)
	})

	router.Mount("/account", accountHandler.Router)
	router.Mount("/payments", paymentsHandler.Router)
	router.Mount("/reports", reportsHandler.Router)

	router.Mount("/docs", swaggerHandler.Router)

	server := &http.Server{
		Addr:    ":" + config.Server.Port,
		Handler: router,
	}

	return &App{
		logger: log,
		server: server,
		ds:     dataSources,
	}, nil
}

func (a *App) Start() {
	a.logger.Info("staring server on port: " + a.server.Addr)
	go func() {
		if err := a.server.ListenAndServe(); err != nil {
			a.logger.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	a.logger.Info("received signal to shutdown server")
	ctx, cancel := context.WithTimeout(context.Background(), shutDownTimeoutSeconds*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error(err)
	}

	<-ctx.Done()

	if err := a.ds.close(); err != nil {
		a.logger.Error(err)
	}
	a.logger.Info("server was shutdown")
}
