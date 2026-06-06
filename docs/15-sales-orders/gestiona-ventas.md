---
title: Órdenes
slug: gestiona-ventas
section: 15-sales-orders
source: https://developers.mercadolibre.com.ve/es_ar/gestiona-ventas
captured: 2026-06-06
---

# Órdenes

> Source: <https://developers.mercadolibre.com.ve/es_ar/gestiona-ventas>

## Endpoints referenced

- `https://api.mercadolibre.com/currency_conversions/search?from=$CURRENCY_ID&to=$CURRENCY_ID`
- `https://api.mercadolibre.com/currency_conversions/search?from=ARS&to=BRL`
- `https://api.mercadolibre.com/orders/$ORDER_ID`
- `https://api.mercadolibre.com/orders/$ORDER_ID/discounts`
- `https://api.mercadolibre.com/orders/$ORDER_ID/product`
- `https://api.mercadolibre.com/orders/2000003508419013`
- `https://api.mercadolibre.com/orders/2000003508419013/discounts`
- `https://api.mercadolibre.com/orders/search`
- `https://api.mercadolibre.com/orders/search?seller=$SELLER_ID&order.date_created.from=2015-07-01T00:00:00.000-00:00&order.date_created.to=2015-07-31T00:00:00.000-00:00`
- `https://api.mercadolibre.com/orders/search?seller=$SELLER_ID&order.status=paid`
- `https://api.mercadolibre.com/orders/search?seller=$SELLER_ID&order.status=paid&sort=date_desc`
- `https://api.mercadolibre.com/orders/search?seller=89660613&q=2032217210`

## Content

Última actualización 18/02/2026

## Gestionar órdenes

Una orden es una solicitud que realiza un cliente para una publicación con intención de comprarlo conforme a una serie de condiciones que seleccionará en el flujo del proceso de compra (checkout). Todas las condiciones de la venta se detallan en la orden, la cual se replicará para las cuentas del comprador y el vendedor. Conoce más **el flujo para gestionar órdenes simples y de carrito, pagos y envíos**.

## Obtener una orden

Una vez que se crea una nueva orden en el usuario, se puede consultar los detalles al realizar una solicitud al recurso de órdenes. Además, te recomendamos suscribirte al nuevo tópico [orders feedback](https://developers.mercadolibre.com.ar/es_ar/productos-recibe-notificaciones#Topics-disponibles) para estar actualizado sobre los feedbacks recibidos.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/$ORDER_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/2000003508419013
```

Respuesta:

```javascript
{
    "id": 2000003508419013,
    "status": "paid",
    "status_detail": null,
    "date_created": "2013-05-27T10:01:50.000-04:00",
    "date_closed": "2013-05-27T10:04:07.000-04:00",
    "order_items": [{
   	 "item": {
   		 "id": "MLB12345678",
   		 "title": "Samsung Galaxy",
   		 "variation_id": null,
   		 "variation_attributes": []
   	 },
   	 "quantity": 1,
   	 "unit_price": 499,
   	 "currency_id": "BRL"
    }],
    "total_amount": 499,
    "currency_id": "BRL",
    "buyer": {
   	 "id": "123456789",
       },
    "seller": {
   	 "id": "123456789",
    },
    "payments": [{
   	 "id": "596707837",
   	 "transaction_amount": 499,
   	 "currency_id": "BRL",
   	 "status": "approved",
   	 "date_created": null,
   	 "date_last_modified": null
    }],
    "feedback": {
   	 "purchase": null,
   	 "sale": null
    },
    "context": {
     "channel": "marketplace",
     "site": "MLB",
     "flows": [
     000]
  },
    "shipping": {
   	 "id": 20676482441
    },
    "tags": [
   	 "no_shipping",
   	 "paid",
   	 "not_delivered"
    ]
}
```

 Notas: 

 - Para obtener el detalle del feedback, debes realizar un llamado al recurso 

/feedbacks/$feedback_id

 con el ID obtenido en la orden.

 -También puedes consumir la información de los feedbacks usando el recurso 

/orders/$order_id/feedback

.

 - Es posible obtener las informaciones del vendedor consultando a la 

API de users

 utilizando el access_token. 

### Campos de respuesta:

**id**: identificador único de la orden. 
 **date_created**: fecha de creación de la orden 
 **date_closed**: fecha de confirmación de la orden. Se define cuando una orden cambia por primera vez al estado: confirmed / paid y se descuenta el stock del ítem. 
 **expiration_date**: fecha límite que tiene el usuario para calificar porque, luego de la misma, se vuelve visible el feedback, se emiten los pagos (si hubiese) y se crean los cargos. 
 **status**: estado de la orden. [Ver los valores posibles](https://developers.mercadolibre.com.ve/es_ar/gestiona-ventas#Estado-de-la-orden). 
 **description**: descripción del estado. 
 **buyer**: información del comprador. 
 **seller**: información del vendedor. 
 **order_items**: publicaciones en la orden.

- **item**: publicación específica.
- **cantidad**: cantidad de items comprados.
- **sale_fee**: comisión de ventas.
- **unit_price**: precio unitário.
- **gross_price**: El atributo `gross_price` es un campo que representa el monto original que el cliente hubiese pagado por todas las unidades del ítem sin descuentos. Este campo permite visualizar claramente el impacto de los descuentos aplicados en cada orden.

payments

: pagos relacionados con la orden. 

feedback

: ID del feedback relacionada con la orden.

context

: detalle de las características de la creación de una orden.

- **channel**: los canales de venta que hoy tiene el ítem. Valores posibles: proximity, mp-channel, marketplace.
- **site**: ID del sitio donde se originó la compra (MLA, MLB, MLM, etc)
- **flows**: es una lista de características del origen de la compra. Valores posibles b2b, cbt, subscription, reservation, catalog, contract, supermarket, 3x_campaign, high_concurrency, lite.
