---
title: Descargas
slug: reportes-descargas
section: 16-billing
source: https://developers.mercadolibre.com.ve/es_ar/reportes-descargas
captured: 2026-06-06
---

# Descargas

> Source: <https://developers.mercadolibre.com.ve/es_ar/reportes-descargas>

## Endpoints referenced

- `https://api.mercadolibre.com/billing/integration/legal_document/$FILE_ID`
- `https://api.mercadolibre.com/billing/integration/legal_document/1234_FE_MEPF00869625_pdf`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/reports`
- `https://api.mercadolibre.com/billing/integration/periods/key/2021-08-01/reports`
- `https://api.mercadolibre.com/billing/integration/reports/$FILE_ID`
- `https://api.mercadolibre.com/billing/integration/reports/$FILE_ID/status`
- `https://api.mercadolibre.com/billing/integration/reports/ML-report-BILL-2021-08-04-11119999-CSV-v2/status?document_type=BILL`
- `https://api.mercadolibre.com/billing/integration/reports/ML-report-BILL-2021-08-04-11119999-CSV-v2?document_type=BILL`

## Content

Última actualización 25/04/2024

## Descarga de documento legal

Te compartimos un paso a paso para descargar las facturas legales de Mercado Libre y Mercado Pago en formato PDF. Debes obtener el file_id consultando el endpoint de /documents. En algunos casos, puedes recibir dos file_id. Solo en estos casos hay que tener en cuenta debes utilizar el file_id correspondiente al PDF. 
 Si el endpoint devuelve un único file_id, entonces debes utilizar ese dato para la descarga del documento legal.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/legal_document/$FILE_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/legal_document/1234_FE_MEPF00869625_pdf
```

Respuesta: **descarga del archivo en formato PDF**.

## Descarga reporte de conciliación

Para descargar los reportes de conciliación de **Mercado Libre**, **Mercado Envíos Flex**, **Insurtech** y **Mercado Pago** en los formatos **XLSX** y **CSV** hay que seguir un proceso de generación del reporte. Ese mismo proceso se aplica al reporte de **Fulfillment y al reporte de Pagos**, que solo admiten el formato de descarga **XLSX**.
 El proceso de generación del reporte consiste en tres pasos:

1. **Generación del reporte**
2. **Estado de generación de reporte**
3. **Descarga del reporte**
