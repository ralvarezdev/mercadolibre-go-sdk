package main

// Page is one documentation page on the MercadoLibre developer site.
type Page struct {
	Section string // output sub-directory under docs/
	Slug    string // URL slug under https://developers.mercadolibre.com.ve/es_ar/
	Title   string // human title (Spanish)
}

// sitemap is the curated table of contents of the MercadoLibre (marketplace)
// developer documentation. Captured from the live sidebar navigation on
// 2026-06-06. Sections map to docs/<section>/ folders.
var sitemap = []Page{
	// 1. Primeros pasos
	{"01-getting-started", "crea-una-aplicacion-en-mercado-libre-es", "Crea una aplicación en Mercado Libre"},
	{"01-getting-started", "permisos-funcionales", "Permisos funcionales"},
	{"01-getting-started", "gestiona-tus-aplicaciones", "Gestiona tus aplicaciones"},
	{"01-getting-started", "autenticacion-y-autorizacion", "Autenticación y Autorización"},
	{"01-getting-started", "realiza-pruebas", "Realiza pruebas"},
	{"01-getting-started", "validador-de-publicaciones", "Validador de publicaciones"},
	{"01-getting-started", "buenas-practicas-para-uso-de-la-plataforma", "Buenas prácticas para uso de la plataforma"},
	{"01-getting-started", "consideraciones-de-diseno", "Consideraciones de diseño"},
	{"01-getting-started", "gestionar-ips-de-una-aplicacion", "Gestionar IPs de una aplicación"},
	{"01-getting-started", "error-403", "Error 403"},

	// 2. Usuarios
	{"02-users", "usuarios-y-aplicaciones", "Usuarios y Aplicaciones"},
	{"02-users", "bloqueo-de-aplicaciones", "Bloqueo de aplicaciones"},
	{"02-users", "producto-consulta-usuarios", "Consulta usuarios"},
	{"02-users", "direcciones-del-usuario", "Direcciones del usuario"},
	{"02-users", "validacion-de-datos", "Preguntas frecuentes sobre validación de datos"},
	{"02-users", "validar-datos-de-vendedores", "Validar datos de vendedores"},
	{"02-users", "marcadores", "Favoritos"},

	// 3. Dominios, Categorías, Ubicación
	{"03-domains-categories", "dominios-y-categorias", "Dominios y Categorías"},
	{"03-domains-categories", "ubicacion-y-monedas", "Ubicación y Monedas"},

	// 4. Ítems y Atributos
	{"04-items", "items-y-busquedas", "Ítems y Búsquedas"},
	{"04-items", "atributos", "Atributos"},

	// 5. Preguntas y Respuestas
	{"05-questions", "preguntas-y-respuestas", "Preguntas y Respuestas"},
	{"05-questions", "gestiona-preguntas-respuestas", "Gestiona preguntas y respuestas"},

	// 6. Pedidos y opiniones
	{"06-orders-feedback", "pedidos-y-opiniones", "Pedidos y opiniones"},

	// 7. Notificaciones y Comunicaciones
	{"07-notifications", "productos-recibe-notificaciones", "Notificaciones"},
	{"07-notifications", "conoce-las-novedades-que-reciben-los-vendedores", "Comunicaciones"},

	// 8. Moderaciones
	{"08-moderations", "gestionar-moderaciones", "Gestionar moderaciones"},
	{"08-moderations", "moderaciones-con-pausado", "Moderaciones con pausado"},
	{"08-moderations", "diagnostico-imagenes", "Diagnóstico de imágenes"},
	{"08-moderations", "moderaciones-de-imagenes", "Moderaciones de imágenes"},

	// 9. Brand Protection Program
	{"09-brand-protection", "que-es-brand-protection-program", "¿Qué es Brand Protection Program?"},
	{"09-brand-protection", "miembros-del-programa", "Miembros del Programa"},
	{"09-brand-protection", "publicaciones-denunciadas", "Publicaciones denunciadas"},

	// 10. Developer Partner Program
	{"10-partner-program", "developer-partner-program", "Developer Partner Program"},

	// 11. Seguridad
	{"11-security", "seguridad-desarrollo-seguro", "Seguridad - Desarrollo Seguro"},
	{"11-security", "introduccion-seguridad", "Introducción"},
	{"11-security", "infraestructura", "Infraestructura (Cifrado y TLS)"},
	{"11-security", "gestion-de-identidades-y-accesos-oauth-y-tokens", "Gestión de Identidades y Accesos (OAuth y Tokens)"},
	{"11-security", "autenticacion-segura", "Autenticación segura"},
	{"11-security", "control-de-acceso-y-autorizacion", "Control de acceso y autorización"},
	{"11-security", "seguridad-apps", "Seguridad de aplicaciones"},
	{"11-security", "monitoreo", "Monitoreo"},
	{"11-security", "gestion-de-incidentes", "Gestión de incidentes"},

	// 12. Guía para productos
	{"12-products", "guia-para-producto", "Introducción"},
	{"12-products", "tipos-de-publicacion-y-actualizaciones-de-articulos", "Tipos de publicación"},
	{"12-products", "categoriza-productos", "Categorización de productos"},
	{"12-products", "compatibilidades-entre-items-y-productos", "Compatibilidades entre ítems y productos de Autopartes"},
	{"12-products", "referencias-de-dominios-productos-y-atributos-para-autopartes", "Referencias de dominios, productos y atributos para Autopartes"},
	{"12-products", "publica-productos", "Publicar productos"},
	{"12-products", "user-products", "User Products"},
	{"12-products", "precio-variacion", "Precio por variación"},
	{"12-products", "stock-distribuido", "Stock distribuido"},
	{"12-products", "stock-multi-origen", "Stock multi origen"},
	{"12-products", "descripcion-de-articulos", "Descripción de productos"},
	{"12-products", "validaciones", "Validaciones"},
	{"12-products", "api-de-precios", "Precios de productos"},
	{"12-products", "precio-por-cantidad", "Precio por cantidad"},
	{"12-products", "comision-por-vender", "Costos por vender"},
	{"12-products", "referencias-de-precios", "Referencias de precios"},
	{"12-products", "automatizaciones-de-precios", "Automatizaciones de precios"},
	{"12-products", "precios-netos", "Precios netos por cantidad"},
	{"12-products", "trabajar-con-imagenes", "Imágenes"},
	{"12-products", "identificadores-de-productos", "Identificadores de productos"},
	{"12-products", "variaciones", "Variaciones"},
	{"12-products", "re-publica", "Republicar ítems"},
	{"12-products", "kits-virtuales", "Kits virtuales"},
	{"12-products", "primeros-pasos-es", "Guías de Talles - Primeros pasos"},
	{"12-products", "guias-de-talles", "Gestionar guía de talles"},
	{"12-products", "validacion-de-guia-de-talles", "Validación de guía de talles"},
	{"12-products", "producto-sincroniza-modifica-publicaciones", "Sincroniza y modifica publicaciones"},
	{"12-products", "convivencia-full-y-flex", "Convivencia Full/Flex (MLA y MLC)"},

	// 13. Mercado Envíos
	{"13-shipping", "mercado-envios", "Gestión Mercado Envíos"},
	{"13-shipping", "mercadoenvios-modo-1", "Mercado Envíos 1"},
	{"13-shipping", "estados-de-ordenes-me1", "Estados de órdenes y seguimiento"},
	{"13-shipping", "flete-dinamico", "Flete dinámico"},
	{"13-shipping", "mercadoenvios-modo-2", "Mercado Envíos 2"},
	{"13-shipping", "costos-de-envios", "Costos de envío"},
	{"13-shipping", "dias-no-laborables", "Envíos en feriados opcionales"},
	{"13-shipping", "envios-colectas-places", "Envíos Colecta y Places"},
	{"13-shipping", "agrupacion-de-paquetes-para-la-colecta", "Agrupación de paquetes para la Colecta"},
	{"13-shipping", "envios-flex", "Envíos Flex"},
	{"13-shipping", "envios-turbo", "Envíos Turbo"},
	{"13-shipping", "envios-fulfillment", "Envíos Fulfillment"},
	{"13-shipping", "envios-personalizados", "Envíos Personalizados"},

	// 14. Central de promociones
	{"14-promotions", "central-de-promociones", "Gestionar promociones"},
	{"14-promotions", "deals", "Campañas tradicionales"},
	{"14-promotions", "campanas-co-fondeadas", "Campañas co-fondeadas"},
	{"14-promotions", "campanas-con-descuento-por-cantidad", "Campañas con descuento por cantidad"},
	{"14-promotions", "descuento-pre-acordado-por-item", "Pre-acordado por ítem y liquidación stock Full"},
	{"14-promotions", "descuento-individual", "Descuento individual"},
	{"14-promotions", "ofertas-del-dia", "Ofertas del día"},
	{"14-promotions", "ofertas-relampago", "Ofertas relámpago"},
	{"14-promotions", "campanas-del-vendedor", "Campañas del vendedor"},
	{"14-promotions", "campanas-smart-price-matching", "Co-fondeada automatizada y precios competitivos"},
	{"14-promotions", "cupones-del-vendedor", "Cupones del vendedor"},
	{"14-promotions", "pix", "Campaña co-fondeada para PIX"},

	// 15. Gestionar ventas
	{"15-sales-orders", "gestiona-ventas", "Órdenes"},
	{"15-sales-orders", "gestion-packs", "Packs"},
	{"15-sales-orders", "envios", "Envíos"},
	{"15-sales-orders", "notas-de-packs", "Notas de Packs"},
	{"15-sales-orders", "pagos", "Pagos"},
	{"15-sales-orders", "feedback-sobre-venta", "Feedback de una venta"},
	{"15-sales-orders", "notas-en-ordenes", "Notas en órdenes"},

	// 16. Facturación
	{"16-billing", "facturacion", "Datos de Facturación"},
	{"16-billing", "cargar-factura", "Cargar y Obtener Facturas - Emisión Propia"},
	{"16-billing", "buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion", "Buenas Prácticas para el Consumo de las APIs de Reportes de Facturación"},
	{"16-billing", "reportes-de-facturacion", "Reportes de Facturación"},
	{"16-billing", "provisiones", "Provisiones"},
	{"16-billing", "reportes-pagos", "Pagos"},
	{"16-billing", "reportes-descargas", "Descargas"},
	{"16-billing", "resumen-percepciones", "Percepciones"},

	// 17. Mensajería posventa
	{"17-messaging", "que-es-mensajeria", "Qué es mensajería"},
	{"17-messaging", "motivos-para-comunicarse", "Motivos para comunicarse"},
	{"17-messaging", "mensajeria-post-venta", "Gestión de mensajes"},
	{"17-messaging", "mensajes-pendientes", "Mensajes pendientes"},
	{"17-messaging", "mensajes-bloqueados", "Mensajes bloqueados"},

	// 18. Reclamos, Devoluciones y Cambios
	{"18-claims", "que-es-un-reclamo", "Gestionar reclamos"},
	{"18-claims", "gestionar-mensaje-de-un-reclamo", "Gestionar mensajes de un reclamo"},
	{"18-claims", "gestionar-resolucion-de-reclamos", "Gestionar resolución de reclamos"},
	{"18-claims", "gestionar-evidencia-de-reclamos", "Gestionar evidencia de reclamos"},
	{"18-claims", "errores", "Errores"},
	{"18-claims", "gestionar-devoluciones", "Devoluciones"},
	{"18-claims", "changes", "Cambios"},

	// 19. Métricas y Tendencias
	{"19-metrics", "recuperacion-reputacion", "Programa de Despegue y Beneficio de Reputación"},
	{"19-metrics", "reputacion-de-vendedores", "Reputación de vendedores"},
	{"19-metrics", "tendencias", "Tendencias"},
	{"19-metrics", "mas-vendidos-en-mercado-libre", "Más vendidos en Mercado Libre"},
	{"19-metrics", "opiniones-sobre-producto", "Opiniones de productos"},
	{"19-metrics", "calidad-de-publicaciones", "Calidad de publicaciones"},
	{"19-metrics", "experiencia-de-compra", "Experiencia de compra"},
	{"19-metrics", "recurso-de-visitas", "Visitas"},
	{"19-metrics", "conoce-como-estan-los-vendedores-frente-la-carga-de-atributos", "Carga de atributos"},

	// 20. MercadoLíder y Tiendas Oficiales
	{"20-official-stores", "tienda-oficial", "Tiendas Oficiales"},

	// 21. Guía para inmuebles (vertical)
	{"21-real-estate", "introduccion-guia-de-inmuebles", "Introducción guía de inmuebles"},
	{"21-real-estate", "primeros-pasos-inmuebles", "Primeros pasos"},
	{"21-real-estate", "configuracion-o-requisitos-previos", "Configuración o requisitos previos"},
	{"21-real-estate", "obtencion-del-access-token", "Obtención del Access Token"},
	{"21-real-estate", "consulta-de-usuarios", "Consulta de usuarios"},
	{"21-real-estate", "pasos-rapidos-para-publicar-un-inmueble-de-prueba", "Pasos rápidos para publicar un inmueble de prueba"},
	{"21-real-estate", "categorias-atributos-inmuebles", "Categorías y atributos"},
	{"21-real-estate", "categorias-inmuebles", "Categorías"},
	{"21-real-estate", "atributos-inmuebles", "Atributos"},
	{"21-real-estate", "localizar-inmuebles", "Localizar Inmuebles"},
	{"21-real-estate", "gestionar-paquetes-de-inmuebles", "Gestionar paquetes de inmuebles"},
	{"21-real-estate", "contratacion-de-paquetes-de-publicacion", "Contratación de paquetes de publicación"},
	{"21-real-estate", "paquetes-y-permisos-para-proyectos-desarrollos-o-emprendimientos-inmobiliarios", "Paquetes y permisos para proyectos inmobiliarios"},
	{"21-real-estate", "publica-inmueble", "Publica Inmuebles"},
	{"21-real-estate", "publicaciones-de-tiendas-oficiales-para-inmuebles", "Publicaciones de Tiendas Oficiales para Inmuebles"},
	{"21-real-estate", "actualiza-tus-publicaciones", "Actualiza tus publicaciones"},
	{"21-real-estate", "ciclo-de-vida-de-las-publicaciones-de-inmuebles", "Ciclo de vida de las publicaciones de Inmuebles"},
	{"21-real-estate", "variaciones-para-inmuebles", "Variaciones"},
	{"21-real-estate", "actualizacion-variacion-inmuebles", "Actualización variación inmuebles"},
	{"21-real-estate", "calidad-de-las-publicaciones-inmuebles", "Calidad de las Publicaciones"},
	{"21-real-estate", "desarrollos-inmobiliarios", "Desarrollos inmobiliarios"},
	{"21-real-estate", "leads-inmuebles", "Leads"},
	{"21-real-estate", "solicitud-de-visita", "Solicitud de visita"},
	{"21-real-estate", "estadisticas-de-interacciones-en-inmuebles", "Estadísticas de interacciones en Inmuebles"},
	{"21-real-estate", "experiencia-para-inmuebles", "Experiencia para inmuebles"},
	{"21-real-estate", "glosario-inmuebles", "Glosario"},

	// 22. Guía para vehículos (vertical)
	{"22-vehicles", "introduccion-vehiculos", "Introducción"},
	{"22-vehicles", "consulta-usuarios", "Consulta Usuarios"},
	{"22-vehicles", "categorias-y-atributos", "Categorías y Atributos"},
	{"22-vehicles", "localizacion-de-vehiculos", "Localiza vehículos"},
	{"22-vehicles", "vehiculos-gestiona-paquetes", "Gestiona paquetes de vehículos"},
	{"22-vehicles", "publica-vehiculos", "Publica vehículos"},
	{"22-vehicles", "vehiculos-sincroniza-publicaciones", "Sincroniza publicaciones (vehículos)"},
	{"22-vehicles", "vehiculos-gestiona-preguntas-y-contactos", "Gestiona preguntas y contactos"},
	{"22-vehicles", "persona-interesadas", "Personas Interesadas"},
	{"22-vehicles", "credits-motors", "Créditos pre aprobados"},
	{"22-vehicles", "calidad-de-publicaciones-vehiculos", "Calidad de publicaciones (vehículos)"},

	// 23. Guía para servicios (vertical)
	{"23-services", "guia-para-servicios", "Introducción"},
	{"23-services", "servicios-consulta-usuarios", "Consulta Usuarios"},
	{"23-services", "elige-tipo-de-servicio", "Elige tipo de servicio"},
	{"23-services", "administra-areas-de-cobertura", "Administra áreas de cobertura"},
	{"23-services", "publica-servicios-vis", "Publica servicios"},
	{"23-services", "servicio-sincroniza-publicaciones", "Sincroniza publicaciones"},
	{"23-services", "consultas-avanzadas-2", "Consultas avanzadas"},

	// 24. Guía para Mercado Ads
	{"24-mercado-ads", "introduccion-a-mercado-ads", "Introducción a Mercado Ads"},
	{"24-mercado-ads", "pads-read", "Product Ads"},
	{"24-mercado-ads", "ads-bads", "Brand Ads"},
	{"24-mercado-ads", "display", "Display Ads"},
	{"24-mercado-ads", "product-ads-para-catalogo-y-user-products-lectura", "Product Ads para Catálogo y User Products"},
	{"24-mercado-ads", "bonificaciones-para-product-ads", "Bonificaciones para Product Ads"},

	// 25. MCP
	{"25-mcp", "mcp-server", "Mercado Libre MCP Server"},

	// 26. FAQs
	{"26-faqs", "faq-preguntas-frecuentes", "Preguntas Frecuentes"},
	{"26-faqs", "rate-limit-error-429", "Rate Limit / Error 429"},
	{"26-faqs", "stock-multiwarehouse", "Gestión de stock multiorigen / User Products"},
	{"26-faqs", "me1-me2-y-envio-gratis", "ME1 / ME2 y envío gratis"},
	{"26-faqs", "mercado-envios-costos-y-cotizaciones", "Mercado Envíos - Costos y cotizaciones"},
	{"26-faqs", "facturacion-billing-info", "Facturación / Billing info"},
	{"26-faqs", "imagenes-y-moderaciones", "Imágenes y moderaciones"},
	{"26-faqs", "promotions-pricing", "Promotions / Pricing"},
	{"26-faqs", "items-atributos-de-envio-y-dimensiones", "Items - Atributos de envío y dimensiones"},
}
