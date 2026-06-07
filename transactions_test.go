package mercadolibre

import (
	"context"
	"net/http"
	"testing"
)

const orderFixture = `{
  "id": 2000003508419013,
  "status": "paid",
  "status_detail": null,
  "date_created": "2013-05-27T10:01:50.000-04:00",
  "date_closed": "2013-05-27T10:04:07.000-04:00",
  "order_items": [{
    "item": {"id":"MLB12345678","title":"Samsung Galaxy","variation_id":null,"variation_attributes":[]},
    "quantity": 1, "unit_price": 499, "currency_id": "BRL"
  }],
  "total_amount": 499,
  "currency_id": "BRL",
  "buyer": {"id": 123456789},
  "seller": {"id": 987654321},
  "payments": [{"id": 596707837, "transaction_amount": 499, "currency_id": "BRL", "status": "approved"}],
  "shipping": {"id": 20676482441},
  "tags": ["paid","not_delivered"]
}`

func TestOrdersGet(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/2000003508419013" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(orderFixture))
	})
	o, err := c.Orders.Get(context.Background(), 2000003508419013)
	if err != nil {
		t.Fatal(err)
	}
	if o.Status != "paid" || o.TotalAmount != 499 || len(o.OrderItems) != 1 {
		t.Errorf("bad order: %+v", o)
	}
	if o.OrderItems[0].Item.ID != "MLB12345678" || o.OrderItems[0].Quantity != 1 {
		t.Errorf("bad order item: %+v", o.OrderItems[0])
	}
	if o.Buyer == nil || o.Buyer.ID != 123456789 {
		t.Errorf("buyer not decoded: %+v", o.Buyer)
	}
	if len(o.Payments) != 1 || o.Payments[0].Status != "approved" {
		t.Errorf("payments not decoded: %+v", o.Payments)
	}
	if o.Shipping == nil || o.Shipping.ID != 20676482441 {
		t.Errorf("shipping not decoded: %+v", o.Shipping)
	}
}

func TestOrdersSearchBySeller(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		if q.Get("seller") != "89660613" || q.Get("order.status") != "paid" {
			t.Errorf("query = %s", r.URL.RawQuery)
		}
		w.Write([]byte(`{"paging":{"total":1,"offset":0,"limit":50},"results":[` + orderFixture + `]}`))
	})
	extra := map[string][]string{"order.status": {"paid"}}
	resp, err := c.Orders.SearchBySeller(context.Background(), 89660613, ListOptions{Limit: 50}, extra)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Paging.Total != 1 || len(resp.Results) != 1 {
		t.Errorf("bad search: %+v", resp)
	}
}

func TestQuestionsGetAndAnswer(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/questions/11751825075":
			if r.URL.Query().Get("api_version") != "4" {
				t.Errorf("missing api_version=4: %s", r.URL.RawQuery)
			}
			w.Write([]byte(`{"id":11751825075,"seller_id":179571326,"item_id":"MLA739200576",
			  "status":"ANSWERED","text":"Texto","date_created":"2021-02-08T17:51:21.746Z",
			  "answer":{"text":"ok","status":"ACTIVE","date_created":"2021-02-16T14:52:13.580-04:00"}}`))
		case r.Method == http.MethodPost && r.URL.Path == "/answers":
			w.Write([]byte(`{"text":"thanks","status":"ACTIVE"}`))
		default:
			t.Errorf("unexpected %s %s", r.Method, r.URL.Path)
		}
	})

	q, err := c.Questions.Get(context.Background(), 11751825075)
	if err != nil {
		t.Fatal(err)
	}
	if q.Status != "ANSWERED" || q.Answer == nil || q.Answer.Text != "ok" {
		t.Errorf("bad question: %+v", q)
	}
	a, err := c.Questions.Answer(context.Background(), 11751825075, "thanks")
	if err != nil {
		t.Fatal(err)
	}
	if a.Text != "thanks" {
		t.Errorf("bad answer: %+v", a)
	}
}

func TestShipmentsHistory(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/shipments/1234567899/history" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`[{"status":"ready_to_ship","substatus":"printed","date":"2016-12-30T12:32:35.000Z"},
		  {"status":"handling","substatus":"waiting_for_label_generation","date":"2016-12-30T12:32:35.000Z"}]`))
	})
	h, err := c.Shipments.History(context.Background(), 1234567899)
	if err != nil {
		t.Fatal(err)
	}
	if len(h) != 2 || h[0].Status != "ready_to_ship" || h[0].Substatus != "printed" {
		t.Errorf("bad history: %+v", h)
	}
}

func TestMissedFeeds(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("app_id") != "3486171129139063" || r.URL.Query().Get("topic") != "payments" {
			t.Errorf("query = %s", r.URL.RawQuery)
		}
		w.Write([]byte(`{"messages":[{"_id":"abc","resource":"/payments/1","user_id":465432224,
		  "topic":"payments","application_id":3486171129139063,"attempts":5,
		  "sent":"2019-10-17T17:15:30.279Z","received":"2019-10-17T17:15:30.259Z"}]}`))
	})
	feeds, err := c.MissedFeeds(context.Background(), MissedFeedsOptions{AppID: 3486171129139063, Topic: "payments"})
	if err != nil {
		t.Fatal(err)
	}
	if len(feeds) != 1 || feeds[0].Topic != "payments" || feeds[0].Attempts != 5 {
		t.Errorf("bad feeds: %+v", feeds)
	}
}
