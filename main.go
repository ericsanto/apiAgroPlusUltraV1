package main

import (
	"github.com/ericsanto/apiAgroPlusUltraV1/config"
	"github.com/ericsanto/apiAgroPlusUltraV1/internal/routes"
)


func main() {

  config.Conect()

  routes := routes.Routes()

  routes.Run(":8080")

}


