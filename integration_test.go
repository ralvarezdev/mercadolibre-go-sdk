//go:build integration

// Integration tests against the live MercadoLibre API.
//
// Prerequisites:
//   - A MercadoLibre app (test or production).
//   - A valid access token for the app's authorized user.
//
// Required environment variables:
//
//	ML_ACCESS_TOKEN   — a valid bearer access token
//
// Optional:
//
//	ML_SELLER_ID      — int64 seller ID to use for order/item queries
//	ML_ITEM_ID        — an existing item ID (e.g. "MLA123456789") for item reads
//
// Run:
//
//	ML_ACCESS_TOKEN=... go test -v -tags integration ./...

package mercadolibre_test

import (
	"context"
	"net/url"
	"os"
	"strconv"
	"testing"

	mercadolibre "github.com/ralvarezdev/mercadolibre-go-sdk"
)

func integrationClient(t *testing.T) *mercadolibre.Client {
	t.Helper()
	tok := os.Getenv("ML_ACCESS_TOKEN")
	if tok == "" {
		t.Skip("ML_ACCESS_TOKEN not set; skipping integration test")
	}
	return mercadolibre.NewClient(mercadolibre.WithStaticToken(tok))
}

func envInt64(key string) int64 {
	v := os.Getenv(key)
	if v == "" {
		return 0
	}
	n, _ := strconv.ParseInt(v, 10, 64)
	return n
}

// ---------------------------------------------------------------------------
// Auth / identity
// ---------------------------------------------------------------------------

func TestIntegration_UsersMe(t *testing.T) {
	c := integrationClient(t)
	me, err := c.Users.Me(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if me.ID == 0 {
		t.Error("user id is 0")
	}
	t.Logf("authenticated as %s (id=%d site=%s)", me.Nickname, me.ID, me.SiteID)
}

// ---------------------------------------------------------------------------
// Sites & catalog
// ---------------------------------------------------------------------------

func TestIntegration_SitesList(t *testing.T) {
	c := integrationClient(t)
	sites, err := c.Sites.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(sites) == 0 {
		t.Fatal("expected at least one site")
	}
	t.Logf("%d sites returned; first: %s (%s)", len(sites), sites[0].ID, sites[0].Name)
}

func TestIntegration_SiteGet(t *testing.T) {
	c := integrationClient(t)
	ctx := context.Background()
	// get the authenticated user's site
	me, err := c.Users.Me(ctx)
	if err != nil {
		t.Skip("cannot get user:", err)
	}
	site, err := c.Sites.Get(ctx, me.SiteID)
	if err != nil {
		t.Fatal(err)
	}
	if site.ID != me.SiteID {
		t.Errorf("site id %q != %q", site.ID, me.SiteID)
	}
	t.Logf("site %s: %s, %d categories", site.ID, site.Name, len(site.Categories))
}

func TestIntegration_CurrenciesList(t *testing.T) {
	c := integrationClient(t)
	cur, err := c.Currencies.List(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if len(cur) == 0 {
		t.Fatal("expected currencies")
	}
	t.Logf("%d currencies", len(cur))
}

// ---------------------------------------------------------------------------
// Items
// ---------------------------------------------------------------------------

func TestIntegration_ItemGet(t *testing.T) {
	itemID := os.Getenv("ML_ITEM_ID")
	if itemID == "" {
		t.Skip("ML_ITEM_ID not set")
	}
	c := integrationClient(t)
	item, err := c.Items.Get(context.Background(), itemID)
	if err != nil {
		t.Fatal(err)
	}
	if item.ID != itemID {
		t.Errorf("item id %q != %q", item.ID, itemID)
	}
	t.Logf("item %s: %q price=%.2f %s status=%s", item.ID, item.Title, item.Price, item.CurrencyID, item.Status)
}

func TestIntegration_SearchUserItems(t *testing.T) {
	sellerID := envInt64("ML_SELLER_ID")
	if sellerID == 0 {
		c := integrationClient(t)
		me, err := c.Users.Me(context.Background())
		if err != nil {
			t.Skip("cannot get user:", err)
		}
		sellerID = me.ID
	}
	c := integrationClient(t)
	resp, err := c.Items.SearchUserItems(context.Background(), sellerID, mercadolibre.ListOptions{Limit: 10}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("seller %d has %d total items; returned %d ids", sellerID, resp.Paging.Total, len(resp.Results))
}

// ---------------------------------------------------------------------------
// Orders
// ---------------------------------------------------------------------------

func TestIntegration_OrdersSearch(t *testing.T) {
	sellerID := envInt64("ML_SELLER_ID")
	if sellerID == 0 {
		c := integrationClient(t)
		me, err := c.Users.Me(context.Background())
		if err != nil {
			t.Skip("cannot get user:", err)
		}
		sellerID = me.ID
	}
	c := integrationClient(t)
	resp, err := c.Orders.SearchBySeller(context.Background(), sellerID, mercadolibre.ListOptions{Limit: 5}, nil)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("seller %d: %d total orders, returned %d", sellerID, resp.Paging.Total, len(resp.Results))
	for _, o := range resp.Results {
		t.Logf("  order %d status=%s amount=%.2f", o.ID, o.Status, o.TotalAmount)
	}
}

// ---------------------------------------------------------------------------
// Questions
// ---------------------------------------------------------------------------

func TestIntegration_QuestionsSearch(t *testing.T) {
	itemID := os.Getenv("ML_ITEM_ID")
	if itemID == "" {
		t.Skip("ML_ITEM_ID not set")
	}
	c := integrationClient(t)
	resp, err := c.Questions.Search(context.Background(), url.Values{"item": {itemID}})
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("item %s: %d questions", itemID, resp.Total)
}

// ---------------------------------------------------------------------------
// Page iterator smoke test
// ---------------------------------------------------------------------------

func TestIntegration_PageIterator(t *testing.T) {
	sellerID := envInt64("ML_SELLER_ID")
	if sellerID == 0 {
		c := integrationClient(t)
		me, err := c.Users.Me(context.Background())
		if err != nil {
			t.Skip("cannot get user:", err)
		}
		sellerID = me.ID
	}
	c := integrationClient(t)
	iter := mercadolibre.NewPageIterator(
		mercadolibre.ListOptions{Limit: 5},
		func(ctx context.Context, o mercadolibre.ListOptions) ([]mercadolibre.Order, mercadolibre.Paging, error) {
			resp, err := c.Orders.SearchBySeller(ctx, sellerID, o, nil)
			if err != nil {
				return nil, mercadolibre.Paging{}, err
			}
			return resp.Results, resp.Paging, nil
		},
	)
	total := 0
	pages := 0
	ctx := context.Background()
	for {
		orders, ok, err := iter.Next(ctx)
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			break
		}
		pages++
		total += len(orders)
		if pages >= 3 { // limit to 3 pages in the test
			break
		}
	}
	t.Logf("iterated %d pages, %d orders", pages, total)
}
