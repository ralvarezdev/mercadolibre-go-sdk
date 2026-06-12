package mercadolibre

import "context"

type (
	// Category is a category node (GET /categories/{id}).
	Category struct {
		Settings                 *CategorySettings `json:"settings,omitempty"`
		ID                       string            `json:"id"`
		Name                     string            `json:"name"`
		Picture                  string            `json:"picture,omitempty"`
		Permalink                string            `json:"permalink,omitempty"`
		AttributeTypes           string            `json:"attribute_types,omitempty"`
		PathFromRoot             []IDName          `json:"path_from_root,omitempty"`
		ChildrenCategories       []ChildCategory   `json:"children_categories,omitempty"`
		TotalItemsInThisCategory int               `json:"total_items_in_this_category,omitempty"`
	}

	// ChildCategory is a direct child reference of a Category.
	ChildCategory struct {
		ID                       string `json:"id"`
		Name                     string `json:"name"`
		TotalItemsInThisCategory int    `json:"total_items_in_this_category,omitempty"`
	}

	// CategorySettings holds listing rules for a category (subset of fields).
	CategorySettings struct {
		Status             string   `json:"status,omitempty"`
		BuyingModes        []string `json:"buying_modes,omitempty"`
		ItemConditions     []string `json:"item_conditions,omitempty"`
		Currencies         []string `json:"currencies,omitempty"`
		MaxPicturesPerItem int      `json:"max_pictures_per_item,omitempty"`
		ListingAllowed     bool     `json:"listing_allowed,omitempty"`
	}

	// CategoriesService accesses categories, their attributes and sale terms.
	CategoriesService struct{ c *Client }

	// Domain is a product domain (GET /domains/{id}).
	Domain struct {
		ID         string      `json:"id"`
		Name       string      `json:"name"`
		Attributes []Attribute `json:"attributes,omitempty"`
	}

	// DomainsService accesses product domains and their technical specs.
	DomainsService struct{ c *Client }
)

// SiteList returns the top-level categories of a site
// (GET /sites/{id}/categories).
func (s *CategoriesService) SiteList(ctx context.Context, siteID string) ([]IDName, error) {
	return Get[[]IDName](ctx, s.c, EPSiteCategories, siteID)
}

// Get returns a category by ID (GET /categories/{id}).
func (s *CategoriesService) Get(ctx context.Context, categoryID string) (*Category, error) {
	return Get[*Category](ctx, s.c, EPCategoryByID, categoryID)
}

// Attributes returns the attributes/technical specs of a category
// (GET /categories/{id}/attributes).
func (s *CategoriesService) Attributes(ctx context.Context, categoryID string) ([]Attribute, error) {
	return Get[[]Attribute](ctx, s.c, EPCategoryAttributes, categoryID)
}

// SaleTerms returns the sale terms applicable to a category
// (GET /categories/{id}/sale_terms).
func (s *CategoriesService) SaleTerms(ctx context.Context, categoryID string) ([]Attribute, error) {
	return Get[[]Attribute](ctx, s.c, EPCategorySaleTerms, categoryID)
}

// Get returns a domain by ID (GET /domains/{id}).
func (s *DomainsService) Get(ctx context.Context, domainID string) (*Domain, error) {
	return Get[*Domain](ctx, s.c, EPDomainByID, domainID)
}

// TechnicalSpecs returns the technical specs of a domain
// (GET /domains/{id}/technical_specs).
func (s *DomainsService) TechnicalSpecs(ctx context.Context, domainID string) (map[string]any, error) {
	return Get[map[string]any](ctx, s.c, EPDomainTechnicalSpecs, domainID)
}
