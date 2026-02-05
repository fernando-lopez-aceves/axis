# 🧠 SERVICIO DE INTELIGENCIA (AXIS ANALYTICS)
--------------------------------------------------------------------------------
El motor predictivo de Axis que transforma datos crudos en decisiones.

## 🎯 RESPONSABILIDADES
- Predicción de demanda de stock para evitar quiebres en `Warehouse`.
- Análisis de comportamiento de clientes para el servicio de `Sales`.
- Detección de anomalías y fraudes en transacciones de `Finance`.

## 🛠️ STACK TECNOLÓGICO
- **Core**: Python (FastAPI para la comunicación).
- **Modelos**: Scikit-learn / TensorFlow para proyecciones.
- **Procesamiento**: Pandas/NumPy para limpieza de datos masivos.

## 🔄 FLUJO DE DATOS
1. Go recolecta transacciones en tiempo real.
2. Los datos se envían a este servicio para re-entrenar modelos.
3. Axis devuelve sugerencias de compra o alertas de riesgo.
--------------------------------------------------------------------------------

