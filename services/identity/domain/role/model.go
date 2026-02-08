package role

import "axis/shared/models"

// Status define el ciclo de vida de un rol dentro del sistema de seguridad.
type Status int

const (
	Inactive Status = iota // 0: El rol no otorga permisos aunque esté asignado
	Active                 // 1: Rol plenamente operativo
	Deleted                // 2: Borrado lógico para integridad histórica
)

// Role representa un perfil de acceso que agrupa permisos granulares.
// Un rol puede ser asignado a un usuario directamente o a un grupo.
type Role struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null;index"`
	Description string `json:"description" gorm:"not null"`

	// Permissions contiene los códigos granulares (ej: "sales:update").
	// Se usa jsonb para permitir flexibilidad y consultas eficientes en PostgreSQL.
	Permissions []string `json:"permissions" gorm:"type:jsonb"`

	// Status determina si el rol está activo, inactivo o borrado.
	Status Status `json:"status" gorm:"default:1"`

	// Embebemos models.Audit para heredar CreatedAt, UpdatedAt y DeletedAt (*time.Time).
	models.Audit
}
