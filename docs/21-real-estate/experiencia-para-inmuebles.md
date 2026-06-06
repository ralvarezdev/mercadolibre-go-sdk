---
title: Experiencia para inmuebles
slug: experiencia-para-inmuebles
section: 21-real-estate
source: https://developers.mercadolibre.com.ve/es_ar/experiencia-para-inmuebles
captured: 2026-06-06
---

# Experiencia para inmuebles

> Source: <https://developers.mercadolibre.com.ve/es_ar/experiencia-para-inmuebles>

## Endpoints referenced

- `https://api.mercadolibre.com/vis-transactions-hub/configurations/provider`
- `https://api.mercadolibre.com/vis-transactions-hub/configurations/provider/{providerId}`
- `https://api.mercadolibre.com/vis-transactions-hub/configurations/provider/{providerId}?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/configurations/provider?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/configurations/seller/{sellerId}`
- `https://api.mercadolibre.com/vis-transactions-hub/configurations/seller/{sellerId}?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/items`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/items/tags`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/items/tags?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/items?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/schedules`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/schedules/{scheduleId}`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/schedules/{scheduleId}?scope=stage`
- `https://api.mercadolibre.com/vis-transactions-hub/{providerId}/entities/schedules?scope=stage`
- `https://api.mercadolibre.com/vis/leads/$LEAD_ID`

## Content

Última actualización 19/01/2026

## Experiencia para inmuebles

 Importante: 

Este recurso fue descontinuado el 31 de marzo de 2025 para los sitios MLM, MLU y MLA, y 

permanece disponible exclusivamente para el sitio MLC

. 

Desde Mercado Libre queremos generar una nueva experiencia de arriendo y venta de propiedades para el buyer a través de publicaciones que ofrezcan una experiencia superior con foco en asegurar la disponibilidad de los inmuebles, y garantizar la respuesta a dudas y solicitudes de agenda en tiempo y forma.

Esta guía tiene como objetivo describir y ejemplificar cómo Mercado Libre se conecta con los partners que hoy están dentro del modelo de solicitud de visita, el customer journey y en qué etapa participa cada uno de los actores involucrados dentro del mismo.

## Objetivos

- Describir brevemente el customer journey de solicitud de visita.
- Describir cada uno de los webhooks y cuales son los dominios de variable que acepta.
- Entregar lineamientos técnicos de cómo integrarse con nuestro sistema para poder hacer uso de la actualizacion en tiempo real y APIs.
