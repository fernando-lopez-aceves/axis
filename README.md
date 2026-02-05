# Axis ERP 📦

Sistema profesional de gestión de inventarios, ventas y logística con un 
enfoque multi-sucursal y multi-rol, diseñado bajo estándares de alta 
disponibilidad y código limpio.

## 🏗️ Arquitectura

Este proyecto implementa **Arquitectura Hexagonal (Puertos y Adaptadores)**.
Esta elección garantiza que la lógica de negocio permanezca pura e 
independiente de factores externos como la base de datos, el framework 
web o librerías de terceros.



## 📂 Estructura del Proyecto

```text
.
├── cmd/                         # Puntos de entrada (Web Server, CLI, Cron)
├── internal/
│   ├── core/
│   │   ├── domain/              # Lógica de negocio pura (Entidades)
│   │   │   ├── identity/        # Usuarios, Roles, Grupos y Permisos (RBAC)
│   │   │   ├── product/         # Catálogo, Precios, Lotes y Presentaciones
│   │   │   ├── movement/        # Entradas, Salidas, Traslados y Auditorías
│   │   │   ├── sales/           # Proceso de Ventas y Gestión de Clientes
│   │   │   └── masters/         # Sucursales, Unidades y Proveedores
│   │   └── ports/               # Interfaces (Contratos del sistema)
│   │
│   ├── adapters/                # Implementaciones técnicas externas
│   │   ├── repository/          # Persistencia de datos (Postgres, GORM)
│   │   ├── security/            # Implementación de JWT y Hashing
│   │   └── handlers/            # Controladores de transporte (Gin, HTTP)
│   │
│   └── platform/                # Código de infraestructura (DB, Logger)
│
├── pkg/                         # Librerías compartidas de utilidad general
├── scripts/                     # Scripts de automatización y migraciones
└── go.mod                       # Definición del módulo y dependencias
```
