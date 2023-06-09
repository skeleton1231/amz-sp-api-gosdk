/*
 * Selling Partner API for Retail Procurement Shipments
 *
 * The Selling Partner API for Retail Procurement Shipments provides programmatic access to retail shipping data for vendors.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An amount of money, including units in the form of currency.
type Money struct {
	// Three digit currency code in ISO 4217 format.
	CurrencyCode string `json:"currencyCode"`
	Amount string `json:"amount"`
}
