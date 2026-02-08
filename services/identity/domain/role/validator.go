package role

// ValidateEntity orquesta la validación completa del objeto Role.
// Se mantiene como punto de entrada único para que la capa de aplicación (servicios)
// no necesite conocer las reglas atómicas internas de rules.go.
func ValidateEntity(role *Role) error {
	// Llamada a la lógica de integridad total definida en las reglas del dominio.
	if err := ValidateFullIntegrity(role); err != nil {
		return err
	}

	// Espacio reservado para validaciones que requieran contexto externo en el futuro,
	// como verificar duplicidad de nombres contra la base de datos o límites globales.

	return nil
}
