---
title: Descuento individual
slug: descuento-individual
section: 14-promotions
source: https://developers.mercadolibre.com.ve/es_ar/descuento-individual
captured: 2026-06-06
---

# Descuento individual

> Source: <https://developers.mercadolibre.com.ve/es_ar/descuento-individual>

## Endpoints referenced

- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&app_version=v2&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLA876768946?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLA876768946?promotion_type=PRICE_DISCOUNT&app_version=v2`

## Content

Última actualización 21/03/2025

## Descuento individual

Los vendedores que deseen ofrecer una oferta particular para sus ítems con los siguientes recursos podrán hacerlo. Tendrán la posibilidad de aplicar, eliminar y consultar el descuento.

Nota:

El plazo máximo para este tipo de campaña es de 31 días. A partir del 24/03/2025 pasara a ser de 14 días.

### Para ofrecer este descuento es necesario:

- Tener reputación verde.
- El item debe tener status igual a activo.
- Condición igual a nuevo.
- La exposición del ítem no puede ser gratuita.
- Y sólo para MLA este tipo de descuento no está disponible en las categorías de libros.

## Ofrecer descuento

Para este tipo de oferta, debes cumplir algunos requisitos. Conoce más sobre [cómo ofrecer descuentos](https://www.mercadolibre.com.ar/ayuda/Como-ofrecer-un-descuento_3992).

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{
   "deal_price": $DEAL_PRICE,
   "top_deal_price": $TOP_DEAL_PRICE,
   "start_date": "$START_DATE",
   "finish_date": "$FINISH_DATE",
   "promotion_type": "PRICE_DISCOUNT"
}'
https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{
   "deal_price": 20,
   "top_deal_price": 30,
   "start_date": "2023-04-19T00:00:00",
   "finish_date": "2024-04-20T00:00:00",
   "promotion_type": "PRICE_DISCOUNT"
}'
https://api.mercadolibre.com/seller-promotions/items/MLA876768946?app_version=v2'
```

Respuesta:

```javascript
{
    "price": 70,
    "original_price": 100
}
```

### Parámetros

**deal_price**: precio del item con descuento para todos los compradores.
 **top_deal_price**: precio del item con descuento para los mejores compradores (con Mercado Puntos nível 3 a 6) (opcional)
 **start_date**: fecha de inicio del descuento.
 **finish_date**: fecha de fin del descuento.

### Consideraciones

- Es posible segmentar la oferta de descuentos, estableciendo un precio general para todos los compradores, y uno menor sólo para nuestros compradores leales (con nivel 3 al 6 de Mercado Puntos).
- El descuento general debe ser como mínimo 5% menor al descuento de usuarios de nivel 3 al 6, para descuentos de hasta 35%. Para descuentos superiores, la diferencia debe ser de mínimo 10%, es decir, damos mejores descuentos a los niveles más altos.
- El descuento máximo debe ser menor a 80% y el descuento mínimo a ofrecer deberá ser mayor o igual al 5%.
- Si se realiza una suba del precio del ítem, los descuentos serán quitados automáticamente.
- Si al iniciar el descuento, el ítem se encuentra participando de un DEAL, dicho descuenConsideto no será aplicado hasta que finalice el DEAL asociado.
- El plazo máximo para un descuento PRICE_DESCOUNT es de 31 días. A partir del 24/03/2025 pasara a ser de 14 dias.
- Las fechas de inicio (start_date) y fin del descuento (finish_date) sólo consideran la fecha en sí, independientemente del horario informado. Por defecto, el descuento comienza a las 00:00:00 del día de inicio y termina a las 23:59:59 del día de finalización.

Nota:

Ten en cuenta que para realizar pruebas, es necesario que el usuario de TEST tenga reputación verde y el ítem tenga como mínimo 1 venta con el precio actual.

## Estado del ítem

En la siguiente tabla puedes encontrar los posibles estados que pueden tomar los ítems cuando deseas aplicarle un descuento individual.

| Estado | Descripción |
| --- | --- |
| **started** | Descuento activo en el ítem. |
| **finished** | Descuento finalizado. |
| **pending** | Descuento programado. |
| **sync_requested** | Proceso de activación pendiente. |
| **restore_requested** | Proceso pendiente de eliminación del descuento. |
| **candidate** | Ítem candidato para participar en la promoción. |

## Eliminar descuento individual a un ítem

Llamada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&app_version=v2&app_version=v2
```

Ejemplo:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/MLA876768946?promotion_type=PRICE_DISCOUNT&app_version=v2
```

Recuerda que para las ofertas individuales (PRICE DISCOUNT) eliminarás toda la oferta, no podrás eliminar por niveles del comprador.

**Respuesta: Status 200 OK**

## Errores

Descuentos fuera de los rangos establecidos

```javascript
{
   "key":"buyer_discount_not_in_range",
   "message":"buyers_discount_percentage parameter must be in range (5, 80)"
}
```

```javascript
{
   "key":"best_buyer_discount_not_in_range",
   "message":"buyers_discount_percentage parameter must be in range (5, 80)"
}
```

Diferencias entre descuentos para niveles 1-2 y niveles 3-6 fuera de los márgenes establecidos

```javascript
{
   "key":"discount_below_10_percent_difference",
   "message":"The best buyer discount difference cannot be below 10% when buyers discount is above 35%"
}
```

```javascript
{
   "key":"discount_below_5_percent_difference",
   "message":"The discount difference cannot be below 5%"
}
```

Cuando el descuento no sea suficiente y el vendedor deba aplicar un descuento mayor, retornaremos:

```javascript
{
    "key": "error_credibility_price",
    "message": "The price is not credible."
}
```

Conoce más sobre [Descuentos en tus publicaciones](https://vendedores.mercadolibre.com.ar/nota/como-crear-descuentos-en-tus-publicaciones-de-mercado-libre/).

**Siguiente**: [Ofertas del día](https://developers.mercadolibre.com.ve/es_ar/ofertas-del-dia)
