# Domain: Group (Identity Service)

El dominio `Group` gestiona las agrupaciones de identidades en **Axis ERP**.
Su propósito es la administración masiva de accesos mediante herencia.

## Responsabilidades
- Gestionar pertenencia de roles a nivel colectivo.
- Administrar permisos granulares directos para el grupo.
- Mantener el ciclo de vida (Activo, Inactive, Deleted).

## Estructura de Datos
Bajo el requerimiento `2026-02-05`, el modelo implementa:
1. **Roles**: Roles heredados por todos los miembros del grupo.
2. **Permisos**: Permisos directos (`modulo:accion`) adicionales.
3. **Auditoría**: Integración nativa con `models.Audit`.

## Reglas de Negocio
- **Nombre**: Único, obligatorio (min. 3 caracteres).
- **Descripción**: Obligatoria (min. 10 caracteres).
- **Formato Permisos**: Patrón estricto `sistema:operacion`.
- **Estados**:
    - `0 (Inactive)`: No propaga permisos.
    - `1 (Active)`: Operativo.
    - `2 (Deleted)`: Borrado lógico.

## Archivos
- `model.go`: Estructura, Iota de Status y etiquetas GORM.
- `rules.go`: Validaciones atómicas y lógica interna.
- `actions.go`: Métodos para modificar el estado (Add/Remove).
- `validator.go`: Orquestador para capas externas.

## Uso sugerido
```go
import "axis/services/identity/domain/group"

// Validar antes de persistir
if err := group.ValidateEntity(newGroup); err != nil {
    return err
}

// Borrado lógico
newGroup.Delete()

