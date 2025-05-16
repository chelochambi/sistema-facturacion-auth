package usuarios

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func EliminarUsuarioHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
			return
		}

		_, err = db.Exec(`UPDATE usuarios SET estado_id = 2 WHERE id = $1`, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al desactivar usuario"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"mensaje": "Usuario desactivado correctamente"})
	}
}
