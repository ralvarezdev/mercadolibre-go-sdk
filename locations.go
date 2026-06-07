package mercadolibre

import "context"

// Country is a country in the classified-locations tree. The list endpoint
// returns id/name only; the detail endpoint also includes states.
type Country struct {
	ID     string   `json:"id"`
	Name   string   `json:"name"`
	States []IDName `json:"states,omitempty"`
}

// State is a state/province with its cities (GET /classified_locations/states/{id}).
type State struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	Country        *IDName         `json:"country,omitempty"`
	Cities         []IDName        `json:"cities,omitempty"`
	GeoInformation *GeoInformation `json:"geo_information,omitempty"`
}

// City is a city (GET /classified_locations/cities/{id}).
type City struct {
	ID             string          `json:"id"`
	Name           string          `json:"name"`
	State          *IDName         `json:"state,omitempty"`
	Country        *IDName         `json:"country,omitempty"`
	GeoInformation *GeoInformation `json:"geo_information,omitempty"`
}

// GeoInformation carries optional geographic metadata for a location.
type GeoInformation struct {
	Location *GeoLocation `json:"location,omitempty"`
}

// LocationsService accesses the classified-locations geographic tree.
type LocationsService struct{ c *Client }

// Countries lists countries (GET /classified_locations/countries).
func (s *LocationsService) Countries(ctx context.Context) ([]Country, error) {
	return Get[[]Country](ctx, s.c, EPClassifiedCountries)
}

// Country returns a country with its states
// (GET /classified_locations/countries/{id}).
func (s *LocationsService) Country(ctx context.Context, countryID string) (*Country, error) {
	return Get[*Country](ctx, s.c, EPClassifiedCountryByID, countryID)
}

// State returns a state with its cities
// (GET /classified_locations/states/{id}).
func (s *LocationsService) State(ctx context.Context, stateID string) (*State, error) {
	return Get[*State](ctx, s.c, EPClassifiedStateByID, stateID)
}

// City returns a city (GET /classified_locations/cities/{id}).
func (s *LocationsService) City(ctx context.Context, cityID string) (*City, error) {
	return Get[*City](ctx, s.c, EPClassifiedCityByID, cityID)
}
