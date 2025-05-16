package usuarios

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ActualizarUsuarioHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var datos struct {
			Username        string `json:"username"`
			Email           string `json:"email"`
			Nombres         string `json:"nombres"`
			PrimerApellido  string `json:"primer_apellido"`
			SegundoApellido string `json:"segundo_apellido"`
		}

		if err := c.ShouldBindJSON(&datos); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
			return
		}

		_, err = db.Exec(`
			UPDATE usuarios 
			SET username = $1, email = $2, nombres = $3, primer_apellido = $4, segundo_apellido = $5 
			WHERE id = $6`,
			datos.Username, datos.Email, datos.Nombres, datos.PrimerApellido, datos.SegundoApellido, id)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario actualizado correctamente"})
	}
}
