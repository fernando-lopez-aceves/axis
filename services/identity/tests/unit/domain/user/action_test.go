package user_test

import (
	"axis/services/identity/domain/user"
	"testing"
)

func TestActions(t *testing.T) {

	t.Run("Activate", func(t *testing.T) {
		u := &user.User{Status: user.Inactive}
		u.Activate()

		if u.Status != user.Active {
			t.Errorf("Se esperaba Active, se obtuvo %v", u.Status)
		}
		if u.UpdatedAt.IsZero() {
			t.Error("UpdatedAt no se actualizó")
		}
	})

	t.Run("SoftDelete", func(t *testing.T) { // Antes era "Delete"
		u := &user.User{Status: user.Active}
		u.SoftDelete()

		if u.Status != user.Deleted {
			t.Errorf("Se esperaba Deleted, se obtuvo %v", u.Status)
		}
		if u.DeletedAt == nil {
			t.Error("DeletedAt no puede ser nil tras SoftDelete")
		}
	})

	t.Run("ChangePassword", func(t *testing.T) { // Antes era "UpdatePassword"
		u := &user.User{Password: "old_hash"}
		newHash := "new_secret_hash"
		u.ChangePassword(newHash)

		if u.Password != newHash {
			t.Error("La contraseña no cambió")
		}
	})

	t.Run("AddGroup", func(t *testing.T) {
		u := &user.User{Groups: []string{"GRP-1"}}
		u.AddGroup("GRP-2")

		if len(u.Groups) != 2 {
			t.Errorf("Se esperaban 2 grupos, se obtuvo %d", len(u.Groups))
		}
	})
}
