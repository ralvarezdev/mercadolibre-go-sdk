---
title: Preguntas y Respuestas
slug: preguntas-y-respuestas
section: 05-questions
source: https://developers.mercadolibre.com.ve/es_ar/preguntas-y-respuestas
captured: 2026-06-06
---

# Preguntas y Respuestas

> Source: <https://developers.mercadolibre.com.ve/es_ar/preguntas-y-respuestas>

## Endpoints referenced

- `https://api.mercadolibre.com/answers`
- `https://api.mercadolibre.com/block-api/search/users/72641919?type=blocked_by_questions`
- `https://api.mercadolibre.com/my/received_questions/search`
- `https://api.mercadolibre.com/questions`
- `https://api.mercadolibre.com/questions/3957150025`
- `https://api.mercadolibre.com/questions/search?item_id=MLA608007087`
- `https://api.mercadolibre.com/users/$SELLER_ID/questions_blacklist/$USER_ID`

## Content

Última actualización 05/06/2025

## Preguntas y Respuestas

 Nota: 

Por seguridad, puedes obtener el email, teléfono y nombre del comprador en en /questions/$QUESTION_ID.

 Utiliza el parámetro api_version=4 y obtén las preguntas y respuestas con la nueva estructura.

Este ejemplo servirá para gestionar preguntas y respuestas

### `/questions/search?item=$ITEM_ID`

Buscar cualquier pregunta realizada en los ítems del usuario.

#### GET

**Obtener preguntas por ID de ítem**

```bash
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/questions/search?item_id=MLA608007087
```

**Respuesta**

```json
{
    "total": 0,
    "limit": 50,
    "questions": [],
    "filters": {
        "limit": 50,
        "offset": 0,
        "api_version": "4",
        "is_admin": false,
        "sorts": [],
        "caller": 447594313,
        "item": "MLA608007087"
    },
    "available_filters": [
        {
            "id": "from",
            "name": "From user id",
            "type": "number"
        },
        {
            "id": "seller",
            "name": "Seller id",
            "type": "number"
        },
        {
            "id": "totalDivisions",
            "name": "total divisions",
            "type": "number"
        },
        {
            "id": "division",
            "name": "Division",
            "type": "number"
        },
        {
            "id": "status",
            "name": "Status",
            "type": "text",
            "values": [
                "ANSWERED",
                "BANNED",
                "CLOSED_UNANSWERED",
                "DELETED",
                "DISABLED",
                "UNANSWERED",
                "UNDER_REVIEW"
            ]
        }
    ],
    "available_sorts": [
        "item_id",
        "from_id",
        "date_created",
        "seller_id"
    ]
}
```
