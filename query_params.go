package mercadolibre

// QueryParam is a query parameter key type.
type QueryParam string

// Query parameter keys for API requests.
// These constants should be used instead of string literals when building url.Values.

// Auth/OAuth parameters
const (
	QueryParamCodeChallenge       QueryParam = "code_challenge"
	QueryParamCodeChallengeMethod QueryParam = "code_challenge_method"
	QueryParamScope               QueryParam = "scope"
	QueryParamResponseType        QueryParam = "response_type"
	QueryParamClientID            QueryParam = "client_id"
	QueryParamRedirectURI         QueryParam = "redirect_uri"
	QueryParamState               QueryParam = "state"
	QueryParamCodeVerifier        QueryParam = "code_verifier"
	QueryParamGrantType           QueryParam = "grant_type"
	QueryParamClientSecret        QueryParam = "client_secret"
	QueryParamCode                QueryParam = "code"
	QueryParamRefreshToken        QueryParam = "refresh_token"
)

// Items/Search parameters
const (
	QueryParamIDs        QueryParam = "ids"
	QueryParamAttributes QueryParam = "attributes"
	QueryParamQ          QueryParam = "q"
	QueryParamPrice      QueryParam = "price"
	QueryParamCategoryID QueryParam = "category_id"
)

// Pagination parameters
const (
	QueryParamOffset QueryParam = "offset"
	QueryParamLimit  QueryParam = "limit"
)

// Metrics/Visits parameters
const (
	QueryParamDateFrom QueryParam = "date_from"
	QueryParamDateTo   QueryParam = "date_to"
	QueryParamLast     QueryParam = "last"
	QueryParamUnit     QueryParam = "unit"
	QueryParamEnding   QueryParam = "ending"
)

// Notification parameters
const (
	QueryParamAppID QueryParam = "app_id"
	QueryParamTopic QueryParam = "topic"
)

// Messaging parameters
const (
	QueryParamTag        QueryParam = "tag"
	QueryParamMarkAsRead QueryParam = "mark_as_read"
)

// Currency parameters
const (
	QueryParamFrom QueryParam = "from"
	QueryParamTo   QueryParam = "to"
)

// Additional search/list parameters
const (
	QueryParamUserID        QueryParam = "user_id"
	QueryParamVersion       QueryParam = "version"
	QueryParamStatus        QueryParam = "status"
	QueryParamSearchType    QueryParam = "search_type"
	QueryParamScrollID      QueryParam = "scroll_id"
	QueryParamSeller        QueryParam = "seller"
	QueryParamPromotionType QueryParam = "promotion_type"
	QueryParamAPIVersion    QueryParam = "api_version"
)

// Nested field filters (for complex queries)
const (
	QueryParamPlayersUserID QueryParam = "players.user_id"
	QueryParamPlayersRole   QueryParam = "players.role"
)

// Shipment and user parameters
const (
	QueryParamShipmentIDs QueryParam = "shipment_ids"
)
