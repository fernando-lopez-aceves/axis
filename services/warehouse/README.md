# 📦 SERVICIO DE ALMACÉN (WAREHOUSE CORE)
--------------------------------------------------------------------------------
Este es el motor de control de existencias y logística física del ERP.

## 🎯 RESPONSABILIDADES
- Gestión de inventario multisucursal en tiempo real.
- Control de SKUs, lotes, fechas de vencimiento y números de serie.
- Procesamiento de entradas (compras) y salidas (ventas/mermas).

## 🚀 OPTIMIZACIÓN ROSS-GO
- El cálculo de algoritmos de ubicación (Sugerencia de estantes) y 
  valorización de inventario masivo se delega a módulos de Rust.
- La concurrencia de pedidos se gestiona con Goroutines.

