---
title: Dominios y Categorías
slug: dominios-y-categorias
section: 03-domains-categories
source: https://developers.mercadolibre.com.ve/es_ar/dominios-y-categorias
captured: 2026-06-06
---

# Dominios y Categorías

> Source: <https://developers.mercadolibre.com.ve/es_ar/dominios-y-categorias>

## Endpoints referenced

- `https://api.mercadolibre.com/categories/MLA109291/attributes`
- `https://api.mercadolibre.com/categories/MLA1743/classifieds_promotion_packs`
- `https://api.mercadolibre.com/categories/MLA5725`
- `https://api.mercadolibre.com/domains/MLA-JACKETS_AND_COATS/technical_specs`
- `https://api.mercadolibre.com/sites`
- `https://api.mercadolibre.com/sites/MLA/categories`
- `https://api.mercadolibre.com/sites/MLA/domain_discovery/search?limit=1&q=celular%20iphone`
- `https://api.mercadolibre.com/sites/MLA/listing_exposures`
- `https://api.mercadolibre.com/sites/MLA/listing_prices?price=1`

## Content

Última actualización 30/12/2025

# Dominios y Categorías

Estos ejemplos servirán para trabajar con las opciones del árbol de categorías y listas de Mercado Libre. Algunos conceptos importantes a tener en cuenta:

- **Sitio**: sitio donde Mercado Libre está disponible para operaciones comerciales identificados con tres letras mayúsculas por ejemplo: MLA para Argentina, MLB para Brasil o MLM para México.
- **Dominio**: el dominio indica a qué familia de productos pertenece la publicación, puede tener una o muchas categorías asociadas, como por ejemplo CELLPHONES, SNEAKERS o BICYCLES.
- **Categoría**: es una clasificación de productos similares que comparten características comunes, dentro de un dominio. Las categorías se utilizan para organizar y presentar los productos de manera estructurada para facilitar la navegación y búsqueda de los clientes, como por ejemplo: CELLPHONE_COVERS o SLIPPERS.

### `/sites`

Devuelve información sobre los sitios donde Mercado Libre está disponible.

#### GET

**Obtener todos los sitios.**

```bash
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/sites
```

**Respuesta**

```json
[
    {
        "default_currency_id": "CRC",
        "id": "MCR",
        "name": "Costa Rica"
    },
    {
        "default_currency_id": "ARS",
        "id": "MLA",
        "name": "Argentina"
    },
    {
        "default_currency_id": "MXN",
        "id": "MLM",
        "name": "Mexico"
    },
    {
        "default_currency_id": "VES",
        "id": "MLV",
        "name": "Venezuela"
    },
    {
        "default_currency_id": "PYG",
        "id": "MPY",
        "name": "Paraguay"
    },
    {
        "default_currency_id": "CLP",
        "id": "MLC",
        "name": "Chile"
    },
    {
        "default_currency_id": "UYU",
        "id": "MLU",
        "name": "Uruguay"
    },
    {
        "default_currency_id": "PAB",
        "id": "MPA",
        "name": "Panamá"
    },
    {
        "default_currency_id": "DOP",
        "id": "MRD",
        "name": "Dominicana"
    },
    {
        "default_currency_id": "CUP",
        "id": "MCU",
        "name": "Cuba"
    },
    {
        "default_currency_id": "PEN",
        "id": "MPE",
        "name": "Perú"
    },
    {
        "default_currency_id": "COP",
        "id": "MCO",
        "name": "Colombia"
    },
    {
        "default_currency_id": "BRL",
        "id": "MLB",
        "name": "Brasil"
    },
    {
        "default_currency_id": "HNL",
        "id": "MHN",
        "name": "Honduras"
    },
    {
        "default_currency_id": "USD",
        "id": "MEC",
        "name": "Ecuador"
    },
    {
        "default_currency_id": "BOB",
        "id": "MBO",
        "name": "Bolivia"
    },
    {
        "default_currency_id": "USD",
        "id": "MSV",
        "name": "El Salvador"
    },
    {
        "default_currency_id": "NIO",
        "id": "MNI",
        "name": "Nicaragua"
    },
    {
        "default_currency_id": "GTQ",
        "id": "MGT",
        "name": "Guatemala"
    }
]
```

## Recategorización automática

Importante:

 A partir del 29/10/2025, implementamos un proceso de recategorización automática para ítems publicados vía API.

Es posible que algunos ítems sean **recategorizados automáticamente** por nuestro sistema, especialmente si detectamos que la categoría seleccionada originalmente no es la más precisa.

Por ese motivo, **recomendamos que tu integración esté escuchando activamente los cambios en los ítems**. De esta manera podrás detectar cambios de categoría en tiempo real y actuar en consecuencia.

Por ejemplo, para determinar el **precio**, costos logísticos, comisiones o condiciones de publicación, te sugerimos que evalúes el impacto que puede tener una recategorización automática. 

En particular:

- Valida si tu sistema necesita recalcular o actualizar el precio cuando cambia la categoría del ítem.
- Considera revisar otras lógicas sensibles a la categoría, como condiciones de envío o filtros de visibilidad.

**Siguiente**: [Ubicación y Moneda](https://developers.mercadolibre.com.ar/es_ar/ubicacion-y-monedas).
