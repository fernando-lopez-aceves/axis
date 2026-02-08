package role

import "time"

// AddPermission agrega un nuevo código de permiso al rol, validando que no esté duplicado.
func (role *Role) AddPermission(permissionCode string) {
	for _, existing := range role.Permissions {
		if existing == permissionCode {
			return
		}
	}
	role.Permissions = append(role.Permissions, permissionCode)
	role.UpdatedAt = time.Now()
}

// RemovePermission elimina un permiso específico del slice de permisos del rol.
func (role *Role) RemovePermission(permissionCode string) {
	var updatedPermissions []string
	for _, existing := range role.Permissions {
		if existing != permissionCode {
			updatedPermissions = append(updatedPermissions, existing)
		}
	}
	role.Permissions = updatedPermissions
	role.UpdatedAt = time.Now()
}

// UpdateBasicInfo actualiza el nombre y la descripción, campos críticos para la auditoría.
func (role *Role) UpdateBasicInfo(name string, description string) {
	role.Name = name
	role.Description = description
	role.UpdatedAt = time.Now()
}

// Activate cambia el estado del rol a Active (1).
func (role *Role) Activate() {
	role.Status = Active
	role.UpdatedAt = time.Now()
}

// Deactivate cambia el estado del rol a Inactive (0), deshabilitando sus permisos.
func (role *Role) Deactivate() {
	role.Status = Inactive
	role.UpdatedAt = time.Now()
}

// Delete ejecuta el borrado lógico del rol utilizando el puntero de models.Audit.
func (role *Role) Delete() {
	role.Status = Deleted
	role.UpdatedAt = time.Now()

	// Sincronizado con metadata.go: asignamos la dirección de memoria del tiempo actual.
	now := time.Now()
	role.DeletedAt = &now
}
