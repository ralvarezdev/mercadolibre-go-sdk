package main

// Page is one documentation page on the MercadoLibre developer site.
type Page struct {
	Section string // output sub-directory under docs/
	Slug    string // URL slug under https://developers.mercadolibre.com.ve/es_ar/
	Title   string // human title (Spanish)
}

const (
	s01GettingStarted = "01-getting-started"
	s02Users          = "02-users"
	s03Domains        = "03-domains-categories"
	s04Items          = "04-items"
	s05Questions      = "05-questions"
	s06Orders         = "06-orders-feedback"
	s07Notifications  = "07-notifications"
	s08Moderations    = "08-moderations"
	s09BrandProtect   = "09-brand-protection"
	s10Partner        = "10-partner-program"
	s11Security       = "11-security"
	s12Products       = "12-products"
	s13Shipping       = "13-shipping"
	s14Promotions     = "14-promotions"
	s15SalesOrders    = "15-sales-orders"
	s16Billing        = "16-billing"
	s17Messaging      = "17-messaging"
	s18Claims         = "18-claims"
	s19Metrics        = "19-metrics"
	s20OfficialStores = "20-official-stores"
	s21RealEstate     = "21-real-estate"
	s22Vehicles       = "22-vehicles"
	s23Services       = "23-services"
	s24MercadoAds     = "24-mercado-ads"
	s25MCP            = "25-mcp"
	s26FAQs           = "26-faqs"

	titleIntro = "Introducción"
)

// sitemap is the curated table of contents of the MercadoLibre (marketplace)
// developer documentation. Captured from the live sidebar navigation on
// 2026-06-06. Sections map to docs/<section>/ folders.
var sitemap = []Page{
	// 1. Primeros pasos
	{s01GettingStarted, "crea-una-aplicacion-en-mercado-libre-es", "Crea una aplicación en Mercado Libre"},
	{s01GettingStarted, "permisos-funcionales", "Permisos funcionales"},
	{s01GettingStarted, "gestiona-tus-aplicaciones", "Gestiona tus aplicaciones"},
	{s01GettingStarted, "autenticacion-y-autorizacion", "Autenticación y Autorización"},
	{s01GettingStarted, "realiza-pruebas", "Realiza pruebas"},
	{s01GettingStarted, "validador-de-publicaciones", "Validador de publicaciones"},
	{s01GettingStarted, "buenas-practicas-para-uso-de-la-plataforma", "Buenas prácticas para uso de la plataforma"},
	{s01GettingStarted, "consideraciones-de-diseno", "Consideraciones de diseño"},
	{s01GettingStarted, "gestionar-ips-de-una-aplicacion", "Gestionar IPs de una aplicación"},
	{s01GettingStarted, "error-403", "Error 403"},

	// 2. Usuarios
	{s02Users, "usuarios-y-aplicaciones", "Usuarios y Aplicaciones"},
	{s02Users, "bloqueo-de-aplicaciones", "Bloqueo de aplicaciones"},
	{s02Users, "producto-consulta-usuarios", "Consulta usuarios"},
	{s02Users, "direcciones-del-usuario", "Direcciones del usuario"},
	{s02Users, "validacion-de-datos", "Preguntas frecuentes sobre validación de datos"},
	{s02Users, "validar-datos-de-vendedores", "Validar datos de vendedores"},
	{s02Users, "marcadores", "Favoritos"},

	// 3. Dominios, Categorías, Ubicación
	{s03Domains, "dominios-y-categorias", "Dominios y Categorías"},
	{s03Domains, "ubicacion-y-monedas", "Ubicación y Monedas"},

	// 4. Ítems y Atributos
	{s04Items, "items-y-busquedas", "Ítems y Búsquedas"},
	{s04Items, "atributos", "Atributos"},

	// 5. Preguntas y Respuestas
	{s05Questions, "preguntas-y-respuestas", "Preguntas y Respuestas"},
	{s05Questions, "gestiona-preguntas-respuestas", "Gestiona preguntas y respuestas"},

	// 6. Pedidos y feedback
	{s06Orders, "pedidos-y-opiniones", "Pedidos y opiniones"},

	// 7. Notificaciones y Comunicaciones
	{s07Notifications, "productos-recibe-notificaciones", "Notificaciones"},
	{s07Notifications, "conoce-las-novedades-que-reciben-los-vendedores", "Comunicaciones"},

	// 8. Moderaciones
	{s08Moderations, "gestionar-moderaciones", "Gestionar moderaciones"},
	{s08Moderations, "moderaciones-con-pausado", "Moderaciones con pausado"},
	{s08Moderations, "diagnostico-imagenes", "Diagnóstico de imágenes"},
	{s08Moderations, "moderaciones-de-imagenes", "Moderaciones de imágenes"},

	// 9. Brand Protection Program
	{s09BrandProtect, "que-es-brand-protection-program", "¿Qué es Brand Protection Program?"},
	{s09BrandProtect, "miembros-del-programa", "Miembros del Programa"},
	{s09BrandProtect, "publicaciones-denunciadas", "Publicaciones denunciadas"},

	// 10. Developer Partner Program
	{s10Partner, "developer-partner-program", "Developer Partner Program"},

	// 11. Seguridad
	{s11Security, "seguridad-desarrollo-seguro", "Seguridad - Desarrollo Seguro"},
	{s11Security, "introduccion-seguridad", titleIntro},
	{s11Security, "infraestructura", "Infraestructura (Cifrado y TLS)"},
	{
		s11Security,
		"gestion-de-identidades-y-accesos-oauth-y-tokens",
		"Gestión de Identidades y Accesos (OAuth y Tokens)",
	},
	{s11Security, "autenticacion-segura", "Autenticación segura"},
	{s11Security, "control-de-acceso-y-autorizacion", "Control de acceso y autorización"},
	{s11Security, "seguridad-apps", "Seguridad de aplicaciones"},
	{s11Security, "monitoreo", "Monitoreo"},
	{s11Security, "gestion-de-incidentes", "Gestión de incidentes"},

	// 12. Guía para items
	{s12Products, "guia-para-producto", titleIntro},
	{s12Products, "tipos-de-publicacion-y-actualizaciones-de-articulos", "Tipos de publicación"},
	{s12Products, "categoriza-productos", "Categorización de productos"},
	{s12Products, "compatibilidades-entre-items-y-productos", "Compatibilidades entre ítems y productos de Autopartes"},
	{
		s12Products,
		"referencias-de-dominios-productos-y-atributos-para-autopartes",
		"Referencias de dominios, productos y atributos para Autopartes",
	},
	{s12Products, "publica-productos", "Publicar productos"},
	{s12Products, "user-products", "User Products"},
	{s12Products, "precio-variacion", "Precio por variación"},
	{s12Products, "stock-distribuido", "Stock distribuido"},
	{s12Products, "stock-multi-origen", "Stock multi origen"},
	{s12Products, "descripcion-de-articulos", "Descripción de productos"},
	{s12Products, "validaciones", "Validaciones"},
	{s12Products, "api-de-precios", "Precios de productos"},
	{s12Products, "precio-por-cantidad", "Precio por cantidad"},
	{s12Products, "comision-por-vender", "Costos por vender"},
	{s12Products, "referencias-de-precios", "Referencias de precios"},
	{s12Products, "automatizaciones-de-precios", "Automatizaciones de precios"},
	{s12Products, "precios-netos", "Precios netos por cantidad"},
	{s12Products, "trabajar-con-imagenes", "Imágenes"},
	{s12Products, "identificadores-de-productos", "Identificadores de productos"},
	{s12Products, "variaciones", "Variaciones"},
	{s12Products, "re-publica", "Republicar ítems"},
	{s12Products, "kits-virtuales", "Kits virtuales"},
	{s12Products, "primeros-pasos-es", "Guías de Talles - Primeros pasos"},
	{s12Products, "guias-de-talles", "Gestionar guía de talles"},
	{s12Products, "validacion-de-guia-de-talles", "Validación de guía de talles"},
	{s12Products, "producto-sincroniza-modifica-publicaciones", "Sincroniza y modifica publicaciones"},
	{s12Products, "convivencia-full-y-flex", "Convivencia Full/Flex (MLA y MLC)"},

	// 13. Mercado Envíos
	{s13Shipping, "mercado-envios", "Gestión Mercado Envíos"},
	{s13Shipping, "mercadoenvios-modo-1", "Mercado Envíos 1"},
	{s13Shipping, "estados-de-ordenes-me1", "Estados de órdenes y seguimiento"},
	{s13Shipping, "flete-dinamico", "Flete dinámico"},
	{s13Shipping, "mercadoenvios-modo-2", "Mercado Envíos 2"},
	{s13Shipping, "costos-de-envios", "Costos de envío"},
	{s13Shipping, "dias-no-laborables", "Envíos en feriados opcionales"},
	{s13Shipping, "envios-colectas-places", "Envíos Colecta y Places"},
	{s13Shipping, "agrupacion-de-paquetes-para-la-colecta", "Agrupación de paquetes para la Colecta"},
	{s13Shipping, "envios-flex", "Envíos Flex"},
	{s13Shipping, "envios-turbo", "Envíos Turbo"},
	{s13Shipping, "envios-fulfillment", "Envíos Fulfillment"},
	{s13Shipping, "envios-personalizados", "Envíos Personalizados"},

	// 14. Central de promociones
	{s14Promotions, "central-de-promociones", "Gestionar promociones"},
	{s14Promotions, "deals", "Campañas tradicionales"},
	{s14Promotions, "campanas-co-fondeadas", "Campañas co-fondeadas"},
	{s14Promotions, "campanas-con-descuento-por-cantidad", "Campañas con descuento por cantidad"},
	{s14Promotions, "descuento-pre-acordado-por-item", "Pre-acordado por ítem y liquidación stock Full"},
	{s14Promotions, "descuento-individual", "Descuento individual"},
	{s14Promotions, "ofertas-del-dia", "Ofertas del día"},
	{s14Promotions, "ofertas-relampago", "Ofertas relámpago"},
	{s14Promotions, "campanas-del-vendedor", "Campañas del vendedor"},
	{s14Promotions, "campanas-smart-price-matching", "Co-fondeada automatizada y precios competitivos"},
	{s14Promotions, "cupones-del-vendedor", "Cupones del vendedor"},
	{s14Promotions, "pix", "Campaña co-fondeada para PIX"},

	// 15. Gestionar ventas
	{s15SalesOrders, "gestiona-ventas", "Órdenes"},
	{s15SalesOrders, "gestion-packs", "Packs"},
	{s15SalesOrders, "envios", "Envíos"},
	{s15SalesOrders, "notas-de-packs", "Notas de Packs"},
	{s15SalesOrders, "pagos", "Pagos"},
	{s15SalesOrders, "feedback-sobre-venta", "Feedback de una venta"},
	{s15SalesOrders, "notas-en-ordenes", "Notas en órdenes"},

	// 16. Facturación
	{s16Billing, "facturacion", "Datos de Facturación"},
	{s16Billing, "cargar-factura", "Cargar y Obtener Facturas - Emisión Propia"},
	{
		s16Billing,
		"buenas-practicas-para-el-consumo-de-las-apis-de-reportes-de-facturacion",
		"Buenas Prácticas para el Consumo de las APIs de Reportes de Facturación",
	},
	{s16Billing, "reportes-de-facturacion", "Reportes de Facturación"},
	{s16Billing, "provisiones", "Provisiones"},
	{s16Billing, "reportes-pagos", "Pagos"},
	{s16Billing, "reportes-descargas", "Descargas"},
	{s16Billing, "resumen-percepciones", "Percepciones"},

	// 17. Mensajería posventa
	{s17Messaging, "que-es-mensajeria", "Qué es mensajería"},
	{s17Messaging, "motivos-para-comunicarse", "Motivos para comunicarse"},
	{s17Messaging, "mensajeria-post-venta", "Gestión de mensajes"},
	{s17Messaging, "mensajes-pendientes", "Mensajes pendientes"},
	{s17Messaging, "mensajes-bloqueados", "Mensajes bloqueados"},

	// 18. Reclamos, Devoluciones y Cambios
	{s18Claims, "que-es-un-reclamo", "Gestionar reclamos"},
	{s18Claims, "gestionar-mensaje-de-un-reclamo", "Gestionar mensajes de un reclamo"},
	{s18Claims, "gestionar-resolucion-de-reclamos", "Gestionar resolución de reclamos"},
	{s18Claims, "gestionar-evidencia-de-reclamos", "Gestionar evidencia de reclamos"},
	{s18Claims, "errores", "Errores"},
	{s18Claims, "gestionar-devoluciones", "Devoluciones"},
	{s18Claims, "changes", "Cambios"},

	// 19. Métricas y Tendencies
	{s19Metrics, "recuperacion-reputacion", "Programa de Despegue y Beneficio de Reputación"},
	{s19Metrics, "reputacion-de-vendedores", "Reputación de vendedores"},
	{s19Metrics, "tendencias", "Tendencias"},
	{s19Metrics, "mas-vendidos-en-mercado-libre", "Más vendidos en Mercado Libre"},
	{s19Metrics, "opiniones-sobre-producto", "Opiniones de productos"},
	{s19Metrics, "calidad-de-publicaciones", "Calidad de publicaciones"},
	{s19Metrics, "experiencia-de-compra", "Experiencia de compra"},
	{s19Metrics, "recurso-de-visitas", "Visitas"},
	{s19Metrics, "conoce-como-estan-los-vendedores-frente-la-carga-de-atributos", "Carga de atributos"},

	// 20. MercadoLíder y Tiendas Oficiales
	{s20OfficialStores, "tienda-oficial", "Tiendas Oficiales"},

	// 21. Guía para inmuebles (vertical)
	{s21RealEstate, "introduccion-guia-de-inmuebles", "Introducción guía de inmuebles"},
	{s21RealEstate, "primeros-pasos-inmuebles", "Primeros pasos"},
	{s21RealEstate, "configuracion-o-requisitos-previos", "Configuración o requisitos previos"},
	{s21RealEstate, "obtencion-del-access-token", "Obtención del Access Token"},
	{s21RealEstate, "consulta-de-usuarios", "Consulta de usuarios"},
	{
		s21RealEstate,
		"pasos-rapidos-para-publicar-un-inmueble-de-prueba",
		"Pasos rápidos para publicar un inmueble de prueba",
	},
	{s21RealEstate, "categorias-atributos-inmuebles", "Categorías y atributos"},
	{s21RealEstate, "categorias-inmuebles", "Categorías"},
	{s21RealEstate, "atributos-inmuebles", "Atributos"},
	{s21RealEstate, "localizar-inmuebles", "Localizar Inmuebles"},
	{s21RealEstate, "gestionar-paquetes-de-inmuebles", "Gestionar paquetes de inmuebles"},
	{s21RealEstate, "contratacion-de-paquetes-de-publicacion", "Contratación de paquetes de publicación"},
	{
		s21RealEstate,
		"paquetes-y-permisos-para-proyectos-desarrollos-o-emprendimientos-inmobiliarios",
		"Paquetes y permisos para proyectos inmobiliarios",
	},
	{s21RealEstate, "publica-inmueble", "Publica Inmuebles"},
	{
		s21RealEstate,
		"publicaciones-de-tiendas-oficiales-para-inmuebles",
		"Publicaciones de Tiendas Oficiales para Inmuebles",
	},
	{s21RealEstate, "actualiza-tus-publicaciones", "Actualiza tus publicaciones"},
	{
		s21RealEstate,
		"ciclo-de-vida-de-las-publicaciones-de-inmuebles",
		"Ciclo de vida de las publicaciones de Inmuebles",
	},
	{s21RealEstate, "variaciones-para-inmuebles", "Variaciones"},
	{s21RealEstate, "actualizacion-variacion-inmuebles", "Actualización variación inmuebles"},
	{s21RealEstate, "calidad-de-las-publicaciones-inmuebles", "Calidad de las Publicaciones"},
	{s21RealEstate, "desarrollos-inmobiliarios", "Desarrollos inmobiliarios"},
	{s21RealEstate, "leads-inmuebles", "Leads"},
	{s21RealEstate, "solicitud-de-visita", "Solicitud de visita"},
	{s21RealEstate, "estadisticas-de-interacciones-en-inmuebles", "Estadísticas de interacciones en Inmuebles"},
	{s21RealEstate, "experiencia-para-inmuebles", "Experiencia para inmuebles"},
	{s21RealEstate, "glosario-inmuebles", "Glosario"},

	// 22. Guía para vehículos (vertical)
	{s22Vehicles, "introduccion-vehiculos", titleIntro},
	{s22Vehicles, "consulta-usuarios", "Consulta Usuarios"},
	{s22Vehicles, "categorias-y-atributos", "Categorías y Atributos"},
	{s22Vehicles, "localizacion-de-vehiculos", "Localiza vehículos"},
	{s22Vehicles, "vehiculos-gestiona-paquetes", "Gestiona paquetes de vehículos"},
	{s22Vehicles, "publica-vehiculos", "Publica vehículos"},
	{s22Vehicles, "vehiculos-sincroniza-publicaciones", "Sincroniza publicaciones (vehículos)"},
	{s22Vehicles, "vehiculos-gestiona-preguntas-y-contactos", "Gestiona preguntas y contactos"},
	{s22Vehicles, "persona-interesadas", "Personas Interesadas"},
	{s22Vehicles, "credits-motors", "Créditos pre aprobados"},
	{s22Vehicles, "calidad-de-publicaciones-vehiculos", "Calidad de publicaciones (vehículos)"},

	// 23. Guía para servicios (vertical)
	{s23Services, "guia-para-servicios", titleIntro},
	{s23Services, "servicios-consulta-usuarios", "Consulta Usuarios"},
	{s23Services, "elige-tipo-de-servicio", "Elige tipo de servicio"},
	{s23Services, "administra-areas-de-cobertura", "Administra áreas de cobertura"},
	{s23Services, "publica-servicios-vis", "Publica servicios"},
	{s23Services, "servicio-sincroniza-publicaciones", "Sincroniza publicaciones"},
	{s23Services, "consultas-avanzadas-2", "Consultas avanzadas"},

	// 24. Guía para Mercado Ads
	{s24MercadoAds, "introduccion-a-mercado-ads", "Introducción a Mercado Ads"},
	{s24MercadoAds, "pads-read", "Product Ads"},
	{s24MercadoAds, "ads-bads", "Brand Ads"},
	{s24MercadoAds, "display", "Display Ads"},
	{s24MercadoAds, "product-ads-para-catalogo-y-user-products-lectura", "Product Ads para Catálogo y User Products"},
	{s24MercadoAds, "bonificaciones-para-product-ads", "Bonificaciones para Product Ads"},

	// 25. MCP
	{s25MCP, "mcp-server", "Mercado Libre MCP Server"},

	// 26. FAQs
	{s26FAQs, "faq-preguntas-frecuentes", "Preguntas Frecuentes"},
	{s26FAQs, "rate-limit-error-429", "Rate Limit / Error 429"},
	{s26FAQs, "stock-multiwarehouse", "Gestión de stock multiorigen / User Products"},
	{s26FAQs, "me1-me2-y-envio-gratis", "ME1 / ME2 y envío gratis"},
	{s26FAQs, "mercado-envios-costos-y-cotizaciones", "Mercado Envíos - Costos y cotizaciones"},
	{s26FAQs, "facturacion-billing-info", "Facturación / Billing info"},
	{s26FAQs, "imagenes-y-moderaciones", "Imágenes y moderaciones"},
	{s26FAQs, "promotions-pricing", "Promotions / Pricing"},
	{s26FAQs, "items-atributos-de-envio-y-dimensiones", "Items - Atributos de envío y dimensiones"},
}
