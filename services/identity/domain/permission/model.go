package permission

import "axis/shared/models"

// Status define si el permiso está disponible para ser asignado en el sistema.
type Status int

const (
	Inactive Status = iota // 0: El permiso no puede ser asignado ni validado
	Active                 // 1: Permiso operativo
	Deleted                // 2: Borrado lógico para integridad referencial
)

// Permission representa la unidad mínima de autorización (ej: "orders:delete").
// Se organiza por módulos para facilitar la administración en la interfaz de Axis ERP.
type Permission struct {
	ID          string `json:"id" gorm:"primaryKey"`
	Code        string `json:"code" gorm:"unique;not null;index"` // Formato: "modulo:accion"
	Description string `json:"description" gorm:"not null"`
	Module      string `json:"module" gorm:"index"` // Agrupador funcional (ej: "Sales")

	// Status indica si el permiso está activo o ha sido revocado/borrado.
	Status Status `json:"status" gorm:"default:1"`

	// Embebemos models.Audit para heredar CreatedAt, UpdatedAt y el puntero DeletedAt.
	models.Audit
}
