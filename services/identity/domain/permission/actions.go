package permission

import "time"

// UpdateBasicInfo permite modificar la metadata del permiso manteniendo la trazabilidad.
func (p *Permission) UpdateBasicInfo(code string, description string, module string) {
	p.Code = code
	p.Description = description
	p.Module = module
	p.UpdatedAt = time.Now()
}

// Activate habilita el permiso para que pueda ser asignado a nuevos roles o grupos.
func (p *Permission) Activate() {
	p.Status = Active
	p.UpdatedAt = time.Now()
}

// Deactivate deshabilita el permiso, impidiendo su uso en validaciones futuras.
func (p *Permission) Deactivate() {
	p.Status = Inactive
	p.UpdatedAt = time.Now()
}

// Delete realiza el borrado lógico del permiso para preservar la integridad referencial.
func (p *Permission) Delete() {
	p.Status = Deleted
	p.UpdatedAt = time.Now()

	// Sincronización con el puntero de auditoría en shared/models.
	now := time.Now()
	p.DeletedAt = &now
}
