package auth

import (
	"net/http"

	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
	"github.com/chelochambi/sistema-facturacion-auth/internal/model"
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
	err := db.DB.QueryRow(`
        SELECT id, username, email, nombre_completo, password_hash
        FROM usuarios WHERE username = $1
    `, req.Usuario).Scan(&user.ID, &user.Username, &user.Email, &user.NombreCompleto, &user.PasswordHash)

	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Clave)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales incorrectas"})
		return
	}

	menus := obtenerMenus(user.ID)
	c.JSON(http.StatusOK, gin.H{
		"usuario": user.Sanitizar(),
		"menus":   menus,
	})
}

func obtenerMenus(usuarioID int) []map[string]interface{} {
	rows, err := db.DB.Query(`
        SELECT m.id, m.nombre, m.ruta, m.icono, m.padre_id, m.orden
        FROM menus m
        JOIN permisos p ON p.menu_id = m.id
        JOIN rol_permiso rp ON rp.permiso_id = p.id
        JOIN usuario_rol ur ON ur.rol_id = rp.rol_id
        WHERE ur.usuario_id = $1
          AND m.estado_id IN (SELECT id FROM estados WHERE codigo = 'ACT')
        ORDER BY m.padre_id NULLS FIRST, m.orden
    `, usuarioID)

	if err != nil {
		return []map[string]interface{}{}
	}
	defer rows.Close()

	var menus []map[string]interface{}
	for rows.Next() {
		var id, orden int
		var nombre string
		var ruta, icono *string
		var padreID *int

		if err := rows.Scan(&id, &nombre, &ruta, &icono, &padreID, &orden); err == nil {
			menus = append(menus, map[string]interface{}{
				"id":     id,
				"nombre": nombre,
				"ruta":   ruta,
				"icono":  icono,
				"padre":  padreID,
				"orden":  orden,
			})
		}
	}
	return menus
}
