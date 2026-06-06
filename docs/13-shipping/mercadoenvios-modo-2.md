---
title: Mercado Envíos 2
slug: mercadoenvios-modo-2
section: 13-shipping
source: https://developers.mercadolibre.com.ve/es_ar/mercadoenvios-modo-2
captured: 2026-06-06
---

# Mercado Envíos 2

> Source: <https://developers.mercadolibre.com.ve/es_ar/mercadoenvios-modo-2>

## Endpoints referenced

- `http://api.mercadolibre.com/catalog_domains/$DOMAIN_ID/shipping_attributes`
- `https://api.mercadolibre.com/catalog_domains/MLB-AUTOMOTIVE_TIRES/shipping_attributes`
- `https://api.mercadolibre.com/categories/MCO7159/shipping_preferences`
- `https://api.mercadolibre.com/categories/MLB278114/attributes`
- `https://api.mercadolibre.com/items`
- `https://api.mercadolibre.com/orders/$ORDER_ID`
- `https://api.mercadolibre.com/orders/$ORDER_ID?options`
- `https://api.mercadolibre.com/orders/2053577644`
- `https://api.mercadolibre.com/shipment_labels?shipment_ids=$SHIPPING_ID1,$SHIPPING_ID2&response_type=pdf`
- `https://api.mercadolibre.com/shipment_labels?shipment_ids=$SHIPPING_ID1,$SHIPPING_ID2&response_type=zpl2`
- `https://api.mercadolibre.com/shipment_labels?shipment_ids=43308302844&response_type=pdf`
- `https://api.mercadolibre.com/shipment_labels?shipment_ids=43308302844&response_type=zpl2`
- `https://api.mercadolibre.com/shipments/$SHIPMENT_ID`
- `https://api.mercadolibre.com/shipments/40173236996`
- `https://api.mercadolibre.com/shipments/43308302844`
- `https://api.mercadolibre.com/users/$USER_ID/shipping_preferences`

## Content

Última actualización 15/07/2025

## Gestionar Mercado Envíos 2

Con estas guías te ayudaremos a publicar un producto con Mercado Envíos 2 y manejar todo el proceso de envío utilizando los recursos de nuestra API. Recuerda que las dimensiones de los paquetes son estipuladas por Mercado Libre y no pueden ser manipuladas por el usuario. Si deseas activar Mercado Envíos 2, puedes hacer desde cada país (Argentina, Brasil, Colombia, México, Chile, Uruguay, Perú y Ecuador).

## Tipos de logísticas

Los diferentes tipos de logísticas de envíos para Mercado Envíos 2 son:

- Mercado Envíos Drop_off
- [Mercado Envíos Colectas (cross_docking) y Places (xd_drop_off)](https://developers.mercadolibre.com.ve/es_ar/envios-colectas-places)
- [Mercado Envíos Flex (self_service)](https://developers.mercadolibre.com.ve/es_ar/envios-flex)

Mercado Envíos Full (fulfillment)

## Agregar ME2 a un ítem

Utiliza POST para publicar el ítem. Asegúrate de informar los [atributos obligatorios requeridos por la categoría](https://developers.mercadolibre.com.mx/es_ar/atributos#atributos-obligatorios) y los atributos requeridos por el dominio.

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'  -H "Content-Type: application/json" -d 
  {
      "title": "Item de teste",
      "category_id": "MLA91727",
      "price": 1200,
      "currency_id": "ARS",
      "available_quantity": 2,
      "buying_mode": "buy_it_now",
      "listing_type_id": "bronze",
      "condition": "new",
      "description": "test",
      "pictures": [
          {
              "source": "http://upload.wikimedia.org/wikipedia/commons/f/fd/Ray_Ban_Original_Wayfarer.jpg"
          },
          {
              "source": "http://en.wikipedia.org/wiki/File:Teashades.gif"
          }
      ],
     "shipping": {
     "mode": "me2",
     "local_pick_up": false,
     "free_shipping": false,
     "free_methods": []
   }
  }
  https://api.mercadolibre.com/items
```

Recuerda que para publicar en categorías que marcadas como Frágil, el usuario también deberá estar marcado como "frágil", para esto deberá tener un acuerdo comercial. En las siguientes llamadas de la API deberás validar los campos que se muestran a continuación:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/users/$USER_ID/shipping_preferences

{
  "local_pick_up": false,
  "modes": [
    "custom",
    "not_specified",
    "me1",
    "me2"
  ],
  "trusted_user": true,
  "custom_calculator": false,
  "picking_type": "cross_docking",
  "thermal_printer": null,
  "option": "in",
  "tags": [
  ],
  "carrier_pickup": false,
  "items_combination": "enabled",
  "services": [
    311,
    591,
    671,
    801,
    881,
    1181,
    1191,
    136261
  ],
  "logistics": [
    { 
      "mode": "me1",
      "types": [
        {
          "type": "default",
          "carrier_pickup": [],
          "services": [
            21,
            23,
            22,
            11
          ],
          "default": true
        }
      ]
    },

      {"mode": "me2",
      "types": [
        {
          "type": "cross_docking",
          "carrier_pickup": [
            17501840
          ],
          "services": [
            311,
            591,
            671,
            801,
            881,
            1181,
            1191
          ],
          "default": false
        },
        {
          "type": "self_service",
          "carrier_pickup": [
          ],
          "services": [
            136261
          ],
          "default": false
        }
      ]
    },
    {
      "mode": "custom",
      "types": [
        {
          "type": "custom",
          "carrier_pickup": [
          ],
          "services": null,
          "default": true
        }
      ]
    },
    {
      "mode": "not_specified",
      "types": [
        {
          "type": "not_specified",
          "carrier_pickup": [
          ],
          "services": null,
          "default": true
        }
      ]
    }
  ],
  "content_declaration_disabled": false,
  "conciliation": {
    "type": null
  },
  "mandatory_invoice_data": false,
  "site_id": "MLA",
  "free_configurations": [
    {
      "condition": {
        "value": null,
        "type": "all"
      },
      "rule": {
        "default": true,
        "free_mode": "country",
        "value": null
      }
    }
  ],
  "mandatory_settings": {
  }
}
```

"restricted": true (API category)

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/categories/MCO7159/shipping_preferences
{
  "category_id": "MCO7159",
  "dimensions": {
    "weight": 50000,
    "height": 20,
    "width": 60,
    "length": 130
  },
  "logistics": [
    {
      "types": [
        "default"
      ],
      "mode": "me1"
    },
    {
      "types": [
        "drop_off",
        "xd_drop_off",
        "cross_docking",
        "fulfillment"
      ],
      "mode": "me2"
    },
    {
      "types": [
        "not_specified"
      ],
      "mode": "not_specified"
    },
    {
      "types": [
        "custom"
      ],
      "mode": "custom"
    }
  ],
  "restricted": true
}
```

## Atributos requeridos por dominio

Deberás validar cuáles son los atributos que acorde al dominio serán requeridos informar de manera obligatoria **para poder determinar si el ítem es candidato a ser envíado por me2 o no** Estos datos son complementares a los resultados de la API disponible en la llamadaEsses dados são [https://api.mercadolibre.com/categories/MLB278114/attributes](https://api.mercadolibre.com/categories/MLB278114/attributes).

. 

Llamada:

```javascript
curl -X GET 'Authorization: Bearer $ACCESS_TOKEN' http://api.mercadolibre.com/catalog_domains/$DOMAIN_ID/shipping_attributes
```

Ejemplo de llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/MLB-AUTOMOTIVE_TIRES/shipping_attributes
```

Ejemplo de respuesta:

```javascript
{
   "domain_id": "MLB-AUTOMOTIVE_TIRES",
   "attributes": [
       {
           "id": "RIM_DIAMETER",
           "type": "NUMBER_UNIT",
           "unit": "\"",
           "index": 1,
           "ranges": null
       },
       {
           "id": "TIRES_NUMBER",
           "type": "INTEGER",
           "unit": "",
           "index": 2,
           "ranges": null
       },
       {
           "id": "SECTION_WIDTH",
           "type": "NUMBER_UNIT",
           "unit": "mm",
           "index": 3,
           "ranges": null
       }
   ],
   "client_id": 3536736322237473,
   "date_created": "2022-03-29T13:04:27.912-03:00",
   "last_modified": "2023-07-18T11:31:20.092-03:00"
}
```

### Los campos indicarán:

- domain_id: ID del dominio consultado.
- attributes: arreglo que contiene los atributos que deben ser informados de manera obligatoria al momento de crear o modificar un ítem y que ayudarán a determinar si el ítem es candidato a ser envíados por me2.

## Consultar fecha de envío del producto

 Importante: 

Esta funcionalidad está disponible en Argentina, Brasil, Chile y México.

Para evitar sobrepasar la capacidad de los transportistas (carriers) y que los compradores reciban los productos a tiempo, es necesario que consultes la fecha de envío de los productos. Identifica los envíos de este tipo realizando un GET a /shipments, incorporando el header 'X-Format-New: true' verificando el nodo “buffering”.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 'X-Format-New: true' https://api.mercadolibre.com/shipments/$SHIPMENT_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 'X-Format-New: true' https://api.mercadolibre.com/shipments/40173236996
```

Respuesta:

```javascript
{
   "id":40173236996,
   "external_reference":null,
   "status":"pending",
   "substatus":"buffered",
   "date_created":"2020-10-20T10:08:30.000-04:00",
   "last_updated":"2020-10-20T15:09:22.000-04:00",
   "declared_value":7000,
   "dimensions":{
      "height":14,
      "width":19,
      "length":38,
      "weight":950
   },
   "logistic":{
      "direction":"forward",
      "mode":"me2",
      "type":"xd_drop_off"
   },
  []
   "lead_time":{
      "option_id":3628548109,
      "shipping_method":{
         "id":510545,
         "name":"Express a domicilio",
         "type":"two_days",
         "deliver_to":"address"
      },
      "currency_id":"ARS",
      "cost":0,
      "list_cost":504.99,
      "cost_type":"free",
      "service_id":831,
      "delivery_type":"estimated",
      "estimated_schedule_limit":{
         "date":null
      },
      "buffering":{
         "date":"2020-10-21T20:18:26.000Z" ---> Fecha que podrá realizar el envío
      },
      "estimated_delivery_time":{
         "type":"known",
         "date":"2020-10-22T00:00:00.000-03:00",
         "unit":"hour",
         "offset":{
            "date":null,
            "shipping":null
         },
         "time_frame":{
            "from":null,
            "to":null
         },
         "pay_before":"2020-10-21T00:00:00.000-03:00",
         "shipping":24,
         "handling":24,
         "schedule":null
      },
      "estimated_delivery_limit":{
         "date":null,
         "offset":null
      },
      "estimated_delivery_final":{
         "date":null,
         "offset":null
      },
      "estimated_delivery_extended":{
         "date":null,
         "offset":null
      },
      "estimated_handling_limit":{
         "date":"2020-10-21T00:00:00.000-03:00"
      }
   },
   "tags":[
      "test_shipment"
   ]
}
```

En el campo buffering “date” del nodo “buffering” estará la fecha correspondiente que se tiene que despachar el paquete y ese mismo día disponibilizaremos la etiqueta para la impresión.

Nota:

Para las 

ventas con envíos DropShipping

, 

Cross Docking

 y 

Cross Docking Drop Off

, si el substatus es 

“buffered”

 deberás chequear el nodo “buffering” y comunicarle al vendedor que podrá imprimir la etiqueta en la fecha mencionada en el campo “date”.

## Imprimir etiquetas de envío

Cuando un comprador completa su compra, es crucial que el vendedor pueda generar rápidamente la etiqueta de envío para agilizar el proceso de despacho.

### Validaciones de tipos de logística:

Antes de intentar obtener la etiqueta, es esencial verificar los campos "mode" y "type" dentro del nodo logístico para garantizar que la etiqueta esté disponible.

**Mode**: siempre debe ser “me2”.
 **Logistic_type**: indica los tipos de logística aceptados tales como:

- “drop_off”
- “xd_drop_off”
- “cross_docking”
- “self_service”
