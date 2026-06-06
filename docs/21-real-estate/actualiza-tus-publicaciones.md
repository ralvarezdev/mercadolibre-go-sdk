---
title: Actualiza tus publicaciones
slug: actualiza-tus-publicaciones
section: 21-real-estate
source: https://developers.mercadolibre.com.ve/es_ar/actualiza-tus-publicaciones
captured: 2026-06-06
---

# Actualiza tus publicaciones

> Source: <https://developers.mercadolibre.com.ve/es_ar/actualiza-tus-publicaciones>

## Endpoints referenced

- `https://api.mercadolibre.com/items/$ITEM_ID`
- `https://api.mercadolibre.com/items/{ITEM_ID}/listing_type`

## Content

Última actualización 04/03/2026

## Actualiza tus publicaciones

Importante:

A partir del 

12 de marzo de 2026

, las solicitudes de actualización para anuncios de los tipos 

silver, gold, gold_special y gold_premium

 que presenten 

requires_picture: true

serán rechazadas (HTTP 400) en caso de que no posean imágenes

.Asegurate de incluir al menos una imagen en el array pictures para estos tipos de anuncios. 

Una vez que tienes publicaciones activas en Mercado Libre, es probable que debas actualizarlas periódicamente. Esto te permitirá sincronizar la información, excluir anuncios de propiedades ya vendidas, pausar publicaciones, optimizar las descripciones y actualizar los precios, entre otras acciones.

Para actualizar publicaciones, envía una solicitud **PUT** a la URL del recurso **/items/{item_id}** con el *id* del ítem a actualizar. En el *body* debe ir el JSON con **solo** los campos a modificar con sus nuevos valores. Los campos omitidos no se alterarán.

Nota:

 Los campos que pueden modificarse son: 

- **Title**
- **Price**
- **Video**
- **Pictures**
- **Description**
- **Location**
- **Atributos de la publicación** (array **attributes**)
- **Category**

Veamos un ejemplo básico de actualización del precio de una publicación. Lo único que necesitas es el **item_id** del inmueble publicado. No olvides incluir el **access_token** para la autenticación.

Llamada:

```javascript
 curl -X PUT \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  'https://api.mercadolibre.com/items/$ITEM_ID' \
  -d '{
    "price": 100000001
  }'
```

| Parámetro | Tipo | Opcional | Valores |
| --- | --- | --- | --- |
| ACCESS_TOKEN | string | No | Recuerda utilizar el token que generaste en la guía de autenticación |
| item_id | string | No | Id del item que obtuviste al momento de publicarlo |
| price | string | No | Nuevo precio de la publicación |

Al ejecutar cualquier actualización, recibirás la respuesta con **status 200** y el cuerpo como si hubieras publicado un ítem, con los datos actualizados. Puedes consultar esta información en [Publica Inmuebles](https://developers.mercadolibre.com.ar/es_ar/publica-inmueble). A continuación, seguimos con otros ejemplos.

## Destacar una publicación

Para aumentar la visibilidad de una publicación, debes realizar una petición **POST** para actualizar el campo **listing_type**.

Importante:

Recuerda que para realizar esta acción, el vendedor debe haber 

contratado previamente un paquete destacado

 (Oro o Oro Premium).

Debes contar con el **id** del ítem que publicaste y que deseas destacar.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/items/{ITEM_ID}/listing_type -d '{"id":"gold_premium"}'
```

| Parámetro | Tipo | Opcional | Valores |
| --- | --- | --- | --- |
| ACCESS_TOKEN | string | No | Recuerda utilizar el token que generaste en la guía de autenticación |
| item_id | string | No | Id del item que obtuviste al momento de publicarlo |
| id | string | No | Nombre del paquete de destaque a usar. |

Recibirás la respuesta con **status 200** y el cuerpo como si hubieras publicado un ítem, con los datos actualizados. Puedes consultar esta información en [Publica Inmuebles](https://developers.mercadolibre.com.ar/es_ar/publica-inmueble), pero con el dato actualizado de **listing_type** con valor **gold** o **gold_premium**.

```javascript
"buying_mode": "classified",
"listing_type_id": "gold_premium",
"start_time": "2025-06-11T21:25:54.260Z",
"stop_time": "2025-12-08T04:00:00.000Z",
"end_time": "2025-12-08T04:00:00.000Z",
```

Es necesario tener en cuenta que:

- No se genera ningún cargo al realizar esta actualización (API call).
- Si se revierte la actualización, el cupo del paquete destacado vuelve a estar disponible.
- No es posible destacar una publicación sin tener contratado al menos un paquete destacado.

## Modifica la localización de tu inmueble.

Si por alguna situación requieres actualizar la ubicación de tu inmueble, una vez tengas toda la información según la guía de [Localizar inmuebles](https://developers.mercadolibre.com.ar/es_ar/localizar-inmuebles), envía una petición con los datos a actualizar.

```javascript
curl -X PUT \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  'https://api.mercadolibre.com/items/$ITEM_ID' \
  -d '{
    "location": {
      "address_line": "My property address NEW 111",
      "zip_code": "5000",
      "neighborhood": {
        "id": "TUxBQlBBTDI1MTVa",
        "name": "Palermo"
      },
      "city": {
        "id": "TUxBQ0NBUGZlZG1sYQ",
        "name": "Capital Federal"
      }
    }
  }
```

Recibirás la respuesta como si hubieras publicado un ítem, con los datos actualizados. Puedes consultar esta información en [Publica Inmuebles](https://developers.mercadolibre.com.ar/es_ar/publica-inmueble)

## Actualiza la tienda oficial del Inmueble

Si requieres actualizar la tienda oficial con la cual fue registrado tu ítem, deberás hacer una petición tipo **PUT** al endpoint **/items**. La respuesta corresponderá a la misma de publicación, con la diferencia en el campo **official_store_id**.

```javascript
curl -X PUT \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  'https://api.mercadolibre.com/items/$ITEM_ID' \
  -d '{
   "official_store_id": 3121
  }
 
```

Recibirás la respuesta como si hubieras publicado un ítem, con los datos actualizados. Puedes consultar esta información en [Publica Inmuebles](https://developers.mercadolibre.com.ar/es_ar/publica-inmueble)

## Cambia el estado de tus publicaciones.

Cada ítem publicado en el marketplace puede encontrarse en uno de varios estados que definen su disponibilidad y visibilidad para los usuarios. A continuación se describen los estados disponibles y su comportamiento.

| Estado | Descripción |
| --- | --- |
| closed | Finaliza la publicación de forma definitiva, el ítem no puede reactivarse. |
| paused | Pausa la publicación temporalmente. Mientras está pausada, el ítem no estará visible para los usuarios ni podrá generar contactos, ya que sus datos se ocultan. |
| active | Reactiva un ítem que se encontraba en estado pausado, haciendo que vuelva a estar visible y disponible para los usuarios. |

Para cambiar el estado de un ítem, envía uno de los valores descritos para el campo **status** en la solicitud de actualización. El valor debe enviarse en minúsculas y respetando exactamente la palabra, ya que es sensible a mayúsculas y minúsculas.

Ejemplo:

```javascript
curl -X PUT \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -H "Accept: application/json" \
  'https://api.mercadolibre.com/items/$ITEM_ID' \
  -d '{
    "status":"paused"
  }'
```

De este modo podrás pausar o activar tu publicación. Por otro lado, si por alguna razón actualizas tu publicación a cerrado (closed) o se encuentra en este estado y quieres volver a publicarla, consulta nuestro artículo sobre [volver a publicar (republicar)](https://www.mercadolibre.com.ar/ayuda/Re-publicar_1071).

Para más información sobre los estados de tu inmueble, puedes consultar [Ciclo de vida de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/ciclo-de-vida-de-las-publicaciones-de-inmuebles).

### Guía para algunos campos

**1. Actualizar descripción**

Al igual que los anteriores campos, la descripción se puede actualizar fácilmente, sin embargo, si deseas hacerlo de la mejor manera, puedes seguir los consejos descritos en la [guía de descripción de productos](https://developers.mercadolibre.com.ar/es_ar/descripcion-de-articulos). Esto te permitirá tener en cuenta algunos puntos importantes.

**2. Actualización de imágenes**

Siempre puedes agregar o reemplazar las imágenes o videos de tus publicaciones; puedes consultar el tutorial de [Trabajar con imágenes](https://developers.mercadolibre.com.ar/es_ar/trabajar-con-imagenes) para conocer la mejor forma de hacerlo.

**3. Actualizar tipo de publicación**

Cuando deseas lograr más exposición para tu inmueble, debes actualizar el tipo de publicación. Para entender más acerca de cómo funciona, consulta la documentación de [Tipos de publicación y cómo obtener mayor exposición](https://developers.mercadolibre.com.ar/es_ar/tipos-de-publicacion-y-actualizaciones-de-articulos).

## Elimina Publicaciones

Eliminar una publicación es una acción irreversible que finaliza y elimina permanentemente el ítem del marketplace. Por esta razón, se recomienda realizar esta operación con precaución.

### Proceso para eliminar un ítem

1. **1.Cambiar el estado del ítem a *closed*:** En este paso se finaliza la publicación, al cambiar su estado a **closed**. `curl -X PUT \ -H "Authorization: Bearer $ACCESS_TOKEN" \ -H "Content-Type: application/json" \ -H "Accept: application/json" \ 'https://api.mercadolibre.com/items/$ITEM_ID' \ -d '{ "status":"closed" }'`
2. **2.Eliminar el ítem definitivamente:** Una vez el ítem se encuentre en estado **closed**, ejecuta el siguiente comando para eliminarlo permanentemente. `curl -X PUT "https://api.mercadolibre.com/items/$ITEM_ID" \ -H "Authorization: Bearer $ACCESS_TOKEN" \ -H "Content-Type: application/json" \ -H "Accept: application/json" \ -d '{ "deleted": true }'`

Al igual que cualquier actualización mencionada, recibirás una respuesta como si hubieras publicado un inmueble, con la variante de que en **sub_status** aparecerá **deleted**:

```javascript
    "status": "closed",
    "sub_status": [
        "deleted",
        "pack_quota_assigned"
    ],
```

### Error común al eliminar tu publicación.

Si al ejecutar el paso anterior recibes el siguiente error:

```javascript
{
  "message": "item optimistic locking error: conflict",
  "status": 409,
  "cause": []
}
```

Significa que la información del ítem aún no ha sido completamente actualizada en el sistema. En este caso, espere unos segundos y vuelva a intentar la llamada.

### Consideraciones finales al eliminar una publicación:

- Después de eliminar tu inmueble, puede seguir apareciendo temporalmente en la vista del vendedor con el mensaje de *"publicación finalizada"*. Esto es normal y desaparecerá automáticamente en poco tiempo.
- Recuerda que la eliminación es definitiva: no es posible reactivar ni recuperar el ítem eliminado.

### Lecturas recomendadas

- [Ciclo de vida de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/ciclo-de-vida-de-las-publicaciones-de-inmuebles).
- [Calidad de las publicaciones](https://developers.mercadolibre.com.ar/es_ar/calidad-de-las-publicaciones-inmuebles).

## Actualizaciones de versión

Esta sección proporciona información sobre las actualizaciones de la API, incluyendo:

### Historial de cambios

| Fecha | Versión | Descripción |
| --- | --- | --- |
| 06/11/2025 | 1.0 | Publicación Inicial |
