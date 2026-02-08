package permission_test

import (
	"axis/services/identity/domain/permission"
	"testing"
)

func TestPermissionRules(t *testing.T) {

	t.Run("ValidateCode", func(t *testing.T) {
		tests := []struct {
			name    string
			input   string
			wantErr error
		}{
			{"Valido_Simple", "inventory:read", nil},
			{"Valido_Con_Guiones", "sales_reports:generate_pdf", nil},
			{"Falla_Sin_Separador", "inventory_read", permission.ErrCodeInvalidFormat},
			{"Falla_Mayusculas", "Inventory:Read", permission.ErrCodeInvalidFormat},
			{"Falla_Vacio", " ", permission.ErrCodeRequired},
			{"Falla_Simbolos", "inventory:read!", permission.ErrCodeInvalidFormat},
		}

		for _, tt := range tests {
			err := permission.ValidateCode(tt.input)
			if err != tt.wantErr {
				t.Errorf("[%s] se esperaba %v, se obtuvo %v", tt.name, tt.wantErr, err)
			}
		}
	})

	t.Run("ValidateModule", func(t *testing.T) {
		if err := permission.ValidateModule("  "); err != permission.ErrModuleRequired {
			t.Error("Debería fallar si el módulo está vacío")
		}
		if err := permission.ValidateModule("Warehouse"); err != nil {
			t.Error("Debería aceptar un módulo válido")
		}
	})
}
