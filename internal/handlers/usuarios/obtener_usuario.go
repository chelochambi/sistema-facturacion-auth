package usuarios

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/chelochambi/sistema-facturacion-auth/internal/model"

	"github.com/gin-gonic/gin"
)

func ObtenerUsuarioHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		var u model.Usuario
		err = db.QueryRow(`
			SELECT id, username, email, nombres, primer_apellido, segundo_apellido 
			FROM usuarios WHERE id = $1 AND estado_id = 1`, id).
			Scan(&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido)

		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Usuario no encontrado"})
			return
		} else if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener usuario"})
			return
		}

		c.JSON(http.StatusOK, u.Sanitizar())
	}
}
