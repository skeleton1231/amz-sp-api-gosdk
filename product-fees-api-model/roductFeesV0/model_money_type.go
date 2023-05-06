/*
 * Selling Partner API for Product Fees
 *
 * The Selling Partner API for Product Fees lets you programmatically retrieve estimated fees for a product. You can then account for those fees in your pricing.
 *
 * API version: v0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

type MoneyType struct {
	// The currency code in ISO 4217 format.
	CurrencyCode string `json:"CurrencyCode,omitempty"`
	// The monetary value.
	Amount float64 `json:"Amount,omitempty"`
}
