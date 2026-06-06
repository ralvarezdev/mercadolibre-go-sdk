---
title: Notas en órdenes
slug: notas-en-ordenes
section: 15-sales-orders
source: https://developers.mercadolibre.com.ve/es_ar/notas-en-ordenes
captured: 2026-06-06
---

# Notas en órdenes

> Source: <https://developers.mercadolibre.com.ve/es_ar/notas-en-ordenes>

## Endpoints referenced

- `https://api.mercadolibre.com/orders/$ORDER_ID/notes`
- `https://api.mercadolibre.com/orders/$ORDER_ID/notes/$NOTE_ID`
- `https://api.mercadolibre.com/users/$CUST_ID/order_blacklist`

## Content

Última actualización 29/12/2025

## Notas en órdenes

## Agregar una nota a una orden

Una nota es una anotación que los vendedores pueden agregar a las órdenes. Las notas pueden tener hasta 300 caracteres cada una y, una vez que publicas una nota, puedes modificarla o eliminarla.

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H "Content-Type: application/json" -d 
{  	
   	"note": "test",
}
https://api.mercadolibre.com/orders/$ORDER_ID/notes
```

### Ver notas de órdenes

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/$ORDER_ID/notes
```

### Modificar una nota

```javascript
curl -X PUT -H 'Authorization: Bearer $ACCESS_TOKEN' -H "Content-Type: application/json" -d 
{
   	"note": "test2",
} 
https://api.mercadolibre.com/orders/$ORDER_ID/notes/$NOTE_ID
```

### Eliminar notas

```javascript
curl -X DELETE -H 'Authorization: Bearer $ACCESS_TOKEN' https://api.mercadolibre.com/orders/$ORDER_ID/notes/$NOTE_ID
```

### Bloquear ofertas

Aunque contamos con una prestación y flujos de prevención de fraude para mantener la seguridad de los compradores y vendedores, en algunas ocasiones puedes encontrar usuarios que, por algún motivo ofertan en las publicaciones. Estos casos pueden ser enviados a la lista negra para evitarles la posibilidad de ofertar.

### Bloquear ofertas de un usuario específico

Realiza una solicitud POST con el cust_id del usuario que deseas bloquear en el JSON y el tuyo en el url, como muestra el siguiente ejemplo:

```javascript
curl -X POST -H 'Authorization: Bearer $ACCESS_TOKEN' -H "Content-Type: application/json" -d
{ user_id: {cust_id} }
https://api.mercadolibre.com/users/$CUST_ID/order_blacklist
```

**Siguiente**: [Bloqueo de usuarios](https://developers.mercadolibre.com.ve/es_ar/usuarios-bloqueados).
