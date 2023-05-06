/*
 * Selling Partner APIs for Fulfillment Outbound
 *
 * The Selling Partner API for Fulfillment Outbound lets you create applications that help a seller fulfill Multi-Channel Fulfillment orders using their inventory in Amazon's fulfillment network. You can get information on both potential and existing fulfillment orders.
 *
 * API version: 2020-07-01
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// A physical address.
type Address struct {
	// The name of the person, business or institution at the address.
	Name string `json:"name"`
	// The first line of the address.
	AddressLine1 string `json:"addressLine1"`
	// Additional address information, if required.
	AddressLine2 string `json:"addressLine2,omitempty"`
	// Additional address information, if required.
	AddressLine3 string `json:"addressLine3,omitempty"`
	// The city where the person, business, or institution is located. This property is required in all countries except Japan. It should not be used in Japan.
	City string `json:"city,omitempty"`
	// The district or county where the person, business, or institution is located.
	DistrictOrCounty string `json:"districtOrCounty,omitempty"`
	// The state or region where the person, business or institution is located.
	StateOrRegion string `json:"stateOrRegion"`
	// The postal code of the address.
	PostalCode string `json:"postalCode"`
	// The two digit country code. In ISO 3166-1 alpha-2 format.
	CountryCode string `json:"countryCode"`
	// The phone number of the person, business, or institution located at the address.
	Phone string `json:"phone,omitempty"`
}
