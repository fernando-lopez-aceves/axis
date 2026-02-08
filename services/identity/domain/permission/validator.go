package permission

// ValidateEntity es el punto de entrada principal para validar la integridad
// de un permiso antes de ser procesado por la capa de infraestructura.
func ValidateEntity(permission *Permission) error {
	// Delegamos la lógica en la función maestra de rules.go
	if err := ValidateFullIntegrity(permission); err != nil {
		return err
	}

	// Espacio reservado para validaciones de contexto que no pertenecen a
	// reglas de formato, como verificar permisos contra un catálogo maestro.

	return nil
}
