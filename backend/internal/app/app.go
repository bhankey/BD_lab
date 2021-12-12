package app

import (
	"context"
	configinternal "finance/internal/config"
	httphandler "finance/internal/delivery/http"
	"finance/internal/delivery/http/accounthandler"
	paymentshandler "finance/internal/delivery/http/paymenthandler"
	"finance/internal/delivery/http/reportshandler"
	"finance/internal/delivery/http/swaggerhandler"
	"finance/internal/repository"
	"finance/internal/repository/accountrepo"
	"finance/internal/repository/paymenthistroryrepo"
	"finance/internal/repository/paymentrepo"
	"finance/internal/service"
	accountservise "finance/internal/service/accountservice"
	"finance/internal/service/paymentsservice"
	"finance/internal/service/reportservice"
	"finance/pkg/logger"
	"github.com/go-chi/chi/v5"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	server *http.Server
	ds     *dataSources
	logger logger.Logger
}

func NewApp(configPath string) (*App, error) {
	config := configinternal.GetConfig(configPath)

	log := logger.GetLogger()

	ds, err := newDataSource(config)
	if err != nil {
		return nil, err
	}

	baseRepository := repository.NewRepository(log)

	accountRepo := accountrepo.NewAccountRepo(baseRepository, ds.db)
	paymentsRepo := paymentrepo.NewPaymentsRepo(baseRepository, ds.db)
	paymentsHistoryRepo := paymenthistroryrepo.NewPaymentsHistoryRepo(baseRepository, ds.db)

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
		ds:     ds,
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
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := a.server.Shutdown(ctx); err != nil {
		a.logger.Error(err)
	}

	<-ctx.Done()

	err := a.ds.close()
	if err != nil {
		a.logger.Error(err)
	}
	a.logger.Info("server was shutdown")
}
