package user_test

import (
	"axis/services/identity/domain/user"
	"testing"
)

func TestRules(t *testing.T) {

	// 1. Test de Username (Mínimo 5 caracteres)
	t.Run("ValidateUsername", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr error
		}{
			{"Valido_Simple", "auferchis", nil},
			{"Valido_Con_Numeros", "admin_2026", nil},
			{"Falla_Muy_Corto", "ui", user.ErrInvalidUsername},
			{"Falla_Con_Espacios", "user name", user.ErrUsernameCharacters}, // El espacio es "especial"
			{"Falla_Con_Arroba", "admin@axis", user.ErrUsernameCharacters},
			{"Valio_Con_Puntos", "juan.perez", nil},
			{"Falla_Con_Simbolos", "admin!", user.ErrUsernameCharacters},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				err := user.ValidateUsername(tt.input)
				if err != tt.wantErr {
					t.Errorf("Username(%s) se esperaba error: %v, se obtuvo: %v", tt.input, tt.wantErr, err)
				}
			})
		}
	})

	// 2. Test de Email
	t.Run("ValidateEmail", func(t *testing.T) {
		tests := []struct {
			input   string
			wantErr error
		}{
			{"info@axis.com", nil},             // OK
			{"axis.com", user.ErrInvalidEmail}, // Falla: sin @
		}
		for _, tt := range tests {
			err := user.ValidateEmail(tt.input)
			if err != tt.wantErr {
				t.Errorf("Email(%s) se esperaba %v", tt.input, tt.wantErr)
			}
		}
	})

	// 3. Test de Status (Usando tu Enum de Status)
	t.Run("ValidateStatus", func(t *testing.T) {
		tests := []struct {
			name    string
			status  user.Status
			wantErr error
		}{
			{"Status_Activo", user.Active, nil},
			{"Status_Invalido", 99, user.ErrInvalidStatus}, // Fuera de rango
		}
		for _, tt := range tests {
			err := user.ValidateStatus(tt.status)
			if err != tt.wantErr {
				t.Errorf("%s falló: se obtuvo %v", tt.name, err)
			}
		}
	})

	// 4. Test de Groups (Lista y IDs vacíos)
	t.Run("ValidateGroups", func(t *testing.T) {
		tests := []struct {
			name    string
			groups  []string
			wantErr error
		}{
			{"Lista_Valida", []string{"GRP-01", "GRP-02"}, nil},
			{"Lista_Vacia", []string{}, user.ErrEmptyGroups},
			{"ID_Vacio", []string{"GRP-01", " "}, user.ErrInvalidGroupID},
		}
		for _, tt := range tests {
			err := user.ValidateGroups(tt.groups)
			if err != tt.wantErr {
				t.Errorf("%s falló: se obtuvo %v", tt.name, err)
			}
		}
	})

	// 5. Test de Formato de Permisos (modulo:accion)
	t.Run("ValidatePermissionFormat", func(t *testing.T) {
		tests := []struct {
			input   string
			wantErr error
		}{
			{"identity:user_create", nil},                  // OK
			{"identity_create", user.ErrInvalidPermission}, // Falla: sin dos puntos
			{" ", user.ErrInvalidPermission},               // Falla: vacío
		}
		for _, tt := range tests {
			err := user.ValidatePermissionFormat(tt.input)
			if err != tt.wantErr {
				t.Errorf("Permiso(%s) falló: %v", tt.input, err)
			}
		}
	})
}
