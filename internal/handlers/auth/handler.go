package auth

import (
	"net/http"

	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
	"github.com/chelochambi/sistema-facturacion-auth/internal/model"
	"github.com/chelochambi/sistema-facturacion-auth/internal/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Usuario string `json:"usuario"`
	Clave   string `json:"clave"`
}

func LoginHandler(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	var user model.Usuario
	var nombres, apellido1, apellido2 string
	var estadoID int

	err := db.DB.QueryRow(`
        SELECT id, username, email, nombres, primer_apellido, segundo_apellido, password_hash, estado_id
        FROM usuarios
        WHERE username = $1
    `, req.Usuario).Scan(&user.ID, &user.Username, &user.Email, &nombres, &apellido1, &apellido2, &user.PasswordHash, &estadoID)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Comprobar si el estado del usuario es ACTIVO
	var estadoCodigo string
	err = db.DB.QueryRow(`
		SELECT codigo FROM estados WHERE id = $1
	`, estadoID).Scan(&estadoCodigo)

	if err != nil || estadoCodigo != "ACT" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Usuario inactivo o no autorizado"})
		return
	}

	// Validar contraseña
	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Clave)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	// Armar nombre completo
	user.Nombres = nombres + " " + apellido1
	if apellido2 != "" {
		user.Nombres += " " + apellido2
	}

	menus := service.ObtenerMenus(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"usuario": user.Sanitizar(),
		"menus":   menus,
	})
}
