package user

import (
	"axis/shared/models"
	"time"
)

// Status define los estados vitales del usuario en el ecosistema Axis.
type Status int

const (
	Inactive   Status = iota // 0: Usuario creado pero sin verificar
	Active                   // 1: Usuario con acceso total
	Suspended                // 2: Bloqueado temporalmente (ej. mora o seguridad)
	Terminated               // 3: Baja administrativa (histórico)
	Deleted                  // 4: Borrado lógico
)

// User es la entidad central de seguridad.
// Los datos personales viven en el servicio de RRHH vinculados por el ID.
type User struct {
	ID       string `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null;index"`
	Email    string `json:"email" gorm:"unique;not null;index"`
	Password string `json:"-"` // Oculto siempre en JSON

	// Roles y Permisos (Requerimiento 05-02-2026)
	// Triada: Roles (directos), Grupos (heredados) y Permisos (excepciones).
	Roles       []string `json:"roles" gorm:"type:jsonb"`
	Groups      []string `json:"groups" gorm:"type:jsonb"`
	Permissions []string `json:"permissions" gorm:"type:jsonb"`

	Status Status `json:"status" gorm:"default:0"`

	// Seguimiento específico de acceso
	LastLogin time.Time `json:"last_login"`

	// Embebemos models.Audit para heredar CreatedAt, UpdatedAt y DeletedAt.
	// Nota: Asegúrate que models.Audit maneje DeletedAt como puntero o gorm.DeletedAt.
	models.Audit
}
