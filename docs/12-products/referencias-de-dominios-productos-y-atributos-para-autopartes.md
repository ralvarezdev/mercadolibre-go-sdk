---
title: Referencias de dominios, productos y atributos para Autopartes
slug: referencias-de-dominios-productos-y-atributos-para-autopartes
section: 12-products
source: https://developers.mercadolibre.com.ve/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes
captured: 2026-06-06
---

# Referencias de dominios, productos y atributos para Autopartes

> Source: <https://developers.mercadolibre.com.ve/es_ar/referencias-de-dominios-productos-y-atributos-para-autopartes>

## Endpoints referenced

- `https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks`
- `https://api.mercadolibre.com/catalog_domains/$DOMAIN_ID`
- `https://api.mercadolibre.com/catalog_domains/$DOMAIN_ID/attributes/$ATTRIBUTE_ID/top_values`
- `https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS`
- `https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/BRAND/top_values`
- `https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/MODEL/top_values`
- `https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/VEHICLE_YEAR/top_values`
- `https://api.mercadolibre.com/categories/$CATEGORY_ID/attributes`
- `https://api.mercadolibre.com/categories/MLA12345/attributes`

## Content

Última actualización 29/08/2024

## Referencias de dominios, productos y atributos para Autopartes

**Dominios disponibles**

| País | Dominio |
| --- | --- |
| ARGENTINA | MLA-CARS_AND_VANS |
| BRASIL | MLB-CARS_AND_VANS |
| MÉXICO | MLM-CARS_AND_VANS_FOR_COMPATIBILITIES |
| URUGUAY | MLU-CARS_AND_VANS |
| CHILE | MLC-CARS_AND_VANS_FOR_COMPATIBILITIES |
| COLOMBIA | MCO-CARS_AND_VANS_FOR_COMPATIBILITIES |

De acuerdo al dominio del site, sugerimos que dentro de tu gestor de compatibilidades habilites filtros primarios, secundarios y opcionales según se muestra a continuación.

Ejemplo de atributos utilizados en el motor de búsqueda de compatibilidades para vehículos del dominio CARS_AND_VANS que aplica en los sites de MLA, MLB y MLU.

Ejemplo de atributos utilizados en motor de búsqueda de compatibilidades para vehículos del dominio CARS_AND_VANS_FOR_COMPATIBILITIES que aplica en los sites de MLM, MLC y MCO.

### Atributos principales

| Descripción de atributos | Atributos de CARS_AND_VANS (MLA, MLB, MLU) | Atributos de CARS_AND_VANS_FOR_COMPATIBILITIES (MLM, MLC y MCO) |
| --- | --- | --- |
| MARCA | BRAND | BRAND |
| MODELO | MODEL | CAR_AND_VAN_MODEL |
| AÑO | VEHICLE_YEAR | YEAR |
| VERSIÓN | SHORT_VERSION | CAR_AND_VAN_SUBMODEL |
| MOTOR | ENGINE | CAR_AND_VAN_ENGINE |

### Atributos secundarios

| Descripción de atributos | Atributos de CARS_AND_VANS (MLA, MLB y MLU) | Atributos de CARS_AND_VANS_FOR_COMPATIBILITIES (MLM, MLC y MCO) |
| --- | --- | --- |
| COMBUSTIBLE | FUEL_TYPE | N/A |
| POTENCIA | POWER | N/A |
| CARROCERÍA | VEHICLE_BODY_TYPE | N/A |
| TRANSMISIÓN | TRANSMISSION_CONTROL_TYPE | N/A |

### Atributos opcionales

| Descripción de atributos | Atributos de CARS_AND_VANS (MLA, MLB y MLU) | Atributos de CARS_AND_VANS_FOR_COMPATIBILITIES (MLM, MLC y MCO) |
| --- | --- | --- |
| MARCHAS | GEAR_NUMBER | CAR_AND_VAN_ENGINE |
| PUERTAS | DOORS | N/A |
| DIRECCIÓN | STEERING | N/A |
| TRACCIÓN | TRACTION_CONTROL | N/A |
| VÁLVULAS | VALVES_PER_CYLINDER | N/A |
| SISTEMA DE DIRECCIÓN | N/A | STEERING_SYSTEM |
| TIPO DE DIRECCIÓN | N/A | STEERING_TYPE |
| TIPO DE TRACCIÓN | N/A | DRIVE_TYPE |
| CARROCERÍA | N/A | CAR_AND_VAN_BODY_TYPE |
| TIPO DE CONTROL DE LA TRANSMISIÓN | N/A | TRANSMISSION_CONTROL_TYPE |
| CANTIDAD DE VELOCIDADES DE LA TRANSMISIÓN | N/A | TRANSMISSION_SPEEDS_NUMBER |
| CANTIDAD DE PUERTAS | N/A | BODY_DOORS_NUMBER |
| FRENOS ABS | N/A | BRAKE_ABS |

## Atributos por dominio

Recuerda que el detalle de los atributos por cada dominio lo puedes obtener con la siguiente llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/$DOMAIN_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS
```

## Atributos por categoría

Recuerda que el detalle de los atributos por cada categoría lo puedes obtener con la siguiente llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/categories/$CATEGORY_ID/attributes
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/categories/MLA12345/attributes
```

## Búsqueda de vehículos

 Nota: 

A partir de abril 2024 disponibilizamos el recurso POST a /catalog_compatibilities/products_search/chunks con el cual 

podrás igualar la experiencia de búsqueda de vehículos a la de nuestro Gestor de Compatibilidades (ABM) de Mercado Libre

.

Con el recurso POST a /catalog_compatibilities/products_search/chunks y por medio de los atributos del dominio es posible realizar búsquedas que te permitan:

1. Obtener los vehículos disponibles en nuestro catálogo.
2. Identificar los vehículos nuevos que fueron añadidos a nuestro catálogo.
3. Identificar sugerencias de compatibilidades.

**Parámetros******

**Limit**: parámetro opcional, que indica la cantidad de productos a devolver. El valor por máximo es = 50.

**Offset**: parámetro opcional que indica el registro a partir del cuál se devolverán resultados. El valor por defecto es = 0.

### Campos de respuesta

**Domain_id**: obligatorio enviar para identificar las compatibilidades.

**Site_id**: obligatorio poner el site que estás buscando.

**Filter**: atributo opcional para indicar qué productos se desean obtener.

- Al ingresar el valor “NEW” se devolverán los vehículos nuevos recientemente añadidos al catálogo.
- Al indicar “SUGGESTED” devolverá el listado de vehículos sugeridos para el ítem o producto indicado.
- En caso de que no se envíe ningún valor (vacío o null), se devolverán todos los vehículos (incluyendo nuevos y sugeridos) tomando en cuenta los atributos de búsqueda.

**Item_id**: atributo obligatorio en caso de que se deseen obtener las sugerencias del ítem. En casos de que no envíe el ítem_id, se lista todas las informaciones del secondary_product_id enviado.

**Secondary_product_id**: ID del producto asociado al ítem, es opcional, pero mejora la performance y los tiempos de respuesta cuando usado. (En casos de que no se envíe, se tomará del item).

**Known_attributes**: siempre deben agregarlos con sus "value_ids" y podrás poner en formato de lista todas las opciones que quieras.

**Sort**: es un campo opcional que permite ordenar resultados por el attribute_id indicado. Por ahora sólo es posible ordenar por BRAND ya sea de manera ascendente o descendente.

Tabla de obligatoriedad de atributos según la necesidad de la búsqueda realizada.

| Atributo | Obtener vehículos disponibles en nuestro catálogo | Identificar vehículos nuevos | Identificar sugerencias |
| --- | --- | --- | --- |
| **domain_id** | Obligatorio | Obligatorio | Obligatorio |
| **site_id** | Obligatorio | Obligatorio | Obligatorio |
| **item_id** | Opcional | Opcional | Opcional |
| **secondary_product_id** | Opcional | Opcional | Opcional, pero en caso de contar con dicho valor se sugiere informarlo, ya que ayuda a mejorar el performance de la petición. |
| **Filter** | No requiere ser informado. | “NEW” | “SUGGESTED” |
| **Sort** | Opcional | Opcional | Opcional |

### Campos de respuesta

**id**: identificador del producto (vehículo).

**attributes**: arreglo de atributos del vehículo.

- id: nombre del atributo del vehículo.
- value_id: identificador valor asociado al atributo del vehículo.
- value_name: valor asociado al atributo del vehículo.

**filters**: este atributo **es una lista** que indica a qué filtro pertenece el producto, es decir:****

- en caso de no enviar en el body ningún filtro, se indica si el producto es nuevo o sugerido o ambos , ejemplo: ( "filters": [“NEW", “SUGGESTED"] ); en caso de que el producto no pertenezca a ninguna de esas opciones se devolverá una lista vacía ("filters": [] ).
- en caso de enviar un filtro en el body, se devuelve en esa lista el filtro que se seleccionó en el request, ejemplo "filters": [“NEW"] o "filters": [“SUGGESTED"].
- value_name: valor asociado al atributo del vehículo.

**total**: total de resultados que coinciden con la búsqueda.

## Obtener vehículos disponibles en nuestro catálogo

Por medio de los atributos del dominio, al indicarlos en “known_attributes”, puedes realizar la búsqueda para obtener los vehículos que se encuentran disponibles en nuestro catálogo.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
{
 "domain_id": "MLB-CARS_AND_VANS",
 "site_id": "MLB",
 "known_attributes": [
     {
         "id": "BRAND",
         "value_ids": ["60249", "66432", "9909"]
     },
     {
         "id": "MODEL",
         "value_ids": ["389648", "17780469"]
     }
 ],
 "sort": {
     "attribute_id": "BRAND",
     "order": "desc"
 }
}
https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Respuesta:

```javascript
{
   "results": [
       {
           "id": "MLB22015088",
           "attributes": [
               {
                   "id": "BRAND",
                   "value_id": "60249",
                   "value_name": "Volkswagen"
               },
               {
                   "id": "MODEL",
                   "value_id": "389648",
                   "value_name": "Voyage"
               },
               {
                   "id": "VEHICLE_YEAR",
                   "value_id": "12023859",
                   "value_name": "2023"
               },
        .
.
.
{
                   "id": "CURRENCY",
                   "value_id": "10837729",
                   "value_name": "r$"
               }
           ],
           "filters": []
       },
       {
           "id": "MLB18230485",
           "attributes": [...],
           "filters": ["NEW"
]
       }
   ],
   "total": 180
}
```

## Identificar vehículos nuevos

Para poder mantener actualizadas siempre las compatibilidades de tus publicaciones, con el siguiente recurso, al indicar el atributo filter = “NEW” en la llamada, podrás conocer cuáles son los vehículos nuevos que fueron agregados a nuestro catálogo a partir del 1er día del mes anterior.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
{
 "domain_id": "MLB-CARS_AND_VANS",
 "site_id": "MLB",
 "known_attributes": [
     {
         "id": "BRAND",
         "value_ids": ["60249", "66432", "9909"]
     },
     {
         "id": "MODEL",
         "value_ids": ["389648", "17780469"]
     }
 ],
 "sort": {
     "attribute_id": "BRAND",
     "order": "desc"
 },
  "filter": "NEW"
}
https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Respuesta:

```javascript
{
   "results": [
       {
           "id": "MLB34236071",
           "attributes": [
               {
                   "id": "BRAND",
                   "name": "Marca",
                   "value_id": "66432",
                   "value_name": "Ford",
                   "values": [
                       {
                           "id": "66432",
                           "name": "Ford"
                       }
                   ]
               },
               {
                   "id": "MODEL",
                   "name": "Modelo",
                   "value_id": "17780469",
                   "value_name": "Corcel Ii",
                   "values": [
                       {
                           "id": "17780469",
                           "name": "Corcel Ii"
                       }
                   ]
               },
               {
                   "id": "CURRENCY",
                   "name": "Moeda",
                   "value_id": "10837729",
                   "value_name": "r$",
                   "values": [
                       {
                           "id": "10837729",
                           "name": "r$"
                       }
                   ]
               }
           ],
           "filters": [
               "NEW"
           ]
       }
   ],
   "total": 1
}
```

## Identificar compatibilidades sugeridas

Para poder mantener actualizadas las compatibilidades de los ítems y/o en casos de que identifique que un ítem tiene el tag “pending_compatibilities” con el siguiente recurso (utilizando el atributo filter = “SUGGESTED”) podrás conocer cuáles son los vehículos **sugeridos** para los ítems.

Recuerda que dentro de la llamada no es obligatorio informar el secondary_product_id, sin embargo, en caso de hacerlo esto te ayudará a tener un mejor performance durante la petición.

 Nota: 

El secondary_product_id corresponde al atributo catalog_product_id que se obtiene al realizar el GET a /items/$ITEM_ID. 

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H 'Content-Type: application/json' -d 
{
 "domain_id": "MLB-CARS_AND_VANS",
 "site_id": "MLB",
 "item_id": "MLB4462690924",
 "secondary_product_id": "MLB31779615",
"known_attributes": [
     {
         "id": "BRAND",
         "value_ids": ["60297"]
     },
     {
         "id": "MODEL",
         "value_ids": ["389399"]
     }
 ],
 "sort": {
     "attribute_id": "BRAND",
     "order": "desc"
 },
 "filter": "SUGGESTED"
}

https://api.mercadolibre.com/catalog_compatibilities/products_search/chunks
```

Respuesta:

```javascript
{
   "results": [
       {
           "id": "MLB7866013",
           "attributes": [
               {
                   "id": "BRAND",
                   "name": "Marca",
                   "value_id": "60297",
                   "value_name": "Toyota",
                   "values": [
                       {
                           "id": "60297",
                           "name": "Toyota"
                       }
                   ]
               },
               {
                   "id": "MODEL",
                   "name": "Modelo",
                   "value_id": "389399",
                   "value_name": "Bandeirante",
                   "values": [
                       {
                           "id": "389399",
                           "name": "Bandeirante"
                       }
                   ]
               },
.
.
.
               {
                   "id": "CURRENCY",
                   "name": "Moeda",
                   "value_id": "10837729",
                   "value_name": "r$",
                   "values": [
                       {
                           "id": "10837729",
                           "name": "r$"
                       }
                   ]
               }
           ],
           "filters": [
               "SUGGESTED"
           ]
       }
   ],
   "total": 293
}
```

Si quieres conocer cómo identificar ítems que tienen sugerencias de compatibilidades puedes ver más detalle en identificar [identificar ítems con sugerencias de compatibilidades](https://developers.mercadolibre.com.ve/es_ar/compatibilidades-entre-items-y-productos?nocache=true#Identificar-ítems-con-sugerencias-de-compatibilidades).

### Posibles errores:

| Error_code | Mensaje del error | Descripción |
| --- | --- | --- |
| 400 | There is no configured compatibility for the category $categoryId | La categoría consultada no está habilitada para informar compatibilidades. |
| 401 | Invalid access token. | Access Token inválido. |
| 403 | Domain is not active. | Dominio inactivo en buybox. |

**400**: formato incorrecto / más de 200 productos para el dominio especificado / más de 10 dominios especificados.

**403**: token inválido o falta de permisos sobre el ítem.

**404**: el ítem o la compatibilidad no existen.

## Top values

A continuación puedes ver cómo implementar por medio del recurso Top values la funcionalidad para cargar distintas listas con valores de atributos e ir filtrando los resultados.
 Con el siguiente recurso se pueden obtener los valores de cada combinación e ir refinando cada vez la búsqueda.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/$DOMAIN_ID/attributes/$ATTRIBUTE_ID/top_values 
```

Ejemplo "BRAND":

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/BRAND/top_values
```

Respuesta:

```javascript
[
   {
       "id": "60249",
       "name": "Volkswagen",
       "metric": 7781
   },
   {
       "id": "66432",
       "name": "Ford",
       "metric": 5616
   },
   {
       "id": "9909",
       "name": "Renault",
       "metric": 4327
   },
   {
       "id": "60279",
       "name": "Peugeot",
       "metric": 4250
   },
   {
       "id": "67781",
       "name": "Fiat",
       "metric": 4172
   },
[…]
]
```

Ejemplo para filtrar modelos (MODEL) de una marca (BRAND):

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/MODEL/top_values 
{
   "known_attributes": [
       {
           "id": "BRAND",
           "value_id": "60249"
       }
   ]
}
```

Respuesta:

```javascript
[
   {
       "id": "63686",
       "name": "Amarok",
       "metric": 1516
   },
   {
       "id": "1252874",
       "name": "Gol Trend",
       "metric": 925
   },
   {
       "id": "62109",
       "name": "Gol",
       "metric": 684
   },
   {
       "id": "1252871",
       "name": "Suran",
       "metric": 604
   },
   {
       "id": "64016",
       "name": "Vento",
       "metric": 585
   },
…
]
```

Ejemplo para obtener los años disponibles (VEHICLE_YEAR) filtrando por marca y modelo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/catalog_domains/MLA-CARS_AND_VANS/attributes/VEHICLE_YEAR/top_values
{
   "known_attributes": [
       {
           "id": "BRAND",
           "value_id": "60249"
       },
       {
           "id": "MODEL",
           "value_id": "63686"
       }
   ]
}
```

Respuesta:

```javascript
[
   {
       "id": "6730991",
       "name": "2020",
       "metric": 732
   },
   {
       "id": "423549",
       "name": "2015",
       "metric": 130
   },
   {
       "id": "436694",
       "name": "2017",
       "metric": 115
   },
   {
       "id": "2451646",
       "name": "2019",
       "metric": 104
   },
[…]
]
```

**Volver**: [Compatibilidades entre ítems y productos de Autopartes](https://developers.mercadolibre.com.ar/es_ar/compatibilidades-entre-items-y-productos).
