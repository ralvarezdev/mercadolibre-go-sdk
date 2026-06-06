---
title: Publica Inmuebles
slug: publica-inmueble
section: 21-real-estate
source: https://developers.mercadolibre.com.ve/es_ar/publica-inmueble
captured: 2026-06-06
---

# Publica Inmuebles

> Source: <https://developers.mercadolibre.com.ve/es_ar/publica-inmueble>

## Endpoints referenced

- `https://api.mercadolibre.com/items`

## Content

Última actualización 04/03/2026

## Publica Inmuebles

 Esta sección detalla cómo utilizar la API de MercadoLibre para publicar anuncios de inmuebles. Al integrar nuestra API, podrás automatizar la publicación de inmuebles, simplificando la gestión de grandes catálogos y la integración con sistemas de gestión de propiedades.

Importante:

Antes de publicar revisa la sección de 

Categorías y Atributos

 para que conozcas los atributos mínimos necesarios para publicar.

# Preparación para la Publicación

Recuerda que necesitarás un **access_token** para realizar la publicación.

- **Obtención del `access_token`:** Si tienes dudas sobre cómo obtener tu `access_token`, consulta la guía de [Autenticación y Autorización](https://developers.mercadolibre.com.ar/es_ar/autenticacion-y-autorizacion).
- **Validación del JSON:** Te recomendamos validar el JSON que envías antes de realizar la solicitud `POST`. Para ello, consulta el [tutorial de validación de artículos](https://developers.mercadolibre.cl/es_ar/validador-de-publicaciones).
- **Creación del JSON:** Puedes crear un JSON para tu inmueble basándote en el ejemplo proporcionado en **Publica tu inmueble**. Alternativamente, puedes enviar el JSON directamente para publicar un producto de prueba en el sitio.

Estás listo, continúa a la próxima etapa>

# Publica tu inmueble

Importante:

A partir del **23 de febrero de 2026**, el envío de **al menos una imagen** (array pictures) pasa a ser obligatorio en la creación de anuncios vía POST /items para **Listing Type Silver** y, a partir del **27 de febrero de 2026**, también para los anuncios gold, gold_special y gold_premium con requires_picture: true. Las solicitudes **sin imágenes serán rechazadas** con HTTP 400 (error 173 – LTP_PICTURE_REQUIRED) y ya no será posible crear el aviso primero para incluir imágenes después.

Asegurate de incluir **al menos una imagen** en el array pictures para estos tipos de anuncios.

Para publicar tu inmueble, debes enviar una solicitud **POST** al endpoint */items*. Asegúrate de incluir toda la información obligatoria, la cual puedes consultar en las guías de [Atributos](https://developers.mercadolibre.com.ar/es_ar/atributos-inmuebles), [Localizar inmuebles](https://developers.mercadolibre.com.ar/es_ar/localizar-inmuebles) y [Calidad de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/calidad-de-las-publicaciones-inmuebles). Esto te ayudará a obtener la información necesaria y a realizar la publicación de la mejor manera.

Ejemplo de llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' "Content-Type: application/json" -d '{
  "title": "Property title",
  "category_id": "MLA401686",
  "price": 100000,
  "currency_id": "ARS",
  "available_quantity": 1,
  "buying_mode": "classified",
  "listing_type_id": "silver",
  "condition": "not_specified",
  "channels": [
    "marketplace"
  ],
  "pictures": [
    {
      "source": "http://mla-d2-p.mlstatic.com/item-de-test-no-ofertar-543605-MLA25041518406_092016-O.jpg?square=false"
    }
  ],
  "seller_contact": {
    "contact": "Contact name",
    "other_info": "Additional contact info",
    "area_code": "011",
    "phone": "4444-5555",
    "area_code2": "",
    "phone2": "",
    "email": "contact-email@somedomain.com",
    "webmail": ""
  },
  "location": {
    "address_line": "My property address 1234",
    "zip_code": "01234567",
    "neighborhood": {
      "id": "TUxBQlBBUzgyNjBa"
    },
    "latitude": -34.48755,
    "longitude": -58.56987
  },
  "attributes": [
    {
      "id": "ROOMS",
      "value_name": "2"
    },
    {
      "id": "FULL_BATHROOMS",
      "value_name": "1"
    },
    {
      "id": "PARKING_LOTS",
      "value_name": "1"
    },
    {
      "id": "BEDROOMS",
      "value_name": "4"
    },
    {
      "id": "COVERED_AREA",
      "value_name": "30 m²"
    },
    {
      "id": "TOTAL_AREA",
      "value_name": "40 m²"
    }
  ],
  "video_id": "gqkEN9poKM;matterport",
  "description": {
    "plain_text": "This is the real estate property description.\n"
  }
}'
https://api.mercadolibre.com/items
```

Ejemplo de respuesta:

```javascript
{
    "id": "MLA839018613",
    "site_id": "MLA",
    "title": "Property Title",
    "subtitle": null,
    "seller_id": 526655030,
    "category_id": "MLA401686",
    "official_store_id": null,
    "price": 100000,
    "base_price": 100000,
    "original_price": null,
    "inventory_id": null,
    "currency_id": "ARS",
    "initial_quantity": 1,
    "available_quantity": 1,
    "sold_quantity": 0,
    "sale_terms": [],
    "buying_mode": "classified",
    "listing_type_id": "silver",
    "start_time": "2020-02-13T19:29:01.499Z",
    "stop_time": "2020-05-14T04:00:00.000Z",
    "end_time": "2020-05-14T04:00:00.000Z",
    "expiration_time": null,
    "condition": "not_specified",
    "permalink": "http://departamento.mercadolibre.com.ar/MLA-839018613-property-title-_JM",
    "pictures": [
        {
            "id": "910707-MLA40763776324_022020",
            "url": "http://http2.mlstatic.com/resources/frontend/statics/processing-image/1.0.0/O-ES.jpg",
            "secure_url": "https://http2.mlstatic.com/resources/frontend/statics/processing-image/1.0.0/O-ES.jpg",
            "size": "500x500",
            "max_size": "500x500",
            "quality": ""
        }
    ],
    "video_id": null,
    "descriptions": [
        {
            "id": "MLA839018613-2516619253"
        }
    ],
    "accepts_mercadopago": false,
    "non_mercado_pago_payment_methods": [],
    "shipping": {
        "mode": "not_specified",
        "local_pick_up": false,
        "free_shipping": false,
        "methods": [],
        "dimensions": null,
        "tags": [],
        "logistic_type": null,
        "store_pick_up": false
    },
    "international_delivery_mode": "none",
    "seller_address": {
        "id": 1087750186,
        "comment": "Referencia: The Testing Cavern",
        "address_line": "Testing Street 1450",
        "zip_code": "1430",
        "city": {
            "id": "TUxBQlNBQTM3Mzda",
            "name": "Saavedra"
        },
        "state": {
            "id": "AR-C",
            "name": "Capital Federal"
        },
        "country": {
            "id": "AR",
            "name": "Argentina"
        },
        "latitude": -34.5545188,
        "longitude": -58.4915986,
        "search_location": {
            "neighborhood": {
                "id": "TUxBQlNBQTM3Mzda",
                "name": "Saavedra"
            },
            "city": {
                "id": "TUxBQ0NBUGZlZG1sYQ",
                "name": "Capital Federal"
            },
            "state": {
                "id": "TUxBUENBUGw3M2E1",
                "name": "Capital Federal"
            }
        }
    },
    "seller_contact": {
        "contact": "Contact name",
        "other_info": "Additional contact info",
        "area_code": "011",
        "phone": "4444-5555",
        "area_code2": "",
        "phone2": "",
        "email": "contact-email@somedomain.com",
        "webpage": "",
        "country_code": "",
        "country_code2": ""
    },
    "location": {
        "address_line": "My property address 1234",
        "zip_code": "01234567",
        "neighborhood": {
            "id": "TUxBQlBBUzgyNjBa",
            "name": "Paso del Rey"
        },
        "city": {
            "id": "TUxBQ01PUmViMTE3",
            "name": "Moreno"
        },
        "state": {
            "id": "TUxBUEdSQWVmNTVm",
            "name": "Bs.as. G.b.a. Oeste"
        },
        "country": {
            "id": "AR",
            "name": "Argentina"
        },
        "latitude": -34.48755,
        "longitude": -58.56987,
        "open_hours": ""
    },
    "geolocation": {
        "latitude": -34.48755,
        "longitude": -58.56987
    },
    "coverage_areas": [],
    "attributes": [
        {
            "id": "ROOMS",
            "name": "Ambientes",
            "value_id": null,
            "value_name": "2",
            "value_struct": null,
            "values": [
                {
                    "id": null,
                    "name": "2",
                    "struct": null
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "FULL_BATHROOMS",
            "name": "Baños",
            "value_id": null,
            "value_name": "1",
            "value_struct": null,
            "values": [
                {
                    "id": null,
                    "name": "1",
                    "struct": null
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "PARKING_LOTS",
            "name": "Cocheras",
            "value_id": null,
            "value_name": "1",
            "value_struct": null,
            "values": [
                {
                    "id": null,
                    "name": "1",
                    "struct": null
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "BEDROOMS",
            "name": "Dormitorios",
            "value_id": null,
            "value_name": "4",
            "value_struct": null,
            "values": [
                {
                    "id": null,
                    "name": "4",
                    "struct": null
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "COVERED_AREA",
            "name": "Superficie cubierta",
            "value_id": null,
            "value_name": "30 m²",
            "value_struct": {
                "number": 30,
                "unit": "m²"
            },
            "values": [
                {
                    "id": null,
                    "name": "30 m²",
                    "struct": {
                        "number": 30,
                        "unit": "m²"
                    }
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "TOTAL_AREA",
            "name": "Superficie total",
            "value_id": null,
            "value_name": "40 m²",
            "value_struct": {
                "number": 40,
                "unit": "m²"
            },
            "values": [
                {
                    "id": null,
                    "name": "40 m²",
                    "struct": {
                        "number": 40,
                        "unit": "m²"
                    }
                }
            ],
            "attribute_group_id": "FIND",
            "attribute_group_name": "Ficha técnica"
        },
        {
            "id": "PROPERTY_TYPE",
            "name": "Inmueble",
            "value_id": "242062",
            "value_name": "Departamento",
            "value_struct": null,
            "values": [
                {
                    "id": "242062",
                    "name": "Departamento",
                    "struct": null
                }
            ],
            "attribute_group_id": "MAIN",
            "attribute_group_name": "Principales"
        },
        {
            "id": "OPERATION",
            "name": "Operación",
            "value_id": "242075",
            "value_name": "Venta",
            "value_struct": null,
            "values": [
                {
                    "id": "242075",
                    "name": "Venta",
                    "struct": null
                }
            ],
            "attribute_group_id": "MAIN",
            "attribute_group_name": "Principales"
        },
        {
            "id": "OPERATION_SUBTYPE",
            "name": "Subtipo de operación",
            "value_id": "244562",
            "value_name": "Propiedad individual",
            "value_struct": null,
            "values": [
                {
                    "id": "244562",
                    "name": "Propiedad individual",
                    "struct": null
                }
            ],
            "attribute_group_id": "MAIN",
            "attribute_group_name": "Principales"
        }
    ],
    "warnings": [],
    "listing_source": "",
    "variations": [],
    "thumbnail": "http://http2.mlstatic.com/resources/frontend/statics/processing-image/1.0.0/I-ES.jpg",
    "secure_thumbnail": "https://http2.mlstatic.com/resources/frontend/statics/processing-image/1.0.0/I-ES.jpg",
    "status": "active",
    "sub_status": [],
    "tags": [
        "test_item"
    ],
    "warranty": null,
    "catalog_product_id": null,
    "domain_id": null,
    "seller_custom_field": null,
    "parent_item_id": null,
    "differential_pricing": null,
    "deal_ids": [],
    "automatic_relist": false,
    "date_created": "2020-02-13T19:29:01.823Z",
    "last_updated": "2020-02-13T19:29:01.823Z",
    "health": null,
    "catalog_listing": false,
    "item_relations": []
}
```

### Parámetros de respuesta

| Parámetro | Tipo | Descripción |
| --- | --- | --- |
| id | String | Identificador único del ítem publicado (e.g., "MLA839018613"). |
| site_id | String | Identificador del sitio donde se publica el ítem (e.g., "MLA"). |
| title | String | Título de la publicación. |
| subtitle | String | Subtítulo del ítem (puede ser null). |
| seller_id | Integer | Identificador del vendedor. |
| category_id | String | Identificador de la categoría seleccionada (e.g., "MLA401686"). |
| official_store_id | Integer | Identificador de la tienda oficial. |
| price | Integer | Precio del ítem publicado. |
| currency_id | String | Identificador de la moneda (e.g., "ARS"). |
| available_quantity | Integer | Cantidad disponible del ítem (para inmuebles siempre será 1). |
| sold_quantity | Integer | Cantidad vendida. |
| buying_mode | String | Modo de compra (e.g., "classified"). |
| listing_type_id | String | Tipo de publicación según paquete contratado. |
| start_time | String | Fecha de inicio de la publicación. |
| end_time | String | Fecha de finalización de la publicación. |
| status | String | Estado de la publicación (e.g., "active"). |
| permalink | String | Enlace público a la publicación. |
| thumbnail | String | URL de la imagen principal del ítem. |
| pictures | Array | Listado de imágenes del ítem. |
| attributes | Array | Atributos de la publicación (e.g., habitaciones, baños, superficie). |
| seller_contact | Object | Datos de contacto del vendedor. |
| location | Object | Datos de ubicación del inmueble. |
| geolocation | Object | Latitud y longitud. |
| date_created | String | Fecha de creación. |
| last_updated | String | Última fecha de actualización. |
| channels | Array | Canales de publicación (e.g., "marketplace"). |

Nota:

 Para publicaciones de los listing types 

silver, gold, gold_special y gold_premium

 que presenten 

requires_picture: true

, es 

obligatorio el envío de al menos 1 imagen

 en el array pictures. Las solicitudes sin imágenes serán rechazadas con HTTP 400. 

## Publica inmuebles en Portal Inmobiliario (Chile)

En Chile (MLC), para que el ítem quede correctamente cargado en MercadoLibre y Portal Inmobiliario, debes incluir dentro del body el atributo **CMG_SITE**, lo que incorporará el **listing_source: portalinmobiliario** dejando visible en ambos portales.

Para incluirlo, ejecutando la siguiente llamada.:

```javascript
{
    "id": "CMG_SITE",
    "name": "Sitio de origen",
    "value_id": null,
    "value name": "POI",
    "value_struct": null,
    "attribute_group_id": "OTHERS",
    "attribute_group_name": "Otros"
}
```

### Lecturas recomendadas

- [Publicaciones de tiendas oficiales](https://developers.mercadolibre.com.ar/es_ar/publicaciones-de-tiendas-oficiales-para-inmuebles)
- [Actualiza tu inmueble](https://developers.mercadolibre.com.ar/es_ar/actualiza-tus-publicaciones)
- [Ciclo de vida de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/ciclo-de-vida-de-las-publicaciones-de-inmuebles)
- [Calidad de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/calidad-de-las-publicaciones-inmuebles)
- [Leads](https://developers.mercadolibre.com.ar/es_ar/leads-inmuebles)

## Actualizaciones de versión

Esta sección proporciona información sobre las actualizaciones de la API, incluyendo:

### Historial de cambios

| Fecha | Versión | Descripción |
| --- | --- | --- |
| 06/11/2025 | 1.0 | Publicación Inicial |
