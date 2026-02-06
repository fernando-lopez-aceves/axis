# Axis ERP 📦

Sistema profesional de gestión de inventarios, ventas y logística con un
enfoque multi-sucursal y multi-rol, diseñado bajo una arquitectura de
servicios desacoplados y alta disponibilidad.

## 🏗️ Arquitectura

Axis utiliza una estructura de **Servicios Especializados**. Cada módulo
en la carpeta `services/` es autónomo, conteniendo su propia lógica de
negocio e infraestructura, comunicándose a través de la capa `shared/`.

## 🌟 Reglas de Oro de Desarrollo (Engineering Standards)

1. **Nombres Descriptivos Obligatorios**: Está estrictamente prohibido el
   uso de variables de una sola letra (ej. `g`, `u`, `p`). Se deben usar
   nombres semánticos (ej. `groupID`, `user`, `permissionCode`).
2. **Triada de Seguridad (2026-02-05)**: El control de acceso es híbrido
   y acumulativo: Roles directos, Grupos y Permisos individuales.
3. **Eficiencia de Estado**: Se utiliza `Iota` para definir estados de
   entidades, optimizando el almacenamiento y la lógica de negocio.

## 📂 Estructura del Proyecto (Nivel 2)

```text
.
├── docs/                # Documentación técnica y de arquitectura
├── interfaces/          # Clientes de entrada (Web, Mobile, Terminal)
├── scripts/             # Automatización, migraciones y despliegue
├── services/            # Dominios de negocio (Identity, Sales, etc.)
│   ├── accounting/      # Contabilidad general
│   ├── audit/           # Trazabilidad y logs
│   ├── finance/         # Gestión financiera
│   ├── human-resources/ # Gestión de personal
│   ├── identity/        # Usuarios, Roles y Grupos
│   ├── intelligence/    # Reportes y analítica
│   ├── sales/           # Facturación y ventas
│   └── warehouse/       # Inventarios y logística
├── shared/              # Recursos compartidos entre servicios
│   ├── models/          # Modelos base (Audit, Metadata)
│   └── proto/           # Contratos de comunicación
└── go.mod               # Dependencias del proyecto

