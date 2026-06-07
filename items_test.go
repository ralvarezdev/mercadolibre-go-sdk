package mercadolibre

import (
	"context"
	"net/http"
	"strings"
	"testing"
)

const itemFixture = `{
  "id": "MLA1136716168",
  "site_id": "MLA",
  "title": "Zapatillas Avid Fof - Test Item",
  "seller_id": 1108966308,
  "category_id": "MLA109027",
  "price": 15000,
  "base_price": 15000,
  "currency_id": "ARS",
  "available_quantity": 2,
  "sold_quantity": 0,
  "buying_mode": "buy_it_now",
  "listing_type_id": "gold_pro",
  "condition": "new",
  "permalink": "https://articulo.mercadolibre.com.ar/MLA-1136716168-zapatillas-avid-fof-test-item-_JM",
  "pictures": [
    {"id":"963513-MLA","url":"http://x/O.jpg","secure_url":"https://x/O.jpg","size":"500x411","max_size":"898x739"}
  ],
  "shipping": {"mode":"not_specified","free_shipping":false,"logistic_type":"not_specified","tags":["adoption_required"]},
  "attributes": [
    {"id":"BRAND","name":"Marca","value_id":"11823494","value_name":"Propia"},
    {"id":"MODEL","name":"Modelo","value_id":null,"value_name":"EQ2122"}
  ],
  "status": "active",
  "tags": ["cart_eligible","test_item"]
}`

func TestItemsGet(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/items/MLA1136716168" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(itemFixture))
	})
	it, err := c.Items.Get(context.Background(), "MLA1136716168")
	if err != nil {
		t.Fatal(err)
	}
	if it.Title != "Zapatillas Avid Fof - Test Item" || it.Price != 15000 || it.CurrencyID != "ARS" {
		t.Errorf("bad item: %+v", it)
	}
	if it.Shipping == nil || it.Shipping.LogisticType != "not_specified" {
		t.Errorf("shipping not decoded: %+v", it.Shipping)
	}
	if len(it.Attributes) != 2 || it.Attributes[0].ID != "BRAND" {
		t.Errorf("attributes not decoded: %+v", it.Attributes)
	}
	// value_id can be null; ensure pointer handling works.
	if it.Attributes[1].ValueID != nil {
		t.Errorf("expected nil value_id for MODEL, got %v", *it.Attributes[1].ValueID)
	}
}

func TestItemsGetMulti(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.URL.Query().Get("ids"); got != "MLA1,MLA2" {
			t.Errorf("ids = %q", got)
		}
		w.Write([]byte(`[{"code":200,"body":` + itemFixture + `},{"code":404,"body":{}}]`))
	})
	res, err := c.Items.GetMulti(context.Background(), "MLA1", "MLA2")
	if err != nil {
		t.Fatal(err)
	}
	if len(res) != 2 || res[0].Code != 200 || res[1].Code != 404 {
		t.Fatalf("bad multiget: %+v", res)
	}
	if res[0].Body.ID != "MLA1136716168" {
		t.Errorf("body not decoded: %+v", res[0].Body)
	}
}

func TestScanUserItems(t *testing.T) {
	// Two pages then an empty page terminates iteration.
	pages := []string{
		`{"paging":{"total":3,"limit":2,"offset":0},"results":["MLA1","MLA2"],"scroll_id":"S1"}`,
		`{"paging":{"total":3,"limit":2,"offset":0},"results":["MLA3"],"scroll_id":"S1"}`,
		`{"paging":{"total":3,"limit":2,"offset":0},"results":[],"scroll_id":"S1"}`,
	}
	var call int
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("search_type") != "scan" {
			t.Errorf("missing search_type=scan: %s", r.URL.RawQuery)
		}
		if call > 0 && r.URL.Query().Get("scroll_id") != "S1" {
			t.Errorf("scroll_id not propagated on call %d: %s", call, r.URL.RawQuery)
		}
		w.Write([]byte(pages[call]))
		call++
	})

	it := c.Items.ScanUserItems(123, 2, nil)
	var all []string
	for {
		ids, ok, err := it.Next(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		if !ok {
			break
		}
		all = append(all, ids...)
	}
	if strings.Join(all, ",") != "MLA1,MLA2,MLA3" {
		t.Errorf("scanned ids = %v", all)
	}
	if call != 3 {
		t.Errorf("calls = %d, want 3", call)
	}
}

func TestSiteSearch(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/sites/MLA/search" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"site_id":"MLA","paging":{"total":1,"offset":0,"limit":50},
		  "results":[{"id":"MLA1","title":"X","price":99,"currency_id":"ARS","category_id":"MLA1","available_quantity":5}],
		  "available_sorts":[{"id":"price_asc","name":"Menor precio"}]}`))
	})
	q := ListOptions{Limit: 50}.Values()
	q.Set("q", "celular")
	resp, err := c.Sites.Search(context.Background(), "MLA", q)
	if err != nil {
		t.Fatal(err)
	}
	if resp.Paging.Total != 1 || len(resp.Results) != 1 || resp.Results[0].Price != 99 {
		t.Errorf("bad search response: %+v", resp)
	}
}
