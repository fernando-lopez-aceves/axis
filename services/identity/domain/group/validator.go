package group

// ValidateEntity actúa como el orquestador principal de validaciones para la entidad Group.
// Se separa en este archivo para mantener la consistencia arquitectónica con el dominio de User.
func ValidateEntity(group *Group) error {
	// Llamamos a la función de integridad total definida en rules.go.
	// Esto asegura que el nombre, descripción, estado y la tríada de seguridad
	// (roles/permisos) sean válidos antes de cualquier operación de persistencia.
	if err := ValidateFullIntegrity(group); err != nil {
		return err
	}

	// Aquí se pueden añadir validaciones futuras que dependan de factores externos
	// o lógica que no pertenezca estrictamente a las reglas atómicas del dominio.

	return nil
}
