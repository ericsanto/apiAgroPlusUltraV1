package main

import (
	"os"

	"github.com/ericsanto/apiAgroPlusUltraV1/config/db"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/routes"
)

func main() {

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbPort := "5432"
	dbUser := os.Getenv("DB_USER")
	sslMode := os.Getenv("SSL_MODE")

	db := db.NewDatabase(dbName, dbUser, dbPassword, dbHost, dbPort, sslMode)

	DB, err := db.Connect()

	if err != nil {
		panic(err)
	}

	db.Migrate(DB)

	routes := routes.SetupRoutes()

	routes.Run(":8080")

}
