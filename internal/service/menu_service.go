package service

import (
	"github.com/chelochambi/sistema-facturacion-auth/internal/db"
)

func ObtenerMenus(usuarioID int) []map[string]interface{} {
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
