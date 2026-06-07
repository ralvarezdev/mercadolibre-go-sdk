package mercadolibre

import (
	"context"
	"net/url"
	"strconv"
	"strings"
)

// User is a MercadoLibre user account (GET /users/{id} and /users/me).
//
// Models the commonly-used fields; the API returns more depending on the caller's
// relationship to the account (e.g. private fields only for /users/me with the
// owner's token). Reach for the generic Get[T] helper if you need a field not
// modeled here.
type User struct {
	ID               int64          `json:"id"`
	Nickname         string         `json:"nickname"`
	RegistrationDate Time           `json:"registration_date"`
	FirstName        string         `json:"first_name,omitempty"`
	LastName         string         `json:"last_name,omitempty"`
	CountryID        string         `json:"country_id"`
	Email            string         `json:"email,omitempty"`
	Identification   *Identification `json:"identification,omitempty"`
	Address          *UserAddress    `json:"address,omitempty"`
	Phone            *Phone          `json:"phone,omitempty"`
	UserType         string          `json:"user_type,omitempty"`
	Tags             []string        `json:"tags,omitempty"`
	Logo             *string         `json:"logo,omitempty"`
	Points           int             `json:"points,omitempty"`
	SiteID           string          `json:"site_id,omitempty"`
	Permalink        string          `json:"permalink,omitempty"`
	SellerExperience string          `json:"seller_experience,omitempty"`
	SellerReputation *SellerReputation `json:"seller_reputation,omitempty"`
}

// Identification is a government/tax identifier.
type Identification struct {
	Type   string `json:"type"`
	Number string `json:"number"`
}

// UserAddress is the short address embedded in a User.
type UserAddress struct {
	State   string `json:"state"`
	City    string `json:"city"`
	Address string `json:"address"`
	ZipCode string `json:"zip_code"`
}

// Phone is a phone number with optional verification flag.
type Phone struct {
	AreaCode  string `json:"area_code"`
	Number    string `json:"number"`
	Extension string `json:"extension"`
	Verified  bool   `json:"verified,omitempty"`
}

// SellerReputation summarizes a seller's standing.
type SellerReputation struct {
	LevelID           *string            `json:"level_id"`
	PowerSellerStatus *string            `json:"power_seller_status"`
	RealLevel         string             `json:"real_level,omitempty"`
	ProtectionEndDate Time               `json:"protection_end_date,omitempty"`
	Transactions      *Transactions      `json:"transactions,omitempty"`
	Metrics           *ReputationMetrics `json:"metrics,omitempty"`
}

// Transactions holds aggregate transaction counts and ratings.
type Transactions struct {
	Period    string   `json:"period"`
	Total     int      `json:"total"`
	Completed int      `json:"completed"`
	Canceled  int      `json:"canceled"`
	Ratings   *Ratings `json:"ratings,omitempty"`
}

// Ratings holds the proportion of positive/neutral/negative ratings.
type Ratings struct {
	Positive float64 `json:"positive"`
	Neutral  float64 `json:"neutral"`
	Negative float64 `json:"negative"`
}

// ReputationMetrics holds the seller's recent-performance metrics.
type ReputationMetrics struct {
	Sales               *MetricEntry `json:"sales,omitempty"`
	Claims              *MetricEntry `json:"claims,omitempty"`
	DelayedHandlingTime *MetricEntry `json:"delayed_handling_time,omitempty"`
	Cancellations       *MetricEntry `json:"cancellations,omitempty"`
}

// MetricEntry is one reputation metric over a period. Completed is used by the
// "sales" metric; Rate/Value by the others. Excluded carries the real values
// for protected sellers.
type MetricEntry struct {
	Period    string          `json:"period,omitempty"`
	Completed int             `json:"completed,omitempty"`
	Rate      float64         `json:"rate,omitempty"`
	Value     float64         `json:"value,omitempty"`
	Excluded  *MetricExcluded `json:"excluded,omitempty"`
}

// MetricExcluded holds the real metric values shown to protected sellers.
type MetricExcluded struct {
	RealValue float64 `json:"real_value"`
	RealRate  float64 `json:"real_rate"`
}

// Address is a full address from GET /users/{id}/addresses.
type Address struct {
	ID           int64   `json:"id"`
	UserID       int64   `json:"user_id"`
	AddressLine  string  `json:"address_line"`
	StreetName   string  `json:"street_name"`
	StreetNumber string  `json:"street_number"`
	ZipCode      string  `json:"zip_code"`
	City         IDName  `json:"city"`
	State        IDName  `json:"state"`
	Country      IDName  `json:"country"`
	Latitude     float64 `json:"latitude,omitempty"`
	Longitude    float64 `json:"longitude,omitempty"`
	Status       string  `json:"status,omitempty"`
}

// UsersService accesses user accounts and their related data.
type UsersService struct{ c *Client }

// Me returns the authenticated user (GET /users/me). Requires a token.
func (s *UsersService) Me(ctx context.Context) (*User, error) {
	return Get[*User](ctx, s.c, EPUsersMe)
}

// Get returns a user by ID (GET /users/{id}).
func (s *UsersService) Get(ctx context.Context, userID int64) (*User, error) {
	return Get[*User](ctx, s.c, EPUserByID, userID)
}

// GetMulti returns several users in one call (GET /users?ids=...).
func (s *UsersService) GetMulti(ctx context.Context, userIDs ...int64) ([]User, error) {
	ids := make([]string, len(userIDs))
	for i, id := range userIDs {
		ids[i] = strconv.FormatInt(id, 10)
	}
	q := url.Values{"ids": {strings.Join(ids, ",")}}
	return GetQ[[]User](ctx, s.c, EPUsersMultiget, q)
}

// Addresses returns the addresses registered for a user
// (GET /users/{id}/addresses).
func (s *UsersService) Addresses(ctx context.Context, userID int64) ([]Address, error) {
	return Get[[]Address](ctx, s.c, EPUserAddresses, userID)
}

// PaymentMethod is one payment method a user has enabled.
type PaymentMethod struct {
	ID            string `json:"id"`
	Name          string `json:"name,omitempty"`
	PaymentTypeID string `json:"payment_type_id,omitempty"`
	Thumbnail     string `json:"thumbnail,omitempty"`
	SecureThumb   string `json:"secure_thumbnail,omitempty"`
}

// AcceptedPaymentMethods returns the payment methods a user accepts
// (GET /users/{id}/accepted_payment_methods).
func (s *UsersService) AcceptedPaymentMethods(ctx context.Context, userID int64) ([]PaymentMethod, error) {
	return Get[[]PaymentMethod](ctx, s.c, EPUserPayMethods, userID)
}

// Brand is an official store brand associated with a user.
type Brand struct {
	ID              int64   `json:"id"`
	Name            string  `json:"name"`
	Permalink       string  `json:"permalink,omitempty"`
	Logo            *string `json:"logo,omitempty"`
	OfficialStoreID int64   `json:"official_store_id,omitempty"`
}

// Brands returns the brands (official stores) associated with a user
// (GET /users/{id}/brands).
func (s *UsersService) Brands(ctx context.Context, userID int64) ([]Brand, error) {
	return Get[[]Brand](ctx, s.c, EPUserBrands, userID)
}
