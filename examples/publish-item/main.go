// Command publish-item demonstrates creating a new item listing, updating its
// description, and then pausing it using the MercadoLibre Go SDK.
//
// Usage:
//
//	ML_ACCESS_TOKEN=<token> go run ./examples/publish-item \
//	    -site MLA -category MLA1055 -title "Widget Pro" \
//	    -price 999.99 -currency ARS
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	mercadolibre "github.com/ralvarezdev/mercadolibre-go-sdk"
)

func main() {
	site := flag.String("site", "MLA", "site ID (MLA, MLB, MLM…)")
	category := flag.String("category", "", "category ID, e.g. MLA1055 (required)")
	title := flag.String("title", "Test item — SDK publish demo", "listing title")
	price := flag.Float64("price", 999.99, "listing price")
	currency := flag.String("currency", "ARS", "currency ID")
	qty := flag.Int("qty", 1, "available_quantity")
	listingType := flag.String("type", "gold_special", "listing_type_id")
	pauseAfter := flag.Bool("pause", true, "pause the item after creating it")
	flag.Parse()

	if *category == "" {
		log.Fatal("-category is required; find one at GET /sites/{site}/categories")
	}

	tok := os.Getenv("ML_ACCESS_TOKEN")
	if tok == "" {
		log.Fatal("ML_ACCESS_TOKEN environment variable not set")
	}

	client := mercadolibre.NewClient(mercadolibre.WithStaticToken(tok))
	ctx := context.Background()

	me, err := client.Users.Me(ctx)
	if err != nil {
		log.Fatalf("users.me: %v", err)
	}
	fmt.Printf("Publishing as %s on site %s\n\n", me.Nickname, *site)

	// Build the listing request with the minimum required fields.
	req := mercadolibre.CreateItemRequest{
		Title:             *title,
		CategoryID:        *category,
		Price:             *price,
		CurrencyID:        *currency,
		AvailableQuantity: *qty,
		BuyingMode:        string(mercadolibre.BuyingModeBuyItNow),
		Condition:         string(mercadolibre.ItemConditionNew),
		ListingTypeID:     *listingType,
	}

	item, err := client.Items.Create(ctx, req)
	if err != nil {
		log.Fatalf("create item: %v", err)
	}
	fmt.Printf("Created item: %s\n  Title:  %s\n  Price:  %.2f %s\n  Status: %s\n  URL:    %s\n\n",
		item.ID, item.Title, item.Price, item.CurrencyID, item.Status, item.Permalink)

	// Add a description.
	desc, err := client.Items.UpdateDescription(ctx, item.ID,
		"This listing was created by the mercadolibre-go-sdk publish-item example.")
	if err != nil {
		log.Printf("update description: %v (non-fatal)", err)
	} else {
		fmt.Printf("Description updated: %q\n\n", desc.PlainText)
	}

	// Optionally pause the item so it doesn't show publicly.
	if *pauseAfter {
		updated, err := client.Items.UpdateStatus(ctx, item.ID, string(mercadolibre.ItemStatusPaused))
		if err != nil {
			log.Fatalf("pause item: %v", err)
		}
		fmt.Printf("Item paused. Status is now: %s\n", updated.Status)
	}
}
