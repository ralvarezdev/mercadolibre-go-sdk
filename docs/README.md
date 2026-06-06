# MercadoLibre API — Documentation Index

This directory documents the **MercadoLibre (marketplace) developer APIs**, as published at
<https://developers.mercadolibre.com.ve/es_ar/api-docs-es>. It is the reference base for the
Go SDK in this repository.

- **API base URL:** `https://api.mercadolibre.com`
- **Docs base URL:** `https://developers.mercadolibre.com.ve/es_ar/<slug>`
- **Source language:** Spanish (es_AR). Page titles below keep the original Spanish, with an
  English gloss where useful.

Each documented page lives under `docs/<section>/<slug>.md` and captures: overview, endpoints
(`METHOD path`), request parameters, request/response JSON models, and notes/caveats.

> Status legend: ⬜ not yet captured · 🟦 raw captured · ✅ curated for SDK
>
> **Raw capture complete (2026-06-06):** all 193 pages below have been fetched into
> `docs/<section>/<slug>.md` via `tools/docgen`. The ⬜ markers track the next phase —
> curating each page into SDK-ready endpoint + model specs.

---

## 1. Primeros pasos — Getting Started

| Page | Slug | Status |
|------|------|--------|
| Crea una aplicación en Mercado Libre | `crea-una-aplicacion-en-mercado-libre-es` | ⬜ |
| Permisos funcionales | `permisos-funcionales` | ⬜ |
| Gestiona tus aplicaciones | `gestiona-tus-aplicaciones` | ⬜ |
| Autenticación y Autorización (OAuth 2.0) | `autenticacion-y-autorizacion` | ⬜ |
| Realiza pruebas (test users) | `realiza-pruebas` | ⬜ |
| Validador de publicaciones | `validador-de-publicaciones` | ⬜ |
| Buenas prácticas para uso de la plataforma | `buenas-practicas-para-uso-de-la-plataforma` | ⬜ |
| Consideraciones de diseño | `consideraciones-de-diseno` | ⬜ |
| Gestionar IPs de una aplicación | `gestionar-ips-de-una-aplicacion` | ⬜ |
| Error 403 | `error-403` | ⬜ |

## 2. Usuarios — Users

| Page | Slug | Status |
|------|------|--------|
| Usuarios y Aplicaciones | `usuarios-y-aplicaciones` | ⬜ |
| Bloqueo de aplicaciones | `bloqueo-de-aplicaciones` | ⬜ |
| Consulta usuarios | `producto-consulta-usuarios` | ⬜ |
| Direcciones del usuario | `direcciones-del-usuario` | ⬜ |
| Preguntas frecuentes sobre validación de datos | `validacion-de-datos` | ⬜ |
| Validar datos de vendedores | `validar-datos-de-vendedores` | ⬜ |
| Favoritos (marcadores) | `marcadores` | ⬜ |

## 3. Dominios, Categorías, Ubicación y Monedas

| Page | Slug | Status |
|------|------|--------|
| Dominios y Categorías | `dominios-y-categorias` | ⬜ |
| Ubicación y Monedas | `ubicacion-y-monedas` | ⬜ |

## 4. Ítems y Atributos — Items

| Page | Slug | Status |
|------|------|--------|
| Ítems y Búsquedas | `items-y-busquedas` | ⬜ |
| Atributos | `atributos` | ⬜ |

## 5. Preguntas y Respuestas — Questions & Answers

| Page | Slug | Status |
|------|------|--------|
| Preguntas y Respuestas | `preguntas-y-respuestas` | ⬜ |
| Gestiona preguntas y respuestas | `gestiona-preguntas-respuestas` | ⬜ |

## 6. Pedidos y opiniones

| Page | Slug | Status |
|------|------|--------|
| Pedidos y opiniones | `pedidos-y-opiniones` | ⬜ |

## 7. Notificaciones y Comunicaciones

| Page | Slug | Status |
|------|------|--------|
| Notificaciones (webhooks / feeds) | `productos-recibe-notificaciones` | ⬜ |
| Comunicaciones | `conoce-las-novedades-que-reciben-los-vendedores` | ⬜ |

## 8. Moderaciones

| Page | Slug | Status |
|------|------|--------|
| Gestionar moderaciones | `gestionar-moderaciones` | ⬜ |
| Moderaciones con pausado | `moderaciones-con-pausado` | ⬜ |
| Diagnóstico de imágenes | `diagnostico-imagenes` | ⬜ |
| Moderaciones de imágenes | `moderaciones-de-imagenes` | ⬜ |

## 9. Brand Protection Program

| Page | Slug | Status |
|------|------|--------|
| ¿Qué es Brand Protection Program? | `que-es-brand-protection-program` | ⬜ |
| Miembros del Programa | `miembros-del-programa` | ⬜ |
| Publicaciones denunciadas | `publicaciones-denunciadas` | ⬜ |

## 10. Developer Partner Program

| Page | Slug | Status |
|------|------|--------|
| Developer Partner Program | `developer-partner-program` | ⬜ |

## 11. Seguridad — Desarrollo Seguro

| Page | Slug | Status |
|------|------|--------|
| Seguridad - Desarrollo Seguro | `seguridad-desarrollo-seguro` | ⬜ |
| Introducción | `introduccion-seguridad` | ⬜ |
| Infraestructura (Cifrado y TLS) | `infraestructura` | ⬜ |
| Gestión de Identidades y Accesos (OAuth y Tokens) | `gestion-de-identidades-y-accesos-oauth-y-tokens` | ⬜ |
| Autenticación segura | `autenticacion-segura` | ⬜ |
| Control de acceso y autorización | `control-de-acceso-y-autorizacion` | ⬜ |
| Seguridad de aplicaciones | `seguridad-apps` | ⬜ |
| Monitoreo | `monitoreo` | ⬜ |
| Gestión de incidentes | `gestion-de-incidentes` | ⬜ |

## 12. Guía para productos — Catalog & Publishing

### 12.1 General
| Page | Slug | Status |
|------|------|--------|
| Introducción | `guia-para-producto` | ⬜ |
| Tipos de publicación | `tipos-de-publicacion-y-actualizaciones-de-articulos` | ⬜ |

### 12.2 Categorización
| Page | Slug | Status |
|------|------|--------|
| Categorización de productos | `categoriza-productos` | ⬜ |
| Compatibilidades entre ítems y productos (Autopartes) | `compatibilidades-entre-items-y-productos` | ⬜ |
| Referencias de dominios, productos y atributos (Autopartes) | `referencias-de-dominios-productos-y-atributos-para-autopartes` | ⬜ |
| Publicar productos | `publica-productos` | ⬜ |

### 12.3 User Products
| Page | Slug | Status |
|------|------|--------|
| User Products | `user-products` | ⬜ |
| Precio por variación | `precio-variacion` | ⬜ |
| Stock distribuido | `stock-distribuido` | ⬜ |
| Stock multi origen | `stock-multi-origen` | ⬜ |
| Descripción de productos | `descripcion-de-articulos` | ⬜ |
| Validaciones | `validaciones` | ⬜ |

### 12.4 Precios y Costos
| Page | Slug | Status |
|------|------|--------|
| Precios de productos | `api-de-precios` | ⬜ |
| Precio por cantidad | `precio-por-cantidad` | ⬜ |
| Costos por vender | `comision-por-vender` | ⬜ |
| Referencias de precios | `referencias-de-precios` | ⬜ |
| Automatizaciones de precios | `automatizaciones-de-precios` | ⬜ |
| Precios netos por cantidad | `precios-netos` | ⬜ |

### 12.5 Imágenes, Identificadores y Variaciones
| Page | Slug | Status |
|------|------|--------|
| Imágenes | `trabajar-con-imagenes` | ⬜ |
| Identificadores de productos | `identificadores-de-productos` | ⬜ |
| Variaciones | `variaciones` | ⬜ |
| Republicar ítems | `re-publica` | ⬜ |
| Kits virtuales | `kits-virtuales` | ⬜ |

### 12.6 Guías de Talles
| Page | Slug | Status |
|------|------|--------|
| Primeros pasos | `primeros-pasos-es` | ⬜ |
| Gestionar guía de talles | `guias-de-talles` | ⬜ |
| Validación de guía de talles | `validacion-de-guia-de-talles` | ⬜ |
| Sincroniza y modifica publicaciones | `producto-sincroniza-modifica-publicaciones` | ⬜ |
| Convivencia Full/Flex (MLA y MLC) | `convivencia-full-y-flex` | ⬜ |

## 13. Mercado Envíos — Shipping

### 13.1 General
| Page | Slug | Status |
|------|------|--------|
| Gestión Mercado Envíos | `mercado-envios` | ⬜ |

### 13.2 Mercado Envíos 1 (ME1)
| Page | Slug | Status |
|------|------|--------|
| Mercado Envíos 1 | `mercadoenvios-modo-1` | ⬜ |
| Estados de órdenes y seguimiento | `estados-de-ordenes-me1` | ⬜ |
| Flete dinámico | `flete-dinamico` | ⬜ |

### 13.3 Mercado Envíos 2 (ME2)
| Page | Slug | Status |
|------|------|--------|
| Mercado Envíos 2 | `mercadoenvios-modo-2` | ⬜ |
| Costos de envío | `costos-de-envios` | ⬜ |
| Envíos en feriados opcionales | `dias-no-laborables` | ⬜ |
| Envíos Colecta y Places | `envios-colectas-places` | ⬜ |
| Agrupación de paquetes para la Colecta | `agrupacion-de-paquetes-para-la-colecta` | ⬜ |
| Envíos Flex | `envios-flex` | ⬜ |
| Envíos Turbo | `envios-turbo` | ⬜ |
| Envíos Fulfillment | `envios-fulfillment` | ⬜ |
| Envíos Personalizados | `envios-personalizados` | ⬜ |

## 14. Central de promociones — Promotions

| Page | Slug | Status |
|------|------|--------|
| Gestionar promociones | `central-de-promociones` | ⬜ |
| Campañas tradicionales | `deals` | ⬜ |
| Campañas co-fondeadas | `campanas-co-fondeadas` | ⬜ |
| Campañas con descuento por cantidad | `campanas-con-descuento-por-cantidad` | ⬜ |
| Pre-acordado por ítem y liquidación stock Full | `descuento-pre-acordado-por-item` | ⬜ |
| Descuento individual | `descuento-individual` | ⬜ |
| Ofertas del día | `ofertas-del-dia` | ⬜ |
| Ofertas relámpago | `ofertas-relampago` | ⬜ |
| Campañas del vendedor | `campanas-del-vendedor` | ⬜ |
| Co-fondeada automatizada y precios competitivos | `campanas-smart-price-matching` | ⬜ |
| Cupones del vendedor | `cupones-del-vendedor` | ⬜ |
| Campaña co-fondeada para PIX | `pix` | ⬜ |

## 15. Gestionar ventas — Orders

| Page | Slug | Status |
|------|------|--------|
| Órdenes | `gestiona-ventas` | ⬜ |
| Packs | `gestion-packs` | ⬜ |
| Envíos | `envios` | ⬜ |
| Notas de Packs | `notas-de-packs` | ⬜ |
| Pagos | `pagos` | ⬜ |
| Feedback de una venta | `feedback-sobre-venta` | ⬜ |
| Notas en órdenes | `notas-en-ordenes` | ⬜ |

## 16. Facturación — Billing

### 16.1 General
| Page | Slug | Status |
|------|------|--------|
| Datos de Facturación | `facturacion` | ⬜ |
| Cargar y Obtener Facturas - Emisión Propia | `cargar-factura` | ⬜ |

### 16.2 Reportes de Facturación
| Page | Slug | Status |
|------|------|--------|
| Buenas Prácticas para el Consumo de las APIs de Reportes | `buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion` | ⬜ |
| Reportes de Facturación | `reportes-de-facturacion` | ⬜ |
| Provisiones | `provisiones` | ⬜ |
| Pagos | `reportes-pagos` | ⬜ |
| Descargas | `reportes-descargas` | ⬜ |
| Percepciones | `resumen-percepciones` | ⬜ |

## 17. Mensajería posventa — Messaging

| Page | Slug | Status |
|------|------|--------|
| Qué es mensajería | `que-es-mensajeria` | ⬜ |
| Motivos para comunicarse | `motivos-para-comunicarse` | ⬜ |
| Gestión de mensajes | `mensajeria-post-venta` | ⬜ |
| Mensajes pendientes | `mensajes-pendientes` | ⬜ |
| Mensajes bloqueados | `mensajes-bloqueados` | ⬜ |

## 18. Reclamos, Devoluciones y Cambios — Claims

| Page | Slug | Status |
|------|------|--------|
| Gestionar reclamos | `que-es-un-reclamo` | ⬜ |
| Gestionar mensajes de un reclamo | `gestionar-mensaje-de-un-reclamo` | ⬜ |
| Gestionar resolución de reclamos | `gestionar-resolucion-de-reclamos` | ⬜ |
| Gestionar evidencia de reclamos | `gestionar-evidencia-de-reclamos` | ⬜ |
| Errores | `errores` | ⬜ |
| Devoluciones | `gestionar-devoluciones` | ⬜ |
| Cambios | `changes` | ⬜ |

## 19. Métricas y Tendencias — Metrics & Trends

| Page | Slug | Status |
|------|------|--------|
| Programa de Despegue y Beneficio de Reputación | `recuperacion-reputacion` | ⬜ |
| Reputación de vendedores | `reputacion-de-vendedores` | ⬜ |
| Tendencias | `tendencias` | ⬜ |
| Más vendidos en Mercado Libre | `mas-vendidos-en-mercado-libre` | ⬜ |
| Opiniones de productos | `opiniones-sobre-producto` | ⬜ |
| Calidad de publicaciones | `calidad-de-publicaciones` | ⬜ |
| Experiencia de compra | `experiencia-de-compra` | ⬜ |
| Visitas | `recurso-de-visitas` | ⬜ |
| Carga de atributos | `conoce-como-estan-los-vendedores-frente-la-carga-de-atributos` | ⬜ |

## 20. MercadoLíder y Tiendas Oficiales

| Page | Slug | Status |
|------|------|--------|
| Tiendas Oficiales | `tienda-oficial` | ⬜ |

---

## Verticals (vertical-specific catalogs)

> These are domain verticals layered on the same marketplace platform. Listed for completeness;
> deprioritized relative to the core marketplace resources above.

### 21. Guía para inmuebles — Real Estate
`introduccion-guia-de-inmuebles`, `primeros-pasos-inmuebles`, `configuracion-o-requisitos-previos`,
`obtencion-del-access-token`, `consulta-de-usuarios`, `pasos-rapidos-para-publicar-un-inmueble-de-prueba`,
`categorias-atributos-inmuebles`, `categorias-inmuebles`, `atributos-inmuebles`, `localizar-inmuebles`,
`gestionar-paquetes-de-inmuebles`, `contratacion-de-paquetes-de-publicacion`,
`paquetes-y-permisos-para-proyectos-desarrollos-o-emprendimientos-inmobiliarios`, `publica-inmueble`,
`publicaciones-de-tiendas-oficiales-para-inmuebles`, `actualiza-tus-publicaciones`,
`ciclo-de-vida-de-las-publicaciones-de-inmuebles`, `variaciones-para-inmuebles`,
`actualizacion-variacion-inmuebles`, `calidad-de-las-publicaciones-inmuebles`, `desarrollos-inmobiliarios`,
`leads-inmuebles`, `solicitud-de-visita`, `estadisticas-de-interacciones-en-inmuebles`,
`experiencia-para-inmuebles`, `glosario-inmuebles`

### 22. Guía para vehículos — Vehicles
`introduccion-vehiculos`, `consulta-usuarios`, `categorias-y-atributos`, `localizacion-de-vehiculos`,
`vehiculos-gestiona-paquetes`, `publica-vehiculos`, `vehiculos-sincroniza-publicaciones`,
`vehiculos-gestiona-preguntas-y-contactos`, `persona-interesadas`, `credits-motors`,
`calidad-de-publicaciones-vehiculos`

### 23. Guía para servicios — Services
`guia-para-servicios`, `servicios-consulta-usuarios`, `elige-tipo-de-servicio`,
`administra-areas-de-cobertura`, `publica-servicios-vis`, `servicio-sincroniza-publicaciones`,
`consultas-avanzadas-2`

### 24. Guía para Mercado Ads
`introduccion-a-mercado-ads`, `pads-read`, `ads-bads`, `display`,
`product-ads-para-catalogo-y-user-products-lectura`, `bonificaciones-para-product-ads`

### 25. MCP Mercado Libre
`mcp-server`

### 26. FAQs
`faq-preguntas-frecuentes`, `rate-limit-error-429`, `stock-multiwarehouse`, `me1-me2-y-envio-gratis`,
`mercado-envios-costos-y-cotizaciones`, `facturacion-billing-info`, `imagenes-y-moderaciones`,
`promotions-pricing`, `items-atributos-de-envio-y-dimensiones`

---

**Total: 193 documentation pages.** This index is generated from the live sidebar navigation of the
MercadoLibre developer site (captured 2026-06-06).
