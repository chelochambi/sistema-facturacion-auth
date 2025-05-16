package model

type Usuario struct {
	ID              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Nombres         string `json:"nombres"`
	PrimerApellido  string `json:"primer_apellido"`
	SegundoApellido string `json:"segundo_apellido"`
	PasswordHash    string `json:"-"` // No exponer hash
}

func (u *Usuario) Sanitizar() map[string]interface{} {
	return map[string]interface{}{
		"id":               u.ID,
		"username":         u.Username,
		"email":            u.Email,
		"nombres":          u.Nombres,
		"primer_apellido":  u.PrimerApellido,
		"segundo_apellido": u.SegundoApellido,
	}
}
