package models

import "time"

// Audit es una estructura de composición (mixin) diseñada para ser embebida
// en las entidades del dominio de Axis. Su propósito es estandarizar la
// trazabilidad temporal y el ciclo de vida de los registros en la base de datos.
type Audit struct {
	// CreatedAt registra el instante exacto (UTC) en que el recurso fue persistido.
	// Se utiliza para reportes de creación y auditoría inicial.
	CreatedAt time.Time `json:"created_at" gorm:"index"`

	// UpdatedAt se actualiza automáticamente en cada modificación del registro.
	// Permite identificar la frescura de los datos y detectar colisiones de escritura.
	UpdatedAt time.Time `json:"updated_at"`

	// DeletedAt implementa el patrón "Soft Delete" (borrado lógico).
	// Si el valor no es nulo, el sistema considera al registro como eliminado,
	// pero los datos permanecen en la DB para integridad referencial e histórica.
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"index"`
}
