---
title: Precios de productos
slug: api-de-precios
section: 12-products
source: https://developers.mercadolibre.com.ve/es_ar/api-de-precios
captured: 2026-06-06
---

# Precios de productos

> Source: <https://developers.mercadolibre.com.ve/es_ar/api-de-precios>

## Endpoints referenced

- `https://api.mercadolibre.com/items/$ITEM_ID/prices`
- `https://api.mercadolibre.com/items/$ITEM_ID/prices/standard`
- `https://api.mercadolibre.com/items/$ITEM_ID/sale_price?context=$CHANNEL,LOYALTY_LEVEL`
- `https://api.mercadolibre.com/items/MLA3191390879/sale_price?context=channel_marketplace,buyer_loyalty_3`
- `https://api.mercadolibre.com/items/MLM237323192/prices`

## Content

Última actualización 26/02/2026

## Precios de productos

 Importante: 

Utiliza los siguientes endpoints para consultar precios, ya que eliminaremos progresivamente: 

 - Los campos 

price, base_price

 y 

original_price

 de la API de /items.

Si publicas y sincronizas items debes consultar los precios relacionados a un producto y contexto (canal y/o nivel de comprador) con las siguientes APIs. Además, puedes conocer el valor exacto de venta. Para **[crear un item](https://developers.mercadolibre.com.ve/es_ar/publica-productos#Publica-un-articulo)** y **[editarlo](https://developers.mercadolibre.com.ve/es_ar/producto-sincroniza-modifica-publicaciones#Consideraciones-para-actualizar-items)** debes continuar haciéndolo mediante la API de /items. **Próximamente, habilitaremos este [endpoint para editar precios](https://developers.mercadolibre.com.ar/es_ar/api-de-precios?nocache=true#Editar-precios-tipo-standard)**.

## Notificaciones sobre precios

Para recibir notificaciones sobre los precios, debes suscribirte al tópico [**items_prices**](https://developers.mercadolibre.com.ar/es_ar/productos-recibe-notificaciones#items-prices), después de recibir la notificación debe consultar el recurso de **/prices**.

## Obtener precio de venta actual

Para conocer el precio de venta, debes saber que **los precios pueden ser de tipo standard** (precios por default sin promociones asociadas) o **promotion** (promocionales, el precio tiene una promoción). 

 Con el siguiente request identificas el precio de venta ganador de un producto y puedes filtrar precios por canal de venta y/o nivel del comprador. Además, recibirás información sobre las promociones asociadas al ítem con el precio ganador, que se muestra al comprador. **Si el token no pertenece al seller e ítem consultado**, no recibirás información del array metadata (tipo de promoción asociada). 

 **Context**: parámetro opcional para filtrar canal de venta y/o nivel de comprador. Te recomendamos enviar por lo menos un channel a la vez y opcional puedes agregar cada loyalty_level. Valores posibles:

**CHANNEL (Canal de venta)**

- **channel_marketplace**: canal Mercado Libre.
- **channel_proximity, mp_merchants y mp_links**: canal de productos publicados en Mercado Pago. Próximamente estarán habilitados.

**LOYALTY_LEVEL (Nivel del comprador)**: no disponible en MLU y MPE

- buyer_loyalty_3
- buyer_loyalty_4
- buyer_loyalty_5
- buyer_loyalty_6

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/sale_price?context=$CHANNEL,LOYALTY_LEVEL
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLA3191390879/sale_price?context=channel_marketplace,buyer_loyalty_3
```

Respuesta:

```javascript
{
    "price_id": "145",
    "amount": 39200.0,
    "regular_amount": 49000.0,
    "currency_id": "ARS",
    "reference_date": "2023-06-08T19:06:56Z",
    "metadata": {
        "promotion_id": "OFFER-MLA13456789-1122334455", 
        "promotion_type": "custom" 
    }
}
```

### Descripción de los campos

**price_id**: ID del precio.
 **amount**: precio de venta del producto.
 **regular_amount**: precio original del producto, en casos que tenga promoción. El precio tachado del precio ganador también es calculado, y se puede tomar de varias fuentes. No necesariamente será el mismo 'regular_amount' del recurso /prices.
 **currency_id**: ID de la moneda a la que se refiere el campo amount y regular_amount.
 **reference_date**: fecha para la cual está calculando el precio de venta.
 **metadata**: información privada del usuario relacionada a la promotion asociada.

 Nota: 

Si el ítem tiene un precio con promoción, los siguientes campos estarán relacionados a Promociones. Puedes utilizarlos como input para realizar consultas específicas sobre 

promociones de marketplace

. 

**promotion_id**: Id de la promoción. Con este ID puedes [consultar la oferta](https://developers.mercadolibre.com.ar/es_ar/central-de-promociones#Consultar-ofertas) (item, promoción y estado). .
 **promotion_type**: tipo de promoción.

### Cuando la promoción está activa:

- Y bajas el precio a un valor por debajo de la oferta, Mercado Libre elimina la promoción y queda el precio standard nuevo.
- Y bajas el precio pero no por debajo del valor de la promoción, esta seguirá activa independientemente del aumento del precio standard.
