package auth

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"net/url"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestAuthCodeURL(t *testing.T) {
	cfg := Config{ClientID: "123", RedirectURI: "https://app.example/cb", Site: "MLB"}
	p, err := GeneratePKCE()
	if err != nil {
		t.Fatal(err)
	}
	raw := cfg.AuthCodeURL("xyz-state", WithPKCE(p), WithScope("offline_access", "read", "write"))
	u, err := url.Parse(raw)
	if err != nil {
		t.Fatal(err)
	}
	if u.Host != "auth.mercadolivre.com.br" { // Brazil = mercadoLIVRE
		t.Errorf("host = %q, want auth.mercadolivre.com.br", u.Host)
	}
	q := u.Query()
	for k, want := range map[string]string{
		"response_type":         "code",
		"client_id":             "123",
		"redirect_uri":          "https://app.example/cb",
		"state":                 "xyz-state",
		"code_challenge":        p.Challenge,
		"code_challenge_method": "S256",
		"scope":                 "offline_access read write",
	} {
		if q.Get(k) != want {
			t.Errorf("query %q = %q, want %q", k, q.Get(k), want)
		}
	}
}

func TestGeneratePKCE_S256(t *testing.T) {
	p, err := GeneratePKCE()
	if err != nil {
		t.Fatal(err)
	}
	sum := sha256.Sum256([]byte(p.Verifier))
	want := base64.RawURLEncoding.EncodeToString(sum[:])
	if p.Challenge != want {
		t.Errorf("challenge mismatch")
	}
	if p.Method != "S256" {
		t.Errorf("method = %q", p.Method)
	}
}

func stubTokenServer(t *testing.T, h http.HandlerFunc) {
	t.Helper()
	srv := httptest.NewServer(h)
	t.Cleanup(srv.Close)
	old := tokenURL
	tokenURL = srv.URL
	t.Cleanup(func() { tokenURL = old })
}

func TestExchange(t *testing.T) {
	stubTokenServer(t, func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		if r.Form.Get("grant_type") != "authorization_code" || r.Form.Get("code") != "the-code" {
			t.Errorf("unexpected form: %v", r.Form)
		}
		if r.Form.Get("code_verifier") != "verifier-123" {
			t.Errorf("missing code_verifier: %v", r.Form)
		}
		w.Write([]byte(`{"access_token":"AT","token_type":"bearer","expires_in":21600,"scope":"read write","user_id":7,"refresh_token":"RT1"}`))
	})

	tok, err := Config{ClientID: "1", ClientSecret: "s"}.Exchange(
		context.Background(), "the-code", WithCodeVerifier("verifier-123"))
	if err != nil {
		t.Fatal(err)
	}
	if tok.AccessToken != "AT" || tok.RefreshToken != "RT1" || tok.UserID != 7 {
		t.Errorf("unexpected token: %+v", tok)
	}
	if tok.Expiry.IsZero() || tok.Expiry.Before(time.Now()) {
		t.Errorf("expiry not set in the future: %v", tok.Expiry)
	}
}

func TestExchange_oauthError(t *testing.T) {
	stubTokenServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid_grant","error_description":"already used","status":400}`))
	})
	_, err := Config{}.Refresh(context.Background(), "RT")
	oerr, ok := err.(*Error)
	if !ok || !oerr.IsInvalidGrant() {
		t.Fatalf("expected invalid_grant *Error, got %T %v", err, err)
	}
}

func TestRefreshingTokenSource_refreshesAndPersists(t *testing.T) {
	var refreshCalls int32
	stubTokenServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&refreshCalls, 1)
		r.ParseForm()
		if r.Form.Get("grant_type") != "refresh_token" || r.Form.Get("refresh_token") != "RT-old" {
			t.Errorf("unexpected refresh form: %v", r.Form)
		}
		w.Write([]byte(`{"access_token":"AT-new","token_type":"bearer","expires_in":21600,"refresh_token":"RT-new"}`))
	})

	store := NewMemoryStore(nil)
	expired := &Token{AccessToken: "AT-old", RefreshToken: "RT-old", Expiry: time.Now().Add(-time.Hour)}
	ts := NewRefreshingTokenSource(Config{ClientID: "1", ClientSecret: "s"}, expired, store, 0)

	got, err := ts.Token(context.Background())
	if err != nil {
		t.Fatal(err)
	}
	if got != "AT-new" {
		t.Errorf("token = %q, want AT-new", got)
	}
	saved, _ := store.Load(context.Background())
	if saved == nil || saved.RefreshToken != "RT-new" {
		t.Errorf("rotated token not persisted: %+v", saved)
	}
}

func TestRefreshingTokenSource_singleFlight(t *testing.T) {
	var refreshCalls int32
	stubTokenServer(t, func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&refreshCalls, 1)
		time.Sleep(20 * time.Millisecond) // widen the race window
		w.Write([]byte(`{"access_token":"AT-new","token_type":"bearer","expires_in":21600,"refresh_token":"RT-new"}`))
	})

	expired := &Token{AccessToken: "AT-old", RefreshToken: "RT-old", Expiry: time.Now().Add(-time.Hour)}
	ts := NewRefreshingTokenSource(Config{ClientID: "1", ClientSecret: "s"}, expired, NewMemoryStore(nil), 0)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if _, err := ts.Token(context.Background()); err != nil {
				t.Error(err)
			}
		}()
	}
	wg.Wait()
	if got := atomic.LoadInt32(&refreshCalls); got != 1 {
		t.Errorf("refresh called %d times, want 1 (single-flight)", got)
	}
}

func TestStaticTokenSource(t *testing.T) {
	ts := StaticTokenSource("AT")
	got, err := ts.Token(context.Background())
	if err != nil || got != "AT" {
		t.Fatalf("got %q, %v", got, err)
	}
}
