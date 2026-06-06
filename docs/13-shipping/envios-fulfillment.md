---
title: Envíos Fulfillment
slug: envios-fulfillment
section: 13-shipping
source: https://developers.mercadolibre.com.ve/es_ar/envios-fulfillment
captured: 2026-06-06
---

# Envíos Fulfillment

> Source: <https://developers.mercadolibre.com.ve/es_ar/envios-fulfillment>

## Endpoints referenced

- `http://api.mercadolibre.com/stock/fulfillment/operations/$OPERATION_ID`
- `http://api.mercadolibre.com/stock/fulfillment/operations/329663159`
- `http://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=$SELLER_ID&inventory_id=$INVENTORY_ID&date_from=$aaammdd&date_to=$aaammdd`
- `https://api.mercadolibre.com/inventories/$INVENTORY_ID/stock/fulfillment`
- `https://api.mercadolibre.com/inventories/$INVENTORY_ID/stock/fulfillment?include_attributes=conditions`
- `https://api.mercadolibre.com/inventories/LCQI05831/stock/fulfillment`
- `https://api.mercadolibre.com/inventories/YLXH33638/stock/fulfillment?include_attributes=conditions`
- `https://api.mercadolibre.com/items/$ITEM_ID`
- `https://api.mercadolibre.com/items/MLB1557246024`
- `https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=$SELLER_ID&inventory_id=$INVENTORY_ID&date_from=$aaammdd&date_to=$aaammdd`
- `https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=$SELLER_ID&inventory_id=$INVENTORY_ID&date_from=$aaammdd&date_to=$aaammdd&scroll=YXBpY29yZS1pdGVtcw==:ZHMtYXBpY29yZS1pdGVtcy0wMQ==:DXF1ZXJ5QW5kRmV0Y2gBAAAAABIu7AgWMXl6anF3SU5SMVNaQXFxTkZubHBqQQ==`
- `https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=384324657&inventory_id=DEHW09303&date_from=2020-06-01&date_to=2020-06-30`
- `https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=384741716&inventory_id=NFWV18668&date_from=2020-06-29&date_to=2020-07-28&type=SALE_CONFIRMATION&external_references.shipment_id=1111?`

## Content

Última actualización 11/02/2026

## Envíos Fulfillment

Importante:

 Actualmente, esta modalidad de envío está disponible para vendedores de Argentina, Brasil, México, Chile y Colombia.

Envíos Full es un servicio integral de Mercado Libre diseñado para simplificar y optimizar la gestión logística de los vendedores. Con Envíos Full, no solo nos encargamos de tus envíos, sino que también almacenamos tu stock y preparamos cada pedido para su envío inmediato al concretar una venta.

Conoce más sobre:

- [Qué es Envíos Full](https://www.mercadolibre.com.ar/ayuda/que-es-mercado-envios-full_5162)
- [Cómo planificar tus envíos a Full](https://vendedores.mercadolibre.com.ar/nota/guia-para-tus-envios-a-full/)
- [Cómo gestionar tu stock de Full](https://www.mercadolibre.com.ar/ayuda/16958)
- [Cómo envío mi stock a Full](https://www.mercadolibre.com.ar/ayuda/Guia-para-tu-primer-envio-de-stock-a-Full_5163)
- [Cómo retirar productos de Full](https://www.mercadolibre.com.ar/ayuda/15999)
- [Cuáles son los costos de Full](https://www.mercadolibre.com.ar/ayuda/20522)
- [Cómo pagar los cargos de Full](https://www.mercadolibre.com.ar/ayuda/17597)
- [Cargos por descartar stock de Full](https://www.mercadolibre.com.ar/ayuda/21881)
- [Cómo ofrecer Envíos Flex y Full en la misma publicación](https://www.mercadolibre.com.ar/ayuda/ofrecer-Flex-y-Full_4980)
- [Qué productos está prohibido enviar a Full](https://www.mercadolibre.com.ar/ayuda/5200)
- [Cuáles son los beneficios de usar Envíos Full](https://www.mercadolibre.com.ar/ayuda/cual-es-el-costo-de-usar-mercado-envios-full_5171)

## Obtener el inventory_id

Para consultar el stock y las operaciones del ítem en fulfillment, primero debes obtener el inventory_id que es el código que identifica al ítem cuando está en fulfillment. Para esto consulta el inventory_id a través del recurso de /items.

Nota:

Cuando el artículo tenga variaciones, tendrá una identificación 

inventory_id

 por variación.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLB1557246024
```

Respuesta:

```javascript
{
  "id": "MLB1557246024",
  "site_id": "MLB",
  "title": "Kit Capa Chuva Test 228",
  "subtitle": null,
  "seller_id": 384324657,
  "category_id": "MLB22675",
  "official_store_id": null,
  "price": 87,
  "base_price": 87,
  "original_price": null,
  "inventory_id": "LCQI05831",
  "currency_id": "BRL",
  "initial_quantity": 50,
  "available_quantity": 50,
  "sold_quantity": 0,
  "sale_terms": []
}
```

## Consultar el stock del vendedor

Además, puedes consultar el stock total de un vendedor en todos los depósitos de Fulfillment y conocer el estado de las piezas no disponibles.

Nota:

Recuerdas que solo disponemos de la información correspondiente a los últimos 12 meses, considerando el día actual de la consulta.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/inventories/$INVENTORY_ID/stock/fulfillment
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/inventories/LCQI05831/stock/fulfillment
```

Respuesta:

```javascript
{
    "inventory_id": "LCQI05831",
    "total": 20,
    "available_quantity": 5,
    "not_available_quantity": 15,
    "not_available_detail":[
        {
            "status": "damage",
            "quantity": 2
        },
        {
            "status": "lost",
            "quantity": 1
        },
        {
            "status": "noFiscalCoverage"
            "quantity": 5
        },
        {
            "status": "withdrawal",
            "quantity": 5
        },
        {
            "status": "internal_process",
            "quantity": 1
        },
        {
            "status": "transfer",
            "quantity": 1
        }
    ],
    "external_references": [
        {
        "type": "item",
        "id": "MLB1557246024",
        "variation_id": 4742223403
      }
   ]
}
```

### Campos de la respuesta

**total**: es la suma de los campos available_quantity y not_available_quantity.
 **available_quantity**: cantidad de ítems disponibles para la venta.
 **not_available_quantity**: total de ítems que no se encuentran disponibles para la venta.
 **not_available_detail**: detalle del status de los ítems que no se encuentran disponibles.

- **damaged**: total de ítems dañados (incluye damaged seller, meli y carrier).
- **lost**: total de ítems que se perdieron y no fueron encontrados.
- **withdrawal**: total de ítems reservados para retiro.
- **internal_process**: total de ítems reservados por procesos de calidad del depósito.
- **transfer**: total reservado para ser transferido de depósito.
- **noFiscalCoverage**: total de ítems que no se encuentran a la venta por no poseer cobertura fiscal.
- **not_supported**: todos los ítems ingresados son no identificables o no procesables.

external_references

: información de la relación del inventory con la publicación de marketplace, y una identificación del tipo.

type

: tipo de relación entre la publicación y el stock almacenado.

id

: identificador del ítem relacionado con el inventory.

variation_id

: identificador de la variación asociada al inventory.

### Manejo de errores

Respuesta con error:

```javascript
{
    "status": 403,
    "error": "forbidden",
    "message": "User 281349747 cannot access to seller_product ESZJ28231",
    "cause": []
}
```

### Posibles errores

| Status | Mensaje | Error | Descripción |
| --- | --- | --- | --- |
| 404 | Inventory not found with id: ESZJ28232 | seller_product_not_found | No se encuentra el vendedor product solicitado |
| 400 | The field inventory_id has an invalid value | validation_error | Parámetro inválido |
| 403 | The caller is not authorized to access this resource | forbidden | El caller no está autorizado a acceder al recurso |
| 401 | No autorizado | unauthorized | El caller no está autenticado en la plataforma |
| 429 | Too many request | too_many_request | El usuario ha superado el número de request permitidas por minuto |
| 500 | Internal server error | internal_error | Error interno al obtener la información |

## Consultar detalle del stock no disponible

Este recurso devuelve información adicional sobre las unidades almacenadas en full que no están disponibles a la venta, describiendo ciertos atributos o condiciones particulares de estas.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/inventories/$INVENTORY_ID/stock/fulfillment?include_attributes=conditions
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/inventories/YLXH33638/stock/fulfillment?include_attributes=conditions
```

Respuesta:

```javascript
{
  "inventory_id": "YLXH33638",
  "total": 20,
  "available_quantity": 5,
  "not_available_quantity": 15,
  "not_available_detail": [
    {
      "status": "damaged",
      "quantity": 2,
      "conditions": [
        {
          "condition": "arrived_damaged",
          "quantity": 1
        },
        {
          "condition": "damaged_in_full",
          "quantity": 1  (meli+carrier)
        }
      ]
    },
    {
      "status": "lost",
      "quantity": 1
    },
    {
      "status": "not_supported",
      "quantity": 3,
      "conditions": [
        {
          "condition": "dimensions_exceeds",
          "quantity": 1
        },
        {
          "condition": "flammable",
          "quantity": 1
        },
        {
          "condition": "multiple_identifier",
          "quantity": 1
        }
      ]
    }
  ]
}
```

Los posibles status son damaged, not_supported, lost, withdrawal, no_fiscal_coverage, internal_process y transfer.

Los status con información adicional son:

**Damaged**: Brinda información sobre las unidades dañadas almacenadas en el depósito que permite a los integradores accionar según se trate de dañados en full o si llegaron dañados.
 **Not supported**: Brinda información de productos not supported (NS) en stock, es decir, productos que no pueden ponerse a la venta por algún problema de identificación o por pertenecer a categorías no aptos.

Conoce las condiciones por las cuales identificamos un producto como dañado o no cumple con los requisitos para ingresar al centro de Fulfillment. También aplica para productos ingresados y que posteriormente detectamos un problema.

| Status Not available | Condition | Descripción del producto |
| --- | --- | --- |
| **damaged** | arrived_damaged | Llegó dañado por el vendedor |
| damaged_in_full | Está dañado dentro del depósito o por el transportista (carrier) |  |
| **not supportted** | dimensions_exceeds | Excede las dimensiones de almacenamiento permitidas del centro de Fulfillment |
| expiration_problem | Tuvo un problema relacionado a su caducidad |  |
| package_problem | No tiene packaging o este no es apropiado para el centro de Fulfillment |  |
| flammable | Es inflamable o explosivo |  |
| regulation_problem | No cumple las regulaciones por el tipo de producto, por ejemplo, sello sanidad |  |
| other | Toda condición no especificada en los puntos anteriores |  |
| multiple_identifier | Tiene el código universal duplicado (EAN) |  |
| empty_identifier | No tiene identificación/etiqueta del vendedor o no tiene EAN en la base de datos del centro de Fulfillment |  |
| multiple_sku | Tiene dos o más SKUs de Mercado Libre |  |
| invalid_identifier | Tiene el código universal incorrecto |  |
| return_problem | Fue devuelto por el comprador por no cumplir con la calidad especificada en la venta |  |

## Consultar operaciones

Obtén el listado de las operaciones de stock para un inventory_id en particular.

### Parámetros

**inventory_id**: lista de identificadores separados por coma.
 **seller_id**: identificador del vendedor.
 **date_from**: fecha de inicio de la búsqueda. Si no lo defines en el GET, por default son 15 días.
 **date_to**: fecha final de búsqueda. Si no lo defines en el GET, por default es la fecha actual.
 **type**: tipo de operación (inbound_reception, sale_confirmation, etc.)
 **external_references**

- **external_references.shipment_id**: identificador del envío al comprador.

**limit**: cantidad de registros a devolver por “página” de resultados.
 **sort**: identificador del campo y orden de búsqueda.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=$SELLER_ID&inventory_id=$INVENTORY_ID&date_from=$aaammdd&date_to=$aaammdd
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=384324657&inventory_id=DEHW09303&date_from=2020-06-01&date_to=2020-06-30
```

Ejemplo con filtros:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/stock/fulfillment/operations/search?seller_id=384741716&inventory_id=NFWV18668&date_from=2020-06-29&date_to=2020-07-28&type=SALE_CONFIRMATION&external_references.shipment_id=1111?
```

Respuesta:

```javascript
{
    "paging": {
        "total": 4,
        "scroll": ""
    },
    "results": [
        {
            "id": 306811273,
            "seller_id": 384324657,
            "inventory_id": "DEHW09303",
            "date_created": "2020-06-18T18:43:26Z",
            "type": "ADJUSTMENT",
            "detail": {
                "available_quantity": -5,
                "not_available_quantity": 5,
                "not_available_detail": [
                    {
                        "status": "lost",
                        "quantity": 5
                    }
                ]
            },
            "result": {
                "total": 100,
                "available_quantity": 95,
                "not_available_quantity": 5,
                "not_available_detail": [
                    {
                        "status": "lost",
                        "quantity": 5
                    }
                ]
            },
            "external_references": []
        },
        {
            "id": 306745917,
            "seller_id": 384324657,
            "inventory_id": "DEHW09303",
            "date_created": "2020-06-18T18:15:13Z",
            "type": "SALE_CANCELATION",
            "detail": {
                "available_quantity": 10,
                "not_available_detail": []
            },
            "result": {
                "total": 100,
                "available_quantity": 100,
                "not_available_quantity": 0,
                "not_available_detail": []
            },
            "external_references": [
                {
                    "type": "shipment_id",
                    "value": "28312959315"
                }
            ]
        },
        {
            "id": 306718974,
            "seller_id": 384324657,
            "inventory_id": "DEHW09303",
            "date_created": "2020-06-18T18:02:33Z",
            "type": "SALE_CONFIRMATION",
            "detail": {
                "available_quantity": -10,
                "not_available_detail": []
            },
            "result": {
                "total": 90,
                "available_quantity": 90,
                "not_available_quantity": 0,
                "not_available_detail": []
            },
            "external_references": [
                {
                    "type": "shipment_id",
                    "value": "28312961122"
                }
            ]
        },
        {
            "id": 306705012,
            "seller_id": 384324657,
            "inventory_id": "DEHW09303",
            "date_created": "2020-06-18T17:55:42Z",
            "type": "INBOUND_RECEPTION",
            "detail": {
                "available_quantity": 100,
                "not_available_detail": []
            },
            "result": {
                "total": 100,
                "available_quantity": 100,
                "not_available_quantity": 0,
                "not_available_detail": []
            },
            "external_references": [
                {
                    "type": "inbound_id",
                    "value": "0001"
                }
            ]
        }
    ],

    "filters": [],
    "available_filters": [],
    "available_sort": [],
    "sort": [], 
    "available_sorts": []
}
```

### Reglas

- En el search de operaciones, el rango máximo de consulta son 60 días
- El filtro de fechas no es obligatorio, pero si no las pones trae los últimos 15 días
