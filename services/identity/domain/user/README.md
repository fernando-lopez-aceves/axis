# 👤 User Domain Module (Axis Identity)

Este módulo centraliza la identidad y el control de acceso en Axis.

## 🏗️ Arquitectura
El dominio se divide en cuatro capas para facilitar el mantenimiento:
- **Model**: Estructura de datos, etiquetas GORM y estados (Iota).
- **Rules**: Validaciones de negocio atómicas y constantes de error.
- **Validator**: Punto de entrada para validar la integridad total.
- **Actions**: Métodos de mutación de estado y gestión de colecciones.

## 🚦 Estados (Status - Iota)
Para optimizar la DB, usamos tipos enteros (SmallInt):
- `0: Inactive`: Creado, pendiente de verificación.
- `1: Active`: Acceso total habilitado.
- `2: Suspended`: Bloqueo temporal por seguridad o administración.
- `3: Terminated`: Baja definitiva (se mantiene para histórico).
- `4: Deleted`: Borrado lógico (ignorado en consultas estándar).

## 🔐 Seguridad (Requerimiento 2026-02-05)
El modelo de autorización es híbrido y acumulativo:
1. **Roles**: Perfiles directos asignados al usuario.
2. **Groups**: Permisos heredados por pertenencia organizacional.
3. **Permissions**: Excepciones granulares (formato `modulo:accion`).

## 🛠️ Restricciones del Dominio
- **Username**: Longitud mínima de 5 caracteres, sin espacios.
- **Email**: Normalización automática a minúsculas y validación RFC 5322.
- **Integridad**: No se permiten strings vacíos en Roles, Groups o Permisos.

