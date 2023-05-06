/*
 * Selling Partner API for Pricing
 *
 * The Selling Partner API for Pricing helps you programmatically retrieve product pricing and offer information for Amazon Marketplace products.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The state and country from where the item is shipped.
type ShipsFromType struct {
	// The state from where the item is shipped.
	State string `json:"State,omitempty"`
	// The country from where the item is shipped.
	Country string `json:"Country,omitempty"`
}
