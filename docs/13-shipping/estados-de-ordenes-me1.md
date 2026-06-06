---
title: Estados de órdenes y seguimiento
slug: estados-de-ordenes-me1
section: 13-shipping
source: https://developers.mercadolibre.com.ve/es_ar/estados-de-ordenes-me1
captured: 2026-06-06
---

# Estados de órdenes y seguimiento

> Source: <https://developers.mercadolibre.com.ve/es_ar/estados-de-ordenes-me1>

## Endpoints referenced

- `https://api.mercadolibre.com/orders/$ORDER_ID/shipments`
- `https://api.mercadolibre.com/orders/2000003508419013/shipments`
- `https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications`
- `https://api.mercadolibre.com/shipments/28264263908/seller_notifications`

## Content

Última actualización 28/05/2026

## Estados de órdenes y seguimiento

El nuevo recurso Estados de órdenes ME1 tiene el objetivo de mejorar la experiencia de los compradores en el acompañamiento de la entrega de los productos. Podrá informarse cuando el producto fuera enviado, si la entrega fue exitosa o no, además del número de seguimiento (tracking number).

## Estados y sub estados de envío

La fusión de la información del campo status y el subestado de envío determina en el site se está leyendo qué se notificará a los compradores. Ahora es posible enviar la información de la compra despachada (shipped) o la entrega fallida (not_delivered):

| Estado | Subestado | Descripción |
| --- | --- | --- |
| shipped | null | Despachado |
| shipped | out_for_delivery | Salió para entrega al comprador |
| shipped | receiver_absent | Visita fallida por comprador ausente |
| shipped | dangerous_area | Visita fallida por zona peligrosa |
| shipped | bad_address | Visita fallida por domicilio incorrecto |
| shipped | unauthorized_receiver | Visita fallida por persona no autorizada para recibir |
| not_delivered | returning_to_sender | No entregado - E devolución al vendedor |
| delivered | null | Entregado al comprador |

## Actualizar el estado de un envío ME1

Para actualizar el estado del envío es necesario que conozcas el shipment_id de la orden. Para obtenerlo, consulta al recurso order.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/$ORDER_ID/shipments
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/2000003508419013/shipments
```

Respuesta:

```javascript
{
    "id": 28264263908,
    "mode": "me1",
    "created_by": "receiver",
    "order_id": 2000003508419013,
    "order_cost": 99.9,
    "base_cost": 22.07,
    "site_id": "MLB",
    "status": "pending",
    "substatus": null,
    ...
}
```

## Marcar compra como despachada

Para marcar la compra despachada es necesario informar el estado como "shipped" y el subestado como "null". Eso es definido cuando la transportadora ha recogido el paquete del vendedor.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json'  \
{
  "payload": {
             "service_id": 154,
    "comment": "despachado",
    "date": "2023-01-16T13:03:51.175-04:00"
  },
  "tracking_number": "OP123456789AR",
  "tracking_url": "http://www.url.test/40886674732",
  "status": "shipped",
  "substatus": "null"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

 Nota: 

Utiliza los siguientes service_id para informar el número de seguimiento (tracking), este representa a ME1:

 - MLB: 11

 - MLA: 154

 - MLM: 231876

 - MLC: 282578

 - MCO: 282579

 - MLU: 282604

 - MPE: 361180

## Marcar compra como salió para entrega

La transportadora comenzó el viaje con destino al comprador y lo visitará durante el día. La promesa de entrega se actualiza automáticamente a "Llega hoy". Puede tener otros estados de eventos fallidos (leer: comprador ausente, domicilio incorrecto, etc). Para notificarlo use el estado "shipped" con el subestado "out_for_delivery".

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json'  \
{
  "payload": {
             "service_id": 154,
    "comment": "salió para entrega",
    "date": "2023-01-16T13:03:51.175-04:00"
  },
  "tracking_number": "OP123456789AR",
  "tracking_url": "http://www.url.test/40886674732",
  "status": "shipped",
  "substatus": "out_for_delivery"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar compra como no entregada

El status "not_delivered" es un estado final e irreversible. Solo debe ser utilizado cuando no hubiera más intentos de entrega. De esa forma, el vendedor tiene que alinear el flujo para que la devolución del dinero del comprador sea realizada. 
 Para marcar la compra como no entregada, debes informar el estado como "not_delivered" y el subestado como "returning_to_sender".

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json'  \
{
   "payload":{
      "comment":"Não entregue",
      "date":"2020-03-05T16:17:51.175-04:00"
   },
   "status":"not_delivered",
   "substatus":"returning_to_sender"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar compra como entregada

Al recibir la información de que un producto fue entregado al comprador, debes realizar un cambio en el estado de la compra para entregada. Para eso, utiliza el estado "delivered" con el subestado "null".
Este status también es finalizador e irreversible.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json'  \
 {
   "payload":{
      "comment":"Pedido entregue",
      "date":"2020-03-06T16:17:51.175-04:00"
   },
   "status":"delivered",
   "substatus":"null"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar entrega como fallida por comprador ausente

El comprador no se encuentra en su domicilio residencial o no responde al timbre al momento de ser visitado.

Posibles soluciones:

- El comprador y vendedor deben coordinar cómo resolver la entrega o cancelar la compra.
- Si se agotaron los intentos de entrega, la transportadora o el vendedor debe marcar No entregado. Recién allí Mercado Libre reembolsará al comprador y el vendedor deberá gestionar la devolución con su transportadora.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' \
{
    "payload":{
        "comment":"Comprador não se encontra no domicílio",
        "date":"2020-03-06T16:17:51.175-04:00"
    },
    "status":"shipped",
    "substatus":"receiver_absent"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar entrega como fallida por zona peligrosa

La transportadora salió a distribución con destino al comprador pero la zona donde se encuentra el domicilio es peligrosa y se decide no ir. También puede aplicar en una planificación antes de salir a destino.

Posibles soluciones:

- El comprador y vendedor deben coordinar cómo resolver la entrega o cancelar la compra.
- Si se agotaron los intentos de entrega, la transportadora o el vendedor debe marcar No entregado. Recién allí Mercado Libre reembolsará al comprador y el vendedor deberá gestionar la devolución con su transportadora.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' \
{
    "payload":{
        "comment":"Região da entrega identificada como zona de risco",
        "date":"2020-03-06T16:17:51.175-04:00"
    },
    "status":"shipped",
    "substatus":"dangerous_area"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar entrega como fallida por domicilio incorrecto

La transportadora no encontró el domicilio de entrega al momento de la visita. Si se agotaron los intentos de entrega Mercado Libre reembolsará al comprador y el vendedor deberá gestionar la devolución con su transportadora.

Posibles soluciones:

- El comprador y vendedor deben coordinar cómo resolver la entrega o cancelar la compra.
- Si se agotaron los intentos de entrega, la transportadora o el vendedor debe marcar No entregado. Recién allí Mercado Libre reembolsará al comprador y el vendedor deberá gestionar la devolución con su transportadora.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' \
{
    "payload":{
        "comment":"O endereço de entrega não foi encontrado",
        "date":"2020-03-06T16:17:51.175-04:00"
    },
    "status":"shipped",
    "substatus":"bad_address"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

## Marcar entrega como fallida por persona no autorizada para recibir el producto

El titular de la compra no se encuentra y la persona que atiende en el domicilio de entrega no está autorizada para recibir en su lugar.

Posibles soluciones:

- El comprador y vendedor deben coordinar cómo resolver la entrega o cancelar la compra.
- Si se agotaron los intentos de entrega, la transportadora o el vendedor debe marcar No entregado. Recién allí Mercado Libre reembolsará al comprador y el vendedor deberá gestionar la devolución con su transportadora.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'
https://api.mercadolibre.com/shipments/$SHIPMENT_ID/seller_notifications
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' \
{
    "payload":{
        "comment":"O titular da compra não se encontra e não há uma pessoa autorizada para receber a entrega",
        "date":"2020-03-06T16:17:51.175-04:00"
    },
    "status":"shipped",
    "substatus":"unauthorized_receiver"
}
https://api.mercadolibre.com/shipments/28264263908/seller_notifications
```

**Siguiente**: [Flete dinámico](https://developers.mercadolibre.com.ve/es_ar/flete-dinamico).
