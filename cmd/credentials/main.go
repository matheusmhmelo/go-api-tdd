package main

import (
	"context"
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
	service := credential.NewService(db)
	cred := handler.NewCredential(service)

	router := server.NewRouter(cred)

	log.Println("waiting requests...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
