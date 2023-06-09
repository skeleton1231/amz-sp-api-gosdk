/*
 * Selling Partner API for Shipping
 *
 * Provides programmatic access to Amazon Shipping APIs.   **Note:** If you are new to the Amazon Shipping API, refer to the latest version of <a href=\"https://developer-docs.amazon.com/amazon-shipping/docs/shipping-api-v2-reference\">Amazon Shipping API (v2)</a> on the <a href=\"https://developer-docs.amazon.com/amazon-shipping/\">Amazon Shipping Developer Documentation</a> site.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The address.
type Address struct {
	// The name of the person, business or institution at that address.
	Name string `json:"name"`
	// First line of that address.
	AddressLine1 string `json:"addressLine1"`
	// Additional address information, if required.
	AddressLine2 string `json:"addressLine2,omitempty"`
	// Additional address information, if required.
	AddressLine3 string `json:"addressLine3,omitempty"`
	StateOrRegion string `json:"stateOrRegion"`
	City string `json:"city"`
	CountryCode string `json:"countryCode"`
	PostalCode string `json:"postalCode"`
	// The email address of the contact associated with the address.
	Email string `json:"email,omitempty"`
	// The email cc addresses of the contact associated with the address.
	CopyEmails []string `json:"copyEmails,omitempty"`
	// The phone number of the person, business or institution located at that address.
	PhoneNumber string `json:"phoneNumber,omitempty"`
}
