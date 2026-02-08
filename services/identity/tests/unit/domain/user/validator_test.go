package user_test

import (
	"axis/services/identity/domain/user"
	"testing"
)

func TestValidateFullIntegrity(t *testing.T) {
	t.Run("Usuario_Completo_Valido", func(t *testing.T) {
		usuarioValido := &user.User{
			ID:          "USR-123",
			Username:    "auferchis",
			Email:       "admin@axis.com",
			Status:      user.Active,
			Groups:      []string{"GRP-01"},
			Roles:       []string{"ROLE-ADMIN"},
			Permissions: []string{"identity:read"},
		}

		err := user.ValidateFullIntegrity(usuarioValido)
		if err != nil {
			t.Errorf("No debería haber fallado: %v", err)
		}
	})

	t.Run("Falla_Por_Falta_De_Grupos", func(t *testing.T) {
		usuarioSinGrupos := &user.User{
			Username: "auferchis",
			Email:    "admin@axis.com",
			Groups:   []string{}, // Lista vacía, tu código debe fallar aquí
		}

		err := user.ValidateFullIntegrity(usuarioSinGrupos)
		if err == nil {
			t.Error("Debería haber fallado porque la lista de grupos está vacía")
		}
	})
}
