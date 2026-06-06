---
title: Ofertas relámpago
slug: ofertas-relampago
section: 14-promotions
source: https://developers.mercadolibre.com.ve/es_ar/ofertas-relampago
captured: 2026-06-06
---

# Ofertas relámpago

> Source: <https://developers.mercadolibre.com.ve/es_ar/ofertas-relampago>

## Endpoints referenced

- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2&promotion_type=$PROMOTION_TYPE`
- `https://api.mercadolibre.com/seller-promotions/items/MLA632979587?app_version=v2&promotion_type=LIGHTNING`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3293401743?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID/items?app_version=v2&promotion_type=LIGHTNING`
- `https://api.mercadolibre.com/seller-promotions/promotions/LGH-MLB1000/items?app_version=v2&promotion_type=LIGHTNING`

## Content

Última actualización 13/03/2026

# Ofertas relámpago

 Importante: 

 El 

nuevo filtro por estado

 ya está disponible para filtrar los ítems de una campaña mediante el query param 

status_item

, que acepta los valores "active" o "paused". 

Los vendedores son invitados periódicamente a participar de diferentes promociones que se realizan en el sitio. Si recibiste la invitación para participar de una **oferta relámpago** y quieres sumarte puedes hacerlo con los siguientes recursos. 
 Ten en cuenta que este tipo de ofertas tiene un stock reservado que al terminarse se finaliza la promoción.

## Consultar ítems

Para conocer los ítems que forman parte de una **oferta relámpago** puedes realizar la siguiente consulta:.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/$PROMOTION_ID/items?app_version=v2&promotion_type=LIGHTNING
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/LGH-MLB1000/items?app_version=v2&promotion_type=LIGHTNING
```

Respuesta:

```javascript
{
   "results": [
       {
           "id": "MLB3293401743",
           "start_date": "2023-04-21T15:00:00",
           "finish_date": "2023-04-21T23:00:00",
           "status": "candidate",
           "price": 4000,
           "original_price": 5000,
           "max_discounted_price": 4950,
           "min_discounted_price": 1500,
           "stock": {
                 "min": 2,
                 "max": 5
             }

       }
   ],
   "paging": {
       "offset": 0,
       "limit": 50,
       "total": 1
   }
}
```

**Campos de respuesta**

**id**: identificador del ítem.
 **start_date**: fecha que empieza la campaña
 **finish_date**: fecha de finalización de la campaña.
 **status**: estado del ítem en la promoción. [(Ver tabla)](https://developers.mercadolibre.com.ve/es_ar/ofertas-relampago#Estado-del-ítem)
 **price**: precio del ítem dentro de la promoción. En el caso de que el estado del ítem sea candidate hace referencia al precio sugerido.
 **original_price**: precio actual del ítem.
 **max_original_price**: es el precio de menor valor al que se puede ofrecer esa promoción.
 **min_discounted_price**: es el mayor precio permitido para esa promoción (es decir, el menor descuento permitido).
 **stock**: valor informativo sobre el rango de stock mínimo y máximo que se debe asignar a la promoción.

## Estado del ítem

En la siguiente tabla puedes encontrar los posibles estados que pueden tomar los ítems dentro de este tipo de promoción.

| Estado | Descripción |
| --- | --- |
| **candidate** | Candidato para participar en la promoción. |
| **pending** | Promoción programada. |
| **started** | Activo en la oferta. |
| **finished** | Eliminado de la campaña. |

## Indicar ítems

Una vez invitado a participar en una **oferta relámpago**, puedes indicar qué productos candidatos deseas incluir en la misma. Teniendo en cuenta que debes informar el stock que estará disponible para esta promoción. Cuando se agote el stock disponible automáticamente se dará por finalizada la promoción en el ítem.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{ 
   "deal_price":$deal_price,
    "stock": $STOCK,
   "promotion_type":$PROMOTION_TYPE
}' 
https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2 
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{ 
   "stock": 4000,
   "stock":2,
   "promotion_type":"LIGHTNING"
}' 
https://api.mercadolibre.com/seller-promotions/items/MLB3293401743?app_version=v2
```

Respuesta:

```javascript
{
  "price": 14999,
  "original_price": 17000
}
```

**Parámetros**

**deal_price**: precio del ítem en la promoción.
 **original_price**: precio del ítem antes de incluirse en la promoción.
 **stock**: stock del ítem que desees reservar para esta promoción.
 **promotion_type**: tipo de promoción **LIGHTNING** (oferta relámpago).

## Eliminar ítems

Nota:

Una vez activadas, las ofertas no pueden ser eliminadas. Debido a su corta duración, esperamos que el vendedor se comprometa a mantenerlas activas durante su ciclo. Sin embargo, si el vendedor no desea mantener la oferta, puede 

pausar el ítem

.

Con este recurso podrás eliminar la oferta relámpago del ítem.

Llamada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2&promotion_type=$PROMOTION_TYPE
```

Ejemplo:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/MLA632979587?app_version=v2&promotion_type=LIGHTNING
```

**Respuesta: Status 200 OK**

Conoce más sobre [Ofertas relámpago](https://vendedores.mercadolibre.com.ar/nota/ofertas-relampago-liquida-tu-stock-en-pocas-horas/).

**Siguiente**: [Campañas del vendedor](https://developers.mercadolibre.com.ve/es_ar/campanas-del-vendedor).
