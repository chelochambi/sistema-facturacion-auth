package usuarios

import (
	"database/sql"
	"net/http"

	"github.com/chelochambi/sistema-facturacion-auth/internal/model"
	"github.com/gin-gonic/gin"
)

func ListarUsuariosHandler(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		rows, err := db.Query(`
			SELECT id, username, email, nombres, primer_apellido, segundo_apellido
			FROM usuarios 
			WHERE estado_id = 1
			ORDER BY id`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al listar usuarios"})
			return
		}
		defer rows.Close()

		var usuarios []map[string]interface{}
		for rows.Next() {
			var u model.Usuario
			if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Nombres, &u.PrimerApellido, &u.SegundoApellido); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error de lectura"})
				return
			}
			usuarios = append(usuarios, u.Sanitizar())
		}

		c.JSON(http.StatusOK, usuarios)
	}
}
