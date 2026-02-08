package user

// ValidateFullUser orquesta todas las reglas de validación sobre la entidad User.
// Actúa como el portero final antes de que los datos toquen la base de datos.
func ValidateFullUser(user *User) error {
	// 1. Validar Identidad básica
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}

	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	// 2. Validar Seguridad
	// Solo validamos contraseña vacía si es un usuario nuevo (ID vacío)
	if user.ID == "" {
		if err := ValidateEmptyPassword(user.Password); err != nil {
			return err
		}
	}

	// 3. Validar Integridad de Acceso (Requerimiento 2026-02-05)
	if err := ValidateStatus(user.Status); err != nil {
		return err
	}

	if err := ValidateGroups(user.Groups); err != nil {
		return err
	}

	// Sincronización: Agregamos la validación de Roles que faltaba
	if err := ValidateRoles(user.Roles); err != nil {
		return err
	}

	// Los permisos individuales son opcionales, pero si existen, deben ser válidos
	if len(user.Permissions) > 0 {
		if err := ValidatePermissionsList(user.Permissions); err != nil {
			return err
		}
	}

	// 4. Normalización final (Garantiza consistencia en la DB)
	user.Email = NormalizeEmail(user.Email)

	return nil
}
