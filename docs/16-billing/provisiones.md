---
title: Provisiones
slug: provisiones
section: 16-billing
source: https://developers.mercadolibre.com.ve/es_ar/provisiones
captured: 2026-06-06
---

# Provisiones

> Source: <https://developers.mercadolibre.com.ve/es_ar/provisiones>

## Endpoints referenced

- `https://api.mercadolibre.com/billing/integration/group/ML/order/details?order_ids=$ORDER_ID`
- `https://api.mercadolibre.com/billing/integration/group/ML/order/details?order_ids=1234567890000`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/details`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/flex/details`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/full/details`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/insurtech/details`
- `https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/MP/details`
- `https://api.mercadolibre.com/billing/integration/periods/key/2021-06-01/group/ML/details?document_type=BILL&limit=1`
- `https://api.mercadolibre.com/billing/integration/periods/key/2022-10-01/group/ML/insurtech/details?document_type=BILL&limit=1`
- `https://api.mercadolibre.com/billing/integration/periods/key/2023-03-01/group/ML/flex/details?document_type=BILL&limit=1`
- `https://api.mercadolibre.com/billing/integration/periods/key/2023-03-01/group/ML/full/details?document_type=BILL&limit=1`
- `https://api.mercadolibre.com/billing/integration/periods/key/2024-05-01/group/MP/details?document_type=BILL&limit=1`
- `https://api.mercadolibre.com/billing/integration/periods/key/2024-11-01/group/MP/details?document_type=BILL&limit=1000&from_id=0`
- `https://api.mercadolibre.com/billing/integration/periods/key/2024-11-01/group/MP/details?document_type=BILL&limit=1000&from_id=12345678`

## Content

Última actualización 16/03/2026

## Provisiones

Obtén el detalle de las provisiones de facturas y cargos de ventas para un período en particular, el grupo de facturación (Mercado Libre o Mercado Pago) y el tipo de documento (Factura o Nota de Crédito) según la unidad de negocio que elijas: Mercado Libre, Mercado Pago, Mercado Envíos Flex, Fulfillment e Insurtech. Puedes filtrar con los parámetros: group (ML, MP) y document_type (BILL, CREDIT_NOTE).

## Parámetros de paginación

Proponemos el uso de 2 parámetros para manejar la paginación:

- limit: limita la cantidad de resultados a obtener. El valor mínimo es 1 y el máximo valor permitido es 1000. Por defecto su valor es 150.
- from_id: permite buscar a partir de un Id de detalle específico. Este valor se devuelve en el campo last_id de la respuesta JSON. Por defecto su valor es 0.

Para poder ordenar y obtener los resultados de forma correcta, se deben anexar a la request los siguientes parámetros:

- sort_by: propiedad por la que se quiere ordenar (ID o DATE)
- order_by: orientación del ordenado (ASC o DESC)

La API de reportes de facturación permite ajustar la cantidad de resultados por página mediante el parámetro limit. Por defecto, este valor es 150, con un máximo permitido de 1000. Esto significa que puedes incrementar el número de registros por solicitud hasta 1000, según tus necesidades.

La frecuencia de consumo depende del volumen de datos y de las necesidades específicas de tu aplicación. Si manejas grandes volúmenes de información, es recomendable realizar solicitudes periódicas, ajustando el limit y utilizando el from_id para paginar los resultados de forma efectiva. Por ejemplo, si deseas obtener los primeros 1000 registros, puedes establecer limit=1000 y from_id=0. Para la siguiente página, mantienes limit=1000, from_id=<last_id de la request anterior> y así sucesivamente. Este enfoque te permite dividir la información en páginas manejables y procesarlas de manera eficiente.

## Filtros opcionales

- date_sort: permite ordenar la búsqueda.
  - asc: ordena los resultados de manera ascendente (valor por default)
  - desc: ordena los resultados de manera descendente
  - Ejemplo: date_sort=asc
- sort_by: permite seleccionar por qué campo ordenar.
  - Valores posibles: ID (valor por default) y DATE
- detail_type: permite buscar por tipos de detalles.
  - charge: trae solamente cargos.
  - bonus: trae solamente bonificaciones.
  - Ejemplo: detail_type=charge
- detail_sub_types: permite filtrar por subtipos de detalles. Se pueden definir varios separados por coma.
  - Valores posibles:
  - Ejemplo: detail_sub_types=CV, BV
- detail_excluded_sub_types: permite excluir de la búsqueda los subtipos de detalles indicados. Se pueden definir varios separados por coma.
  - Ejemplo: not_subtypes=CXD, BXD
- marketplace_type: permite buscar por el market del cargo y/o bonus.
  - Valores posibles:
  - Ejemplo: marketplace_type=SHIPPING
- order_ids: permite buscar por uno o varios id de la order. Disponible para Mercado Libre.
  - Ejemplo: order_ids=2294412230
- item_ids: permite buscar por uno o más id de la publicación.
  - Ejemplo: item_ids=724159812
- document_ids: permite buscar por uno o más id de la factura.
  - Ejemplo: document_ids=987046992
- detail_ids: permite buscar por uno o más id del detalle.
  - Ejemplo: detail_ids=724159812
- offset: permite buscar desde un número de resultado en adelante. El valor mínimo permitido es 0 y el valor máximo permitido es 10000. Por defecto el valor es 0 – Te recomendamos utilizar más filtros y acotar los resultados.
- limit: limita la cantidad de resultados. Por defecto el mínimo es 1 y el máximo valor permitido: 1000.
- from_id: permite buscar a partir de un Id de detalle específico. Este valor se devuelve en el campo last_id de la respuesta JSON. Por defecto su valor es 0.

## Ejemplo de paginación: Detalles de Mercado Pago

**Primera página:**

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2024-11-01/group/MP/details?document_type=BILL&limit=1000&from_id=0
```

**Segunda página:**

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2024-11-01/group/MP/details?document_type=BILL&limit=1000&from_id=12345678
```

## Consideraciones

¿Cómo hay que integrarse para asegurar que, aunque se hagan varias consultas, la información no quede duplicada?

Para evitar duplicados al realizar múltiples consultas, es fundamental utilizar correctamente los parámetros limit y from_id en cada solicitud. El parámetro limit define la cantidad de registros a obtener y from_id permite indicar un id de detalle específico. Al incrementar el limit en cada solicitud y enviar el from_id, aseguras que cada página de resultados sea única y no se repitan registros.

- Para obtener la primera página: limit=1000 y from_id=0.
- Para la segunda página: limit=1000 y from_id=<last_id de request anterior>.
- Y así sucesivamente hasta llegar a consultar todos los detalles.

Este método garantiza una paginación efectiva sin duplicados.

Van a encontrar también el parámetro offset. El offset permite buscar desde un número de resultado en adelante. El valor mínimo permitido es 0 y el valor máximo permitido es 9999. Este parámetro solo se recomienda utilizar en casos donde la cantidad de detalles es menor a 10000.

## Mercado Libre

Verás los cargos facturados, información de la venta, descuentos, envíos y la publicación.

Importante:

 La estructura de la tarifa de venta para MLB se actualizó para separar el costo por vender en la plataforma, el costo por cobrar con Mercado Pago y la tasa de parcelamento la cual depende del método y la cantidad de cuotas elegidas por el comprador (no genera Nota Fiscal, ya que no corresponde a un servicio o transacción). Además, algunos productos pueden incluir un costo fijo adicional a la tarifa de venta. Para más información visitar 

Saiba más sobre las tarifas del Mercado Libre

. 

Para MLB, la respuesta de la API incluirá una nueva entidad con información detallada sobre la composición de la tarifa de venta (sale_fee). Esta mejora permitirá visualizar de forma más clara los componentes de la tarifa asociados a la venta, separando los descuentos aplicados y los rebates recibidos por cada order.

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/details
```

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2021-06-01/group/ML/details?document_type=BILL&limit=1
```

### Respuesta:

```javascript
{
   "offset": 0,
   "limit": 150,
   "total": 1,
   "last_id": 12345678,
   "results": [
       {
           "charge_info": {
               "legal_document_number": "0011A02842483",
               "legal_document_status": "PROCESSED",
               "legal_document_status_description": "Procesado",
               "creation_date_time": "2023-11-19T00:00:30",
               "detail_id": 12345678,
               "transaction_detail": "Cargo por vender",
               "debited_from_operation": "YES",
               "debited_from_operation_description": "Si",
               "status": null,
               "status_description": null,
               "charge_bonified_id": null,
               "detail_amount": 615.95,
               "detail_type": "CHARGE",
               "detail_sub_type": "CV"
           },
           "discount_info": {
               "charge_amount_without_discount": 815.95,
               "discount_amount": 200.00,
               "discount_reason": "Descuento general",
               "applied_percentage": 9.61,
           },
           "sales_info": [
               {
                   "order_id": 2000005750612628,
                   "operation_id": 58682854541,
                   "sale_date_time": "2023-11-18T23:59:26",
                   "sales_channel": "Mercado Libre",
                   "payer_nickname": "RAFACLORAFACLO",
                   "state_name": "Formosa",
                   "transaction_amount": 8490.00,
                   "financing_transfer_total": 102.75,
                   "financing_fee": 2.75
                   "sale_fee": {
                       "gross":13.84,
                       "net": 8.39,
                       "rebate": 5.45,
                       "discount": 0.0,
                       "discount_reason": reason
                   }
               }
           ],
           "shipping_info": {
               "shipping_id": "42313873858",
               "pack_id": null,
               "receiver_shipping_cost": 868.49
           },
           "items_info": [
               {
                   "item_id": "MLA781129295",
                   "item_title": "Anteojos Sol Polarizados Aviador Policarbonato Filtro Uv400",
                   "item_type": "gold_pro",
                   "item_category": "Ropa y Accesorios > Accesorios de Moda > Anteojos y Accesorios > Marcos de Anteojos",
                   "inventory_id": null,
                   "item_amount": 1,
                   "item_price": 8490.00,
                   "order_id": 2000005750612628,
                   "fees_added_in_publication": "Si"
               }
           ],
           "document_info": {
               "document_id": 2761583612
           },
           "marketplace_info": {
               "marketplace": "CORE"
           },
           "currency_info": {
               "currency_id": "ARS"
           },
       }
   ]
}
```

### Campos de respuesta Mercado Libre:

- charge_info: información del cargo.
  - legal_document_number: número del documento.
  - legal_document_status: estado de generación del documento.
    - Valores posibles: PROCESSING, PROCESSED.
  - legal_document_status_description: descripción internacionalizada del estado del documento legal_document_status.
  - creation_date_time: fecha de creación del cargo.
  - detail_id: identificador del cargo.
  - transaction_detail: detalle del cargo.
  - debited_from_operation: indica si está descontado de la operación.
    - Valores posibles: YES, NO, INAPPLICABLE.
  - debited_from_operation_description: descripción internacionalizada del campo debited_from_operation.
  - status: estado del cargo.
    - Valores posibles:
      - BONUS_ON_CREDIT_NOTE,
      - BONUS_PART_ON_CREDIT_NOTE,
      - BONUS_ON_BILL,
      - BONUS_PART_ON_BILL,
      - BONUS_ON, BONUS_PART_ON.
  - status_description: descripción internacionalizada de status.
  - charge_bonified_id: identificador del cargo que bonifica.
  - detail_amount: monto del cargo.
  - detail_type: tipo de detalle.
    - Valores posibles:
  - detail_sub_type: subtipos de detalles.
    - Valores posibles:
- discount_info: información sobre descuentos.
  - applied_percentage: porcentaje que se aplicó para calcular el monto del cargo. [Exclusivo para Argentina]
  - charge_amount_without_discount: valor del cargo sin descuento.
  - discount_amount: valor del descuento.
  - discount_reason: motivo del descuento.
  - rebate: valor del descuento por participación en campaña comercial.
- sales_info: información de las ventas.
  - order_id: identificador de la venta.
  - operation_id: identificador del pago.
  - sale_date_time: fecha y hora de venta.
  - sales_channel: canal de venta.
  - payer_nickname: cliente.
  - state_name: provincia.
  - transaction_amount: monto total de la venta.
  - financing_fee: Diferenciación en el precio según el número de cuotas elegidas por el comprador [Exclusivo para Brasil].
  - financing_transfer_total: importe total pagado por el cliente por el producto [Exclusivo para Brasil].
  - sale_fee: información sobre la tarifa de la venta (Exclusivo para Brasil).
    - gross: monto del cobro sin descuento.
    - net: monto del cobro.
    - rebate: monto del descuento por participación en campaña comercial.
    - discount: monto del descuento.
    - discount_reason: motivo del descuento.
- shipping_info: información del envío.
  - shipping_id: identificador del envío.
  - pack_id: identificador del paquete.
  - receiver_shipping_cost: envío a cargo del cliente.
- items_info: información sobre las publicaciones.
  - item_id: identificador de la publicación.
  - item_kit_id: identificador del kit. [Disponible apenas para Argentina, Brasil y México]
  - item_title: título de la publicación.
    - Kits Virtuales: En el caso que un producto pertenezca a un kit se concatena el nombre del item con *Producto en Kit: <nombre del kit>*. [Disponible apenas para Argentina, Brasil y México]
  - item_type: tipo de publicación.
  - item_category: categoría de publicación.
  - inventory_id: código de Mercado Libre.
  - item_amount: cantidad de ítems vendidos.
  - item_price: precio unitario del ítem.
  - order_id: orden a la que pertenece el ítem.
  - fees_added_in_publication: indica si la publicación ofrece cuotas. [Disponible apenas para Argentina]
- document_info: información del documento.
  - document_id: número Id de documentos.
- marketplace_info: información del marketplace.
  - marketplace: nombre del marketplace.
- currency_info: información de la moneda de acuerdo al site_id.
  - currency_id: identificador de la moneda de acuerdo al site_id.
- store_info: información de la sucursal.
  - store_id: identificador de la sucursal. [Disponible solo para MLM, MLC, MCO y MLA]
  - store_name: nombre de la sucursal. [Disponible solo para MLM, MLC, MCO y MLA]

## Reportes de Facturación por Orders y Packs

Este endpoint permite obtener los reportes de facturación por el filtro de orders y packs.

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/group/ML/order/details?order_ids=$ORDER_ID
```

### Parámetros de consulta:

- order_ids: Permite buscar por uno o varios id de order.
- pack_id: Permite buscar por un id de pack.
- sort_by:
  - Valores posibles: ID y DATE;
  - Valor por defecto: ID
- order_by: Permite ordenar la búsqueda.
  - Posibles valores: ASC, DESC;
  - Valor por defecto: ASC

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/group/ML/order/details?order_ids=1234567890000
```

### Respuesta:

```javascript
{
    "offset": 0,
    "limit": 150,
    "total": 1,
    "results": [
        {
            "order_id": 1234567890000,
            "payment_info": [
                {
                    "payment_id": 99999999999,
                    "date_approved": "2024-04-23T03:11:47",
                    "date_created": "2024-04-23T03:11:43",
                    "money_release_date": "2024-05-02T19:40:45",
                    "money_release_days": 28,
                    "money_release_status": "released",
                    "payer_id": 12345678,
                    "payment_method_id": "visa",
                    "payment_type_id": "credit_card",
                    "status": "approved",
                    "status_details": null,
                    "tax_details": [
                        {
                            "from": "collector",
                            "to": "mp",
                            "original_amount": 2018.99,
                            "refunded_amount": 0,
                            "mov_detail": "tax_withholding",
                            "mov_financial_entity": "retencion_ganancias",
                            "tax_id": 9999999997,
                            "tax_status": "applied"
                        },
                        {
                            "from": "collector",
                            "to": "mp",
                            "original_amount": 6056.97,
                            "refunded_amount": 0,
                            "mov_detail": "tax_withholding",
                            "mov_financial_entity": "retencion_iva",
                            "tax_id": 9999999998,
                            "tax_status": "applied"
                        },
                        {
                            "from": "collector",
                            "to": "mp",
                            "original_amount": 1211.39,
                            "refunded_amount": 0,
                            "mov_detail": "tax_withholding_collector",
                            "mov_financial_entity": "debitos_creditos",
                            "tax_id": 9999999999,
                            "tax_status": "applied"
                        },
                        {
                            "from": "collector",
                            "to": "mp",
                            "original_amount": 201.9,
                            "refunded_amount": 0,
                            "mov_detail": "tax_withholding_sirtac",
                            "mov_financial_entity": "cordoba",
                            "tax_id": 9999999990,
                            "tax_status": "applied"
                        }
                    ]
                }
            ],
            "sale_fee": {
               "gross": 120
               "net": 100
               "rebate":20
               "discount": 0
               "discount_reason": null
            }
            "details": [
                {
                    "charge_info": {
                        "legal_document_number": "0011A03800000",
                        "legal_document_status": "PROCESSED",
                        "legal_document_status_description": "Procesado",
                        "creation_date_time": "2024-04-22T23:12:02",
                        "detail_id": 5555566666,
                        "transaction_detail": "Cargo por venta",
                        "debited_from_operation": "YES",
                        "debited_from_operation_description": "Si",
                        "status": null,
                        "status_description": null,
                        "charge_bonified_id": null,
                        "detail_amount": 28265.86,
                        "detail_type": "CHARGE",
                        "detail_sub_type": "CV"
                    },
                    "discount_info": {
                        "charge_amount_without_discount": 28265.86,
                        "discount_amount": 0,
                        "discount_reason": "Descuento general",
                        "applied_percentage": 14,
                        "rebate": null
                    },
                    "sales_info": [
                        {
                            "order_id": 1234567890000,
                            "operation_id": 99999999999,
                            "sale_date_time": "2024-04-22T23:11:42",
                            "sales_channel": "Mercado Libre",
                            "payer_nickname": "NICKNAME",
                            "state_name": "Córdoba",
                            "transaction_amount": 201899
                            "financing_transfer_total": 102.75,
                            "financing_fee": 2.75
                        }
                    ],
                    "shipping_info": {
                        "shipping_id": "5555566666",
                        "pack_id": null,
                        "receiver_shipping_cost": null
                    },
                    "items_info": [
                        {
                            "item_id": "MLA920316309",
                            "item_title": "Calefactor A Gas Eskabe Miniconvex 5000 S21p Marfil Clase A",
                            "item_type": "gold_special",
                            "item_category": "Electrodomésticos y Aires Ac. > Climatización > Estufas y Calefactores > A Gas",
                            "inventory_id": null,
                            "item_amount": 1,
                            "item_price": 201899,
                            "order_id": 1234567890000,
                            "fees_added_in_publication": "No"
                        }
                    ],
                    "document_info": {
                        "document_id": 5555566666
                    },
                    "marketplace_info": {
                        "marketplace": "CORE"
                    },
                    "currency_info": {
                        "currency_id": "ARS"
                    }
                },
                {
                    "charge_info": {
                        "legal_document_number": "0011A03800000",
                        "legal_document_status": "PROCESSED",
                        "legal_document_status_description": "Procesado",
                        "creation_date_time": "2024-04-22T23:12:02",
                        "detail_id": 5555566666,
                        "transaction_detail": "Cargo por Mercado Envíos",
                        "debited_from_operation": "YES",
                        "debited_from_operation_description": "Si",
                        "status": null,
                        "status_description": null,
                        "charge_bonified_id": null,
                        "detail_amount": 9380.99,
                        "detail_type": "CHARGE",
                        "detail_sub_type": "CXD"
                    },
                    "discount_info": {
                        "charge_amount_without_discount": 18761.99,
                        "discount_amount": 9381,
                        "discount_reason": "Descuento general",
                        "rebate": null
                    },
                    "sales_info": [
                        {
                            "order_id": 1234567890000,
                            "operation_id": 99999999999,
                            "sale_date_time": "2024-04-22T23:11:42",
                            "sales_channel": "Mercado Libre",
                            "payer_nickname": "NICKNAME",
                            "state_name": "Córdoba",
                            "transaction_amount": 201899
                        }
                    ],
                    "shipping_info": {
                        "shipping_id": "5555566666",
                        "pack_id": null,
                        "receiver_shipping_cost": 0
                    },
                    "items_info": [
                        {
                            "item_id": "MLA920316309",
                            "item_title": "Calefactor A Gas Eskabe Miniconvex 5000 S21p Marfil Clase A",
                            "item_type": "gold_special",
                            "item_category": "Electrodomésticos y Aires Ac. > Climatización > Estufas y Calefactores > A Gas",
                            "inventory_id": null,
                            "item_amount": 1,
                            "item_price": 201899,
                            "order_id": 1234567890000,
                            "fees_added_in_publication": "No"
                        }
                    ],
                    "document_info": {
                        "document_id": 5555566666
                    },
                    "marketplace_info": {
                        "marketplace": "SHIPPING"
                    },
                    "currency_info": {
                        "currency_id": "ARS"
                    }
                }
            ]
        }
    ]
}
```

### Parámetros de respuesta

- order_id: Identificador de la venta.
- payment_info: Información del pago.
  - payment_id: Identificador del pago.
  - date_approved: Fecha de aprobación.
  - date_created: Fecha de creación.
  - money_release_date: Fecha de liberación del pago.
  - money_release_days: Días para la liberación del pago.
  - money_release_status: Estado de la liberación del pago.
  - payer_id: Identificador de cliente.
  - payment_method_id: Método de pago.
  - payment_type_id: Tipo de medio de pago.
  - status: Estado del pago.
  - status_details: Detalles del estado del pago.
  - tax_details: Detalles de impuestos.
- details: Detalles de cargos.
  - charge_info: Información del cobro.
  - discount_info: Información sobre descuentos.
  - sales_info: Información sobre la venta.
  - shipping_info: Información de la entrega.
  - items_info: Información de la publicación.
  - document_info: Información del documento.
  - marketplace_info: Información del marketplace.
  - currency_info: Información de la moneda según el site_id.
- sale_fee: Información sobre la tarifa de la venta (Exclusivo para Brasil).
  - gross: Monto del cobro sin descuento.
  - net: Monto del cobro.
  - rebate: Monto del descuento por participación en campaña comercial.
  - discount: Monto del descuento.
  - discount_reason: Motivo del descuento.

## Links útiles

### 1. Valores a recibir

- **[GET /orders](https://developers.mercadolibre.com.ar/es_ar/gestiona-ventas)**: datos del pedido.
  - **unit_price**: valor unitario del ítem con el descuento "de/por" ya aplicado.
  - **quantity**: cantidad de ítems del pedido.
  - **sale_fee**: tarifa por unidad.
  - **marketplace_fee**: tarifa totalizada en el pedido.
- **[GET /packs](https://developers.mercadolibre.com.ar/es_ar/gestion-packs)**: identificar las orders dentro de un pack.
  - **orders_ids**: identificadores de los pedidos que componen el pack.
- **[GET /shipments](https://developers.mercadolibre.com.ar/es_ar/envios)**: identificar el costo de envío.
  - **seller.cost**: costo de envío subvencionado por el vendedor.

**Ejemplo de cálculo simplificado:**
 (unit_price * quantity) - marketplace_fee - seller.cost = valor neto del pedido.

### 2. Costos y descuentos aplicados

- **[GET /orders/{order_id}/discounts](https://developers.mercadolibre.com.ar/es_ar/gestiona-ventas#Obtener-descuentos-aplicados)**: información de descuentos y campañas aplicadas al pedido.
  - **discounts**, **coupon**: tipos de descuentos aplicados.
  - **supplier**: proveedor de la campaña.
  - **meli_campaign**: campaña de descuentos asociada al cupón.
  - **offer_id**: identificador de la oferta (útil para rastrear la promoción).
  - **funding_mode**: tipo de promoción (ej.: sale_fee).
  - **amounts.total**: monto total del descuento (parte MELI + parte vendedor).
- **GET /items/{item_id}/sale_price**: identificar el precio de venta aplicado a un ítem.
  - **amount**: precio vigente del producto (ya con descuento).
  - **regular_amount**: precio original antes de la promoción.
  - **metadata.promotion_id**: identificador de la promoción asociada.
  - **metadata.promotion_type**: tipo de promoción (ej.: custom, deal).
- **GET /seller-promotions/offers/{offer_id}**: identificar cambios y estado de las ofertas promocionales.
  - **promotion_id**: ID de la promoción asociada.
  - **type**: tipo de promoción (ej.: DEAL).
  - **status.id**: estado actual de la promoción (ej.: ACTIVE, FINISHED).

### 3. Conciliación financiera

La conciliación se realiza a partir de la combinación de los siguientes recursos:

- GET /orders
- GET /orders/{id}/discounts
- GET /shipments
- GET /packs

El cálculo consolidado considera:

- Valor del ítem (unit_price * quantity).
- Tarifas (sale_fee, marketplace_fee).
- Costos de envío (seller.cost).
- Descuentos aplicados (discounts, coupon).

**Resultado:** Visión consolidada de los valores netos a recibir por pedido o pack.

## Mercado Pago

Verás el detalle de cargos facturados con información complementaria sobre la operación de Mercado Pago, como los movimientos, medios de pago, *payer*, sucursal, punto de venta, entre otros.

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 

https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/MP/details
```

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2024-05-01/group/MP/details?document_type=BILL&limit=1
```

### Respuesta:

```javascript
{
  "offset": 0,
  "limit": 1,
  "total": 1,
  "results": [
    {
      "charge_info": {
        "legal_document_number": "0029A01508173",
        "legal_document_status": "PROCESSED",
        "legal_document_status_description": "Procesado",
        "detail_id": 24168819712,
        "movement_id": "199835301597",
        "transaction_detail": "Cargo de Mercado Pago",
        "debited_from_operation": "INAPPLICABLE",
        "debited_from_operation_description": "No aplica",
        "status": "BONUS_ON_BILL",
        "status_description": "Anulado en factura",
        "charge_bonified_id": null,
        "creation_date_time": "2023-07-19T07:29:02",
        "detail_amount": 3122.76,
        "detail_type": "CHARGE",
        "detail_sub_type": "CCMP"
      },
      "operation_info": {
        "operation_type": "BUY",
        "operation_type_description": "Pago",
        "reference_id": 60833750481,
        "sales_channel": "Checkout",
        "store_id": null,
        "store_name": null,
        "external_reference": "385080",
        "payer_nickname": "SALADO1958",
        "financing_fee": 9.2, // solo Brasil
        "financing_transfer_total": 109.2, // solo Brasil
        "transaction_amount": 73999
      },
      "perception_info": {
        "aliquot": null,
        "taxable_amount": null
      },
      "document_info": {
        "document_id": 2589999426
      },
      "marketplace_info": {
        "marketplace": "MP"
      },
      "currency_info": {
        "currency_id": "ARS"
      }
    }
  ]
}
```

### Campos de respuesta:

- charge_info: información del cargo.
  - legal_document_number: número del documento.
  - detail_id: identificador del cargo.
  - legal_document_status: estado de generación del documento.
    - **Valores posibles:** PROCESSING, PROCESSED, NOT_APPLICABLE.
  - legal_document_status_description: descripción internacionalizada.
  - movement_id: número de movimiento.
  - transaction_detail: detalle.
  - debited_from_operation:
    - **Valores posibles:** YES, NO, INAPPLICABLE.
  - debited_from_operation_description: descripción internacionalizada.
  - status: estado del cargo.
    - **Valores posibles:** BONUS_ON_CREDIT_NOTE, BONUS_PART_ON_CREDIT_NOTE, BONUS_ON_BILL, BONUS_PART_ON_BILL, BONUS_ON, BONUS_PART_ON.
  - status_description: descripción internacionalizada de status.
  - charge_bonified_id: identificador del cargo que bonifica.
  - creation_date_time: fecha del cargo.
  - detail_amount: monto del cargo.
  - detail_type: tipo de detalle.
  - detail_sub_type: subtipos de detalles.
- operation_info: información de la operación sobre la que aplica.
  - operation_type: tipo de operación.
    - **Valores posibles:** BUY, TAX.
  - operation_type_description: descripción internacionalizada.
  - reference_id: número de operación relacionada.
  - sales_channel: tipo de pago.
  - store_id: número de sucursal.
  - store_name: nombre de la sucursal.
  - external_reference: número de referencia externa.
  - payer_nickname: cliente.
  - financing_fee / financing_transfer_total: Exclusivo para Brasil.
  - transaction_amount: monto de la operación.
- perception_info: información de la percepción.
  - aliquot: alícuota.
  - taxable_amount: monto imponible.
- document_info: información del documento (document_id).
- marketplace_info: información del marketplace.
- currency_info: información de la moneda (currency_id).

## Mercado Envíos Flex

Importante:

 Disponible en MLA, MLC y MCO. 

Verás el detalle para conciliar las bonificaciones y anulaciones de Flex para un período en particular, el grupo de facturación Mercado Libre y el tipo de documento (nota de débito o nota de crédito). Además, información sobre el envío e información de la venta.

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/flex/details
```

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2023-03-01/group/ML/flex/details?document_type=BILL&limit=1
```

### Respuesta:

```javascript
{
  "offset": 0,
  "limit": 1,
  "total": 100,
  "last_id": 12345678,
  "results": [{
    "charge_info": {
      "legal_document_number": "00AA11AA00",
      "legal_document_status": "PROCESSED",
      "legal_document_status_description": "Procesado",
      "creation_date_time": "2023-02-21T12:35:58",
      "detail_id": 12345678,
      "detail_associated_id": 4040404040,
      "detail_amount": 163,
      "transaction_detail": "Anulación de bonificación por Mercado Envíos Flex",
      "detail_type": "CHARGE",
      "detail_sub_type": "CFLX",
      "concept_type": "FLEX"
    },
    "shipping_info": {
      "shipping_id": 4444455555,
      "receiver_nickname": "NICKNAME",
      "pack_id": "12345678",
      "receiver_shipping_cost": 814.99,
      "order": {
        "order_id": 9000000008888888,
        "date_created": "2023-02-15T11:54:51",
        "total_amount": 29499,
        "payment_id": 998899889988,
        "buyer_nickname": "NICKNAME"
      }
    },
    "document_info": { "document_id": 776677667711 }
  }],
  "errors": []
}
```

### Campos de respuesta:

- charge_info: información del cargo.
  - legal_document_number: número del documento.
  - legal_document_status: estado de generación del documento.
    - **Valores posibles:** PROCESSING, PROCESSED.
  - legal_document_status_description: descripción internacionalizada del estado del documento legal_document_status.
  - creation_date_time: fecha de creación del cargo.
  - detail_id: identificador del cargo.
  - detail_associated_id: identificador del cargo asociado (en caso de anulación de bonificación).
  - detail_amount: monto del cargo.
  - transaction_detail: detalle del cargo.
  - detail_type: tipo de detalle.
  - detail_sub_type: subtipos de detalles.
  - concept_type: tipo de concepto.
- shipping_info: información sobre envío.
  - shipping_id: identificador del envío.
  - receiver_nickname: cliente.
  - pack_id: número de paquete.
  - receiver_shipping_cost: costo del envío.
- order: información de la venta.
  - order_id: identificador de la venta.
  - date_created: fecha de la orden.
  - total_amount: total de la orden.
  - payment_id: identificador del pago.
  - buyer_nickname: cliente.
- document_info: información del documento.
  - document_id: id de documento.

## Fulfillment

Importante:

 Está disponible en MLA, MLB, MLM, MCO y MLC. 

Verás los cargos y las bonificaciones por colecta y/o almacenamiento para un período en particular, el grupo de facturación Mercado Libre y el tipo de documento (factura o nota de crédito). También información del producto almacenado o recolectado. Los tipos de cargos para el reporte de Fulfillment pueden ser por: retiro de stock, almacenamiento prolongado, servicio de colecta, incumplimiento, almacenamiento.

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/full/details
```

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2023-03-01/group/ML/full/details?document_type=BILL&limit=1
```

### Respuesta:

```javascript
{
  "offset": 0,
  "limit": 100,
  "total": 634,
  "last_id": 12345678,
  "results": [{
    "charge_info": {
      "legal_document_number": "000AAA00000000",
      "legal_document_status": "PROCESSED",
      "legal_document_status_description": "Procesado",
      "creation_date_time": "2021-07-23T16:37:58",
      "detail_id": 12345678,
      "detail_amount": 2.54,
      "transaction_detail": "Cargo por servicio de colecta Full",
      "charge_bonified_id": null,
      "detail_type": "CHARGE",
      "detail_sub_type": "CFCB",
      "concept_type": "FULFILLMENT",
      "payment_id": 222222222
    },
    "fulfillment_info": {
      "type": "INBOUND_COLLECT",
      "amount_per_unit": 2.54,
      "amount": 2.54,
      "sku": "3125404000009",
      "ean": "3125404000009",
      "item_id": "MLM788740252",
      "item_title": "VESTIDO CORTO AZUL MARINO BORDADO EN PECHO DEVENDI",
      "variation": "AZUL MARINO | EG",
      "quantity": 1,
      "volume_type": null,
      "inventory_id": "LLLGGKK12",
      "space": "Almacenamiento"
      "inbound_id": 555555,
      "volume_unit": "m3",
      "amount_per_volume_unit": 500,
      "volume": 0.00507,
      "volume_total": 0.00507
    },
    "document_info": { "document_id": 333333333 }
  }],
  "errors": []
}
```

### Campos de respuesta Mercado Envíos Fulfillment:

- charge_info: información del cargo.
  - legal_document_number: número del documento.
  - legal_document_status: estado de generación del documento.
    - **Valores posibles:** PROCESSING, PROCESSED.
  - legal_document_status_description: descripción internacionalizada del estado del documento legal_document_status.
  - creation_date_time: fecha de creación del cargo.
  - detail_id: identificador del cargo.
  - detail_associated_id: identificador del cargo asociado (en caso de anulación de bonificación).
  - detail_amount: monto del cargo.
  - transaction_detail: detalle del cargo.
  - detail_type: tipo de detalle.
  - detail_sub_type: subtipos de detalles.
  - concept_type: tipo de concepto.

Importante:

 [SOLAMENTE MLC y MLA] Para 

fulfillment_info

 con 

type

 = 

WITHDRAWAL

, se quitará el campo 

amount_per_unit

.

 El campo 

volume_type

 ahora puede tener los siguientes valores: small, medium, large, extralarge. 

- fulfillment_info: información de fulfillment.
  - type: tipo de fulfillment.
    - **Valores posibles:** WITHDRAWAL, AGING, INBOUND_COLLECT, INBOUND_PENALTY, WAREHOUSING, OVERAGE, SPACE_PURCHASE, SPACE_CANCELLATION.
  - amount_per_unit: monto por unidad.
  - amount: monto total.
  - sku: stock keeping unit.
  - item_id: número de publicación.
  - item_title: título de la publicación.
  - variation: variante del producto.
  - quantity: unidades almacenadas o recolectadas.
  - volume_type: tamaño de la unidad.
  - inventory_id: código del inventario de ML.
  - withdrawal_id: número de retiro – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - shipment_type: forma de retiro – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - volume_unit: unidad de medida (m3) – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - amount_per_volume_unit: monto por m3 – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - volume: volumen unitario (cm3) – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - volume_total: volumen total – TYPE WITHDRAWAL: Cargo por retiro de stock.
  - months_range: antigüedad en meses – TYPE AGING: Cargo por almacenamiento prolongado.
  - stock_details: detalles del stock – TYPE AGING: Cargo por almacenamiento prolongado.
  - inventory_status: estado del inventario – TYPE AGING: Cargo por almacenamiento prolongado.
  - inbound_id: número de envío – TYPE INBOUND_COLLECT / INBOUND_PENALTY.
  - penalty_type: tipo de incumplimiento – TYPE INBOUND_PENALTY: Cargo por incumplimiento.
  - warehouse_id: identificador de warehouse – TYPE WAREHOUSING: Cargo por almacenamiento.
  - size: tamaño de la unidad – TYPE WAREHOUSING: Cargo por almacenamiento.
  - item_quantity: unidades almacenadas – TYPE WAREHOUSING: Cargo por almacenamiento.
  - space: indica el tipo de espacio que se compra o cancela – TYPE SPACE_PURCHASE / SPACE_CANCELLATION.
    - **Valores posibles (2 espacios):** Pequeños y medianos, Grandes y extra grandes.
    - **Valores posibles (1 espacio):** Almacenamiento.
- document_info: información del documento.
  - document_id: ID del documento.

## Insurtech

Importante:

 Está disponible en MLA, MLB, MLC y MLM. 

Verás el detalle para conciliar los cargos y bonificaciones de las garantías aplicadas sobre los productos para un período en particular, el grupo de facturación Mercado Libre y el tipo de documento (Factura o Nota de Crédito).

### Llamada:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 

https://api.mercadolibre.com/billing/integration/periods/key/$KEY/group/ML/insurtech/details
```

### Ejemplo:

```javascript
curl -X GET -H 'Authorization: Bearer $ACCESS_TOKEN' 
https://api.mercadolibre.com/billing/integration/periods/key/2022-10-01/group/ML/insurtech/details?document_type=BILL&limit=1
```

### Respuesta:

```javascript
{
  "offset": 0,
  "limit": 150,
  "total": 1,
  "last_id": 12345678,
  "results": [
    {
      "charge_info": {
        "legal_document_number": "001112131415",
        "legal_document_status": "PROCESSED",
        "legal_document_status_description": "Procesado",
        "creation_date_time": "2022-10-04T22:24:18",
        "detail_id": 12345678,
        "detail_amount": 520.01,
        "transaction_detail": "Cargo por seguro de garantía extendida",
        "status": null,
        "status_description": null,
        "charge_bonified_id": null,
        "detail_type": "CHARGE",
        "detail_sub_type": "CEW",
        "concept_type": "WARRANTY"
      },
      "warranty_info": {
        "warranty_id": "11111111-43c2-44ea-8436-00000000",
        "certificate_id": "MLA999999",
        "warranty_product": "GAREX",
        "buyer_nickname": "TEST",
        "buyer_state_name": "Salta",
        "order": {
          "order_id": 102030405060,
          "order_items": [
            {
              "unit_price": 10791,
              "listing_type_id": "gold_special",
              "item": {
                "item_id": "MLA88888888",
                "title": "Auriculares Inalámbricos Jbl Tune 510bt Negro",
                "category_id": "MLA1234",
                "category_name": "Auriculares"
              }
            }
          ]
        },
        "quote_model": null,
        "quote_brand": null,
        "quote_description": ""
      },
      "prepaid_info": {
        "operation_id": 55558888,
        "movement_id": 123456789,
        "doc_id": 11111111111,
        "payment": {
          "payment_id": 5555555555,
          "date_created": "2022-10-04T22:23:40",
          "transaction_amount": 736.98,
          "money_release_date": "2023-03-03T22:23:41"
        }
      },
      "document_info": { "document_id": 3333333333 }
    }
  ]
}
```

### Campos de respuesta Insurtech:

- charge_info: información del cargo.
  - legal_document_number: número del documento.
  - legal_document_status: estado de generación del documento.
    - **Valores posibles:** PROCESSING, PROCESSED.
  - legal_document_status_description: descripción internacionalizada del estado del documento legal_document_status.
  - creation_date_time: fecha de creación del cargo.
  - detail_id: identificador del cargo.
  - detail_amount: monto del cargo.
  - transaction_detail: detalle del cargo.
  - status: estado del cargo.
    - **Valores posibles:** BONUS_ON_CREDIT_NOTE, BONUS_PART_ON_CREDIT_NOTE, BONUS_ON_BILL, BONUS_PART_ON_BILL, BONUS_ON, BONUS_PART_ON.
  - status_description: descripción internacionalizada de status.
  - charge_bonified_id: identificador del cargo que bonifica.
  - detail_type: tipo de detalle.
  - detail_sub_type: subtipos de detalles.
  - concept_type: tipo de concepto.
- warranty_info: información de la garantía.
  - warranty_id: identificador de la garantía.
  - certificate_id: identificador del certificado.
  - warranty_product: tipo de garantía.
    - **Valores posibles:** CARDS, GAREX, RODA.
  - buyer_nickname: número del comprador.
  - buyer_state_name: provincia del comprador.
  - order: información de la orden.
    - order_id: identificador de la orden.
    - order_items: lista de ítems de la orden.
    - listing_type_id: tipo de publicación.
    - item: información del ítem (item_id, title, category_id, category_name).
  - quote_model: modelo del producto. Aplica a RODA.
  - quote_brand: marca del producto. Aplica a RODA.
  - quote_description: descripción adicional. Aplica a RODA.
- prepaid_info: información del prepago.
  - operation_id: identificador de la operación.
  - movement_id: identificador del movimiento.
  - doc_id: identificador del documento.
  - payment: información del pago.
    - payment_id: identificador del pago.
    - date_created: fecha de pago.
    - transaction_amount: monto del pago.
    - money_release_date: fecha de liberación del dinero.
- document_info: información del documento.
  - document_id: ID del documento.

**Siguiente**: [Pagos](https://developers.mercadolibre.com.ve/es_ar/reportes-pagos).
