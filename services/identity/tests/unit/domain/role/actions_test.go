package role_test

import (
	"axis/services/identity/domain/role"
	"testing"
)

func TestRoleActions(t *testing.T) {

	t.Run("AddPermission", func(t *testing.T) {
		r := &role.Role{Permissions: []string{"inventory:read"}}

		// 1. Añadir nuevo permiso
		r.AddPermission("inventory:write")
		if len(r.Permissions) != 2 {
			t.Errorf("Se esperaban 2 permisos, se obtuvieron %d", len(r.Permissions))
		}

		// 2. Probar duplicado (No debe añadirlo)
		r.AddPermission("inventory:write")
		if len(r.Permissions) != 2 {
			t.Error("Se detectó un duplicado en la lista de permisos")
		}
	})

	t.Run("RemovePermission", func(t *testing.T) {
		r := &role.Role{Permissions: []string{"admin:write", "admin:read"}}

		r.RemovePermission("admin:write")

		if len(r.Permissions) != 1 || r.Permissions[0] != "admin:read" {
			t.Error("El permiso no fue eliminado correctamente")
		}
	})

	t.Run("UpdateBasicInfo", func(t *testing.T) {
		r := &role.Role{Name: "Old Name", Description: "Old Description"}

		newName := "Super Admin"
		newDesc := "Control total sobre todos los módulos"

		r.UpdateBasicInfo(newName, newDesc)

		if r.Name != newName || r.Description != newDesc {
			t.Error("La información básica del rol no se actualizó")
		}

		if r.UpdatedAt.IsZero() {
			t.Error("UpdatedAt no se disparó al actualizar información")
		}
	})

	t.Run("LifeCycleTransitions", func(t *testing.T) {
		r := &role.Role{Status: role.Inactive}

		// Activar
		r.Activate()
		if r.Status != role.Active {
			t.Error("Fallo al activar el rol")
		}

		// Desactivar
		r.Deactivate()
		if r.Status != role.Inactive {
			t.Error("Fallo al desactivar el rol")
		}

		// Borrado lógico
		r.Delete()
		if r.Status != role.Deleted {
			t.Error("Fallo al realizar borrado lógico del rol")
		}
		if r.DeletedAt == nil {
			t.Error("DeletedAt debería tener la estampa de tiempo")
		}
	})
}
