package mercadolibre

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// Time wraps time.Time to tolerate the date/time formats MercadoLibre returns
// (RFC3339 with millis and offset, and date-only values). A null/empty value
// decodes to the zero Time.
type Time struct{ time.Time }

var timeLayouts = []string{
	time.RFC3339Nano,
	time.RFC3339,
	"2006-01-02T15:04:05.000Z07:00",
	"2006-01-02T15:04:05Z07:00",
	"2006-01-02",
}

// UnmarshalJSON parses the supported layouts.
func (t *Time) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), `"`)
	if s == "" || s == "null" {
		t.Time = time.Time{}
		return nil
	}
	for _, l := range timeLayouts {
		if parsed, err := time.Parse(l, s); err == nil {
			t.Time = parsed
			return nil
		}
	}
	return fmt.Errorf("mercadolibre: cannot parse time %q", s)
}

// MarshalJSON renders RFC3339, or null for the zero value.
func (t Time) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte("null"), nil
	}
	return []byte(`"` + t.Format(time.RFC3339) + `"`), nil
}

// FlexID is an identifier that MercadoLibre returns inconsistently as either a
// JSON number or a JSON string (e.g. billing cust_id). It always decodes to its
// string form.
type FlexID string

// UnmarshalJSON accepts both quoted strings and bare numbers.
func (f *FlexID) UnmarshalJSON(b []byte) error {
	b = bytes.TrimSpace(b)
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	if b[0] == '"' {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*f = FlexID(s)
		return nil
	}
	*f = FlexID(b)
	return nil
}
