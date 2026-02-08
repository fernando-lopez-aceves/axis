package group_test

import (
	"axis/services/identity/domain/group"
	"testing"
)

func TestGroupActions(t *testing.T) {

	t.Run("AddRole", func(t *testing.T) {
		g := &group.Group{Roles: []string{"ROLE_VIEWER"}}

		// 1. Añadir nuevo
		g.AddRole("ROLE_ADMIN")
		if len(g.Roles) != 2 {
			t.Errorf("Se esperaban 2 roles, se obtuvieron %d", len(g.Roles))
		}

		// 2. Probar duplicado (No debería añadirlo)
		g.AddRole("ROLE_ADMIN")
		if len(g.Roles) != 2 {
			t.Error("Se detectó un duplicado en la lista de roles")
		}
	})

	t.Run("RemoveRole", func(t *testing.T) {
		g := &group.Group{Roles: []string{"ROLE_1", "ROLE_2"}}

		g.RemoveRole("ROLE_1")

		if len(g.Roles) != 1 || g.Roles[0] != "ROLE_2" {
			t.Error("El rol no fue eliminado correctamente")
		}
	})

	t.Run("AddPermission", func(t *testing.T) {
		g := &group.Group{Permissions: []string{"inventory:read"}}

		g.AddPermission("inventory:write")

		if len(g.Permissions) != 2 {
			t.Error("El permiso no fue añadido")
		}

		// Verificar que el tiempo de actualización cambió
		if g.UpdatedAt.IsZero() {
			t.Error("UpdatedAt debería haberse actualizado")
		}
	})

	t.Run("UpdateBasicInfo", func(t *testing.T) {
		g := &group.Group{Name: "Old Name", Description: "Old Description"}

		newName := "Warehouse Team"
		newDesc := "Equipos encargados del inventario físico"

		g.UpdateBasicInfo(newName, newDesc)

		if g.Name != newName || g.Description != newDesc {
			t.Error("La información básica no se actualizó correctamente")
		}
	})

	t.Run("StatusTransitions", func(t *testing.T) {
		g := &group.Group{Status: group.Inactive}

		// Activar
		g.Activate()
		if g.Status != group.Active {
			t.Error("Fallo al activar el grupo")
		}

		// Eliminar
		g.Delete()
		if g.Status != group.Deleted {
			t.Error("Fallo al realizar SoftDelete")
		}
		if g.DeletedAt == nil {
			t.Error("DeletedAt debería tener una fecha asignada")
		}
	})
}
