---
title: Más vendidos en Mercado Libre
slug: mas-vendidos-en-mercado-libre
section: 19-metrics
source: https://developers.mercadolibre.com.ve/es_ar/mas-vendidos-en-mercado-libre
captured: 2026-06-06
---

# Más vendidos en Mercado Libre

> Source: <https://developers.mercadolibre.com.ve/es_ar/mas-vendidos-en-mercado-libre>

## Endpoints referenced

- `https://api.mercadolibre.com/highlights/$SITE_ID/category/$CATEGORY_ID`
- `https://api.mercadolibre.com/highlights/$SITE_ID/category/$CATEGORY_ID?attribute=BRAND&attributeValue=59387`
- `https://api.mercadolibre.com/highlights/$SITE_ID/item/$ITEM_ID`
- `https://api.mercadolibre.com/highlights/$SITE_ID/product/$PRODUCT_ID`
- `https://api.mercadolibre.com/highlights/MLA/category/MLA1055?attribute=BRAND&attributeValue=59387`
- `https://api.mercadolibre.com/highlights/MLB/category/MLB432825`
- `https://api.mercadolibre.com/highlights/MLB/product/MLB15960724`
- `https://api.mercadolibre.com/highlights/MLM/item/MLM16140709`

## Content

Última actualización 28/04/2023

## Más vendidos en Mercado Libre

Actualiza tu integración con el recurso **/highlights** para consultar el listado de los 20 productos más vendidos en Mercado Libre. Puedes consultarlos según la [categoría](https://developers.mercadolibre.com.ve/categoriza-productos), marca, el [producto](https://developers.mercadolibre.com.ve/buscador-de-productos#Producto-por-ID) y/o [ítem](https://developers.mercadolibre.com.ve/items-y-busquedas#Multiget).

## Más vendidos por categoría

Consulta el top 20 de items/productos de una categoría específica.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/$SITE_ID/category/$CATEGORY_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/MLB/category/MLB432825
```

Respuesta:

```javascript
{
    "query_data": {
        "highlight_type": "BEST_SELLER",
        "criteria": "CATEGORY",
        "id": "MLB432825"
    },
    "content": [
        {
            "id": "MLB1481736854",
            "position": 1,
            "type": "ITEM"
        },
        {
            "id": "MLB1530545830",
            "position": 2,
            "type": "ITEM"
        },
        {
            "id": "MLB1692302041",
            "position": 3,
            "type": "ITEM"
        },
        {
            "id": "MLB1480437372",
            "position": 4,
            "type": "ITEM"
        },
        {
            "id": "MLB1415573454",
            "position": 5,
            "type": "ITEM"
        },
        {
            "id": "MLB1775202155",
            "position": 6,
            "type": "ITEM"
        },
        {
            "id": "MLB1470331312",
            "position": 7,
            "type": "ITEM"
        },
        {
            "id": "MLB1917662040",
            "position": 8,
            "type": "ITEM"
        },
        {
            "id": "MLB1775121349",
            "position": 9,
            "type": "ITEM"
        },
        {
            "id": "MLB1646755234",
            "position": 10,
            "type": "ITEM"
        },
        {
            "id": "MLB1914214929",
            "position": 11,
            "type": "ITEM"
        },
        {
            "id": "MLB1773591467",
            "position": 12,
            "type": "ITEM"
        },
        {
            "id": "MLB1775162236",
            "position": 13,
            "type": "ITEM"
        },
        {
            "id": "MLB1690035398",
            "position": 14,
            "type": "ITEM"
        },
        {
            "id": "MLB1480477864",
            "position": 15,
            "type": "ITEM"
        },
        {
            "id": "MLB1775187545",
            "position": 16,
            "type": "ITEM"
        },
        {
            "id": "MLB1671977510",
            "position": 17,
            "type": "ITEM"
        },
        {
            "id": "MLB1658084769",
            "position": 18,
            "type": "ITEM"
        },
        {
            "id": "MLB1614538979",
            "position": 19,
            "type": "ITEM"
        },
        {
            "id": "MLB1627711451",
            "position": 20,
            "type": "ITEM"
        }
    ]
}
```

### Campos de la respuesta

**query_data**: información del tipo de filtro aplicado
 **content**:

- **id**: id del ítem o producto
- **position**: posicionamiento dentro del listado
- **type**: tipo de elemento (item/product)

## Más vendidos por categoría y atributo marca

Puedes obtener el Top 20 de items/productos de una marca en una categoría específica.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/$SITE_ID/category/$CATEGORY_ID?attribute=BRAND&attributeValue=59387
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/MLA/category/MLA1055?attribute=BRAND&attributeValue=59387
```

Respuesta:

```javascript
{
    "query_data": {
        "highlight_type": "BEST_SELLER",
        "criteria": "CATEGORY",
        "id": "MLA1055"
    },
    "content": [
        {
            "id": "MLA17706115",
            "position": 1,
            "type": "PRODUCT"
        },
        {
            "id": "MLA16150548",
            "position": 2,
            "type": "PRODUCT"
        },
        {
            "id": "MLA16141545",
            "position": 3,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17706117",
            "position": 4,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17369333",
            "position": 5,
            "type": "PRODUCT"
        },
        {
            "id": "MLA15946515",
            "position": 6,
            "type": "PRODUCT"
        },
        {
            "id": "MLA16141547",
            "position": 7,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17369332",
            "position": 8,
            "type": "PRODUCT"
        },
        {
            "id": "MLA16107499",
            "position": 9,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17464694",
            "position": 10,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17706118",
            "position": 11,
            "type": "PRODUCT"
        },
        {
            "id": "MLA918309355",
            "position": 12,
            "type": "ITEM"
        },
        {
            "id": "MLA17324819",
            "position": 13,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17426238",
            "position": 14,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17415926",
            "position": 15,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17457352",
            "position": 16,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17464698",
            "position": 17,
            "type": "PRODUCT"
        },
        {
            "id": "MLA16217275",
            "position": 18,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17706116",
            "position": 19,
            "type": "PRODUCT"
        },
        {
            "id": "MLA17493810",
            "position": 20,
            "type": "PRODUCT"
        }
    ]
}
```

## Posicionamiento del producto

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/$SITE_ID/product/$PRODUCT_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/MLB/product/MLB15960724
```

Respuesta:

```javascript
{
    "dimension": "category",
    "id": "MLB1055",
    "label": "Celulares e Smartphones",
    "position": 19
}
```

## Posicionamiento del ítem

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/$SITE_ID/item/$ITEM_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/highlights/MLM/item/MLM16140709
```

Respuesta:

```javascript
{
    "dimension": "category",
    "id": "MLM126793",
    "label": "Smartwatches",
    "position": 5
}
```

### Campos de la respuesta

**dimension**: criterio utilizado (category)

- **id**: id de la dimensión
- **label**: título de la dimensión
- **position**: posición del producto o ítem (depende la llamada) dentro del listado

## Errores

| Error | Descripción | Solución |
| --- | --- | --- |
| 400 | Error site: ML | Enviar site_id válido |
| 400 | Invalid product id MLB | Enviar product/item_id válido |
| 401 | unspecified_token | Enviar access token válido |
| 404 | item/product with id MLB15960724 not found | El ítem/producto no está listado en Más vendidos |
| 404 | Dimension CATEGORY with id MLB432825 not found | La categoría no tiene listado de Más vendidos |
