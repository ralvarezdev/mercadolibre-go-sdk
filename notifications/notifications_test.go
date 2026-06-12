package notifications

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	n, err := Parse([]byte(`{"resource":"/orders/123","user_id":465432224,"topic":"orders",
	  "application_id":999,"attempts":1,"sent":"2019-10-17T17:15:30.279Z","received":"2019-10-17T17:15:30.259Z"}`))
	if err != nil {
		t.Fatal(err)
	}
	if n.Topic != TopicOrders || n.Resource != "/orders/123" || n.UserID != 465432224 {
		t.Errorf("bad notification: %+v", n)
	}
}

func TestHandler(t *testing.T) {
	var got Notification
	h := Handler(func(r *http.Request, n Notification) error {
		got = n
		return nil
	})
	srv := httptest.NewServer(h)
	defer srv.Close()

	body := `{"resource":"/items/MLA1","user_id":1,"topic":"items","application_id":2,"attempts":1}`
	req, err := http.NewRequestWithContext(context.Background(), http.MethodPost, srv.URL, strings.NewReader(body))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want 200", resp.StatusCode)
	}
	if got.Topic != TopicItems || got.Resource != "/items/MLA1" {
		t.Errorf("handler got %+v", got)
	}
}

func TestHandler_rejectsGET(t *testing.T) {
	srv := httptest.NewServer(Handler(func(*http.Request, Notification) error { return nil }))
	defer srv.Close()
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, srv.URL, http.NoBody)
	if err != nil {
		t.Fatal(err)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusMethodNotAllowed {
		t.Errorf("status = %d, want 405", resp.StatusCode)
	}
}
