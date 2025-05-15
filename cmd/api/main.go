package main

import (
	"log"

	"github.com/chelochambi/sistema-facturacion-auth/internal/auth"
	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect() // Conecta a PostgreSQL

	router := gin.Default()
	router.POST("/api/login", auth.LoginHandler)

	log.Println("Servidor corriendo en http://localhost:8080")
	router.Run(":8080")
}
