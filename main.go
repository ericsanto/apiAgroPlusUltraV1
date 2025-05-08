package main

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/routes"
)

func main() {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("erro ao carregar variáveis do .env %w", err)
	// }

	config.Conect()

	routes := routes.SetupRoutes()

	routes.Run(":8080")

}
