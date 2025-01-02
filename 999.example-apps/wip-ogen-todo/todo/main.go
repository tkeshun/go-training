package main

import (
	"log"
	"net/http"
	"pkg/logger"
	"todo/api"
	"todo/ogen"
)

func main() {
	logger := logger.InitLogger(nil)

	server := api.APIHandler{
		Logger: logger,
	}
	handler, err := ogen.NewServer(&server)
	if err != nil {
		log.Fatalf("failed start server: %v", err)
	}

	// HTTPサーバーの起動
	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	// サーバー起動
	println("Server is running on http://localhost:8080")
	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatalf("failed Listen And Serve: %v", err)
	}
}
