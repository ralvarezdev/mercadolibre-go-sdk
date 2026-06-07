# mercadolibre-go-sdk

A Go SDK for the [MercadoLibre](https://www.mercadolibre.com) (marketplace) developer APIs.

## Repository layout

```
auth/                 Self-contained OAuth package (no import cycle).
docs/                 Curated API reference, one Markdown file per documentation page.
  README.md           Master index / table of contents (26 sections, 193 pages).
  01-getting-started/ OAuth, app management, test users, best practices…
  02-users/           Users, addresses, account data…
  04-items/           Items & search…
  12-products/        Catalog, user products, pricing, size guides…
  13-shipping/        Mercado Envíos (ME1 / ME2, Flex, Fulfillment…)
  14-promotions/      Central de promociones…
  15-sales-orders/    Orders, packs, payments, feedback…
  16-billing/         Billing data and reports…
  17-messaging/       Post-sale messaging…
  18-claims/          Claims, returns, changes…
  …                   (see docs/README.md for the full map)
examples/             Runnable examples.
  quickstart/         OAuth + PKCE flow → typed + generic calls.
  list-orders/        Paginate orders with PageIterator.
  publish-item/       Create, describe, and pause a listing.
notifications/        Webhook handler package (Topic constants + http.Handler).
tools/docgen/         Generator that fetches the live docs and writes docs/*.md.
```

## Documentation generator

`tools/docgen` fetches each page from
`https://developers.mercadolibre.com.ve/es_ar/<slug>`, extracts the article body, converts it
to Markdown (headings, tables, fenced code blocks, links), and renders MercadoLibre's interactive
"resource reference" tables as structured per-endpoint sections with `bash` request and `json`
response examples. It also lists every `api.mercadolibre.com` endpoint referenced on the page.

```bash
go run ./tools/docgen              # regenerate every page in the sitemap
go run ./tools/docgen -only items  # only slugs containing "items"
go run ./tools/docgen -section 04-items
go run ./tools/docgen -list        # print the sitemap and exit
```

The page inventory lives in `tools/docgen/sitemap.go`.

## API basics

- **REST base URL:** `https://api.mercadolibre.com`
- **Auth:** OAuth 2.0 (authorization code + PKCE); rotating refresh tokens, persisted via `auth.Store`.
- **Sites:** country sites such as `MLA` (Argentina), `MLB` (Brazil), `MLM` (Mexico), `MLC` (Chile)…
- **Source docs language:** Spanish (es_AR).

## Quickstart

```go
// Static token (scripts / testing)
client := mercadolibre.NewClient(mercadolibre.WithStaticToken(accessToken))

// Refreshing token source (production)
client = mercadolibre.NewClient(
    mercadolibre.WithTokenSource(auth.NewRefreshingTokenSource(cfg, tok, store, 0)),
    mercadolibre.WithRequestsPerSecond(8, 16), // 8 req/s, burst 16
)

// Typed service call
me, err := client.Users.Me(ctx)

// Generic escape hatch for any endpoint without a typed wrapper
site, err := mercadolibre.Get[mercadolibre.SiteDetail](ctx, client, mercadolibre.EPSiteByID, "MLA")
```

### Running the examples

```bash
# OAuth PKCE flow → save token → call Users.Me
go run ./examples/quickstart -client-id $ID -client-secret $SECRET -redirect $URI -site MLA

# List orders (reads ML_ACCESS_TOKEN from env)
ML_ACCESS_TOKEN=... go run ./examples/list-orders -pages 5

# Publish a test item, add a description, then pause it
ML_ACCESS_TOKEN=... go run ./examples/publish-item -site MLA -category MLA1055
```

### Integration tests

Run against the live API by providing a real access token:

```bash
ML_ACCESS_TOKEN=... go test -v -tags integration ./...
# Optional extras:
ML_SELLER_ID=...  # seller ID for order/item queries
ML_ITEM_ID=...    # existing item ID for item reads
```

## Typed services

| Service | Methods (read) | Methods (write) |
|---|---|---|
| `client.Users` | `Me`, `Get`, `GetMulti`, `Addresses`, `AcceptedPaymentMethods`, `Brands` | — |
| `client.Sites` | `List`, `Get`, `Search`, `ListingTypes`, `ListingType`, `ListingPrices`, `PredictCategory` | — |
| `client.Items` | `Get`, `GetMulti`, `GetMultiFields`, `Description`, `SearchUserItems`, `ScanUserItems`, `GetVariations` | `Create`, `Update`, `UpdateDescription`, `UpdateStatus`, `UpdatePrice`, `UpdateStock`, `UploadPictureFromURL`, `CreateVariation`, `UpdateVariation`, `DeleteVariation` |
| `client.Categories` | `SiteList`, `Get`, `Attributes`, `SaleTerms` | — |
| `client.Domains` | `Get`, `TechnicalSpecs` | — |
| `client.Currencies` | `List`, `Get`, `Convert` | — |
| `client.Locations` | `Countries`, `Country`, `State`, `City` | — |
| `client.Orders` | `Get`, `Search`, `SearchBySeller` | `PostFeedback`, `AddNote` |
| `client.Shipments` | `Get`, `Items`, `History`, `Payments`, `Costs`, `Carrier`, `SLA`, `LeadTime`, `Labels` | — |
| `client.Questions` | `Get`, `Search`, `ResponseTime` | `Ask`, `Answer`, `Delete` |
| `client.Messaging` | `PackMessages`, `Get`, `GetAttachment` | `Send` |
| `client.Claims` | `Get`, `Search`, `SearchByPlayer`, `History` | `TakeAction` |
| `client.Metrics` | `ItemVisits`, `UserVisits`, `ItemVisitsTimeWindow`, `SiteTrends`, `CategoryTrends`, `Reputation` | — |
| `client.Promotions` | `Get`, `User`, `Items`, `ItemPromotions` | `EnrollItem`, `UnenrollItem` |
| `client.Billing` | `Info`, `Reports`, `DownloadReport` | — |
| `client.MissedFeeds(…)` | paginated feed catch-up | — |

Anything not yet wrapped: use `mercadolibre.Get[T]`, `Post[T,B]`, `Put[T,B]`, or `Delete[T]` directly.

## Pagination helpers

```go
// Offset/limit pagination — works with any list endpoint
iter := mercadolibre.NewPageIterator(
    mercadolibre.ListOptions{Limit: 50},
    func(ctx context.Context, o mercadolibre.ListOptions) ([]mercadolibre.Order, mercadolibre.Paging, error) {
        resp, err := client.Orders.SearchBySeller(ctx, sellerID, o, nil)
        if err != nil { return nil, mercadolibre.Paging{}, err }
        return resp.Results, resp.Paging, nil
    },
)
for {
    orders, ok, err := iter.Next(ctx)
    if !ok || err != nil { break }
    // handle orders
}

// Scan/scroll (>1000 seller items)
scanner := client.Items.ScanUserItems(sellerID, 100, nil)
for {
    ids, ok, err := scanner.Next(ctx)
    if !ok || err != nil { break }
    // hydrate with GetMulti
}
```

## Status constants

```go
// Order statuses
mercadolibre.OrderStatusPaid        // "paid"
mercadolibre.OrderStatusCancelled   // "cancelled"

// Shipment statuses
mercadolibre.ShipmentStatusDelivered     // "delivered"
mercadolibre.ShipmentStatusReadyToShip   // "ready_to_ship"

// Item statuses / conditions / buying modes
mercadolibre.ItemStatusActive    // "active"
mercadolibre.ItemStatusPaused    // "paused"
mercadolibre.ItemConditionNew    // "new"
mercadolibre.BuyingModeBuyItNow  // "buy_it_now"

// Question, claim, payment statuses also defined in statuses.go
```

## Roadmap

1. ✅ Capture the full MercadoLibre documentation set into `docs/` (raw → curated).
2. ✅ **M1 — Core client + auth**: generic request core, endpoint constants, `APIError`, retries/backoff, `x/time/rate` limiter, full OAuth (AuthCodeURL/Exchange/Refresh/PKCE), rotating-refresh `TokenSource`, memory/file `Store`. First typed resources: Users + Sites. Runnable `examples/quickstart`.
3. ✅ **M2 — Backbone resources**: Items (get/multiget, search, scan/scroll), Categories/Domains, Currencies, Locations, public site search.
4. ✅ **M3 — Transactions**: Orders, Shipments, Questions & Answers, `missed_feeds`, `notifications/` webhook handler.
5. ✅ **M4 — Breadth**: Messaging (post-sale), Claims/returns, Metrics (visits, trends, reputation), Promotions.
6. ✅ **M5 — Billing + CI**: Billing info + reports + download. CI workflow (vet/build/race on Go 1.26).
7. ✅ **M6 — Writes + polish**: Item create/update/pause/price/stock/description + picture upload + variations CRUD. Order feedback + notes. Promotion item enrollment/unenrollment. `GetRaw` for binary downloads. Typed status constants (`statuses.go`). Generic offset/limit `PageIterator`. Integration test harness (`//go:build integration`). `list-orders` and `publish-item` examples.
8. ✅ **M7 — Surface completion**: `Users.AcceptedPaymentMethods/Brands`. `Sites.ListingTypes/ListingType/ListingPrices/PredictCategory`. `Shipments.Carrier/SLA/LeadTime/Labels` (label PDF). `Claims.History/TakeAction`. `Messaging.GetAttachment`. `Questions.ResponseTime`. `Promotions.Get`.

## License

See [LICENSE](LICENSE).
