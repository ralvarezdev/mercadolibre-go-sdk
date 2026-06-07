package mercadolibre

import (
	"context"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// userFixture is a trimmed real /users/{id} response body.
const userFixture = `{
  "id": 206946886,
  "nickname": "TETE6838590",
  "registration_date": "2016-02-24T15:18:42.000-04:00",
  "first_name": "Pedro",
  "last_name": "Picapiedras",
  "country_id": "AR",
  "email": "test_user_15879541@testuser.com",
  "identification": { "type": "DNI", "number": "33333333" },
  "address": { "state": "AR-C", "city": "CapitalFederal", "address": "Triunvirato5555", "zip_code": "1414" },
  "user_type": "normal",
  "tags": ["normal", "test_user"],
  "logo": null,
  "points": 100,
  "site_id": "MLA",
  "seller_reputation": { "level_id": null, "power_seller_status": null,
    "transactions": { "period": "historic", "total": 0, "completed": 0, "canceled": 0 } }
}`

func newTestClient(t *testing.T, h http.HandlerFunc, opts ...Option) *Client {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	base := []Option{WithBaseURL(srv.URL), WithStaticToken("APP_USR-test")}
	return NewClient(append(base, opts...)...)
}

func TestUsersGet_decodesAndAuthenticates(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Authorization"); got != "Bearer APP_USR-test" {
			t.Errorf("Authorization = %q, want bearer token", got)
		}
		if r.URL.Path != "/users/206946886" {
			t.Errorf("path = %q", r.URL.Path)
		}
		w.Write([]byte(userFixture))
	})

	u, err := c.Users.Get(context.Background(), 206946886)
	if err != nil {
		t.Fatal(err)
	}
	if u.Nickname != "TETE6838590" || u.SiteID != "MLA" {
		t.Errorf("unexpected user: %+v", u)
	}
	if u.Identification == nil || u.Identification.Number != "33333333" {
		t.Errorf("identification not decoded: %+v", u.Identification)
	}
	if u.RegistrationDate.Year() != 2016 {
		t.Errorf("registration date not parsed: %v", u.RegistrationDate)
	}
	if u.SellerReputation == nil || u.SellerReputation.Transactions.Period != "historic" {
		t.Errorf("seller reputation not decoded: %+v", u.SellerReputation)
	}
}

func TestMe_usesMeEndpoint(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/users/me" {
			t.Errorf("path = %q, want /users/me", r.URL.Path)
		}
		w.Write([]byte(userFixture))
	})
	if _, err := c.Users.Me(context.Background()); err != nil {
		t.Fatal(err)
	}
}

func TestAPIError_forbidden(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(`{"status":403,"error":"Invalid scopes","code":"FORBIDDEN"}`))
	})

	_, err := c.Users.Me(context.Background())
	apiErr, ok := IsAPIError(err)
	if !ok {
		t.Fatalf("expected *APIError, got %T: %v", err, err)
	}
	if !apiErr.IsForbidden() || apiErr.Code != "FORBIDDEN" || apiErr.ErrName != "Invalid scopes" {
		t.Errorf("unexpected APIError: %+v", apiErr)
	}
}

func TestRetry_on429ThenSuccess(t *testing.T) {
	var calls int32
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		if atomic.AddInt32(&calls, 1) == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			w.Write([]byte(`{"status":429,"error":"local_rate_limited"}`))
			return
		}
		w.Write([]byte(userFixture))
	}, WithRetryPolicy(RetryPolicy{MaxAttempts: 3, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}))

	if _, err := c.Users.Me(context.Background()); err != nil {
		t.Fatalf("expected success after retry, got %v", err)
	}
	if got := atomic.LoadInt32(&calls); got != 2 {
		t.Errorf("calls = %d, want 2 (one 429 then success)", got)
	}
}

func TestRetry_exhausts(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"status":503,"error":"unavailable"}`))
	}, WithRetryPolicy(RetryPolicy{MaxAttempts: 2, BaseDelay: time.Millisecond, MaxDelay: time.Millisecond}))

	_, err := c.Users.Me(context.Background())
	apiErr, ok := IsAPIError(err)
	if !ok || apiErr.StatusCode != http.StatusServiceUnavailable {
		t.Fatalf("expected 503 APIError, got %v", err)
	}
}

func TestGenericGet_escapeHatch(t *testing.T) {
	c := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"site_id":"MLA","id":"PAY1","name":"Visa"}`))
	})
	type payMethod struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
	pm, err := Get[payMethod](context.Background(), c, "/anything")
	if err != nil {
		t.Fatal(err)
	}
	if pm.ID != "PAY1" || pm.Name != "Visa" {
		t.Errorf("generic Get decoded wrong: %+v", pm)
	}
}
