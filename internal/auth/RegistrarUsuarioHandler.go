package auth

import (
	"net/http"

	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
	"github.com/chelochambi/sistema-facturacion-auth/internal/model"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegistroRequest struct {
	Username       string `json:"username" binding:"required"`
	Email          string `json:"email" binding:"required,email"`
	Clave          string `json:"clave" binding:"required"`
	NombreCompleto string `json:"nombre_completo"`
	EstadoID       int    `json:"estado_id" binding:"required"`
}

func RegistrarUsuarioHandler(c *gin.Context) {
	var req RegistroRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos", "detalles": err.Error()})
		return
	}

	// Hashear la clave
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Clave), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al hashear la contraseña"})
		return
	}

	// Usamos model.Usuario
	user := model.Usuario{
		Username:       req.Username,
		Email:          req.Email,
		NombreCompleto: req.NombreCompleto,
		PasswordHash:   string(hashedPassword),
	}

	// Insertamos en la base
	_, err = db.DB.Exec(`
        INSERT INTO usuarios (username, email, password_hash, nombre_completo, estado_id, creado_en)
        VALUES ($1, $2, $3, $4, $5, CURRENT_TIMESTAMP)
    `, user.Username, user.Email, user.PasswordHash, user.NombreCompleto, req.EstadoID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo registrar el usuario", "detalle": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "Usuario registrado exitosamente",
		"usuario": user.Sanitizar(),
	})
}
