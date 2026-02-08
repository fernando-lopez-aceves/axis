package role_test

import (
	"axis/services/identity/domain/role"
	"testing"
)

func TestRoleRules(t *testing.T) {

	t.Run("ValidateName", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr error
		}{
			{"Valido", "Administrator", nil},
			{"Vacio", "   ", role.ErrNameRequired},
			{"Corto", "Ad", role.ErrNameTooShort},
		}
		for _, tt := range tests {
			err := role.ValidateName(tt.input)
			if err != tt.wantErr {
				t.Errorf("[%s] se esperaba %v, se obtuvo %v", tt.name, tt.wantErr, err)
			}
		}
	})

	t.Run("ValidatePermissionsList", func(t *testing.T) {
		t.Run("Formatos", func(t *testing.T) {
			tests := []struct {
				name    string
				perms   []string
				wantErr error
			}{
				{"Valido", []string{"users:read", "billing:write"}, nil},
				{"Invalido_Sin_Dos_Puntos", []string{"usersread"}, role.ErrInvalidPermission},
				{"Invalido_Muy_Corto", []string{"u:"}, role.ErrInvalidPermission},
			}
			for _, tt := range tests {
				err := role.ValidatePermissionsList(tt.perms)
				if err != tt.wantErr {
					t.Errorf("[%s] se esperaba %v, se obtuvo %v", tt.name, tt.wantErr, err)
				}
			}
		})

		t.Run("Limite_Maximo", func(t *testing.T) {
			// Creamos una lista que supere los 100 permisos
			manyPerms := make([]string, 101)
			for i := 0; i < 101; i++ {
				manyPerms[i] = "module:action"
			}
			err := role.ValidatePermissionsList(manyPerms)
			if err != role.ErrTooManyPermissions {
				t.Errorf("Se esperaba error de límite máximo, se obtuvo %v", err)
			}
		})
	})
}
