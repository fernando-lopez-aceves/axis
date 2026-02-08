package role

import (
	"errors"
	"strings"
)

// Constantes de configuración de dominio
const MaxPermissionsPerRole = 100

// Errores de dominio con nombres simplificados (el contexto del paquete 'role' ya está implícito)
var (
	ErrNameRequired        = errors.New("role_rule: el nombre del rol es obligatorio")
	ErrNameTooShort        = errors.New("role_rule: el nombre debe tener al menos 3 caracteres")
	ErrDescriptionRequired = errors.New("role_rule: la descripción es obligatoria")
	ErrInvalidPermission   = errors.New("role_rule: formato de permiso inválido (debe ser modulo:accion)")
	ErrInvalidStatus       = errors.New("role_rule: el estado del rol no es válido")
	ErrTooManyPermissions  = errors.New("role_rule: el rol supera el límite máximo de 100 permisos")
)

// ValidateFullIntegrity verifica que el rol cumpla con todas las especificaciones de negocio.
func ValidateFullIntegrity(role *Role) error {
	if err := ValidateName(role.Name); err != nil {
		return err
	}
	if err := ValidateDescription(role.Description); err != nil {
		return err
	}
	if err := ValidateStatus(role.Status); err != nil {
		return err
	}
	if err := ValidatePermissionsList(role.Permissions); err != nil {
		return err
	}
	return nil
}

// ValidateName asegura la presencia y longitud mínima del identificador del rol.
func ValidateName(name string) error {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ErrNameRequired
	}
	if len(trimmed) < 3 {
		return ErrNameTooShort
	}
	return nil
}

// ValidateDescription garantiza que el rol tenga documentación suficiente para auditoría.
func ValidateDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return ErrDescriptionRequired
	}
	return nil
}

// ValidateStatus verifica que el estado esté dentro del rango del Iota (Inactive, Active, Deleted).
func ValidateStatus(status Status) error {
	if status < Inactive || status > Deleted {
		return ErrInvalidStatus
	}
	return nil
}

// ValidatePermissionsList valida el formato y limita la cantidad de permisos para evitar el "Privilege Creep".
func ValidatePermissionsList(permissions []string) error {
	if len(permissions) > MaxPermissionsPerRole {
		return ErrTooManyPermissions
	}

	for _, p := range permissions {
		cleanPermission := strings.TrimSpace(p)
		if !strings.Contains(cleanPermission, ":") || len(cleanPermission) < 3 {
			return ErrInvalidPermission
		}
	}
	return nil
}
