/*
 * Selling Partner API for Direct Fulfillment Shipping
 *
 * The Selling Partner API for Direct Fulfillment Shipping provides programmatic access to a direct fulfillment vendor's shipping data.
 *
 * API version: 2021-12-28
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Physical dimensional measurements of a container.
type Dimensions struct {
	Length string `json:"length"`
	Width string `json:"width"`
	Height string `json:"height"`
	// The unit of measure for dimensions.
	UnitOfMeasure string `json:"unitOfMeasure"`
}
