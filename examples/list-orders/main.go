// Command list-orders demonstrates paginating over a seller's orders using the
// generic PageIterator and printing a summary table.
//
// Usage:
//
//	ML_ACCESS_TOKEN=<token> go run ./examples/list-orders [-seller <id>] [-limit <n>]
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
	sellerID := flag.Int64("seller", 0, "seller user ID (defaults to the authenticated user)")
	pageSize := flag.Int("limit", 20, "orders per page")
	maxPages := flag.Int("pages", 3, "maximum number of pages to fetch (0 = all)")
	flag.Parse()

	tok := os.Getenv("ML_ACCESS_TOKEN")
	if tok == "" {
		log.Fatal("ML_ACCESS_TOKEN environment variable not set")
	}

	client := mercadolibre.NewClient(
		mercadolibre.WithStaticToken(tok),
		mercadolibre.WithRequestsPerSecond(8, 16),
	)
	ctx := context.Background()

	// resolve seller
	if *sellerID == 0 {
		me, err := client.Users.Me(ctx)
		if err != nil {
			log.Fatalf("users.me: %v", err)
		}
		*sellerID = me.ID
		fmt.Printf("Authenticated as %s (id=%d)\n\n", me.Nickname, me.ID)
	}

	iter := mercadolibre.NewPageIterator(
		mercadolibre.ListOptions{Limit: *pageSize},
		func(ctx context.Context, opts mercadolibre.ListOptions) ([]mercadolibre.Order, mercadolibre.Paging, error) {
			resp, err := client.Orders.SearchBySeller(ctx, *sellerID, opts, nil)
			if err != nil {
				return nil, mercadolibre.Paging{}, err
			}
			return resp.Results, resp.Paging, nil
		},
	)

	fmt.Printf("%-15s %-22s %10s  %s\n", "Order ID", "Status", "Amount", "Date")
	fmt.Println("-------------- ---------------------- ----------  ----------")

	total := 0
	pages := 0
	for {
		orders, ok, err := iter.Next(ctx)
		if err != nil {
			log.Fatalf("page %d: %v", pages+1, err)
		}
		if !ok {
			break
		}
		pages++
		for _, o := range orders {
			fmt.Printf("%-15d %-22s %10.2f  %s\n",
				o.ID, o.Status, o.TotalAmount, o.DateCreated.Format("2006-01-02"))
			total++
		}
		if *maxPages > 0 && pages >= *maxPages {
			fmt.Printf("\n[stopped after %d pages; use -pages 0 for all]\n", pages)
			break
		}
	}
	fmt.Printf("\nTotal orders shown: %d\n", total)
}
