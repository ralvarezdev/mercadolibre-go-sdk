---
title: Campañas tradicionales
slug: deals
section: 14-promotions
source: https://developers.mercadolibre.com.ve/es_ar/deals
captured: 2026-06-06
---

# Campañas tradicionales

> Source: <https://developers.mercadolibre.com.ve/es_ar/deals>

## Endpoints referenced

- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&promotion_id=$PROMOTION_ID&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3295112047?app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/items/MLB3295112047?promotion_type=DEAL&promotion_id=P-MLB1806019=&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1806019/items?promotion_type=DEAL&app_version=v2`
- `https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1806019?promotion_type=DEAL&app_version=v2`

## Content

Última actualización 13/03/2026

# Campañas tradicionales

 Importante: 

 El 

nuevo filtro por estado

 ya está disponible para filtrar los ítems de una campaña mediante el query param 

status_item

, que acepta los valores "active" o "paused". 

La **campaña tradicional**, conocida como **DEAL**, es un tipo de promoción organizada por Mercado Libre, en la cual los vendedores invitados pueden ofrecer sus productos con precios promocionales.

Los vendedores son invitados periódicamente por Mercado Libre para participar en campañas DEAL. Si aceptan la invitación, pueden incluir productos y definir precios promocionales dentro de los parámetros establecidos por la plataforma.

## Consultar detalles de una campaña

Para obtener los detalles de una oferta del tipo DEAL, utiliza el siguiente endpoint:

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1806019?promotion_type=DEAL&app_version=v2
```

Respuesta:

```javascript
{
  "id": "P-MLB1806019",
  "type": "DEAL",
  "status": "started",
  "start_date": "2023-04-20T03:00:00Z",
  "finish_date": "2023-08-01T02:00:00Z",
  "deadline_date": "2023-08-01T01:00:00Z",
  "name": "HOTSALE"
}
```

## Estados

Estos son los distintos estados por lo que puede pasar una campaña tradicional.

| Estado | Descripción |
| --- | --- |
| **pending** | Promoción aprobada que aún no inició. |
| **started** | Promoción activa. |
| **finished** | Promoción finalizada. |

## Consultar items de una campaña

Para conocer los items que forman parte de una campaña tradicional, utiliza el siguiente endpoint

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/promotions/P-MLB1806019/items?promotion_type=DEAL&app_version=v2
```

Respuesta:

```javascript
{
  "results": [
      {
          "id": "MLB3538191898",
          "status": "candidate",
          "price": 0,
          "original_price": 5000,
          "max_discounted_price": 4800,
          "suggested_discounted_price": 4150,
          "min_discounted_price": 1600,
          "start_date": "2024-11-27T00:00:00",
          "end_date": "2024-12-05T00:00:00",
          "sub_type": "FLEXIBLE_PERCENTAGE"
   },
      {
          "id": "MLB3538191900",
          "status": "started",
          "price": 1000,
          "original_price": 1500,
          "max_discounted_price": 1300,
          "suggested_discounted_price": 1200,
          "min_discounted_price": 1000,
          "start_date": "2024-12-01T00:00:00",
          "end_date": "2024-12-10T00:00:00",
          "sub_type": "FIXED_AMOUNT",
          "top_deal_price": 1100,
          "discount_percentage": 33.33,
          "currency": "BRL"
      },
      {
          "id": "MLB3538191901",
          "status": "pending",
          "price": 0,
          "original_price": 2000,
          "max_discounted_price": 1800,
          "suggested_discounted_price": 1700,
          "min_discounted_price": 1500,
          "start_date": "2024-12-05T00:00:00",
          "end_date": "2024-12-15T00:00:00",
          "sub_type": "FLEXIBLE_PERCENTAGE",
          "currency": "BRL"
      }
  ],
  "paging": {
      "offset": 0,
      "limit": 50,
      "total": 2
  }
}
```

### Parámetros

**id**: id del item.
 **status**: estado del item en la campaña.
 **price**: precio del item en la campaña.
 **original_price**: precio del item sin descuento.
 **min_discounted_price**: precio mínimo que el vendedor puede colocar en el item en la campaña. Es decir, es el mayor descuento que el item puede tener.
 **max_discounted_price**: Precio máximo de descuento considerado creíble.
 **suggested_discounted_price**: Sugerencia de precio promocional atractivo. Puede ser null si no hay una sugerencia confiable para el item.
 **top_deal_price**: precio del elemento para los mejores compradores, con nivel 3 y 6 de Mercado Puntos. Este campo aparecerá solo si el elemento está activo en la campaña y si el vendedor estableció un valor para el mismo en el momento en que agregó el elemento a la campaña.

 Importante: 

Comenzaremos a enviar oficialmente los siguientes campos a partir del 31 de marzo:

- **max_discounted_price**
- **min_discounted_price**
- **suggested_discounted_price**

Se generará un error 

400

 si el valor de 

deal_price

 informado al asociar ítems a una campaña, no corresponde con los descuentos sugeridos.

## Estado de los ítems

En la siguiente tabla puedes encontrar los posibles estados que pueden tomar los ítems dentro de este tipo de campaña.

| Status | Descripción |
| --- | --- |
| **candidate** | Ítem elegible a la deal. |
| **pending** | Ítem enviado a deal pero no se inició. |
| **started** | Ítem con la deal ya iniciada. |
| **finished** | Ítem eliminado de la campaña. |

## Sugerencia de Descuentos para Promociones

La funcionalidad **Sugerencia de Descuentos** permite que los vendedores reciban sugerencias automáticas de precios promocionales, con el fin de aumentar la competitividad de sus ofertas. Esta sugerencia considera el historial de precios y otros factores internos para generar un descuento adecuado.

El recurso **seller-promotions/promotions** permite obtener información sobre los elementos candidatos para una campaña de promoción.

## **Tipos de campañas soportadas**

- Campañas tradicionales
- Descuento individual
- Campañas del vendedor

 Importante: 

 Los nuevos campos no serán retornados para ofertas ya creadas. 

## Indicar ítems para una campaña

Una vez invitado a participar en una campaña tradicional, puedes indicar qué productos deseas incluir en la misma.

Nota:

Es opcional informar el precio para 

top_deal_price

. 

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{
  "top_deal_price":$TOP_DEAL_PRICE
  "promotion_id":"$PROMOTION_ID"
   "deal_price":$DEAL_PRICE,
   "promotion_type":"$PROMOTION_TYPE"
}'
https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
-d '{
  "deal_price": 4000,
  "top_deal_price": 3000,
  "promotion_id": "P-MLB1806019",
  "promotion_type": "DEAL"
 }' 
https://api.mercadolibre.com/seller-promotions/items/MLB3295112047?app_version=v2
```

Respuesta:

```javascript
{
  "price": 4000,
  "top_price": 3000,
  "original_price": 5000
}
```

Respuesta con error:

Error 400: Se produce cuando el **deal_price** informado no cumple con los requisitos para el "Precio Sugerido".

```javascript
{
  {
    "message": "Errors: ERROR_CREDIBILITY_DISCOUNTED_PRICE - The discounted price is not credible.",
    "error": "bad_request",
    "status": 400,
    "cause": [
      {
        "error_code": "ERROR_CREDIBILITY_DISCOUNTED_PRICE",
        "error_message": "The discounted price is not credible."

   }
   
}
```

### Parámetros

**deal_price**: precio del ítem en la promoción.
 **top_deal_price**: precio del ítem para los mejores compradores, con nivel Mercado Puntos 3 a 6 (es opcional informar este precio)
 **promotion_id**: identificación de la promoción.
 **promotion_type**: tipo de promoción (DEAL.)

## Modificar ítems

Para modificar los ítems que están participando en una promoción realiza la siguiente operación:

Llamada:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN'
-d'{
   "deal_price":$DEAL_PRICE,
   "top_deal_price":$TOP_DEAL_PRICE,
   "promotion_id":"$PROMOTION_ID"
   "promotion_type":"DEAL"
}'
https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?app_version=v2
```

Ejemplo:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN'
-d'{
  "deal_price": 3900,
  "top_deal_price": 3000,
  "promotion_id": "P-MLB1806019",
  "promotion_type": "DEAL"
 }'
https://api.mercadolibre.com/seller-promotions/items/MLB3295112047?app_version=v2
```

Respuesta:

```javascript
{
  "price": 3900,
  "top_price": 3000,
  "original_price": 5000
}
```

## Eliminar ítems

Con este recurso podrás eliminar la oferta del ítem.

Llamada:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/$ITEM_ID?promotion_type=$PROMOTION_TYPE&promotion_id=$PROMOTION_ID&app_version=v2
```

Ejemplo:

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/seller-promotions/items/MLB3295112047?promotion_type=DEAL&promotion_id=P-MLB1806019=&app_version=v2
```

Respuesta: **Status 200 OK**

## Posibles Mensajes de Error

Al interactuar con la API, es importante estar al tanto de los mensajes de error que pueden ocurrir, especialmente en casos de solicitudes inválidas o falta de acceso. En caso de problemas, se recomienda verificar los permisos de acceso y los parámetros de solicitud, además de mantener el token de autenticación actualizado.

| Código de Error | Mensaje de Error | Descripción |
| --- | --- | --- |
| **400** | **Bad Request** | La solicitud es inválida o está malformada. Verifique los parámetros o el cuerpo de la solicitud. |
| **401** | **Unauthorized** | El token de autenticación proporcionado es inválido o ha expirado. Solicite un nuevo token. |
| **403** | **Forbidden** | El acceso a la API está prohibido para el usuario o para el tipo de operación solicitada. |
| **404** | **Not Found** | El recurso solicitado no se encontró. Verifique el endpoint o el ID del recurso. |
| **422** | **Unprocessable Entity** | El servidor entiende la solicitud, pero no puede procesarla debido a datos inválidos o inconsistentes. |
| **429** | **Too Many Requests** | El número de solicitudes realizadas excedió el límite permitido. Intente nuevamente más tarde. |
| **500** | **Internal Server Error** | Ocurrió un error inesperado en el servidor. Intente nuevamente más tarde. |
| **503** | **Service Unavailable** | El servicio está temporalmente fuera de servicio. Intente nuevamente más tarde. |

**Next**: [Campañas co-fondeadas](https://developers.mercadolibre.com.ve/es_ar/campanas-co-fondeadas)
