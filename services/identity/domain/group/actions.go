package group

import "time"

// AddRole añade un rol al grupo si no existe previamente para evitar duplicados.
func (group *Group) AddRole(role string) {
	for _, existing := range group.Roles {
		if existing == role {
			return
		}
	}
	group.Roles = append(group.Roles, role)
	group.UpdatedAt = time.Now()
}

// RemoveRole elimina un rol específico de la lista del grupo.
func (group *Group) RemoveRole(role string) {
	var updated []string
	for _, existing := range group.Roles {
		if existing != role {
			updated = append(updated, existing)
		}
	}
	group.Roles = updated
	group.UpdatedAt = time.Now()
}

// AddPermission agrega un permiso granular al grupo.
func (group *Group) AddPermission(permission string) {
	for _, existing := range group.Permissions {
		if existing == permission {
			return
		}
	}
	group.Permissions = append(group.Permissions, permission)
	group.UpdatedAt = time.Now()
}

// RemovePermission quita un permiso específico del grupo.
func (group *Group) RemovePermission(permission string) {
	var updated []string
	for _, existing := range group.Permissions {
		if existing != permission {
			updated = append(updated, existing)
		}
	}
	group.Permissions = updated
	group.UpdatedAt = time.Now()
}

// UpdateBasicInfo actualiza los datos descriptivos del grupo.
func (group *Group) UpdateBasicInfo(name string, description string) {
	group.Name = name
	group.Description = description
	group.UpdatedAt = time.Now()
}

// Deactivate cambia el estado del grupo a Inactive (0).
func (group *Group) Deactivate() {
	group.Status = Inactive
	group.UpdatedAt = time.Now()
}

// Activate pone al grupo en estado operativo Active (1).
func (group *Group) Activate() {
	group.Status = Active
	group.UpdatedAt = time.Now()
}

// Delete marca el grupo como borrado lógico utilizando el puntero de models.Audit.
func (group *Group) Delete() {
	group.Status = Deleted
	group.UpdatedAt = time.Now()

	// Ajuste según tu metadata.go: Usamos un puntero a time.Now()
	now := time.Now()
	group.DeletedAt = &now
}
