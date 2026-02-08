# Domain: Permission (Identity Service)

El dominio `Permission` representa la unidad mínima de autorización dentro
de **Axis ERP**. Define acciones específicas que pueden ser ejecutadas
en los diferentes módulos del sistema.

## Responsabilidades
- Definir identificadores técnicos únicos (`Code`) para cada acción.
- Agrupar acciones por áreas funcionales (`Module`).
- Garantizar la consistencia de los códigos mediante reglas estrictas.

## Estructura de Datos
1. **Code**: Identificador único con formato `modulo:accion` (snake_case).
2. **Module**: Categoría funcional para organización en la interfaz.
3. **Status**: Ciclo de vida (Active, Inactive, Deleted).
4. **Audit**: Metadatos de trazabilidad (CreatedAt, UpdatedAt, DeletedAt).

## Reglas de Negocio
- **Formato de Código**: Debe cumplir con el Regex `^[a-z0-9_]+:[a-z0-9_]+$`.
- **Snake Case**: Solo se permiten minúsculas, números y guiones bajos.
- **Descripción**: Documentación obligatoria del alcance del permiso.
- **Integridad**: El borrado es siempre lógico para preservar históricos.

## Archivos del Dominio
- `model.go`: Definición de la estructura y constantes de estado.
- `rules.go`: Validación maestra y lógica de formato técnico (Regex).
- `actions.go`: Métodos para gestión de estado y actualización de datos.
- `validator.go`: Punto de entrada único para validaciones de entidad.

## Ejemplo de Uso
```go
import "axis/services/identity/domain/permission"

// El código asegura consistencia técnica y de auditoría
p := &permission.Permission{
    Code:   "billing:generate_invoice",
    Module: "Facturación",
}

if err := permission.ValidateEntity(p); err != nil {
    return err
}

