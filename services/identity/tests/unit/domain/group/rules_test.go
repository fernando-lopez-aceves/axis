package group_test

import (
	"axis/services/identity/domain/group"
	"testing"
)

func TestGroupRules(t *testing.T) {

	t.Run("ValidateName", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr error
		}{
			{"Valido", "Administradores", nil},
			{"Vacio", "   ", group.ErrNameEmpty},
			{"Muy_Corto", "Ad", group.ErrNameTooShort},
		}
		for _, tt := range tests {
			err := group.ValidateName(tt.input)
			if err != tt.wantErr {
				t.Errorf("[%s] se esperaba %v, se obtuvo %v", tt.name, tt.wantErr, err)
			}
		}
	})

	t.Run("ValidateDescription", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr error
		}{
			{"Valido", "Acceso total al sistema de identidad", nil},
			{"Vacio", "", group.ErrDescriptionEmpty},
			{"Insuficiente", "Acceso", group.ErrDescriptionTooShort},
		}
		for _, tt := range tests {
			err := group.ValidateDescription(tt.input)
			if err != tt.wantErr {
				t.Errorf("[%s] se esperaba %v, se obtuvo %v", tt.name, tt.wantErr, err)
			}
		}
	})

	t.Run("ValidateRolesList", func(t *testing.T) {
		tests := []struct {
			name    string
			roles   []string
			wantErr error
		}{
			{"Lista_Valida", []string{"ROLE-01", "ROLE-02"}, nil},
			{"ID_Vacio", []string{"ROLE-01", " "}, group.ErrInvalidRoleID},
		}
		for _, tt := range tests {
			err := group.ValidateRolesList(tt.roles)
			if err != tt.wantErr {
				t.Errorf("[%s] falló", tt.name)
			}
		}
	})

	t.Run("ValidatePermissionsList", func(t *testing.T) {
		tests := []struct {
			name    string
			perms   []string
			wantErr error
		}{
			{"Formato_Correcto", []string{"inventory:read"}, nil},
			{"Formato_Incorrecto", []string{"inventory_read"}, group.ErrInvalidPermission},
			{"Vacio", []string{" "}, group.ErrInvalidPermission},
		}
		for _, tt := range tests {
			err := group.ValidatePermissionsList(tt.perms)
			if err != tt.wantErr {
				t.Errorf("[%s] falló", tt.name)
			}
		}
	})
}
