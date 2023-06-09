/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Address of the party.
type Address struct {
	// The name of the person, business or institution at that address.
	Name string `json:"name"`
	// First line of the address.
	AddressLine1 string `json:"addressLine1"`
	// Additional street address information, if required.
	AddressLine2 string `json:"addressLine2,omitempty"`
	// Additional street address information, if required.
	AddressLine3 string `json:"addressLine3,omitempty"`
	// The city where the person, business or institution is located.
	City string `json:"city,omitempty"`
	// The county where person, business or institution is located.
	County string `json:"county,omitempty"`
	// The district where person, business or institution is located.
	District string `json:"district,omitempty"`
	// The state or region where person, business or institution is located.
	StateOrRegion string `json:"stateOrRegion,omitempty"`
	// The postal code of that address. It contains a series of letters or digits or both, sometimes including spaces or punctuation.
	PostalCode string `json:"postalCode,omitempty"`
	// The two digit country code in ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode"`
	// The phone number of the person, business or institution located at that address.
	Phone string `json:"phone,omitempty"`
}
