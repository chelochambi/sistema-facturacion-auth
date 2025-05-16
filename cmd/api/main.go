package main

import (
	"log"
	"time"

	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
	"github.com/chelochambi/sistema-facturacion-auth/internal/handlers/auth"
	"github.com/chelochambi/sistema-facturacion-auth/internal/handlers/usuarios"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.Connect() // Conecta a PostgreSQL

	router := gin.Default()

	// Middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// Ruta autenticación
	api := router.Group("/api")
	{
		api.POST("/auth/login", auth.LoginHandler)
		// Rutas de usuarios

		api.GET("/usuarios", usuarios.ListarUsuariosHandler(db.DB))
		api.GET("/usuarios/:id", usuarios.ObtenerUsuarioHandler(db.DB))
		api.PUT("/usuarios/:id", usuarios.ActualizarUsuarioHandler(db.DB))
		api.DELETE("/usuarios/:id", usuarios.BloquearUsuarioHandler(db.DB))
		api.POST("/usuarios/registro", usuarios.CrearUsuarioHandler(db.DB))

		log.Println("Servidor corriendo en http://localhost:8080")
	}
	router.Run(":8080")
}
