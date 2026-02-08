package permission_test

import (
	"axis/services/identity/domain/permission"
	"testing"
)

func TestPermissionActions(t *testing.T) {

	t.Run("UpdateBasicInfo", func(t *testing.T) {
		p := &permission.Permission{
			Code:        "old:code",
			Description: "Old Desc",
			Module:      "OldModule",
		}

		newCode := "inventory:stock_update"
		newDesc := "Permite actualizar el stock físico"
		newMod := "Inventory"

		p.UpdateBasicInfo(newCode, newDesc, newMod)

		if p.Code != newCode || p.Description != newDesc || p.Module != newMod {
			t.Error("La información básica no se actualizó correctamente")
		}
		if p.UpdatedAt.IsZero() {
			t.Error("UpdatedAt debería tener una fecha asignada")
		}
	})

	t.Run("LifeCycle", func(t *testing.T) {
		p := &permission.Permission{Status: permission.Inactive}

		// Probar Activación
		p.Activate()
		if p.Status != permission.Active {
			t.Error("Fallo al activar el permiso")
		}

		// Probar Desactivación
		p.Deactivate()
		if p.Status != permission.Inactive {
			t.Error("Fallo al desactivar el permiso")
		}

		// Probar Borrado Lógico
		p.Delete()
		if p.Status != permission.Deleted {
			t.Error("Fallo al marcar como borrado")
		}
		if p.DeletedAt == nil {
			t.Error("DeletedAt no puede ser nil tras un Delete()")
		}
	})
}
