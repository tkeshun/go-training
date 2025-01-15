package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"pkg/logger"
	"syscall"
	"time"
	"todo/api"
	"todo/infra"
	"todo/usecase"
)

var Logger *logger.Logger

func main() {
	ctx := context.Background()
	Logger = logger.InitLogger(&logger.LoggerConfig{})
	pool, err := infra.NewDBConnection()
	if err != nil {
		Logger.Error(ctx, err.Error(), map[string]any{})
		log.Fatal()
	}

	todoUsecase := usecase.NewTodoService(
		infra.NewTodoRepo(),
		infra.NewPgxTransactionManager(pool),
	)
	userUsecase := usecase.NewUserService(
		infra.NewUserRepo(pool),
	)

	handler := api.NewHandler(
		todoUsecase,
		userUsecase,
	)

	mux := api.SetupRoutes(handler)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		fmt.Println("Starting server on :8080")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Could not listen on :8080: %v\n", err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	<-sigChan
	fmt.Println("\nShutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
	fmt.Println("Server gracefully stopped")
}
