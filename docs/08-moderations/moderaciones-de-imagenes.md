---
title: Moderaciones de imágenes
slug: moderaciones-de-imagenes
section: 08-moderations
source: https://developers.mercadolibre.com.ve/es_ar/moderaciones-de-imagenes
captured: 2026-06-06
---

# Moderaciones de imágenes

> Source: <https://developers.mercadolibre.com.ve/es_ar/moderaciones-de-imagenes>

## Endpoints referenced

- `https://api.mercadolibre.com/moderations/last_moderation/$MODERATION_REFERENCE_ID`
- `https://api.mercadolibre.com/pictures/items/upload`

## Content

Última actualización 21/07/2025

## Moderaciones de imágenes

Las publicaciones pueden ser moderadas debido a problemas de calidad en las imágenes. Estas moderaciones pueden presentarse con el status **active** o **paused** y el tag **poor_quality_thumbnail**. 
 Conoce recomendaciones para sacar buenas fotos de [Productos](https://www.mercadolibre.com.ar/ayuda/Sacar-bue-nas-fotos-productos_805), [Moda](https://www.mercadolibre.com.ar/ayuda/22140), [Vehículos](https://www.mercadolibre.com.ar/ayuda/Sacar-bue-nas-fotos-vehiculos_806), [Inmuebles](https://www.mercadolibre.com.ar/ayuda/Sacar-buenas-fotos-de-inmuebles_807) y Servicios.

## Flujo recomendado para carga sin moderaciones

1. Utiliza la [API de diagnóstico de imágenes](https://developers.mercadolibre.com.ve/diagnostico-imagenes?nocache=true) para imágenes base64, picture_id (para fotos existentes en el CDN de Meli) y URL
2. Para cargar imágenes en el CDN de Meli, podes utilizar el endpoint https://api.mercadolibre.com/pictures/items/upload.
