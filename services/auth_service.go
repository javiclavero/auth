// services/auth_service.go
package services

import (
	"context"
	"errors"
	"github.com/javiclavero/go-auth-service/ent"
	"github.com/javiclavero/go-auth-service/ent/user"
	"golang.org/x/crypto/bcrypt"
)

// Login valida las credenciales de un usuario.
func Login(client *ent.Client, username, password string) (bool, error) {
	ctx := context.Background()

	// Buscar al usuario por nombre de usuario
	u, err := client.User.
		Query().
		Where(user.UseUsername(username)).
		Only(ctx)
	if err != nil {
		return false, errors.New("usuario no encontrado")
	}

	// Comparar la contraseña hash almacenada con la proporcionada
	if err := bcrypt.CompareHashAndPassword([]byte(u.UsePwd), []byte(password)); err != nil {
		return false, errors.New("contraseña incorrecta")
	}

	return true, nil
}
