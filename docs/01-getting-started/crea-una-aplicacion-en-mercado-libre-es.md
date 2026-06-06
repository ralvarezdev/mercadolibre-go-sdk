---
title: Crea una aplicación en Mercado Libre
slug: crea-una-aplicacion-en-mercado-libre-es
section: 01-getting-started
source: https://developers.mercadolibre.com.ve/es_ar/crea-una-aplicacion-en-mercado-libre-es
captured: 2026-06-06
---

# Crea una aplicación en Mercado Libre

> Source: <https://developers.mercadolibre.com.ve/es_ar/crea-una-aplicacion-en-mercado-libre-es>

## Content

Puedes usar esta documentación para las siguientes unidades de negocio:

Última actualización 16/01/2026

## Crea una aplicación en Mercado Libre

Para crear una aplicación, debes loguearte e ingresar a **Mis aplicaciones** ([Argentina](https://developers.mercadolibre.com.ar/devcenter/), [Brasil](https://developers.mercadolivre.com.br/devcenter/), [Chile](https://developers.mercadolibre.cl/devcenter/), [México](https://developers.mercadolibre.com.mx/devcenter/), [Colombia](https://developers.mercadolibre.com.co/devcenter/), [Uruguay](https://developers.mercadolibre.com.uy/devcenter/), [Peru](https://developers.mercadolibre.com.pe/devcenter/), [Ecuador](https://developers.mercadolibre.com.ec/devcenter/) e [Venezuela](https://developers.mercadolibre.com.ve/devcenter/)) y completar la información que se solicita. Luego, obtendrás un Client Id y Secret Key, necesarios para autenticarte con nuestra API.

Antes de crear una aplicación, asegúrate que la cuenta que estás utilizando es **la cuenta del propietario** en la solución que va a desarrollar, evitando así problemas de transferencia de cuenta en el futuro. Recomendamos que la cuenta se cree bajo una entidad legal.

Estos son los pasos para crear una aplicación dentro de Mercado Libre, que te permitirá acceder al ecosistema de APIs públicas desde una integración:

1. Accede a nuestro [DevCenter](https://developers.mercadolibre.com.ve/devcenter).
2. Haz clic en **Crear nueva aplicación**, y completa todos los datos obligatorios.
3. En URIs de redirect se suman las posibles URIs de redireccionamiento en donde se devolverá el Code de la autorización recibida. Completa con la raíz del dominio.
4. Scopes
5. **Tópicos**: contiene un check list que clasifica por temas específicos, podrás seleccionar únicamente [los intereses que tienes en recibir notificaciones](https://developers.mercadolibre.com.ar/es_ar/productos-recibe-notificaciones) y en el campo **"Notificaciones callbacks URL"** se debe configurar una ruta para recibir estas notificaciones.****
6. Guarda el proyecto y serás redirigido a la página de inicio donde incluiremos tu aplicación. Puedes ver el ID y la clave secreta (Secret_Key) que tu aplicación ha expuesto. Con estos valores, podemos comenzar nuestra integración.

## Gestionar mis aplicaciones

Una vez creada la aplicación dentro de Mercado Libre puedes acceder en [DevCenter](https://developers.mercadolibre.com.ve/devcenter). Si tienes alguna aplicación ya generada, ingresa a la vista de **Configurar** para visualizar y gestionar tu aplicación.

## Configurar

Existen cuatro grupos de información en este formulario:

- Configuración de la aplicación
- Información Básica de la aplicación
- [Autenticación y Seguridad](https://developers.mercadolibre.com.ve/es_ar/autenticacion-y-autorizacion)
- [Configuraciones de Notificaciones](https://developers.mercadolibre.com.ar/es_ar/productos-recibe-notificaciones)

### Configuración de la aplicación

**client id**: es el APP ID de la aplicación que se creó.

**client secret**: clave de tu aplicación en Mercado Libre. **Este código es secreto, no lo compartas con nadie**.

**programar renovación**: acción para programar la actualización del Client Secret.

**renovar ahora**: acción para la renovación del Client Secret.

**Editar aplicación**

Siempre que quieras modificar el Client Secret puedes hacerlo de forma manual siguiendo estos pasos:

1. Accede a "Configuración de la aplicación".
2. Cambia el modo para "Ocultar" o "Mostrar" el Client Secret.

Haz clic en los tres puntos y selecciona una de las acciones que se muestran para programar la forma de renovar el Client secret: **"Renovar ahora"** o **"Programar renovación"**.

**Renovar ahora**

Esta es la confirmación para realizar la renovación de Client Secret. Al seleccionar esta opción, se generará automáticamente una nueva clave en ese mismo momento, la clave anterior caducará y se llevará a cabo la renovación.

Recomendamos actualizar la nueva clave en tus desarrollos a la brevedad, debido a que en ese lapso de tiempo los nuevos usuarios que quieran dar permiso a la aplicación tendrán error.

**Programar renovación**

Esta es la opción que recomendamos usar, ya que maximiza la seguridad de tu integración, tendrás la posibilidad de preparar tu desarrollo y sus diferentes ambientes (desarrollo/test), para el cambio de clave en la fecha de actualización programada.

Para esto:

1. Selecciona la fecha en la que desees que la clave actual expire, el selector se desplegará hasta 7 días.
2. También podrás seleccionar la hora, el selector se desplegará mostrándote opciones cada 30 minutos.
3. Por último, haz clic en **Renovar**, para confirmar la actualización programada del Client Secret en la fecha y hora que indicaste.****

Al programar la renovación tendrás 2 Client Secret “vigentes” **Client secret nuevo** y **Client secret actual** antes de que culmine el plazo.

![](https://http2.mlstatic.com/storage/developers-site-cms-admin/191728505714-1e.png)

Por otro lado, una vez realizada la confirmación para la actualización programada, tendrás disponibles las opciones de Cancelar renovación (acción para la cancelación de la renovación del Client Secret ) o Expirar ahora (acción para la renovación del Client Secret ).

**Cancelar renovación**

Una vez programada la renovación del Client Secret , es posible cancelarla. Si se cancela la renovación, el Client Secret generado expirará y continuará siendo válido el Client Secret vigente.

**Expirar Ahora**

Esta acción te permitirá adelantar la renovación programada, el Client Secret nuevo es el que queda funcional y el Client Secret vigente expirará en ese mismo momento.

![](https://http2.mlstatic.com/storage/developers-site-cms-admin/191728101849-1h.png)

### Consideraciones sobre scopes

Existen varios tipos de aplicaciones. No obstante, las dividiremos en tres grupos para explicar los scopes requeridos.

#### Aplicaciones de solo lectura

Aplicación que permite a un usuario anónimo o autenticado acceder a información personalizada de MELI. En este caso, un usuario anónimo podría buscar artículos, leer descripciones, etc. y un usuario autenticado puede ver la información personal. Si no realizas ninguna modificación a los datos de MELI (ninguna actualización de la información de usuario, publicación de artículos ni compra de artículos), todo lo que necesitas es un scope de lectura. Recuerda que cualquier intento por modificar datos a través de las API de Meli arrojaría error.

#### Aplicaciones online de lectura/escritura

Este tipo de aplicación permite que un usuario anónimo realice ciertas operaciones de solo lectura en MELI, así como también permite a un usuario autenticado modificar datos, publicar nuevos artículos (vender), publicar pedidos (comprar), etc. En este caso, la aplicación requiere un scope de escritura para que el usuario pueda otorgar permisos de escritura y la aplicación actúe en su nombre. La aplicación podrá modificar datos en nombre del usuario mientras el access token tenga validez. Una vez expirado, el usuario debe renovar el access token para volver a tener acceso.

#### Aplicaciones offline de lectura/escritura

Si tu aplicación debe actuar en nombre del usuario aún cuando este último está offline, requerirá permiso de acceso offline por parte del usuario. Al solicitar este scope, una vez aceptada por el usuario, la aplicación tendrá tanto el access token para actuar en nombre del usuario como un refresh token para obtener un nuevo access token válido cuando expire el anterior.

## Administrar permisos

Podrás acceder al listado de usuarios que le [dieron permisos a tu aplicación.](https://developers.mercadolibre.com.ve/es_ar/gestiona-tus-aplicaciones#Usuarios-que-le-dieron-permisos-a-tu-aplicación)

**Nuevo**: autorización creada en las últimas 24 horas.

**Inactivo (bullet gray)**: autorización sin usar por más de 3 meses.

**Inactivo (bullet blue)**: autorización sin usar menos de 3 meses.

**Activo**: autorización con uso constante.

## Eliminar

Ingresando a **Mis aplicaciones** tienes la opción de **Eliminar**, esta acción te permite borrar la aplicación. Una vez realizada esta acción, no hay forma de recuperar la aplicación eliminada.

**Conoce más:** sobre nuestro [programa de certificación](https://developers.mercadolibre.com.ve/es_ar/developer-partner-program).

**Siguiente:** [Autenticación y Autorización](https://developers.mercadolibre.com.ar/es_ar/autenticacion-y-autorizacion).
