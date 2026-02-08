package permission_test

import (
	"axis/services/identity/domain/permission"
	"testing"
)

func TestValidateEntity(t *testing.T) {
	t.Run("Validacion_Exitosa", func(t *testing.T) {
		p := &permission.Permission{
			Code:        "finance:authorize_payment",
			Description: "Permite la autorización de pagos a proveedores",
			Module:      "Finance",
			Status:      permission.Active,
		}

		// Probamos el punto de entrada principal
		err := permission.ValidateEntity(p)
		if err != nil {
			t.Errorf("Se esperaba validación exitosa, se obtuvo error: %v", err)
		}
	})

	t.Run("Validacion_Fallida_Por_Integridad", func(t *testing.T) {
		p := &permission.Permission{
			Code:        "INVALID CODE", // Espacios no permitidos por el Regex en rules.go
			Description: "Desc",
			Module:      "Finance",
			Status:      permission.Active,
		}

		err := permission.ValidateEntity(p)
		if err == nil {
			t.Error("Se esperaba un error de validación debido al formato del código")
		}
	})
}
