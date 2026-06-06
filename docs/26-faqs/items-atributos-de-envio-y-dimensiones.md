---
title: Items - Atributos de envío y dimensiones
slug: items-atributos-de-envio-y-dimensiones
section: 26-faqs
source: https://developers.mercadolibre.com.ve/es_ar/items-atributos-de-envio-y-dimensiones
captured: 2026-06-06
---

# Items - Atributos de envío y dimensiones

> Source: <https://developers.mercadolibre.com.ve/es_ar/items-atributos-de-envio-y-dimensiones>

## Content

Última actualización 05/05/2026

FAQs

›

Items - Atributos de envío y dimensiones

## Items - Atributos de envío y dimensiones

 ¿Qué atributos debo enviar para que la API acepte dimensiones y peso del paquete (seller_package_*) y en qué unidades? 

Las dimensiones deben enviarse como valores numéricos con unidades: height/width/length en centímetros y weight en gramos (entero). En shipping.dimensions suele utilizarse el formato AxBxC, peso, con dimensiones en cm y peso en gramos. La API valida límites y formatos; en algunos sites enviar el peso en gramos entero evita errores por fracciones.

 Recomendación 

 Normalice unidades en su integración (cm y g), convierta fracciones a enteros cuando corresponda y valide contra los límites de la API antes de enviar. 

 ¿Por qué recibo error "seller_package dimensions are too small / invalid" si envío dimensiones reales muy pequeñas (p. ej. delineadores)? 

El sistema valida el realismo de las dimensiones y puede rechazar medidas extremadamente pequeñas; la validación espera dimensiones de embalaje (caja o sobre) y no medidas del producto crudo si la categoría exige mínimos.

 Recomendación 

 Envíe las dimensiones del embalaje real utilizado (caja o sobre) y ajuste a mínimos de categoría cuando sea requerido. 

 Al actualizar un ítem con ME2 me aparecen warnings y las dimensiones no cambian. ¿Por qué? 

Para logística ME2/fulfillment, ciertas dimensiones son gestionadas por la operación logística y no siempre son modificables vía API; la plataforma puede imponer valores o considerar algunas medidas como read-only hasta el ingreso físico en depósito.

 Recomendación 

 Si el ítem está en modo fulfillment, considere que algunas dimensiones serán gestionadas por la logística y planifique actualizaciones coordinadas con el flujo operativo.
