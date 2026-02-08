package group

import "axis/shared/models"

// Status define los estados operativos de un grupo mediante un enumerado basado en iota.
type Status int

const (
	Inactive Status = iota // 0: El grupo existe pero sus permisos no aplican
	Active                 // 1: Grupo plenamente operativo
	Deleted                // 2: Borrado lógico para auditoría histórica
)

// Group permite la gestión masiva de accesos mediante la tríada de herencia:
// Roles y Permisos asignados a un conjunto de usuarios.
type Group struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"unique;not null;index"`
	Description string `json:"description" gorm:"not null"`

	// Roles que heredarán todos los usuarios pertenecientes a este grupo.
	// Se almacena como jsonb para flexibilidad en la base de datos.
	Roles []string `json:"roles" gorm:"type:jsonb"`

	// Permisos granulares específicos (ej: "warehouse:create").
	Permissions []string `json:"permissions" gorm:"type:jsonb"`

	// Status representa el ciclo de vida del grupo (0=Inactive, 1=Active, 2=Deleted).
	Status Status `json:"status" gorm:"default:0"`

	// Embebemos models.Audit para heredar CreatedAt, UpdatedAt y DeletedAt (*time.Time).
	models.Audit
}
