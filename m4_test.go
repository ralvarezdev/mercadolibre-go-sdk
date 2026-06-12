package mercadolibre

import (
	"context"
	"net/http"
	"testing"
	"time"
)

func TestMessagingPackMessages(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/messages/packs/2000000089077943/sellers/415458330" {
			t.Errorf("path = %q", r.URL.Path)
		}
		if r.URL.Query().Get("tag") != "post_sale" {
			t.Errorf("missing tag=post_sale: %s", r.URL.RawQuery)
		}
		if r.URL.Query().Get("mark_as_read") != "false" {
			t.Errorf("expected mark_as_read=false: %s", r.URL.RawQuery)
		}
		w.Write([]byte(`{"paging":{"limit":10,"offset":0,"total":1},
		  "conversation_status":{"status":"active"},
		  "messages":[{"id":"fd1d","site_id":"MLB","from":{"user_id":1},"to":{"user_id":2},
		    "status":"available","text":"Mensaje de prueba","message_resources":[{"id":"x","name":"packs"}]}],
		  "seller_max_message_length":350,"buyer_max_message_length":3500}`))
	})
	resp, err := c.Messaging.PackMessages(context.Background(), 2000000089077943, 415458330, ListOptions{}, false)
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Messages) != 1 || resp.Messages[0].Text != "Mensaje de prueba" || resp.SellerMaxLength != 350 {
		t.Errorf("bad messages: %+v", resp)
	}
}

func TestMessagingSend(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			t.Errorf("method = %s", r.Method)
		}
		w.Write([]byte(`{"id":"new1","from":{"user_id":415458330},"to":{"user_id":3037675074},"text":"hola"}`))
	})
	msg, err := c.Messaging.Send(context.Background(), 2000000089077943, 415458330, SendMessageRequest{
		From: MessageParty{UserID: 415458330},
		To:   MessageParty{UserID: 3037675074},
		Text: "hola",
	})
	if err != nil {
		t.Fatal(err)
	}
	if msg.ID != "new1" {
		t.Errorf("bad message: %+v", msg)
	}
}

func TestClaimsGet(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/post-purchase/v1/claims/5256749420" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"id":5256749420,"resource_id":2000007819609432,"status":"closed",
		  "type":"mediations","stage":"claim","resource":"order","reason_id":"PDD9549","fulfilled":true,
		  "players":[{"role":"complainant","type":"buyer","user_id":1325224382,"available_actions":[]},
		    {"role":"respondent","type":"seller","user_id":1330467461}],
		  "resolution":{"reason":"payment_refunded","benefited":["complainant"],"closed_by":"respondent","applied_coverage":false},
		  "site_id":"MLB","date_created":"2024-03-14T08:28:44.000-04:00","last_updated":"2024-03-21T05:19:22.000-04:00"}`))
	})
	cl, err := c.Claims.Get(context.Background(), 5256749420)
	if err != nil {
		t.Fatal(err)
	}
	if cl.Status != "closed" || cl.Type != "mediations" || len(cl.Players) != 2 {
		t.Errorf("bad claim: %+v", cl)
	}
	if cl.Resolution == nil || cl.Resolution.Reason != "payment_refunded" {
		t.Errorf("resolution not decoded: %+v", cl.Resolution)
	}
}

func TestClaimsSearchByPlayer(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("players.user_id") != "123456789" || q.Get("players.role") != "respondent" ||
			q.Get("status") != "opened" {
			t.Errorf("query = %s", r.URL.RawQuery)
		}
		w.Write(
			[]byte(
				`{"paging":{"total":1,"offset":0,"limit":30},"data":[{"id":1,"status":"opened","type":"returns","resource":"order","resource_id":2,"site_id":"MLM","players":[]}]}`,
			),
		)
	})
	resp, err := c.Claims.SearchByPlayer(
		context.Background(),
		123456789,
		"respondent",
		"opened",
		ListOptions{Limit: 30},
	)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Paging.Total != 1 || len(resp.Data) != 1 {
		t.Errorf("bad search: %+v", resp)
	}
}

func TestMetricsItemVisits(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/items/visits" || r.URL.Query().Get("ids") != "MCO1" {
			t.Errorf("req = %s %s", r.URL.Path, r.URL.RawQuery)
		}
		w.Write([]byte(`{"item_id":"MCO1","date_from":"2021-01-01T00:00:00Z","date_to":"2021-02-01T00:00:00Z",
		  "total_visits":536,"visits_detail":[{"company":"mercadolibre","quantity":536}]}`))
	})
	from := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2021, 2, 1, 0, 0, 0, 0, time.UTC)
	v, err := c.Metrics.ItemVisits(context.Background(), "MCO1", from, to)
	if err != nil {
		t.Fatal(err)
	}
	if v.TotalVisits != 536 || len(v.VisitsDetail) != 1 {
		t.Errorf("bad visits: %+v", v)
	}
}

func TestMetricsTrends(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/trends/MLA" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`[{"keyword":"labial","url":"https://listado.mercadolibre.com.ar/labial#trend"}]`))
	})
	tr, err := c.Metrics.SiteTrends(context.Background(), "MLA")
	if err != nil {
		t.Fatal(err)
	}
	if len(tr) != 1 || tr[0].Keyword != "labial" {
		t.Errorf("bad trends: %+v", tr)
	}
}

func TestPromotionsUser(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/seller-promotions/users/1356551933" || r.URL.Query().Get("app_version") != "v2" {
			t.Errorf("req = %s %s", r.URL.Path, r.URL.RawQuery)
		}
		w.Write([]byte(`{"results":[{"id":"P-MLB1806015","type":"MARKETPLACE_CAMPAIGN","status":"started",
		  "name":"Campanha","benefits":{"type":"REBATE","meli_percent":5,"seller_percent":25}}],
		  "paging":{"offset":0,"limit":50,"total":1}}`))
	})
	resp, err := c.Promotions.User(context.Background(), 1356551933, ListOptions{})
	if err != nil {
		t.Fatal(err)
	}
	if len(resp.Results) != 1 || resp.Results[0].Benefits.SellerPercent != 25 {
		t.Errorf("bad promotions: %+v", resp)
	}
}

func TestPromotionsItems_loosePaging(t *testing.T) {
	// The promotions items endpoint sometimes returns paging as an empty array.
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("promotion_type") != "DEAL" {
			t.Errorf("missing promotion_type: %s", r.URL.RawQuery)
		}
		w.Write(
			[]byte(
				`{"results":[{"id":"MLA639970000","status":"started","price":4037,"original_price":4427}],"paging":[]}`,
			),
		)
	})
	resp, err := c.Promotions.Items(context.Background(), "MLA1111", "DEAL", nil)
	if err != nil {
		t.Fatalf("loose paging should not fail: %v", err)
	}
	if len(resp.Results) != 1 || resp.Results[0].Price != 4037 {
		t.Errorf("bad items: %+v", resp)
	}
}
