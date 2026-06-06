---
title: Co-fondeada automatizada y precios competitivos
slug: campanas-smart-price-matching
section: 14-promotions
source: https://developers.mercadolibre.com.ve/es_ar/campanas-smart-price-matching
captured: 2026-06-06
---

# Co-fondeada automatizada y precios competitivos

> Source: <https://developers.mercadolibre.com.ve/es_ar/campanas-smart-price-matching>

## Endpoints referenced

- `https://api.mercadolibre.com/seller-promotions//P-MLB3528002/items?promotion_type=PRICE_MATCHING_MELI_ALL&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/P-MLB2087012/items?promotion_type=PRICE_MATCHING&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&promotiondeal_id=$PROMOTION&offer_id=$OFFER_ID`
- `https://api.mercadolibre.com/seller-promotions/items/MLA1387793467?promotion_type=PRICE_MATCHING_MELI_ALL&promotion_id=P-MLA2072013&offer_id=OFFER-MLA1387793467-1000000151&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?promotion_type=SMART&promotion_id=P-MLB1812010&offer_id=OFFER-MLB3538191898-177685&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB4048719074?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB4048719074?promotion_type=PRICE_MATCHING&promotion_id=P-MLB2087012&offer_id=OFFER-MLB4048719074-10000001972&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1812010/items?promotion_type=SMART&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1812010?promotion_type=SMART&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB2087012?promotion_type=PRICE_MATCHING&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB35280024?promotion_type=PRICE_MATCHING_MELI_ALL&app_version=v2`

## Content

Última actualización 02/06/2026

## Campañas co-fondeada automatizada y campañas de precios competitivos

 Importante: 

 El 

nuevo filtro por estado

 ya está disponible para filtrar los ítems de una campaña mediante el query param 

status_item

, que acepta los valores "active" o "paused". 

 Importante: 

-Estas campañas se encuentran en la misma documentación porque trabajan con la misma lógica y los mismos parámetros. 

 - A partir del día 14 de octubre de 2024, las campañas de precios competitivos cuentan con dos types: PRICE_MATCHING y PRICE_MATCHING_MELI_ALL (100% a cargo de Mercado Libre). 

 - En el caso en que el vendedor desee excluir su participación en las campañas de precios competitivos 100% Mercado Libre (PRICE_MATCHING_MELI_ALL), deberá acceder a de Mercado Libre para realizar desactivar dicha funcionalidad. Para remover ítems específicos de 

Los vendedores son invitados de manera periódica a participar en diversas campañas que se realizan en el sitio. En el caso de las **campañas co-fondeadas automatizadas y las de precios competitivos**, Mercado Libre asume un porcentaje del descuento ofrecido. 
 Las campañas co-fondeadas automatizadas funcionan de manera similar a las co-fondeadas tradicionales, pero utilizan un proceso automatizado para seleccionar los ítems que serán invitados a participar. En cuanto a las campañas de precios competitivos, su objetivo es asegurar que los productos alcancen el mejor precio en comparación con otros sitios web y marketplaces. Los candidatos para estas campañas se actualizan diariamente, lo que significa que un ítem puede ser elegible hoy, pero no necesariamente mañana.
 
 A partir de ahora, las campañas de precios competitivos ofrecen dos tipos de promociones:

- **PRICE_MATCHING**: El descuento es cofinanciado entre el vendedor y Mercado Libre.
- **PRICE_MATCHING_MELI_ALL**: El descuento es 100% financiado por Mercado Libre, y la participación del vendedor se gestiona automáticamente, sin necesidad de ninguna acción por su parte.

Esta estructura proporciona mayor flexibilidad en la implementación de descuentos, adaptándose a las características de cada campaña. Si el vendedor recibió una invitación y quiere sumarse, puede hacerlo con los siguientes recursos.

## Consultar detalle de campaña

Nota:

Para estas campañas, las respuestas tienen los mismos campos, cambiando únicamente la información del "type" (SMART, o PRICE_MATCHING o PRICE_MATCHING_MELI_ALL). 

Para obtener los detalles de una promoción, realiza la siguiente consulta:

Ejemplo de co-fondeada automatizada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1812010?promotion_type=SMART&app_version=v2
```

Respuesta de co-fondeada automatizada:

```javascript
{
  "id": "P-MLB1812010",
  "type": "SMART",
  "status": "started",
  "start_date": "2023-04-26T23:00:00Z",
  "finish_date": "2023-05-10T23:59:00Z",
  "deadline_date": "2023-05-10T23:59:00Z",
  "name": "test-smart-2"
}
```

Ejemplo de precios competitivos co-fondeada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB2087012?promotion_type=PRICE_MATCHING&app_version=v2
```

Respuesta de precios competitivos co-fondeada:

```javascript
{
  "id": "P-MLB2087012",
  "type": "PRICE_MATCHING",
  "status": "pending",
  "start_date": "2023-09-19T18:15:00Z",
  "finish_date": "2023-10-01T05:59:59Z",
  "deadline_date": "2023-10-01T05:59:59Z",
  "name": "Gánale a la competencia con un aporte de Mercado Libre"
}
```

Ejemplo de precios competitivos 100% Mercado Libre:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB35280024?promotion_type=PRICE_MATCHING_MELI_ALL&app_version=v2
```

Respuesta de precios competitivos 100% Mercado Libre:

```javascript
{
    "id": "P-MLB3528002",
    "type": "PRICE_MATCHING_MELI_ALL",
    "status": "started",
    "start_date": "2024-09-26T15:20:04Z",
    "finish_date": "2024-10-01T15:18:04Z",
    "deadline_date": "2024-10-01T15:18:04Z",
    "name": "100% a cargo de Mercado Libre"
}
```

### Campos de la respuesta

- **id**: identificador de la campaña.
- **type**: tipo de campaña (SMART, PRICE_MATCHING o PRICE_MATCHING_MELI_ALL).
- **status**: status de la campaña.
- **start_date**: fecha que empieza la campaña.
- **finish_date**: fecha que se cierra la campaña.
- **deadline_date**: fecha límite para crear la campaña.
- **name**: nombre de la campaña.

## Estados

Estos son los distintos estados por los que puede pasar en las campañas co-fondeada automatizada y precios competitivos.

| Estado | Descripción |
| --- | --- |
| **pending** | Promoción aprobada, pero aún no inició. |
| **started** | Promoción activa. |
| **finished** | Promoción finalizada. |

## Consultar ítems en una campaña

Para conocer los ítems que forman parte de una campaña puedes realizar la siguiente consulta:

Ejemplo de co-fondeada automatizada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1812010/items?promotion_type=SMART&app_version=v2
```

Respuesta de co-fondeada automatizada:

```javascript
{
  "results": [
      {
          "id": "MLB3538191898",
          "status": "candidate",
          "price": 3000,
          "original_price": 5000,
          "offer_id": "CANDIDATE-MLB3538191898-25593903",
          "meli_percentage": 20,
          "seller_percentage": 20,
          "start_date": "2023-04-26T11:40:00Z",
          "end_date": "2023-05-30T15:47:00Z"
      }
  ],
  "paging": {
      "offset": 0,
      "limit": 50,
      "total": 1
  }
}
```

Ejemplo de precios competitivos co-fondeada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/P-MLB2087012/items?promotion_type=PRICE_MATCHING&app_version=v2
```

Respuesta de precios competitivos co-fondeada:

```javascript
{
   "results": [
       {
           "id": "MLB4048719074",
           "status": "candidate",
           "price": 3000,
           "original_price": 5000,
           "offer_id": "CANDIDATE-MLB4048719074-70000001705",
           "meli_percentage": 20,
           "seller_percentage": 20,
           "start_date": "2023-09-19T03:00:00Z",
           "end_date": "2023-09-26T02:59:59Z"
       }
   ],
   "paging": {
       "total": 1,
       "limit": 50
   }
}
```

Ejemplo de precios competitivos 100% Mercado Libre:

Nota:

Para este tipo de campaña, no se visualizará un estado de "candidate", dado que el proceso de activar es ejecutado automáticamente por Mercado Libre. Cuando la campaña es mostrada al vendedor, indica que los ítems ya han sido activados en la campaña de manera automática. 

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 'https://api.mercadolibre.com/seller-promotions//P-MLB3528002/items?promotion_type=PRICE_MATCHING_MELI_ALL&app_version=v2''
```

Respuesta de precios competitivos 100% Mercado Libre:

```javascript
{
    "results": [
        {
            "id": "MLB3845318745",
            "status": "started",
            "price": 121.5,
            "original_price": 135,
            "offer_id": "OFFER-MLB3845318745-10000115845",
            "meli_percentage": 10,
            "seller_percentage": 0,
            "start_date": "2024-09-26T15:24:35Z",
            "end_date": "2024-09-28T23:59:59Z"
        }
    ],
    "paging": {
        "total": 1,
        "limit": 50
    }
}
```

Al crearse una nueva campaña del tipo SMART y PRICE_MATCHING se seleccionan todos los ítems aplicables a la misma. El estado inicial (**status**) de los ítems es **candidate** y sin offer id asignado. Al momento que el vendedor incorpora un ítem a la campaña su status se modifica y se le asigna un **offer_id** único.

## Estado de los ítems

En la siguiente tabla puedes encontrar los posibles estados que pueden tomar los ítems dentro de estos tipos de campañas.

| Estado | Descripción |
| --- | --- |
| **candidate** | Ítem candidato para participar de la promoción. |
| **pending** | Ítem con promoción aprobada y programada. |
| **started** | Ítem activo en la campaña. |
| **finished** | Ítem eliminado de la campaña |

## Indicar ítems para una campaña

Nota:

La campaña co-fondeada automatizada puede tener una duración máxima de 30 días, mientras que la campaña de precios competitivos puede tener hasta 10 días. 

Una vez que has sido invitado a participar en una de estas campañas, puedes indicar qué productos deseas incluir en las mismas.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' \
  -d '{
     "promotion_id":"$PROMOTIONDEAL_ID",
     "promotion_type":"$PROMOTION_TYPE",
     "offer_id":"$OFFER_ID"
  }'
  https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2
```

Ejemplo de co-fondeada automatizada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' \
  -d '{
    "promotion_id":"P-MLB1812010",
    "promotion_type":"SMART",
    "offer_id":"CANDIDATE-MLB3538191898-25593903"
  }
  '
  https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?app_version=v2
```

Respuesta de co-fondeada automatizada:

```javascript
{
  "offer_id": "OFFER-MLB3538191898-177685",
  "price": 3000,
  "original_price": 5000
}
```

Ejemplo de precios competitivos:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' \
  -d '{
     "promotion_id": "P-MLB2087012",
     "offer_id": "CANDIDATE-MLB4048719074-70000001705",
     "promotion_type": "PRICE_MATCHING"
  }
  '
  https://api.mercadolibre.com/seller-promotions/items/MLB4048719074?app_version=v2
```

Respuesta de precios competitivos:

```javascript
{
    "offer_id": "OFFER-MLB4048719074-10000001972",
    "price": 3000,
    "original_price": 5000
 }
 
```

### Parámetros

- **promotion_id**: identificación de la promoción.
- **promotion_type**: tipo de promoción (SMART o PRICE_MATCHING).
- **offer_id**: identificación de la oferta acordada.

## Eliminar campaña

Llamada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&promotiondeal_id=$PROMOTION&offer_id=$OFFER_ID
```

Ejemplo de co-fondeada automatizada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' 'https://api.mercadolibre.com/seller-promotions/items/MLB3538191898?promotion_type=SMART&promotion_id=P-MLB1812010&offer_id=OFFER-MLB3538191898-177685&app_version=v2
```

Respuesta: **Status 200 OK**

Ejemplo de precios competitivos co-fondeada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' 'https://api.mercadolibre.com/seller-promotions/items/MLB4048719074?promotion_type=PRICE_MATCHING&promotion_id=P-MLB2087012&offer_id=OFFER-MLB4048719074-10000001972&app_version=v2
```

Respuesta: **Status 200 OK**

Ejemplo de precios competitivos 100% Mercado Libre: :

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' 'https://api.mercadolibre.com/seller-promotions/items/MLA1387793467?promotion_type=PRICE_MATCHING_MELI_ALL&promotion_id=P-MLA2072013&offer_id=OFFER-MLA1387793467-1000000151&app_version=v2
```

Respuesta: **Status 200 OK**
