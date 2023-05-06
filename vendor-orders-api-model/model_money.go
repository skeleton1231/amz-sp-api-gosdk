/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// An amount of money, including units in the form of currency.
type Money struct {
	// Three digit currency code in ISO 4217 format. String of length 3.
	CurrencyCode string `json:"currencyCode,omitempty"`
	Amount string `json:"amount,omitempty"`
}