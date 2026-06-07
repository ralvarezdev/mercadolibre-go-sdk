package mercadolibre

// MultiGetResult wraps one entry of a multiget (verbose) response, where each
// element carries its own HTTP status code alongside the body.
type MultiGetResult[T any] struct {
	Code int `json:"code"`
	Body T   `json:"body"`
}

// AttributeValue is an {id, name} option, sometimes with a struct value.
type AttributeValue struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Struct any    `json:"struct,omitempty"`
}

// Attribute is used both on items (value_id/value_name set) and on category /
// domain technical specs (value_type, values, tags set). Fields not present in a
// given context are simply empty.
type Attribute struct {
	ID            string           `json:"id"`
	Name          string           `json:"name,omitempty"`
	ValueID       *string          `json:"value_id,omitempty"`
	ValueName     *string          `json:"value_name,omitempty"`
	ValueStruct   any              `json:"value_struct,omitempty"`
	ValueType     string           `json:"value_type,omitempty"`
	Values        []AttributeValue `json:"values,omitempty"`
	Tags          any              `json:"tags,omitempty"`
	AttributeGroupID   string      `json:"attribute_group_id,omitempty"`
	AttributeGroupName string      `json:"attribute_group_name,omitempty"`
}

// SaleTerm is a sale condition such as warranty type/time.
type SaleTerm struct {
	ID        string  `json:"id"`
	Name      string  `json:"name,omitempty"`
	ValueID   *string `json:"value_id,omitempty"`
	ValueName string  `json:"value_name,omitempty"`
}

// Picture is an item image in its various sizes.
type Picture struct {
	ID        string `json:"id,omitempty"`
	URL       string `json:"url,omitempty"`
	SecureURL string `json:"secure_url,omitempty"`
	Source    string `json:"source,omitempty"` // used on input when creating items
	Size      string `json:"size,omitempty"`
	MaxSize   string `json:"max_size,omitempty"`
	Quality   string `json:"quality,omitempty"`
}

// Shipping describes an item's shipping configuration.
type Shipping struct {
	Mode                string         `json:"mode,omitempty"`
	LocalPickUp         bool           `json:"local_pick_up,omitempty"`
	FreeShipping        bool           `json:"free_shipping,omitempty"`
	StorePickUp         bool           `json:"store_pick_up,omitempty"`
	LogisticType        string         `json:"logistic_type,omitempty"`
	Methods             []any          `json:"methods,omitempty"`
	Dimensions          *string        `json:"dimensions,omitempty"`
	Tags                []string       `json:"tags,omitempty"`
}

// GeoLocation is a latitude/longitude pair.
type GeoLocation struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// SearchLocation holds normalized neighborhood/city/state for an address.
type SearchLocation struct {
	Neighborhood *IDName `json:"neighborhood,omitempty"`
	City         *IDName `json:"city,omitempty"`
	State        *IDName `json:"state,omitempty"`
}

// SellerAddress is the address attached to an item.
type SellerAddress struct {
	ID             int64           `json:"id,omitempty"`
	Comment        string          `json:"comment,omitempty"`
	AddressLine    string          `json:"address_line,omitempty"`
	ZipCode        string          `json:"zip_code,omitempty"`
	City           *IDName         `json:"city,omitempty"`
	State          *IDName         `json:"state,omitempty"`
	Country        *IDName         `json:"country,omitempty"`
	Latitude       float64         `json:"latitude,omitempty"`
	Longitude      float64         `json:"longitude,omitempty"`
	SearchLocation *SearchLocation `json:"search_location,omitempty"`
}

// Filter is an applied or available search filter.
type Filter struct {
	ID     string        `json:"id"`
	Name   string        `json:"name"`
	Type   string        `json:"type,omitempty"`
	Values []FilterValue `json:"values"`
}

// FilterValue is one option within a Filter; Results is the match count for
// available filters.
type FilterValue struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Results int    `json:"results,omitempty"`
}

// Sort is an applied or available sort option.
type Sort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
