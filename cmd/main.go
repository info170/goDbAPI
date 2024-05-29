package main

import (
	"directories/pkg/db"
	"directories/pkg/router"
	"directories/pkg/server"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
)

/**
 * MicroService for fetching data from database
 * It create web-server for use as REST API
 * Get all data from the table: GET http://localhost:8001/api/data/regions?limit=2&offset=10
 * Get specific entity: GET http://localhost:8001/api/data/regions/2
 */

import (
	"context"
	_ "github.com/lib/pq"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbConn, err := db.NewPostgresDB(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USERNAME"),
		DBName:   os.Getenv("DB_DBNAME"),
		SSLMode:  os.Getenv("DB_SSL"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	port := os.Getenv("PORT")
	routers := router.NewRouter()
	webserver := new(server.Server)
	go func() {
		if err := webserver.Run(port, routers.GetRouter(dbConn)); err != nil {
			log.Fatalf("error occured while running http webserver: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	log.Print("TodoApp Shutting Down")

	if err := webserver.Shutdown(context.Background()); err != nil {
		log.Fatalf("error occured on webserver shutting down: %s", err.Error())
	}

}
