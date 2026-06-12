package mercadolibre

import "fmt"

// Endpoint is a typed API path template. Path parameters use fmt verbs (%s, %d)
// and are filled with Path / the variadic args of the generic helpers.
//
// Endpoints are relative to the client BaseURL (https://api.mercadolibre.com).
type Endpoint string

// Path substitutes args into the endpoint template.
//
//	EPItemByID.Path("MLA123") // "/items/MLA123"
func (e Endpoint) Path(args ...any) string {
	if len(args) == 0 {
		return string(e)
	}
	return fmt.Sprintf(string(e), args...)
}

// Endpoint constants. This registry grows as resources are typed; it is seeded
// from the endpoints captured in docs/. Keep grouped by resource.
const (
	// EPOAuthToken is the OAuth token endpoint (also defined locally in the auth package to avoid an import cycle).
	EPOAuthToken Endpoint = "/oauth/token"

	// EPSites and related endpoints for sites, currencies, locations.
	EPSites                 Endpoint = "/sites"
	EPSiteByID              Endpoint = "/sites/%s"
	EPSiteCategories        Endpoint = "/sites/%s/categories"
	EPSiteSearch            Endpoint = "/sites/%s/search"
	EPSiteListingTypes      Endpoint = "/sites/%s/listing_types"
	EPSiteListingTypeByID   Endpoint = "/sites/%s/listing_types/%s"
	EPSiteListingPrices     Endpoint = "/sites/%s/listing_prices"
	EPSiteCategoryPredictor Endpoint = "/sites/%s/category_predictor/predict"

	// EPCurrencies and related currency endpoints.
	EPCurrencies          Endpoint = "/currencies"
	EPCurrencyByID        Endpoint = "/currencies/%s"
	EPCurrencyConversions Endpoint = "/currency_conversions/search" // ?from=&to=

	// EPCountries and country endpoints.
	EPCountries   Endpoint = "/countries"
	EPCountryByID Endpoint = "/countries/%s"

	// EPClassifiedCountries and geographic tree endpoints.
	EPClassifiedCountries   Endpoint = "/classified_locations/countries"
	EPClassifiedCountryByID Endpoint = "/classified_locations/countries/%s"
	EPClassifiedStateByID   Endpoint = "/classified_locations/states/%s"
	EPClassifiedCityByID    Endpoint = "/classified_locations/cities/%s"

	// EPUsersMe and related user endpoints.
	EPUsersMe         Endpoint = "/users/me"
	EPUserByID        Endpoint = "/users/%d"
	EPUsersMultiget   Endpoint = "/users" // ?ids=
	EPUserAddresses   Endpoint = "/users/%d/addresses"
	EPUserPayMethods  Endpoint = "/users/%d/accepted_payment_methods"
	EPUserBrands      Endpoint = "/users/%d/brands"
	EPUserItemsSearch Endpoint = "/users/%d/items/search"
	EPUserAppByID     Endpoint = "/users/%d/applications/%d"

	// EPApplicationByID is the applications endpoint.
	EPApplicationByID Endpoint = "/applications/%d"

	// EPItems and related item endpoints.
	EPItems                       Endpoint = "/items" // ?ids= for multiget; POST to create
	EPItemByID                    Endpoint = "/items/%s"
	EPItemDescription             Endpoint = "/items/%s/description"
	EPItemVariations              Endpoint = "/items/%s/variations"
	EPItemVariationByID           Endpoint = "/items/%s/variations/%d"
	EPUserItemsSearchRestrictions Endpoint = "/users/%d/items/search/restrictions"

	// EPPictures is the pictures endpoint.
	EPPictures    Endpoint = "/pictures"    // POST to upload by URL
	EPPictureByID Endpoint = "/pictures/%s" // GET a single picture

	// EPCategoryByID and related category and domain endpoints.
	EPCategoryByID         Endpoint = "/categories/%s"
	EPCategoryAttributes   Endpoint = "/categories/%s/attributes"
	EPCategorySaleTerms    Endpoint = "/categories/%s/sale_terms"
	EPDomainByID           Endpoint = "/domains/%s"
	EPDomainTechnicalSpecs Endpoint = "/domains/%s/technical_specs"

	// EPQuestions and related question endpoints.
	EPQuestions                Endpoint = "/questions"
	EPQuestionsSearch          Endpoint = "/questions/search"
	EPQuestionByID             Endpoint = "/questions/%d"
	EPAnswers                  Endpoint = "/answers"
	EPUserQuestionResponseTime Endpoint = "/users/%d/questions/response_time"

	// EPOrdersSearch and related order endpoints.
	EPOrdersSearch     Endpoint = "/orders/search"
	EPOrderByID        Endpoint = "/orders/%d"
	EPOrderFeedback    Endpoint = "/orders/%d/feedback"
	EPOrderNotes       Endpoint = "/orders/%d/notes"
	EPOrderProduct     Endpoint = "/orders/%d/product"
	EPOrderDiscounts   Endpoint = "/orders/%d/discounts"
	EPOrderBillingInfo Endpoint = "/orders/billing-info/%s/%s" // site, billing_info_id

	// EPShipmentByID and related shipment endpoints.
	EPShipmentByID     Endpoint = "/shipments/%d"
	EPShipmentItems    Endpoint = "/shipments/%d/items"
	EPShipmentCosts    Endpoint = "/shipments/%d/costs"
	EPShipmentPayments Endpoint = "/shipments/%d/payments"
	EPShipmentHistory  Endpoint = "/shipments/%d/history"
	EPShipmentCarrier  Endpoint = "/shipments/%d/carrier"
	EPShipmentSLA      Endpoint = "/shipments/%d/sla"
	EPShipmentLeadTime Endpoint = "/shipments/%d/lead_time"
	EPShipmentLabels   Endpoint = "/shipment_labels"

	// EPMissedFeeds is the notifications endpoint.
	EPMissedFeeds Endpoint = "/missed_feeds" // ?app_id=&topic=&offset=&limit=

	// EPMessagesPackSeller and related messaging endpoints.
	EPMessagesPackSeller Endpoint = "/messages/packs/%d/sellers/%d"
	EPMessageByID        Endpoint = "/messages/%s"
	EPMessageAttachments Endpoint = "/messages/attachments"
	EPMessageAttachByID  Endpoint = "/messages/attachments/%s"

	// EPClaimByID and related claims endpoints.
	EPClaimByID    Endpoint = "/post-purchase/v1/claims/%d"
	EPClaimsSearch Endpoint = "/post-purchase/v1/claims/search"
	EPClaimActions Endpoint = "/post-purchase/v1/claims/%d/actions"
	EPClaimHistory Endpoint = "/post-purchase/v1/claims/%d/history"

	// EPUserItemsVisits and related metrics endpoints.
	EPUserItemsVisits           Endpoint = "/users/%d/items_visits"             // ?date_from=&date_to=
	EPUserItemsVisitsTimeWindow Endpoint = "/users/%d/items_visits/time_window" // ?last=&unit=&ending=
	EPItemsVisits               Endpoint = "/items/visits"                      // ?ids=&date_from=&date_to=
	EPItemVisitsTimeWindow      Endpoint = "/items/%s/visits/time_window"       // ?last=&unit=&ending=

	// EPTrendsSite is the trends endpoint.
	EPTrendsSite     Endpoint = "/trends/%s"
	EPTrendsCategory Endpoint = "/trends/%s/%s"

	// EPSellerPromotionsUser and related promotion endpoints.
	EPSellerPromotionsUser     Endpoint = "/seller-promotions/users/%d"
	EPSellerPromotionByID      Endpoint = "/seller-promotions/promotions/%s"
	EPSellerPromotionItems     Endpoint = "/seller-promotions/promotions/%s/items"
	EPSellerPromotionItemByID  Endpoint = "/seller-promotions/promotions/%s/items/%s"
	EPSellerPromotionByItem    Endpoint = "/seller-promotions/items/%s"
	EPSellerPromotionOffers    Endpoint = "/seller-promotions/offers/%s"
	EPSellerPromotionCandidate Endpoint = "/seller-promotions/candidates/%s"

	// EPBillingReports is the billing reports endpoint.
	EPBillingReports  Endpoint = "/billing/integration/reports/%s"             // ?user_id=&version=
	EPBillingDownload Endpoint = "/billing/integration/reports/%s/download/%s" // report_type, doc_id
)
