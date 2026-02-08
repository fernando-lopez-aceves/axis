package permission

import (
	"errors"
	"regexp"
	"strings"
)

// Errores de dominio simplificados para el contexto de Permission.
var (
	ErrCodeRequired        = errors.New("permission_rule: el código es obligatorio")
	ErrCodeInvalidFormat   = errors.New("permission_rule: formato inválido (ej. modulo:accion_permitida)")
	ErrDescriptionRequired = errors.New("permission_rule: la descripción es obligatoria")
	ErrModuleRequired      = errors.New("permission_rule: el módulo es obligatorio")
	ErrInvalidStatus       = errors.New("permission_rule: el estado no es válido")
)

// ValidateFullIntegrity es la FUNCIÓN MAESTRA que orquesta la validación total.
// Se asegura de que el permiso sea consistente antes de entrar a la base de datos.
func ValidateFullIntegrity(p *Permission) error {
	if err := ValidateCode(p.Code); err != nil {
		return err
	}
	if err := ValidateDescription(p.Description); err != nil {
		return err
	}
	if err := ValidateModule(p.Module); err != nil {
		return err
	}
	if err := ValidateStatus(p.Status); err != nil {
		return err
	}
	return nil
}

// ValidateCode garantiza el uso de minúsculas, números y guiones bajos con el separador ':'.
func ValidateCode(code string) error {
	trimmed := strings.TrimSpace(code)
	if trimmed == "" {
		return ErrCodeRequired
	}

	// Regex estricto: snake_case para módulo y acción, separados por ':'
	var validCode = regexp.MustCompile(`^[a-z0-9_]+:[a-z0-9_]+$`)
	if !validCode.MatchString(trimmed) {
		return ErrCodeInvalidFormat
	}
	return nil
}

// ValidateDescription asegura que exista una explicación clara del permiso.
func ValidateDescription(description string) error {
	if strings.TrimSpace(description) == "" {
		return ErrDescriptionRequired
	}
	return nil
}

// ValidateModule verifica que el permiso esté categorizado en un módulo funcional.
func ValidateModule(module string) error {
	if strings.TrimSpace(module) == "" {
		return ErrModuleRequired
	}
	return nil
}

// ValidateStatus asegura que el estado esté dentro del rango definido (0-2).
func ValidateStatus(status Status) error {
	if status < Inactive || status > Deleted {
		return ErrInvalidStatus
	}
	return nil
}
