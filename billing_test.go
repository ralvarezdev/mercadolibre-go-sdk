package mercadolibre

import (
	"context"
	"net/http"
	"testing"
)

func TestBillingInfo_numberCustID(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/orders/billing-info/MLA/677487519924852462" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(`{"site_id":"MLA","buyer":{"cust_id":123123123,
		  "billing_info":{"name":"Juan","last_name":"Sanchez",
		    "identification":{"type":"DNI","number":"307722738"},
		    "taxes":{"taxpayer_type":{"id":"01","description":"Consumidor Final"}},
		    "address":{"street_name":"Aysen","street_number":"30","zip_code":"5000","country_id":"AR",
		      "state":{"code":"01","name":"Buenos Aires"}},
		    "attributes":{"cust_type":"CO","is_normalized":true}}},
		  "seller":{"cust_id":0}}`))
	})
	bi, err := c.Billing.Info(context.Background(), "MLA", "677487519924852462")
	if err != nil {
		t.Fatal(err)
	}
	if bi.Buyer.CustID != "123123123" { // number decoded to string form
		t.Errorf("cust_id = %q, want 123123123", bi.Buyer.CustID)
	}
	if bi.Buyer.BillingInfo == nil || bi.Buyer.BillingInfo.Identification.Type != "DNI" {
		t.Errorf("billing info not decoded: %+v", bi.Buyer.BillingInfo)
	}
	if bi.Buyer.BillingInfo.Taxes.TaxpayerType.Description != "Consumidor Final" {
		t.Errorf("taxes not decoded: %+v", bi.Buyer.BillingInfo.Taxes)
	}
}

func TestBillingInfo_stringCustID(t *testing.T) {
	// MPE/MLU return cust_id as a quoted string — FlexID must handle both.
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"buyer":{"cust_id":"325999999","billing_info":{"name":"Adrian",
		  "identification":{"type":"RUC","number":"01234567891"}}},
		  "seller":{"cust_id":"1469999999"},"buyer_id":325999999,"seller_id":1469999999}`))
	})
	bi, err := c.Billing.Info(context.Background(), "MPE", "abc")
	if err != nil {
		t.Fatal(err)
	}
	if bi.Buyer.CustID != "325999999" || bi.Seller.CustID != "1469999999" {
		t.Errorf("cust_id flex decode failed: buyer=%q seller=%q", bi.Buyer.CustID, bi.Seller.CustID)
	}
	if bi.BuyerID != 325999999 {
		t.Errorf("buyer_id = %d", bi.BuyerID)
	}
}
