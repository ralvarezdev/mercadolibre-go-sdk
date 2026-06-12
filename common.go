package mercadolibre

type (
	// MultiGetResult wraps one entry of a multiget (verbose) response, where each
	// element carries its own HTTP status code alongside the body.
	MultiGetResult[T any] struct {
		Body T   `json:"body"`
		Code int `json:"code"`
	}

	// AttributeValue is an {id, name} option, sometimes with a struct value.
	AttributeValue struct {
		Struct any    `json:"struct,omitempty"`
		ID     string `json:"id"`
		Name   string `json:"name"`
	}

	// Attribute is used both on items (value_id/value_name set) and on category /
	// domain technical specs (value_type, values, tags set). Fields not present in a
	// given context are simply empty.
	Attribute struct {
		Tags               any              `json:"tags,omitempty"`
		ValueStruct        any              `json:"value_struct,omitempty"`
		ValueID            *string          `json:"value_id,omitempty"`
		ValueName          *string          `json:"value_name,omitempty"`
		ID                 string           `json:"id"`
		Name               string           `json:"name,omitempty"`
		ValueType          string           `json:"value_type,omitempty"`
		AttributeGroupID   string           `json:"attribute_group_id,omitempty"`
		AttributeGroupName string           `json:"attribute_group_name,omitempty"`
		Values             []AttributeValue `json:"values,omitempty"`
	}

	// SaleTerm is a sale condition such as warranty type/time.
	SaleTerm struct {
		ValueID   *string `json:"value_id,omitempty"`
		ID        string  `json:"id"`
		Name      string  `json:"name,omitempty"`
		ValueName string  `json:"value_name,omitempty"`
	}

	// Picture is an item image in its various sizes.
	Picture struct {
		ID        string `json:"id,omitempty"`
		URL       string `json:"url,omitempty"`
		SecureURL string `json:"secure_url,omitempty"`
		Source    string `json:"source,omitempty"` // used on input when creating items
		Size      string `json:"size,omitempty"`
		MaxSize   string `json:"max_size,omitempty"`
		Quality   string `json:"quality,omitempty"`
	}

	// Shipping describes an item's shipping configuration.
	Shipping struct {
		Dimensions   *string  `json:"dimensions,omitempty"`
		Mode         string   `json:"mode,omitempty"`
		LogisticType string   `json:"logistic_type,omitempty"`
		Methods      []any    `json:"methods,omitempty"`
		Tags         []string `json:"tags,omitempty"`
		LocalPickUp  bool     `json:"local_pick_up,omitempty"`
		FreeShipping bool     `json:"free_shipping,omitempty"`
		StorePickUp  bool     `json:"store_pick_up,omitempty"`
	}

	// GeoLocation is a latitude/longitude pair.
	GeoLocation struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}

	// SearchLocation holds normalized neighborhood/city/state for an address.
	SearchLocation struct {
		Neighborhood *IDName `json:"neighborhood,omitempty"`
		City         *IDName `json:"city,omitempty"`
		State        *IDName `json:"state,omitempty"`
	}

	// SellerAddress is the address attached to an item.
	SellerAddress struct {
		SearchLocation *SearchLocation `json:"search_location,omitempty"`
		Country        *IDName         `json:"country,omitempty"`
		State          *IDName         `json:"state,omitempty"`
		City           *IDName         `json:"city,omitempty"`
		Comment        string          `json:"comment,omitempty"`
		AddressLine    string          `json:"address_line,omitempty"`
		ZipCode        string          `json:"zip_code,omitempty"`
		Latitude       float64         `json:"latitude,omitempty"`
		Longitude      float64         `json:"longitude,omitempty"`
		ID             int64           `json:"id,omitempty"`
	}

	// Filter is an applied or available search filter.
	Filter struct {
		Type   string        `json:"type,omitempty"`
		ID     string        `json:"id"`
		Name   string        `json:"name"`
		Values []FilterValue `json:"values"`
	}

	// FilterValue is one option within a Filter; Results is the match count for
	// available filters.
	FilterValue struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Results int    `json:"results,omitempty"`
	}

	// Sort is an applied or available sort option.
	Sort struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
