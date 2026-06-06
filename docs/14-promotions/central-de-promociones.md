---
title: Gestionar promociones
slug: central-de-promociones
section: 14-promotions
source: https://developers.mercadolibre.com.ve/es_ar/central-de-promociones
captured: 2026-06-06
---

# Gestionar promociones

> Source: <https://developers.mercadolibre.com.ve/es_ar/central-de-promociones>

## Endpoints referenced

- `https://api.mercadolibre.com/seller-promotions/candidates/$CANDIDATE_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/candidates/CANDIDATE-`
- `https://api.mercadolibre.com/seller-promotions/exclusion-list/item?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/exclusion-list/seller/{item_id}?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/exclusion-list/seller?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLA1399846831?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/offers/$OFFERS_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/offers/OFFER-MLB1970246686-42701792?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTIONS_ID/items?promotion_type=$PROMOTION_TYPE&app_version=v2&limit=50&search_after={$SEARCH_AFTER}`
- `https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTIONS_ID/items?promotion_type=PROMOTIONS_TYPE&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID/items?promotion_type=$PROMOTION_TYPE&status=$STATUS&item_id=$ITEM_ID&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID?promotion_type=$PROMOTION_TYPE&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/MLA1111/items?promotion_type=DEAL&item_id=MLA604400000&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB13451002MLB9377/items?promotion_type=DEAL&app_version=v2&limit=50&search_after=d3e3fb02371ca8e49ceb3e998c4a1b8da189235497375e55d3027c7dacf5a4ef0175b2aaca4f4a0fdf31299947d82ba661037482172ba7f9cfb1b0250d3134aa71c889367aa7c1401e4c3ff5bd70ee14337dfaa18c99bbe5e52dc3a2c1b55b195131903ecbc60a1c639e01dbecf11b15126d4b38cdb6122182acde2eca1b1a42`
- `https://api.mercadolibre.com/seller-promotions/users/$USER_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/users/1356551933?app_version=v2`

## Content

Última actualización 13/03/2026

## Gestionar promociones

 Importante: 

A partir del 28/08, el límite máximo de descuento permitido pasa de 70% a 80%.

Este cambio aplica a todas las modalidades configurables por el seller: LIGHTNING, DOD, SELLER_CAMPAIGN, DEAL, PRICE_DISCOUNT

.

Con el recurso **/seller-promotions** puedes centralizar todos los tipos de promociones disponibles como **campañas tradicionales** (DEAL),**campañas co-fondeadas** por Mercado Libre (MARKETPLACE CAMPAIGN), **descuentos individuales** (PRICE DISCOUNT), **ofertas relámpago** (LIGHTNING), **ofertas del día** (DOD), **descuento por volumen** (VOLUME), **descuento pre-acordado por item** (PRE NEGOTIATED), **campaña del vendedor** (SELLER_CAMPAIGN), **campañas co-fondeada automatizada** (SMART),**campañas de precios competitivos** (PRICE_MATCHING), **campaña de liquidación stock Full** (UNHEALTHY_STOCK) y **campañas de cupones del vendedor** (SELLER_COUPON_CAMPAIGN). Además de los nuevos tipos de ofertas que disponibilicemos.

## Características de las promociones

| Nombre de la campaña | Tipo de campaña | Definición de precio | Sugerencia de precio | Bonificación MELI | Stock para participar | Deadline | Aprobación |
| --- | --- | --- | --- | --- | --- | --- | --- |
| **Tradiconal** | DEAL | Usuario define | No | No | No | Sí | Sí |
| **Co-fondeada** | MARKETPLACE CAMPAIGN | Usuario acepta | No | Sí | No | Sí | No |
| **Descuento por volumen** | VOLUME | Usuario acepta | No | Sí | No | Sí | No |
| **Oferta del día** | DOD | Usuario define | Sí | No | Sí, informativo | No | No |
| **Oferta relámpago** | LIGHTNING | Usuario define | Sí | No | Sí, mandatorio | No | No |
| **Descuento pre-acordado por ítem** | PRE_NEGOTIATED | Usuario acuerda y acepta | No | Sí | Sí | Sí | No |
| **Campaña del vendedor** | SELLER CAMPAIGN | Usuario define y acepta | No | No | No | Sí | No |
| **Campaña co-fondeada automatizada** | SMART | Usuario acepta | No | Sí | No | Sí | No |
| **Campaña de precios competitivos** | PRICE_MATCHING | Usuario acepta | No | Sí | No | Sí | No |
| **Campaña de liquidación stock Full** | UNHEALTHY_STOCK | Usuario acuerda y acepta | No | Sí | Sí | Sí | No |

## Disponibilidad por país

| Sitio | **Campañas tradicionales** (DEAL) | **Campaña co-fondeada** (MARKETPLACE CAMPAIGN) | **Descuento individual** (PRICE DISCOUNT) | **Descuento por volumen** (VOLUME) | **Descuento pre-acordado por ítem** (PRE_NEGOTIATED) | **Oferta del día** (DOD) | **Oferta relámpago** (LIGHTNING) | **Campaña co-fondeada automatizada** (SMART) | **Campaña de precios competitivos** (PRICE_MATCHING) | **Campaña de liquidación stock Full** (UNHEALTHY_STOCK) | **Campaña del vendedor** (SELLER_CAMPAIGN) |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| **MLA, MLB, MLM, MCO, MLC, MLU, MPE** |  |  |  |  |  |  |  |  |  |  |  |
| **MLV y MEC** |  |  |  |  |  |  |  |  |  |  |  |

 Nota: 

 La campañas de cupones del vendedor (SELLER_COUPON_CAMPAIGN) esta disponible solo para MLB.

## Promociones del vendedor

Recuerda que un usuario puede tener más de una invitación y de diferentes tipos.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/users/$USER_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/users/1356551933?app_version=v2
```

Respuesta:

```javascript
{
  "results": [
    {
      "id": "P-MLB1806015",
      "type": "MARKETPLACE_CAMPAIGN",
      "status": "started",
      "start_date": "2023-04-20T02:00:00Z",
      "finish_date": "2023-08-01T02:00:00Z",
      "deadline_date": "2023-08-01T01:00:00Z",
      "name": "Campanha de teste v2",
      "benefits": {
        "type": "REBATE",
        "meli_percent": 5,
        "seller_percent": 25
      }
    },
    {
      "id": "P-MLB1806017",
      "type": "VOLUME",
      "status": "started",
      "start_date": "2023-04-20T03:00:00Z",
      "finish_date": "2023-08-01T02:00:00Z",
      "deadline_date": "2023-08-01T01:00:00Z",
      "name": "Leva 3 paga 2",
      "benefits": {
        "type": "VOLUME",
        "meli_percent": 9.9999,
        "seller_percent": 23.3331,
        "name": "3x2",
        "buy_quantity": 3,
        "pay_quantity": 2,
        "item_discount_percent": 33.333
      }
    },
    {
      "id": "P-MLB1806019",
      "type": "DEAL",
      "status": "started",
      "start_date": "2023-04-20T03:00:00Z",
      "finish_date": "2023-08-01T02:00:00Z",
      "deadline_date": "2023-08-01T01:00:00Z",
      "name": "deals de teste v2"
    },
    {
      "id": "P-MLB1809008",
      "type": "DEAL",
      "status": "started",
      "start_date": "2023-04-21T21:00:00Z",
      "finish_date": "2023-08-01T02:00:00Z",
      "deadline_date": "2023-08-01T01:00:00Z",
      "name": "Deals de test v2"
    },
    {
      "id": "P-MLB1809009",
      "type": "DEAL",
      "status": "started",
      "start_date": "2023-04-21T23:00:00Z",
      "finish_date": "2023-08-01T02:00:00Z",
      "deadline_date": "2023-08-01T01:00:00Z",
      "name": "campanha de teste"
    }
  ],
  "paging": {
    "offset": 0,
    "limit": 50,
    "total": 5
  }
}
```

### Campos de la respuesta

**id**: código de identificación de la oferta.
 **type**: tipo de la oferta (DEAL, MARKETPLACE_CAMPAIGN, DOD, LIGHTNING, VOLUME, PRICE DISCOUNT, PRE_NEGOTIATED, SELLER_CAMPAIGN, SMART, PRICE_MATCHING, UNHEALTHY_STOCK y SELLER_COUPON_CAMPAIGN).
 **status**: Estado
 **start_date**: fecha de inicio de la oferta.
 **finish_date**: fecha de fin de la oferta.
 **deadline_date**: plazo máximo para aceptar la invitación.
 **name**: nombre de la promoción.
 **deadline_date**: plazo máximo para incorporar los ítems a la promoción.
 **benefits**: configuración de beneficios de la promoción.

## Consultar items candidatos

El recurso **/seller-promotions/candidates** permite identificar los ítems invitados a participar de una promoción. Siempre que un ítem obtiene el status de **candidate** en una promoción se envía una notificación con el **candidate_id**, con este recurso es posible identificar el ítem, la promoción y el status.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN'  https://api.mercadolibre.com/seller-promotions/candidates/$CANDIDATE_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN'  https://api.mercadolibre.com/seller-promotions/candidates/CANDIDATE- MLB1254949426-803130663?app_version=v2
```

Respuesta:

```javascript
{
  "id": "CANDIDATE-MLB1254949426-803130663",
  "item_id": "MLB1254949426",
  "promotion_id": "P-MLB4629001",
  "type": "MARKETPLACE_CAMPAIGN",
  "status": {
    "id": "candidate"
  }
}
```

**Campos de respuesta**

**id**: código de identificación del candidato.

**item_id**: ítem asociado al candidato.

**promotion_id**: id de la promoción.

**type**: tipo de promoción (DEAL, MARKETPLACE_CAMPAIGN, DOD, LIGHTNING, VOLUME, PRICE DISCOUNT, PRE_NEGOTIATED, SELLER_CAMPAIGN, SMART, PRICE_MATCHING, UNHEALTHY_STOCK y SELLER_COUPON_CAMPAIGN).

**status**: estado del candidato.

 Nota: 

 El id del candidato se obtiene a través de la notificación del topic 

public candidate

.

## Consultar ofertas

El recurso **/seller-promotions/offers** permite identificar cambios en la oferta de un ítem. Todos los cambios se envían por medio de notificaciones con el **offer_id**, es posible identificar el item, la promoción y el estado.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/offers/$OFFERS_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/offers/OFFER-MLB1970246686-42701792?app_version=v2
```

Respuesta:

```javascript
{
  "id": "OFFER-MLB1970246686-42701792",
  "item_id": "MLB1970246686",
  "promotion_id": "P-MLB3329001",
  "type": "DEAL",
  "status": {
    "id": "ACTIVE"
  }
}
```

### Campos de la respuesta

**id**: código de identificación de la oferta.
 **item_id**: ítem asociado a la oferta.
 **promotion_id**: id de la promoción.
 **type**: tipo de promoción (DEAL, MARKETPLACE_CAMPAIGN, DOD, LIGHTNING, VOLUME, PRICE DISCOUNT, PRE_NEGOTIATED, SELLER_CAMPAIGN, SMART, PRICE_MATCHING, UNHEALTHY_STOCK y SELLER_COUPON_CAMPAIGN).
 **status**: estado de la oferta. (programmed, active, e inactive).

 Nota: 

 El id de la oferta lo obtienes por medio de una notificación del tópico 

public offers

. 

## Consultar detalles de la promoción

Realiza la siguiente consulta para acceder a los detalles particulares de una campaña tradicional, campaña co-fondeada y para los descuentos por volumen.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID?promotion_type=$PROMOTION_TYPE&app_version=v2
```

Para obtener más información, acceda a las documentaciones de cada campaña.

## Estado

A continuación puedes encontrar los posibles estados que pueden tener los distintos tipos de promociones:

- [Estados de campaña tradicional](https://developers.mercadolibre.com.ar/es_ar/deals?#Estados)
- [Estado de una campaña co-fondeada](https://developers.mercadolibre.com.ar/es_ar/campanas-co-fondeadas?#Estados)
- [Estado de campaña descuento por cantidad](https://developers.mercadolibre.com.ar/es_ar/campanas-con-descuento-por-cantidad#Estados)
- [Estado de campaña pre-acordado por ítem y Campaña de liquidación stock Full](https://developers.mercadolibre.com.ar/es_ar/descuento-pre-acordado-por-item#Estados-de-las-campa%C3%B1as)
- [Estado campaña co-fondeada automatizada y campañas de precios competitivos](https://developers.mercadolibre.com.ar/es_ar/campanas-smart-price-matching#Estados)
- [Cupones del vendedor](https://developers.mercadolibre.com.ar/es_ar/cupones-del-vendedor)

## Consultar ítems de la promoción

 Nota: 

 En esta consulta se obtiene el estado del ítem en la campaña.

Para conocer los ítems que forman parte de una determinada oferta puedes realizar la siguiente consulta:

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTIONS_ID/items?promotion_type=PROMOTIONS_TYPE&app_version=v2
```

Además, puedes consultar ítems de una campaña:

- [Tradicional](https://developers.mercadolibre.com.ar/es_ar/deals#Consultar-%C3%ADtems-en-una-campa%C3%B1a-tradicional)
- [Co-fondeada](https://developers.mercadolibre.com.ar/es_ar/campanas-co-fondeadas#Consultar-%C3%ADtems-en-una-campa%C3%B1a-co-fondeada)
- [Descuento por cantidad](https://developers.mercadolibre.com.ar/es_ar/campanas-con-descuento-por-cantidad)
- [Descuento pre-acordado por ítem y Campaña de liquidación stock Full](https://developers.mercadolibre.com.ar/es_ar/descuento-pre-acordado-por-item)
- [Oferta del día](https://developers.mercadolibre.com.ar/es_ar/ofertas-del-dia#Consultar-%C3%ADtems-en-una-oferta-del-d%C3%ADa)
- [Oferta relámpago](https://developers.mercadolibre.com.ar/es_ar/ofertas-relampago#Consultar-%C3%ADtems-en-una-oferta-rel%C3%A1mpago)
- [Del vendedor](https://developers.mercadolibre.com.ar/es_ar/campanas-del-vendedor?nocache=true#)
- [Co-fondeada automatizada y campañas de precios competitivos](https://developers.mercadolibre.com.ar/es_ar/campanas-smart-price-matching)
- [Cupones del vendedor](https://developers.mercadolibre.com.ar/es_ar/cupones-del-vendedor)

## Filtros

Puedes aplicar filtros por item_id, status y status_item:

- **item_id:** Permite filtrar por un ítem específico.
- **status:** Permite filtrar por el estado de la oferta: **started**, **pending** o **candidate**.
- **status_item:** Permite filtrar por el estado de los ítems que forman parte de la campaña, pudiendo ser **active** o **paused**.

 Nota: 

Cuando se envía el filtro status_item, solo se devuelven los ítems correspondientes al estado consultado: "active" o "paused". Si no se incluye este parámetro, la consulta, por defecto, devuelve únicamente los ítems activos en Mercado Libre.

En caso de enviar un valor distinto de "active" o "paused", se responderá con un 

400 - Bad Request.

 . 

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID/items?promotion_type=$PROMOTION_TYPE&status=$STATUS&item_id=$ITEM_ID&app_version=v2
```

Ejemplo de filtro por ítem_id:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/MLA1111/items?promotion_type=DEAL&item_id=MLA604400000&app_version=v2
```

Respuesta:

```javascript
{
  "results": [
    {
      "id": "MLA604400000",
      "status": "started",
      "price": 23968,
      "original_price": 28549
    }
  ],
  "paging": {}
}
```

Ejemplo de filtro por status:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' /seller-promotions/promotions/MLA1111/items?promotion_type=DEAL&status=started&app_version=v2
```

Respuesta:

```javascript
{
  "results": [
    {
      "id": "MLA639970000",
      "status": "started",
      "price": 4037,
      "original_price": 4427
    },
    {
      "id": "MLA639973333",
      "status": "started",
      "price": 6007,
      "original_price": 6587
    }
  ],
  "paging": []
}
```

Ejemplo de filtro por status_item:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' /seller-promotions/promotions/MLA1111/items?promotion_type=DEAL&status_item=active&app_version=v2
```

Respuesta:

```javascript
{
  "results": [
    {
      "id": "MLA639970000",
      "status": "started",
      "price": 4037,
      "original_price": 4427
    },
    {
      "id": "MLA639973333",
      "status": "started",
      "price": 6007,
      "original_price": 6587
    }
  ],
  "paging": []
}
```

## Paginación

 Importante: 

- El query param para enviar este valor 

comienza a llamarse search_after

 y no más searchAfter. Pero se continuará aceptando searchAfter por un tiempo. 

- 

Se unifica el valor de search_after

 para que solo utilice valores distintos, eliminando la ambigüedad.

Para realizar la paginación deberás utilizar el parámetro search_after. 
En la respuesta del GET, devolvemos el parámetro searchAfter, el cual servirá para poder recorrer los resultados. Para ello se deberá recuperar dicho ID y realizar la siguiente request con el query param search_after={search_after}. Este ID es un string, por eso tienen que aceptar el string y usarlo luego en sus pegadas.

 Nota: 

 Si no utilizas el parámetro de limit, se retornarán por defecto 50 ítems del total. Puedes agregar un limit máximo de 50.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTIONS_ID/items?promotion_type=$PROMOTION_TYPE&app_version=v2&limit=50&search_after={$SEARCH_AFTER}
```

Ejemplo de paginación:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB13451002MLB9377/items?promotion_type=DEAL&app_version=v2&limit=50&search_after=d3e3fb02371ca8e49ceb3e998c4a1b8da189235497375e55d3027c7dacf5a4ef0175b2aaca4f4a0fdf31299947d82ba661037482172ba7f9cfb1b0250d3134aa71c889367aa7c1401e4c3ff5bd70ee14337dfaa18c99bbe5e52dc3a2c1b55b195131903ecbc60a1c639e01dbecf11b15126d4b38cdb6122182acde2eca1b1a42
```

Respuesta:

```javascript
{
  "results": [
    {
      "id": "MLB2674512266",
      "status": "candidate",
      "price": 0,
      "original_price": 0
    },
    {
      "id": "MLB2674506199",
      "status": "candidate",
      "price": 0,
      "original_price": 0
    },
    {
      "id": "MLB2674506138",
      "status": "candidate",
      "price": 0,
      "original_price": 0
    },
    {
      "id": "MLB2674505931",
      "status": "candidate",
      "price": 0,
      "original_price": 0
    },
    {
      "id": "MLB2674505924",
      "status": "candidate",
      "price": 0,
      "original_price": 0
    }
  ],
  "paging": {
    "searchAfter": "d3e3fb02371ca8e49ceb3e998c4a1b8da189235497375e55d3027c7dacf5a4ef0175b2aaca4f4a0fdf31299947d82ba661037482172ba7f9cfb1b0250d3134aa71c889367aa7c1401e4c3ff5bd70ee14337dfaa18c99bbe5e52dc3a2c1b55b195131903ecbc60a1c639e01dbecf11b15126d4b38cdb6122182acde2eca3b5b55",
    "limit": 50,
    "total": 14424
  }
}
```

### Consideraciones

- Se devolverá search_after en todas las páginas, excepto en la última.
- La única forma de avanzar en la respuesta (paginar) es a través del uso de este parámetro.
- Al iterar los resultados, cada llamada retornará el search_after que deberá ser utilizado en la siguiente llamada.
- Siempre se debe utilizar el search_after que fue proporcionado por la respuesta del request, ya que este puede cambiar y expirar (tienen un TTL de 5 minutos).
- No es posible realizar paginados hacia atrás.

## Cómo participar

Puedes participar en distintos tipos de promociones e incluso ofrecer un descuento individual para los ítems:

- [Indicando ítems para una campaña tradicional](https://developers.mercadolibre.com.ve/es_ar/deals?#Indicar-ítems-para-una-campaña-tradicional).
- [Indicando ítems para una campaña co-fondeada](https://developers.mercadolibre.com.ve/es_ar/campanas-co-fondeadas?#Indicar-ítems-para-una-campaña-co-fondeada).
- [Indicando ítems para descuento por volumen](https://developers.mercadolibre.com.ve/es_ar/campanas-con-descuento-por-volumen#Indicar-ítems-para-una-campaña-con-descuento-por-volumen).
- [Aceptando descuento pre-acordado por ítem](https://developers.mercadolibre.com.ve/es_ar/descuento-pre-acordado-por-item#Aceptar-descuentopre-acordado-por-%C3%ADtem).
- [Indicando ítems para una oferta del día](https://developers.mercadolibre.com.ve/es_ar/ofertas-del-dia#Indicar-ítems-para-una-oferta-del-día).
- [Indicando ítems para una oferta oferta ralámpago](https://developers.mercadolibre.com.ve/es_ar/es_ar/ofertas-relampago#Indicar-ítems-para-una-oferta-relámpago).
- [Ofreciendo un descuento individual para un ítem](https://developers.mercadolibre.com.ve/es_ar/descuento-individual #Ofrecer-un-descuento-para-un-ítem).
- [Indicando items para una campaña del vendedor.](https://developers.mercadolibre.com.ve/es_ar/campanas-del-vendedor?nocache=true#)
- [Indicando items para una campaña smart.](https://developers.mercadolibre.com.ve/es_ar/campanas-smart?nocache=true)

## Consultar promociones del ítem

Este recurso devuelve todas las promociones asociadas a un ítem. La respuesta indica el estado de participación del ítem en cada promoción y el precio correspondiente en el momento de la consulta. No incluye información general de la promoción.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?app_version=v2
```

Respuesta:

```javascript
{
  "id": "C-MLB13967",
  "type": "SELLER_COUPON_CAMPAIGN",
  "sub_type": "FIXED_PERCENTAGE",
  "fixed_percentage": 15,
  "status": "candidate",
  "price": 0,
  "original_price": 20,
  "start_date": "2025-09-03T00:00:00",
  "finish_date": "2025-09-09T23:59:59",
  "name": "Test"
}
{
  "id": "C-MLB139959",
  "type": "SELLER_COUPON_CAMPAIGN",
  "sub_type": "FIXED_AMOUNT",
  "fixed_amount": 10,
  "status": "started",
  "price": 0,
  "original_price": 20,
  "start_date": "2025-09-03T00:00:00",
  "finish_date": "2025-09-09T23:59:59",
  "name": "Esportivo"
}
{
  "id": "P-MLB4912006",
  "type": "PRICE_MATCHING_MELI_ALL",
  "ref_id": "OFFER-MLB4141870813-10000236971",
  "status": "started",
  "price": 18.2,
  "meli_percentage": 9,
  "seller_percentage": 0,
  "original_price": 20,
  "name": "test_Livre"
}
{
  "id": "P-MLB5108026",
  "type": "SMART",
  "ref_id": "CANDIDATE-MLB5457079540-70000147624",
  "status": "candidate",
  "price": 82.67,
  "meli_percentage": 4.6,
  "seller_percentage": 12.7,
  "original_price": 100,
  "name": "Promo test "
}
{
  "type": "PRICE_DISCOUNT",
  "status": "candidate",
  "price": 0,
  "original_price": 20,
  "name": "",
  "min_discounted_price": 6,
  "max_discounted_price": 19,
  "suggested_discounted_price": 18
}
{
  "type": "PRICE_DISCOUNT",
  "status": "started",
  "price": 14,
  "original_price": 20,
  "start_date": "2025-09-09T00:00:00",
  "finish_date": "2025-09-10T23:59:59",
  "name": ""
}
{
  "id": "C-MLB13962",
  "type": "SELLER_CAMPAIGN",
  "sub_type": "FLEXIBLE_PERCENTAGE",
  "status": "candidate",
  "price": 0,
  "original_price": 20,
  "start_date": "2025-09-02T00:00:00",
  "finish_date": "2025-09-08T23:59:59",
  "name": "desc por porc",
  "min_discounted_price": 6,
  "max_discounted_price": 19,
  "suggested_discounted_price": 18
}
{
  "id": "C-MLB14160",
  "type": "SELLER_CAMPAIGN",
  "sub_type": "FLEXIBLE_PERCENTAGE",
  "status": "started",
  "price": 14,
  "original_price": 20,
  "start_date": "2025-09-09T00:00:00",
  "finish_date": "2025-09-15T23:59:59",
  "name": "seller campaign"
},
{
  "id": "LGH-MLB1000",
  "type": "LIGHTNING",
  "ref_id": "CANDIDATE-MLB4141935865-70000146658",
  "status": "candidate",
  "price": 0,
  "original_price": 18.99,
  "min_discounted_price": 5.7,
  "max_discounted_price": 18.04,
  "suggested_discounted_price": 17.1,
  "stock": {
    "min": 1,
    "max": 2
  }
},
{
  "id": "LGH-MLB1000",
  "type": "LIGHTNING",
  "ref_id": "OFFER-MLB4141870813-10000237588",
  "status": "pending",
  "price": 12.9,
  "original_price": 20,
  "stock": {
    "remaining_stock": 1
  }
},
{
  "id": "DOD-MLB1000",
  "type": "DOD",
  "ref_id": "CANDIDATE-MLB4141935865-70000146656",
  "status": "candidate",
  "price": 10.99,
  "original_price": 18.99,
  "min_discounted_price": 5.7,
  "max_discounted_price": 18.04,
  "suggested_discounted_price": 17.1,
  "stock": {
    "min": 1,
    "max": 2
  }
},
{
  "id": "DOD-MLB1000",
  "type": "DOD",
  "ref_id": "OFFER-MLB4141870813-10000237589",
  "status": "pending",
  "price": 12.9,
  "original_price": 20,
  "stock": {}
}
{
  "id": "P-MLB5106026",
  "type": "DEAL",
  "status": "started",
  "price": 13.26,
  "original_price": 18.99,
  "start_date": "2025-07-30T16:00:00-03:00",
  "finish_date": "2025-10-04T23:00:00-03:00",
  "name": "Test_nuevo "
},
{
  "id": "P-MLB5106026",
  "type": "DEAL",
  "status": "candidate",
  "price": 0,
  "original_price": 18.99,
  "start_date": "2025-07-30T16:00:00-03:00",
  "finish_date": "2025-10-04T23:00:00-03:00",
  "name": "Test Tier 1 - nuevo ",
  "min_discounted_price": 5.7,
  "max_discounted_price": 18.04,
  "suggested_discounted_price": 17.1
},
{
  "id": "P-MLB4914006",
  "type": "MARKETPLACE_CAMPAIGN",
  "status": "candidate",
  "price": 80,
  "meli_percentage": 4,
  "seller_percentage": 16,
  "original_price": 100,
  "start_date": "2025-07-10T18:00:00Z",
  "finish_date": "2025-09-27T02:00:00Z",
  "name": "Test_refresh "
},
{
  "id": "P-MLB4914006",
  "type": "MARKETPLACE_CAMPAIGN",
  "status": "started",
  "price": 80,
  "meli_percentage": 4,
  "seller_percentage": 16,
  "original_price": 100,
  "start_date": "2025-07-10T18:00:00Z",
  "finish_date": "2025-09-27T02:00:00Z",
  "name": "Test_refresh "
},
{
  "id": "P-MLU4942004",
  "type": "PRE_NEGOTIATED",
  "ref_id": "CANDIDATE-MLU760878674-70000115453",
  "status": "candidate",
  "price": 4338,
  "meli_percentage": 4.2,
  "seller_percentage": 4.2,
  "original_price": 4738,
  "name": "prene_test"
},
{
  "id": "P-MLB5128002",
  "type": "BANK",
  "ref_id": "OFFER-MLB4141935865-10000237044",
  "sub_type": "COFINANCED",
  "status": "started",
  "meli_percentage": 6,
  "seller_percentage": 8,
  "original_price": 18.99,
  "start_date": "2025-09-04T11:54:50Z",
  "finish_date": "2025-10-05T02:00:00Z",
  "name": "Desconto no Pix",
  "payment_method": "PIX"
},
{
  "id": "P-MLB5128002",
  "type": "BANK",
  "ref_id": "CANDIDATE-MLB4141921895-70000146801",
  "sub_type": "COFINANCED",
  "status": "candidate",
  "meli_percentage": 6,
  "seller_percentage": 8,
  "original_price": 239.78,
  "start_date": "2025-08-01T14:00:00Z",
  "finish_date": "2025-10-05T02:00:00Z",
  "name": "Desconto no Pix",
  "payment_method": "PIX"
}
```

### Campos de respuesta:

**id**: Identificador de la promoción

**status**: Estado específico del ítem en la promoción:

- **candidate**: El ítem es elegible y puede participar en la promoción
- **started**: El ítem participa activamente en la promoción
- **pending**: El ítem fue optineado pero la oferta aún no comenzó

**original_price**: precio del ítem sin descuento.

**min_discounted_price**: Precio mínimo permitido en la promoción. Refleja el mayor descuento posible para el ítem.

**max_discounted_price**: Precio máximo al que puede ofrecerse el ítem en la promoción, garantizando descuentos creíbles.

**suggested_discounted_price**: Precio sugerido para una oferta atractiva, basado en el historial y contexto del ítem. Puede ser null si no hay una sugerencia disponible.

#### **Según promoción**

**Deal**

**top_deal_price**: Precio exclusivo disponible únicamente para compradores destacados (niveles 3 y 6 de Mercado Puntos). Este campo solo aparece si el ítem está activo en la campaña y el vendedor lo configuró al momento de sumarse.

**Marketplace campaign**

**ref_id**: id de la oferta o candidato (presente solo cuando el estado es started).

**meli_percentage**: Porcentaje de descuento aportado por Mercado Libre.

**seller_percentaje**: Porcentaje de descuento aportado por el vendedor.

**price**: precio del ítem en la campaña

**Seller campaign**

**sub_type**: FLEXIBLE_PERCENTAGE.

**price**: precio del ítem en la campaña

**Volume**

**buy_quantity/pay_quantity_discount_percentage**: se completa de acuerdo al subtipo de promo.

**allow_combination**: permite la combinación de items.

**sub_type**: pudiendo ser BNGM - BNSP - SPONTH.

**Oferta del día y Oferta relámpago**

**stock**: Información sobre el stock mínimo y máximo requerido para que el ítem pueda sumarse como candidato a la promoción.

**Cupones**

**fixed_percentage**: Porcentaje de descuento ofertado (solo para subtipo FIXED_PERCENTAGE).

**sub_type**: Subtipo de la campaña. Indica si el cupón es de monto fijo (FIXED_AMOUNT) o porcentaje (FIXED_PERCENTAGE).

**fixed_amount**: Monto fijo de descuento otorgado (solo para subtipo FIXED_AMOUNT).

## Modificar ítems

Puedes modificar los ítems que están participando en una determinada oferta:

- [Modificando ítems en una campaña tradicional](https://developers.mercadolibre.com.ve/es_ar/deals?#Modificar-ítems).
- [Modificando ítems en una campaña co-fondeada](https://developers.mercadolibre.com.ve/es_ar/es_ar/campanas-co-fondeadas?#Modificar-ítems).
- [Modificando ítems en una campaña con descuento por volumen](https://developers.mercadolibre.com.ve/es_ar/campanas-con-descuento-por-volumen#Modificar-ítems).

 Nota: 

Para editar los descuentos individuales (PRICE DISCOUNT), las ofertas del día (DOD) y las ofertas relámpago (LIGHTNING) debes eliminar la promoción y aplicarla nuevamente.

## Delete masivo de ofertas

Puedes eliminar de forma masiva todas las ofertas que están en el ítem.

 Nota: 

Este delete masivo no se aplica en casos de ofertas de campañas del tipo DOD y LIGHTNING. Para estas ofertas, es necesario continuar eliminando una campaña por vez.

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/MLA1399846831?app_version=v2
```

Respuesta:

```javascript
{
  "successful_ids": [
    {
      "offer_id": "OFFER-MLA1399846831-10000081416",
      "error": null
    },
    {
      "offer_id": "OFFER-MLA1399846831-10000081567",
      "error": null
    }
  ],
  "errors": []
}
```

En casos donde el ítem tenga campañas que no pueden ser eliminadas de forma masiva, la llamada tendrá una respuesta http 200, pero la respuesta contendrá los mensajes de error.

Ejemplo donde ninguna oferta puede ser eliminada:

```javascript
{
  "successful_ids": [],
  "errors": [
    {
      "offer_id": "OFFER-MLA1399846831-10000081416",
      "error": "The offer of type DOD not allowed for delete."
    },
    {
      "offer_id": "OFFER-MLA1399846831-10000081828",
      "error": "The offer could not be deleted. Try again."
    }
  ]
}
```

Ejemplo donde las ofertas fueron eliminadas correctamente y también ocurrieron errores:

```javascript
{
  "successful_ids": [
    {
      "offer_id": "OFFER-MLA1399846831-10000081416",
      "error": null
    },
    {
      "offer_id": "OFFER-MLA1399846831-10000081417",
      "error": null
    }
  ],
  "errors": [
    {
      "offer_id": "OFFER-MLA1399846831-10000081418",
      "error": "The offer of type DOD not allowed for delete."
    },
    {
      "offer_id": "OFFER-MLA1399846831-10000081419",
      "error": "The offer could not be deleted. Try again."
    }
  ]
}
```

### Posibles errores

**423_ENTITY_LOCKED**: La solicitud no pudo ser procesada porque el ítem está temporalmente bloqueado para realizar solicitudes. La solicitud puede intentarse nuevamente después de unos segundos.

**400_BAD_REQUEST**: Cuando el formato del ítem es inválido.

## Eliminar ítems

Puedes eliminar los ítems que están participando en una determinada oferta:

- [Eliminando ítems en una campaña tradicional](https://developers.mercadolibre.com.ve/es_ar/deals?#Eliminar-%C3%ADtems).
- [Eliminando ítems en una campaña co-fondeada](https://developers.mercadolibre.com.ve/es_ar/campanas-co-fondeadas?#Eliminar-%C3%ADtems).
- [Eliminando ítems en una campaña con descuento por volumen](https://developers.mercadolibre.com.ve/es_ar/campanas-con-descuento-por-volumen#Eliminar-%C3%ADtems).
- [Eliminando descuento pre-acordado por ítem](https://developers.mercadolibre.com.ve/es_ar/descuento-pre-acordado-por-item#Eliminar-%C3%ADtems).
- [Eliminando ítems en una oferta del día](https://developers.mercadolibre.com.ve/es_ar/ofertas-del-dia#Eliminar%20%C3%ADtems).
- [Eliminando ítems en una oferta relámpago](https://developers.mercadolibre.com.ve/es_ar/ofertas-relampago#Eliminar%20%C3%ADtems).
- [Eliminando descuento individual a un ítem](https://developers.mercadolibre.com.ve/es_ar/descuento-individual?#Eliminar-descuento-individual-a-un-%C3%ADtem).

## Gestión de lista de exclusión para Campañas Automáticas

Con este recurso podrás administrar la lista de exclusión automática para las promociones en Mercado Libre. Si deseas evitar que determinados **sellers** o **productos** participen en campañas de forma automática, esta guía te mostrará cómo hacerlo.

### Consulta por Seller

Puedes verificar si un seller está excluido de la participación automática en promociones.

**Llamada:**

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/seller-promotions/exclusion-list/seller?app_version=v2
```

**Respuesta:**

```javascript
{
  "excluded": "not_excluded"
}
```

**Parámetros:**

- **excluded**: Indica si el seller está excluido.
  - **"not_excluded"**: No está excluido.
  - **"excluded"**: Está excluido.

### Gestionar sellers de la Lista de Exclusión

Puedes agregar o eliminar un seller de la lista de exclusión para controlar su participación en promociones automáticas.

**Importante:** Mercado Libre no creará ofertas de participación automática para sellers excluidos.

**Llamada:**

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
    https://api.mercadolibre.com/seller-promotions/exclusion-list/seller?app_version=v2
    --data '{
    "exclusion_status": "true" 
    }'
```

**Parámetros:**

- **exclusion_status**: define la acción a realizar.
  - **"true"**: Excluir al seller.
  - **"false"**: Eliminar de la exclusión.

### Consulta por items

Puedes verificar si un item está excluido de la participación automática en promociones.

**Llamada:**

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
    https://api.mercadolibre.com/seller-promotions/exclusion-list/seller/{item_id}?app_version=v2
```

**Respuesta:**

```javascript
{
  "excluded": "excluded"
}
```

**Parámetros:**

- **excluded**: Indica si el seller está excluido.
  - **"excluded"**: Está excluido.
  - **"not_excluded"**: No está excluido.

### Gestionar items de la Lista de Exclusión

Puedes agregar o eliminar un item de la lista de exclusión para evitar o permitir su participación en promociones automáticas.

**Importante:**Mercado Libre no creará ofertas de participación automática para items excluidos.

**Llamada:**

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
   https://api.mercadolibre.com/seller-promotions/exclusion-list/item?app_version=v2
--data '{
    "item_id": "12345678",
    "exclusion_status": "false"
}'
```

**Parámetros:**

- **item_id**: ID del producto a modificar.
- **exclusion**: Define la acción a realizar.
  - **"true"**: Excluir el item.
  - **"false"**: Eliminar de la exclusión.

## Asignar campañas de pruebas

Para realizar pruebas con campañas de test, envíanos los datos de tu usuario e ítems en el siguiente **Formulario:**.

- [Formulario](https://docs.google.com/forms/d/e/1FAIpQLSenA_USmZQb8deHLrjhO_Rx1oOqfsj--Rhv-f_L1SebEJRBjA/viewform)

Recuerda que tanto los usuarios como los ítems deben ser de test.

 Nota: 

Debes agregar el parámetro 

version=test

 dentro de las llamadas para interactuar con las promociones de test.

**Next post**: [Campañas co-fondeadas](https://developers.mercadolibre.com.ve/es_ar/campanas-co-fondeadas)
