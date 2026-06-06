---
title: Envíos Personalizados
slug: envios-personalizados
section: 13-shipping
source: https://developers.mercadolibre.com.ve/es_ar/envios-personalizados
captured: 2026-06-06
---

# Envíos Personalizados

> Source: <https://developers.mercadolibre.com.ve/es_ar/envios-personalizados>

## Endpoints referenced

- `https://api.mercadolibre.com/items/$ITEM_ID`
- `https://api.mercadolibre.com/items/$ITEM_ID/shipping_options?zip_code=$ZIP_CODE`
- `https://api.mercadolibre.com/items/MLA803066380/shipping_options?zip_code=$1234`
- `https://api.mercadolibre.com/shipments/$SHIPMENT_ID`

## Content

Última actualización 13/03/2026

## Envíos Personalizados

El envío personalizado permite a los vendedores gestionar la logística de sus ventas de manera propia, ofreciendo flexibilidad en los costos de envío para sus clientes.

- **Custom:**El vendedor carga una tabla con precios de envío específicos para cada región. Este se encarga completamente de la logística y entrega.
- **Not_Specified:** El vendedor no define precios de envío. El comprador debe contactar al vendedor para acordar los detalles y costos de envío. Solo para estos casos no hay shipment_id.
