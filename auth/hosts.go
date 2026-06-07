package auth

// SiteID is an alias for the MercadoLibre site identifier string.
// Using an alias (not a distinct type) keeps it assignment-compatible with
// plain string, so existing callers need no changes.
type SiteID = string

// ISO 3166-1 alpha-2 country-code constants for every MercadoLibre site.
// Each constant holds the site ID used in API calls (MLA, MLB, …).
const (
	AR SiteID = "MLA" // Argentina
	BR SiteID = "MLB" // Brazil (mercadolivre.com.br)
	MX SiteID = "MLM" // Mexico
	CL SiteID = "MLC" // Chile
	CO SiteID = "MCO" // Colombia
	UY SiteID = "MLU" // Uruguay
	PE SiteID = "MPE" // Peru
	VE SiteID = "MLV" // Venezuela
	EC SiteID = "MEC" // Ecuador
	BO SiteID = "MBO" // Bolivia
	PY SiteID = "MPY" // Paraguay
	GT SiteID = "MGT" // Guatemala
	PA SiteID = "MPA" // Panama
	DO SiteID = "MRD" // Dominican Republic
	CR SiteID = "MCR" // Costa Rica
	HN SiteID = "MHN" // Honduras
	SV SiteID = "MSV" // El Salvador
	NI SiteID = "MNI" // Nicaragua
)

// authHosts maps a site ID to its OAuth authorization host.
// Brazil uses "mercadolivre" (Portuguese spelling).
var authHosts = map[SiteID]string{
	AR: "auth.mercadolibre.com.ar",
	BR: "auth.mercadolivre.com.br",
	MX: "auth.mercadolibre.com.mx",
	CL: "auth.mercadolibre.cl",
	CO: "auth.mercadolibre.com.co",
	UY: "auth.mercadolibre.com.uy",
	PE: "auth.mercadolibre.com.pe",
	VE: "auth.mercadolibre.com.ve",
	EC: "auth.mercadolibre.com.ec",
	BO: "auth.mercadolibre.com.bo",
	PY: "auth.mercadolibre.com.py",
	GT: "auth.mercadolibre.com.gt",
	PA: "auth.mercadolibre.com.pa",
	DO: "auth.mercadolibre.com.do",
	CR: "auth.mercadolibre.co.cr",
	HN: "auth.mercadolibre.com.hn",
	SV: "auth.mercadolibre.com.sv",
	NI: "auth.mercadolibre.com.ni",
}

const defaultAuthHost = "auth.mercadolibre.com"

// AuthHost returns the OAuth authorization host for the given site ID,
// falling back to the global mercadolibre.com host for unknown sites.
func AuthHost(site SiteID) string {
	if h, ok := authHosts[site]; ok {
		return h
	}
	return defaultAuthHost
}
