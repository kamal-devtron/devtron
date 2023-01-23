package main

import (
	"fmt"
	"github.com/devtron-labs/devtron/internal/middleware"
	"github.com/go-pg/pg"
	"go.uber.org/zap"
	"net/http"
	"os"
)

type App struct {
	db *pg.DB
	//sessionManager *authMiddleware.SessionManager
	MuxRouter *MuxRouter
	Logger    *zap.SugaredLogger
	server    *http.Server
	//telemetry      telemetry.TelemetryEventClient
	//posthogClient  *telemetry.PosthogClient
}

func NewApp(db *pg.DB,
	//sessionManager *authMiddleware.SessionManager,
	MuxRouter *MuxRouter,
	//telemetry telemetry.TelemetryEventClient,
	//posthogClient *telemetry.PosthogClient,
	Logger *zap.SugaredLogger) *App {
	return &App{
		db: db,
		//sessionManager: sessionManager,
		MuxRouter: MuxRouter,
		Logger:    Logger,
		//telemetry:      telemetry,
		//posthogClient:  posthogClient,
	}
}
func (app *App) Start() {
	fmt.Println("starting k8s client App")

	port := 8080 //TODO: extract from environment variable
	app.Logger.Debugw("starting server")
	app.Logger.Infow("starting server on ", "port", port)
	app.MuxRouter.Init()
	//authEnforcer := casbin2.Create()

	//_, err := app.telemetry.SendTelemetryInstallEventEA()

	//if err != nil {
	//	app.Logger.Warnw("telemetry installation success event failed", "err", err)
	//}
	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: app.MuxRouter.Router}

	app.MuxRouter.Router.Use(middleware.PrometheusMiddleware)
	app.server = server

	err := server.ListenAndServe()
	if err != nil {
		app.Logger.Errorw("error in startup", "err", err)
		os.Exit(2)
	}
}

func (app *App) Stop() {
	app.Logger.Info("stopping k8s client App")
	//posthogCl := app.posthogClient.Client
	//if posthogCl != nil {
	//	app.Logger.Info("flushing messages of posthog")
	//	posthogCl.Close()
	//}
}