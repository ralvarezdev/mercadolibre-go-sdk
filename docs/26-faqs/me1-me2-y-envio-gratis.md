---
title: ME1 / ME2 y envío gratis
slug: me1-me2-y-envio-gratis
section: 26-faqs
source: https://developers.mercadolibre.com.ve/es_ar/me1-me2-y-envio-gratis
captured: 2026-06-06
---

# ME1 / ME2 y envío gratis

> Source: <https://developers.mercadolibre.com.ve/es_ar/me1-me2-y-envio-gratis>

## Content

Última actualización 05/05/2026

FAQs

›

ME1 / ME2 y envío gratis

## ME1 / ME2 y envío gratis

 ¿Por qué setear shipping.mode = "me1" vía API no muestra envío gratis en el frontend aunque free_shipping:true esté en el item? 

En ME1 el flag free_shipping:true puede existir en los datos, pero Mercado Libre no lo interpreta automáticamente como "envío gratis" en el frontend ni le aplica los beneficios automáticos de ME2. La visibilidad y beneficios de "Envío gratis" están diseñados para ME2; para que el frontend muestre el envío gratis con los beneficios automáticos, el ítem debe estar en modo me2.

 Recomendación 

 Si necesita la exhibición y beneficios automáticos de envío gratis, verifique que el ítem esté en modo ME2 y cumpla los requisitos de elegibilidad del mercado. 

 ¿Por qué al intentar cambiar un ítem de ME2 a ME1 la API lo deja como not_specified? 

Si la cuenta o las preferencias de usuario tienen coexistencia priorizada a ME2 (por ejemplo, tags como me2_available o preferencias), la plataforma puede impedir la conversión a ME1 y devolver not_specified. La razón suele ser la configuración de la cuenta y las preferencias comerciales; la conversión a ME1 puede requerir intervención comercial (KAM) para ajustar las preferencias de envío.

 Recomendación 

 Revise las preferencias comerciales y tags del seller; si la conversión falla, coordine con el equipo comercial para ajustar la configuración de convivencia. 

 ¿Es obligatorio estar en ME2 para que se activen automáticamente opciones como "envío gratis a todo el país"? 

Sí: muchas de las reglas automáticas de muestra y beneficios (por ejemplo, "Envío gratis" con beneficios) se aplican preferentemente a ítems en modo ME2 y dependen además del umbral de precio y otras políticas internas. No es posible forzar vía API el beneficio global si no se cumplen esas condiciones.

 Recomendación 

 Compruebe elegibilidad para ME2 y los umbrales requeridos; no intente forzar beneficios por API sin cumplir las condiciones del mercado. 

 ¿Se puede forzar por API el envío gratis a todo el país? 

No. La activación de envío gratis a nivel nacional depende de reglas automáticas (precio, tipo de logística como ME2) y no se puede forzar vía API sin cumplir esos requisitos.

 Recomendación 

 Ajuste precios y logística para cumplir las condiciones del mercado; no espere poder habilitar envío gratis nacional únicamente por API.
