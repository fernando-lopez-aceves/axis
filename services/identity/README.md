# 🔑 SERVICIO DE IDENTIDAD (IDENTITY CORE)
--------------------------------------------------------------------------------
Este servicio es el "Gatekeeper" de Axis. Gestiona la autenticación, 
autorización y perfiles de usuario para todo el ecosistema.

## 🎯 RESPONSABILIDADES
- Autenticación segura (Login/Logout).
- Gestión de Permisos Basados en Roles (RBAC).
- Emisión y validación de tokens (JWT/Paseto).
- Auditoría de accesos y seguridad de cuentas.

## 🛠️ ARQUITECTURA ROSS-GOLANG
- **Capa de Dominio (Go)**: Define la entidad User y las reglas de negocio.
- **Capa de Criptografía (Rust)**: Procesamiento de hashing (Argon2id) para 
  máxima seguridad y velocidad ante ataques de fuerza bruta.
- **API (gRPC)**: Provee validación de identidad ultra rápida a los demás 
  microservicios de Axis.

## 🗄️ MODELO DE DATOS
- Los usuarios se almacenan con estados: Activo, Inactivo o Bloqueado.
- Los roles definen el alcance en Warehouse, Sales, Finance y Accounting.
--------------------------------------------------------------------------------

