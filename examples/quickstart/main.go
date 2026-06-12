// Command quickstart demonstrates the MercadoLibre Go SDK: the OAuth
// authorization-code + PKCE flow, a refreshing/persisted token source, and a
// couple of API calls (typed service + generic escape hatch).
//
// Usage:
//
//	# 1) print the authorization URL, open it, copy the ?code= from the redirect
//	go run ./examples/quickstart -client-id $ID -client-secret $SECRET \
//	    -redirect https://your.app/cb -site MLA
//
//	# 2) exchange the code (reuses the printed PKCE verifier)
//	go run ./examples/quickstart ... -code $CODE -verifier $VERIFIER
package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	mercadolibre "github.com/ralvarezdev/mercadolibre-go-sdk"
	"github.com/ralvarezdev/mercadolibre-go-sdk/auth"
)

func main() {
	clientID := flag.String("client-id", "", "application APP ID")
	clientSecret := flag.String("client-secret", "", "application Secret Key")
	redirect := flag.String("redirect", "", "redirect URI (must match the app config)")
	site := flag.String("site", string(mercadolibre.SiteArgentina), "site ID, e.g. MLA")
	code := flag.String("code", "", "authorization code from the redirect")
	verifier := flag.String("verifier", "", "PKCE code verifier from step 1")
	tokenFile := flag.String("token-file", "meli-token.json", "where to persist the token")
	flag.Parse()

	cfg := auth.Config{
		ClientID:     *clientID,
		ClientSecret: *clientSecret,
		RedirectURI:  *redirect,
		Site:         *site,
	}
	ctx := context.Background()

	// Step 1: no code yet -> print the authorization URL and a fresh PKCE pair.
	if *code == "" {
		pkce, err := auth.GeneratePKCE()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Open this URL, authorize, then copy the ?code= from the redirect:")
		fmt.Println(cfg.AuthCodeURL("state-"+pkce.Verifier[:8], auth.WithPKCE(pkce)))
		fmt.Println("\nThen re-run with:")
		fmt.Printf("  -code <CODE> -verifier %s\n", pkce.Verifier)
		return
	}

	// Step 2: exchange the code, persist the token, build a client.
	tok, err := cfg.Exchange(ctx, *code, auth.WithCodeVerifier(*verifier))
	if err != nil {
		log.Fatalf("exchange: %v", err)
	}
	store := auth.NewFileStore(*tokenFile)
	if err = store.Save(ctx, tok); err != nil {
		log.Fatalf("save token: %v", err)
	}

	client := mercadolibre.NewClient(
		mercadolibre.WithTokenSource(auth.NewRefreshingTokenSource(cfg, tok, store, 0)),
		mercadolibre.WithRequestsPerSecond(8, 16),
	)

	// Typed service call.
	me, err := client.Users.Me(ctx)
	if err != nil {
		log.Fatalf("users.me: %v", err)
	}
	fmt.Printf("Authenticated as %s (id=%d, site=%s)\n", me.Nickname, me.ID, me.SiteID)

	// Generic escape hatch for an endpoint without a typed wrapper yet.
	site2, err := mercadolibre.Get[mercadolibre.SiteDetail](ctx, client, mercadolibre.EPSiteByID, me.SiteID)
	if err != nil {
		log.Fatalf("site: %v", err)
	}
	fmt.Printf("Site %s default currency: %s\n", site2.ID, site2.DefaultCurrencyID)
}
