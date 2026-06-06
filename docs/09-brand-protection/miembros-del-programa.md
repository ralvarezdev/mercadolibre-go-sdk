---
title: Miembros del Programa
slug: miembros-del-programa
section: 09-brand-protection
source: https://developers.mercadolibre.com.ve/es_ar/miembros-del-programa
captured: 2026-06-06
---

# Miembros del Programa

> Source: <https://developers.mercadolibre.com.ve/es_ar/miembros-del-programa>

## Endpoints referenced

- `https://api.mercadolibre.com/moderations/pppi/case/$DENOUNCE_ID`
- `https://api.mercadolibre.com/moderations/pppi/case/123`
- `https://api.mercadolibre.com/moderations/pppi/denounces/$SITE_ID/ITM/options`
- `https://api.mercadolibre.com/moderations/pppi/denounces/MLA/ITM/options`
- `https://api.mercadolibre.com/moderations/pppi/denounces/items/$ITEM_ID`
- `https://api.mercadolibre.com/moderations/pppi/denounces/items/MLA123`

## Content

Última actualización 06/05/2026

## Miembros del Programa

 Importante: 

Esta funcionalidad puede utilizarla solo quien sea Miembro del Brand Protection Program. Si no lo eres, puedes 

adherirte como Miembro al Programa

. 

Con las siguientes APIs, el BPP invita a quienes sean titulares de derechos, o sus representantes legales, a proteger todo su portafolio de derechos de propiedad intelectual mediante la denuncia de cualquier publicación que, presuntamente, pueda infringir sus derechos de propiedad intelectual.

## Consultar motivos habilitados

Conoce cuáles son los motivos que tienes habilitados como Miembro para denunciar.

 Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/denounces/$SITE_ID/ITM/options
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/denounces/MLA/ITM/options
```

Respuesta OK (200):

```javascript
[
   {
      "id":"PPPI2",
      "group":"PPPI",
      "type":"Product",
      "description":"Uso ilegítimo de marca registrada",
      "description_en":"Unlawful use of trademark",
      "sub_text":"Por ejemplo, dice que es mi distribuidor oficial cuando no lo es, incluye mis logos en la descripción o en las imágenes de la publicación.",
      "sub_text_en":"For example, says it is my official distributor when it is not, includes my logos in the description or in the images of the listing."
   },
   "..."
]
```

### Realizar denuncia

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/denounces/items/$ITEM_ID
```

Ejemplo:

```javascript
curl -X POST  -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/denounces/items/MLA123 -d
{
  "report_reason_id": "PPPI1",
  "comment": "Comment example.",
}
```

## Derecho de autor para imágenes

Si quieres realizar denuncias en derechos de autor para imágenes, deberás indicar en el body cuáles son las imágenes infractoras. Esto será obligatorio para motivos de denuncia PPPI4, PPPI6, PPPI7, PPPI17. Puedes obtener los ID de imágenes infractoras API de /items, del atributo pictures.

```javascript
curl -X POST  -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/denounces/items/MLA123 -d
{   "report_reason_id":"PPPI1",
   "comment":"Comment example.",
   "photos_denounced": [
      "666591-MLA26622267232_012016",
      "666591-MLA26622267232_012017",
      "666591-MLA26622267232_012018"
   ]
}
```

Respuesta ok (200):

```javascript
{
  "status": 201,
  "denounce_id": 12547408
}
```

## Consultar estado de denuncia

Obtén información del estado actual de la denuncia para luego responder al vendedor que has denunciado.

Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/case/$DENOUNCE_ID
```

Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/case/123
```

Respuesta ok (200):

```javascript
{
  "item_info" {
    "title": ,
    "description": < string | null >,
    "price": < number | null >, 
    "pictures": < array[object: "url":  ] >,
  },
  "user_type": < string(seller|member) >,
  "reason_text": ,
  "member_name": ,
  "member_quittance": < string | null >,
  "seller_name": ,
  "seller_quittance": < string | null >,
  "document_url": < string | null >,
  "document_name": < string | null >,
  "due_date": ,
"current_status": < string(WAITING_DOCUMENTATION|DOCUMENTATION_PRESENTED|DOCUMENTATION_NOT_PRESENTED|DOCUMENTATION_APPROVED|DOCUMENTATION_NOT_APPROVED|MEMBER_NOT_RESPOND|ROLLBACK|DISCARD_DUE_RESTRICTION) >,
"reject_option_member": [
    {
      "sub_text_en": null,
      "text_en": "The documentation does not correspond to the reported product",
      "id": 1,
      "text_pt": "A documentação não corresponde ao produto denunciado",
      "sub_text_pt": null,
      "text_es": "La documentación no se corresponde con el producto denunciado",
      "sub_text_es": null
    },
    {
      "sub_text_en": null,
      "text_en": "The documentation is illegible",
      "id": 2,
      "text_pt": "A documentação está ilegível",
      "sub_text_pt": null,
      "text_es": "La documentación es ilegible",
      "sub_text_es": null
    },
    {
      "sub_text_en": "The documentation does not prove that they are authorized to use my brands, logos, or that they are official distributors",
      "text_en": "You are not authorized to use this content",
      "id": 3,
      "text_pt": "Não está autorizado a usar este conteúdo",
      "sub_text_pt": "A documentação não comprova que está autorizado a usar minhas marcas, logotipos ou que é um distribuidor oficial",
      "text_es": "No está autorizado a utilizar este contenido",
      "sub_text_es": "La documentación no prueba que está autorizado a usar mis marcas, logos, ni que es un distribuidor oficial"
    }
],
"photos_denounced": [
       {
           "id": "670708-MLA40946169781_022020"
           "status": "REMOVED",
           "src": "http://mla-s2-p.mlstatic.com/670708-MLA40946169781_022020.jpg"
       }
   ],
   "photos_new": [
       {
           "id": "8889-MLA26622267232_012016",
           "src": "http://mla-s2-p.mlstatic.com/670708-MLA40946169781_022020.jpg"
       },
       {
          "id": "792503-MLA40997189396_032020",
           "src": "http://mla-s2-p.mlstatic.com/792503-MLA40997189396_032020-O.jpg"
       }
   ]
}
```

## Responder vendedor

Una vez que realizas una denuncia como Miembro del Programa, el vendedor tiene **3 días** para responder con la documentación y luego tendrás **3 días corridos** para responder a este.

- **Si no respondes dentro del plazo**, la publicación denunciada del vendedor será reactivada.
- **Para responder rechazando la respuesta del vendedor**, debes agregar dentro del body el campo **reject_member_id** con el id del motivo de rechazo obtenido previamente en la llamada GET dentro de campo **reject_option_member**.

 Nota: 

Sólo puedes responder ante casos con estado DOCUMENTATION_PRESENTED.

Llamada:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/moderations/pppi/case/$DENOUNCE_ID -d
{
  "documentation_approved":"false" ,
  "member_quittance": ,
  "reject_member_id": "1"
}
```

Ejemplo aprobar denuncia:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'  https://api.mercadolibre.com/moderations/pppi/case/$DENOUNCE_ID -d
{
  "documentation_approved":"true" ,
  "member_quittance": 
}
```

Ejemplo rechazar denuncia:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN'  https://api.mercadolibre.com/moderations/pppi/case/$DENOUNCE_ID -d
{
  "documentation_approved":"false",
  "member_quittance": < string | null >,
 "reject_member_id": "1"
}
```

Respuesta:

```javascript
{
  "item_info" {
    "title": ,
    "description": < string |null >,
    "price": < number | null >, 
    "pictures": < array[object: "url":  ] >,
  },
  "user_type": < string(seller|member) >,
  "reason_text": ,
  "member_name": ,
  "member_quittance": < string | null >,
  "seller_name": ,
  "seller_quittance": < string |null >,
  "document_url": < string |null >,
  "document_name": < string |null >,
  "due_date": ,
"current_status": < string(WAITING_DOCUMENTATION|DOCUMENTATION_PRESENTED|DOCUMENTATION_NOT_PRESENTED|DOCUMENTATION_APPROVED|DOCUMENTATION_NOT_APPROVED|MEMBER_NOT_RESPOND|ROLLBACK|DISCARD_DUE_RESTRICTION) >,
"reject_option_member": [
    {
      "sub_text_en": null,
      "text_en": "The documentation does not correspond to the reported product",
      "id": 1,
      "text_pt": "A documentação não corresponde ao produto denunciado",
      "sub_text_pt": null,
      "text_es": "La documentación no se corresponde con el producto denunciado",
      "sub_text_es": null
    },
    {
      "sub_text_en": null,
      "text_en": "The documentation is illegible",
      "id": 2,
      "text_pt": "A documentação está ilegível",
      "sub_text_pt": null,
      "text_es": "La documentación es ilegible",
      "sub_text_es": null
    },
    {
      "sub_text_en": "The documentation does not prove that they are authorized to use my brands, logos, or that they are official distributors",
      "text_en": "You are not authorized to use this content",
      "id": 3,
      "text_pt": "Não está autorizado a usar este conteúdo",
      "sub_text_pt": "A documentação não comprova que está autorizado a usar minhas marcas, logotipos ou que é um distribuidor oficial",
      "text_es": "No está autorizado a utilizar este contenido",
      "sub_text_es": "La documentación no prueba que está autorizado a usar mis marcas, logos, ni que es un distribuidor oficial"
    }
],
"photos_denounced": [
       {
           "id": "670708-MLA40946169781_022020"
           "status": "REMOVED",
           "src": "http://mla-s2-p.mlstatic.com/670708-MLA40946169781_022020.jpg"
       }
   ],
   "photos_new": [
       {
           "id": "8889-MLA26622267232_012016",
           "src": "http://mla-s2-p.mlstatic.com/670708-MLA40946169781_022020.jpg"
       },
       {
          "id": "792503-MLA40997189396_032020",
           "src": "http://mla-s2-p.mlstatic.com/792503-MLA40997189396_032020-O.jpg"
       }
   ]
}
```

**Siguiente**: [Publicaciones denunciadas](https://developers.mercadolibre.com.ve/es_ar/publicaciones-denunciadas?nocache=true).
