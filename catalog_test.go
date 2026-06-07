package mercadolibre

import (
	"context"
	"net/http"
	"testing"
)

func TestCurrencies(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/currencies/ARS" {
			w.Write([]byte(`{"id":"ARS","description":"Peso argentino","symbol":"$","decimal_places":2}`))
			return
		}
		w.Write([]byte(`[{"id":"ARS","description":"Peso argentino","symbol":"$","decimal_places":2},
		  {"id":"BRL","description":"Real","symbol":"R$","decimal_places":2}]`))
	})

	all, err := c.Currencies.List(context.Background())
	if err != nil || len(all) != 2 {
		t.Fatalf("list: %v %d", err, len(all))
	}
	ars, err := c.Currencies.Get(context.Background(), "ARS")
	if err != nil || ars.Symbol != "$" || ars.DecimalPlaces != 2 {
		t.Fatalf("get: %v %+v", err, ars)
	}
}

func TestLocationsState(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/classified_locations/states/UY-RO" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"id":"UY-RO","name":"Rocha","country":{"id":"UY","name":"Uruguay"},
		  "cities":[{"id":"C1","name":"Cabo Polonio"},{"id":"C2","name":"Chuy"}]}`))
	})
	st, err := c.Locations.State(context.Background(), "UY-RO")
	if err != nil {
		t.Fatal(err)
	}
	if st.Name != "Rocha" || st.Country == nil || st.Country.ID != "UY" || len(st.Cities) != 2 {
		t.Errorf("bad state: %+v", st)
	}
}

func TestCategoriesGet(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"id":"MLA5725","name":"Accesorios para Vehículos",
		  "path_from_root":[{"id":"MLA1743","name":"Autos"},{"id":"MLA5725","name":"Accesorios"}],
		  "children_categories":[{"id":"MLA1747","name":"Audio","total_items_in_this_category":10}]}`))
	})
	cat, err := c.Categories.Get(context.Background(), "MLA5725")
	if err != nil {
		t.Fatal(err)
	}
	if cat.Name == "" || len(cat.PathFromRoot) != 2 || len(cat.ChildrenCategories) != 1 {
		t.Errorf("bad category: %+v", cat)
	}
	if cat.ChildrenCategories[0].TotalItemsInThisCategory != 10 {
		t.Errorf("child count not decoded: %+v", cat.ChildrenCategories[0])
	}
}
