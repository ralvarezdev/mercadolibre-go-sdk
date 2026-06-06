---
title: Opiniones de productos
slug: opiniones-sobre-producto
section: 19-metrics
source: https://developers.mercadolibre.com.ve/es_ar/opiniones-sobre-producto
captured: 2026-06-06
---

# Opiniones de productos

> Source: <https://developers.mercadolibre.com.ve/es_ar/opiniones-sobre-producto>

## Endpoints referenced

- `https://api.mercadolibre.com/reviews/item/$ITEM_ID`
- `https://api.mercadolibre.com/reviews/item/MLB1625519814`
- `https://api.mercadolibre.com/reviews/item/MLB1632704547?catalog_product_id=MLB14186226`

## Content

Última actualización 28/04/2023

## Opiniones de productos

Una vez entregado el producto al comprador, este podrá opinar de acuerdo a su experiencia indicando cuántas estrellas le daría al producto, realizar un comentario y, en caso de ser un producto de Moda, indicar si le quedó como esperaba o no. De esta manera, el vendedor podrá saber el promedio de estrellas sobre sus productos. En Mercado Libre, verán las estrellas y cantidad de opiniones debajo del título de la publicación de la siguiente manera:

Con la siguiente llamada, puedes averiguar las estrellas (campo rating_average) que tiene un producto según sus opiniones.
 Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/reviews/item/$ITEM_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/reviews/item/MLB1625519814
```

Respuesta:

```javascript
{
  "paging": {},
  "reviews": [
    {},
    {},
    {
      "id": 41672915,
      "reviewable_object": {},
      "date_created": "2019-06-08T14:12:29Z",
      "status": "published",
      "title": "Iincreíble, lo amo",
      "content": "Impresionante, muy satisfecha con el samsung s9, tenía un s7 y son todos excelentes! lo único que le encontré es q luego de una actualización de software, que x lo gral lo programo para la noche, el telefono no vuelve a encender solo, y por eso no suena la alarma de la mañana. Todavia no he terminado de descubrir todo, pero la rapidez y reaccion de todo el teléfono es genial.",
      "rate": 5,
      "valorization": 0,
      "likes": 0,
      "dislikes": 0,
      "reviewer_id": 0,
      "buying_date": "2019-04-12T04:00:00Z",
      "relevance": 71,
      "forbidden_words": 0
    },
    {},
    {}
  ],
  "rating_average": 4.8,
  "rating_levels": {
    "one_star": 2,
    "two_star": 1,
    "three_star": 0,
    "four_star": 9,
    "five_star": 80
  },
  "helpful_reviews": {},
  "attributes": [
  ]
}
```

 Nota: 

La información del parámetro 

reviewer_id

 se encuentra ofuscada y el mismo devuelve el valor (0), en todos los casos.

### Campos de la respuesta

**reviews**: opiniones de producto

- **id**: id de la opinión
- **date_created**: fecha de la creación de la opinión
- **status**: estado de la opinión (por ejemplo, publicada)
- **tittle**: título de la opinión
- **content**: contenido de la opinión
- **rate**: valoración en estrellas (de 1 a 5)
- **buying_date**: fecha de la compra del producto

**rating_average**: promedio de las valoraciones totales
 **rating_levels**: clasificación por estrellas recibidas

## Consultar opiniones de ítem de catálogo

Utilizando el catalog_product_id correspondiente al ítem con optin a catálogo, puedes consultar sus opiniones. Conoce más cómo [publicaciones de Catálogo](https://developers.mercadolibre.com.ve/es_ar/llego-catalogo-conoce-como-adaptar-tu-integracion).

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/reviews/item/MLB1632704547?catalog_product_id=MLB14186226
```

Respuesta:

```javascript
{
  "paging": {},
  "reviews": [
    {},
    {},
    {},
    {
      "id": 43083524,
      "reviewable_object": {},
      "date_created": "2019-07-17T23:11:29Z",
      "status": "published",
      "title": "Perfecto¡¡¡",
      "content": "Excelente producto, recomendable 100%.",
      "rate": 5,
      "valorization": 2,
      "likes": 2,
      "dislikes": 0,
      "reviewer_id": 0,
      "buying_date": "2019-06-22T04:00:00Z",
      "relevance": 4,
      "forbidden_words": 0,
      "attributes": []
    },
    {}
  ],
  "rating_average": 4.8,
  "rating_levels": {},
  "helpful_reviews": [
  ],
  "attributes": [
  ]
}
```

Conoce más sobre [Métricas de venta](https://vendedores.mercadolibre.com.ar/nota/seccion-de-metricas-de-venta-conoce-de-que-se-trata/).
