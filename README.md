# mercadolibre-go-sdk

A Go SDK for the [MercadoLibre](https://www.mercadolibre.com) (marketplace) developer APIs.

> **Status: early / documentation phase.** The API reference is being captured and curated
> before the client surface is implemented.

## Repository layout

```
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
- **Auth:** OAuth 2.0 (authorization code + refresh token); bearer access token per request.
- **Sites:** country sites such as `MLA` (Argentina), `MLB` (Brazil), `MLM` (Mexico), `MLC` (Chile)…
- **Source docs language:** Spanish (es_AR).

## Roadmap

1. ✅ Capture the full MercadoLibre documentation set into `docs/` (raw → curated).
2. ⬜ Curate the core marketplace resources (users, items, orders, shipping, questions) into
   SDK-ready endpoint + model specs.
3. ⬜ Implement the Go client: HTTP core, OAuth/token handling, typed resources and models.
4. ⬜ Tests and examples.

## License

See [LICENSE](LICENSE).
