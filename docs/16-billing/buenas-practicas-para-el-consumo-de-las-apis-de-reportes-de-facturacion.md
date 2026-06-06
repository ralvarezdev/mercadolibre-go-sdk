---
title: Buenas Prácticas para el Consumo de las APIs de Reportes de Facturación
slug: buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion
section: 16-billing
source: https://developers.mercadolibre.com.ve/es_ar/buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion
captured: 2026-06-06
---

# Buenas Prácticas para el Consumo de las APIs de Reportes de Facturación

> Source: <https://developers.mercadolibre.com.ve/es_ar/buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion>

## Content

Última actualización 27/02/2026

## Buenas Prácticas para el Consumo de las APIs de Reportes de Facturación

Este es el resumen de los recursos disponibles para la integración de reportes de facturación de Mercado Libre y Mercado Pago.

## Información General

**Funcionalidad:** Permite conocer los reportes de facturación, documentos asociados (facturas y notas de crédito) y el resumen de cargos y bonificaciones de los vendedores.

**Parámetro Global Obligatorio:** Todos los endpoints requieren el parámetro `group` para especificar el grupo de facturación: **ML** (Mercado Libre) o **MP** (Mercado Pago). Si no se especifica, se obtiene información de ambos grupos.

## Propósito de los Recursos

Los recursos de **Reportes de Facturación** están **estrictamente destinados a operaciones de Post-Venta** y tienen como única finalidad la **conciliación fiscal y generación de reportes de facturación**.

Estos endpoints **no deben utilizarse como fuente de datos primaria** para gestión de ventas, seguimiento de pedidos en tiempo real, o cualquier otra finalidad operacional. Para esas necesidades, utiliza los recursos apropiados listados en la sección "Recursos Alternativos".

## Listado de Endpoints

### [1. Obtener Períodos de Facturación](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#obtener-periodo)

```javascript
GET /billing/integration/monthly/periods
```

Recupera información de los períodos de facturación. Por defecto entrega los últimos 6, con un máximo de 12 mediante paginación.

- **Parámetro obligatorio:** `document_type` (Valores: `Bill` o `Credit_note`)
- **Parámetros opcionales:** `offset`, `limit`

### [2. Obtener Documentos de un Período](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#obtener-documentos)

```javascript
GET /billing/integration/periods/key/{key}/documents
```

Permite obtener el listado de facturas y notas de crédito para un período específico, identificado por su `{key}` (primer día del mes, ej: `2024-01-01`).

- **Parámetros opcionales:** `document_id`, `document_type` (BILL, CREDIT_NOTE), `offset`, `limit`

### [3. Resumen de Facturación](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#resumen-facturacion)

```javascript
GET /billing/integration/periods/key/{key}/summary/details
```

Provee el resumen de los cargos, bonificaciones e impuestos aplicados al vendedor en un período determinado.

**Restricción de Uso:** No se recomienda su uso en procesamientos masivos (batch). Su uso debe ser secuencial y se recomienda una consulta diaria por usuario, ya que la información es estática durante el día.

### 4. Obtener Detalles de Facturación

#### [Mercado Libre](https://developers.mercadolibre.com.ar/es_ar/provisiones#mercado-libre)

```javascript
GET /billing/integration/periods/key/{KEY}/group/ML/details
```

Recupera el detalle completo de cargos por venta, bonificaciones e información de envíos asociados a las ventas de Mercado Libre para un período específico.

- **Parámetro obligatorio:** `document_type` (Valores: `BILL` o `CREDIT_NOTE`)
- **Parámetros opcionales:** `limit`, `from_id`, `sort_by`, `order_by`, `date_sort`, `detail_type`, `detail_sub_types`, `marketplace_type`, `order_ids`, `item_ids`, `document_ids`, `detail_ids`

#### [Mercado Pago](https://developers.mercadolibre.com.ar/es_ar/provisiones#mercado-pago)

```javascript
GET /billing/integration/periods/key/{KEY}/group/MP/details
```

Recupera el detalle de cargos y movimientos de la cuenta de Mercado Pago, incluyendo información de medios de pago, sucursales y referencias externas.

- **Parámetro obligatorio:** `document_type` (Valores: `BILL` o `CREDIT_NOTE`)
- **Parámetros opcionales:** `limit`, `from_id`, `sort_by`, `order_by`, `detail_type`

## Conciliación General

Para realizar una conciliación efectiva, los pasos propuestos son los siguientes:

1. **Consultar los períodos de facturación** para visualizar los períodos disponibles y el status de cada uno.
2. **Consultar el resumen de facturación**, a partir del período de facturación obtenido en el paso (1).
3. **Consultar el reporte de detalles del período**, a partir del período de facturación obtenido en el paso (1).

**Objetivo de la conciliación:** Los cargos expuestos en la factura y en el resumen de facturación deben coincidir con la sumatoria de cargos del mismo tipo que se muestran en el detalle de facturación, para un mismo período de facturación.

#### Ejemplo - Resumen de facturación:

```javascript
"charges": [
  {
    "label": "Campañas de publicidad - Product Ads",
    "amount": 48600,
    "type": "PADS",
    "groupId": 24
  },
  {
    "label": "Cargo por Mercado Envíos",
    "amount": 11195255.36,
    "type": "CXD",
    "groupId": 24
  },
  {
    "label": "Cargo por venta",
    "amount": 131285530.48,
    "type": "CV",
    "groupId": 28
  }
]
```

#### Ejemplo - Detalle de facturación:

```javascript
"results": [
  {
    "charge_info": {
      "legal_document_number": "0011A11111111",
      "legal_document_status": "PROCESSED",
      "legal_document_status_description": "Procesado",
      "creation_date_time": "2023-11-19T00:00:30",
      "detail_id": 12345678,
      "transaction_detail": "Cargo por vender",
      "debited_from_operation": "YES",
      "debited_from_operation_description": "Si",
      "status": null,
      "status_description": null,
      "charge_bonified_id": null,
      "detail_amount": 615.95,
      "detail_type": "CHARGE",
      "detail_sub_type": "CV"
    }
  }
]
```

**Regla de Conciliación:** En el endpoint de detalle de facturación, la suma de los `detail_amount` para un mismo `detail_sub_type` debe ser igual al `amount` para el mismo `type` en el endpoint de Resumen de facturación. Se sugiere para tareas de conciliación la utilización del filtro `detail_sub_types`.

## Frecuencia de Consulta de los Endpoints

La frecuencia de consulta recomendada depende del **status del Período de facturación**:

| Status del Período | Comportamiento | Frecuencia Recomendada |
| --- | --- | --- |
| **Abierto** | El resumen y el detalle de facturación varían diariamente a medida que se van generando los cargos. | En diferentes instancias del día. Por ejemplo: una vez al comienzo del día y una vez al finalizar. |
| **Cerrado** | El resumen y el detalle de facturación no varían. Pueden existir excepciones: devoluciones por cancelación de ventas y generación del PDF (cada site tiene una cantidad de días hábiles para disponibilizar el PDF una vez que cierra el período). | Una vez al día. Una vez que se cuente con todos los documentos fiscales, puede hacerse consultas periódicas para validar que no haya entrado algún bonus que afecte el total del período. |

- **Períodos de facturación:** Pueden cambiar de estado (ej: de abierto a cerrado). Frecuencia recomendada: **una vez por día**.
- **Documentos:** Se crean una vez cerrado el período. Cada site cuenta con días hábiles de tolerancia para disponibilizar el PDF. Frecuencia recomendada: **una vez por día**, con la mejora de realizar estas consultas al comienzo del período y desestimar en la parte final.

## Cómo Paginar Correctamente (Clave para Evitar Duplicados)

Para manejar grandes volúmenes de datos y garantizar que no se repitan registros entre llamadas, se debe utilizar una **paginación basada en IDs (`from_id`)** en lugar de solo desplazamientos (`offset`).

#### Parámetros Esenciales

| Parámetro | Descripción | Valores |
| --- | --- | --- |
| `limit` | Cantidad de registros por página | Mínimo: 1, Máximo: 1000, Default: 150 |
| `from_id` | El ID a partir del cual buscar | Default: 0 |
| `sort_by` | Propiedad de ordenamiento | `ID` o `DATE`. **Se recomienda ID** para paginar. |
| `order_by` | Orientación del ordenamiento | `ASC` o `DESC` |

#### Estrategia Recomendada

1. **Primera página:** Envía `limit=1000` y `from_id=0`.
2. **Siguientes páginas:** Obtén el valor del campo `last_id` de la respuesta JSON anterior y envíalo en el parámetro `from_id` de la nueva solicitud.
3. **Repetición:** Continúa hasta que la respuesta no devuelva más resultados.

```javascript
// Primera página
GET .../details?limit=1000&from_id=0&sort_by=ID&order_by=ASC

// Segunda página (usa el last_id de la respuesta anterior)
GET .../details?limit=1000&from_id={last_id}&sort_by=ID&order_by=ASC

// Continúa hasta que no haya más resultados...
```

**Evita el uso de `offset`** si tienes más de 10,000 registros, ya que este parámetro tiene un límite máximo de 10,000. El método `from_id` es el único que garantiza la integridad total en listados extensos.

## Estrategia de Consumo Recomendada

### 1. Evita Procesamiento en Batch Masivo

No realices solicitudes en paralelo masivo (batch) para obtener información de facturación. Los endpoints de billing **no fueron diseñados para consumo en alta frecuencia**. Almacena los resultados y pagina usando `from_id`.

INCORRECTO

Múltiples llamadas en batch:

```javascript
for seller in all_sellers:
    for order in seller.orders:
        GET /billing/integration/group/ML/order/details?order_ids={order}
```

CORRECTO

Una llamada por seller, una vez al día:

```javascript
GET /billing/integration/periods/key/{key}/group/ML/details?limit=1000&from_id=0
```

### 2. Implementa Caché Local

Dado que los datos se actualizan según el status del período, es **extremadamente recomendado** implementar una estrategia de caché:

- Almacena los datos consultados en tu base de datos
- Define una política de actualización basada en el status del período (abierto/cerrado)
- Antes de hacer una nueva solicitud, verifica si ya tienes los datos actualizados

### 3. No Utilices /monthly/periods en Batch

El endpoint `/billing/integration/monthly/periods` retorna información que **rara vez cambia**. La key del período es siempre el **primer día del mes** (ej: `2024-01-01`). No necesitas consultar `/monthly/periods` repetidamente. Construye la key directamente usando el primer día del mes deseado: `YYYY-MM-01`.

## Recursos Alternativos para Necesidades Operacionales

| Necesidad | Recurso Recomendado |
| --- | --- |
| Datos del pedido en tiempo real | [GET /orders](https://developers.mercadolibre.com.ar/es_ar/gestiona-ventas) |
| Identificar orders en packs | [GET /packs](https://developers.mercadolibre.com.ar/es_ar/gestion-packs) |
| Costo de envío | [GET /shipments](https://developers.mercadolibre.com.ar/es_ar/envios) |
| Descuentos aplicados | [GET /orders/{id}/discounts](https://developers.mercadolibre.com.ar/es_ar/gestiona-ventas#Obtener-descuentos-aplicados) |
| Precio de venta de un ítem | [GET /items/{item_id}/sale_price](https://developers.mercadolibre.com.ar/es_ar/api-de-precios#Obtener-precios-del-%C3%ADtem) |

## Dónde Obtener Cada Información

| Información Necesaria | Endpoint | Observación |
| --- | --- | --- |
| [Períodos de facturación](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#obtener-periodo) | `/billing/integration/monthly/periods` | Consulta **una vez** para obtener historial. Las keys siguen el patrón `YYYY-MM-01` |
| [Documentos (facturas/notas de crédito)](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#obtener-documentos) | `/billing/integration/periods/key/{key}/documents` | Filtra por `group` (ML/MP) y `document_type` |
| [Resumen de facturación](https://developers.mercadolibre.com.ar/es_ar/reportes-de-facturacion#resumen-facturacion) | `/billing/integration/periods/key/{key}/summary/details` | **No usar en batch**. Consumo secuencial, una vez al día |
| [Detalles de provisiones ML](https://developers.mercadolibre.com.ar/es_ar/provisiones#mercado-libre) | `/billing/integration/periods/key/{key}/group/ML/details` | Usa paginación con `limit` y `from_id`. Requiere `document_type` |
| [Detalles de provisiones MP](https://developers.mercadolibre.com.ar/es_ar/provisiones#mercado-pago) | `/billing/integration/periods/key/{key}/group/MP/details` | Usa paginación con `limit` y `from_id`. Requiere `document_type` |
| [Detalles por Order/Pack](https://developers.mercadolibre.com.ar/es_ar/provisiones#reportes-de-facturaci%C3%B3n-por-orders-y-packs) | `/billing/integration/group/ML/order/details` | **Consulta solo orders que aún no procesaste** |
| [Reportes de pagos](https://developers.mercadolibre.com.ar/es_ar/reportes-pagos) | `/billing/integration/periods/key/{key}/group/ML/payment/details` | Detalles de facturas abonadas |
| [Descarga de documento legal](https://developers.mercadolibre.com.ar/es_ar/reportes-descargas) | `/billing/integration/legal_document/{file_id}` | Obtén el `file_id` de `/documents` |
| [Descarga de reporte (CSV/XLSX)](https://developers.mercadolibre.com.ar/es_ar/reportes-descargas#Descarga-del-reporte-de-conciliaci%C3%B3n) | `/billing/integration/reports/{file_id}` | Requiere creación previa vía POST |
| [Percepciones (Argentina)](https://developers.mercadolibre.com.ar/es_ar/resumen-percepciones) | `/billing/integration/periods/key/{key}/perceptions/summary` | Exclusivo para MLA |

## Endpoints con Alta Incidencia de Error 429

Los siguientes endpoints son los más afectados por uso inadecuado:

| Endpoint | Causa Común del Problema |
| --- | --- |
| `/billing/integration/group/ML/order/details` | Consultas repetitivas por order_id o pack_id ya procesados |
| `/billing/integration/monthly/periods` | Llamadas en batch innecesarias o polling excesivo |

## Tratamiento de Errores

| Código | Tipo | Acción Recomendada |
| --- | --- | --- |
| 206 | Partial Content | Algunos datos están incompletos. Espera e intenta nuevamente más tarde (próximo ciclo de actualización). |
| 429 | Too Many Requests | **Bloqueo preventivo por IP.** Revisa tu implementación: reduce la frecuencia de llamadas, implementa caché y evita batch masivo. |

## Resumen de Buenas Prácticas

- Usa estos recursos **solo para conciliación fiscal** - no como fuente de datos operacionales
- Siempre especifica el parámetro `group` (ML o MP) para optimizar las consultas
- Ajusta la frecuencia de consulta según el **status del período** (abierto o cerrado)
- **Almacena en caché** toda la información consultada
- Usa **paginación basada en `from_id`** - evita `offset` para más de 10,000 registros
- Construye la key del período directamente (`YYYY-MM-01`) en lugar de consultar `/monthly/periods` repetidamente
- **Evita batch masivo** y llamadas paralelas en alta frecuencia
- **No consultes el mismo order/pack** más de una vez
- Para conciliación, usa el filtro `detail_sub_types` y valida que la suma de los `detail_amount` coincida con el `amount` del resumen
