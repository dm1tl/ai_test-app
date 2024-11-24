package main

import (
	"ai_test-app/clients/sso"
	"ai_test-app/clients/testgen"
	"ai_test-app/internal/config"
	"ai_test-app/internal/config/env"
	"ai_test-app/internal/handler"
	"ai_test-app/internal/repository"
	"ai_test-app/internal/repository/connector"
	"ai_test-app/internal/services"
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("loading env variables")
	if err := config.Load(); err != nil {
		logrus.Fatalf("couldn't load env variables %s", err.Error())
	}
	logrus.Info("done_loading env variables")

	logrus.Info("loading db config")
	dbcfg, err := env.NewDBConfig()
	if err != nil {
		logrus.Fatalf("couldn't load db config %s", err.Error())
	}
	logrus.Info("done_loading db config")
	logrus.Info("db config", dbcfg)

	logrus.Info("conecting to postgres-docker db")
	db, err := connector.NewPostgresDB(dbcfg)
	if err != nil {
		logrus.Fatalf("couldn't initialize db %s", err.Error())
	}
	logrus.Info("done_conecting to postgres-docker db")

	logrus.Info("loading sso cfg")
	ssoConfig, err := env.NewSSOConfig()
	if err != nil {
		logrus.Fatalf("couldn't load sso config %s", err.Error())
	}
	logrus.Info("done_loading sso cfg")
	logrus.Info("sso cfg", ssoConfig)

	logrus.Info("loading testgen cfg")
	genTestConfig, err := env.NewGenTestConfig()
	if err != nil {
		logrus.Fatalf("couldn't load genTest config %s", err.Error())
	}
	logrus.Info("done_loading testgen cfg")
	logrus.Info("gentestcfg", genTestConfig)

	logrus.Info("Initializing mock SSO gRPC client")
	ssoGRPCMockClient := sso.NewMockSSOServiceClient(ssoConfig)
	logrus.Info("done_Initializing mock SSO gRPC client")

	logrus.Info("Wrapping SSO gRPC client with SSOClientWrapper")
	ssoClient := sso.NewSSOClientWrapper(ssoGRPCMockClient)
	logrus.Info("done_Wrapping SSO gRPC client with SSOClientWrapper")

	logrus.Info("Initializing mock gen test gRPC client")
	genTestMockClient := testgen.NewMockTestGenServiceClient(genTestConfig)
	genTestClient := testgen.NewGenClient(genTestMockClient)
	logrus.Info("done_Initializing mock gen test gRPC client")

	logrus.Info("initializing repo layer")
	repository := repository.NewRepository(db)
	logrus.Info("done_initializing repo layer")

	logrus.Info("initializing service layer")
	services := services.NewService(repository, ssoClient, genTestClient)
	logrus.Info("done_initializing service layer")

	logrus.Info("initializing handler layer")
	handler := handler.NewHandler(services)
	logrus.Info("done_initializing handler layer")

	logrus.Info("loading server cfg")
	serverConfig, err := env.NewHTTPServerConfig()
	if err != nil {
		logrus.Fatalf("couldn't load httpServer config %s", err.Error())
	}
	logrus.Info("done_loading server cfg")
	logrus.Info("srv cfg", serverConfig)

	logrus.Info("starting server")
	srv := new(env.Server)
	go func() {
		if err := srv.Run(*serverConfig, handler.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("cannot start server %s", err.Error())
		}
	}()
	logrus.Info("server started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	logrus.Print("Todo app shutting down")
	if err := srv.ShutDown(context.Background()); err != nil {
		logrus.Errorf("couldn't shut down an app %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("couldn't close db connection %s", err.Error())
	}
	logrus.Print("Todo app shutted down")
}
