package role_test

import (
	"axis/services/identity/domain/role"
	"testing"
)

func TestValidateFullIntegrity(t *testing.T) {
	t.Run("Rol_Valido", func(t *testing.T) {
		r := &role.Role{
			Name:        "InventoryManager",
			Description: "Capacidad total sobre el módulo de inventarios",
			Status:      role.Active,
			Permissions: []string{"inventory:create", "inventory:update"},
		}

		err := role.ValidateFullIntegrity(r)
		if err != nil {
			t.Errorf("No debería haber fallado la integridad: %v", err)
		}
	})

	t.Run("Falla_Nombre_Y_Evita_Validar_Resto", func(t *testing.T) {
		r := &role.Role{
			Name: " ", // Error aquí primero
		}

		err := role.ValidateFullIntegrity(r)
		if err != role.ErrNameRequired {
			t.Errorf("Se esperaba ErrNameRequired, se obtuvo %v", err)
		}
	})
}
