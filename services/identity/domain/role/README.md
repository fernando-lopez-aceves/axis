# Domain: Role (Identity Service)

El dominio `Role` representa perfiles de acceso que agrupan permisos
granulares. Es el puente entre las acciones atómicas (`Permission`)
y los sujetos (`User` o `Group`).

## Responsabilidades
- Agrupar permisos específicos (ej. "invoice:create", "invoice:read").
- Gestionar el ciclo de vida de los perfiles de acceso.
- Validar que no se excedan los límites de seguridad definidos.

## Estructura de Datos
Implementa la lógica de seguridad bajo el requerimiento `2026-02-05`:
1. **Permissions**: Un slice de strings (`jsonb`) con códigos granulares.
2. **Status**: Control de estado (Active, Inactive, Deleted).
3. **Auditoría**: Integración con `models.Audit` para histórico.

## Reglas de Negocio
- **Nombre**: Único y obligatorio (min. 3 caracteres).
- **Descripción**: Obligatoria para fines de auditoría clara.
- **Límite de Permisos**: Máximo 100 permisos por rol para evitar el
  "Privilege Creep" y optimizar el rendimiento.
- **Formato**: Los permisos deben seguir el patrón `modulo:accion`.

## Archivos del Dominio
- `model.go`: Estructura del rol y constantes de estado.
- `rules.go`: Validaciones de formato, longitud y límites de permisos.
- `actions.go`: Métodos para añadir/quitar permisos y borrado lógico.
- `validator.go`: Orquestador de validación para la capa de servicio.

## Uso sugerido
```go
import "axis/services/identity/domain/role"

// Validar integridad antes de guardar
if err := role.ValidateEntity(newRole); err != nil {
    return err
}

// Agregar permiso de forma segura
currentRole.AddPermission("sales:view")

