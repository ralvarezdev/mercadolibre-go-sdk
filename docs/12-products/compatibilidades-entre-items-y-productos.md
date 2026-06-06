---
title: Compatibilidades entre ítems y productos de Autopartes
slug: compatibilidades-entre-items-y-productos
section: 12-products
source: https://developers.mercadolibre.com.ve/es_ar/compatibilidades-entre-items-y-productos
captured: 2026-06-06
---

# Compatibilidades entre ítems y productos de Autopartes

> Source: <https://developers.mercadolibre.com.ve/es_ar/compatibilidades-entre-items-y-productos>

## Endpoints referenced

- `http://api.mercadolibre.com/catalog/dumps/domains/$SITE_ID/compatibilities`
- `https://api.mercadolibre.com/catalog/dumps/domains/MLB/compatibilities`
- `https://api.mercadolibre.com/catalog_compatibilities/products_search/count_family_products`
- `https://api.mercadolibre.com/catalog_compatibilities/restrictions/values?main_domain_id=MLA-CARS_AND_VANS&secondary_domain_id=MLA-VEHICLE_ENGINE_MOUNTS`
- `https://api.mercadolibre.com/catalog_compatibilities/restrictions/values?main_domain_id=MLA-CARS_AND_VANS&secondary_domain_id=MLA-VEHICLE_SHOCK_ABSORBERS`
- `https://api.mercadolibre.com/catalog_domains/MLB-CARS_AND_VANS/compatibilities/cards`
- `https://api.mercadolibre.com/compats-snapshots/orders/$ORDER_ID`
- `https://api.mercadolibre.com/compats-snapshots/orders/2000006372967416`
- `https://api.mercadolibre.com/compats-snapshots/orders/2000006372967424`
- `https://api.mercadolibre.com/compats-snapshots/orders/2000006372967684`
- `https://api.mercadolibre.com/items/$ITEM_ID`
- `https://api.mercadolibre.com/items/$ITEM_ID/compatibilities`
- `https://api.mercadolibre.com/items/$ITEM_ID/compatibilities/$COMPATIBILITY_ID`
- `https://api.mercadolibre.com/items/$ITEM_ID/compatibilities/$COMPATIBILITY_ID/note`
- `https://api.mercadolibre.com/items/$ITEM_ID/compatibilities?extended=true`
- `https://api.mercadolibre.com/items/$item_id/compatibilities/exception`
- `https://api.mercadolibre.com/items/MLA1417560910`
- `https://api.mercadolibre.com/items/MLA794706391/compatibilities`
- `https://api.mercadolibre.com/items/MLB12345678/compatibilities/exception`
- `https://api.mercadolibre.com/items/MLB3863034063/compatibilities`
- `https://api.mercadolibre.com/items/MLB3863097751/compatibilities`
- `https://api.mercadolibre.com/items/MLB4462690924`
- `https://api.mercadolibre.com/items/MLM12456789/compatibilities`
- `https://api.mercadolibre.com/items/MLM12456789/compatibilities/bcbd413f-cd65-0e0f-88c9-5eb4aebb5372/note`
- `https://api.mercadolibre.com/items/MLM794706391/compatibilities`
- `https://api.mercadolibre.com/items/MLM794706391/compatibilities/$COMPATIBILITY_ID`
- `https://api.mercadolibre.com/items/MLM794706391/compatibilities/4cb9af35-8e9b-ebfd-9e7f-2245ac363d10`
- `https://api.mercadolibre.com/items/MLM794706391/compatibilities?extended=true`
- `https://api.mercadolibre.com/items/catalog_domains/$DOMAIN_ID/compatibilities/cards`
- `https://api.mercadolibre.com/items/compatibilities_summary`
- `https://api.mercadolibre.com/user-products/$USER_PRODUCT_ID/compatibilities/exception`
- `https://api.mercadolibre.com/user-products/&USER_PRODUCT_ID/compatibilities?main_domain_id=MLM-CARS_AND_VANS_FOR_COMPATIBILITIES&extended=true`
- `https://api.mercadolibre.com/user-products/MLAU2339836140/compatibilities/exception`
- `https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities`
- `https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities/copy-paste`
- `https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities?main_domain_id=MLM-CARS_AND_VANS_FOR_COMPATIBILITIES&extended=true`
- `https://api.mercadolibre.com/user-products/MLU1443000156/compatibilities/exception`
- `https://api.mercadolibre.com/users/$SELLER_ID/items/search?tags=incomplete_compatibilities`
- `https://api.mercadolibre.com/users/$SELLER_ID/items/search?tags=pending_compatibilities`
- `https://api.mercadolibre.com/users/$USER_ID/items/search?status=active&has_compatibilities=true`
- `https://api.mercadolibre.com/users/1695976736/items/search?tags=pending_compatibilities`
- `https://api.mercadolibre.com/users/1871678972/items/search?tags=incomplete_position_compatibilities`

## Content

Última actualización 31/03/2026

## Compatibilidades entre ítems y productos de Autopartes

 Importante: 

- Este recurso está disponible sólo para Argentina, México, Brasil, Uruguay, Chile y Colombia. 

 - A partir de ahora ya esta disponible la gestión de compatibilidades a través del user products 

Las compatibilidades permiten asociar los ítems publicados a los productos compatibles del marketplace por ejemplo, si tienes un “Amortiguador Corven Plus” publicado podrás definir atributos como Marca, Modelo, Año y Motor para los cuales este repuesto es compatible. De esta manera, mejoras la calidad de las publicaciones y reduces la cantidad de publicaciones por ítem. 
Para esto, deberás acceder al dump y verificar que el dominio de los ítems y el dominio de los productos sean compatibles. Luego, podrás agregar las compatibilidades de 3 (tres) maneras diferentes y por último, listarlas.
Si la compatibilidad no es la adecuada, podrás eliminar aquellas definidas por los usuarios (vendedores).

## Verificar compatibilidad entre dominios

 Nota: 

- Antes de publicar te recomendamos 

validar si la categoría del ítem contiene el atributo categories.required = true

. 

 - Una vez creada la publicación, podrás 

identificar los ítems en donde es obligatorio informar compatibilidades con el atributo tag incomplete_compatibilities

. Puedes ver más detalle en 

identificar ítems que requieren compatibilizarse

. 

Antes de crear compatibilidades entre ítems y productos, debes verificar que los dominios y catagorías de ítems y productos sean compatibles.

Consultando el siguiente dump, obtienes el listado de dominios y categorías en los cuales puede o necesita informar compatibilidades por site.

Llamada:

```javascript
curl -X GET http://api.mercadolibre.com/catalog/dumps/domains/$SITE_ID/compatibilities
```

Ejemplo llamada:

```javascript
curl -X GET https://api.mercadolibre.com/catalog/dumps/domains/MLB/compatibilities
```

 Importante: 

A partir de ahora, la respuesta de este endpoint incluye el atributo restrictions_required, el cual indica si es obligatorio informar posiciones para la categoría correspondiente. 

Ejemplo respuesta:

```javascript
[
[
   {
       "domain_id": "MLB-VEHICLE_ALTERNATOR_BRACKETS",
       "main": false,
       "compatibilities": [
           {
               "compatible_domain_id": "MLB-CARS_AND_VANS",
               "type": "EXTENSION",
               "categories": [
                   {
                       "id": "MLB439572",
                       "required": true,
                       "note_status": "ENABLED",
                       "restrictions_status": "ENABLED",
                       "universal_status": "DISABLED",
                       "display_new_products_status": "ENABLED",
                       "restrictions_required": false,
                       "enabled_clients": [
                           "SUPPLY"
                       ]
                   }
               ]
           }
       ]
   },
   {
       "domain_id": "MLB-AUTOMOTIVE_WIRE_HARNESSES",
       "main": false,
       "compatibilities": [
           {
               "compatible_domain_id": "MLB-CARS_AND_VANS",
               "type": "EXTENSION",
               "required": false,
               "categories": [
                   {
                       "id": "MLB431130",
                       "required": true,
                       "note_status": "ENABLED",
                       "restrictions_status": "ENABLED",
                       "universal_status": "DISABLED",
                       "display_new_products_status": "ENABLED",
                       "restrictions_required": false,
                       "enabled_clients": [
                           "SUPPLY"
                       ]
                   }
               ]
           }
       ]
   },

    …
}]
```

Los nuevos campos indican:

- **categories**: categorías que admiten la carga de compatibilidades.
- **required**: categorías en las que es obligatorio cargar compatibilidades.
- **type**: tipo de compatibilidad. Solo el tipo EXTENSION admite carga de compatibilidades.
- **note_status y restrictions_status**: indican si la categoría permite informar notas y restricciones de posición.
- **universal_status**: indica si la categoría permite informar compatibilidades universales. **ENABLED**: permite informar compatibilidades universales o **DISABLED**: no permite informar compatibilidades universales.
- **restrictions_required**: indica la obligatoriedad de informar la posición.

## Múltiples publicaciones elegibles con multiget

Obtén más información de [dominios, productos y atributos de autopartes](https://developers.mercadolibre.com.ar/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes).

## Contar productos de un dominio

Para consultar la cantidad de productos existentes por dominio (familia de producto) que cumplen con determinados atributos y valores puedes realizar el siguiente POST. Esto te permitirá validar, previo a agregar las compatibilidades, la cantidad de productos y evitar errores en la asignación de compatibilidades.
Esto es importante ya que solo se pueden asignar un máximo de 200 productos por llamado.

Nota:

El límite de tráfico por APP_ID es de 100 rpm.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/products_search/count_family_products
{
  "domain_id": "$domainId",
  "attributes": [{
    "id": "$attributeId1",
    "value_id": "$valueId1"
  }, {
    "id": "$attributeId2",
    "value_name": "$valueName2"
  }]
}
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/products_search/count_family_products
{
  "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
  "attributes": [{
      "id": "BRAND",
      "value_name": "Volkswagen"
    },
    {
      "id": "CAR_AND_VAN_MODEL",
      "value_name": "VENTO"
    }
  ]
}
```

Respuesta:

```javascript
{
   "count":141
}
```

## Identificar ítems que requieren compatibilizarse

Con el siguiente recurso a través del tag incomplete_compatibilities puedes identificar los ítems que requieren informar compatibilidades de manera obligatoria para evitar moderaciones.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLA1417560910
```

Respuesta:

```javascript
{
  "id": "MLA743626587",
  "site_id": "MLA",
  "title": "Paragolpe Trasero Peugeot 307 Linea Nueva 5 Puertas",
  "seller_id": 65711207,
  "category_id": "MLA60954",
  "official_store_id": null,
.
.
.
.

  "tags": [
    "good_quality_picture",
    "brand_verified",
    "loyalty_discount_eligible",
    "good_quality_thumbnail",
    "ahora-paid-by-buyer",
    "incomplete_compatibilities",
    "immediate_payment"
  ],
  "warranty": "CON GARANTIA DEL FABRICANTE",
  "catalog_product_id": null,
  "domain_id": "MLA-VEHICLE_REAR_BUMPERS",
  "parent_item_id": null,
  "deal_ids": [
  ],
  "automatic_relist": false,
  "date_created": "2018-08-17T20:36:54.000Z",
  "last_updated": "2023-12-15T17:18:35.000Z",
  "health": 0.83,
  "catalog_listing": false
}
```

### Filtrar ítems que requieren compatibilizarse

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/users/$SELLER_ID/items/search?tags=incomplete_compatibilities
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/users/$SELLER_ID/items/search?tags=incomplete_compatibilities
```

Respuesta:

```javascript
{
   "seller_id": "1373279576",
   "results": [
       "MLA1417560910",
       "MLA1373565211",
       "MLA1371734513",
       "MLA1396609243"
   ],
   "paging": {
       "limit": 50,
       "offset": 0,
       "total": 4
   },
   "query": null,
.
.
.
.
}
```

## Obtener restricciones de posiciones

 Importante: 

El campo 

combined_values

 proporciona combinaciones predefinidas que facilitan al vendedor la selección de posiciones comunes, mejorando la experiencia de uso de la API.

El objetivo es proporcionar una lista de los valores permitidos para la restricción "POSITION", junto con las combinaciones más frecuentes, con el fin de simplificar el proceso y ayudar al vendedor. Para lograr esto, se añade una lista de combinaciones posibles (**combined_values**). Esta API es útil para obtener las opciones de posición válidas antes de crear o actualizar productos en el catálogo de vehículos y sus accesorios.

Llamada:

```javascript
curl --location -H 'Authorization: Bearer $ACCESS_TOKEN' \
'https://api.mercadolibre.com/catalog_compatibilities/restrictions/values?main_domain_id=MLA-CARS_AND_VANS&secondary_domain_id=MLA-VEHICLE_ENGINE_MOUNTS'
```

Respuesta:

```javascript
{
  "attributes_values": [
    {
      "attribute_id": "POSITION",
      "values": [
        {
          "value_id": "13373175",
          "value_name": "Conductor",
          "opposite_values": [
            "13373176"
          ]
        },
        {
          "value_id": "13373176",
          "value_name": "Acompañante",
          "opposite_values": [
            "13373175"
          ]
        }
      ],
      "combined_values": [
        {
          "values": [
            {
              "value_id": "13701104",
              "value_name": "Delantera"
            },
            {
              "value_id": "2262158",
              "value_name": "Izquierda"
            },
            {
              "value_id": "4774239",
              "value_name": "Inferior"
            }
          ]
        },
        {
          "values": [
            {
              "value_id": "13701105",
              "value_name": "Trasera"
            },
            {
              "value_id": "2262158",
              "value_name": "Izquierda"
            },
            {
              "value_id": "4774239",
              "value_name": "Inferior"
            }
          ]
        }
      ]
    }
  ]
}
```

## Filtrar ítems con posiciones incompletas

 Importante: 

Esta funcionalidad permite identificar productos que requieren completar su información de compatibilidad de posiciones antes de ser publicados o actualizados en el catálogo.

Para identificar y filtrar los ítems que tienen posiciones de compatibilidades incompletas, la API permite realizar búsquedas específicas utilizando el tag **incomplete_position_compatibilities**. Esta funcionalidad es útil para detectar productos que requieren completar su información de compatibilidad antes de ser publicados o actualizados en el catálogo.

Llamada:

```javascript
curl -H 'Authorization: Bearer $ACCESS_TOKEN' \
'https://api.mercadolibre.com/users/1871678972/items/search?tags=incomplete_position_compatibilities'
```

Respuesta:

```javascript
{
  "seller_id": "1871678972",
  "results": [
    "MLB3845403593",
    "MLB3845392401"
  ],
  "paging": {
    "limit": 50,
    "offset": 0,
    "total": 2
  },
  "query": null,
  "orders": [
    {
      "id": "stop_time_asc",
      "name": "Order by stop time ascending"
    }
  ],
  "available_orders": [
    {
      "id": "stop_time_asc",
      "name": "Order by stop time ascending"
    },
    {
      "id": "stop_time_desc",
      "name": "Order by stop time descending"
    },
    {
      "id": "start_time_asc",
      "name": "Order by start time ascending"
    },
    {
      "id": "start_time_desc",
      "name": "Order by start time descending"
    },
    {
      "id": "available_quantity_asc",
      "name": "Order by available quantity ascending"
    },
    {
      "id": "available_quantity_desc",
      "name": "Order by available quantity descending"
    },
    {
      "id": "sold_quantity_asc",
      "name": "Order by sold quantity ascending"
    },
    {
      "id": "sold_quantity_desc",
      "name": "Order by sold quantity descending"
    },
    {
      "id": "price_asc",
      "name": "Order by price ascending"
    },
    {
      "id": "price_desc",
      "name": "Order by price descending"
    },
    {
      "id": "last_updated_desc",
      "name": "Order by lastUpdated descending"
    },
    {
      "id": "last_updated_asc",
      "name": "Order by last updated ascending"
    },
    {
      "id": "total_sold_quantity_asc",
      "name": "Order by total sold quantity ascending"
    },
    {
      "id": {
        "id": "total_sold_quantity_desc",
        "field": "sold_quantity",
        "missing": "_last",
        "order": "desc"
      },
      "name": "Order by total sold quantity descending"
    },
    {
      "id": {
        "id": "inventory_id_asc",
        "field": "inventory_id",
        "missing": "_last",
        "order": "asc"
      },
      "name": "Order by inventory id ascending"
    }
  ]
}
```

 Importante: 

#### Actualización en la funcionalidad de copiar y pegar compatibilidades

A partir del 1 de septiembre, modificaremos el flujo para copiar y pegar compatibilidades. Ahora, al copiar vehículos, siempre se incluirá la información de posiciones. La opción de copiar solo vehículos ha sido eliminada.

Además, mantenemos la posibilidad de elegir si deseas incluir o no las observaciones (notas) de las compatibilidades. Esto te permitirá tener mayor control sobre la información que migras entre tus publicaciones.

Con esto queremos evitar la doble carga de posiciones (en características y compatibilidades) y datos obsoletos (título y/o descripción indican info defasada).

## Identificar ítems con sugerencias de compatibilidades

Con el siguiente recurso podrás conocer el listado de todos los ítems del vendedor que tienen sugerencias de compatibilidades, es decir que tienen el tag “pending_compatibilities”.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/users/$SELLER_ID/items/search?tags=pending_compatibilities
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/users/1695976736/items/search?tags=pending_compatibilities
```

Respuesta:

```javascript
{
   "seller_id": "1373279576",
   "results": [
      "MLB4462690924"
   ],
   "paging": {
       "limit": 50,
       "offset": 0,
       "total": 4
   },
   "query": null,
.
.
.
.
}
```

Con la siguiente llamada y a través del tag pending_compatibilities que se encuentra en el response, puedes identificar para un ítem en específico si tiene sugerencias de compatibilidades.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLB4462690924
```

Respuesta:

```javascript
{
  "id": "MLB4462690924",
  "site_id": "MLB",
  "title": "Pastilha Freio Dianteira Gm Blazer 4.3 V6 Mpfi 1996 À 2003 ",
  "seller_id": 1695976736,
  "category_id": "MLB439261",
  "official_store_id": null,

.
.
.
.
.
  "tags": [
    "pending_compatibilities"
  ],
 "catalog_product_id": "MLB31779615",
  "domain_id": "MLB-VEHICLE_BRAKE_PADS",
  "parent_item_id": null,
  "deal_ids": [
  ],
  "automatic_relist": false,
  "date_created": "2024-02-22T16:51:37.577Z",
  "last_updated": "2024-03-22T20:49:39.772Z",
  "health": 0.75,
  "catalog_listing": false
}
```

Filtrar ítems que tienen compatibilidades sugeridas.

## Obtener la cantidad de sugerencias y reclamos

Con el siguiente recurso podrás conocer la cantidad de sugerencias de compatibilidades y de reclamos que tiene una publicación en específico.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/catalog_domains/$DOMAIN_ID/compatibilities/cards
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
{
"item_id": "MLB4462690924",
"product_id": "MLB31779615",
"filters":   ["REPUTATION",   "SUGGESTED"]
}
https://api.mercadolibre.com/catalog_domains/MLB-CARS_AND_VANS/compatibilities/cards
```

### Atributos request

- **item_id**: atributo obligatorio que corresponde al identificador del ítem.
- **product_id**: atributo que corresponde al producto que está asociado al ítem, es opcional, sin embargo se mejora la performance y los tiempos de respuesta cuando dicho atributo es informado.
- **filters**: atributo obligatorio para indicar si se desean obtener la cantidad de reclamos ["REPUTATION"], la cantidad de sugerencias de compatibilidades ["SUGGESTED"] que tiene el ítem; o ambas cosas ["REPUTATION", "SUGGESTED"].

Respuesta:

```javascript
[
   {
       "title": "Tienes 293 vehículos sugeridos pendientes por agregar",
       "subtitle": "Identificamos vehículos compatibles con tu producto que te ayudarán a vender más",
       "filter": "SUGGESTED",
       "quantity": 293
   },
   {
       "title": "Vehículos con reclamos por incompatibilidad",
       "subtitle": "Revisalos y eliminalos para evitar nuevos reclamos.",
       "filter": "REPUTATION",
       "quantity": 1
   }
]
```

**Reputation**: son todas las compatibilidades que están generando reclamos.

**Suggested**: son todos los casos de compatibilidades nuevas sugeridas a sus publicaciones.

Si quieres conocer cómo obtener las compatibilidades sugeridas para cada ítem puedes ver más detalle en [identificar compatibilidades sugeridas](https://developers.mercadolibre.com.ve/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes#Identificar-compatibilidades-sugeridas).

Si quieres conocer cuáles son las compatibilidades problemáticas que están generando reclamos y afectando la reputación del vendedor seller puedes ver más detalle en [Conocer compatibilidades que generan reclamos](https://developers.mercadolibre.com.ve/es_ar/compatibilidades-entre-items-y-productos#Conocer-compatibilidades-que-generan-reclamos).

## Agregar compatibilidades

 Nota: 

- Disponibilizamos todas las funcionalidades de compatibilidades en las publicaciones de autopartes de las categorías 

MLA1747

, 

MLM1748

, 

MLB22693

, 

MLU1748

 ,

MLC1748

 y 

MCO87919

. 

 -En los sites de MLA, MLM, MLB, MLU, MLC y MCO 

deberán informar compatibilidades de manera obligatoria en los ítems marcados con el tag incomplete_compabilities

 para evitar que las publicaciones de autopartes sean pausadas. 

Para agregar compatibilidades de un ítem con un producto y/o dominio podrás consultar hasta un máximo de 200 productos por llamado (incluidos los definidos en los dominios) y hacerlo de 3 maneras diferentes:

- **Por producto**: para agregar nuevas compatibilidades a un ítem debes enviar las compatibilidades que quieras añadir. No es necesario enviar las existentes para mantener las actuales.
- **Por dominio de producto**: puedes especificar un conjunto de atributos que definen el dominio de productos. Para cada dominio, debes especificar su dominio y para cada atributo, un value conformado por id y/o name.
- **Por producto y dominio**: puedes agergar compatibilidades a un ítem publicado de otro producto y una dominio de productos, es decir, te permite agregar de manera conjunta las 2 primeras.

Nota:

El límite de tráfico por APP_ID es de 100 RPM (request por minuto).

## Valores posibles para una restricción de posición

Al agregar una compatibilidad también podrás indicar la restricción de posición de la misma, con la siguiente llamada podrás conocer los valores posibles para informarla.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/restrictions/values?main_domain_id=MLA-CARS_AND_VANS&secondary_domain_id=MLA-VEHICLE_SHOCK_ABSORBERS
```

Respuesta:

```javascript
{
      "attributes_values": [   
          {
               "attribute_id": "POSITION",
               "values":
                [
                    {
                        "value_id": "23536",
                        "value_name": "Superior",
                    },
                    {
                        "value_id": "23537",
                        "value_name": "Inferior"
                    }
               ]
         }
    ]
}
```

 Importante: 

**El atributo **creation_source** ahora es obligatorio en todas las solicitudes de creación de compatibilidades.**

Este atributo indica el origen de las compatibilidades creadas, con los siguientes valores permitidos:

- **ITEM_SUGGESTIONS:** Compatibilidades generadas por [sugerencias](https://developers.mercadolibre.com.ve/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes#:~:text=%3A%201%0A%7D-,Identificar%20compatibilidades%20sugeridas,-Para%20poder%20mantener).
- **NEW_VEHICLES:** Compatibilidades con [vehículos nuevos.](https://developers.mercadolibre.com.ve/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes#:~:text=%3A%20180%0A%7D-,Identificar%20veh%C3%ADculos%20nuevos,-Para%20poder%20mantener).
- **DEFAULT:** Valor por defecto cuando no se utiliza un filtro específico.

**Casos afectados:**

- Agregar compatibilidad por producto, dominio o ambos.
- Actualizar compatibilidades.

**Casos no afectados:**

- Copiar y pegar compatibilidades.
- Agregar compatibilidad universal.

### Agregar compatibilidad por producto

Para agregar una compatibilidad a uno o varios productos individuales, puedes [utilizar el product search](https://developers.mercadolibre.com.ar/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes#Product-search).

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products": [{
       "id": "$PRODUCT_ID",
       "creation_source": "$CREATION_SOURCE",
       "note": "texto",
       "restrictions": [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      },
                      {
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}, 
                                  {"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      }
                    ]
                }]
   }]
}'
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products": [{
       "id": "MLB155254",
         "creation_source": "ITEM_SUGGESTIONS",
       "note": "Modelos posteriores a Mayo de 2018",
       "restrictions": [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "12456","value_name": "Delantero"}]
                      },
                      {
                        "values":[{"value_id": "65432","value_name": "Trasero"}, 
                                  {"value_id": "87675","value_name": "Inferior"}]
                      }
                    ]
                }]
   }]
}
```

Respuesta:

```javascript
{
 "created_compatibilities_count": 72
}
```

También es posible agregar una nota y restricción de posición a más de una compatibilidad, para esto es necesario reemplazar el nodo **products** por **products_group**, ejemplo:

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products_group": [{
       "ids": ["MLB155254", "MLB155255"],
        "creation_source": "ITEM_SUGGESTIONS"
       "note": "texto",
       "restrictions": [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      },
                      {
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}, 
                                  {"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      }
                    ]
                }]
   }]
}
```

### Agregar compatibilidad por dominio

Para agregar compatibilidades definidas por los atributos que determinan un dominio, [conoce los dominios y atributos de autopartes](https://developers.mercadolibre.com.ar/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes#Atributos-por-dominos).

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products_families": [{
       "domain_id": "$DOMAIN_ID",
    "creation_source": "$CREATION_SOURCE",
       "attributes": [{
               "id": "$ATTRIBUTE_ID",
               "value_id": "$VALUE_ID"
           },
           {
               "id": "$ATTRIBUTE_ID",
               "value_id": "$VALUE_ID"
           },
       ],
     "note": "Texto",
     "restrictions":
                [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      },
                      {
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}, 
                                  {"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      }
                    ]
                }]
}
```

Ejemplo (excepto MLM):

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLA794706391/compatibilities
{
   "products_families": [{
       "domain_id": "MLA-CARS_AND_VANS",
       "creation_source": "ITEM_SUGGESTIONS",
       "attributes": [{
               "id": "BRAND",
               "value_id": "60249"
           },
           {
               "id": "YEAR",
               "value_name": "2010"
           },
       ],
     "note": "Solamente para vehículos de fabricación Europea",
     "restrictions":
                [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "12456","value_name": "Delantero"}]
                      },
                      {
                        "values":[{"value_id": "65432","value_name": "Trasero"}, 
                                  {"value_id": "87675","value_name": "Inferior"}]
                      }
                    ]

                }]
}
```

Respuesta:

```javascript
{
 "created_compatibilities_count": 23
}
```

Ejemplo MLM:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLM794706391/compatibilities
{
   "products_families": [{
           "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
       "creation_source": "ITEM_SUGGESTIONS"
           "attributes": [{
                   "id": "DRIVE_TYPE",
                   "value_id": "8182649"
                  
               },
               {
                   "id": "CAR_AND_VAN_BODY_TYPE",
                   "value_id": "8183109"
                  
               },
               {
                   "id": "YEAR",
                   "value_name": "2010"
                  
               }
           ],
     "note": "Solamente para vehículos de fabricación Europea",
     "restrictions":
                [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "12456","value_name": "Delantero"}]
                      },
                      {
                        "values":[{"value_id": "65432","value_name": "Trasero"}, 
                                  {"value_id": "87675","value_name": "Inferior"}]
                      }
                    ]
       }
   ]
}
```

Respuesta:

```javascript
{
 "created_compatibilities_count": 23
}
```

### Agregar por producto y dominio de productos

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products": [{
       id": "$PRODUCTIID",
        "creation_source": "$CREATION_SOURCE",
       "note": "texto",
       "restrictions": [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      },
                      {
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}, 
                                  {"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      }
                    ]
   }],
   "products_families": [{
       "domain_id": "$DOMAIN_ID",
       "attributes": [{
               "id": "ATTRIBUTE_ID",
               "value_id": "$VALUE_ID"
           },
           {
               "id": "ATTRIBUTE_ID",
               "value_id": "$VALUE_ID"
           },
       ],
     "note": "Texto",
     "restrictions":
                [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      },
                      {
                        "values":[{"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}, 
                                  {"value_id": "$VALUE_ID","value_name": "$VALUE_NAME"}]
                      }
                    ]
                }]
}
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/MLM794706391/compatibilities
{
   "products": [{
       "id": "MLB155254",
     "creation_source": "ITEM_SUGGESTIONS",
       "note": "Modelos posteriores a Mayo de 2018",
       "restrictions": []
   }],
   "products_families": [{
       "domain_id": "MLB-CARS_AND_VANS",
       "attributes": [{
               "id": "BRAND",
               "value_id": "60249"
           },
           {
               "id": "YEAR",
               "value_name": "2010"
           },
       ],
     "note": "Solamente para vehículos de fabricación Europea",
     "restrictions":
                [{
                    "attribute_id": "POSITION",
                    "attribute_values":
                    [{
                        "values":[{"value_id": "12456","value_name": "Delantero"}]
                      },
                      {
                        "values":[{"value_id": "65432","value_name": "Trasero"}, 
                                  {"value_id": "87675","value_name": "Inferior"}]
                      }
                    ]
                }]
}
```

En el campo **note** permitimos hasta 500 caracteres y el texto es moderado.

Respuesta:

```javascript
{
 "created_compatibilities_count": 23
}
```

### Agregar por user products

Para crear las compatibilidades asociadas al UP el campo **domain_id** en el body es requerido y debe enviarse siempre por fuera de las listas de **products_families** y **products**.

Ejemplo:

```javascript
curl --location 'https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities' \
--header 'Content-Type: application/json' \
--data '{ 
    "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
    "category_id": "MLM12344"
    "products": [
        {
            "id": "MLM15870230",
            "note": "nota sobre la compats",
            "restrictions": [
                {
                    "attribute_id": "POSITION",
                    "attribute_values": [
                        {
                            "values": [
                                {
                                    "value_id": "13701104",
                                    "value_name": "Delantero"
                                },
                                {
                                    "value_id": "2262158",
                                    "value_name": "Izquierdo"
                                }
                            ]
                        }
                    ]
                }
            ]
        }
    ],
    "products_families": [
        {
     "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
            "attributes": [
                {
                    "id": "BRAND",
                    "value_id": "60249"
                },
                {
                    "id": "YEAR",
                    "value_name": "2010"
                },
                {
                    "id": "CAR_AND_VAN_MODEL",
                    "value_id": "8240568"
                },
                {
                    "id": "CAR_AND_VAN_SUBMODEL",
                    "value_id": "8238217"
                }
            ]
        }
    ]
}
```

Respuesta:

```javascript
{
 "created_compatibilities_count": 23
}
```

### Posibles errores

**400**: validaciones de consistencia:

- Los campos obligatorios están incompletos.
- El formato de los ids es incorrecto.
- Se encontraron y/o especificaron más de 200 productos para los dominios de productos.
- Se especificaron más de 10 dominios de productos.
- Los productos y/o dominios no pertenecen al mismo site que el ítem.
- Los productos deben ser todos hijos.
- El dominio del ítem es compatible con los dominios de los productos especificados y/o con los dominios especificados en los dominos de productos especificadas.
- No puede haber más de 4 posiciones configuradas en una restricción de posición.
- Cada restricción puede tener máximo 4 ids.
- Los ids deben pertenecer al conjunto cerrado de ids definidos.
- La combinación de valores debe ser única, es decir no puede haber dos listas de ids iguales dentro de una restricción.
- La categoría del ítem debe tener habilitadas las notas y restricciones.
- La nota no puede superar los 500 caracteres.

**403**:token inválido o falta de permisos sobre el ítem.

**404**:o existe el ítem, los productos o los dominios especificados.

## Copiar y pegar compatibilidades

El primer paso para copiar las compatibilidades es obtener una lista de los ítems activos que cuenten con compatibilidades configuradas. Para ello, es necesario comenzar obteniendo todos los ítems activos del vendedor mediante la siguiente consulta:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 'https://api.mercadolibre.com/users/$USER_ID/items/search?status=active&has_compatibilities=true'
```

Todos los items deben conter el atributo "HAS_COMPATIBILITIES".

Para obtener el listado de ítems con la cantidad de compatibilidades, notas y posiciones, realice la siguiente llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/compatibilities_summary
{
   "domain_id": "$domainId",
   "items": [
       "$item_id",
       "item_id",
       "item_id"
   ]
}
'
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN 'https://api.mercadolibre.com/items/compatibilities_summary' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "domain_id": "MLB-CARS_AND_VANS",
   "items": [
       "MLB4816242302",
       "MLB3845318745",
       "MLB3845318797"
   ]
}
```

Respuesta:

```javascript
[
   {
       "item_id": "MLB4816242302",
       "item_title": "Pastilha Dianteira Anuncio De Teste 2",
       "compatibilities_count": 1001,
       "compatibilities_claims_count": 0,
       "notes_count": 65,
       "restrictions_count": 64
   },
   {
       "item_id": "MLB3845318797",
       "item_title": "Pastilha Dianteira Anuncio De Teste",
       "compatibilities_count": 400,
       "compatibilities_claims_count": 0,
       "notes_count": 64,
       "restrictions_count": 64
   },
   {
       "item_id": "MLB3845318745",
       "item_title": "Pastilha Dianteira Anuncio De Teste",
       "compatibilities_count": 200,
       "compatibilities_claims_count": 0,
       "notes_count": 64,
       "restrictions_count": 64
   }
]
```

 Nota: 

 - El límite de la cantidad de ítems a recibir será de 10 ítems.

 - El campo “compatibilities_count” contara todas las compatibilidades que tiene el ítem incluyendo las que tienen claims. 

Para copiar las compatibilidades de un ítem para un ítem sin compatibilidades cargadas, haga esta llamada:

 Nota: 

 En los dos métodos de POST y PUT se agrega un nuevo campo tipo objeto 

“item_to_copy”

, que indicará que va a copiar las compatibilidades del ítem especificado y tomará en cuenta si debe o no copiar las notas y restricciones mediante el campo 

“extended_information”

. 

 Los campos: 

products, products_families, products_group y universal

, no estarán como requeridos para estos casos que se quieran crear o editar ítems mediante el copiar y pegar. 

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN 'https://api.mercadolibre.com/items/$ITEM_ID/compatibilities' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "item_to_copy": {
       "item_id": "$ITEM_ID",
       "extended_information": true
   }
}'
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'https://api.mercadolibre.com/items/MLB3863097751/compatibilities' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
   "item_to_copy": {
       "item_id": "MLB4816242302",
       "extended_information": true
   }
}'
```

Respuesta:

```javascript
200 OK
```

Para copiar las compatibilidades de un ítem para un ítem que ya tiene compatibilidades cargadas, haga esta llamada:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN 'https://api.mercadolibre.com/items/$ITEM_ID/compatibilities' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
 "create": {
   "item_to_copy": {
     "item_id": "$ITEM_ID",
     "extended_information": true
   }
 }
}
```

Ejemplo:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN'https://api.mercadolibre.com/items/MLB3863034063/compatibilities' \
--header 'Content-Type: application/json' \
--header 'Authorization: ••••••' \
--data '{
 "create": {
   "item_to_copy": {
     "item_id": "MLB4816242302",
     "extended_information": true
   }
 }
}
'
```

Respuesta:

```javascript
{
   "create": {
       "products": [
           {
               "id": "MLB7864691",
               "note": "frente",
               "restrictions": [
                   {
                       "attribute_id": "POSITION",
                       "attribute_code": 1,
                       "attribute_values": [
                           {
                               "values": [
                                   {
                                       "value_id": "13701104",
                                       "value_code": 5
                                   }
                               ]
                           }
                       ]
                   }
               ]
           },
                          
.
.
.
       "universal": false,
       "item_to_copy": {
           "item_id": "MLB4816242302",
           "extended_information": true
       }
   }
}
```

### Copiar y pegar compatibilidades por User Products

Para copiar compatibilidades para un user product se debe indicar el origen de las compatibilidades (ya sea desde un ítem o un user product) y, opcionalmente, si se desea copiar también las notas y restricciones, se debe utilizar el campo `"extended_information"`.

Reglas para los campos:

- Si se desea copiar desde un **User Product**, deben completar el campo `"user_product_id"` con el ID correspondiente y enviar el campo `"item_id"` como `null`.
- Si se desea copiar desde un **ítem**, deben completar el campo `"item_id"` y enviar `"user_product_id"` como `null`.

**Importante:** No se deben completar ambos campos al mismo tiempo. Solo uno debe tener valor y el otro debe ser `null`.

Ejemplo copiar compatibilidades de un user product a un user product :

```javascript
curl --location curl -X POST \
'https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities/copy-paste' \
  -H 'content-type: application/json' \
  -d '{
       "domain_id": "MLB-CARS_AND_VANS",
       "category_id": "MLB12344",
       "user_product_id": "MLBU1552549",
       "extended_information": true
}
```

Ejemplo copiar compatibilidades de un item a un user product:

```javascript
curl --location curl -X POST \
'https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities/copy-paste' \
  -H 'content-type: application/json' \
  -d '{
       "domain_id": "MLB-CARS_AND_VANS",
       "category_id": "MLB12344",
       "item_id": "MLB1552432",
       "extended_information": true
}
```

Respuesta:

```javascript
200 OK
```

### Consideraciones

 En la publicación de destino, hay dos escenarios:

 No se copiarán compatibilidades que tengan reclamos, es decir, se excluyen en el proceso de guardar.

Notas y Posiciones: El servicio del POST, PUT para crear compatibilidades recibirá la opción de copiar solo la compatibilidad o la compatibilidad más las notas y restricciones.

 Si el ítem de origen supera las 6K compatibilidades, se generará una excepción de límite al publicar/modificar. En las modificaciones, se debe considerar la intersección de compatibilidades entre los ítems. Por ejemplo, si el ítem de origen tiene 5K y el ítem destino 2K, y 1K son compatibilidades comunes, solo se copiarán 4K del ítem de origen al ítem destino.

En los servicios PUT y POST, se guardarán de forma síncrona hasta 200 compatibilidades, mientras que el resto se creará de forma asíncrona para evitar problemas de rendimiento. Esto puede generar demoras en ver todas las compatibilidades creadas, por lo que se debe notificar al seller sobre esta posibilidad.

## Agregar una compatibilidad universal

 Nota: 

 En este momento, esta funcionalidad no está disponible en producción. 

Con el objetivo de mejorar la calidad de las publicaciones de autopartes de las categorías de accesorios para autos y camionetas ([MLA6520](https://api.mercadolibre.com/categories/MLA6520), [MLM5320](https://api.mercadolibre.com/categories/MLM5320), [MLU1747](https://api.mercadolibre.com/categories/MLU1747), [MLB1747](https://api.mercadolibre.com/categories/MLB1747)) se podrán informar **compatibilidades universales** para indicar que un ítem es compatible con cualquier producto.

Para indicar que un ítem es compatible con cualquier producto, dentro de la petición, se cuenta con el campo **universal**, el cual deberá ser informado en **true**. Lo anterior indica que este ítem es universal (por lo tanto no se le debe de agregar compatibilidades dado que es compatible con todos los productos dentro del mismo dominio).

Al indicar una compatibilidad universal, no es posible especificar **productos ni familias**. En caso de que se envíen ambos campos se obtendrá un error.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
{
   "products": [],
   "products_families": [],
   "products_group": [],
   "universal": true
}
```

Ejemplo:

```javascript
curl -X PUT -H ' https://api.mercadolibre.com/items/MLM12456789/compatibilities
{
   "products": [],
   "products_families": [],
   "products_group": [],
   "universal": true,
}
```

Respuesta:

```javascript
{
 "created_compatibilities_count": 1
}
```

### Posibles errores al asignar compatibilidades universales

**400**: validaciones de consistencia:

En caso de que se envíen productos y familias al mismo tiempo que se informa una compatibilidad universal, se obtendrá:

En el caso que se tenga alguna compatibilidad en el ítem registrada anteriormente y se intenta volver universal, se obtendrá:

- Message: Item has compatibilities and these must be removed before setting it as universal.

Cuando se intenta generar un ítem universal pero la categoria no esta habilitada para esta experiencia, se obtendrá:

- Message: There is no configured compatibility for the category $CATEGORY_ID

Si el ítem anteriormente es universal y se intenta cargar alguna compatibilidad, se obtendrá:

- Message: Item has universal setting and must be removed before creating compatibilities.

**403**: token inválido o falta de permisos sobre el ítem.

**404**: no existe el ítem.

## Actualizar compatibilidades

Con este método es posible crear, actualizar y eliminar las compatibilidades, notas y restricciones de un ítem, utilizando la misma estructura de datos que la creación de compatibilidades. Esta acción puede ser hecha para una o múltiples compatibilidades.

Para estos casos, solo es necesario enviar el **creation_source** dentro del **create**

 Importante: 

- Este PUT es diferente de los otros utilizados para ítems, en este caso, la información no se sobrescribe. Es necesario especificar claramente la acción deseada: crear (create), modificar (update) o eliminar (delete). 

 - Los ejemplos a continuación incluyen "create", "update" y "delete". No es necesario especificar todas las acciones juntas, también se pueden indicar por separado. 

 - Para este método solo se puede indicar 200 productos por llamada. 

Llamada:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN' \
https://api.mercadolibre.com/items/$ITEM_ID/compatibilities
```

 Nota: 

Al borrar compatibilidades, puedes optar por hacerlo a través de un listado de productos (products) o por familias de productos (products_families). Ambas opciones están disponibles en el ejemplo a continuación para que puedas elegir la que mejor se ajuste a tus necesidades y veas cómo enviar los datos correctamente.

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
'{
   "create": {
       "products": [
           {
               "id": "MLB22015088",
    "creation_source": "NEW_VEHICLES"
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60249"
                   },
                   {
                       "id": "MODEL",
                       "value_name": "Gol"
                   },
                   {
                       "id": "YEAR",
                       "value_name": "2023"
                   }
               ]
           }
       ]
   },
   "delete": {
       "products": [
           {
               "id": "MLB22015074"
           },
           {
               "id": "MLB7427549"
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60249"
                   },
                   {
                       "id": "MODEL",
                       "value_id": "62109"
                   },
                   {
                       "id": "YEAR",
                       "value_name": "2023"
                   }
               ]
           }
       ]
   }
}
```

Respuesta:

```javascript
{
   "create": {
       "products": [
           {
               "id": "MLB22015088"
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60249"
                   },
                   {
                       "id": "MODEL",
                       "value_name": "Gol"
                   },
                   {
                       "id": "YEAR",
                       "value_name": "2023"
                   }
               ]
           }
       ],
       "universal": false
   },
   "delete": {
       "products": [
           {
               "id": "MLB22015074"
           },
           {
               "id": "MLB7427549"
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60249"
                   },
                   {
                       "id": "MODEL",
                       "value_id": "62109"
                   },
                   {
                       "id": "YEAR",
                       "value_name": "2023"
                   }
               ]
           }
       ],
       "universal": false
   }
}
```

Ejemplo para crear y actualizar compatibilidades con notas y/o restricciones:

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
'{
   "create": {
       "products": [
           {
               "id": "MLB28049481",
               "note": "nota para teste",
               "restrictions": []
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60395"
                   },
                   {
                       "id": "MODEL",
                       "value_ids": [
                           "389577",
                           "16696"
                       ]
                   }
               ]
           }
       ]
   },
   "update": {
       "products": [
           {
               "id": "MLB22567898",
               "note": "nota para teste 10",
               "restrictions": [
                   {
                       "attribute_id": "POSITION",
                       "attribute_values": [
                           {
                               "values": [
                                   {
                                       "value_id": "13701104",
                                       "value_name": "Dianteira"
                                   }
                               ]
                           }
                       ]
                   }
               ]
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "67781"
                   },
                   {
                       "id": "MODEL",
                       "value_name": "ARGO"
                   }
               ],
               "note": "somente as versões de freios traseiros",
               "restrictions": [
                   {
                       "attribute_id": "POSITION",
                       "attribute_values": [
                           {
                               "values": [
                                   {
                                       "value_id": "13701104",
                                       "value_name": "Dianteira"
                                   }
                               ]
                           }
                       ]
                   }
               ]
           }
       ]
   }
}
```

Respuesta:

```javascript
{
   "create": {
       "products": [
           {
               "id": "MLB28049481",
               "note": "nota para teste",
               "restrictions": []
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "60395"
                   },
                   {
                       "id": "MODEL",
                       "value_ids": [
                           "389577",
                           "16696"
                       ]
                   }
               ]
           }
       ],
       "universal": false
   },
   "update": {
       "products": [
           {
               "id": "MLB22567898",
               "note": "nota para teste 10",
               "restrictions": [
                   {
                       "attribute_id": "POSITION",
                       "attribute_code": 1,
                       "attribute_values": [
                           {
                               "values": [
                                   {
                                       "value_id": "13701104",
                                       "value_name": "Dianteira",
                                       "value_code": 5
                                   }
                               ]
                           }
                       ]
                   }
               ]
           }
       ],
       "products_families": [
           {
               "domain_id": "MLB-CARS_AND_VANS",
               "attributes": [
                   {
                       "id": "BRAND",
                       "value_id": "67781"
                   },
                   {
                       "id": "MODEL",
                       "value_name": "ARGO"
                   }
               ],
               "note": "somente as versões de freios traseiros",
               "restrictions": [
                   {
                       "attribute_id": "POSITION",
                       "attribute_code": 1,
                       "attribute_values": [
                           {
                               "values": [
                                   {
                                       "value_id": "13701104",
                                       "value_name": "Dianteira",
                                       "value_code": 5
                                   }
                               ]
                           }
                       ]
                   }
               ]
           }
       ],
       "universal": false
   }
}
```

 Nota: 

En el caso de que se quiera borrar las notas y/o restricciones se debe enviar en la sección 'update' con note: "" o restrictions: []

### Actualizar compatibilidades por user products

Este servicio permite crear o eliminar las compatibilidades de un **user product** y también actualizar las notas y/o restricciones de las mismas.

Al igual que en la creación, el campo `domain_id` pasa a ser requerido, y debe ser enviado siempre por fuera de las listas de `products_families` y `products`.

```javascript
curl --location --request PUT 'https://api.mercadolibre.com/user-products/MLMU427597763/compatibilities' \
--header 'Content-Type: application/json' \
--data '{
    "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
    "category_id": "MLM12344",
    "create": {
        "products": [
            {
                "id": "MLM15861466",
                "note": "Nota producto creado individual desde upd compatibilidad 21/08",
                "restrictions": [
                    {
                        "attribute_id": "POSITION",
                        "attribute_values": [
                            {
                                "values": [
                                    {
                                        "value_id": "13701104",
                                        "value_name": "Dianteiro"
                                    }
                                ]
                            },
                            {
                                "values": [
                                    {
                                        "value_id": "13373176",
                                        "value_name": "Passageiro"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "update": {
        "products": [
            {
                "id": "MLM15871318",
                "note": "Nota individual compatibilidad actualizada 21/08 por product",
                "restrictions": [
                    {
                        "attribute_id": "POSITION",
                        "attribute_values": [
                            {
                                "values": [
                                    {
                                        "value_id": "13701104",
                                        "value_name": "Dianteiro"
                                    }
                                ]
                            },
                            {
                                "values": [
                                    {
                                        "value_id": "13373176",
                                        "value_name": "Passageiro"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ],
        "products_families": [
            {
                "domain_id": "MLM-CARS_AND_VANS_FOR_COMPATIBILITIES",
                "attributes": [
                    {
                        "id": "BRAND",
                        "value_name": "CHEVROLET"
                    },
                    {
                        "id": "YEAR",
                        "value_name": "2017"
                    }
                ],
                "note": "Nota familiar actualizada 21/08 por family upd",
                "restrictions": [
                    {
                        "attribute_id": "POSITION",
                        "attribute_values": [
                            {
                                "values": [
                                    {
                                        "value_id": "13701104",
                                        "value_name": "Dianteiro"
                                    }
                                ]
                            },
                            {
                                "values": [
                                    {
                                        "value_id": "13373176",
                                        "value_name": "Passageiro"
                                    },
                                    {
                                        "value_id": "2262158",
                                        "value_name": "Esquerdo"
                                    }
                                ]
                            }
                        ]
                    }
                ]
            }
        ]
    },
    "delete": {
        "products": [
            {
                "id": "MLM15847272"
            }
        ]
    }
}
```

### Posibles errores

**400** - Validaciones de consistencia:

- Completitud de los campos obligatorios.
- Correctitud en el formato de los ids.
- Productos y/o dominios pertenecen al mismo site que el item.
- Dominio del item es compatible con los dominios de los productos especificados.
- Se ha excedido el máximo de 200 productos para una sola solicitud.
