package model

type Usuario struct {
	ID             int    `json:"id"`
	Username       string `json:"username"`
	Email          string `json:"email"`
	NombreCompleto string `json:"nombre"`
	PasswordHash   string `json:"-"` // con "-" evitamos que se serialice
}

func (u *Usuario) Sanitizar() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"nombre":   u.NombreCompleto,
	}
}
