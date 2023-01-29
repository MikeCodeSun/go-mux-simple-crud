package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MikeCodeSun/go-mux-api/app"
	"github.com/joho/godotenv"
)

func main() {
	if err:= godotenv.Load(); err !=nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	
	var app app.App
	app.Initilize(dbHost, dbPort, dbUser,dbPassword,dbName)
	fmt.Println("Server is on Port:" + port)
	log.Fatal(http.ListenAndServe(":"+port, app.Router))
	

}