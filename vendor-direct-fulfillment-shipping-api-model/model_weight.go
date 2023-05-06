/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// The weight.
type Weight struct {
	// The unit of measurement.
	UnitOfMeasure string `json:"unitOfMeasure"`
	Value string `json:"value"`
}
