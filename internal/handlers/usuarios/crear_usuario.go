package usuarios

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CrearUsuarioHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var datos struct {
			Username        string `json:"username"`
			Email           string `json:"email"`
			Nombres         string `json:"nombres"`
			PrimerApellido  string `json:"primer_apellido"`
			SegundoApellido string `json:"segundo_apellido"`
			PasswordHash    string `json:"password"` // Se espera recibir el hash
		}

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		_, err := db.Exec(`
			INSERT INTO usuarios (username, email, nombres, primer_apellido, segundo_apellido, password_hash, estado_id) 
			VALUES ($1, $2, $3, $4, $5, $6, 1)`,
			datos.Username, datos.Email, datos.Nombres, datos.PrimerApellido, datos.SegundoApellido, datos.PasswordHash)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear usuario"})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"mensaje": "Usuario creado correctamente"})
	}
}
