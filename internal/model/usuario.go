package model

type Usuario struct {
	ID             int
	Username       string
	Email          string
	NombreCompleto string
	PasswordHash   string
}

func (u *Usuario) Sanitizar() map[string]interface{} {
	return map[string]interface{}{
		"id":       u.ID,
		"username": u.Username,
		"email":    u.Email,
		"nombre":   u.NombreCompleto,
	}
}
