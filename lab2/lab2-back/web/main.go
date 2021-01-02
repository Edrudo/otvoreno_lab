package main

import (
	"context"
	"fmt"
	"github.com/Edrudo/otvoreno_lab/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"net/http"
	"os"
)

type application struct {
	repo     storage.Repository
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// conecting to mongoDB
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// checking if the connection is established
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	app := application{
		repo:     storage.Repository{Client: client},
		errorLog: log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile),
		infoLog:  log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime),
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	fmt.Print("Starting server on port 8080")

	err = srv.ListenAndServe()
	log.Fatal(err)
}
