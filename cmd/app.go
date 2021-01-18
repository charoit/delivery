package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	_ "github.com/jackc/pgx/v4/pgxpool"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"delivery/services/orders"
	"delivery/services/orders/api"
	"delivery/services/orders/service"
	"delivery/services/orders/storage"
)

type App struct {
	httpServer *http.Server
	orders     orders.Service
}

func NewApp() *App {
	pgHost := viper.GetString("DATABASE_URL")

	db := initDB(pgHost)

	os := storage.NewStorage(db)
	return &App{
		httpServer: nil,
		orders:     service.NewService(os),
	}
}

func (a *App) Run(port string) error {
	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	root := router.Group("/api")

	// Orders API endpoints
	api.RegisterHTTPEndpoints(root, a.orders)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}

func initDB(host string) *sqlx.DB {
	db := sqlx.MustConnect("postgres", host)
	if err := db.Ping(); err != nil {
		log.Fatal(errors.WithStack(err))
	}
	return db
}
