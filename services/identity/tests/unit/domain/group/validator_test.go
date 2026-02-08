package group_test

import (
	"axis/services/identity/domain/group"
	"testing"
)

func TestValidateFullIntegrity(t *testing.T) {
	t.Run("Grupo_Completo_Valido", func(t *testing.T) {
		g := &group.Group{
			Name:        "Recursos Humanos",
			Description: "Gestión de personal y nóminas",
			Status:      group.Active,
			Roles:       []string{"ROLE-HR-ADMIN"},
			Permissions: []string{"hr:write", "hr:read"},
		}

		err := group.ValidateFullIntegrity(g)
		if err != nil {
			t.Errorf("No debería haber fallado la integridad: %v", err)
		}
	})

	t.Run("Falla_Por_Descripcion_Pobre", func(t *testing.T) {
		g := &group.Group{
			Name:        "Soporte IT",
			Description: "Soporte", // Menos de 10 chars
			Status:      group.Active,
		}

		err := group.ValidateFullIntegrity(g)
		if err != group.ErrDescriptionTooShort {
			t.Errorf("Se esperaba error de descripción corta, se obtuvo %v", err)
		}
	})
}
