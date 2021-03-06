package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/matheusmhmelo/go-api-tdd/config"
	"github.com/matheusmhmelo/go-api-tdd/internal/handler"
	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/credential"
	"github.com/matheusmhmelo/go-api-tdd/internal/pkg/repository"
	"github.com/matheusmhmelo/go-api-tdd/internal/server"
	"log"
	"net/http"
)

func main() {
	config.Init()

	conn := config.NewDatabaseConn(context.Background())
	db := repository.NewDatabase(conn)
	service := credential.NewService(db, uuid.New)
	h := handler.NewCredential(service)

	router := server.NewRouter(h)

	log.Println("waiting requests...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
