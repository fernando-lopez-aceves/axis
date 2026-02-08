package group

import (
	"errors"
	"strings"
)

// Errores de dominio con prefijo claro para trazabilidad en logs.
var (
	ErrNameEmpty           = errors.New("group_rule: el nombre es obligatorio")
	ErrNameTooShort        = errors.New("group_rule: el nombre debe tener al menos 3 caracteres")
	ErrDescriptionEmpty    = errors.New("group_rule: la descripción es obligatoria")
	ErrDescriptionTooShort = errors.New("group_rule: la descripción debe ser detallada (min. 10 caracteres)")
	ErrInvalidStatus       = errors.New("group_rule: el estado proporcionado no es válido")
	ErrInvalidRoleID       = errors.New("group_rule: el listado contiene un ID de rol vacío o nulo")
	ErrInvalidPermission   = errors.New("group_rule: el permiso no cumple el formato 'modulo:accion'")
)

// ValidateFullIntegrity orquesta la integridad total de la entidad Group.
// Esta es la función maestra que garantiza que el grupo cumple con todas las reglas de negocio.
func ValidateFullIntegrity(group *Group) error {
	if err := ValidateName(group.Name); err != nil {
		return err
	}
	if err := ValidateDescription(group.Description); err != nil {
		return err
	}
	if err := ValidateStatus(group.Status); err != nil {
		return err
	}
	if err := ValidateRolesList(group.Roles); err != nil {
		return err
	}
	if err := ValidatePermissionsList(group.Permissions); err != nil {
		return err
	}
	return nil
}

// ValidateName asegura que el nombre sea consistente y no solo espacios en blanco.
func ValidateName(name string) error {
	trimmedName := strings.TrimSpace(name)
	if trimmedName == "" {
		return ErrNameEmpty
	}
	if len(trimmedName) < 3 {
		return ErrNameTooShort
	}
	return nil
}

// ValidateDescription garantiza que la descripción no esté vacía y aporte valor informativo.
func ValidateDescription(description string) error {
	trimmedDescription := strings.TrimSpace(description)
	if trimmedDescription == "" {
		return ErrDescriptionEmpty
	}
	if len(trimmedDescription) < 10 {
		return ErrDescriptionTooShort
	}
	return nil
}

// ValidateStatus verifica que el estado esté dentro del rango definido por el Iota de Status.
func ValidateStatus(status Status) error {
	if status < Inactive || status > Deleted {
		return ErrInvalidStatus
	}
	return nil
}

// ValidateRolesList limpia y valida que no se intenten guardar roles inexistentes o vacíos.
func ValidateRolesList(roles []string) error {
	for _, roleID := range roles {
		if strings.TrimSpace(roleID) == "" {
			return ErrInvalidRoleID
		}
	}
	return nil
}

// ValidatePermissionsList verifica el formato granular del permiso (ej. "warehouse:create").
func ValidatePermissionsList(permissions []string) error {
	for _, code := range permissions {
		cleanCode := strings.TrimSpace(code)
		if cleanCode == "" || !strings.Contains(cleanCode, ":") {
			return ErrInvalidPermission
		}
	}
	return nil
}
