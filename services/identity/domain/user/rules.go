package user

import (
	"errors"
	"net/mail"
	"regexp"
	"strings"
)

var (
	ErrInvalidUsername    = errors.New("user_rule: el nombre de usuario debe tener al menos 5 caracteres")
	ErrUsernameCharacters = errors.New("user_rule: el nombre de usuario solo puede contener letras, números y guiones bajos")
	ErrInvalidEmail       = errors.New("user_rule: el formato del correo electrónico no es válido")
	ErrEmptyPassword      = errors.New("user_rule: la contraseña no puede estar vacía")
	ErrInvalidStatus      = errors.New("user_rule: el estado del usuario no es válido")
	ErrEmptyGroups        = errors.New("user_rule: un usuario debe pertenecer al menos a un grupo")
	ErrInvalidGroupID     = errors.New("user_rule: se detectó un ID de grupo vacío")
	ErrInvalidRoleID      = errors.New("user_rule: se detectó un ID de rol vacío")
	ErrInvalidPermission  = errors.New("user_rule: formato de permiso no válido (debe ser modulo:accion)")
)

// --- VALIDACIONES INDIVIDUALES ---
// Definimos el patrón: Solo letras (a-z, A-Z), números (0-9) y guiones bajos (_)
var usernameRegex = regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)

func ValidateUsername(username string) error {
	trimmed := strings.TrimSpace(username)

	if len(trimmed) < 5 {
		return ErrInvalidUsername
	}

	if !usernameRegex.MatchString(trimmed) {
		return ErrUsernameCharacters
	}

	return nil
}

func ValidateEmail(emailAddr string) error {
	_, err := mail.ParseAddress(emailAddr)
	if err != nil {
		return ErrInvalidEmail
	}
	return nil
}

func ValidateEmptyPassword(password string) error {
	if len(strings.TrimSpace(password)) == 0 {
		return ErrEmptyPassword
	}
	return nil
}

func ValidateStatus(status Status) error {
	if status < Inactive || status > Deleted {
		return ErrInvalidStatus
	}
	return nil
}

func ValidateGroups(groups []string) error {
	if len(groups) == 0 {
		return ErrEmptyGroups
	}
	for _, groupID := range groups {
		if strings.TrimSpace(groupID) == "" {
			return ErrInvalidGroupID
		}
	}
	return nil
}

func ValidateRoles(roles []string) error {
	for _, roleID := range roles {
		if strings.TrimSpace(roleID) == "" {
			return ErrInvalidRoleID
		}
	}
	return nil
}

func ValidatePermissionFormat(permission string) error {
	rawPermission := strings.TrimSpace(permission)
	if len(rawPermission) == 0 || !strings.Contains(rawPermission, ":") {
		return ErrInvalidPermission
	}
	return nil
}

func ValidatePermissionsList(permissions []string) error {
	for _, permissionCode := range permissions {
		if err := ValidatePermissionFormat(permissionCode); err != nil {
			return err
		}
	}
	return nil
}

// --- REGLA MAESTRA (LA INTEGRIDAD TOTAL) ---

func ValidateFullIntegrity(user *User) error {
	if err := ValidateUsername(user.Username); err != nil {
		return err
	}
	if err := ValidateEmail(user.Email); err != nil {
		return err
	}

	if user.ID == "" {
		if err := ValidateEmptyPassword(user.Password); err != nil {
			return err
		}
	}

	if err := ValidateStatus(user.Status); err != nil {
		return err
	}
	if err := ValidateGroups(user.Groups); err != nil {
		return err
	}
	if err := ValidateRoles(user.Roles); err != nil {
		return err
	}
	if err := ValidatePermissionsList(user.Permissions); err != nil {
		return err
	}
	return nil
}

func NormalizeEmail(emailAddr string) string {
	return strings.ToLower(strings.TrimSpace(emailAddr))
}
