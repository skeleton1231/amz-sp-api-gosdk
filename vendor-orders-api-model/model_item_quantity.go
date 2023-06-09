/*
 * Selling Partner API for Retail Procurement Orders
 *
 * The Selling Partner API for Retail Procurement Orders provides programmatic access to vendor orders data.
 *
 * API version: v1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Details of quantity ordered.
type ItemQuantity struct {
	// Acknowledged quantity. This value should not be zero.
	Amount int32 `json:"amount,omitempty"`
	// Unit of measure for the acknowledged quantity.
	UnitOfMeasure string `json:"unitOfMeasure,omitempty"`
	// The case size, in the event that we ordered using cases.
	UnitSize int32 `json:"unitSize,omitempty"`
}
